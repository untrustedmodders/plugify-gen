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
	invalidNames := []string{
		"and", "break", "do", "else", "elseif", "end", "false",
		"for", "function", "goto", "if", "in", "local", "nil",
		"not", "or", "repeat", "return", "then", "true", "until", "while",
	}

	return &LuaGenerator{
		BaseGenerator: NewBaseGenerator("lua", NewLuaTypeMapper(), invalidNames),
	}
}

// Generate generates Lua bindings
func (g *LuaGenerator) Generate(m *manifest.Manifest) (*GeneratorResult, error) {
	g.ResetCaches()

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

	// Generate classes
	if len(m.Classes) > 0 {
		classesCode, err := g.generateClasses(m)
		if err != nil {
			return nil, err
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
	var sb strings.Builder

	for _, method := range m.Methods {
		// Check return type
		if method.RetType.Enum != nil && !g.IsEnumCached(method.RetType.Enum.Name) {
			enumCode := g.generateEnum(method.RetType.Enum)
			sb.WriteString(enumCode)
			sb.WriteString("\n")
			g.CacheEnum(method.RetType.Enum.Name)
		}

		// Check parameters
		for _, param := range method.ParamTypes {
			if param.Enum != nil && !g.IsEnumCached(param.Enum.Name) {
				enumCode := g.generateEnum(param.Enum)
				sb.WriteString(enumCode)
				sb.WriteString("\n")
				g.CacheEnum(param.Enum.Name)
			}

			// Check nested prototypes
			if param.Prototype != nil {
				g.processPrototypeEnums(param.Prototype, &sb)
			}
		}

		// Check return type prototype
		if method.RetType.Prototype != nil {
			g.processPrototypeEnums(method.RetType.Prototype, &sb)
		}
	}

	return sb.String(), nil
}

func (g *LuaGenerator) processPrototypeEnums(proto *manifest.Prototype, sb *strings.Builder) {
	if proto.RetType.Enum != nil && !g.IsEnumCached(proto.RetType.Enum.Name) {
		sb.WriteString(g.generateEnum(proto.RetType.Enum))
		sb.WriteString("\n")
		g.CacheEnum(proto.RetType.Enum.Name)
	}

	for _, param := range proto.ParamTypes {
		if param.Enum != nil && !g.IsEnumCached(param.Enum.Name) {
			sb.WriteString(g.generateEnum(param.Enum))
			sb.WriteString("\n")
			g.CacheEnum(param.Enum.Name)
		}
		if param.Prototype != nil {
			g.processPrototypeEnums(param.Prototype, sb)
		}
	}
}

func (g *LuaGenerator) generateEnum(enum *manifest.EnumType) string {
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

	return sb.String()
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
	if class.Description != "" {
		sb.WriteString(fmt.Sprintf("--- %s\n", class.Description))
	} else {
		sb.WriteString(fmt.Sprintf("--- Class: %s\n", class.Name))
	}

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
		sb.WriteString(fmt.Sprintf("--- Constructor for %s\n", class.Name))
		sb.WriteString(fmt.Sprintf("function %s.new() end\n\n", class.Name))

		// Main constructor if no constructors and destructors specified
		if !hasDtor {
			sb.WriteString(fmt.Sprintf("--- Constructor for %s from existing handle\n", class.Name))
			sb.WriteString(fmt.Sprintf("function %s.new(handle) end\n\n", class.Name))
		}
	}

	// Generate utility methods (valid, get, release, and close if destructor exists)
	sb.WriteString(g.generateUtilityMethods(class, hasDtor))

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

func (g *LuaGenerator) generateUtilityMethods(class *manifest.Class, hasDtor bool) string {
	var sb strings.Builder

	// valid() method
	sb.WriteString("--- Check if the handle is valid.\n")
	sb.WriteString("-- @return boolean True if the handle is valid, false otherwise\n")
	sb.WriteString(fmt.Sprintf("function %s:valid() end\n\n", class.Name))

	// get() method
	sb.WriteString("--- Get the raw handle value without transferring ownership.\n")
	sb.WriteString(fmt.Sprintf("-- @return %s The underlying handle value\n", class.HandleType))
	sb.WriteString(fmt.Sprintf("function %s:get() end\n\n", class.Name))

	// release() method
	sb.WriteString("--- Release ownership of the handle and return it.\n")
	sb.WriteString(fmt.Sprintf("-- @return %s The released handle value\n", class.HandleType))
	sb.WriteString(fmt.Sprintf("function %s:release() end\n\n", class.Name))

	// reset() method
	sb.WriteString("--- Reset the handle by closing it.\n")
	sb.WriteString(fmt.Sprintf("function %s:reset() end\n\n", class.Name))

	// close() method - only if destructor exists
	if hasDtor {
		sb.WriteString("--- Close and destroy the handle if owned.\n")
		sb.WriteString(fmt.Sprintf("function %s:close() end\n\n", class.Name))
	}

	return sb.String()
}

func (g *LuaGenerator) generateConstructor(m *manifest.Manifest, class *manifest.Class, methodName string) (string, error) {
	// Find the method in the manifest
	var method *manifest.Method
	for i := range m.Methods {
		if m.Methods[i].Name == methodName || m.Methods[i].FuncName == methodName {
			method = &m.Methods[i]
			break
		}
	}
	if method == nil {
		return "", fmt.Errorf("constructor method %s not found", methodName)
	}

	var sb strings.Builder

	// Constructor documentation
	if method.Description != "" {
		sb.WriteString(fmt.Sprintf("--- %s\n", method.Description))
	} else {
		sb.WriteString(fmt.Sprintf("--- Constructor for %s\n", class.Name))
	}

	// Parameters
	for _, param := range method.ParamTypes {
		paramName := param.Name
		if paramName == "" {
			paramName = "unnamed"
		}
		sb.WriteString(fmt.Sprintf("-- @param %s %s %s\n",
			g.SanitizeName(paramName), param.Type, param.Description))
	}

	// Return type (the class instance)
	sb.WriteString(fmt.Sprintf("-- @return %s\n", class.Name))

	// Constructor signature (using .new convention)
	params := g.formatParameters(method.ParamTypes)
	sb.WriteString(fmt.Sprintf("function %s.new(%s) end\n", class.Name, params))

	return sb.String(), nil
}

func (g *LuaGenerator) generateBinding(m *manifest.Manifest, class *manifest.Class, binding *manifest.Binding) (string, error) {
	// Find the method in the manifest
	var method *manifest.Method
	for i := range m.Methods {
		if m.Methods[i].Name == binding.Method || m.Methods[i].FuncName == binding.Method {
			method = &m.Methods[i]
			break
		}
	}
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
	if method.Description != "" {
		sb.WriteString(fmt.Sprintf("--- %s\n", method.Description))
	} else {
		sb.WriteString(fmt.Sprintf("--- %s\n", binding.Name))
	}

	// Parameters (excluding self if bindSelf)
	for _, param := range methodParams {
		paramName := param.Name
		if paramName == "" {
			paramName = "unnamed"
		}
		// Check if this parameter has an alias
		paramType := param.Type
		for i, alias := range binding.ParamAliases {
			if i < len(methodParams) && methodParams[i].Name == param.Name && alias.Name != "" {
				paramType = alias.Name
				break
			}
		}
		sb.WriteString(fmt.Sprintf("-- @param %s %s %s\n",
			g.SanitizeName(paramName), paramType, param.Description))
	}

	// Return type
	if method.RetType.Type != "void" {
		retType := method.RetType.Type
		// Handle return alias
		if binding.RetAlias != nil && binding.RetAlias.Name != "" {
			retType = binding.RetAlias.Name
		}
		sb.WriteString(fmt.Sprintf("-- @return %s %s\n", retType, method.RetType.Description))
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

	// Generate function signature
	params := g.formatParameters(method.ParamTypes)
	sb.WriteString(fmt.Sprintf("function %s(%s) end\n", g.SanitizeName(method.Name), params))

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
		if paramName == "" {
			paramName = fmt.Sprintf("p%d", i)
		}

		result += g.SanitizeName(paramName)
	}

	return result
}

func (g *LuaGenerator) generateDocumentation(method *manifest.Method) string {
	var sb strings.Builder

	// Main description
	if method.Description != "" {
		sb.WriteString(fmt.Sprintf("--- %s\n", method.Description))
	} else {
		sb.WriteString(fmt.Sprintf("--- %s\n", method.Name))
	}

	// Parameters
	for _, param := range method.ParamTypes {
		paramName := param.Name
		if paramName == "" {
			paramName = "unnamed"
		}
		paramType := param.Type
		paramDesc := param.Description

		sb.WriteString(fmt.Sprintf("-- @param %s %s %s\n",
			g.SanitizeName(paramName), paramType, paramDesc))
	}

	// Return type
	if method.RetType.Type != "void" {
		retDesc := method.RetType.Description
		sb.WriteString(fmt.Sprintf("-- @return %s %s\n", method.RetType.Type, retDesc))
	}

	// Callback prototypes
	for _, param := range method.ParamTypes {
		if param.Prototype != nil {
			proto := param.Prototype
			protoName := proto.Name
			protoDesc := proto.Description
			if protoDesc == "" {
				protoDesc = "Callback function"
			}

			sb.WriteString(fmt.Sprintf("-- @callback %s %s - %s\n",
				protoName, g.SanitizeName(param.Name), protoDesc))

			// Callback parameters
			for _, protoParam := range proto.ParamTypes {
				sb.WriteString(fmt.Sprintf("-- @param %s %s %s\n",
					g.SanitizeName(protoParam.Name), protoParam.Type, protoParam.Description))
			}

			// Callback return
			if proto.RetType.Type != "void" {
				sb.WriteString(fmt.Sprintf("-- @return %s %s\n",
					proto.RetType.Type, proto.RetType.Description))
			}
		}
	}

	return sb.String()
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

func (m *LuaTypeMapper) MapParamType(param *manifest.ParamType, context TypeContext) (string, error) {
	return "", nil
}

func (m *LuaTypeMapper) MapReturnType(retType *manifest.TypeInfo) (string, error) {
	return "", nil
}

func (m *LuaTypeMapper) MapHandleType(class *manifest.Class) (string, string) {
	return "", ""
}
