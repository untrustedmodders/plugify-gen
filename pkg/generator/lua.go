package generator

import (
	"fmt"
	"strings"

	"github.com/untrustedmodders/plugify-gen/pkg/manifest"
)

// LuaGenerator generates Lua stub files
type LuaGenerator struct {
	*BaseGenerator
}

// NewLuaGenerator creates a new Lua generator
func NewLuaGenerator() *LuaGenerator {
	return &LuaGenerator{
		BaseGenerator: NewBaseGenerator("lua", NewLuaTypeMapper(), LuaReservedWords),
	}
}

// Generate generates Lua bindings
func (g *LuaGenerator) Generate(m *manifest.Manifest, opts *GeneratorOptions) (*GeneratorResult, error) {
	g.ResetCaches()
	m.Sanitize(g.Sanitizer)
	opts = EnsureOptions(opts)

	var sb strings.Builder

	// Header comment
	sb.WriteString(fmt.Sprintf("-- Generated from %s.pplugin\n\n", m.Name))

	// Generate enums
	enumsCode, err := g.generateEnums(m)
	if err != nil {
		return nil, err
	}
	if enumsCode != "" {
		sb.WriteString(enumsCode)
		sb.WriteString("\n")
	}

	// Generate methods
	for _, method := range m.Methods {
		methodCode, err := g.generateMethod(&method)
		if err != nil {
			return nil, fmt.Errorf("failed to generate method %s: %w", method.Name, err)
		}
		sb.WriteString(methodCode)
		sb.WriteString("\n")
	}

	// Generate classes (if enabled)
	if opts.GenerateClasses && len(m.Classes) > 0 {
		classesCode, err := g.generateClasses(m)
		if err != nil {
			return nil, fmt.Errorf("generating classes: %w", err)
		}
		sb.WriteString(classesCode)
		sb.WriteString("\n")
	}

	result := &GeneratorResult{
		Files: map[string]string{
			fmt.Sprintf("pps/%s.lua", m.Name): sb.String(),
		},
	}

	return result, nil
}

func (g *LuaGenerator) generateEnums(m *manifest.Manifest) (string, error) {
	// Use the base generator's CollectEnums helper
	return g.CollectEnums(m, g.generateEnum)
}

func (g *LuaGenerator) generateEnum(enum *manifest.Enum, underlyingType string) (string, error) {
	var sb strings.Builder

	// Enum comment
	if enum.Description != "" {
		sb.WriteString(fmt.Sprintf("-- %s: %s\n", enum.Name, enum.Description))
	} else {
		sb.WriteString(fmt.Sprintf("-- Enum: %s\n", enum.Name))
	}

	// Lua table
	sb.WriteString(fmt.Sprintf("%s = {\n", enum.Name))

	for _, val := range enum.Values {
		if val.Description != "" {
			sb.WriteString(fmt.Sprintf("  -- %s\n", val.Description))
		}
		sb.WriteString(fmt.Sprintf("  %s = %d,\n", val.Name, val.Value))
	}

	sb.WriteString("}\n")

	return sb.String(), nil
}

func (g *LuaGenerator) generateClasses(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	for i, class := range m.Classes {
		if i > 0 {
			sb.WriteString("\n")
		}
		classCode, err := g.generateClass(m, &class)
		if err != nil {
			return "", fmt.Errorf("failed to generate class %s: %w", class.Name, err)
		}
		sb.WriteString(classCode)
	}

	return sb.String(), nil
}

func (g *LuaGenerator) generateClass(m *manifest.Manifest, class *manifest.Class) (string, error) {
	var sb strings.Builder

	hasCtor := len(class.Constructors) > 0
	hasDtor := class.Destructor != nil

	// Class comment
	sb.WriteString(g.formatDescriptionComment(class.Description, fmt.Sprintf("Class: %s", class.Name)))

	// Class table declaration
	sb.WriteString(fmt.Sprintf("%s = {}\n\n", class.Name))

	// Generate constructors
	if hasCtor {
		for _, ctorName := range class.Constructors {
			ctorCode, err := g.generateConstructor(m, class, ctorName)
			if err != nil {
				return "", err
			}
			sb.WriteString(ctorCode)
			sb.WriteString("\n")
		}
	} else {
		// Default constructor if no constructors specified
		hasDefaultConstructor := g.HasConstructorWithNoParam(m, class)
		if !hasDefaultConstructor {
			sb.WriteString(fmt.Sprintf("--- Constructor for %s\n", class.Name))
			sb.WriteString(fmt.Sprintf("function %s.new() end\n\n", class.Name))
		}

		// Main constructor if no constructors and destructors specified
		if !hasDtor {
			sb.WriteString(fmt.Sprintf("--- Constructor for %s from existing handle\n", class.Name))
			sb.WriteString(fmt.Sprintf("function %s.new(handle) end\n\n", class.Name))
		}
	}

	// Generate utility methods (valid, get, release, and close if destructor exists)
	sb.WriteString(g.generateUtilityMethods(class))

	// Generate bindings (methods)
	for _, binding := range class.Bindings {
		methodCode, err := g.generateBinding(m, class, &binding)
		if err != nil {
			return "", err
		}
		sb.WriteString(methodCode)
		sb.WriteString("\n")
	}

	return sb.String(), nil
}

