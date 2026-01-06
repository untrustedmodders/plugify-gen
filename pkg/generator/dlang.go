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
	return &DlangGenerator{
		BaseGenerator: NewBaseGenerator("dlang", NewDlangTypeMapper(), DReservedWords),
	}
}

// Generate generates D language bindings
func (g *DlangGenerator) Generate(m *manifest.Manifest, opts *GeneratorOptions) (*GeneratorResult, error) {
	g.ResetCaches()
	m.Sanitize(g.Sanitizer)
	opts = EnsureOptions(opts)

	// Module declaration
	moduleName := strings.ToLower(m.Name)

	// Collect all unique groups from both methods and classes
	groups := g.GetGroups(m)
	// Generate files for each group
	files := make(map[string]string)

	// First, generate enums file
	enumsCode, err := g.generateEnumsFile(m)
	if err != nil {
		return nil, fmt.Errorf("generating enums file: %w", err)
	}
	files[fmt.Sprintf("source/imported/%s/enums.d", moduleName)] = enumsCode

	// Secondly separate aliases file
	aliasesCode, err := g.generateAliasesFile(m)
	if err != nil {
		return nil, fmt.Errorf("generating aliases file: %w", err)
	}
	files[fmt.Sprintf("source/imported/%s/aliases.d", moduleName)] = aliasesCode

	// Thirdly, generate delegates file
	delegateCode, err := g.generateDelegatesFile(m)
	if err != nil {
		return nil, fmt.Errorf("generating delegates file: %w", err)
	}
	files[fmt.Sprintf("source/imported/%s/delegates.d", moduleName)] = delegateCode

	// Generate package.d file with all public imports
	packageCode := g.generatePackageFile(m, moduleName, groups)
	files[fmt.Sprintf("source/imported/%s/package.d", moduleName)] = packageCode

	// Generate a file for each group
	for groupName := range groups {
		groupCode, err := g.generateModuleFile(m, moduleName, groupName, opts)
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

	sb.WriteString("/// Ownership type for RAII wrappers\n")
	sb.WriteString(fmt.Sprintf("enum %s : bool {\n", OwnershipEnumName))
	sb.WriteString("\tBorrowed = false,\n")
	sb.WriteString("\tOwned = true\n")
	sb.WriteString("}\n\n")

	// Use the base generator's CollectEnums helper
	enumsCode, err := g.generateEnums(m)
	if err != nil {
		return "", fmt.Errorf("generating enums file: %w", err)
	}
	sb.WriteString(enumsCode)

	return sb.String(), nil
}

func (g *DlangGenerator) generateEnums(m *manifest.Manifest) (string, error) {
	return g.CollectEnums(m, g.generateEnum)
}

func (g *DlangGenerator) generateAliasesFile(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	moduleName := strings.ToLower(m.Name)
	sb.WriteString(fmt.Sprintf("module imported.%s.aliases;\n\n", moduleName))

	// Use the base generator's CollectAliases helper
	aliasesCode, err := g.generateAliases(m)
	if err != nil {
		return "", fmt.Errorf("generating alias file: %w", err)
	}
	sb.WriteString(aliasesCode)

	return sb.String(), nil
}

func (g *DlangGenerator) generateAliases(m *manifest.Manifest) (string, error) {
	return g.CollectAliases(m, g.generateAlias)
}

func (g *DlangGenerator) generateDelegatesFile(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	moduleName := strings.ToLower(m.Name)
	sb.WriteString(fmt.Sprintf("module imported.%s.delegates;\n\n", moduleName))
	sb.WriteString(fmt.Sprintf("public import imported.%s.enums;\n", moduleName))
	sb.WriteString(fmt.Sprintf("public import imported.%s.aliases;\n\n", moduleName))

	delegatesCode, err := g.generateDelegates(m)
	if err != nil {
		return "", fmt.Errorf("generating delegates file: %w", err)
	}
	sb.WriteString(delegatesCode)

	return sb.String(), nil
}

func (g *DlangGenerator) generateDelegates(m *manifest.Manifest) (string, error) {
	return g.CollectDelegates(m, g.generateDelegate)
}

func (g *DlangGenerator) generatePackageFile(m *manifest.Manifest, moduleName string, groups map[string]struct{}) string {
	var sb strings.Builder

	// Module declaration
	sb.WriteString(fmt.Sprintf("module imported.%s;\n\n", moduleName))

	// Always import enums first
	sb.WriteString(fmt.Sprintf("public import imported.%s.enums;\n", moduleName))
	sb.WriteString(fmt.Sprintf("public import imported.%s.aliases;\n", moduleName))
	sb.WriteString(fmt.Sprintf("public import imported.%s.delegates;\n", moduleName))

	// Import all group modules
	for groupName := range groups {
		sb.WriteString(fmt.Sprintf("public import imported.%s.%s;\n", moduleName, groupName))
	}

	return sb.String()
}

func (g *DlangGenerator) generateEnum(enum *manifest.Enum, underlyingType string) (string, error) {
	var sb strings.Builder

	if enum.Description != "" {
		sb.WriteString(fmt.Sprintf("/// %s\n", enum.Description))
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
	return sb.String(), nil
}

func (g *DlangGenerator) generateAlias(alias *manifest.Alias, underlyingType string) (string, error) {
	var sb strings.Builder

	if alias.Description != "" {
		sb.WriteString(fmt.Sprintf("/// %s\n", alias.Description))
	}

	sb.WriteString(fmt.Sprintf("alias %s = %s;\n", alias.Name, underlyingType))

	return sb.String(), nil
}

func (g *DlangGenerator) generateModuleFile(m *manifest.Manifest, moduleName, groupName string, opts *GeneratorOptions) (string, error) {
	var sb strings.Builder

	// Module declaration
	sb.WriteString(fmt.Sprintf("module imported.%s.%s;\n\n", moduleName, groupName))

	// Imports
	sb.WriteString("import plugify.internals;\n")
	sb.WriteString("public import plugify;\n")
	sb.WriteString(fmt.Sprintf("public import imported.%s.enums;\n", moduleName))
	sb.WriteString(fmt.Sprintf("public import imported.%s.aliases;\n", moduleName))
	sb.WriteString(fmt.Sprintf("public import imported.%s.delegates;\n", moduleName))

	// Find which other groups this group depends on (for method calls from classes)
	if len(m.Classes) > 0 {
		dependentGroups := g.FindDependentGroups(m, groupName)
		if len(dependentGroups) > 0 {
			for depGroup := range dependentGroups {
				sb.WriteString(fmt.Sprintf("import imported.%s.%s;\n", moduleName, depGroup))
			}
			sb.WriteString("\n")
		}
	}

	sb.WriteString("import std.exception : enforce;\n")
	sb.WriteString("import std.algorithm.mutation : swap;\n\n")

	// Generate methods for this group
	for _, method := range m.Methods {
		methodGroup := method.Group
		if methodGroup == groupName {
			methodCode, err := g.generateMethodWrapper(&method, m.Name)
			if err != nil {
				return "", fmt.Errorf("failed to generate method wrapper %s: %w", method.Name, err)
			}
			sb.WriteString(methodCode)
			aliasCode, err := g.generateMethodAlias(&method, m.Name)
			if err != nil {
				return "", fmt.Errorf("failed to generate method alias %s: %w", method.Name, err)
			}
			sb.WriteString(aliasCode)
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

// generateDocumentation generates DDoc-style comments for D language
func (g *DlangGenerator) generateDocumentation(opts DocOptions) string {
	var sb strings.Builder

	sb.WriteString(opts.Indent + "/++\n")
	if opts.Description != "" {
		sb.WriteString(fmt.Sprintf("%s\t%s\n", opts.Indent, opts.Description))
	}

	// Parameters section
	if len(opts.Params) > 0 {
		if opts.Description != "" {
			sb.WriteString(opts.Indent + "\n")
		}
		sb.WriteString(opts.Indent + "\tParams:\n")
		for _, param := range opts.Params {
			sb.WriteString(fmt.Sprintf("%s\t\t%s = %s\n", opts.Indent, param.Name, param.Description))
		}
	}

	// Return type section
	if opts.RetType.Type != "void" {
		if len(opts.Params) > 0 || opts.Description != "" {
			sb.WriteString(opts.Indent + "\n")
		}
		sb.WriteString(opts.Indent + "\tReturns:\n")
		if opts.RetType.Description != "" {
			sb.WriteString(fmt.Sprintf("%s\t\t%s\n", opts.Indent, opts.RetType.Description))
		} else {
			sb.WriteString(opts.Indent + "\t\tReturn value\n")
		}
	}

	// Throws section
	if opts.AddThrows {
		sb.WriteString(opts.Indent + "\n")
		sb.WriteString(opts.Indent + "\tThrows: Exception if handle is null\n")
	}

	sb.WriteString(opts.Indent + "+/\n")

	// Add deprecation attribute if present
	if opts.Deprecated != "" {
		sb.WriteString(fmt.Sprintf("\tdeprecated(\"%s\")\n", opts.Deprecated))
	}

	return sb.String()
}

// generateUtilityMethods generates utility methods for RAII wrapper classes
func (g *DlangGenerator) generateUtilityMethods(className, handleType, invalidValue string, hasDtor bool) string {
	var sb strings.Builder

	// Get method
	sb.WriteString("\t/// Get the underlying handle\n")
	sb.WriteString(fmt.Sprintf("\t@property %s get() const pure nothrow @nogc {\n", handleType))
	sb.WriteString("\t\treturn _handle;\n")
	sb.WriteString("\t}\n\n")

	// Release method
	sb.WriteString("\t/// Release ownership of the handle\n")
	sb.WriteString(fmt.Sprintf("\t%s release() nothrow @nogc {\n", handleType))
	sb.WriteString("\t\tauto handle = _handle;\n")
	if hasDtor {
		sb.WriteString("\t\tnullify();\n")
	} else {
		sb.WriteString(fmt.Sprintf("\t\t_handle = %s;\n", invalidValue))
	}
	sb.WriteString("\t\treturn handle;\n")
	sb.WriteString("\t}\n\n")

	// Reset method
	sb.WriteString("\t/// Reset the handle\n")
	if hasDtor {
		sb.WriteString("\tvoid reset() nothrow {\n")
		sb.WriteString("\t\tdestroy();\n")
		sb.WriteString("\t\tnullify();\n")
	} else {
		sb.WriteString("\tvoid reset() nothrow @nogc {\n")
		sb.WriteString(fmt.Sprintf("\t\t_handle = %s;\n", invalidValue))
	}
	sb.WriteString("\t}\n\n")

	// Swap method
	sb.WriteString("\t/// Swap two " + className + " instances\n")
	sb.WriteString(fmt.Sprintf("\tvoid swap(ref %s other) nothrow @nogc {\n", className))
	sb.WriteString("\t\timport std.algorithm.mutation : swap;\n")
	sb.WriteString("\t\tswap(_handle, other._handle);\n")
	if hasDtor {
		sb.WriteString("\t\tswap(_ownership, other._ownership);\n")
	}
	sb.WriteString("\t}\n\n")

	// Boolean conversion operator
	sb.WriteString("\t/// Boolean conversion operator\n")
	sb.WriteString("\tbool opCast(T : bool)() const pure nothrow @nogc {\n")
	sb.WriteString(fmt.Sprintf("\t\treturn _handle !is %s;\n", invalidValue))
	sb.WriteString("\t}\n\n")

	// Comparison operators
	sb.WriteString("\t/// Comparison operators\n")
	sb.WriteString(fmt.Sprintf("\tint opCmp(ref const %s other) const pure nothrow @nogc {\n", className))
	sb.WriteString("\t\tif (_handle < other._handle) return -1;\n")
	sb.WriteString("\t\tif (_handle > other._handle) return 1;\n")
	sb.WriteString("\t\treturn 0;\n")
	sb.WriteString("\t}\n\n")

	sb.WriteString(fmt.Sprintf("\tbool opEquals(ref const %s other) const pure nothrow @nogc {\n", className))
	sb.WriteString("\t\treturn _handle == other._handle;\n")
	sb.WriteString("\t}\n\n")

	return sb.String()
}

// generateMoveSemantics generates move constructor and move assignment operator
func (g *DlangGenerator) generateMoveSemantics(className string) string {
	var sb strings.Builder

	// Move constructor
	sb.WriteString("\t/// Move constructor (called when moving)\n")
	sb.WriteString(fmt.Sprintf("\tthis(ref return scope %s other) {\n", className))
	sb.WriteString("\t\t_handle = other._handle;\n")
	sb.WriteString("\t\t_ownership = other._ownership;\n")
	sb.WriteString("\t\tother.nullify();\n")
	sb.WriteString("\t}\n\n")

	// Move assignment
	sb.WriteString("\t/// Move assignment\n")
	sb.WriteString(fmt.Sprintf("\tref %s opAssign(%s other) return {\n", className, className))
	sb.WriteString("\t\tswap(this, other);\n")
	sb.WriteString("\t\treturn this;\n")
	sb.WriteString("\t}\n\n")

	return sb.String()
}

// generatePrivateHelpers generates private helper methods (destroy and nullify)
func (g *DlangGenerator) generatePrivateHelpers(className, invalidValue, destructor string) string {
	var sb strings.Builder

	sb.WriteString("\tprivate void destroy() const nothrow {\n")
	sb.WriteString(fmt.Sprintf("\t\tif (_handle !is %s && _ownership == %s.Owned) {\n", invalidValue, OwnershipEnumName))
	sb.WriteString(fmt.Sprintf("\t\t\t%s(_handle);\n", destructor))
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n\n")

	sb.WriteString("\tprivate void nullify() nothrow @nogc {\n")
	sb.WriteString(fmt.Sprintf("\t\t_handle = %s;\n", invalidValue))
	sb.WriteString(fmt.Sprintf("\t\t_ownership = %s.Borrowed;\n", OwnershipEnumName))
	sb.WriteString("\t}\n")

	return sb.String()
}

func (g *DlangGenerator) generateDelegate(proto *manifest.Prototype) (string, error) {
	var sb strings.Builder

	sb.WriteString(g.generateDocumentation(DocOptions{
		Description: proto.Description,
		Params:      proto.ParamTypes,
		Indent:      "",
	}))

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
	var params []string
	for _, param := range proto.ParamTypes {
		paramName := param.Name
		paramType, err := g.typeMapper.MapParamType(&param)
		if err != nil {
			return "", err
		}

		// Convert to C ABI type
		cType, err := g.toCType(paramType, &manifest.RetType{
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

func (g *DlangGenerator) generateMethodWrapper(method *manifest.Method, pluginName string) (string, error) {
	var sb strings.Builder

	// Documentation
	sb.WriteString(g.generateDocumentation(DocOptions{
		Description: method.Description,
		Deprecated:  method.Deprecated,
		Params:      method.ParamTypes,
		RetType:     method.RetType,
		Indent:      "",
	}))

	// Function signature
	retType, err := g.typeMapper.MapReturnType(&method.RetType)
	if err != nil {
		return "", err
	}

	sb.WriteString(fmt.Sprintf("%s %s(", retType, method.Name))

	// Parameters
	var params []string
	for _, param := range method.ParamTypes {
		paramName := param.Name
		paramType, err := g.typeMapper.MapParamType(&param)
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
		paramName := param.Name
		paramType, _ := g.typeMapper.MapParamType(&param)

		// Check if we need conversion
		cType, _ := g.toCType(paramType, &manifest.RetType{
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

	sb.WriteString(fmt.Sprintf("__%s_%s(", pluginName, method.Name))
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

func (g *DlangGenerator) generateMethodAlias(method *manifest.Method, pluginName string) (string, error) {
	var sb strings.Builder

	// Get C ABI return type
	retType, _ := g.typeMapper.MapReturnType(&method.RetType)
	cRetType, _ := g.toCType(retType, &method.RetType, false)

	// Remove 'ref' prefix if present
	if strings.HasPrefix(cRetType, "ref ") {
		cRetType = strings.TrimPrefix(cRetType, "ref ")
	}

	sb.WriteString(fmt.Sprintf("alias _%s = extern (C) %s function(", method.Name, cRetType))

	// Parameters
	var params []string
	for _, param := range method.ParamTypes {
		paramName := param.Name
		paramType, _ := g.typeMapper.MapParamType(&param)

		cType, _ := g.toCType(paramType, &manifest.RetType{
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

	sb.WriteString(fmt.Sprintf("export __gshared _%s __%s_%s = null;\n", method.Name, pluginName, method.Name))

	return sb.String(), nil
}

// toCType converts D native types to C ABI types
func (g *DlangGenerator) toCType(nativeType string, typeInfo *manifest.RetType, isRef bool) (string, error) {
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
		"Vec2":   "ref Vec2",
		"Vec3":   "ref Vec3",
		"Vec4":   "ref Vec4",
		"Mat4x4": "ref Mat4x4",
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

func (g *DlangGenerator) generateClass(m *manifest.Manifest, class *manifest.Class) (string, error) {
	var sb strings.Builder

	// Get handle type and invalid value
	invalidValue, handleType, err := g.typeMapper.MapHandleType(class)
	if err != nil {
		return "", err
	}

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

	// Class documentation
	sb.WriteString(g.generateDocumentation(DocOptions{
		Description: class.Description,
		Deprecated:  class.Deprecated,
		Indent:      "",
	}))

	// Struct declaration
	sb.WriteString(fmt.Sprintf("struct %s {\n", class.Name))

	// Only generate handle members if class has a handle
	if hasHandle {
		// Private members
		sb.WriteString(fmt.Sprintf("\tprivate %s _handle = %s;\n", handleType, invalidValue))
		if hasDtor {
			sb.WriteString(fmt.Sprintf("\tprivate %s _ownership = %s.Borrowed;\n\n", OwnershipEnumName, OwnershipEnumName))

			// Default constructor
			hasDefaultConstructor := g.HasConstructorWithNoParam(m, class)
			if !hasDefaultConstructor {
				sb.WriteString("\tthis()\n\n")
			}

			// Disable default postblit to prevent accidental copies
			sb.WriteString("\t/// Disable default postblit to prevent accidental copies\n")
			sb.WriteString("\t@disable this(this);\n\n")
		} else {
			sb.WriteString("\n")
		}
	}

	// Only generate handle-related code if class has a handle
	if hasHandle {
		// Generate constructors from methods
		for _, ctorName := range class.Constructors {
			ctorCode, err := g.generateConstructor(m, class, ctorName)
			if err != nil {
				return "", err
			}
			sb.WriteString(ctorCode)
			sb.WriteString("\n")
		}

		// Constructor from handle
		sb.WriteString("\t/**\n")
		sb.WriteString(fmt.Sprintf("\t * Creates a %s from an existing handle\n", class.Name))
		sb.WriteString("\t * Params:\n")
		sb.WriteString(fmt.Sprintf("\t *   handle = The %s handle\n", class.Name))
		if hasDtor {
			sb.WriteString("\t *   ownership = Whether this wrapper owns the handle\n")
		} else {
			sb.WriteString("\t *   ownership = Ownership flag (unused in this version)\n")
		}
		sb.WriteString("\t */\n")
		if hasDtor {
			// Check if there's a constructor with exactly 1 param of handle type to avoid ambiguity
			hasHandleOnlyConstructor := g.HasConstructorWithSingleHandleParam(m, class)
			ownershipDefault := ""
			if !hasHandleOnlyConstructor {
				ownershipDefault = fmt.Sprintf(" = %s.Borrowed", OwnershipEnumName)
			}
			sb.WriteString(fmt.Sprintf("\tthis(%s handle, %s ownership%s) {\n", handleType, OwnershipEnumName, ownershipDefault))
			sb.WriteString("\t\t_handle = handle;\n")
			sb.WriteString("\t\t_ownership = ownership;\n")
			sb.WriteString("\t}\n\n")
		} else {
			ownershipTag := ""
			if hasCtor {
				ownershipTag = fmt.Sprintf(", Ownership ownership = %s.Borrowed", OwnershipEnumName)
			}
			sb.WriteString(fmt.Sprintf("\tthis(%s handle%s) {\n", handleType, ownershipTag))
			sb.WriteString("\t\t_handle = handle;\n")
			sb.WriteString("\t}\n\n")
		}

		// Destructor (only if needed)
		if hasDtor {
			sb.WriteString("\t~this() {\n")
			sb.WriteString("\t\tdestroy();\n")
			sb.WriteString("\t}\n\n")

			sb.WriteString(g.generateMoveSemantics(class.Name))
		}

		sb.WriteString(g.generateUtilityMethods(class.Name, handleType, invalidValue, hasDtor))
	}

	// Generate class bindings
	for _, binding := range class.Bindings {
		methodCode, err := g.generateBinding(m, class, &binding)
		if err != nil {
			return "", err
		}
		sb.WriteString(methodCode)
		sb.WriteString("\n")
	}

	// Private helper methods for RAII version
	if hasDtor {
		sb.WriteString(g.generatePrivateHelpers(class.Name, invalidValue, *class.Destructor))
	}

	sb.WriteString("}\n")

	return sb.String(), nil
}

func (g *DlangGenerator) generateConstructor(m *manifest.Manifest, class *manifest.Class, methodName string) (string, error) {
	// Find the method in the manifest
	method := FindMethod(m, methodName)
	if method == nil {
		return "", fmt.Errorf("constructor method %s not found", methodName)
	}

	var sb strings.Builder

	// Generate documentation
	description := method.Description
	if description == "" {
		description = fmt.Sprintf("Creates a new %s instance", class.Name)
	}

	sb.WriteString(g.generateDocumentation(DocOptions{
		Description: description,
		Deprecated:  method.Deprecated,
		Params:      method.ParamTypes,
		Indent:      "\t",
	}))

	// Constructor signature
	sb.WriteString("\tthis(")

	// Parameters
	var params []string
	for _, param := range method.ParamTypes {
		paramName := param.Name
		paramType, err := g.typeMapper.MapParamType(&param)
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

	// Constructor body - call the method and capture handle
	var callArgs []string
	for _, param := range method.ParamTypes {
		paramName := param.Name
		callArgs = append(callArgs, paramName)
	}

	sb.WriteString(fmt.Sprintf("\t\tthis(%s(", method.Name))
	sb.WriteString(strings.Join(callArgs, ", "))
	sb.WriteString(fmt.Sprintf("), %s.Owned);\n", OwnershipEnumName))
	sb.WriteString("\t}\n")

	return sb.String(), nil
}

func (g *DlangGenerator) generateBinding(m *manifest.Manifest, class *manifest.Class, binding *manifest.Binding) (string, error) {
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

	// Add deprecation attribute if present (check both binding and underlying method)
	deprecationReason := binding.Deprecated
	if deprecationReason == "" {
		deprecationReason = method.Deprecated
	}

	// Generate documentation
	sb.WriteString(g.generateDocumentation(DocOptions{
		Description: method.Description,
		Deprecated:  deprecationReason,
		Params:      methodParams,
		RetType:     method.RetType,
		Indent:      "\t",
		AddThrows:   binding.BindSelf,
	}))

	// Method signature
	retType, err := g.typeMapper.MapReturnType(&method.RetType)
	if err != nil {
		return "", err
	}

	// Handle return type alias
	if binding.RetAlias != nil && binding.RetAlias.Name != "" {
		retType = binding.RetAlias.Name
	}

	if !binding.BindSelf {
		sb.WriteString(fmt.Sprintf("\tstatic %s %s(", retType, binding.Name))
	} else {
		sb.WriteString(fmt.Sprintf("\t%s %s(", retType, binding.Name))
	}

	// Parameters (excluding self if bindSelf)
	var paramStrs []string
	for i, param := range methodParams {
		paramName := param.Name
		paramType, err := g.typeMapper.MapParamType(&param)
		paramRef := param.Ref
		paramMode := "ref"
		if err != nil {
			return "", err
		}

		// Check if this parameter has an alias
		if i < len(binding.ParamAliases) && binding.ParamAliases[i] != nil {
			paramType = binding.ParamAliases[i].Name
			paramRef = true
			if !binding.ParamAliases[i].Owner {
				paramMode = "const ref"
			}
		}

		if paramRef {
			paramStrs = append(paramStrs, fmt.Sprintf("%s %s %s", paramMode, paramType, paramName))
		} else {
			paramStrs = append(paramStrs, fmt.Sprintf("%s %s", paramType, paramName))
		}
	}

	sb.WriteString(strings.Join(paramStrs, ", "))
	sb.WriteString(") {\n")

	// Method body
	// Generate null check if needed (only for non-static methods)
	if binding.BindSelf {
		invalidValue, _, err := g.typeMapper.MapHandleType(class)
		if err != nil {
			return "", err
		}
		sb.WriteString(fmt.Sprintf("\t\tenforce(_handle !is %s, \"%s: %s\");\n", invalidValue, class.Name, EmptyHandleError))
	}

	// Build call arguments
	var callArgs []string
	if binding.BindSelf {
		callArgs = append(callArgs, "_handle")
	}

	for i, param := range methodParams {
		paramName := param.Name

		// Check if parameter has alias and needs .release() or .get()
		if i < len(binding.ParamAliases) && binding.ParamAliases[i] != nil {
			if binding.ParamAliases[i].Owner {
				callArgs = append(callArgs, paramName+".release()")
			} else {
				callArgs = append(callArgs, paramName+".get()")
			}
		} else {
			callArgs = append(callArgs, paramName)
		}
	}

	// Function call
	if method.RetType.Type != "void" {
		sb.WriteString("\t\treturn ")
	} else {
		sb.WriteString("\t\t")
	}

	//hasCtor := len(class.Constructors) > 0
	hasDtor := class.Destructor != nil

	// Check if the return value should be wrapped
	if binding.RetAlias != nil && binding.RetAlias.Name != "" {
		ownership := ""
		if hasDtor /*|| hasCtor*/ {
			if binding.RetAlias.Owner {
				ownership = fmt.Sprintf(", %s.Owned", OwnershipEnumName)
			} else {
				ownership = fmt.Sprintf(", %s.Borrowed", OwnershipEnumName)
			}
		}
		sb.WriteString(fmt.Sprintf("%s(%s(", binding.RetAlias.Name, method.Name))
		sb.WriteString(strings.Join(callArgs, ", "))
		sb.WriteString(fmt.Sprintf(")%s);\n", ownership))
	} else {
		sb.WriteString(fmt.Sprintf("%s(", method.Name))
		sb.WriteString(strings.Join(callArgs, ", "))
		sb.WriteString(");\n")
	}

	sb.WriteString("\t}\n")

	return sb.String(), nil
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

func (m *DlangTypeMapper) MapParamType(param *manifest.ParamType) (string, error) {
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

	// Regular type mapping
	return m.MapType(typeName, ctx, param.IsArray())
}

func (m *DlangTypeMapper) MapReturnType(retType *manifest.RetType) (string, error) {
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

	// Regular type mapping
	return m.MapType(typeName, TypeContextReturn, retType.IsArray())
}

func (m *DlangTypeMapper) MapHandleType(class *manifest.Class) (string, string, error) {
	invalidValue := class.InvalidValue
	handleType, err := m.MapType(class.HandleType, TypeContextReturn, false)
	if err != nil {
		return "", "", err
	}

	nullptr := invalidValue == "0" || invalidValue == "" || invalidValue == "NULL" || invalidValue == "nullptr"
	if strings.HasPrefix(class.HandleType, "ptr") && nullptr {
		invalidValue = "null"
	} else if invalidValue == "" {
		invalidValue = fmt.Sprintf("%s.init", handleType)
	}

	return invalidValue, handleType, err
}
