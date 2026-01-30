package generator

import (
	"fmt"
	"strings"

	"github.com/untrustedmodders/plugify-gen/pkg/manifest"
)

// CppGenerator generates C++ header files
type CppGenerator struct {
	*BaseGenerator
}

// NewCppGenerator creates a new C++ generator
func NewCppGenerator() *CppGenerator {
	return &CppGenerator{
		BaseGenerator: NewBaseGenerator("cpp", NewCppCommonTypeMapper(), CppReservedWords),
	}
}

// Generate generates C++ bindings
func (g *CppGenerator) Generate(m *manifest.Manifest, opts *GeneratorOptions) (*GeneratorResult, error) {
	g.ResetCaches()
	m.Sanitize(g.Sanitizer)
	opts = EnsureOptions(opts)

	// Collect all unique groups from both methods and classes
	groups := g.GetGroups(m)

	files := make(map[string]string)
	folder := fmt.Sprintf("external/plugify/include/%s", m.Name)

	// Generate separate enums file
	enumsCode, err := g.generateEnumsFile(m)
	if err != nil {
		return nil, fmt.Errorf("generating enums file: %w", err)
	}
	files[fmt.Sprintf("%s/%s/enums.hpp", folder, m.Name)] = enumsCode

	// Generate separate aliases file
	aliasesCode, err := g.generateAliasesFile(m)
	if err != nil {
		return nil, fmt.Errorf("generating aliases file: %w", err)
	}
	files[fmt.Sprintf("%s/%s/aliases.hpp", folder, m.Name)] = aliasesCode

	// Generate separate delegates file
	delegatesCode, err := g.generateDelegatesFile(m)
	if err != nil {
		return nil, fmt.Errorf("generating delegates file: %w", err)
	}
	files[fmt.Sprintf("%s/%s/delegates.hpp", folder, m.Name)] = delegatesCode

	// Generate a file for each group
	for groupName := range groups {
		groupCode, err := g.generateGroupFile(m, groupName, opts)
		if err != nil {
			return nil, fmt.Errorf("failed to generate group %s: %w", groupName, err)
		}
		files[fmt.Sprintf("%s/%s/%s.hpp", folder, m.Name, groupName)] = groupCode
	}

	// Generate main header that includes all pieces
	mainHeader, err := g.generateMainHeader(m, groups)
	if err != nil {
		return nil, fmt.Errorf("generating main header: %w", err)
	}
	files[fmt.Sprintf("%s/%s.hpp", folder, m.Name)] = mainHeader

	// Generate main impl that includes all pieces
	mainImpl, err := g.generateMainImpl(m)
	if err != nil {
		return nil, fmt.Errorf("generating main impl: %w", err)
	}
	files[fmt.Sprintf("%s/%s.cpp", folder, m.Name)] = mainImpl

	result := &GeneratorResult{
		Files: files,
	}

	return result, nil
}

// generateDocumentation generates C++ Doxygen-style documentation comments (/** */)
func (g *CppGenerator) generateDocumentation(opts DocOptions) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("%s/**\n", opts.Indent))

	// Brief description
	if opts.Description != "" {
		sb.WriteString(fmt.Sprintf("%s * @brief %s\n", opts.Indent, opts.Description))
	}

	// Parameters
	for i, param := range opts.Params {
		paramType := param.Type
		if param.Ref {
			paramType += "&"
		}

		// Apply parameter alias if provided
		if i < len(opts.ParamAliases) && opts.ParamAliases[i] != nil {
			paramType = opts.ParamAliases[i].Name
		}

		sb.WriteString(fmt.Sprintf("%s * @param %s (%s)", opts.Indent, param.Name, paramType))
		if param.Description != "" {
			sb.WriteString(fmt.Sprintf(": %s", param.Description))
		}
		sb.WriteString("\n")
	}

	// Return type section
	if opts.RetType.Type != "void" {
		returnType := opts.RetType.Type

		// Apply return alias if provided
		if opts.RetAlias != nil && opts.RetAlias.Name != "" {
			returnType = opts.RetAlias.Name
		}

		sb.WriteString(fmt.Sprintf("%s * @return %s", opts.Indent, returnType))
		if opts.RetType.Description != "" {
			sb.WriteString(fmt.Sprintf(": %s", opts.RetType.Description))
		}
		sb.WriteString("\n")
	}

	sb.WriteString(fmt.Sprintf("%s */\n", opts.Indent))

	// Add deprecation attribute if present
	if opts.Deprecated != "" {
		sb.WriteString(fmt.Sprintf("    [[deprecated(\"%s\")]]\n", opts.Deprecated))
	}

	return sb.String()
}

