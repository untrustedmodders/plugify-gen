package generator

import (
	"fmt"
	"strings"

	"github.com/untrustedmodders/plugify-gen/pkg/manifest"
)

// DlangGenerator generates D language bindings
type DlangGenerator struct {
	*BaseGenerator
}

// NewDlangGenerator creates a new D language generator
func NewDlangGenerator() *DlangGenerator {
	invalidNames := []string{
		"out", "version", "module", "function", "body", "in", "ref",
		"abstract", "alias", "align", "asm", "assert", "auto",
		"bool", "break", "byte", "case", "cast", "catch", "cdouble",
		"cent", "cfloat", "char", "class", "const", "continue", "creal",
		"dchar", "debug", "default", "delegate", "delete", "deprecated",
		"do", "double", "else", "enum", "export", "extern", "false",
		"final", "finally", "float", "for", "foreach", "foreach_reverse",
		"goto", "idouble", "if", "ifloat", "immutable", "import",
		"inout", "int", "interface", "invariant", "ireal", "is",
		"lazy", "long", "macro", "mixin", "new", "nothrow", "null",
		"override", "package", "pragma", "private", "protected", "public",
		"pure", "real", "return", "scope", "shared", "short", "static",
		"struct", "super", "switch", "synchronized", "template", "this",
		"throw", "true", "try", "typedef", "typeid", "typeof", "ubyte",
		"ucent", "uint", "ulong", "union", "unittest", "ushort", "void",
		"volatile", "wchar", "while", "with",
	}

	return &DlangGenerator{
		BaseGenerator: NewBaseGenerator("dlang", NewDlangTypeMapper(), invalidNames),
	}
}

// Generate generates D language bindings
func (g *DlangGenerator) Generate(m *manifest.Manifest) (*GeneratorResult, error) {
	g.ResetCaches()

	var sb strings.Builder

	// Module declaration
	moduleName := strings.ToLower(m.Name)
	for _, method := range m.Methods {
		groupName := strings.ToLower(method.Group)
		if groupName == "" {
			continue
		}

		// Generate one file per module/group
		moduleCode, err := g.generateModule(m, groupName)
		if err != nil {
			return nil, err
		}

		// Store the generated code (we'll collect all groups first)
		sb.WriteString(moduleCode)
	}

	// Collect all unique groups
	groups := make(map[string]bool)
	for _, method := range m.Methods {
		groupName := strings.ToLower(method.Group)
		if groupName != "" {
			groups[groupName] = true
		}
	}

	// Generate files for each group
	files := make(map[string]string)

	// First, generate enums file
	enumsCode, err := g.generateEnumsFile(m)
	if err != nil {
		return nil, err
	}
	files[fmt.Sprintf("source/imported/%s/enums.d", moduleName)] = enumsCode

	// Generate a file for each group
	for groupName := range groups {
		groupCode, err := g.generateModuleFile(m, moduleName, groupName)
		if err != nil {
			return nil, err
		}
		files[fmt.Sprintf("source/imported/%s/%s.d", moduleName, groupName)] = groupCode
	}

	result := &GeneratorResult{
		Files: files,
	}

	return result, nil
}

