package generator

import (
	"fmt"
	"strings"

	"github.com/untrustedmodders/plugify-gen/pkg/manifest"
)

// RustGenerator generates Rust bindings
type RustGenerator struct {
	*BaseGenerator
}

// NewRustGenerator creates a new Rust generator
func NewRustGenerator() *RustGenerator {
	return &RustGenerator{
		BaseGenerator: NewBaseGenerator("rust", NewRustTypeMapper(), RustReservedWords),
	}
}

// Generate generates Rust bindings
func (g *RustGenerator) Generate(m *manifest.Manifest, opts *GeneratorOptions) (*GeneratorResult, error) {
	g.ResetCaches()
	opts = EnsureOptions(opts)

	// Collect all unique groups from both methods and classes
	groups := g.GetGroups(m)

	files := make(map[string]string)
	folder := fmt.Sprintf("src/%s", m.Name)

	// Generate separate enums file
	enumsCode, err := g.generateEnumsFile(m)
	if err != nil {
		return nil, fmt.Errorf("generating enums file: %w", err)
	}
	files[fmt.Sprintf("%s/enums.rs", folder)] = enumsCode

	// Generate separate delegates file
	delegatesCode, err := g.generateDelegatesFile(m)
	if err != nil {
		return nil, fmt.Errorf("generating delegates file: %w", err)
	}
	files[fmt.Sprintf("%s/delegates.rs", folder)] = delegatesCode

	// Generate a file for each group
	for groupName := range groups {
		groupCode, err := g.generateGroupFile(m, groupName, opts)
		if err != nil {
			return nil, fmt.Errorf("failed to generate group %s: %w", groupName, err)
		}
		files[fmt.Sprintf("%s/%s.rs", folder, groupName)] = groupCode
	}

	// Generate lib.rs that re-exports all pieces
	libRs, err := g.generateLibFile(m, groups)
	if err != nil {
		return nil, fmt.Errorf("generating lib.rs: %w", err)
	}
	files[fmt.Sprintf("%s/lib.rs", folder)] = libRs

	result := &GeneratorResult{
		Files: files,
	}

	return result, nil
}

func (g *RustGenerator) generateEnums(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	for _, method := range m.Methods {
		// Check return type
		if method.RetType.Enum != nil && !g.IsEnumCached(method.RetType.Enum.Name) {
			enumCode := g.generateEnum(method.RetType.Enum, method.RetType.Type)
			sb.WriteString(enumCode)
			sb.WriteString("\n")
			g.CacheEnum(method.RetType.Enum.Name)
		}

		// Check return type prototype
		if method.RetType.Prototype != nil {
			g.processPrototypeEnums(method.RetType.Prototype, &sb)
		}

		// Check parameters
		for _, param := range method.ParamTypes {
			if param.Enum != nil && !g.IsEnumCached(param.Enum.Name) {
				enumCode := g.generateEnum(param.Enum, param.Type)
				sb.WriteString(enumCode)
				sb.WriteString("\n")
				g.CacheEnum(param.Enum.Name)
			}

			// Check nested prototypes
			if param.Prototype != nil {
				g.processPrototypeEnums(param.Prototype, &sb)
			}
		}
	}

	return sb.String(), nil
}

func (g *RustGenerator) processPrototypeEnums(proto *manifest.Prototype, sb *strings.Builder) {
	if proto.RetType.Enum != nil && !g.IsEnumCached(proto.RetType.Enum.Name) {
		enumCode := g.generateEnum(proto.RetType.Enum, proto.RetType.Type)
		sb.WriteString(enumCode)
		sb.WriteString("\n")
		g.CacheEnum(proto.RetType.Enum.Name)
	}

	for _, param := range proto.ParamTypes {
		if param.Enum != nil && !g.IsEnumCached(param.Enum.Name) {
			enumCode := g.generateEnum(param.Enum, param.Type)
			sb.WriteString(enumCode)
			sb.WriteString("\n")
			g.CacheEnum(param.Enum.Name)
		}
		if param.Prototype != nil {
			g.processPrototypeEnums(param.Prototype, sb)
		}
	}
}

func (g *RustGenerator) generateEnum(enum *manifest.EnumType, enumType string) string {
	var sb strings.Builder

	if enum.Description != "" {
		sb.WriteString(fmt.Sprintf("/// %s\n", enum.Description))
	}

	underlyingType, _ := g.typeMapper.MapType(enumType, TypeContextReturn, false)

	sb.WriteString("#[repr(")
	sb.WriteString(underlyingType)
	sb.WriteString(")]\n")
	sb.WriteString("#[derive(Debug, Clone, Copy, PartialEq, Eq)]\n")
	sb.WriteString(fmt.Sprintf("pub enum %s {\n", enum.Name))

	for _, val := range enum.Values {
		if val.Description != "" {
			sb.WriteString(fmt.Sprintf("    /// %s\n", val.Description))
		}
		sb.WriteString(fmt.Sprintf("    %s = %d,\n", val.Name, val.Value))
	}

	sb.WriteString("}\n")
	return sb.String()
}