// buildCallArguments builds the call arguments string for method invocations
func (g *CppGenerator) buildCallArguments(params []manifest.ParamType, aliases []*manifest.ParamAlias, bindSelf bool) string {
	var parts []string

	if bindSelf {
		parts = append(parts, "_handle")
	}

	for i, param := range params {
		var arg string

		// Check for alias
		if i < len(aliases) && aliases[i] != nil {
			if aliases[i].Owner {
				arg = param.Name + ".release()"
			} else {
				arg = param.Name + ".get()"
			}
		} else {
			arg = param.Name
		}
		parts = append(parts, arg)
	}

	return strings.Join(parts, ", ")
}

// generateUtilityMethods generates get, release, reset, and swap methods
func (g *CppGenerator) generateUtilityMethods(className, invalidValue, handleType string, hasDtor bool) string {
	var sb strings.Builder

	// get method
	sb.WriteString(fmt.Sprintf("    [[nodiscard]] auto get() const noexcept { return _handle; }\n\n"))

	// release method
	sb.WriteString("    [[nodiscard]] auto release() noexcept {\n")
	sb.WriteString("      auto handle = _handle;\n")
	if hasDtor {
		sb.WriteString("      nullify();\n")
	} else {
		sb.WriteString(fmt.Sprintf("      _handle = %s;\n", invalidValue))
	}
	sb.WriteString("      return handle;\n")
	sb.WriteString("    }\n\n")

	// reset method
	sb.WriteString("    void reset() noexcept {\n")
	if hasDtor {
		sb.WriteString("      destroy();\n")
		sb.WriteString("      nullify();\n")
	} else {
		sb.WriteString(fmt.Sprintf("      _handle = %s;\n", invalidValue))
	}
	sb.WriteString("    }\n\n")

	// swap method
	sb.WriteString(fmt.Sprintf("    void swap(%s& other) noexcept {\n", className))
	sb.WriteString("      using std::swap;\n")
	sb.WriteString("      swap(_handle, other._handle);\n")
	if hasDtor {
		sb.WriteString("      swap(_ownership, other._ownership);\n")
	}
	sb.WriteString("    }\n\n")

	sb.WriteString(fmt.Sprintf("    friend void swap(%s& lhs, %s& rhs) noexcept { lhs.swap(rhs); }\n\n", className, className))

	return sb.String()
}

// generateComparisonOperators generates operator bool, operator<=>, and operator==
func (g *CppGenerator) generateComparisonOperators(className, invalidValue string) string {
	return fmt.Sprintf(`    explicit operator bool() const noexcept { return _handle != %s; }
    [[nodiscard]] auto operator<=>(const %s& other) const noexcept { return _handle <=> other._handle; }
    [[nodiscard]] bool operator==(const %s& other) const noexcept { return _handle == other._handle; }

`, invalidValue, className, className)
}

