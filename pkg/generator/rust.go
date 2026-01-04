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
	m.Sanitize(g.Sanitizer)
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

// RustDocOptions configures Rust documentation generation
type RustDocOptions struct {
	Description string
	Params      []manifest.ParamType
	Returns     string
	Indent      string
}

// generateRustDocumentation generates Rust-style documentation comments (///)
func (g *RustGenerator) generateRustDocumentation(opts RustDocOptions) string {
	var sb strings.Builder

	// Main description
	if opts.Description != "" {
		sb.WriteString(fmt.Sprintf("%s/// %s\n", opts.Indent, opts.Description))
	}

	// Parameters
	if len(opts.Params) > 0 {
		if opts.Description != "" {
			sb.WriteString(fmt.Sprintf("%s///\n", opts.Indent))
		}
		sb.WriteString(fmt.Sprintf("%s/// # Arguments\n", opts.Indent))
		for _, param := range opts.Params {
			paramType := param.Type
			if param.Ref {
				paramType += "&"
			}
			sb.WriteString(fmt.Sprintf("%s/// * `%s` - (%s)", opts.Indent, param.Name, paramType))
			if param.Description != "" {
				sb.WriteString(fmt.Sprintf(": %s", param.Description))
			}
			sb.WriteString("\n")
		}
	}

	// Returns
	if opts.Returns != "" {
		sb.WriteString(fmt.Sprintf("%s///\n", opts.Indent))
		sb.WriteString(fmt.Sprintf("%s/// # Returns\n", opts.Indent))
		sb.WriteString(fmt.Sprintf("%s/// %s\n", opts.Indent, opts.Returns))
	}

	return sb.String()
}

// formatParameters is a generic helper that formats parameters using a custom formatter function
func (g *RustGenerator) formatParameters(params []manifest.ParamType, formatter func(int, *manifest.ParamType) (string, error)) (string, error) {
	if len(params) == 0 {
		return "", nil
	}

	parts := []string{}
	for i, param := range params {
		part, err := formatter(i, &param)
		if err != nil {
			return "", err
		}
		if part != "" {
			parts = append(parts, part)
		}
	}
	return strings.Join(parts, ", "), nil
}

func (g *RustGenerator) generateEnums(m *manifest.Manifest) (string, error) {
	return g.CollectEnums(m, g.generateEnum)
}

