package generator

import (
	"fmt"
	"strings"

	"github.com/untrustedmodders/plugify-generator/pkg/manifest"
)

// CppGenerator generates C++ header files
type CppGenerator struct {
	*BaseGenerator
}

// NewCppGenerator creates a new C++ generator
func NewCppGenerator() *CppGenerator {
	invalidNames := []string{
		"alignas", "alignof", "and", "and_eq", "asm", "auto", "bitand", "bitor",
		"bool", "break", "case", "catch", "char", "char8_t", "char16_t", "char32_t",
		"class", "compl", "concept", "const", "consteval", "constexpr", "constinit",
		"const_cast", "continue", "co_await", "co_return", "co_yield", "decltype",
		"default", "delete", "do", "double", "dynamic_cast", "else", "enum", "explicit",
		"export", "extern", "false", "float", "for", "friend", "goto", "if", "inline",
		"int", "long", "mutable", "namespace", "new", "noexcept", "not", "not_eq",
		"nullptr", "operator", "or", "or_eq", "private", "protected", "public",
		"register", "reinterpret_cast", "requires", "return", "short", "signed",
		"sizeof", "static", "static_assert", "static_cast", "struct", "switch",
		"template", "this", "thread_local", "throw", "true", "try", "typedef",
		"typeid", "typename", "union", "unsigned", "using", "virtual", "void",
		"volatile", "wchar_t", "while", "xor", "xor_eq",
	}

	return &CppGenerator{
		BaseGenerator: NewBaseGenerator("cpp", NewCppTypeMapper(), invalidNames),
	}
}

// Generate generates C++ bindings
func (g *CppGenerator) Generate(m *manifest.Manifest) (*GeneratorResult, error) {
	g.ResetCaches()

	var sb strings.Builder

	// Header guard and includes
	sb.WriteString("#pragma once\n\n")
	sb.WriteString("#include <plg/plugin.hpp>\n")
	sb.WriteString("#include <plg/any.hpp>\n")
	sb.WriteString("#include <cstdint>\n\n")
	sb.WriteString(fmt.Sprintf("// Generated from %s.pplugin\n\n", m.Name))

	// Namespace
	sb.WriteString(fmt.Sprintf("namespace %s {\n\n", m.Name))

	// Generate enums and delegates first
	enumsCode, err := g.generateEnumsAndDelegates(m)
	if err != nil {
		return nil, err
	}
	sb.WriteString(enumsCode)

	// Generate methods
	for _, method := range m.Methods {
		methodCode, err := g.generateMethod(m.Name, &method)
		if err != nil {
			return nil, fmt.Errorf("failed to generate method %s: %w", method.Name, err)
		}
		sb.WriteString(methodCode)
		sb.WriteString("\n")
	}

	// Close namespace
	sb.WriteString(fmt.Sprintf("} // namespace %s\n", m.Name))

	result := &GeneratorResult{
		Files: map[string]string{
			fmt.Sprintf("%s.hpp", m.Name): sb.String(),
		},
	}

	return result, nil
}

