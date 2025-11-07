package generator

import (
	"fmt"
	"strings"

	"github.com/untrustedmodders/plugify-generator/pkg/manifest"
)

// V8Generator generates TypeScript definition files for V8/JavaScript
type V8Generator struct {
	*BaseGenerator
}

// NewV8Generator creates a new V8/JavaScript generator
func NewV8Generator() *V8Generator {
	invalidNames := []string{
		"abstract", "arguments", "await", "boolean", "break", "byte", "case",
		"catch", "char", "class", "const", "continue", "debugger", "default",
		"delete", "do", "double", "else", "enum", "eval", "export", "extends",
		"false", "final", "finally", "float", "for", "function", "goto", "if",
		"implements", "import", "in", "instanceof", "int", "interface", "let",
		"long", "native", "new", "null", "package", "private", "protected",
		"public", "return", "short", "static", "super", "switch", "synchronized",
		"this", "throw", "throws", "transient", "true", "try", "typeof", "var",
		"void", "volatile", "while", "with", "yield",
	}

	return &V8Generator{
		BaseGenerator: NewBaseGenerator("v8", NewV8TypeMapper(), invalidNames),
	}
}

// Generate generates V8/JavaScript TypeScript definitions
func (g *V8Generator) Generate(m *manifest.Manifest) (*GeneratorResult, error) {
	g.ResetCaches()

	var sb strings.Builder

	// Header comment
	sb.WriteString(fmt.Sprintf("// Generated from %s.pplugin\n\n", m.Name))

	// Import plugify base types
	sb.WriteString("declare module \"plugify\" {\n")
	sb.WriteString("  type Vector2 = { x: number; y: number };\n")
	sb.WriteString("  type Vector3 = { x: number; y: number; z: number };\n")
	sb.WriteString("  type Vector4 = { x: number; y: number; z: number; w: number };\n")
	sb.WriteString("  type Matrix4x4 = { data: number[] };\n")
	sb.WriteString("}\n\n")

	// Module declaration
	sb.WriteString(fmt.Sprintf("declare module \":%s\" {\n", m.Name))
	sb.WriteString("  import { Vector2, Vector3, Vector4, Matrix4x4 } from \"plugify\";\n\n")

	// Generate enums and delegates
	enumsCode, err := g.generateEnumsAndDelegates(m)
	if err != nil {
		return nil, err
	}
	sb.WriteString(enumsCode)

	// Generate methods
	for _, method := range m.Methods {
		methodCode, err := g.generateMethod(&method)
		if err != nil {
			return nil, fmt.Errorf("failed to generate method %s: %w", method.Name, err)
		}
		sb.WriteString(methodCode)
	}

	// Close module
	sb.WriteString("}\n")

	result := &GeneratorResult{
		Files: map[string]string{
			fmt.Sprintf("%s.d.ts", m.Name): sb.String(),
		},
	}

	return result, nil
}

func (g *V8Generator) generateEnumsAndDelegates(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	for _, method := range m.Methods {
		// Check return type
		if method.RetType.Enum != nil && !g.IsEnumCached(method.RetType.Enum.Name) {
			enumCode := g.generateEnum(method.RetType.Enum)
			sb.WriteString(enumCode)
			g.CacheEnum(method.RetType.Enum.Name)
		}
		if method.RetType.Prototype != nil && !g.IsDelegateCached(method.RetType.Prototype.Name) {
			delegateCode, err := g.generateDelegate(method.RetType.Prototype)
			if err != nil {
				return "", err
			}
			sb.WriteString(delegateCode)
			g.CacheDelegate(method.RetType.Prototype.Name)
		}

		// Check parameters
		for _, param := range method.ParamTypes {
			if param.Enum != nil && !g.IsEnumCached(param.Enum.Name) {
				enumCode := g.generateEnum(param.Enum)
				sb.WriteString(enumCode)
				g.CacheEnum(param.Enum.Name)
			}
			if param.Prototype != nil && !g.IsDelegateCached(param.Prototype.Name) {
				delegateCode, err := g.generateDelegate(param.Prototype)
				if err != nil {
					return "", err
				}
				sb.WriteString(delegateCode)
				g.CacheDelegate(param.Prototype.Name)
			}
		}
	}

	if sb.Len() > 0 {
		sb.WriteString("\n")
	}

	return sb.String(), nil
}

