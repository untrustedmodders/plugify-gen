package generator

import (
	"fmt"

	"github.com/untrustedmodders/plugify-generator/pkg/manifest"
)

// GolangTypeMapper implements type mapping for Go
type GolangTypeMapper struct {
	// Primary type map for Go types
	typesMap map[string]string
	// C return types map
	rtypesMap map[string]string
	// CGo types map
	ctypesMap map[string]string
	// Function type names
	ftypesMap map[string]string
	// Value type cast constructors
	valTypeCastMap map[string]string
	// Return type cast names
	retTypeCastMap map[string]string
	// Assignment type cast helpers
	assTypeCastMap map[string]string
	// Destructor functions
	delTypeCastMap map[string]string
}

// NewGolangTypeMapper creates a new Go type mapper
func NewGolangTypeMapper() *GolangTypeMapper {
	return &GolangTypeMapper{
		typesMap: map[string]string{
			"void":      "",
			"bool":      "bool",
			"char8":     "int8",
			"char16":    "uint16",
			"int8":      "int8",
			"int16":     "int16",
			"int32":     "int32",
			"int64":     "int64",
			"uint8":     "uint8",
			"uint16":    "uint16",
			"uint32":    "uint32",
			"uint64":    "uint64",
			"ptr64":     "uintptr",
			"float":     "float32",
			"double":    "float64",
			"function":  "delegate",
			"string":    "string",
			"any":       "interface{}",
			"bool[]":    "[]bool",
			"char8[]":   "[]int8",
			"char16[]":  "[]uint16",
			"int8[]":    "[]int8",
			"int16[]":   "[]int16",
			"int32[]":   "[]int32",
			"int64[]":   "[]int64",
			"uint8[]":   "[]uint8",
			"uint16[]":  "[]uint16",
			"uint32[]":  "[]uint32",
			"uint64[]":  "[]uint64",
			"ptr64[]":   "[]uintptr",
			"float[]":   "[]float32",
			"double[]":  "[]float64",
			"string[]":  "[]string",
			"any[]":     "[]interface{}",
			"vec2[]":    "[]plugify.Vector2",
			"vec3[]":    "[]plugify.Vector3",
			"vec4[]":    "[]plugify.Vector4",
			"mat4x4[]":  "[]plugify.Matrix4x4",
			"vec2":      "plugify.Vector2",
			"vec3":      "plugify.Vector3",
			"vec4":      "plugify.Vector4",
			"mat4x4":    "plugify.Matrix4x4",
		},
		rtypesMap: map[string]string{
			"void":      "void",
			"bool":      "bool",
			"char8":     "int8_t",
			"char16":    "uint16_t",
			"int8":      "int8_t",
			"int16":     "int16_t",
			"int32":     "int32_t",
			"int64":     "int64_t",
			"uint8":     "uint8_t",
			"uint16":    "uint16_t",
			"uint32":    "uint32_t",
			"uint64":    "uint64_t",
			"ptr64":     "uintptr_t",
			"float":     "float",
			"double":    "double",
			"function":  "void*",
			"string":    "String*",
			"any":       "Variant*",
			"bool[]":    "Vector*",
			"char8[]":   "Vector*",
			"char16[]":  "Vector*",
			"int8[]":    "Vector*",
			"int16[]":   "Vector*",
			"int32[]":   "Vector*",
			"int64[]":   "Vector*",
			"uint8[]":   "Vector*",
			"uint16[]":  "Vector*",
			"uint32[]":  "Vector*",
			"uint64[]":  "Vector*",
			"ptr64[]":   "Vector*",
			"float[]":   "Vector*",
			"double[]":  "Vector*",
			"string[]":  "Vector*",
			"any[]":     "Vector*",
			"vec2[]":    "Vector*",
			"vec3[]":    "Vector*",
			"vec4[]":    "Vector*",
			"mat4x4[]":  "Vector*",
			"vec2":      "Vector2*",
			"vec3":      "Vector3*",
			"vec4":      "Vector4*",
			"mat4x4":    "Matrix4x4*",
		},
		ctypesMap: map[string]string{
			"void":      "",
			"bool":      "bool",
			"char8":     "int8",
			"char16":    "uint16",
			"int8":      "int8",
			"int16":     "int16",
			"int32":     "int32",
			"int64":     "int64",
			"uint8":     "uint8",
			"uint16":    "uint16",
			"uint32":    "uint32",
			"uint64":    "uint64",
			"ptr64":     "uintptr",
			"float":     "float32",
			"double":    "float64",
			"function":  "func",
			"string":    "C.String",
			"any":       "C.Variant",
			"bool[]":    "C.Vector",
			"char8[]":   "C.Vector",
			"char16[]":  "C.Vector",
			"int8[]":    "C.Vector",
			"int16[]":   "C.Vector",
			"int32[]":   "C.Vector",
			"int64[]":   "C.Vector",
			"uint8[]":   "C.Vector",
			"uint16[]":  "C.Vector",
			"uint32[]":  "C.Vector",
			"uint64[]":  "C.Vector",
			"ptr64[]":   "C.Vector",
			"float[]":   "C.Vector",
			"double[]":  "C.Vector",
			"string[]":  "C.Vector",
			"any[]":     "C.Vector",
			"vec2[]":    "C.Vector",
			"vec3[]":    "C.Vector",
			"vec4[]":    "C.Vector",
			"mat4x4[]":  "C.Vector",
			"vec2":      "C.Vector2",
			"vec3":      "C.Vector3",
			"vec4":      "C.Vector4",
			"mat4x4":    "C.Matrix4x4",
		},
		ftypesMap: map[string]string{
			"void":      "",
			"bool":      "Bool",
			"char8":     "Int8",
			"char16":    "UInt16",
			"int8":      "Int8",
			"int16":     "Int16",
			"int32":     "Int32",
			"int64":     "Int64",
			"uint8":     "UInt8",
			"uint16":    "UInt16",
			"uint32":    "UInt32",
			"uint64":    "UInt64",
			"ptr64":     "Pointer",
			"ptr32":     "Pointer",
			"float":     "Float",
			"double":    "Double",
			"string":    "String",
			"any":       "Variant",
			"vec2":      "Vector2",
			"vec3":      "Vector3",
			"vec4":      "Vector4",
			"mat4x4":    "Matrix4x4",
			"bool[]":    "Bool",
			"char8[]":   "Char8",
			"char16[]":  "Char16",
			"int8[]":    "Int8",
			"int16[]":   "Int16",
			"int32[]":   "Int32",
			"int64[]":   "Int64",
			"uint8[]":   "UInt8",
			"uint16[]":  "UInt16",
			"uint32[]":  "UInt32",
			"uint64[]":  "UInt64",
			"ptr32[]":   "Pointer",
			"ptr64[]":   "Pointer",
			"float[]":   "Float",
			"double[]":  "Double",
			"string[]":  "String",
			"any[]":     "Variant",
			"vec2[]":    "Vector2",
			"vec3[]":    "Vector3",
			"vec4[]":    "Vector4",
			"mat4x4[]":  "Matrix4x4",
		},
		valTypeCastMap: map[string]string{
			"void":      "",
			"bool":      "C.bool",
			"char8":     "C.int8_t",
			"char16":    "C.uint16_t",
			"int8":      "C.int8_t",
			"int16":     "C.int16_t",
			"int32":     "C.int32_t",
			"int64":     "C.int64_t",
			"uint8":     "C.uint8_t",
			"uint16":    "C.uint16_t",
			"uint32":    "C.uint32_t",
			"uint64":    "C.uint64_t",
			"ptr64":     "C.uintptr_t",
			"float":     "C.float",
			"double":    "C.double",
			"function":  "plugify.GetFunctionPointerForDelegate",
			"string":    "plugify.ConstructString",
			"any":       "plugify.ConstructVariant",
			"bool[]":    "plugify.ConstructVectorBool",
			"char8[]":   "plugify.ConstructVectorChar8",
			"char16[]":  "plugify.ConstructVectorChar16",
			"int8[]":    "plugify.ConstructVectorInt8",
			"int16[]":   "plugify.ConstructVectorInt16",
			"int32[]":   "plugify.ConstructVectorInt32",
			"int64[]":   "plugify.ConstructVectorInt64",
			"uint8[]":   "plugify.ConstructVectorUInt8",
			"uint16[]":  "plugify.ConstructVectorUInt16",
			"uint32[]":  "plugify.ConstructVectorUInt32",
			"uint64[]":  "plugify.ConstructVectorUInt64",
			"ptr64[]":   "plugify.ConstructVectorPointer",
			"float[]":   "plugify.ConstructVectorFloat",
			"double[]":  "plugify.ConstructVectorDouble",
			"string[]":  "plugify.ConstructVectorString",
			"any[]":     "plugify.ConstructVectorVariant",
			"vec2[]":    "plugify.ConstructVectorVector2",
			"vec3[]":    "plugify.ConstructVectorVector3",
			"vec4[]":    "plugify.ConstructVectorVector4",
			"mat4x4[]":  "plugify.ConstructVectorMatrix4x4",
			"vec2":      "C.Vector2",
			"vec3":      "C.Vector3",
			"vec4":      "C.Vector4",
			"mat4x4":    "C.Matrix4x4",
		},
		retTypeCastMap: map[string]string{
			"void":      "",
			"bool":      "bool",
			"char8":     "int8",
			"char16":    "uint16",
			"int8":      "int8",
			"int16":     "int16",
			"int32":     "int32",
			"int64":     "int64",
			"uint8":     "uint8",
			"uint16":    "uint16",
			"uint32":    "uint32",
			"uint64":    "uint64",
			"ptr64":     "uintptr",
			"float":     "float32",
			"double":    "float64",
			"function":  "",
			"string":    "plugify.PlgString",
			"any":       "plugify.PlgVariant",
			"bool[]":    "plugify.PlgVector",
			"char8[]":   "plugify.PlgVector",
			"char16[]":  "plugify.PlgVector",
			"int8[]":    "plugify.PlgVector",
			"int16[]":   "plugify.PlgVector",
			"int32[]":   "plugify.PlgVector",
			"int64[]":   "plugify.PlgVector",
			"uint8[]":   "plugify.PlgVector",
			"uint16[]":  "plugify.PlgVector",
			"uint32[]":  "plugify.PlgVector",
			"uint64[]":  "plugify.PlgVector",
			"ptr64[]":   "plugify.PlgVector",
			"float[]":   "plugify.PlgVector",
			"double[]":  "plugify.PlgVector",
			"string[]":  "plugify.PlgVector",
			"any[]":     "plugify.PlgVector",
			"vec2[]":    "plugify.PlgVector",
			"vec3[]":    "plugify.PlgVector",
			"vec4[]":    "plugify.PlgVector",
			"mat4x4[]":  "plugify.PlgVector",
			"vec2":      "plugify.Vector2",
			"vec3":      "plugify.Vector3",
			"vec4":      "plugify.Vector4",
			"mat4x4":    "plugify.Matrix4x4",
		},
		assTypeCastMap: map[string]string{
			"void":      "",
			"bool":      "bool",
			"char8":     "int8",
			"char16":    "uint16",
			"int8":      "int8",
			"int16":     "int16",
			"int32":     "int32",
			"int64":     "int64",
			"uint8":     "uint8",
			"uint16":    "uint16",
			"uint32":    "uint32",
			"uint64":    "uint64",
			"ptr64":     "uintptr",
			"float":     "float32",
			"double":    "float64",
			"function":  "",
			"string":    "plugify.GetStringData",
			"any":       "plugify.GetVariantData",
			"bool[]":    "plugify.GetVectorDataBool",
			"char8[]":   "plugify.GetVectorDataChar8",
			"char16[]":  "plugify.GetVectorDataChar16",
			"int8[]":    "plugify.GetVectorDataInt8",
			"int16[]":   "plugify.GetVectorDataInt16",
			"int32[]":   "plugify.GetVectorDataInt32",
			"int64[]":   "plugify.GetVectorDataInt64",
			"uint8[]":   "plugify.GetVectorDataUInt8",
			"uint16[]":  "plugify.GetVectorDataUInt16",
			"uint32[]":  "plugify.GetVectorDataUInt32",
			"uint64[]":  "plugify.GetVectorDataUInt64",
			"ptr64[]":   "plugify.GetVectorDataPointer",
			"float[]":   "plugify.GetVectorDataFloat",
			"double[]":  "plugify.GetVectorDataDouble",
			"string[]":  "plugify.GetVectorDataString",
			"any[]":     "plugify.GetVectorDataVariant",
			"vec2[]":    "plugify.GetVectorDataVector2",
			"vec3[]":    "plugify.GetVectorDataVector3",
			"vec4[]":    "plugify.GetVectorDataVector4",
			"mat4x4[]":  "plugify.GetVectorDataMatrix4x4",
			"vec2":      "plugify.Vector2",
			"vec3":      "plugify.Vector3",
			"vec4":      "plugify.Vector4",
			"mat4x4":    "plugify.Matrix4x4",
		},
		delTypeCastMap: map[string]string{
			"void":      "",
			"bool":      "",
			"char8":     "",
			"char16":    "",
			"int8":      "",
			"int16":     "",
			"int32":     "",
			"int64":     "",
			"uint8":     "",
			"uint16":    "",
			"uint32":    "",
			"uint64":    "",
			"ptr64":     "",
			"float":     "",
			"double":    "",
			"function":  "",
			"string":    "plugify.DestroyString",
			"any":       "plugify.DestroyVariant",
			"bool[]":    "plugify.DestroyVectorBool",
			"char8[]":   "plugify.DestroyVectorChar8",
			"char16[]":  "plugify.DestroyVectorChar16",
			"int8[]":    "plugify.DestroyVectorInt8",
			"int16[]":   "plugify.DestroyVectorInt16",
			"int32[]":   "plugify.DestroyVectorInt32",
			"int64[]":   "plugify.DestroyVectorInt64",
			"uint8[]":   "plugify.DestroyVectorUInt8",
			"uint16[]":  "plugify.DestroyVectorUInt16",
			"uint32[]":  "plugify.DestroyVectorUInt32",
			"uint64[]":  "plugify.DestroyVectorUInt64",
			"ptr64[]":   "plugify.DestroyVectorPointer",
			"float[]":   "plugify.DestroyVectorFloat",
			"double[]":  "plugify.DestroyVectorDouble",
			"string[]":  "plugify.DestroyVectorString",
			"any[]":     "plugify.DestroyVectorVariant",
			"vec2[]":    "plugify.DestroyVectorVector2",
			"vec3[]":    "plugify.DestroyVectorVector3",
			"vec4[]":    "plugify.DestroyVectorVector4",
			"mat4x4[]":  "plugify.DestroyVectorMatrix4x4",
			"vec2":      "",
			"vec3":      "",
			"vec4":      "",
			"mat4x4":    "",
		},
	}
}

