package generator

import (
	"fmt"
	"strings"

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
	if context != TypeContextAlias && isArray {
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
	objectLikeTypes := map[string]struct{}{
		"string": {},
		"any":    {},
		"vec2":   {},
		"vec3":   {},
		"vec4":   {},
		"mat4x4": {},
	}
	_, ok := objectLikeTypes[baseType]
	return ok
}

func (m *CppCommonTypeMapper) MapParamType(param *manifest.ParamType) (string, error) {
	// Regular type mapping
	ctx := TypeContextValue
	if param.Ref {
		ctx = TypeContextRef
	}

	var typeName string
	switch {
	case param.Enum != nil:
		typeName = param.Enum.Name

	case param.Alias != nil:
		typeName = param.Alias.Name

	case param.Prototype != nil:
		return param.Prototype.Name, nil

	default:
		typeName = param.BaseType()
	}

	return m.MapType(typeName, ctx, param.IsArray())
}

func (m *CppCommonTypeMapper) MapReturnType(retType *manifest.RetType) (string, error) {
	var typeName string

	switch {
	case retType.Enum != nil:
		typeName = retType.Enum.Name

	case retType.Alias != nil:
		typeName = retType.Alias.Name

	case retType.Prototype != nil:
		return retType.Prototype.Name, nil

	default:
		typeName = retType.BaseType()
	}

	// Regular type mapping - returns always by value
	return m.MapType(typeName, TypeContextReturn, retType.IsArray())
}

func (m *CppCommonTypeMapper) MapHandleType(class *manifest.Class) (string, string, error) {
	invalidValue := class.InvalidValue
	handleType, err := m.MapType(class.HandleType, TypeContextReturn, false)
	if err != nil {
		return "", "", err
	}

	nullptr := invalidValue == "0" || invalidValue == "" || invalidValue == "NULL" || invalidValue == "nullptr"
	if strings.HasPrefix(class.HandleType, "ptr") && nullptr {
		invalidValue = "nullptr"
	} else if invalidValue == "" {
		invalidValue = "{}"
	}

	return invalidValue, handleType, nil
}