// generateCopyMoveSemantics generates copy/move constructors and assignment operators
func (g *CppGenerator) generateCopyMoveSemantics(className string, hasDtor bool) string {
	if hasDtor {
		return fmt.Sprintf(`    ~%s() {
      destroy();
    }

    %s(const %s&) = delete;
    %s& operator=(const %s&) = delete;

    %s(%s&& other) noexcept
      : _handle(other._handle)
      , _ownership(other._ownership) {
      other.nullify();
    }

    %s& operator=(%s&& other) noexcept {
      if (this != &other) {
        destroy();
        _handle = other._handle;
        _ownership = other._ownership;
        other.nullify();
      }
      return *this;
    }

`, className, className, className, className, className,
			className, className, className, className)
	}

	// Without destructor (defaulted)
	return fmt.Sprintf(`    %s(const %s&) = default;
    %s& operator=(const %s&) = default;
    %s(%s&&) noexcept = default;
    %s& operator=(%s&&) noexcept = default;
    ~%s() = default;

`, className, className, className, className, className,
		className, className, className, className)
}

// generatePrivateHelpers generates private helper methods (destroy and nullify)
func (g *CppGenerator) generatePrivateHelpers(invalidValue, destructor string) string {
	return fmt.Sprintf(`    void destroy() const noexcept {
      if (_handle != %s && _ownership == %s::Owned) {
        %s(_handle);
      }
    }

    void nullify() noexcept {
      _handle = %s;
      _ownership = %s::Borrowed;
    }
`, invalidValue, OwnershipEnumName, destructor, invalidValue, OwnershipEnumName)
}

func (g *CppGenerator) generateEnums(m *manifest.Manifest) (string, error) {
	// Use the base generator's CollectEnums helper
	return g.CollectEnums(m, g.generateEnum)
}

func (g *CppGenerator) generateEnum(enum *manifest.Enum, underlyingType string) (string, error) {
	var sb strings.Builder

	if enum.Description != "" {
		sb.WriteString(fmt.Sprintf("  // %s\n", enum.Description))
	}

	sb.WriteString(fmt.Sprintf("  enum class %s : %s {\n", enum.Name, underlyingType))

	for i, val := range enum.Values {
		if val.Description != "" {
			sb.WriteString(fmt.Sprintf("    %s = %d, // %s\n", val.Name, val.Value, val.Description))
		} else {
			sb.WriteString(fmt.Sprintf("    %s = %d", val.Name, val.Value))
			if i < len(enum.Values)-1 {
				sb.WriteString(",")
			}
			sb.WriteString("\n")
		}
	}

	sb.WriteString("  };\n")
	return sb.String(), nil
}

func (g *CppGenerator) generateAliases(m *manifest.Manifest) (string, error) {
	// Use the base generator's CollectAliases helper
	return g.CollectAliases(m, g.generateAlias)
}

func (g *CppGenerator) generateAlias(alias *manifest.Alias, underlyingType string) (string, error) {
	var sb strings.Builder

	if alias.Description != "" {
		sb.WriteString(fmt.Sprintf("  // %s\n", alias.Description))
	}

	sb.WriteString(fmt.Sprintf("  using %s = %s;\n", alias.Name, underlyingType))

	return sb.String(), nil
}

func (g *CppGenerator) generateDelegates(m *manifest.Manifest) (string, error) {
	return g.CollectDelegates(m, g.generateDelegate)
}

func (g *CppGenerator) generateDelegate(proto *manifest.Prototype) (string, error) {
	var sb strings.Builder

	if proto.Description != "" {
		sb.WriteString(fmt.Sprintf("  // %s\n", proto.Description))
	}

	// Generate return type
	retType, err := g.typeMapper.MapReturnType(&proto.RetType)
	if err != nil {
		return "", err
	}

	// Generate parameters
	params, err := FormatParameters(proto.ParamTypes, ParamFormatTypes, g.typeMapper)
	if err != nil {
		return "", err
	}

	sb.WriteString(fmt.Sprintf("  using %s = %s (*)(%s);\n\n", proto.Name, retType, params))
	return sb.String(), nil
}

