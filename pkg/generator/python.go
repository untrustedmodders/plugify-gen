package generator

import (
	"fmt"
	"strings"

	"github.com/untrustedmodders/plugify-gen/pkg/manifest"
)

// PythonGenerator generates Python stub (.pyi) files
type PythonGenerator struct {
	*BaseGenerator
}

// NewPythonGenerator creates a new Python generator
func NewPythonGenerator() *PythonGenerator {
	return &PythonGenerator{
		BaseGenerator: NewBaseGenerator("python", NewPythonTypeMapper(), PythonReservedWords),
	}
}

// Generate generates Python bindings
func (g *PythonGenerator) Generate(m *manifest.Manifest, opts *GeneratorOptions) (*GeneratorResult, error) {
	g.ResetCaches()
	m.Sanitize(g.Sanitizer)
	opts = EnsureOptions(opts)

	var sb strings.Builder

	// Check if we need Callable import
	needsCallable := g.needsCallable(m)

	// Header imports
	if needsCallable {
		sb.WriteString("from collections.abc import Callable\n")
	}
	sb.WriteString("from enum import IntEnum\n")
	sb.WriteString("from plugify.plugin import Vector2, Vector3, Vector4, Matrix4x4\n\n")
	sb.WriteString(fmt.Sprintf("# Generated from %s.pplugin\n\n", m.Name))

	// Generate enums
	enumsCode, err := g.generateEnums(m)
	if err != nil {
		return nil, err
	}
	if enumsCode != "" {
		sb.WriteString(enumsCode)
		sb.WriteString("\n")
	}
	// Generate aliases
	aliasesCode, err := g.generateAliases(m)
	if err != nil {
		return nil, err
	}
	if aliasesCode != "" {
		sb.WriteString(aliasesCode)
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
			fmt.Sprintf("pps/%s.pyi", m.Name): sb.String(),
		},
	}

	return result, nil
}

// needsCallable checks if any method uses function/delegate types
func (g *PythonGenerator) needsCallable(m *manifest.Manifest) bool {
	for _, method := range m.Methods {
		if method.RetType.Prototype != nil {
			return true
		}
		for _, param := range method.ParamTypes {
			if param.Prototype != nil {
				return true
			}
		}
	}
	return false
}

func (g *PythonGenerator) generateEnums(m *manifest.Manifest) (string, error) {
	// Use the base generator's CollectEnums helper
	return g.CollectEnums(m, g.generateEnum)
}

func (g *PythonGenerator) generateEnum(enum *manifest.Enum, underlyingType string) (string, error) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("class %s(IntEnum):\n", enum.Name))

	if enum.Description != "" {
		sb.WriteString(fmt.Sprintf("    \"\"\"\n    %s\n    \"\"\"\n", enum.Description))
	}

	for _, val := range enum.Values {
		if val.Description != "" {
			sb.WriteString(fmt.Sprintf("    # %s\n", val.Description))
		}
		sb.WriteString(fmt.Sprintf("    %s = %d\n", val.Name, val.Value))
	}

	return sb.String(), nil
}

func (g *PythonGenerator) generateAliases(m *manifest.Manifest) (string, error) {
	// Use the base generator's CollectAliases helper
	return g.CollectAliases(m, g.generateAlias)
}

func (g *PythonGenerator) generateAlias(alias *manifest.Alias, underlyingType string) (string, error) {
	var sb strings.Builder

	if alias.Description != "" {
		sb.WriteString(fmt.Sprintf("    \"\"\"\n    %s\n    \"\"\"\n", alias.Description))
	}

	sb.WriteString(fmt.Sprintf("    %s = %s\n", alias.Name, underlyingType))

	return sb.String(), nil
}

