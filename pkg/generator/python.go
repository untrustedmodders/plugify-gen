package generator

import (
	"fmt"
	"strings"

	"github.com/untrustedmodders/plugify-generator/pkg/manifest"
)

// PythonGenerator generates Python stub (.pyi) files
type PythonGenerator struct {
	*BaseGenerator
}

// NewPythonGenerator creates a new Python generator
func NewPythonGenerator() *PythonGenerator {
	invalidNames := []string{
		"False", "await", "else", "import", "pass",
		"None", "break", "except", "in", "raise",
		"True", "class", "finally", "is", "return",
		"and", "continue", "for", "lambda", "try",
		"as", "def", "from", "nonlocal", "while",
		"assert", "del", "global", "not", "with",
		"async", "elif", "if", "or", "yield",
	}

	return &PythonGenerator{
		BaseGenerator: NewBaseGenerator("python", NewPythonTypeMapper(), invalidNames),
	}
}

// Generate generates Python bindings
func (g *PythonGenerator) Generate(m *manifest.Manifest) (*GeneratorResult, error) {
	g.ResetCaches()

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

	// Generate methods
	for _, method := range m.Methods {
		methodCode, err := g.generateMethod(&method)
		if err != nil {
			return nil, fmt.Errorf("failed to generate method %s: %w", method.Name, err)
		}
		sb.WriteString(methodCode)
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

func (g *PythonGenerator) processPrototypeEnums(proto *manifest.Prototype, sb *strings.Builder) {
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

func (g *PythonGenerator) generateEnum(enum *manifest.EnumType) string {
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

	return sb.String()
}

func (g *PythonGenerator) generateMethod(method *manifest.Method) (string, error) {
	var sb strings.Builder

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

	sb.WriteString(fmt.Sprintf("def %s(%s) -> %s:\n", g.SanitizeName(method.Name), params, retType))

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

		typeName, err := g.typeMapper.MapParamType(&param, TypeContextValue)
		if err != nil {
			return "", err
		}

		paramName := param.Name
		if paramName == "" {
			paramName = fmt.Sprintf("p%d", i)
		}

		result += g.SanitizeName(paramName) + ": " + typeName
	}

	return result, nil
}

func (g *PythonGenerator) generateReturnType(retType *manifest.TypeInfo, params []manifest.ParamType) (string, error) {
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
			paramType, err := g.typeMapper.MapParamType(&param, TypeContextValue)
			if err != nil {
				return "", err
			}
			types = append(types, paramType)
		}
	}

	return fmt.Sprintf("tuple[%s]", strings.Join(types, ", ")), nil
}

func (g *PythonGenerator) generateDocstring(method *manifest.Method) string {
	var sb strings.Builder

	sb.WriteString("    \"\"\"\n")
	if method.Description != "" {
		sb.WriteString(fmt.Sprintf("    %s\n\n", method.Description))
	}

	if len(method.ParamTypes) > 0 {
		sb.WriteString("    Args:\n")
		for _, param := range method.ParamTypes {
			paramName := param.Name
			if paramName == "" {
				paramName = "unnamed"
			}
			sb.WriteString(fmt.Sprintf("        %s (%s): %s\n",
				g.SanitizeName(paramName), param.Type, param.Description))
		}
	}

	if method.RetType.Type != "void" {
		sb.WriteString(fmt.Sprintf("\n    Returns:\n        %s: %s\n",
			method.RetType.Type, method.RetType.Description))
	}

	// Add callback prototypes
	for _, param := range method.ParamTypes {
		if param.Prototype != nil {
			sb.WriteString(fmt.Sprintf("\n    Callback Prototype (%s):\n", param.Prototype.Name))
			if param.Prototype.Description != "" {
				sb.WriteString(fmt.Sprintf("        %s\n\n", param.Prototype.Description))
			}
			if len(param.Prototype.ParamTypes) > 0 {
				sb.WriteString("        Args:\n")
				for _, protoParam := range param.Prototype.ParamTypes {
					sb.WriteString(fmt.Sprintf("            %s (%s): %s\n",
						protoParam.Name, protoParam.Type, protoParam.Description))
				}
			}
			if param.Prototype.RetType.Type != "void" {
				sb.WriteString(fmt.Sprintf("\n        Returns:\n            %s: %s\n",
					param.Prototype.RetType.Type, param.Prototype.RetType.Description))
			}
		}
	}

	sb.WriteString("    \"\"\"\n")
	return sb.String()
}

// PythonTypeMapper implements type mapping for Python
type PythonTypeMapper struct{}

func NewPythonTypeMapper() *PythonTypeMapper {
	return &PythonTypeMapper{}
}

func (m *PythonTypeMapper) MapType(baseType string, context TypeContext, isArray bool) (string, error) {
	typeMap := map[string]string{
		"void":    "None",
		"bool":    "bool",
		"char8":   "str",
		"char16":  "str",
		"int8":    "int",
		"int16":   "int",
		"int32":   "int",
		"int64":   "int",
		"uint8":   "int",
		"uint16":  "int",
		"uint32":  "int",
		"uint64":  "int",
		"ptr64":   "int",
		"float":   "float",
		"double":  "float",
		"string":  "str",
		"any":     "object",
		"vec2":    "Vector2",
		"vec3":    "Vector3",
		"vec4":    "Vector4",
		"mat4x4":  "Matrix4x4",
	}

	mapped, ok := typeMap[baseType]
	if !ok {
		// Assume it's a custom type
		mapped = baseType
	}

	if isArray {
		mapped = fmt.Sprintf("list[%s]", mapped)
	}

	return mapped, nil
}

func (m *PythonTypeMapper) MapParamType(param *manifest.ParamType, context TypeContext) (string, error) {
	// Check for enum
	if param.Enum != nil {
		typeName := param.Enum.Name
		if param.IsArray() {
			typeName = fmt.Sprintf("list[%s]", typeName)
		}
		return typeName, nil
	}

	// Check for function/delegate
	if param.Prototype != nil {
		return m.generateCallableType(param.Prototype)
	}

	return m.MapType(param.BaseType(), context, param.IsArray())
}

func (m *PythonTypeMapper) MapReturnType(retType *manifest.TypeInfo) (string, error) {
	// Check for enum
	if retType.Enum != nil {
		typeName := retType.Enum.Name
		if retType.IsArray() {
			typeName = fmt.Sprintf("list[%s]", typeName)
		}
		return typeName, nil
	}

	// Check for function/delegate
	if retType.Prototype != nil {
		return m.generateCallableType(retType.Prototype)
	}

	return m.MapType(retType.BaseType(), TypeContextReturn, retType.IsArray())
}

func (m *PythonTypeMapper) generateCallableType(proto *manifest.Prototype) (string, error) {
	// Generate parameter types
	paramTypes := []string{}
	for _, param := range proto.ParamTypes {
		pType, err := m.MapParamType(&param, TypeContextValue)
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