func (g *CppGenerator) generateMethod(pluginName string, method *manifest.Method) (string, error) {
	var sb strings.Builder

	formattedParams, err := FormatParameters(method.ParamTypes, ParamFormatTypesAndNames, g.typeMapper)
	if err != nil {
		return "", err
	}

	// Generate function signature
	retType, err := g.typeMapper.MapReturnType(&method.RetType)
	if err != nil {
		return "", err
	}

	// Generate type alias for function pointer
	funcTypeParams, err := FormatParameters(method.ParamTypes, ParamFormatTypes, g.typeMapper)
	if err != nil {
		return "", err
	}
	sb.WriteString(fmt.Sprintf("  using _%s = %s (*)(%s);\n", method.Name, retType, funcTypeParams))
	sb.WriteString("}\n")

	// Generate exported wrapper function
	paramNames, err := FormatParameters(method.ParamTypes, ParamFormatNames, g.typeMapper)
	if err != nil {
		return "", err
	}

	// Generate global exported function pointer
	sb.WriteString(fmt.Sprintf("extern \"C\" PLUGIN_API %s::_%s __%s_%s;\n", pluginName, method.Name, pluginName, method.Name))
	sb.WriteString(fmt.Sprintf("namespace %s {\n", pluginName))

	sb.WriteString(g.generateDocumentation(DocOptions{
		Description: method.Description,
		Deprecated:  method.Deprecated,
		Params:      method.ParamTypes,
		RetType:     method.RetType,
		Indent:      "  ",
	}))

	sb.WriteString(fmt.Sprintf("  inline %s %s(%s) {\n", retType, method.Name, formattedParams))
	if method.RetType.Type == "void" {
		sb.WriteString(fmt.Sprintf("    return __%s_%s(%s);\n", pluginName, method.Name, paramNames))
	} else {
		sb.WriteString(fmt.Sprintf("    return __%s_%s(%s);\n", pluginName, method.Name, paramNames))
	}
	sb.WriteString("  }\n")

	return sb.String(), nil
}

func (g *CppGenerator) generateClasses(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	for i, class := range m.Classes {
		if i > 0 {
			sb.WriteString("\n")
		}
		classCode, err := g.generateClass(m, &class)
		if err != nil {
			return "", fmt.Errorf("failed to generate class %s: %w", class.Name, err)
		}
		sb.WriteString(classCode)
	}

	return sb.String(), nil
}

func (g *CppGenerator) generateClass(m *manifest.Manifest, class *manifest.Class) (string, error) {
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
		Indent:      "  ",
	}))

	// Class declaration
	sb.WriteString(fmt.Sprintf("  class %s final {\n", class.Name))

	sb.WriteString("  public:\n")

	// Only generate handle-related code if class has a handle
	if hasHandle {
		// Default constructor
		hasDefaultConstructor := g.HasConstructorWithNoParam(m, class)
		if !hasDefaultConstructor {
			sb.WriteString(fmt.Sprintf("    %s() = default;\n\n", class.Name))
		}

		// Generate constructors from methods
		for _, ctorName := range class.Constructors {
			ctorCode, err := g.generateConstructor(m, class, ctorName)
			if err != nil {
				return "", err
			}
			sb.WriteString(ctorCode)
		}

		// Destructor and copy/move semantics
		sb.WriteString(g.generateCopyMoveSemantics(class.Name, hasDtor))

		// Constructor from handle
		if hasDtor {
			// Check if there's a constructor with exactly 1 param of handle type to avoid ambiguity
			hasHandleOnlyConstructor := g.HasConstructorWithSingleHandleParam(m, class)
			ownershipDefault := ""
			if !hasHandleOnlyConstructor {
				ownershipDefault = fmt.Sprintf(" = %s::Borrowed", OwnershipEnumName)
			}
			sb.WriteString(fmt.Sprintf("    %s(%s handle, %s ownership%s) : _handle(handle), _ownership(ownership) {}\n\n", class.Name, handleType, OwnershipEnumName, ownershipDefault))
		} else {
			ownershipTag := ""
			if hasCtor {
				ownershipTag = fmt.Sprintf(", [[maybe_unused]] %s ownership = %s::Borrowed", OwnershipEnumName, OwnershipEnumName)
			}
			sb.WriteString(fmt.Sprintf("    explicit %s(%s handle%s) : _handle(handle) {}\n\n", class.Name, handleType, ownershipTag))
		}

		// Utility methods
		sb.WriteString(g.generateUtilityMethods(class.Name, invalidValue, handleType, hasDtor))

		// Comparison operators
		sb.WriteString(g.generateComparisonOperators(class.Name, invalidValue))
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

	// Only generate private section if class has a handle
	if hasHandle {
		// Private section
		sb.WriteString("  private:\n")

		if hasDtor {
			sb.WriteString(g.generatePrivateHelpers(invalidValue, *class.Destructor))
			sb.WriteString("\n")
		}

		// Member variables
		sb.WriteString(fmt.Sprintf("    %s _handle{%s};\n", handleType, invalidValue))
		if hasDtor {
			sb.WriteString(fmt.Sprintf("    %s _ownership{%s::Borrowed};\n", OwnershipEnumName, OwnershipEnumName))
		}
	}

	sb.WriteString("  };\n\n")

	return sb.String(), nil
}