func (g *PythonGenerator) generateClasses(m *manifest.Manifest) (string, error) {
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

func (g *PythonGenerator) generateClass(m *manifest.Manifest, class *manifest.Class) (string, error) {
	var sb strings.Builder

	hasCtor := len(class.Constructors) > 0
	hasDtor := class.Destructor != nil

	// Add deprecation decorator if present
	if class.Deprecated != "" {
		sb.WriteString(fmt.Sprintf("@deprecated(reason=\"%s\")\n", class.Deprecated))
	}

	// Class declaration with docstring
	sb.WriteString(fmt.Sprintf("class %s:\n", class.Name))
	if class.Description != "" {
		sb.WriteString(fmt.Sprintf("    \"\"\"\n    %s\n    \"\"\"\n", class.Description))
	}

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
			sb.WriteString("    def __init__(self) -> None:\n")
			sb.WriteString("        ...\n\n")
		}

		// Main constructor if no constructors and destructors specified
		if !hasDtor {
			sb.WriteString("    def __init__(self, handle) -> None:\n")
			sb.WriteString("        ...\n\n")
		}
	}

	// Generate destructor and context manager methods for classes with destructors
	if hasDtor {
		dtorCode := g.generateDestructorMethods(class)
		sb.WriteString(dtorCode)
	}

	// Generate utility methods (get, release, reset, valid) for all classes
	utilCode, err := g.generateUtilityMethods(class)
	if err != nil {
		return "", err
	}
	sb.WriteString(utilCode)

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

func (g *PythonGenerator) generateDestructorMethods(class *manifest.Class) string {
	return fmt.Sprintf(`    def __del__(self) -> None:
        """
        Destructor. Releases the underlying handle if owned.
        """
        ...

    def close(self) -> None:
        """
        Close/destroy the handle if owned.
        """
        ...

    def __enter__(self) -> "%s":
        """
        Enter the runtime context for this object.

        Returns:
            %s: self
        """
        ...

    def __exit__(self, exc_type: type[BaseException] | None, exc_val: BaseException | None, exc_tb: object) -> None:
        """
        Exit the runtime context and release resources.

        Args:
            exc_type (type[BaseException] | None): Exception type if an exception was raised
            exc_val (BaseException | None): Exception value if an exception was raised
            exc_tb (object): Traceback if an exception was raised
        """
        ...

`, class.Name, class.Name)
}

func (g *PythonGenerator) generateUtilityMethods(class *manifest.Class) (string, error) {
	// Get the handle type mapped to Python
	_, handleType, err := g.typeMapper.MapHandleType(class)
	if err != nil {
		return "", err
	}

	methods := []struct {
		name        string
		returnType  string
		description string
		returnDesc  string
	}{
		{"get", handleType, "Get the raw handle value without transferring ownership.", "The underlying handle value"},
		{"release", handleType, "Release ownership of the handle and return it.", "The released handle value"},
		{"reset", "None", "Reset the handle by closing it.", ""},
		{"valid", "bool", "Check if the handle is valid.", "True if the handle is valid, False otherwise"},
	}

	var sb strings.Builder
	for _, method := range methods {
		sb.WriteString(fmt.Sprintf("    def %s(self) -> %s:\n", method.name, method.returnType))
		sb.WriteString("        \"\"\"\n")
		sb.WriteString(fmt.Sprintf("        %s\n", method.description))
		if method.returnDesc != "" {
			sb.WriteString(fmt.Sprintf("\n        Returns:\n            %s: %s\n", method.returnType, method.returnDesc))
		}
		sb.WriteString("        \"\"\"\n")
		sb.WriteString("        ...\n\n")
	}

	return sb.String(), nil
}

