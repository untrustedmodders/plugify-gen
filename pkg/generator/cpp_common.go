package generator

import (
	"fmt"

	"github.com/untrustedmodders/plugify-gen/pkg/manifest"
)

// CppCommonTypeMapper implements type mapping for both C++ generators (cpp and cxx)
// This is shared between traditional C++ headers and C++20 modules
type CppCommonTypeMapper struct{}

func NewCppCommonTypeMapper() *CppCommonTypeMapper {
	return &CppCommonTypeMapper{}
}

func (m *CppCommonTypeMapper) MapType(baseType string, context TypeContext, isArray bool) (string, error) {
	// Base type mapping
	typeMap := map[string]string{
		"void":   "void",
		"bool":   "bool",
		"char8":  "char",
		"char16": "char16_t",
		"int8":   "int8_t",
		"int16":  "int16_t",
		"int32":  "int32_t",
		"int64":  "int64_t",
		"uint8":  "uint8_t",
		"uint16": "uint16_t",
		"uint32": "uint32_t",
		"uint64": "uint64_t",
		"ptr64":  "void*",
		"float":  "float",
		"double": "double",
		"string": "plg::string",
		"any":    "plg::any",
		"vec2":   "plg::vec2",
		"vec3":   "plg::vec3",
		"vec4":   "plg::vec4",
		"mat4x4": "plg::mat4x4",
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
	if context == TypeContextValue && baseType != "void" {
		if m.isObjectLikeType(baseType) || isArray {
			mapped = fmt.Sprintf("const %s&", mapped)
		}
	}

	// Handle reference context (ref=true parameters)
	if context == TypeContextRef && baseType != "void" {
		mapped = mapped + "&"
	}

	return mapped, nil
}

// isObjectLikeType returns true for types that should be passed by const& in parameters
func (m *CppCommonTypeMapper) isObjectLikeType(baseType string) bool {
	objectLikeTypes := map[string]bool{
		"string": true,
		"any":    true,
		"vec2":   true,
		"vec3":   true,
		"vec4":   true,
		"mat4x4": true,
	}
	return objectLikeTypes[baseType]
}

func (m *CppCommonTypeMapper) MapParamType(param *manifest.ParamType, context TypeContext) (string, error) {
	// Check for enum type first
	if param.Enum != nil {
		typeName := param.Enum.Name
		if param.IsArray() {
			typeName = fmt.Sprintf("plg::vector<%s>", typeName)
		}
		// Enums are primitive types, pass by value or reference
		if param.Ref {
			return typeName + "&", nil
		} else if param.IsArray() {
			return fmt.Sprintf("const %s&", typeName), nil
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

func (m *CppCommonTypeMapper) MapReturnType(retType *manifest.TypeInfo) (string, error) {
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

func (m *CppCommonTypeMapper) MapHandleType(class *manifest.Class) (string, string, error) {
	invalidValue := class.InvalidValue
	handleType, err := m.MapType(class.HandleType, TypeContextReturn, false)
	if err != nil {
		return "", "", err
	}

	nullptr := invalidValue == "0" || invalidValue == "" || invalidValue == "NULL" || invalidValue == "nullptr"
	if class.HandleType == "ptr64" && nullptr {
		invalidValue = "nullptr"
	} else if invalidValue == "" {
		invalidValue = "{}"
	}

	return invalidValue, handleType, nil
}
