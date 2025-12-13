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

	// Secondly, generate delegates file
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
	sb.WriteString("enum Ownership : bool {\n")
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

func (g *DlangGenerator) generateDelegatesFile(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	moduleName := strings.ToLower(m.Name)
	sb.WriteString(fmt.Sprintf("module imported.%s.delegates;\n\n", moduleName))
	sb.WriteString(fmt.Sprintf("public import imported.%s.enums;\n\n", moduleName))

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

func (g *DlangGenerator) generatePackageFile(m *manifest.Manifest, moduleName string, groups map[string]bool) string {
	var sb strings.Builder

	// Module declaration
	sb.WriteString(fmt.Sprintf("module imported.%s;\n\n", moduleName))

	// Always import enums first
	sb.WriteString(fmt.Sprintf("public import imported.%s.enums;\n", moduleName))
	sb.WriteString(fmt.Sprintf("public import imported.%s.delegates;\n", moduleName))

	// Import all group modules
	for groupName := range groups {
		sb.WriteString(fmt.Sprintf("public import imported.%s.%s;\n", moduleName, groupName))
	}

	return sb.String()
}

func (g *DlangGenerator) generateEnum(enum *manifest.EnumType, underlyingType string) (string, error) {
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

func (g *DlangGenerator) generateModuleFile(m *manifest.Manifest, moduleName, groupName string, opts *GeneratorOptions) (string, error) {
	var sb strings.Builder

	// Module declaration
	sb.WriteString(fmt.Sprintf("module imported.%s.%s;\n\n", moduleName, groupName))

	// Imports
	sb.WriteString("import plugify.internals;\n")
	sb.WriteString("public import plugify;\n")
	sb.WriteString(fmt.Sprintf("public import imported.%s.enums;\n", moduleName))
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
		methodGroup := g.GetGroupName(method.Group)
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

func (g *DlangGenerator) generateMethodWrapper(method *manifest.Method, pluginName string) (string, error) {
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

	sb.WriteString(fmt.Sprintf("%s %s(", retType, g.SanitizeName(method.Name)))

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

	sb.WriteString(fmt.Sprintf("export __gshared _%s __%s_%s = null;\n", method.Name, pluginName, method.Name))

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
	sb.WriteString("/**\n")
	if class.Description != "" {
		sb.WriteString(fmt.Sprintf(" * %s\n", class.Description))
	}
	if hasHandle {
		if hasDtor {
			sb.WriteString(fmt.Sprintf(" * RAII wrapper for %s handle.\n", class.Name))
		} else {
			sb.WriteString(fmt.Sprintf(" * Wrapper for %s handle.\n", class.Name))
		}
	} else {
		sb.WriteString(fmt.Sprintf(" * Static utility class for %s.\n", class.Name))
	}
	sb.WriteString(" */\n")

	// Struct declaration
	sb.WriteString(fmt.Sprintf("struct %s {\n", class.Name))

	// Only generate handle members if class has a handle
	if hasHandle {
		// Private members
		sb.WriteString(fmt.Sprintf("\tprivate %s _handle = %s;\n", handleType, invalidValue))
		if hasDtor {
			sb.WriteString("\tprivate Ownership _ownership = Ownership.Borrowed;\n\n")

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
				ownershipDefault = " = Ownership.Borrowed"
			}
			sb.WriteString(fmt.Sprintf("\tthis(%s handle, Ownership ownership%s) {\n", handleType, ownershipDefault))
			sb.WriteString("\t\t_handle = handle;\n")
			sb.WriteString("\t\t_ownership = ownership;\n")
			sb.WriteString("\t}\n\n")
		} else {
			ownershipTag := ""
			if hasCtor {
				ownershipTag = ", Ownership ownership = Ownership.Borrowed"
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

			// Move constructor
			sb.WriteString("\t/// Move constructor (called when moving)\n")
			sb.WriteString(fmt.Sprintf("\tthis(ref return scope %s other) {\n", class.Name))
			sb.WriteString("\t\t_handle = other._handle;\n")
			sb.WriteString("\t\t_ownership = other._ownership;\n")
			sb.WriteString("\t\tother.nullify();\n")
			sb.WriteString("\t}\n\n")

			// Move assignment
			sb.WriteString("\t/// Move assignment\n")
			sb.WriteString(fmt.Sprintf("\tref %s opAssign(%s other) return {\n", class.Name, class.Name))
			sb.WriteString("\t\tswap(this, other);\n")
			sb.WriteString("\t\treturn this;\n")
			sb.WriteString("\t}\n\n")
		}

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
		sb.WriteString("\t/// Swap two " + class.Name + " instances\n")
		sb.WriteString(fmt.Sprintf("\tvoid swap(ref %s other) nothrow @nogc {\n", class.Name))
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
		sb.WriteString(fmt.Sprintf("\tint opCmp(ref const %s other) const pure nothrow @nogc {\n", class.Name))
		sb.WriteString("\t\tif (_handle < other._handle) return -1;\n")
		sb.WriteString("\t\tif (_handle > other._handle) return 1;\n")
		sb.WriteString("\t\treturn 0;\n")
		sb.WriteString("\t}\n\n")

		sb.WriteString(fmt.Sprintf("\tbool opEquals(ref const %s other) const pure nothrow @nogc {\n", class.Name))
		sb.WriteString("\t\treturn _handle == other._handle;\n")
		sb.WriteString("\t}\n\n")
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
		sb.WriteString("\tprivate void destroy() const nothrow {\n")
		sb.WriteString(fmt.Sprintf("\t\tif (_handle !is %s && _ownership == Ownership.Owned) {\n", invalidValue))
		sb.WriteString(fmt.Sprintf("\t\t\t%s(_handle);\n", *class.Destructor))
		sb.WriteString("\t\t}\n")
		sb.WriteString("\t}\n\n")

		sb.WriteString("\tprivate void nullify() nothrow @nogc {\n")
		sb.WriteString(fmt.Sprintf("\t\t_handle = %s;\n", invalidValue))
		sb.WriteString("\t\t_ownership = Ownership.Borrowed;\n")
		sb.WriteString("\t}\n")
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
	sb.WriteString("\t/**\n")
	if method.Description != "" {
		sb.WriteString(fmt.Sprintf("\t * %s\n", method.Description))
	} else {
		sb.WriteString(fmt.Sprintf("\t * Creates a new %s instance\n", class.Name))
	}

	// Document parameters
	if len(method.ParamTypes) > 0 {
		sb.WriteString("\t * Params:\n")
		for _, param := range method.ParamTypes {
			paramName := g.SanitizeName(param.Name)
			description := param.Description
			if description == "" {
				description = "Parameter"
			}
			sb.WriteString(fmt.Sprintf("\t *   %s = %s\n", paramName, description))
		}
	}
	sb.WriteString("\t */\n")

	// Constructor signature
	sb.WriteString("\tthis(")

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

	// Constructor body - call the method and capture handle
	callArgs := []string{}
	for _, param := range method.ParamTypes {
		paramName := g.SanitizeName(param.Name)
		callArgs = append(callArgs, paramName)
	}

	sb.WriteString(fmt.Sprintf("\t\tthis(%s(", method.Name))
	sb.WriteString(strings.Join(callArgs, ", "))
	sb.WriteString("), Ownership.Owned);\n")
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

	// Generate documentation
	sb.WriteString("\t/**\n")
	if method.Description != "" {
		sb.WriteString(fmt.Sprintf("\t * %s\n", method.Description))
	}

	// Document parameters (excluding self if bindSelf)
	if len(methodParams) > 0 {
		sb.WriteString("\t * Params:\n")
		for _, param := range methodParams {
			paramName := g.SanitizeName(param.Name)
			description := param.Description
			if description == "" {
				description = "Parameter"
			}
			sb.WriteString(fmt.Sprintf("\t *   %s = %s\n", paramName, description))
		}
	}

	// Return documentation
	if method.RetType.Type != "void" {
		sb.WriteString("\t * Returns:\n")
		if method.RetType.Description != "" {
			sb.WriteString(fmt.Sprintf("\t *   %s\n", method.RetType.Description))
		} else {
			sb.WriteString("\t *   Return value\n")
		}
	}

	// Determine if method is non static
	if binding.BindSelf {
		sb.WriteString("\t * Throws: Exception if handle is null\n")
	}

	sb.WriteString("\t */\n")

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
	paramStrs := []string{}
	for i, param := range methodParams {
		paramName := g.SanitizeName(param.Name)
		paramType, err := g.typeMapper.MapParamType(&param, TypeContextValue)
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
		sb.WriteString(fmt.Sprintf("\t\tenforce(_handle !is %s, \"%s: Empty handle\");\n", invalidValue, class.Name))
	}

	// Build call arguments
	callArgs := []string{}
	if binding.BindSelf {
		callArgs = append(callArgs, "_handle")
	}

	for i, param := range methodParams {
		paramName := g.SanitizeName(param.Name)

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

	hasCtor := len(class.Constructors) > 0
	hasDtor := class.Destructor != nil

	// Check if the return value should be wrapped
	if binding.RetAlias != nil && binding.RetAlias.Name != "" {
		ownership := ""
		if hasDtor || hasCtor {
			if binding.RetAlias.Owner {
				ownership = ", Ownership.Owned"
			} else {
				ownership = ", Ownership.Borrowed"
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

func (m *DlangTypeMapper) MapHandleType(class *manifest.Class) (string, string, error) {
	invalidValue := class.InvalidValue
	handleType, err := m.MapType(class.HandleType, TypeContextReturn, false)
	if err != nil {
		return "", "", err
	}

	nullptr := invalidValue == "0" || invalidValue == "" || invalidValue == "NULL" || invalidValue == "nullptr"
	if class.HandleType == "ptr64" && nullptr {
		invalidValue = "null"
	} else if invalidValue == "" {
		invalidValue = fmt.Sprintf("%s.init", handleType)
	}

	return invalidValue, handleType, err
}