func (g *PythonGenerator) generateConstructor(m *manifest.Manifest, class *manifest.Class, methodName string) (string, error) {
	// Find the method in the manifest
	method := FindMethod(m, methodName)
	if method == nil {
		return "", fmt.Errorf("constructor method %s not found", methodName)
	}

	var sb strings.Builder

	// Add deprecation decorator if present
	if method.Deprecated != "" {
		sb.WriteString(fmt.Sprintf("    @deprecated(reason=\"%s\")\n", method.Deprecated))
	}

	// Generate __init__ signature
	params, err := g.formatParameters(method.ParamTypes)
	if err != nil {
		return "", err
	}

	if params != "" {
		sb.WriteString(fmt.Sprintf("    def __init__(self, %s) -> None:\n", params))
	} else {
		sb.WriteString("    def __init__(self) -> None:\n")
	}

	// Generate docstring
	if method.Description != "" {
		sb.WriteString(g.generateDocumentation(DocOptions{
			Description: method.Description,
			Params:      method.ParamTypes,
			Indent:      "    ",
		}))
	}

	// Stub body
	sb.WriteString("        ...\n")

	return sb.String(), nil
}

func (g *PythonGenerator) generateBinding(m *manifest.Manifest, class *manifest.Class, binding *manifest.Binding) (string, error) {
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

	// Format parameters
	formattedParams, err := g.formatParameters(methodParams)
	if err != nil {
		return "", err
	}

	// Apply parameter aliases (replace types with class names)
	if len(binding.ParamAliases) > 0 {
		formattedParams = g.applyParamAliases(formattedParams, methodParams, binding.ParamAliases)
	}

	// Generate return type
	retType, err := g.typeMapper.MapReturnType(&method.RetType)
	if err != nil {
		return "", err
	}

	// Handle return alias
	if binding.RetAlias != nil && binding.RetAlias.Name != "" {
		retType = binding.RetAlias.Name
	}

	// Add deprecation decorator if present (check both binding and underlying method)
	deprecationReason := binding.Deprecated
	if deprecationReason == "" {
		deprecationReason = method.Deprecated
	}
	if deprecationReason != "" {
		sb.WriteString(fmt.Sprintf("    @deprecated(reason=\"%s\")\n", deprecationReason))
	}

	// Determine if method is static
	if !binding.BindSelf {
		sb.WriteString("    @staticmethod\n")
		if formattedParams != "" {
			sb.WriteString(fmt.Sprintf("    def %s(%s) -> %s:\n", binding.Name, formattedParams, retType))
		} else {
			sb.WriteString(fmt.Sprintf("    def %s() -> %s:\n", binding.Name, retType))
		}
	} else {
		if formattedParams != "" {
			sb.WriteString(fmt.Sprintf("    def %s(self, %s) -> %s:\n", binding.Name, formattedParams, retType))
		} else {
			sb.WriteString(fmt.Sprintf("    def %s(self) -> %s:\n", binding.Name, retType))
		}
	}

	// Generate docstring
	if method.Description != "" {
		sb.WriteString(g.generateDocumentation(DocOptions{
			Description:  method.Description,
			Params:       methodParams,
			RetType:      method.RetType,
			ParamAliases: binding.ParamAliases,
			RetAlias:     binding.RetAlias,
			Indent:       "    ",
		}))
	}

	// Stub body
	sb.WriteString("        ...\n")

	return sb.String(), nil
}

func (g *PythonGenerator) applyParamAliases(formattedParams string, params []manifest.ParamType, aliases []*manifest.ParamAlias) string {
	result := formattedParams
	for i, alias := range aliases {
		if i < len(params) && alias.Name != "" {
			paramName := params[i].Name

			// Find and replace the type for this parameter
			// Look for "paramName: OldType" and replace with "paramName: NewType"
			oldPattern := paramName + ": "
			idx := strings.Index(result, oldPattern)
			if idx != -1 {
				// Find the end of the type (either comma or end of string)
				start := idx + len(oldPattern)
				end := start
				for end < len(result) && result[end] != ',' {
					end++
				}
				// Replace the type
				result = result[:start] + alias.Name + result[end:]
			}
		}
	}
	return result
}