func (g *CppGenerator) generateConstructor(m *manifest.Manifest, class *manifest.Class, methodName string) (string, error) {
	// Find the method in the manifest
	method := FindMethod(m, methodName)
	if method == nil {
		return "", fmt.Errorf("constructor method %s not found", methodName)
	}

	var sb strings.Builder

	// Generate documentation
	sb.WriteString(g.generateDocumentation(DocOptions{
		Description: method.Description,
		Deprecated:  method.Deprecated,
		Params:      method.ParamTypes,
		Indent:      "    ",
	}))

	// Generate constructor signature
	formattedParams, err := FormatParameters(method.ParamTypes, ParamFormatTypesAndNames, g.typeMapper)
	if err != nil {
		return "", err
	}

	sb.WriteString(fmt.Sprintf("    explicit %s(%s)\n", class.Name, formattedParams))

	// Generate initialization list
	paramNames, err := FormatParameters(method.ParamTypes, ParamFormatNames, g.typeMapper)
	if err != nil {
		return "", err
	}

	sb.WriteString(fmt.Sprintf("      : %s(%s(%s), %s::Owned) {}\n\n", class.Name, method.FuncName, paramNames, OwnershipEnumName))

	return sb.String(), nil
}

func (g *CppGenerator) generateBinding(m *manifest.Manifest, class *manifest.Class, binding *manifest.Binding) (string, error) {
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

	sb.WriteString(g.generateDocumentation(DocOptions{
		Description:  method.Description,
		Deprecated:   deprecationReason,
		Params:       methodParams,
		RetType:      method.RetType,
		ParamAliases: binding.ParamAliases,
		RetAlias:     binding.RetAlias,
		Indent:       "    ",
	}))

	// Generate method signature
	retType, err := g.typeMapper.MapReturnType(&method.RetType)
	if err != nil {
		return "", err
	}

	// Handle return type alias
	if binding.RetAlias != nil && binding.RetAlias.Name != "" {
		retType = binding.RetAlias.Name
	}

	formattedParams, err := FormatParameters(methodParams, ParamFormatTypesAndNames, g.typeMapper)
	if err != nil {
		return "", err
	}

	// Apply parameter aliases to signature
	if len(binding.ParamAliases) > 0 {
		formattedParams = g.applyParamAliases(formattedParams, methodParams, binding.ParamAliases)
	}

	// Determine if method is static
	if !binding.BindSelf {
		sb.WriteString(fmt.Sprintf("    static %s %s(%s) {\n", retType, binding.Name, formattedParams))
	} else {
		sb.WriteString(fmt.Sprintf("    %s %s(%s) {\n", retType, binding.Name, formattedParams))
	}

	// Generate null check if needed
	nullPolicy := class.NullPolicy
	if nullPolicy == "" {
		nullPolicy = "throw"
	}

	if binding.BindSelf && nullPolicy == "throw" {
		invalidValue, _, err := g.typeMapper.MapHandleType(class)
		if err != nil {
			return "", err
		}
		sb.WriteString(fmt.Sprintf("      if (_handle == %s) throw std::runtime_error(\"%s: %s\");\n", invalidValue, class.Name, EmptyHandleError))
	}

	// Build call arguments
	callArgs := g.buildCallArguments(methodParams, binding.ParamAliases, binding.BindSelf)

	//hasCtor := len(class.Constructors) > 0
	hasDtor := class.Destructor != nil

	// Generate return statement
	if method.RetType.Type == "void" {
		sb.WriteString(fmt.Sprintf("      %s::%s(%s);\n", m.Name, method.FuncName, callArgs))
	} else {
		// Handle return alias
		if binding.RetAlias != nil && binding.RetAlias.Name != "" {
			ownership := ""
			if hasDtor /*|| hasCtor*/ {
				if binding.RetAlias.Owner {
					ownership = fmt.Sprintf(", %s::Owned", OwnershipEnumName)
				} else {
					ownership = fmt.Sprintf(", %s::Borrowed", OwnershipEnumName)
				}
			}
			sb.WriteString(fmt.Sprintf("      return %s(%s::%s(%s)%s);\n", binding.RetAlias.Name, m.Name, method.FuncName, callArgs, ownership)) // always pass ownership just as a tag
		} else {
			sb.WriteString(fmt.Sprintf("      return %s::%s(%s);\n", m.Name, method.FuncName, callArgs))
		}
	}

	sb.WriteString("    }\n")

	return sb.String(), nil
}