func (g *V8Generator) generateEnum(enum *manifest.EnumType) string {
	var sb strings.Builder

	if enum.Description != "" {
		sb.WriteString(fmt.Sprintf("  /** %s */\n", enum.Description))
	}

	sb.WriteString(fmt.Sprintf("  export const enum %s {\n", enum.Name))

	for i, val := range enum.Values {
		if val.Description != "" {
			sb.WriteString(fmt.Sprintf("    /** %s */\n", val.Description))
		}
		sb.WriteString(fmt.Sprintf("    %s = %d", val.Name, val.Value))
		if i < len(enum.Values)-1 {
			sb.WriteString(",")
		}
		sb.WriteString("\n")
	}

	sb.WriteString("  }\n\n")
	return sb.String()
}

func (g *V8Generator) generateDelegate(proto *manifest.Prototype) (string, error) {
	var sb strings.Builder

	if proto.Description != "" {
		sb.WriteString(fmt.Sprintf("  /** %s */\n", proto.Description))
	}

	// Generate return type
	retType, err := g.typeMapper.MapReturnType(&proto.RetType)
	if err != nil {
		return "", err
	}

	// Generate parameters
	params, err := FormatParameters(proto.ParamTypes, ParamFormatTypesAndNames, g.typeMapper, g.SanitizeName)
	if err != nil {
		return "", err
	}

	sb.WriteString(fmt.Sprintf("  export type %s = (%s) => %s;\n\n", proto.Name, params, retType))
	return sb.String(), nil
}

func (g *V8Generator) generateMethod(method *manifest.Method) (string, error) {
	var sb strings.Builder

	// JSDoc comment
	sb.WriteString("  /**\n")
	if method.Description != "" {
		sb.WriteString(fmt.Sprintf("   * %s\n", method.Description))
	}
	for _, param := range method.ParamTypes {
		sb.WriteString(fmt.Sprintf("   * @param %s", g.SanitizeName(param.Name)))
		if param.Description != "" {
			sb.WriteString(fmt.Sprintf(" - %s", param.Description))
		}
		sb.WriteString("\n")
	}
	if method.RetType.Type != "void" && method.RetType.Description != "" {
		sb.WriteString(fmt.Sprintf("   * @returns %s\n", method.RetType.Description))
	}
	sb.WriteString("   */\n")

	// Function signature
	retType, err := g.typeMapper.MapReturnType(&method.RetType)
	if err != nil {
		return "", err
	}

	params, err := FormatParameters(method.ParamTypes, ParamFormatTypesAndNames, g.typeMapper, g.SanitizeName)
	if err != nil {
		return "", err
	}

	sb.WriteString(fmt.Sprintf("  export function %s(%s): %s;\n\n", g.SanitizeName(method.Name), params, retType))

	return sb.String(), nil
}

// V8TypeMapper implements type mapping for V8/JavaScript
type V8TypeMapper struct{}

func NewV8TypeMapper() *V8TypeMapper {
	return &V8TypeMapper{}
}

func (m *V8TypeMapper) MapType(baseType string, context TypeContext, isArray bool) (string, error) {
	typeMap := map[string]string{
		"void":    "void",
		"bool":    "boolean",
		"char8":   "number",
		"char16":  "number",
		"int8":    "number",
		"int16":   "number",
		"int32":   "number",
		"int64":   "number",
		"uint8":   "number",
		"uint16":  "number",
		"uint32":  "number",
		"uint64":  "bigint",
		"ptr64":   "bigint",
		"float":   "number",
		"double":  "number",
		"string":  "string",
		"any":     "any",
		"vec2":    "Vector2",
		"vec3":    "Vector3",
		"vec4":    "Vector4",
		"mat4x4":  "Matrix4x4",
	}

	mapped, ok := typeMap[baseType]
	if !ok {
		// Assume it's a custom type (enum or delegate)
		mapped = baseType
	}

	// Handle arrays
	if isArray {
		mapped = mapped + "[]"
	}

	return mapped, nil
}

func (m *V8TypeMapper) MapParamType(param *manifest.ParamType, context TypeContext) (string, error) {
	return m.MapType(param.BaseType(), context, param.IsArray())
}

func (m *V8TypeMapper) MapReturnType(retType *manifest.TypeInfo) (string, error) {
	return m.MapType(retType.BaseType(), TypeContextReturn, retType.IsArray())
}