func (g *PythonGenerator) generateMethod(method *manifest.Method) (string, error) {
	var sb strings.Builder

	// Add deprecation decorator if present
	if method.Deprecated != "" {
		sb.WriteString(fmt.Sprintf("@deprecated(reason=\"%s\")\n", method.Deprecated))
	}

	// Generate function signature
	params, err := g.formatParameters(method.ParamTypes)
	if err != nil {
		return "", err
	}

	// Generate return type (with tuple for ref parameters)
	retType, err := g.generateReturnType(&method.RetType, method.ParamTypes)
	if err != nil {
		return "", err
	}

	sb.WriteString(fmt.Sprintf("def %s(%s) -> %s:\n", method.Name, params, retType))

	// Generate docstring
	sb.WriteString(g.generateDocstring(method))

	// Stub body
	sb.WriteString("    ...\n")

	return sb.String(), nil
}

func (g *PythonGenerator) formatParameters(params []manifest.ParamType) (string, error) {
	if len(params) == 0 {
		return "", nil
	}

	result := ""
	for i, param := range params {
		if i > 0 {
			result += ", "
		}

		typeName, err := g.typeMapper.MapParamType(&param)
		if err != nil {
			return "", err
		}

		paramName := param.Name

		result += paramName + ": " + typeName
	}

	return result, nil
}

func (g *PythonGenerator) generateReturnType(retType *manifest.RetType, params []manifest.ParamType) (string, error) {
	// Check if any parameters are ref (out parameters)
	hasRef := false
	for _, param := range params {
		if param.Ref {
			hasRef = true
			break
		}
	}

	baseRetType, err := g.typeMapper.MapReturnType(retType)
	if err != nil {
		return "", err
	}

	if !hasRef {
		return baseRetType, nil
	}

	// Build tuple with return value + ref parameters
	types := []string{baseRetType}
	for _, param := range params {
		if param.Ref {
			paramType, err := g.typeMapper.MapParamType(&param)
			if err != nil {
				return "", err
			}
			types = append(types, paramType)
		}
	}

	return fmt.Sprintf("tuple[%s]", strings.Join(types, ", ")), nil
}

// generateDocumentation generates a Python docstring with Args and Returns sections
func (g *PythonGenerator) generateDocumentation(opts DocOptions) string {
	var sb strings.Builder

	sb.WriteString(opts.Indent + "    \"\"\"\n")
	if opts.Description != "" {
		sb.WriteString(fmt.Sprintf("%s    %s\n", opts.Indent, opts.Description))
	}

	// Parameters section
	if len(opts.Params) > 0 {
		sb.WriteString(fmt.Sprintf("\n%s    Args:\n", opts.Indent))
		for i, param := range opts.Params {
			paramName := param.Name
			paramType := param.Type

			// Apply parameter alias if provided
			if i < len(opts.ParamAliases) && opts.ParamAliases[i] != nil {
				paramType = opts.ParamAliases[i].Name
			}

			sb.WriteString(fmt.Sprintf("%s        %s (%s): %s\n",
				opts.Indent, paramName, paramType, param.Description))
		}
	}

	// Return type section
	if opts.RetType.Type != "void" {
		returnType := opts.RetType.Type

		// Apply return alias if provided
		if opts.RetAlias != nil && opts.RetAlias.Name != "" {
			returnType = opts.RetAlias.Name
		}

		sb.WriteString(fmt.Sprintf("\n%s    Returns:\n%s        %s: %s\n",
			opts.Indent, opts.Indent, returnType, opts.RetType.Description))
	}

	// Callback prototypes
	if opts.IncludeCallbacks {
		for _, param := range opts.Params {
			if param.Prototype != nil {
				sb.WriteString(fmt.Sprintf("\n%s    Callback Prototype (%s):\n", opts.Indent, param.Prototype.Name))
				if param.Prototype.Description != "" {
					sb.WriteString(fmt.Sprintf("%s        %s\n\n", opts.Indent, param.Prototype.Description))
				}
				if len(param.Prototype.ParamTypes) > 0 {
					sb.WriteString(fmt.Sprintf("%s        Args:\n", opts.Indent))
					for _, protoParam := range param.Prototype.ParamTypes {
						sb.WriteString(fmt.Sprintf("%s            %s (%s): %s\n",
							opts.Indent, protoParam.Name, protoParam.Type, protoParam.Description))
					}
				}
				if param.Prototype.RetType.Type != "void" {
					sb.WriteString(fmt.Sprintf("\n%s        Returns:\n%s            %s: %s\n",
						opts.Indent, opts.Indent, param.Prototype.RetType.Type, param.Prototype.RetType.Description))
				}
			}
		}
	}

	sb.WriteString(opts.Indent + "    \"\"\"\n")
	return sb.String()
}