func (g *CppGenerator) applyParamAliases(formattedParams string, params []manifest.ParamType, aliases []*manifest.ParamAlias) string {
	// This is a simplified implementation
	// In reality, you might need more sophisticated parsing
	result := formattedParams

	for i, param := range params {
		if i < len(aliases) && aliases[i] != nil {
			// Build the parameter type to search for
			paramType, _ := g.typeMapper.MapParamType(&param)

			// Determine the replacement type
			replacementType := aliases[i].Name
			if aliases[i].Owner {
				replacementType = aliases[i].Name + "&&"
			} else {
				replacementType = "const " + aliases[i].Name + "&"
			}

			// Replace in the formatted params
			oldPattern := fmt.Sprintf("%s %s", paramType, param.Name)
			newPattern := fmt.Sprintf("%s %s", replacementType, param.Name)
			result = strings.ReplaceAll(result, oldPattern, newPattern)
		}
	}

	return result
}

// generateEnumsFile generates a file containing all enums
func (g *CppGenerator) generateEnumsFile(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	// Header guard and includes
	sb.WriteString("#pragma once\n\n")
	sb.WriteString("#include <cstdint>\n")
	sb.WriteString(fmt.Sprintf("// Generated from %s.pplugin\n\n", m.Name))

	// Namespace
	sb.WriteString(fmt.Sprintf("namespace %s {\n\n", m.Name))

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
	sb.WriteString(fmt.Sprintf("  /// %s type for RAII wrappers\n", OwnershipEnumName))
	sb.WriteString(fmt.Sprintf("  enum class %s : bool { Borrowed, Owned };\n\n", OwnershipEnumName))

	// Close namespace
	sb.WriteString(fmt.Sprintf("} // namespace %s\n", m.Name))

	return sb.String(), nil
}

// generateAliasesFile generates a file containing all aliases
func (g *CppGenerator) generateAliasesFile(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	// Header guard and includes
	sb.WriteString("#pragma once\n\n")
	sb.WriteString("#include <cstdint>\n")
	sb.WriteString(fmt.Sprintf("// Generated from %s.pplugin\n\n", m.Name))

	// Namespace
	sb.WriteString(fmt.Sprintf("namespace %s {\n\n", m.Name))

	// Generate aliases
	aliasesCode, err := g.generateAliases(m)
	if err != nil {
		return "", err
	}
	if aliasesCode != "" {
		sb.WriteString(aliasesCode)
		sb.WriteString("\n")
	}

	// Close namespace
	sb.WriteString(fmt.Sprintf("} // namespace %s\n", m.Name))

	return sb.String(), nil
}