func (g *LuaGenerator) generateUtilityMethods(class *manifest.Class) string {
	var sb strings.Builder

	methods := []struct {
		name        string
		description string
		returnType  string
		returnDesc  string
	}{
		{"valid", "Check if the handle is valid.", "boolean", "True if the handle is valid, false otherwise"},
		{"get", "Get the raw handle value without transferring ownership.", class.HandleType, "The underlying handle value"},
		{"release", "Release ownership of the handle and return it.", class.HandleType, "The released handle value"},
		{"reset", "Reset the handle by closing it.", "", ""},
	}

	for _, method := range methods {
		sb.WriteString(fmt.Sprintf("--- %s\n", method.description))
		if method.returnType != "" {
			sb.WriteString(fmt.Sprintf("-- @return %s %s\n", method.returnType, method.returnDesc))
		}
		sb.WriteString(fmt.Sprintf("function %s:%s() end\n\n", class.Name, method.name))
	}

	// close() method - only if destructor exists
	if class.Destructor != nil {
		sb.WriteString("--- Close and destroy the handle if owned.\n")
		sb.WriteString(fmt.Sprintf("function %s:close() end\n\n", class.Name))
	}

	return sb.String()
}

func (g *LuaGenerator) generateConstructor(m *manifest.Manifest, class *manifest.Class, methodName string) (string, error) {
	// Find the method in the manifest
	method := FindMethod(m, methodName)
	if method == nil {
		return "", fmt.Errorf("constructor method %s not found", methodName)
	}

	var sb strings.Builder

	// Constructor documentation
	classRetType := manifest.RetType{Type: class.Name, Description: ""}
	sb.WriteString(g.generateLuaDocumentation(LuaDocOptions{
		Description:  method.Description,
		FallbackName: fmt.Sprintf("Constructor for %s", class.Name),
		Params:       method.ParamTypes,
		RetType:      &classRetType,
	}))

	// Add deprecation annotation if present
	if method.Deprecated != "" {
		sb.WriteString(fmt.Sprintf("@[deprecated {reason = \"%s\"}]\n", method.Deprecated))
	}

	// Constructor signature (using .new convention)
	params := g.formatParameters(method.ParamTypes)
	sb.WriteString(fmt.Sprintf("function %s.new(%s) end\n", class.Name, params))

	return sb.String(), nil
}

func (g *LuaGenerator) generateBinding(m *manifest.Manifest, class *manifest.Class, binding *manifest.Binding) (string, error) {
	// Find the method in the manifest
	method := FindMethod(m, binding.Method)
	if method == nil {
		return "", fmt.Errorf("method %s not found", binding.Method)
	}

	var sb strings.Builder

	// Determine parameters (skip first if bindSelf)
	params := method.ParamTypes
	startIdx := 0
	if binding.BindSelf && len(params) > 0 {
		startIdx = 1
	}
	methodParams := params[startIdx:]

	// Method documentation
	sb.WriteString(g.generateLuaDocumentation(LuaDocOptions{
		Description:  method.Description,
		FallbackName: binding.Name,
		Params:       methodParams,
		ParamAliases: binding.ParamAliases,
		RetType:      &method.RetType,
		RetAlias:     binding.RetAlias,
	}))

	// Add deprecation annotation if present (check both binding and underlying method)
	deprecationReason := binding.Deprecated
	if deprecationReason == "" {
		deprecationReason = method.Deprecated
	}
	if deprecationReason != "" {
		sb.WriteString(fmt.Sprintf("@[deprecated {reason = \"%s\"}]\n", deprecationReason))
	}

	// Method signature
	formattedParams := g.formatParameters(methodParams)

	// Determine if method is static or instance
	if !binding.BindSelf {
		// Static method: ClassName.method()
		sb.WriteString(fmt.Sprintf("function %s.%s(%s) end\n", class.Name, binding.Name, formattedParams))
	} else {
		// Instance method: ClassName:method() (colon notation includes implicit self)
		sb.WriteString(fmt.Sprintf("function %s:%s(%s) end\n", class.Name, binding.Name, formattedParams))
	}

	return sb.String(), nil
}