func (g *PythonGenerator) generateDocstring(method *manifest.Method) string {
	return g.generateDocumentation(DocOptions{
		Description:      method.Description,
		Deprecated:       method.Deprecated,
		Params:           method.ParamTypes,
		RetType:          method.RetType,
		IncludeCallbacks: true,
		Indent:           "",
	})
}

// PythonTypeMapper implements type mapping for Python
type PythonTypeMapper struct{}

func NewPythonTypeMapper() *PythonTypeMapper {
	return &PythonTypeMapper{}
}

func (m *PythonTypeMapper) MapType(baseType string, context TypeContext, isArray bool) (string, error) {
	typeMap := map[string]string{
		"void":   "None",
		"bool":   "bool",
		"char8":  "str",
		"char16": "str",
		"int8":   "int",
		"int16":  "int",
		"int32":  "int",
		"int64":  "int",
		"uint8":  "int",
		"uint16": "int",
		"uint32": "int",
		"uint64": "int",
		"ptr64":  "int",
		"float":  "float",
		"double": "float",
		"string": "str",
		"any":    "object",
		"vec2":   "Vector2",
		"vec3":   "Vector3",
		"vec4":   "Vector4",
		"mat4x4": "Matrix4x4",
	}

	mapped, ok := typeMap[baseType]
	if !ok {
		// Assume it's a custom type
		mapped = baseType
	}

	if isArray && context&TypeContextAlias == 0 {
		mapped = fmt.Sprintf("list[%s]", mapped)
	}

	return mapped, nil
}

func (m *PythonTypeMapper) MapParamType(param *manifest.ParamType) (string, error) {
	ctx := TypeContextValue

	var typeName string
	switch {
	case param.Alias != nil:
		typeName = param.Alias.Name
		if !param.Alias.Element {
			ctx |= TypeContextAlias
		}

	case param.Enum != nil:
		typeName = param.Enum.Name

	case param.Prototype != nil:
		return m.generateCallableType(param.Prototype)

	default:
		typeName = param.BaseType()
	}

	return m.MapType(typeName, ctx, param.IsArray())
}

func (m *PythonTypeMapper) MapReturnType(retType *manifest.RetType) (string, error) {
	ctx := TypeContextReturn

	var typeName string
	switch {
	case retType.Alias != nil:
		typeName = retType.Alias.Name
		if !retType.Alias.Element {
			ctx |= TypeContextAlias
		}

	case retType.Enum != nil:
		typeName = retType.Enum.Name

	case retType.Prototype != nil:
		return m.generateCallableType(retType.Prototype)

	default:
		typeName = retType.BaseType()
	}

	return m.MapType(typeName, ctx, retType.IsArray())
}

func (m *PythonTypeMapper) generateCallableType(proto *manifest.Prototype) (string, error) {
	// Generate parameter types
	var paramTypes []string
	for _, param := range proto.ParamTypes {
		pType, err := m.MapParamType(&param)
		if err != nil {
			return "", err
		}
		paramTypes = append(paramTypes, pType)
	}

	// Generate return type
	retType, err := m.MapReturnType(&proto.RetType)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Callable[[%s], %s]", strings.Join(paramTypes, ", "), retType), nil
}

func (m *PythonTypeMapper) MapHandleType(class *manifest.Class) (string, string, error) {
	handleType, err := m.MapType(class.HandleType, TypeContextReturn, false)
	if err != nil {
		return "", "", err
	}
	return "None", handleType, nil
}