func (g *CppGenerator) generateEnumsAndDelegates(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	// Collect all enums and delegates
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

func (g *CppGenerator) generateEnum(enum *manifest.EnumType) string {
	var sb strings.Builder

	if enum.Description != "" {
		sb.WriteString(fmt.Sprintf("  // %s\n", enum.Description))
	}

	underlyingType := "int32_t"
	if enum.Type != "" {
		mapped, _ := g.typeMapper.MapType(enum.Type, TypeContextValue, false)
		underlyingType = mapped
	}

	sb.WriteString(fmt.Sprintf("  enum class %s : %s {\n", enum.Name, underlyingType))

	for i, val := range enum.Values {
		if val.Description != "" {
			sb.WriteString(fmt.Sprintf("    %s = %d, // %s\n", val.Name, val.Value, val.Description))
		} else {
			sb.WriteString(fmt.Sprintf("    %s = %d", val.Name, val.Value))
			if i < len(enum.Values)-1 {
				sb.WriteString(",")
			}
			sb.WriteString("\n")
		}
	}

	sb.WriteString("  };\n\n")
	return sb.String()
}

func (g *CppGenerator) generateDelegate(proto *manifest.Prototype) (string, error) {
	var sb strings.Builder

	if proto.Description != "" {
		sb.WriteString(fmt.Sprintf("  // %s\n", proto.Description))
	}

	// Generate return type
	retType, err := g.typeMapper.MapReturnType(&proto.RetType)
	if err != nil {
		return "", err
	}

	// Generate parameters
	params, err := FormatParameters(proto.ParamTypes, ParamFormatTypes, g.typeMapper, g.SanitizeName)
	if err != nil {
		return "", err
	}

	sb.WriteString(fmt.Sprintf("  using %s = %s (*)(%s);\n\n", proto.Name, retType, params))
	return sb.String(), nil
}

func (g *CppGenerator) generateMethod(pluginName string, method *manifest.Method) (string, error) {
	var sb strings.Builder

	// Generate documentation comment
	sb.WriteString("  /**\n")
	if method.Description != "" {
		sb.WriteString(fmt.Sprintf("   * @brief %s\n", method.Description))
	}
	sb.WriteString(fmt.Sprintf("   * @function %s\n", method.Name))
	for _, param := range method.ParamTypes {
		paramType := param.Type
		if param.Ref {
			paramType += "&"
		}
		sb.WriteString(fmt.Sprintf("   * @param %s (%s)", g.SanitizeName(param.Name), paramType))
		if param.Description != "" {
			sb.WriteString(fmt.Sprintf(": %s", param.Description))
		}
		sb.WriteString("\n")
	}
	if method.RetType.Type != "void" {
		sb.WriteString(fmt.Sprintf("   * @return %s", method.RetType.Type))
		if method.RetType.Description != "" {
			sb.WriteString(fmt.Sprintf(": %s", method.RetType.Description))
		}
		sb.WriteString("\n")
	}
	sb.WriteString("   */\n")

	// Generate function signature
	retType, err := g.typeMapper.MapReturnType(&method.RetType)
	if err != nil {
		return "", err
	}

	params, err := FormatParameters(method.ParamTypes, ParamFormatTypesAndNames, g.typeMapper, g.SanitizeName)
	if err != nil {
		return "", err
	}

	sb.WriteString(fmt.Sprintf("  inline %s %s(%s) {\n", retType, g.SanitizeName(method.Name), params))

	// Generate function body
	// Type alias for function pointer
	funcTypeParams, err := FormatParameters(method.ParamTypes, ParamFormatTypes, g.typeMapper, g.SanitizeName)
	if err != nil {
		return "", err
	}
	sb.WriteString(fmt.Sprintf("    using %sFn = %s (*)(%s);\n", method.Name, retType, funcTypeParams))
	sb.WriteString(fmt.Sprintf("    static %sFn __func = nullptr;\n", method.Name))
	sb.WriteString(fmt.Sprintf("    if (__func == nullptr) plg::GetMethodPtr2(\"%s.%s\", reinterpret_cast<void**>(&__func));\n", pluginName, method.Name))

	// Call function
	paramNames, err := FormatParameters(method.ParamTypes, ParamFormatNames, g.typeMapper, g.SanitizeName)
	if err != nil {
		return "", err
	}

	if method.RetType.Type == "void" {
		sb.WriteString(fmt.Sprintf("    __func(%s);\n", paramNames))
	} else {
		sb.WriteString(fmt.Sprintf("    return __func(%s);\n", paramNames))
	}

	sb.WriteString("  }\n")

	return sb.String(), nil
}

// CppTypeMapper implements type mapping for C++
type CppTypeMapper struct{}

func NewCppTypeMapper() *CppTypeMapper {
	return &CppTypeMapper{}
}

func (m *CppTypeMapper) MapType(baseType string, context TypeContext, isArray bool) (string, error) {
	// Base type mapping
	typeMap := map[string]string{
		"void":    "void",
		"bool":    "bool",
		"char8":   "char",
		"char16":  "char16_t",
		"int8":    "int8_t",
		"int16":   "int16_t",
		"int32":   "int32_t",
		"int64":   "int64_t",
		"uint8":   "uint8_t",
		"uint16":  "uint16_t",
		"uint32":  "uint32_t",
		"uint64":  "uint64_t",
		"ptr64":   "void*",
		"float":   "float",
		"double":  "double",
		"string":  "plg::string",
		"any":     "plg::any",
		"vec2":    "plg::vec2",
		"vec3":    "plg::vec3",
		"vec4":    "plg::vec4",
		"mat4x4":  "plg::mat4x4",
	}

	mapped, ok := typeMap[baseType]
	if !ok {
		// Assume it's a custom type (enum or delegate)
		mapped = baseType
	}

	// Handle arrays
	if isArray {
		mapped = fmt.Sprintf("plg::vector<%s>", mapped)
	}

	// Handle parameter context (value parameters)
	// Object-like types pass by const& even when not ref=true
	if context == TypeContextValue && !isArray && baseType != "void" {
		if m.isObjectLikeType(baseType) {
			mapped = fmt.Sprintf("const %s&", mapped)
		}
	}

	// Handle reference context (ref=true parameters)
	if context == TypeContextRef && !isArray && baseType != "void" {
		mapped = mapped + "&"
	}

	return mapped, nil
}

// isObjectLikeType returns true for types that should be passed by const& in parameters
func (m *CppTypeMapper) isObjectLikeType(baseType string) bool {
	objectLikeTypes := map[string]bool{
		"string":  true,
		"any":     true,
		"vec2":    true,
		"vec3":    true,
		"vec4":    true,
		"mat4x4":  true,
	}
	return objectLikeTypes[baseType]
}

func (m *CppTypeMapper) MapParamType(param *manifest.ParamType, context TypeContext) (string, error) {
	// Check for enum type first
	if param.Enum != nil {
		typeName := param.Enum.Name
		if param.IsArray() {
			typeName = fmt.Sprintf("plg::vector<%s>", typeName)
		}
		// Enums are primitive types, pass by value or reference
		if param.Ref && !param.IsArray() {
			return typeName + "&", nil
		}
		return typeName, nil
	}

	// Check for function/delegate type
	if param.Prototype != nil {
		return param.Prototype.Name, nil
	}

	// Regular type mapping
	ctx := TypeContextValue
	if param.Ref {
		ctx = TypeContextRef
	}

	return m.MapType(param.BaseType(), ctx, param.IsArray())
}

func (m *CppTypeMapper) MapReturnType(retType *manifest.TypeInfo) (string, error) {
	// Check for enum type
	if retType.Enum != nil {
		typeName := retType.Enum.Name
		if retType.IsArray() {
			typeName = fmt.Sprintf("plg::vector<%s>", typeName)
		}
		return typeName, nil
	}

	// Check for function/delegate type
	if retType.Prototype != nil {
		return retType.Prototype.Name, nil
	}

	// Regular type mapping - returns always by value
	return m.MapType(retType.BaseType(), TypeContextReturn, retType.IsArray())
}
