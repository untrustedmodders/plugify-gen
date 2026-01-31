package generator

import (
	"strings"

	"github.com/untrustedmodders/plugify-gen/pkg/manifest"
)

// DotnetTypeMapper implements type mapping for C#/.NET
type DotnetTypeMapper struct {
	// Public C# types
	typesMap map[string]string
	// C types for unmanaged calls
	ctypesMap map[string]string
	// Constructor functions for marshaling
	constructorMap map[string]string
	// Return type wrappers
	returnTypeMap map[string]string
	// Data converter functions
	dataConverterMap map[string]string
	// Size functions for arrays
	sizeFunctionMap map[string]string
	// Destructor/cleanup functions
	destructorMap map[string]string
}

func NewDotnetTypeMapper() *DotnetTypeMapper {
	return &DotnetTypeMapper{
		typesMap:         initTypesMap(),
		ctypesMap:        initCTypesMap(),
		constructorMap:   initConstructorMap(),
		returnTypeMap:    initReturnTypeMap(),
		dataConverterMap: initDataConverterMap(),
		sizeFunctionMap:  initSizeFunctionMap(),
		destructorMap:    initDestructorMap(),
	}
}

func initTypesMap() map[string]string {
	return map[string]string{
		"void":   "void",
		"bool":   "Bool8",
		"char8":  "Char8",
		"char16": "Char16",
		"int8":   "sbyte",
		"int16":  "short",
		"int32":  "int",
		"int64":  "long",
		"uint8":  "byte",
		"uint16": "ushort",
		"uint32": "uint",
		"uint64": "ulong",
		"ptr64":  "nint",
		"float":  "float",
		"double": "double",
		"string": "string",
		"any":    "object",
		"vec2":   "Vector2",
		"vec3":   "Vector3",
		"vec4":   "Vector4",
		"mat4x4": "Matrix4x4",
	}
}

func initCTypesMap() map[string]string {
	return map[string]string{
		"void":   "void",
		"bool":   "Bool8",
		"char8":  "Char8",
		"char16": "Char16",
		"int8":   "sbyte",
		"int16":  "short",
		"int32":  "int",
		"int64":  "long",
		"uint8":  "byte",
		"uint16": "ushort",
		"uint32": "uint",
		"uint64": "ulong",
		"ptr64":  "nint",
		"float":  "float",
		"double": "double",
		"string": "String192*",
		"any":    "Variant256*",
		"vec2":   "Vector2*",
		"vec3":   "Vector3*",
		"vec4":   "Vector4*",
		"mat4x4": "Matrix4x4*",
	}
}

func initConstructorMap() map[string]string {
	return map[string]string{
		"function": "Marshalling.GetFunctionPointerForDelegate",
		"string":   "NativeMethods.ConstructString",
		"any":      "NativeMethods.ConstructVariant",
		"bool[]":   "NativeMethods.ConstructVectorBool",
		"char8[]":  "NativeMethods.ConstructVectorChar8",
		"char16[]": "NativeMethods.ConstructVectorChar16",
		"int8[]":   "NativeMethods.ConstructVectorInt8",
		"int16[]":  "NativeMethods.ConstructVectorInt16",
		"int32[]":  "NativeMethods.ConstructVectorInt32",
		"int64[]":  "NativeMethods.ConstructVectorInt64",
		"uint8[]":  "NativeMethods.ConstructVectorUInt8",
		"uint16[]": "NativeMethods.ConstructVectorUInt16",
		"uint32[]": "NativeMethods.ConstructVectorUInt32",
		"uint64[]": "NativeMethods.ConstructVectorUInt64",
		"ptr64[]":  "NativeMethods.ConstructVectorIntPtr",
		"float[]":  "NativeMethods.ConstructVectorFloat",
		"double[]": "NativeMethods.ConstructVectorDouble",
		"string[]": "NativeMethods.ConstructVectorString",
		"any[]":    "NativeMethods.ConstructVectorVariant",
		"vec2[]":   "NativeMethods.ConstructVectorVector2",
		"vec3[]":   "NativeMethods.ConstructVectorVector3",
		"vec4[]":   "NativeMethods.ConstructVectorVector4",
		"mat4x4[]": "NativeMethods.ConstructVectorMatrix4x4",
	}
}