// generateDelegatesFile generates a file containing all delegate typedefs
func (g *CppGenerator) generateDelegatesFile(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	// Header guard and includes
	sb.WriteString("#pragma once\n\n")
	sb.WriteString("#include \"enums.hpp\"\n")
	sb.WriteString("#include \"aliases.hpp\"\n")
	sb.WriteString("#include <plg/plugin.hpp>\n")
	sb.WriteString("#include <plg/any.hpp>\n\n")
	sb.WriteString(fmt.Sprintf("// Generated from %s.pplugin\n\n", m.Name))

	// Namespace
	sb.WriteString(fmt.Sprintf("namespace %s {\n\n", m.Name))

	// Generate delegates
	delegatesCode, err := g.generateDelegates(m)
	if err != nil {
		return "", err
	}
	if delegatesCode != "" {
		sb.WriteString(delegatesCode)
	}

	// Close namespace
	sb.WriteString(fmt.Sprintf("} // namespace %s\n", m.Name))

	return sb.String(), nil
}

// generateGroupFile generates a file for a specific group containing methods and classes
func (g *CppGenerator) generateGroupFile(m *manifest.Manifest, groupName string, opts *GeneratorOptions) (string, error) {
	var sb strings.Builder

	// Header guard and includes
	sb.WriteString("#pragma once\n\n")
	sb.WriteString("#include \"enums.hpp\"\n")
	sb.WriteString("#include \"aliases.hpp\"\n")
	sb.WriteString("#include \"delegates.hpp\"\n")
	sb.WriteString("#include <plugin_export.h>\n\n")

	// Find which other groups this group depends on (for method calls from classes)
	if len(m.Classes) > 0 {
		dependentGroups := g.FindDependentGroups(m, groupName)
		if len(dependentGroups) > 0 {
			for depGroup := range dependentGroups {
				sb.WriteString(fmt.Sprintf("#include \"%s.hpp\"\n", depGroup))
			}
			sb.WriteString("\n")
		}
	}

	sb.WriteString(fmt.Sprintf("// Generated from %s.pplugin (group: %s)\n\n", m.Name, groupName))

	// Namespace
	sb.WriteString(fmt.Sprintf("namespace %s {\n\n", m.Name))

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

	// Close namespace
	sb.WriteString(fmt.Sprintf("} // namespace %s\n", m.Name))

	return sb.String(), nil
}

// generateMainHeader generates the main header file that includes all group files
func (g *CppGenerator) generateMainHeader(m *manifest.Manifest, groups map[string]struct{}) (string, error) {
	var sb strings.Builder

	// Header guard
	sb.WriteString("#pragma once\n\n")
	sb.WriteString(fmt.Sprintf("// Generated from %s.pplugin\n", m.Name))
	sb.WriteString("// This header includes all generated components\n\n")

	// Include enums, aliases and delegates first
	sb.WriteString(fmt.Sprintf("#include \"%s/enums.hpp\"\n", m.Name))
	sb.WriteString(fmt.Sprintf("#include \"%s/aliases.hpp\"\n", m.Name))
	sb.WriteString(fmt.Sprintf("#include \"%s/delegates.hpp\"\n", m.Name))

	// Import all group headers
	for groupName := range groups {
		sb.WriteString(fmt.Sprintf("#include \"%s/%s.hpp\"\n", m.Name, groupName))
	}

	return sb.String(), nil
}

// generateMainImpl generates the main impl file that includes all group files
func (g *CppGenerator) generateMainImpl(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("#include \"%s.hpp\"\n\n", m.Name))

	for i := range m.Methods {
		method := &m.Methods[i]

		// Generate global exported function pointer impl
		sb.WriteString(fmt.Sprintf("%s::_%s __%s_%s = nullptr;\n\n", m.Name, method.Name, m.Name, method.Name))
	}

	return sb.String(), nil
}