func (g *RustGenerator) generateDelegates(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	for _, method := range m.Methods {
		// Check return type
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

	return sb.String(), nil
}

func (g *RustGenerator) generateDelegate(proto *manifest.Prototype) (string, error) {
	var sb strings.Builder

	if proto.Description != "" {
		sb.WriteString(fmt.Sprintf("/// %s\n", proto.Description))
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

	sb.WriteString(fmt.Sprintf("pub type %s = unsafe extern \"C\" fn(", proto.Name))
	sb.WriteString(params)
	sb.WriteString(")")
	if retType != "" && retType != "()" {
		sb.WriteString(" -> ")
		sb.WriteString(retType)
	}
	sb.WriteString(";\n\n")

	return sb.String(), nil
}

func (g *RustGenerator) generateMethod(pluginName string, method *manifest.Method) (string, error) {
	var sb strings.Builder

	// Generate documentation comment
	if method.Description != "" {
		sb.WriteString(fmt.Sprintf("/// %s\n", method.Description))
		sb.WriteString("///\n")
	}
	for _, param := range method.ParamTypes {
		paramType := param.Type
		if param.Ref {
			paramType += "&"
		}
		sb.WriteString(fmt.Sprintf("/// # Arguments\n"))
		sb.WriteString(fmt.Sprintf("/// * `%s` - (%s)", g.SanitizeName(param.Name), paramType))
		if param.Description != "" {
			sb.WriteString(fmt.Sprintf(": %s", param.Description))
		}
		sb.WriteString("\n")
	}
	if method.RetType.Type != "void" {
		sb.WriteString("///\n")
		sb.WriteString("/// # Returns\n")
		sb.WriteString(fmt.Sprintf("/// %s", method.RetType.Type))
		if method.RetType.Description != "" {
			sb.WriteString(fmt.Sprintf(": %s", method.RetType.Description))
		}
		sb.WriteString("\n")
	}

	// Generate function signature
	retType, err := g.typeMapper.MapReturnType(&method.RetType)
	if err != nil {
		return "", err
	}

	params, err := FormatParameters(method.ParamTypes, ParamFormatTypesAndNames, g.typeMapper, g.SanitizeName)
	if err != nil {
		return "", err
	}

	sb.WriteString("#[inline]\n")
	sb.WriteString(fmt.Sprintf("pub fn %s(%s)", g.SanitizeName(method.Name), params))
	if retType != "" && retType != "()" {
		sb.WriteString(" -> ")
		sb.WriteString(retType)
	}
	sb.WriteString(" {\n")

	// Generate function body using lazy static pattern
	funcTypeParams, err := FormatParameters(method.ParamTypes, ParamFormatTypes, g.typeMapper, g.SanitizeName)
	if err != nil {
		return "", err
	}

	sb.WriteString("    unsafe {\n")
	sb.WriteString(fmt.Sprintf("        static FUNC: OnceLock<unsafe extern \"C\" fn(%s)", funcTypeParams))
	if retType != "" && retType != "()" {
		sb.WriteString(" -> ")
		sb.WriteString(retType)
	}
	sb.WriteString("> = OnceLock::new();\n")

	sb.WriteString("        let func = FUNC.get_or_init(|| {\n")
	sb.WriteString(fmt.Sprintf("            let name = \"%s.%s\";\n", pluginName, method.Name))
	sb.WriteString("            let ptr = init_get_method_ptr(name.as_ptr(), name.len());\n")
	sb.WriteString("            std::mem::transmute(ptr)\n")
	sb.WriteString("        });\n")

	// Call function
	paramNames, err := FormatParameters(method.ParamTypes, ParamFormatNames, g.typeMapper, g.SanitizeName)
	if err != nil {
		return "", err
	}

	if method.RetType.Type == "void" {
		sb.WriteString(fmt.Sprintf("        func(%s);\n", paramNames))
	} else {
		sb.WriteString(fmt.Sprintf("        func(%s)\n", paramNames))
	}

	sb.WriteString("    };\n")
	sb.WriteString("}\n")

	return sb.String(), nil
}

// generateEnumsFile generates a file containing all enums
func (g *RustGenerator) generateEnumsFile(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("// Generated from %s.pplugin\n\n", m.Name))

	// Generate enums
	enumsCode, err := g.generateEnums(m)
	if err != nil {
		return "", err
	}
	if enumsCode != "" {
		sb.WriteString(enumsCode)
		sb.WriteString("\n")
	}

	// Ownership enum (if any class has destructor)
	sb.WriteString("/// Ownership type for RAII wrappers\n")
	sb.WriteString("#[derive(Debug, Clone, Copy, PartialEq, Eq)]\n")
	sb.WriteString("pub enum Ownership {\n")
	sb.WriteString("    Borrowed,\n")
	sb.WriteString("    Owned,\n")
	sb.WriteString("}\n")

	return sb.String(), nil
}

// generateDelegatesFile generates a file containing all delegate type definitions
func (g *RustGenerator) generateDelegatesFile(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("// Generated from %s.pplugin\n\n", m.Name))
	sb.WriteString("use super::enums::*;\n\n")

	// Generate delegates
	delegatesCode, err := g.generateDelegates(m)
	if err != nil {
		return "", err
	}
	if delegatesCode != "" {
		sb.WriteString(delegatesCode)
	}

	return sb.String(), nil
}