func initReturnTypeMap() map[string]string {
	return map[string]string{
		"string":   "String192",
		"any":      "Variant256",
		"bool[]":   "Vector192",
		"char8[]":  "Vector192",
		"char16[]": "Vector192",
		"int8[]":   "Vector192",
		"int16[]":  "Vector192",
		"int32[]":  "Vector192",
		"int64[]":  "Vector192",
		"uint8[]":  "Vector192",
		"uint16[]": "Vector192",
		"uint32[]": "Vector192",
		"uint64[]": "Vector192",
		"ptr64[]":  "Vector192",
		"float[]":  "Vector192",
		"double[]": "Vector192",
		"string[]": "Vector192",
		"any[]":    "Vector192",
		"vec2[]":   "Vector192",
		"vec3[]":   "Vector192",
		"vec4[]":   "Vector192",
		"mat4x4[]": "Vector192",
		"vec2":     "Vector2",
		"vec3":     "Vector3",
		"vec4":     "Vector4",
		"mat4x4":   "Matrix4x4",
	}
}

func initDataConverterMap() map[string]string {
	return map[string]string{
		"string":   "NativeMethods.GetStringData",
		"any":      "NativeMethods.GetVariantData",
		"bool[]":   "NativeMethods.GetVectorDataBool",
		"char8[]":  "NativeMethods.GetVectorDataChar8",
		"char16[]": "NativeMethods.GetVectorDataChar16",
		"int8[]":   "NativeMethods.GetVectorDataInt8",
		"int16[]":  "NativeMethods.GetVectorDataInt16",
		"int32[]":  "NativeMethods.GetVectorDataInt32",
		"int64[]":  "NativeMethods.GetVectorDataInt64",
		"uint8[]":  "NativeMethods.GetVectorDataUInt8",
		"uint16[]": "NativeMethods.GetVectorDataUInt16",
		"uint32[]": "NativeMethods.GetVectorDataUInt32",
		"uint64[]": "NativeMethods.GetVectorDataUInt64",
		"ptr64[]":  "NativeMethods.GetVectorDataIntPtr",
		"float[]":  "NativeMethods.GetVectorDataFloat",
		"double[]": "NativeMethods.GetVectorDataDouble",
		"string[]": "NativeMethods.GetVectorDataString",
		"any[]":    "NativeMethods.GetVectorDataVariant",
		"vec2[]":   "NativeMethods.GetVectorDataVector2",
		"vec3[]":   "NativeMethods.GetVectorDataVector3",
		"vec4[]":   "NativeMethods.GetVectorDataVector4",
		"mat4x4[]": "NativeMethods.GetVectorDataMatrix4x4",
	}
}

func initSizeFunctionMap() map[string]string {
	return map[string]string{
		"bool[]":   "NativeMethods.GetVectorSizeBool",
		"char8[]":  "NativeMethods.GetVectorSizeChar8",
		"char16[]": "NativeMethods.GetVectorSizeChar16",
		"int8[]":   "NativeMethods.GetVectorSizeInt8",
		"int16[]":  "NativeMethods.GetVectorSizeInt16",
		"int32[]":  "NativeMethods.GetVectorSizeInt32",
		"int64[]":  "NativeMethods.GetVectorSizeInt64",
		"uint8[]":  "NativeMethods.GetVectorSizeUInt8",
		"uint16[]": "NativeMethods.GetVectorSizeUInt16",
		"uint32[]": "NativeMethods.GetVectorSizeUInt32",
		"uint64[]": "NativeMethods.GetVectorSizeUInt64",
		"ptr64[]":  "NativeMethods.GetVectorSizeIntPtr",
		"float[]":  "NativeMethods.GetVectorSizeFloat",
		"double[]": "NativeMethods.GetVectorSizeDouble",
		"string[]": "NativeMethods.GetVectorSizeString",
		"any[]":    "NativeMethods.GetVectorSizeVariant",
		"vec2[]":   "NativeMethods.GetVectorSizeVector2",
		"vec3[]":   "NativeMethods.GetVectorSizeVector3",
		"vec4[]":   "NativeMethods.GetVectorSizeVector4",
		"mat4x4[]": "NativeMethods.GetVectorSizeMatrix4x4",
	}
}

