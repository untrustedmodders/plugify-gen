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

	// Generate mod.rs that re-exports all pieces
	modRs, err := g.generateModFile(m, groups)
	if err != nil {
		return nil, fmt.Errorf("generating mod.rs: %w", err)
	}
	files[fmt.Sprintf("%s/mod.rs", folder)] = modRs

	result := &GeneratorResult{
		Files: files,
	}

	return result, nil
}

func (g *RustGenerator) generateEnums(m *manifest.Manifest) (string, error) {
	return g.CollectEnums(m, g.generateEnum)
}

func (g *RustGenerator) generateEnum(enum *manifest.EnumType, underlyingType string) (string, error) {
	var sb strings.Builder

	if enum.Description != "" {
		sb.WriteString(fmt.Sprintf("/// %s\n", enum.Description))
	}

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
	sb.WriteString(fmt.Sprintf("vector_enum_traits!(%s, %s);\n", enum.Name, underlyingType))
	return sb.String(), nil
}

func (g *RustGenerator) generateDelegates(m *manifest.Manifest) (string, error) {
	return g.CollectDelegates(m, g.generateDelegate)
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

	// Generate parameters in Rust format (name: type)
	params, err := g.formatRustParams(proto.ParamTypes, false)
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

// formatRustParams formats parameters in Rust style (name: type)
func (g *RustGenerator) formatRustParams(params []manifest.ParamType, includeNames bool) (string, error) {
	if len(params) == 0 {
		return "", nil
	}

	result := ""
	for i, param := range params {
		if i > 0 {
			result += ", "
		}

		// Get the type
		typeName, err := g.typeMapper.MapParamType(&param, TypeContextValue)
		if err != nil {
			return "", err
		}

		if includeNames {
			// Rust format: name: type
			paramName := param.Name
			if paramName == "" {
				paramName = fmt.Sprintf("p%d", i)
			}
			result += g.SanitizeName(paramName) + ": " + typeName
		} else {
			// Just type (for extern "C" fn signatures)
			result += typeName
		}
	}

	return result, nil
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

	// Generate parameters in Rust format (name: type)
	params, err := g.formatRustParams(method.ParamTypes, true)
	if err != nil {
		return "", err
	}

	sb.WriteString("#[allow(dead_code, non_snake_case)]\n")
	sb.WriteString(fmt.Sprintf("pub fn %s(%s)", g.SanitizeName(method.Name), params))
	if retType != "" && retType != "()" {
		sb.WriteString(" -> ")
		sb.WriteString(retType)
	}
	sb.WriteString(" {\n")

	// Generate function body using lazy static pattern
	// For extern "C" fn, we don't need names, just types
	funcTypeParams, err := g.formatRustParams(method.ParamTypes, false)
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

	sb.WriteString("        let __func = FUNC.get_or_init(|| {\n")
	sb.WriteString(fmt.Sprintf("            let name = \"%s.%s\";\n", pluginName, method.Name))
	sb.WriteString("            let ptr = get_method_ptr(name.as_ptr(), name.len());\n")
	sb.WriteString("            std::mem::transmute(ptr)\n")
	sb.WriteString("        });\n")

	// Call function - just use parameter names
	paramNames, err := FormatParameters(method.ParamTypes, ParamFormatNames, g.typeMapper, g.SanitizeName)
	if err != nil {
		return "", err
	}

	if method.RetType.Type == "void" {
		sb.WriteString(fmt.Sprintf("        __func(%s);\n", paramNames))
	} else {
		sb.WriteString(fmt.Sprintf("        __func(%s)\n", paramNames))
	}

	sb.WriteString("    }\n")
	sb.WriteString("}\n")

	return sb.String(), nil
}

// generateEnumsFile generates a file containing all enums
func (g *RustGenerator) generateEnumsFile(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("// Generated from %s.pplugin\n\n", m.Name))
	sb.WriteString("#[allow(unused_imports)]\n")
	sb.WriteString("use plugify::{vector_enum_traits};\n\n")

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
	sb.WriteString("#[allow(unused_imports)]\n")
	sb.WriteString("use super::enums::*;\n")
	sb.WriteString("#[allow(unused_imports)]\n")
	sb.WriteString("use plugify::{PlgString, PlgVector, PlgVariant, Vector2, Vector3, Vector4, Matrix4x4};\n\n")

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
	sb.WriteString("#[allow(unused_imports)]\n")
	sb.WriteString("use std::sync::OnceLock;\n")
	sb.WriteString("#[allow(unused_imports)]\n")
	sb.WriteString("use super::enums::*;\n")
	sb.WriteString("#[allow(unused_imports)]\n")
	sb.WriteString("use super::delegates::*;\n")
	sb.WriteString("#[allow(unused_imports)]\n")
	sb.WriteString("use plugify::{get_method_ptr, PlgString, PlgVector, PlgVariant, Vector2, Vector3, Vector4, Matrix4x4};\n\n")

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

	// Generate classes for this group (if enabled)
	if opts.GenerateClasses {
		for _, class := range m.Classes {
			classGroup := g.GetGroupName(class.Group)
			if classGroup == groupName {
				classCode, err := g.generateClass(m, &class)
				if err != nil {
					return "", fmt.Errorf("failed to generate class %s: %w", class.Name, err)
				}
				sb.WriteString(classCode)
				sb.WriteString("\n")
			}
		}
	}

	return sb.String(), nil
}