// MapType implements TypeMapper interface
func (m *GolangTypeMapper) MapType(baseType string, context TypeContext, isArray bool) (string, error) {
	typeKey := baseType
	if isArray {
		typeKey += "[]"
	}

	result, ok := m.typesMap[typeKey]
	if !ok {
		return "", fmt.Errorf("unsupported type: %s", typeKey)
	}

	return result, nil
}

// MapParamType implements TypeMapper interface
func (m *GolangTypeMapper) MapParamType(param *manifest.ParamType, context TypeContext) (string, error) {
	// Handle delegate/function types
	if param.Type == "function" && param.Prototype != nil {
		return param.Prototype.Name, nil
	}

	// Handle enum types
	if param.Enum != nil {
		enumName := param.Enum.Name
		if param.Type[len(param.Type)-2:] == "[]" {
			if param.Ref {
				return "*[]" + enumName, nil
			}
			return "[]" + enumName, nil
		}
		if param.Ref {
			return "*" + enumName, nil
		}
		return enumName, nil
	}

	// Get base type
	goType, err := m.MapType(param.Type, context, false)
	if err != nil {
		return "", err
	}

	// Handle reference parameters
	if param.Ref {
		return "*" + goType, nil
	}

	return goType, nil
}

// MapReturnType implements TypeMapper interface
func (m *GolangTypeMapper) MapReturnType(retType *manifest.TypeInfo) (string, error) {
	if retType == nil || retType.Type == "void" {
		return "", nil
	}

	// Handle delegate/function types
	if retType.Type == "function" && retType.Prototype != nil {
		return retType.Prototype.Name, nil
	}

	// Handle enum types
	if retType.Enum != nil {
	    enumName := retType.Enum.Name
        if retType.Type[len(retType.Type)-2:] == "[]" {
            return "[]" + enumName, nil
        }
        return enumName, nil
	}

	// Get base type
	goType, err := m.MapType(retType.Type, TypeContextReturn, false)
	if err != nil {
		return "", err
	}

	return goType, nil
}

// GetCType returns the C type for a given manifest type
func (m *GolangTypeMapper) GetCType(baseType string, isRef bool, isRet bool) string {
	result := m.rtypesMap[baseType]

	// If it contains a pointer, handle return type removal
	if isRet && result != "void*" && len(result) > 0 && result[len(result)-1] == '*' {
		return result[:len(result)-1] // Remove trailing *
	}

	// If it's a reference and doesn't already have a pointer
	if isRef && len(result) > 0 && result[len(result)-1] != '*' {
		return result + "*"
	}

	return result
}

// Helper methods for type checking
func (m *GolangTypeMapper) IsObjType(baseType string) bool {
	return baseType == "string" || baseType == "any" || len(baseType) > 2 && baseType[len(baseType)-2:] == "[]"
}

func (m *GolangTypeMapper) IsPodType(baseType string) bool {
	return baseType == "vec2" || baseType == "vec3" || baseType == "vec4" || baseType == "mat4x4"
}