func (g *RustGenerator) generateEnum(enum *manifest.EnumType, underlyingType string) (string, error) {
	var sb strings.Builder

	if enum.Description != "" {
		sb.WriteString(g.generateRustDocumentation(RustDocOptions{
			Description: enum.Description,
			Indent:      "",
		}))
	}

	sb.WriteString("#[repr(")
	sb.WriteString(underlyingType)
	sb.WriteString(")]\n")
	sb.WriteString("#[derive(Debug, Clone, Copy, PartialEq, Eq)]\n")
	sb.WriteString(fmt.Sprintf("pub enum %s {\n", enum.Name))

	for _, val := range enum.Values {
		if val.Description != "" {
			sb.WriteString(g.generateRustDocumentation(RustDocOptions{
				Description: val.Description,
				Indent:      "    ",
			}))
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
		sb.WriteString(g.generateRustDocumentation(RustDocOptions{
			Description: proto.Description,
			Indent:      "",
		}))
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
	return g.formatParameters(params, func(_ int, param *manifest.ParamType) (string, error) {
		typeName, err := g.typeMapper.MapParamType(param)
		if err != nil {
			return "", err
		}

		if includeNames {
			return param.Name + ": " + typeName, nil
		}
		return typeName, nil
	})
}

func (g *RustGenerator) generateMethod(pluginName string, method *manifest.Method) (string, error) {
	var sb strings.Builder

	// Generate documentation comment
	returns := ""
	if method.RetType.Type != "void" {
		returns = method.RetType.Type
		if method.RetType.Description != "" {
			returns += ": " + method.RetType.Description
		}
	}

	sb.WriteString(g.generateRustDocumentation(RustDocOptions{
		Description: method.Description,
		Params:      method.ParamTypes,
		Returns:     returns,
		Indent:      "",
	}))

	// Generate function signature
	retType, err := g.typeMapper.MapReturnType(&method.RetType)
	if err != nil {
		return "", err
	}

	// Generate parameters in Rust format (name: type)
	paramTypes, err := g.formatRustParams(method.ParamTypes, true)
	if err != nil {
		return "", err
	}

	// Call function - just use parameter names
	paramNames, err := FormatParameters(method.ParamTypes, ParamFormatNames, g.typeMapper)
	if err != nil {
		return "", err
	}

	// Generate function body using lazy static pattern
	// For extern "C" fn, we don't need names, just types
	funcTypeParams, err := g.formatRustParams(method.ParamTypes, false)
	if err != nil {
		return "", err
	}

	// Generate wrapper function
	sb.WriteString(fmt.Sprintf("#[allow(dead_code, non_snake_case)]\n"))
	sb.WriteString(fmt.Sprintf("pub fn %s(%s)", method.Name, paramTypes))
	if retType != "" && retType != "()" {
		sb.WriteString(" -> ")
		sb.WriteString(retType)
	}
	sb.WriteString(" {\n")
	sb.WriteString("    unsafe { ")
	sb.WriteString(fmt.Sprintf("__%s_%s", pluginName, method.Name))
	sb.WriteString(".expect(\"")
	sb.WriteString(method.Name)
	sb.WriteString(" function was not found\")(")
	sb.WriteString(paramNames)
	sb.WriteString(") }\n")
	sb.WriteString("}\n")

	// pub type FuncType = unsafe extern "C" fn(params...) -> ret;
	sb.WriteString(fmt.Sprintf("pub type _%s = unsafe extern \"C\" fn(%s)",
		method.Name,
		funcTypeParams,
	))

	if retType != "" && retType != "()" {
		sb.WriteString(" -> ")
		sb.WriteString(retType)
	}

	sb.WriteString(";\n")

	// Generate static extern function pointer
	sb.WriteString(fmt.Sprintf("#[allow(dead_code, non_upper_case_globals)]\n#[unsafe(no_mangle)]\n"))
	sb.WriteString(fmt.Sprintf("pub static mut __%s_%s: Option<_%s> = None;\n", pluginName, method.Name, method.Name))

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
	sb.WriteString("use plugify::{Str, Arr, Var, Vec2, Vec3, Vec4, Mat4x4};\n\n")

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
	sb.WriteString("use std::sync::RwLock;\n")
	sb.WriteString("#[allow(unused_imports)]\n")
	sb.WriteString("use super::enums::*;\n")
	sb.WriteString("#[allow(unused_imports)]\n")
	sb.WriteString("use super::delegates::*;\n")
	sb.WriteString("#[allow(unused_imports)]\n")
	sb.WriteString("use plugify::{get_method_ptr, Str, Arr, Var, Vec2, Vec3, Vec4, Mat4x4};\n\n")

	// Generate methods for this group
	for _, method := range m.Methods {
		methodGroup := method.Group
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
			classGroup := class.Group
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
		"string": "Str",
		"any":    "Var",
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

	// Handle void -> ()
	if mapped == "" {
		mapped = "()"
	}

	// Handle arrays
	if isArray {
		mapped = fmt.Sprintf("Arr<%s>", mapped)
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

func (m *RustTypeMapper) MapParamType(param *manifest.ParamType) (string, error) {
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
		typeName = *param.Alias

	case param.Prototype != nil:
		return param.Prototype.Name, nil

	default:
		typeName = param.BaseType()
	}

	return m.MapType(typeName, ctx, param.IsArray())
}

func (m *RustTypeMapper) MapReturnType(retType *manifest.RetType) (string, error) {
	var typeName string
	switch {
	case retType.Enum != nil:
		typeName = retType.Enum.Name

	case retType.Alias != nil:
		typeName = *retType.Alias

	case retType.Prototype != nil:
		return retType.Prototype.Name, nil

	default:
		typeName = retType.BaseType()
	}

	// Regular type mapping - returns always by value
	return m.MapType(typeName, TypeContextReturn, retType.IsArray())
}

func (m *RustTypeMapper) MapHandleType(class *manifest.Class) (string, string, error) {
	invalidValue := class.InvalidValue
	handleType, err := m.MapType(class.HandleType, TypeContextReturn, false)
	if err != nil {
		return "", "", err
	}

	nullptr := invalidValue == "0" || invalidValue == "" || invalidValue == "NULL" || invalidValue == "nullptr"
	if class.HandleType == "ptr64" && nullptr {
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
		sb.WriteString(g.generateRustDocumentation(RustDocOptions{
			Description: class.Description,
			Indent:      "",
		}))
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

// generateSafeConstructorBody generates the body of a constructor that validates the handle
func (g *RustGenerator) generateSafeConstructorBody(m *manifest.Manifest, class *manifest.Class, method *manifest.Method, paramNames string) (string, error) {
	var sb strings.Builder

	invalidValue, _, err := g.typeMapper.MapHandleType(class)
	if err != nil {
		return "", err
	}

	hasDtor := class.Destructor != nil

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
	sb.WriteString(g.generateRustDocumentation(RustDocOptions{
		Description: method.Description,
		Params:      method.ParamTypes,
		Indent:      "    ",
	}))

	// Generate function signature
	params, err := g.formatRustParams(method.ParamTypes, true)
	if err != nil {
		return "", err
	}

	funcName := "new"
	if method.Name != class.Name {
		// If constructor method name is different from class name, use it as suffix
		funcName = fmt.Sprintf("new_%s", method.Name)
	}

	sb.WriteString("    #[allow(dead_code, non_snake_case)]\n")
	sb.WriteString(fmt.Sprintf("    pub fn %s(%s) -> Result<Self, %sError> {\n", funcName, params, class.Name))

	// Generate call to underlying FFI function
	paramNames, err := FormatParameters(method.ParamTypes, ParamFormatNames, g.typeMapper)
	if err != nil {
		return "", err
	}

	// Generate safe constructor body
	body, err := g.generateSafeConstructorBody(m, class, method, paramNames)
	if err != nil {
		return "", err
	}
	sb.WriteString(body)

	sb.WriteString("    }\n\n")

	return sb.String(), nil
}

func (g *RustGenerator) generateUtilityMethods(m *manifest.Manifest, class *manifest.Class, invalidValue, handleType string) (string, error) {
	var sb strings.Builder

	hasDtor := class.Destructor != nil

	// Define utility methods metadata
	type utilityMethod struct {
		doc       string
		signature string
		body      func() string
	}

	methods := []utilityMethod{
		{
			doc:       "Returns the underlying handle",
			signature: fmt.Sprintf("pub fn get(&self) -> %s", handleType),
			body: func() string {
				return "        self.handle\n"
			},
		},
		{
			doc:       "Release ownership and return the handle. Wrapper becomes empty & borrowed.",
			signature: fmt.Sprintf("pub fn release(&mut self) -> %s", handleType),
			body: func() string {
				var b strings.Builder
				b.WriteString("        let h = self.handle;\n")
				b.WriteString(fmt.Sprintf("        self.handle = %s;\n", invalidValue))
				if hasDtor {
					b.WriteString("        self.ownership = Ownership::Borrowed;\n")
				}
				b.WriteString("        h\n")
				return b.String()
			},
		},
		{
			doc:       "Destroys and resets the handle",
			signature: "pub fn reset(&mut self)",
			body: func() string {
				var b strings.Builder
				if hasDtor {
					b.WriteString(fmt.Sprintf("        if self.handle != %s && self.ownership == Ownership::Owned {\n", invalidValue))
					b.WriteString(fmt.Sprintf("            crate::%s::%s(self.handle);\n", m.Name, *class.Destructor))
					b.WriteString("        }\n")
				}
				b.WriteString(fmt.Sprintf("        self.handle = %s;\n", invalidValue))
				if hasDtor {
					b.WriteString("        self.ownership = Ownership::Borrowed;\n")
				}
				return b.String()
			},
		},
		{
			doc:       fmt.Sprintf("Swaps two %s instances", class.Name),
			signature: fmt.Sprintf("pub fn swap(&mut self, other: &mut %s)", class.Name),
			body: func() string {
				var b strings.Builder
				b.WriteString("        std::mem::swap(&mut self.handle, &mut other.handle);\n")
				if hasDtor {
					b.WriteString("        std::mem::swap(&mut self.ownership, &mut other.ownership);\n")
				}
				return b.String()
			},
		},
		{
			doc:       "Returns true if handle is valid (not empty)",
			signature: "pub fn is_valid(&self) -> bool",
			body: func() string {
				return fmt.Sprintf("        self.handle != %s\n", invalidValue)
			},
		},
	}

	// Generate methods
	for _, method := range methods {
		sb.WriteString(fmt.Sprintf("    /// %s\n", method.doc))
		sb.WriteString("    #[allow(dead_code)]\n")
		sb.WriteString(fmt.Sprintf("    %s {\n", method.signature))
		sb.WriteString(method.body())
		sb.WriteString("    }\n\n")
	}

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
	returns := ""
	if method.RetType.Type != "void" && method.RetType.Description != "" {
		returns = method.RetType.Description
	}

	sb.WriteString(g.generateRustDocumentation(RustDocOptions{
		Description: method.Description,
		Params:      methodParams,
		Returns:     returns,
		Indent:      "    ",
	}))

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
	return g.formatParameters(params, func(i int, param *manifest.ParamType) (string, error) {
		name := param.Name
		var typeName string

		// Check if this parameter has an alias
		if i < len(aliases) && aliases[i] != nil {
			if aliases[i].Owner {
				typeName = aliases[i].Name
			} else {
				typeName = fmt.Sprintf("&%s", aliases[i].Name)
			}
		} else {
			var err error
			typeName, err = g.typeMapper.MapParamType(param)
			if err != nil {
				return "", err
			}
		}
		return fmt.Sprintf("%s: %s", name, typeName), nil
	})
}

func (g *RustGenerator) formatClassCallArgs(params []manifest.ParamType, binding *manifest.Binding, bindSelf bool) (string, error) {
	var parts []string

	// Add self if bindSelf
	if bindSelf {
		parts = append(parts, "self.handle")
	}

	// Add other parameters using generic helper
	paramParts, err := g.formatParameters(params, func(i int, param *manifest.ParamType) (string, error) {
		name := param.Name

		// Check if parameter has alias
		if i < len(binding.ParamAliases) && binding.ParamAliases[i] != nil {
			if binding.ParamAliases[i].Owner {
				return fmt.Sprintf("%s.release()", name), nil
			}
			return fmt.Sprintf("%s.get()", name), nil
		}
		return name, nil
	})
	if err != nil {
		return "", err
	}

	if paramParts != "" {
		if len(parts) > 0 {
			parts = append(parts, paramParts)
			return strings.Join(parts, ", "), nil
		}
		return paramParts, nil
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
	className := class.Name
	return fmt.Sprintf(`impl std::cmp::PartialEq for %s {
    fn eq(&self, other: &Self) -> bool {
        self.handle == other.handle
    }
}
impl std::cmp::Eq for %s {}
impl std::cmp::PartialOrd for %s {
    fn partial_cmp(&self, other: &Self) -> Option<std::cmp::Ordering> {
        (self.handle).partial_cmp(&(other.handle))
    }
}
impl std::cmp::Ord for %s {
    fn cmp(&self, other: &Self) -> std::cmp::Ordering {
        (self.handle).cmp(&(other.handle))
    }
}

`, className, className, className, className)
}