// generateModFile generates the mod.rs file that re-exports all modules
func (g *RustGenerator) generateModFile(m *manifest.Manifest, groups map[string]bool) (string, error) {
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
	sb.WriteString("#[allow(unused_imports)]\n")
	sb.WriteString("pub use enums::*;\n")
	sb.WriteString("#[allow(unused_imports)]\n")
	sb.WriteString("pub use delegates::*;\n")
	for groupName := range groups {
		sb.WriteString("#[allow(unused_imports)]\n")
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

	// Handle parameter context (value parameters)
	// Object-like types pass by const& (in Rust: &) even when not ref=true
	if context == TypeContextValue && baseType != "void" {
		if m.isObjectLikeType(baseType) || isArray {
			mapped = fmt.Sprintf("&%s", mapped)
		}
	}

	// Handle reference context (ref=true parameters)
	if context == TypeContextRef && baseType != "void" {
		mapped = fmt.Sprintf("&mut %s", mapped)
	}

	return mapped, nil
}

// isObjectLikeType returns true for types that should be passed by reference in parameters
func (m *RustTypeMapper) isObjectLikeType(baseType string) bool {
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

func (m *RustTypeMapper) MapParamType(param *manifest.ParamType, context TypeContext) (string, error) {
	// Check for enum type first
	if param.Enum != nil {
		typeName := param.Enum.Name
		if param.IsArray() {
			typeName = fmt.Sprintf("PlgVector<%s>", typeName)
		}
		// Handle reference
		if param.Ref {
			return fmt.Sprintf("&mut %s", typeName), nil
		} else if param.IsArray() {
			return fmt.Sprintf("&%s", typeName), nil
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

func (g *RustGenerator) generateClass(m *manifest.Manifest, class *manifest.Class) (string, error) {
	var sb strings.Builder

	// Check if this is a handleless class (static methods only)
	hasHandle := class.HandleType != "" && class.HandleType != "void"

	hasCtor := len(class.Constructors) > 0
	hasDtor := class.Destructor != nil

	// Validate: handleless classes should only have static methods
	if !hasHandle {
		for _, binding := range class.Bindings {
			if binding.BindSelf {
				return "", fmt.Errorf("class %s: handleless classes (handleType is void/empty) cannot have instance methods (bindSelf=true for %s)", class.Name, binding.Name)
			}
		}
		if hasCtor || hasDtor {
			return "", fmt.Errorf("class %s: handleless classes cannot have constructors or destructors", class.Name)
		}
	}

	// Map handle type
	invalidValue, handleType, err := g.typeMapper.MapHandleType(class)
	if err != nil {
		return "", err
	}

	// Generate error type for this class if it has a handle
	if hasHandle {
		sb.WriteString("#[derive(Debug)]\n")
		sb.WriteString(fmt.Sprintf("pub enum %sError {\n", class.Name))
		sb.WriteString("    EmptyHandle,\n")
		sb.WriteString("}\n\n")

		sb.WriteString(fmt.Sprintf("impl std::fmt::Display for %sError {\n", class.Name))
		sb.WriteString("    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {\n")
		sb.WriteString("        match self {\n")
		sb.WriteString(fmt.Sprintf("            %sError::EmptyHandle => write!(f, \"empty handle\"),\n", class.Name))
		sb.WriteString("        }\n")
		sb.WriteString("    }\n")
		sb.WriteString("}\n\n")

		sb.WriteString(fmt.Sprintf("impl std::error::Error for %sError {}\n\n", class.Name))
	}

	// Class documentation
	if class.Description != "" {
		sb.WriteString(fmt.Sprintf("/// %s\n", class.Description))
	}

	// Struct definition
	if hasHandle {
		if hasDtor {
			sb.WriteString("#[derive(Debug)]\n")
		} else {
			sb.WriteString("#[derive(Debug, Clone, Copy)]\n")
		}
	} else {
		sb.WriteString("#[derive(Debug, Clone, Copy)]\n")
	}

	sb.WriteString(fmt.Sprintf("pub struct %s {\n", class.Name))
	if hasHandle {
		sb.WriteString(fmt.Sprintf("    handle: %s,\n", handleType))
		if hasDtor {
			sb.WriteString("    ownership: Ownership,\n")
		}
	}
	sb.WriteString("}\n\n")

	// Only generate handle-related code if class has a handle
	if hasHandle {
		sb.WriteString(fmt.Sprintf("impl %s {\n", class.Name))

		// Generate constructors
		for _, ctorName := range class.Constructors {
			ctorCode, err := g.generateConstructor(m, class, ctorName)
			if err != nil {
				return "", err
			}
			sb.WriteString(ctorCode)
		}

		// Generate from_raw helper constructor
		if hasDtor {
			sb.WriteString(fmt.Sprintf("    /// Construct from raw handle with specified ownership\n"))
			sb.WriteString("    #[allow(dead_code)]\n")
			sb.WriteString(fmt.Sprintf("    pub unsafe fn from_raw(handle: %s, ownership: Ownership) -> Self {\n", handleType))
			sb.WriteString("        Self { handle, ownership }\n")
			sb.WriteString("    }\n\n")
		} else {
			sb.WriteString(fmt.Sprintf("    /// Construct from raw handle (does not assume ownership)\n"))
			sb.WriteString("    #[allow(dead_code)]\n")
			sb.WriteString(fmt.Sprintf("    pub unsafe fn from_raw(handle: %s) -> Self {\n", handleType))
			sb.WriteString("        Self { handle }\n")
			sb.WriteString("    }\n\n")
		}

		// Generate utility methods
		utilityCode, err := g.generateUtilityMethods(m, class, invalidValue, handleType)
		if err != nil {
			return "", err
		}
		sb.WriteString(utilityCode)

		// Generate class bindings
		for _, binding := range class.Bindings {
			methodCode, err := g.generateBinding(m, class, &binding, invalidValue)
			if err != nil {
				return "", err
			}
			sb.WriteString(methodCode)
		}

		sb.WriteString("}\n\n")

		// Generate Drop trait if destructor exists
		if hasDtor {
			dropCode, err := g.generateDropTrait(m, class, invalidValue)
			if err != nil {
				return "", err
			}
			sb.WriteString(dropCode)
		}

		// Generate comparison traits
		comparisonCode := g.generateComparisonTraits(class)
		sb.WriteString(comparisonCode)
	} else {
		// For handleless classes, generate static methods
		sb.WriteString(fmt.Sprintf("impl %s {\n", class.Name))
		for _, binding := range class.Bindings {
			methodCode, err := g.generateBinding(m, class, &binding, "")
			if err != nil {
				return "", err
			}
			sb.WriteString(methodCode)
		}
		sb.WriteString("}\n\n")
	}

	return sb.String(), nil
}

func (g *RustGenerator) generateConstructor(m *manifest.Manifest, class *manifest.Class, methodName string) (string, error) {
	// Find the method in the manifest
	method := FindMethod(m, methodName)
	if method == nil {
		return "", fmt.Errorf("constructor method %s not found", methodName)
	}

	var sb strings.Builder

	// Generate documentation
	if method.Description != "" {
		sb.WriteString(fmt.Sprintf("    /// %s\n", method.Description))
	}

	// Generate parameters documentation
	for _, param := range method.ParamTypes {
		if param.Description != "" {
			sb.WriteString(fmt.Sprintf("    /// @param %s: %s\n", param.Name, param.Description))
		}
	}

	// Generate function signature
	params, err := g.formatRustParams(method.ParamTypes, true)
	if err != nil {
		return "", err
	}

	funcName := fmt.Sprintf("new")
	if method.Name != class.Name {
		// If constructor method name is different from class name, use it as suffix
		funcName = fmt.Sprintf("new_%s", method.Name)
	}

	hasDtor := class.Destructor != nil
	sb.WriteString("    #[allow(dead_code, non_snake_case)]\n")
	sb.WriteString(fmt.Sprintf("    pub fn %s(%s) -> Result<Self, %sError> {\n", funcName, params, class.Name))

	// Generate call to underlying FFI function
	paramNames, err := FormatParameters(method.ParamTypes, ParamFormatNames, g.typeMapper, g.SanitizeName)
	if err != nil {
		return "", err
	}

	invalidValue, _, err := g.typeMapper.MapHandleType(class)
	if err != nil {
		return "", err
	}

	sb.WriteString(fmt.Sprintf("        let h = crate::%s::%s(%s);\n", m.Name, method.FuncName, paramNames))
	sb.WriteString(fmt.Sprintf("        if h == %s {\n", invalidValue))
	sb.WriteString(fmt.Sprintf("            return Err(%sError::EmptyHandle);\n", class.Name))
	sb.WriteString("        }\n")
	sb.WriteString("        Ok(Self {\n")
	sb.WriteString("            handle: h,\n")
	if hasDtor {
		sb.WriteString("            ownership: Ownership::Owned,\n")
	}
	sb.WriteString("        })\n")
	sb.WriteString("    }\n\n")

	return sb.String(), nil
}

func (g *RustGenerator) generateUtilityMethods(m *manifest.Manifest, class *manifest.Class, invalidValue, handleType string) (string, error) {
	var sb strings.Builder

	hasDtor := class.Destructor != nil

	// get method
	sb.WriteString("    /// Returns the underlying handle\n")
	sb.WriteString("    #[allow(dead_code)]\n")
	sb.WriteString(fmt.Sprintf("    pub fn get(&self) -> %s {\n", handleType))
	sb.WriteString("        self.handle\n")
	sb.WriteString("    }\n\n")

	// release method
	sb.WriteString("    /// Release ownership and return the handle. Wrapper becomes empty & borrowed.\n")
	sb.WriteString("    #[allow(dead_code)]\n")
	sb.WriteString(fmt.Sprintf("    pub fn release(&mut self) -> %s {\n", handleType))
	sb.WriteString("        let h = self.handle;\n")
	sb.WriteString(fmt.Sprintf("        self.handle = %s;\n", invalidValue))
	if hasDtor {
		sb.WriteString("        self.ownership = Ownership::Borrowed;\n")
	}
	sb.WriteString("        h\n")
	sb.WriteString("    }\n\n")

	// reset method
	sb.WriteString("    /// Destroys and resets the handle\n")
	sb.WriteString("    #[allow(dead_code)]\n")
	sb.WriteString("    pub fn reset(&mut self) {\n")
	if hasDtor {
		sb.WriteString(fmt.Sprintf("        if self.handle != %s && self.ownership == Ownership::Owned {\n", invalidValue))
		sb.WriteString(fmt.Sprintf("            crate::%s::%s(self.handle);\n", m.Name, *class.Destructor))
		sb.WriteString("        }\n")
	}
	sb.WriteString(fmt.Sprintf("        self.handle = %s;\n", invalidValue))
	if hasDtor {
		sb.WriteString("        self.ownership = Ownership::Borrowed;\n")
	}
	sb.WriteString("    }\n\n")

	// swap method
	sb.WriteString(fmt.Sprintf("    /// Swaps two %s instances\n", class.Name))
	sb.WriteString("    #[allow(dead_code)]\n")
	sb.WriteString(fmt.Sprintf("    pub fn swap(&mut self, other: &mut %s) {\n", class.Name))
	sb.WriteString("        std::mem::swap(&mut self.handle, &mut other.handle);\n")
	if hasDtor {
		sb.WriteString("        std::mem::swap(&mut self.ownership, &mut other.ownership);\n")
	}
	sb.WriteString("    }\n\n")

	// is_valid method
	sb.WriteString("    /// Returns true if handle is valid (not empty)\n")
	sb.WriteString("    #[allow(dead_code)]\n")
	sb.WriteString("    pub fn is_valid(&self) -> bool {\n")
	sb.WriteString(fmt.Sprintf("        self.handle != %s\n", invalidValue))
	sb.WriteString("    }\n\n")

	return sb.String(), nil
}

func (g *RustGenerator) generateBinding(m *manifest.Manifest, class *manifest.Class, binding *manifest.Binding, invalidValue string) (string, error) {
	// Find the underlying method
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

	// Generate documentation
	if method.Description != "" {
		sb.WriteString(fmt.Sprintf("    /// %s\n", method.Description))
	}

	// Document parameters
	for _, param := range methodParams {
		if param.Description != "" {
			sb.WriteString(fmt.Sprintf("    /// @param %s: %s\n", param.Name, param.Description))
		}
	}

	// Document return type
	if method.RetType.Type != "void" && method.RetType.Description != "" {
		sb.WriteString(fmt.Sprintf("    /// @return %s\n", method.RetType.Description))
	}

	// Generate method signature
	formattedParams, err := g.formatClassParams(methodParams, binding.ParamAliases)
	if err != nil {
		return "", err
	}

	// Build return type
	returnSignature := ""
	hasHandle := class.HandleType != "" && class.HandleType != "void"

	// Determine if method is static
	if !binding.BindSelf {
		if method.RetType.Type != "void" {
			if binding.RetAlias != nil && binding.RetAlias.Name != "" {
				returnSignature = fmt.Sprintf(" -> %s", binding.RetAlias.Name)
			} else {
				retType, err := g.typeMapper.MapReturnType(&method.RetType)
				if err != nil {
					return "", err
				}
				returnSignature = fmt.Sprintf(" -> %s", retType)
			}
		}
		sb.WriteString("    #[allow(dead_code, non_snake_case)]\n")
		sb.WriteString(fmt.Sprintf("    pub fn %s(%s)%s {\n", binding.Name, formattedParams, returnSignature))
	} else {
		selfRef := "&self"
		// Determine if we need &mut self
		for i, param := range methodParams {
			if param.Ref || (i < len(binding.ParamAliases) && binding.ParamAliases[i] != nil && binding.ParamAliases[i].Owner) {
				selfRef = "&mut self"
				break
			}
		}
		// Check if method modifies the handle (for ref parameters)
		if len(params) > 0 && params[0].Ref {
			selfRef = "&mut self"
		}

		if method.RetType.Type != "void" {
			if binding.RetAlias != nil && binding.RetAlias.Name != "" {
				returnSignature = fmt.Sprintf(" -> Result<%s, %sError>", binding.RetAlias.Name, class.Name)
			} else {
				retType, err := g.typeMapper.MapReturnType(&method.RetType)
				if err != nil {
					return "", err
				}
				returnSignature = fmt.Sprintf(" -> Result<%s, %sError>", retType, class.Name)
			}
		} else {
			returnSignature = fmt.Sprintf(" -> Result<(), %sError>", class.Name)
		}
		sb.WriteString("    #[allow(dead_code, non_snake_case)]\n")
		sb.WriteString(fmt.Sprintf("    pub fn %s(%s", binding.Name, selfRef))
		if formattedParams != "" {
			sb.WriteString(fmt.Sprintf(", %s", formattedParams))
		}
		sb.WriteString(fmt.Sprintf(")%s {\n", returnSignature))
	}

	// Generate null check
	nullPolicy := class.NullPolicy
	if nullPolicy == "" {
		nullPolicy = "throw"
	}

	if binding.BindSelf && hasHandle && nullPolicy == "throw" {
		sb.WriteString(fmt.Sprintf("        if self.handle == %s {\n", invalidValue))
		sb.WriteString(fmt.Sprintf("            return Err(%sError::EmptyHandle);\n", class.Name))
		sb.WriteString("        }\n")
	}

	// Build call arguments
	callArgs, err := g.formatClassCallArgs(methodParams, binding, binding.BindSelf)
	if err != nil {
		return "", err
	}

	hasCtor := len(class.Constructors) > 0
	hasDtor := class.Destructor != nil

	functionName := fmt.Sprintf("crate::%s::%s", m.Name, method.FuncName)

	// Generate call
	if method.RetType.Type == "void" {
		sb.WriteString(fmt.Sprintf("        %s(%s);\n", functionName, callArgs))
		if binding.BindSelf {
			sb.WriteString("        Ok(())\n")
		}
	} else {
		if binding.RetAlias != nil && binding.RetAlias.Name != "" {
			ownership := ""
			if hasDtor || hasCtor {
				if binding.RetAlias.Owner {
					ownership = ", Ownership::Owned"
				} else {
					ownership = ", Ownership::Borrowed"
				}
			}
			if binding.BindSelf {
				sb.WriteString(fmt.Sprintf("        Ok(unsafe { %s::from_raw(%s(%s)%s) })\n", binding.RetAlias.Name, functionName, callArgs, ownership))
			} else {
				sb.WriteString(fmt.Sprintf("        unsafe { %s::from_raw(%s(%s)%s) }\n", binding.RetAlias.Name, functionName, callArgs, ownership))
			}
		} else {
			if binding.BindSelf {
				sb.WriteString(fmt.Sprintf("        Ok(%s(%s))\n", functionName, callArgs))
			} else {
				sb.WriteString(fmt.Sprintf("        %s(%s)\n", functionName, callArgs))
			}
		}
	}

	sb.WriteString("    }\n\n")

	return sb.String(), nil
}

func (g *RustGenerator) formatClassParams(params []manifest.ParamType, aliases []*manifest.ParamAlias) (string, error) {
	if len(params) == 0 {
		return "", nil
	}

	var parts []string
	for i, param := range params {
		name := g.SanitizeName(param.Name)

		var typeName string
		var err error

		// Check if this parameter has an alias
		if i < len(aliases) && aliases[i] != nil {
			if aliases[i].Owner {
				typeName = aliases[i].Name
			} else {
				typeName = fmt.Sprintf("&%s", aliases[i].Name)
			}
		} else {
			typeName, err = g.typeMapper.MapParamType(&param, TypeContextValue)
			if err != nil {
				return "", err
			}
		}
		parts = append(parts, fmt.Sprintf("%s: %s", name, typeName))
	}

	return strings.Join(parts, ", "), nil
}

func (g *RustGenerator) formatClassCallArgs(params []manifest.ParamType, binding *manifest.Binding, bindSelf bool) (string, error) {
	var parts []string

	// Add self if bindSelf
	if bindSelf {
		parts = append(parts, "self.handle")
	}

	// Add other parameters
	for i, param := range params {
		name := g.SanitizeName(param.Name)

		// Check if parameter has alias
		if i < len(binding.ParamAliases) && binding.ParamAliases[i] != nil {
			if binding.ParamAliases[i].Owner {
				parts = append(parts, fmt.Sprintf("%s.release()", name))
			} else {
				parts = append(parts, fmt.Sprintf("%s.get()", name))
			}
		} else {
			parts = append(parts, name)
		}
	}

	return strings.Join(parts, ", "), nil
}

func (g *RustGenerator) generateDropTrait(m *manifest.Manifest, class *manifest.Class, invalidValue string) (string, error) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("impl Drop for %s {\n", class.Name))
	sb.WriteString("    fn drop(&mut self) {\n")
	sb.WriteString(fmt.Sprintf("        if self.handle != %s && self.ownership == Ownership::Owned {\n", invalidValue))
	sb.WriteString(fmt.Sprintf("            crate::%s::%s(self.handle);\n", m.Name, *class.Destructor))
	sb.WriteString("        }\n")
	sb.WriteString("    }\n")
	sb.WriteString("}\n\n")

	return sb.String(), nil
}

func (g *RustGenerator) generateComparisonTraits(class *manifest.Class) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("impl std::cmp::PartialEq for %s {\n", class.Name))
	sb.WriteString(fmt.Sprintf("    fn eq(&self, other: &Self) -> bool {\n"))
	sb.WriteString("        self.handle == other.handle\n")
	sb.WriteString("    }\n")
	sb.WriteString("}\n")
	sb.WriteString(fmt.Sprintf("impl std::cmp::Eq for %s {}\n", class.Name))
	sb.WriteString(fmt.Sprintf("impl std::cmp::PartialOrd for %s {\n", class.Name))
	sb.WriteString("    fn partial_cmp(&self, other: &Self) -> Option<std::cmp::Ordering> {\n")
	sb.WriteString("        (self.handle).partial_cmp(&(other.handle))\n")
	sb.WriteString("    }\n")
	sb.WriteString("}\n")
	sb.WriteString(fmt.Sprintf("impl std::cmp::Ord for %s {\n", class.Name))
	sb.WriteString("    fn cmp(&self, other: &Self) -> std::cmp::Ordering {\n")
	sb.WriteString("        (self.handle).cmp(&(other.handle))\n")
	sb.WriteString("    }\n")
	sb.WriteString("}\n\n")

	return sb.String()
}