// generateGroupFile generates a file for a specific group containing methods
func (g *RustGenerator) generateGroupFile(m *manifest.Manifest, groupName string, opts *GeneratorOptions) (string, error) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("// Generated from %s.pplugin (group: %s)\n\n", m.Name, groupName))
	sb.WriteString("use super::enums::*;\n")
	sb.WriteString("use super::delegates::*;\n\n")

	// External function declaration for getting method pointers
	sb.WriteString("extern \"C\" {\n")
	sb.WriteString("    fn plugify_get_method_ptr2(export_name: *const i8, out_func_ptr: *mut *mut std::ffi::c_void);\n")
	sb.WriteString("}\n\n")

	// Generate methods for this group
	for _, method := range m.Methods {
		methodGroup := g.GetGroupName(method.Group)
		if methodGroup == groupName {
			methodCode, err := g.generateMethod(m.Name, &method)
			if err != nil {
				return "", fmt.Errorf("failed to generate method %s: %w", method.Name, err)
			}
			sb.WriteString(methodCode)
			sb.WriteString("\n")
		}
	}

	// TODO: Generate classes for this group (if enabled) in future PR
	// if opts.GenerateClasses {
	//     // Class generation will be implemented later
	// }

	return sb.String(), nil
}

// generateLibFile generates the lib.rs file that re-exports all modules
func (g *RustGenerator) generateLibFile(m *manifest.Manifest, groups map[string]bool) (string, error) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("// Generated from %s.pplugin\n", m.Name))
	sb.WriteString("// This module re-exports all generated components\n\n")

	// Declare modules
	sb.WriteString("pub mod enums;\n")
	sb.WriteString("pub mod delegates;\n")
	for groupName := range groups {
		sb.WriteString(fmt.Sprintf("pub mod %s;\n", groupName))
	}
	sb.WriteString("\n")

	// Re-export everything
	sb.WriteString("pub use enums::*;\n")
	sb.WriteString("pub use delegates::*;\n")
	for groupName := range groups {
		sb.WriteString(fmt.Sprintf("pub use %s::*;\n", groupName))
	}

	return sb.String(), nil
}

// RustTypeMapper implements type mapping for Rust
type RustTypeMapper struct{}

func NewRustTypeMapper() *RustTypeMapper {
	return &RustTypeMapper{}
}

func (m *RustTypeMapper) MapType(baseType string, context TypeContext, isArray bool) (string, error) {
	// Base type mapping
	typeMap := map[string]string{
		"void":   "",
		"bool":   "bool",
		"char8":  "i8",
		"char16": "u16",
		"int8":   "i8",
		"int16":  "i16",
		"int32":  "i32",
		"int64":  "i64",
		"uint8":  "u8",
		"uint16": "u16",
		"uint32": "u32",
		"uint64": "u64",
		"ptr64":  "usize",
		"float":  "f32",
		"double": "f64",
		"string": "PlgString",
		"any":    "PlgVariant",
		"vec2":   "Vector2",
		"vec3":   "Vector3",
		"vec4":   "Vector4",
		"mat4x4": "Matrix4x4",
	}

	mapped, ok := typeMap[baseType]
	if !ok {
		// Assume it's a custom type (enum or delegate)
		mapped = baseType
	}

	// Handle void -> ()
	if mapped == "" {
		mapped = "()"
	}

	// Handle arrays
	if isArray {
		mapped = fmt.Sprintf("PlgVector<%s>", mapped)
	}

	// Handle reference context (ref=true parameters)
	if context == TypeContextRef && baseType != "void" {
		mapped = fmt.Sprintf("&mut %s", mapped)
	}

	return mapped, nil
}

func (m *RustTypeMapper) MapParamType(param *manifest.ParamType, context TypeContext) (string, error) {
	// Check for enum type first
	if param.Enum != nil {
		typeName := param.Enum.Name
		if param.IsArray() {
			typeName = fmt.Sprintf("PlgVector<%s>", typeName)
		}
		// Handle reference
		if param.Ref && !param.IsArray() {
			return fmt.Sprintf("&mut %s", typeName), nil
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

func (m *RustTypeMapper) MapReturnType(retType *manifest.TypeInfo) (string, error) {
	// Check for enum type
	if retType.Enum != nil {
		typeName := retType.Enum.Name
		if retType.IsArray() {
			typeName = fmt.Sprintf("PlgVector<%s>", typeName)
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

func (m *RustTypeMapper) MapHandleType(class *manifest.Class) (string, string, error) {
	invalidValue := class.InvalidValue
	handleType, err := m.MapType(class.HandleType, TypeContextReturn, false)
	if err != nil {
		return "", "", err
	}

	if class.HandleType == "ptr64" && invalidValue == "0" {
		invalidValue = "0"
	} else if invalidValue == "" {
		invalidValue = "Default::default()"
	}

	return invalidValue, handleType, nil
}