func initDestructorMap() map[string]string {
	return map[string]string{
		"string":   "NativeMethods.DestroyString",
		"any":      "NativeMethods.DestroyVariant",
		"bool[]":   "NativeMethods.DestroyVectorBool",
		"char8[]":  "NativeMethods.DestroyVectorChar8",
		"char16[]": "NativeMethods.DestroyVectorChar16",
		"int8[]":   "NativeMethods.DestroyVectorInt8",
		"int16[]":  "NativeMethods.DestroyVectorInt16",
		"int32[]":  "NativeMethods.DestroyVectorInt32",
		"int64[]":  "NativeMethods.DestroyVectorInt64",
		"uint8[]":  "NativeMethods.DestroyVectorUInt8",
		"uint16[]": "NativeMethods.DestroyVectorUInt16",
		"uint32[]": "NativeMethods.DestroyVectorUInt32",
		"uint64[]": "NativeMethods.DestroyVectorUInt64",
		"ptr64[]":  "NativeMethods.DestroyVectorIntPtr",
		"float[]":  "NativeMethods.DestroyVectorFloat",
		"double[]": "NativeMethods.DestroyVectorDouble",
		"string[]": "NativeMethods.DestroyVectorString",
		"any[]":    "NativeMethods.DestroyVectorVariant",
		"vec2[]":   "NativeMethods.DestroyVectorVector2",
		"vec3[]":   "NativeMethods.DestroyVectorVector3",
		"vec4[]":   "NativeMethods.DestroyVectorVector4",
		"mat4x4[]": "NativeMethods.DestroyVectorMatrix4x4",
	}
}

func (m *DotnetTypeMapper) MapType(baseType string, context TypeContext, isArray bool) (string, error) {
	typeName := baseType
	if isArray {
		typeName = baseType + "[]"
	}

	mapped, ok := m.typesMap[typeName]
	if !ok {
		mapped, ok = m.typesMap[baseType]
		if !ok {
			// Custom type (enum or delegate)
			mapped = baseType
		}
		if isArray && context&TypeContextAlias == 0 {
			mapped = mapped + "[]"
		}
	}

	// Handle reference context (ref=true parameters)
	if context&TypeContextRef != 0 && baseType != "void" {
		mapped = "ref " + mapped
	}

	return mapped, nil
}

func (m *DotnetTypeMapper) MapParamType(param *manifest.ParamType) (string, error) {
	// Regular type mapping
	ctx := TypeContextValue
	if param.Ref {
		ctx = TypeContextRef
	}

	var typeName string
	switch {
	case param.Alias != nil:
		typeName = param.Alias.Name
		ctx |= TypeContextAlias

	case param.Enum != nil:
		typeName = param.Enum.Name

	case param.Prototype != nil:
		return param.Prototype.Name, nil

	default:
		typeName = param.BaseType()
	}

	return m.MapType(typeName, ctx, param.IsArray())
}

func (m *DotnetTypeMapper) MapReturnType(retType *manifest.RetType) (string, error) {
	ctx := TypeContextReturn

	var typeName string
	switch {
	case retType.Alias != nil:
		typeName = retType.Alias.Name
		ctx |= TypeContextAlias

	case retType.Enum != nil:
		typeName = retType.Enum.Name

	case retType.Prototype != nil:
		return retType.Prototype.Name, nil

	default:
		typeName = retType.BaseType()
	}

	return m.MapType(typeName, ctx, retType.IsArray())
}

func (m *DotnetTypeMapper) MapHandleType(class *manifest.Class) (string, string, error) {
	invalidValue := class.InvalidValue
	handleType, err := m.MapType(class.HandleType, TypeContextReturn, false)
	if err != nil {
		return "", "", err
	}

	nullptr := invalidValue == "0" || invalidValue == "" || invalidValue == "NULL" || invalidValue == "nullptr"
	if strings.HasPrefix(class.HandleType, "ptr") && nullptr {
		invalidValue = "nint.Zero"
	} else if invalidValue == "" {
		invalidValue = "default"
	}

	return invalidValue, handleType, err
}