func (g *LuaGenerator) generateMethod(method *manifest.Method) (string, error) {
	var sb strings.Builder

	// Generate LDoc-style documentation
	sb.WriteString(g.generateDocumentation(method))

	// Add deprecation annotation if present
	if method.Deprecated != "" {
		sb.WriteString(fmt.Sprintf("@[deprecated {reason = \"%s\"}]\n", method.Deprecated))
	}

	// Generate function signature
	params := g.formatParameters(method.ParamTypes)
	sb.WriteString(fmt.Sprintf("function %s(%s) end\n", method.Name, params))

	return sb.String(), nil
}

func (g *LuaGenerator) formatParameters(params []manifest.ParamType) string {
	if len(params) == 0 {
		return ""
	}

	result := ""
	for i, param := range params {
		if i > 0 {
			result += ", "
		}

		paramName := param.Name
		result += paramName
	}

	return result
}

// LuaDocOptions configures LDoc-style documentation generation
type LuaDocOptions struct {
	Description      string
	FallbackName     string
	Params           []manifest.ParamType
	ParamAliases     []*manifest.ParamAlias
	RetType          *manifest.RetType
	RetAlias         *manifest.RetAlias
	IncludeCallbacks bool
}

// formatDescriptionComment generates description comment with fallback
func (g *LuaGenerator) formatDescriptionComment(description, fallbackName string) string {
	if description != "" {
		return fmt.Sprintf("--- %s\n", description)
	}
	return fmt.Sprintf("--- %s\n", fallbackName)
}

// generateLuaDocumentation generates LDoc-style documentation for methods
func (g *LuaGenerator) generateLuaDocumentation(opts LuaDocOptions) string {
	var sb strings.Builder

	// Main description
	sb.WriteString(g.formatDescriptionComment(opts.Description, opts.FallbackName))

	// Parameters
	for i, param := range opts.Params {
		paramName := param.Name
		paramType := param.Type

		// Apply parameter alias if provided
		if i < len(opts.ParamAliases) && opts.ParamAliases[i] != nil && opts.ParamAliases[i].Name != "" {
			paramType = opts.ParamAliases[i].Name
		}

		sb.WriteString(fmt.Sprintf("-- @param %s %s %s\n",
			paramName, paramType, param.Description))
	}

	// Return type
	if opts.RetType != nil && opts.RetType.Type != "void" {
		retType := opts.RetType.Type
		retDesc := opts.RetType.Description

		// Apply return alias if provided
		if opts.RetAlias != nil && opts.RetAlias.Name != "" {
			retType = opts.RetAlias.Name
		}

		sb.WriteString(fmt.Sprintf("-- @return %s %s\n", retType, retDesc))
	}

	// Callback prototypes
	if opts.IncludeCallbacks {
		for _, param := range opts.Params {
			if param.Prototype != nil {
				proto := param.Prototype
				sb.WriteString(fmt.Sprintf("-- @callback %s %s - %s\n",
					proto.Name, param.Name, proto.Description))

				// Callback parameters
				for _, protoParam := range proto.ParamTypes {
					sb.WriteString(fmt.Sprintf("-- @param %s %s %s\n",
						protoParam.Name, protoParam.Type, protoParam.Description))
				}

				// Callback return
				if proto.RetType.Type != "void" {
					sb.WriteString(fmt.Sprintf("-- @return %s %s\n",
						proto.RetType.Type, proto.RetType.Description))
				}
			}
		}
	}

	return sb.String()
}

func (g *LuaGenerator) generateDocumentation(method *manifest.Method) string {
	return g.generateLuaDocumentation(LuaDocOptions{
		Description:      method.Description,
		FallbackName:     method.Name,
		Params:           method.ParamTypes,
		RetType:          &method.RetType,
		IncludeCallbacks: true,
	})
}

// LuaTypeMapper is a placeholder - Lua doesn't need explicit type mapping in stubs
type LuaTypeMapper struct{}

func NewLuaTypeMapper() *LuaTypeMapper {
	return &LuaTypeMapper{}
}

func (m *LuaTypeMapper) MapType(baseType string, context TypeContext, isArray bool) (string, error) {
	// Lua is dynamically typed, return empty string (not used)
	return "", nil
}

func (m *LuaTypeMapper) MapParamType(param *manifest.ParamType) (string, error) {
	return "", nil
}

func (m *LuaTypeMapper) MapReturnType(retType *manifest.RetType) (string, error) {
	return "", nil
}

func (m *LuaTypeMapper) MapHandleType(class *manifest.Class) (string, string, error) {
	return "", "", nil
}