func (g *DlangGenerator) generateEnumsFile(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	moduleName := strings.ToLower(m.Name)
	sb.WriteString(fmt.Sprintf("module imported.%s.enums;\n\n", moduleName))

	// Collect all enums
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

func (g *DlangGenerator) processPrototypeEnums(proto *manifest.Prototype, sb *strings.Builder) {
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

func (g *DlangGenerator) generateEnum(enum *manifest.EnumType) string {
	var sb strings.Builder

	if enum.Description != "" {
		sb.WriteString(fmt.Sprintf("/// %s\n", enum.Description))
	}

	underlyingType := "int"
	if enum.Type != "" {
		mapped, _ := g.typeMapper.MapType(enum.Type, TypeContextValue, false)
		underlyingType = mapped
	}

	sb.WriteString(fmt.Sprintf("enum %s : %s {\n", enum.Name, underlyingType))

	for _, val := range enum.Values {
		// Format the value with underscores for readability (like the legacy version)
		valueStr := fmt.Sprintf("%d", val.Value)
		if len(valueStr) > 3 {
			// Add underscores for thousands separator
			var formatted strings.Builder
			for i, c := range valueStr {
				if i > 0 && (len(valueStr)-i)%3 == 0 {
					formatted.WriteRune('_')
				}
				formatted.WriteRune(c)
			}
			valueStr = formatted.String()
		}

		sb.WriteString(fmt.Sprintf("\t%s = %s,", val.Name, valueStr))
		if val.Description != "" {
			sb.WriteString(fmt.Sprintf(" /// %s", val.Description))
		}
		sb.WriteString("\n")
	}

	sb.WriteString("}\n")
	return sb.String()
}

func (g *DlangGenerator) generateModuleFile(m *manifest.Manifest, moduleName, groupName string) (string, error) {
	var sb strings.Builder

	// Module declaration
	sb.WriteString(fmt.Sprintf("module imported.%s.%s;\n\n", moduleName, groupName))

	// Imports
	sb.WriteString("import plugify.internals;\n")
	sb.WriteString("public import plugify;\n")
	sb.WriteString(fmt.Sprintf("public import imported.%s.enums;\n\n", moduleName))

	// Collect methods and delegates for this group
	var methods []*manifest.Method
	delegates := make(map[string]*manifest.Prototype)

	for i := range m.Methods {
		method := &m.Methods[i]
		if strings.ToLower(method.Group) == groupName {
			methods = append(methods, method)

			// Collect delegates from parameters
			for _, param := range method.ParamTypes {
				if param.Prototype != nil {
					delegates[param.Prototype.Name] = param.Prototype
				}
			}
			// Collect delegates from return type
			if method.RetType.Prototype != nil {
				delegates[method.RetType.Prototype.Name] = method.RetType.Prototype
			}
		}
	}

	// Generate delegate aliases
	for _, proto := range delegates {
		delegateCode, err := g.generateDelegate(proto)
		if err != nil {
			return "", err
		}
		sb.WriteString(delegateCode)
		sb.WriteString("\n")
	}

	// Generate wrapper functions
	for _, method := range methods {
		wrapperCode, err := g.generateMethodWrapper(method)
		if err != nil {
			return "", fmt.Errorf("failed to generate method %s: %w", method.Name, err)
		}
		sb.WriteString(wrapperCode)
		sb.WriteString("\n")
	}

	// Generate private section with function pointers
	sb.WriteString("private {\n")
	for _, method := range methods {
		aliasCode, err := g.generateMethodAlias(method)
		if err != nil {
			return "", err
		}
		sb.WriteString(aliasCode)
	}
	sb.WriteString("}\n\n")

	// Generate static initialization
	sb.WriteString("shared static this() {\n")
	for _, method := range methods {
		sb.WriteString(fmt.Sprintf("\t__%s = cast(_%s)_MethodPointer(\"%s\");\n",
			method.Name, method.Name, method.FuncName))
	}
	sb.WriteString("}\n")

	return sb.String(), nil
}

func (g *DlangGenerator) generateDelegate(proto *manifest.Prototype) (string, error) {
	var sb strings.Builder

	sb.WriteString("/++\n")
	if proto.Description != "" {
		sb.WriteString(fmt.Sprintf("\t%s\n\n", proto.Description))
	}

	// Generate parameter documentation
	if len(proto.ParamTypes) > 0 {
		sb.WriteString("\tParams:\n")
		for _, param := range proto.ParamTypes {
			paramName := g.SanitizeName(param.Name)
			sb.WriteString(fmt.Sprintf("\t\t%s = %s\n", paramName, param.Description))
		}
	}
	sb.WriteString("+/\n")

	// Generate return type
	retType, err := g.typeMapper.MapReturnType(&proto.RetType)
	if err != nil {
		return "", err
	}

	// Convert to C ABI type
	cRetType, err := g.toCType(retType, &proto.RetType, false)
	if err != nil {
		return "", err
	}

	sb.WriteString(fmt.Sprintf("alias %s = extern (C) %s function(", proto.Name, cRetType))

	// Generate parameters
	params := []string{}
	for _, param := range proto.ParamTypes {
		paramName := g.SanitizeName(param.Name)
		paramType, err := g.typeMapper.MapParamType(&param, TypeContextValue)
		if err != nil {
			return "", err
		}

		// Convert to C ABI type
		cType, err := g.toCType(paramType, &manifest.TypeInfo{
			Type: param.Type,
			Enum: param.Enum,
		}, param.Ref)
		if err != nil {
			return "", err
		}

		if !param.Ref && strings.HasPrefix(cType, "ref") {
			cType = "const " + cType
		}

		params = append(params, fmt.Sprintf("%s %s", cType, paramName))
	}

	sb.WriteString(strings.Join(params, ", "))
	sb.WriteString(");\n")

	return sb.String(), nil
}

func (g *DlangGenerator) generateMethodWrapper(method *manifest.Method) (string, error) {
	var sb strings.Builder

	// Documentation
	sb.WriteString("/++\n")
	if method.Description != "" {
		sb.WriteString(fmt.Sprintf("\t%s\n\n", method.Description))
	}

	// Parameters documentation
	if len(method.ParamTypes) > 0 {
		sb.WriteString("\tParams:\n")
		for _, param := range method.ParamTypes {
			paramName := g.SanitizeName(param.Name)
			sb.WriteString(fmt.Sprintf("\t\t%s = %s\n", paramName, param.Description))
		}
	}

	// Return documentation
	if method.RetType.Type != "void" {
		sb.WriteString("\tReturns:\n")
		sb.WriteString(fmt.Sprintf("\t\t%s\n", method.RetType.Description))
	}
	sb.WriteString("+/\n")

	// Function signature
	retType, err := g.typeMapper.MapReturnType(&method.RetType)
	if err != nil {
		return "", err
	}

	sb.WriteString(fmt.Sprintf("%s %s(", retType, method.Name))

	// Parameters
	params := []string{}
	for _, param := range method.ParamTypes {
		paramName := g.SanitizeName(param.Name)
		paramType, err := g.typeMapper.MapParamType(&param, TypeContextValue)
		if err != nil {
			return "", err
		}

		if param.Ref {
			params = append(params, fmt.Sprintf("ref %s %s", paramType, paramName))
		} else {
			params = append(params, fmt.Sprintf("%s %s", paramType, paramName))
		}
	}

	sb.WriteString(strings.Join(params, ", "))
	sb.WriteString(") {\n")

	// Function body - handle type conversions
	var conversions strings.Builder
	var callArgs []string
	var cleanups strings.Builder

	for i, param := range method.ParamTypes {
		paramName := g.SanitizeName(param.Name)
		paramType, _ := g.typeMapper.MapParamType(&param, TypeContextValue)

		// Check if we need conversion
		cType, _ := g.toCType(paramType, &manifest.TypeInfo{
			Type: param.Type,
			Enum: param.Enum,
		}, param.Ref)

		if strings.HasPrefix(cType, "ref") && (strings.Contains(cType, "PlgA") || strings.Contains(cType, "PlgS")) {
			// Need conversion for arrays and strings
			varName := fmt.Sprintf("_var%d", i)
			valType := strings.TrimPrefix(cType, "ref ")
			conversions.WriteString(fmt.Sprintf("\t%s %s = (%s);\n", valType, varName, paramName))

			if param.Ref {
				cleanups.WriteString(fmt.Sprintf("\tscope(exit) { %s = %s; }\n", paramName, varName))
			}
			conversions.WriteString("\n")
			callArgs = append(callArgs, varName)
		} else {
			callArgs = append(callArgs, paramName)
		}
	}

	// Write conversions
	sb.WriteString(conversions.String())

	// Write function call
	sb.WriteString("\t")

	// Check if return type needs conversion
	cRetType, _ := g.toCType(retType, &method.RetType, false)
	needsReturnConversion := strings.Contains(cRetType, "!")

	if method.RetType.Type != "void" {
		sb.WriteString("return ")
	}

	sb.WriteString(fmt.Sprintf("__%s(", method.Name))
	sb.WriteString(strings.Join(callArgs, ", "))
	sb.WriteString(")")

	if needsReturnConversion {
		sb.WriteString(".value")
	}

	sb.WriteString(";\n")

	// Write cleanups (scope(exit) statements were already written before)

	sb.WriteString("}\n")

	return sb.String(), nil
}

func (g *DlangGenerator) generateMethodAlias(method *manifest.Method) (string, error) {
	var sb strings.Builder

	// Get C ABI return type
	retType, _ := g.typeMapper.MapReturnType(&method.RetType)
	cRetType, _ := g.toCType(retType, &method.RetType, false)

	// Remove 'ref' prefix if present
	if strings.HasPrefix(cRetType, "ref ") {
		cRetType = strings.TrimPrefix(cRetType, "ref ")
	}

	sb.WriteString(fmt.Sprintf("\talias _%s = extern (C) %s function(", method.Name, cRetType))

	// Parameters
	params := []string{}
	for _, param := range method.ParamTypes {
		paramName := g.SanitizeName(param.Name)
		paramType, _ := g.typeMapper.MapParamType(&param, TypeContextValue)

		cType, _ := g.toCType(paramType, &manifest.TypeInfo{
			Type: param.Type,
			Enum: param.Enum,
		}, param.Ref)

		if strings.HasPrefix(cType, "ref") {
			if !param.Ref {
				cType = "const " + cType
			}
		} else if param.Ref {
			cType = "ref " + cType
		}

		// Handle function types
		if param.Type == "function" && param.Prototype != nil {
			cType = param.Prototype.Name
		}

		params = append(params, fmt.Sprintf("%s %s", cType, paramName))
	}

	sb.WriteString(strings.Join(params, ", "))
	sb.WriteString(");\n")

	sb.WriteString(fmt.Sprintf("\t__gshared _%s __%s;\n", method.Name, method.Name))

	return sb.String(), nil
}

// toCType converts D native types to C ABI types
func (g *DlangGenerator) toCType(nativeType string, typeInfo *manifest.TypeInfo, isRef bool) (string, error) {
	// Map native types to C types
	typeMap := map[string]string{
		"void":     "void",
		"bool":     "bool",
		"char":     "char",
		"wchar":    "wchar",
		"byte":     "byte",
		"short":    "short",
		"int":      "int",
		"long":     "long",
		"ubyte":    "ubyte",
		"ushort":   "ushort",
		"uint":     "uint",
		"ulong":    "ulong",
		"void*":    "void*",
		"float":    "float",
		"double":   "double",
		"function": "function",
		"string":   "ref PlgS",
		"PlgV":     "ref PlgV",
	}

	// Check for array types
	if strings.HasSuffix(nativeType, "[]") {
		baseType := strings.TrimSuffix(nativeType, "[]")
		return fmt.Sprintf("ref PlgA!%s", baseType), nil
	}

	// Check for vector types
	vectorTypes := map[string]string{
		"Vec2":    "ref Vec2",
		"Vec3":    "ref Vec3",
		"Vec4":    "ref Vec4",
		"Mat4x4":  "ref Mat4x4",
	}

	if cType, ok := vectorTypes[nativeType]; ok {
		return cType, nil
	}

	// Check type map
	if cType, ok := typeMap[nativeType]; ok {
		return cType, nil
	}

	// If it's an enum, return as-is
	if typeInfo != nil && typeInfo.Enum != nil {
		return nativeType, nil
	}

	// Default: assume it's a custom type
	return nativeType, nil
}

func (g *DlangGenerator) generateModule(m *manifest.Manifest, groupName string) (string, error) {
	// This is now handled by generateModuleFile
	return "", nil
}

// DlangTypeMapper implements type mapping for D language
type DlangTypeMapper struct{}

func NewDlangTypeMapper() *DlangTypeMapper {
	return &DlangTypeMapper{}
}

func (m *DlangTypeMapper) MapType(baseType string, context TypeContext, isArray bool) (string, error) {
	typeMap := map[string]string{
		"void":   "void",
		"bool":   "bool",
		"char8":  "char",
		"char16": "wchar",
		"int8":   "byte",
		"int16":  "short",
		"int32":  "int",
		"int64":  "long",
		"uint8":  "ubyte",
		"uint16": "ushort",
		"uint32": "uint",
		"uint64": "ulong",
		"ptr64":  "void*",
		"float":  "float",
		"double": "double",
		"string": "string",
		"any":    "PlgV",
		"vec2":   "Vec2",
		"vec3":   "Vec3",
		"vec4":   "Vec4",
		"mat4x4": "Mat4x4",
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

func (m *DlangTypeMapper) MapParamType(param *manifest.ParamType, context TypeContext) (string, error) {
	// Check for enum type first
	if param.Enum != nil {
		typeName := param.Enum.Name
		if param.IsArray() {
			typeName = typeName + "[]"
		}
		return typeName, nil
	}

	// Check for function/delegate type
	if param.Prototype != nil {
		return param.Prototype.Name, nil
	}

	// Regular type mapping
	return m.MapType(param.BaseType(), context, param.IsArray())
}

func (m *DlangTypeMapper) MapReturnType(retType *manifest.TypeInfo) (string, error) {
	// Check for enum type
	if retType.Enum != nil {
		typeName := retType.Enum.Name
		if retType.IsArray() {
			typeName = typeName + "[]"
		}
		return typeName, nil
	}

	// Check for function/delegate type
	if retType.Prototype != nil {
		return retType.Prototype.Name, nil
	}

	// Regular type mapping
	return m.MapType(retType.BaseType(), TypeContextReturn, retType.IsArray())
}