// MapDelegateParamType maps parameter types for delegate definitions
func (m *DotnetTypeMapper) MapDelegateParamType(param *manifest.ParamType) (string, error) {
	// Delegates use ref for POD types automatically, enums only if ref=true
	if param.Enum != nil {
		typeName := param.Enum.Name
		if param.IsArray() {
			typeName = typeName + "[]"
		}
		// Enums only get ref if explicitly marked in manifest
		if param.Ref {
			return "ref " + typeName, nil
		}
		return typeName, nil
	}

	if param.Prototype != nil {
		return param.Prototype.Name, nil
	}

	typeName, err := m.MapType(param.BaseType(), TypeContextValue, param.IsArray())
	if err != nil {
		return "", err
	}

	// Add ref for:
	// 1. POD types (always, even when ref=false)
	// 2. Any type with ref=true in manifest
	if m.isPODType(param.Type) || param.Ref {
		return "ref " + typeName, nil
	}

	return typeName, nil
}

// MapDelegateReturnType maps return type for delegate definitions
func (m *DotnetTypeMapper) MapDelegateReturnType(retType *manifest.RetType) (string, error) {
	// Same as regular return type
	return m.MapReturnType(retType)
}

// MapUnmanagedParamType maps parameters for unmanaged function pointer declarations
func (m *DotnetTypeMapper) MapUnmanagedParamType(param *manifest.ParamType) (string, error) {
	// Check for enum
	if param.Enum != nil && !param.IsArray() {
		typeName := param.Enum.Name
		if param.Ref {
			return typeName + "*", nil
		}
		return typeName, nil
	}

	// Check for function pointer
	if param.Prototype != nil {
		return "nint", nil
	}

	// Get C type
	typeName := param.Type
	if param.IsArray() {
		typeName = typeName + "[]"
	}

	cType, ok := m.ctypesMap[typeName]
	if !ok {
		cType, ok = m.ctypesMap[param.BaseType()]
		if !ok {
			cType = param.BaseType()
		}
		if param.IsArray() {
			cType = "Vector192*"
		}
	}

	// Add pointer for ref parameters
	if param.Ref && !strings.Contains(cType, "*") {
		cType = cType + "*"
	}

	return cType, nil
}

// MapUnmanagedReturnType maps return type for unmanaged function pointer declarations
func (m *DotnetTypeMapper) MapUnmanagedReturnType(retType *manifest.RetType) (string, error) {
	// Check for enum
	if retType.Enum != nil && !retType.IsArray() {
		return retType.Enum.Name, nil
	}

	// Check for function pointer
	if retType.Prototype != nil {
		return "nint", nil
	}

	typeName := retType.Type
	if retType.IsArray() {
		typeName = typeName + "[]"
	}

	cType, ok := m.ctypesMap[typeName]
	if !ok {
		cType, ok = m.ctypesMap[retType.BaseType()]
		if !ok {
			cType = retType.BaseType()
		}
		if retType.IsArray() {
			cType = "Vector192*"
		}
	}

	// Remove pointer for return types
	if strings.HasSuffix(cType, "*") {
		cType = cType[:len(cType)-1]
	}

	return cType, nil
}

// Helper methods

func (m *DotnetTypeMapper) isObjectReturn(typeName string) bool {
	return strings.Contains(typeName, "[]") || typeName == "string" || typeName == "any"
}

func (m *DotnetTypeMapper) isPODType(typeName string) bool {
	return typeName == "vec2" || typeName == "vec3" || typeName == "vec4" || typeName == "mat4x4"
}

func (m *DotnetTypeMapper) isFunction(typeName string) bool {
	return typeName == "function"
}

func (m *DotnetTypeMapper) getConstructor(typeName string) string {
	if strings.Contains(typeName, "[]") {
		return m.constructorMap[typeName]
	}
	return m.constructorMap[typeName]
}

func (m *DotnetTypeMapper) getNativeReturnType(typeName string) string {
	if strings.Contains(typeName, "[]") {
		return m.returnTypeMap[typeName]
	}
	return m.returnTypeMap[typeName]
}

func (m *DotnetTypeMapper) getDataConverter(typeName string) string {
	if strings.Contains(typeName, "[]") {
		return m.dataConverterMap[typeName]
	}
	return m.dataConverterMap[typeName]
}

func (m *DotnetTypeMapper) getSizeFunction(typeName string) string {
	if strings.Contains(typeName, "[]") {
		return m.sizeFunctionMap[typeName]
	}
	return ""
}

func (m *DotnetTypeMapper) getDestructor(typeName string) string {
	if strings.Contains(typeName, "[]") {
		return m.destructorMap[typeName]
	}
	return m.destructorMap[typeName]
}
