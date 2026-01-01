package generator

import (
	"fmt"
	"strings"

	"github.com/untrustedmodders/plugify-gen/pkg/manifest"
)

// CxxGenerator generates C++ header files
type CxxGenerator struct {
	*BaseGenerator
}

// NewCxxGenerator creates a new C++ generator
func NewCxxGenerator() *CxxGenerator {
	return &CxxGenerator{
		BaseGenerator: NewBaseGenerator("cxx", NewCppCommonTypeMapper(), CppReservedWords),
	}
}

// Generate generates C++ bindings
func (g *CxxGenerator) Generate(m *manifest.Manifest, opts *GeneratorOptions) (*GeneratorResult, error) {
	g.ResetCaches()
	m.Sanitize(g.Sanitizer)
	opts = EnsureOptions(opts)

	// Collect all unique groups from both methods and classes
	groups := g.GetGroups(m)

	files := make(map[string]string)
	folder := fmt.Sprintf("external/plugify/modules/%s", m.Name)

	// Generate separate enums module
	enumsCode, err := g.generateEnumsFile(m)
	if err != nil {
		return nil, fmt.Errorf("generating enums file: %w", err)
	}
	files[fmt.Sprintf("%s/enums.ixx", folder)] = enumsCode

	// Generate separate aliases file
	aliasesCode, err := g.generateAliasesFile(m)
	if err != nil {
		return nil, fmt.Errorf("generating aliases file: %w", err)
	}
	files[fmt.Sprintf("%s/%s/aliases.hpp", folder, m.Name)] = aliasesCode

	// Generate separate delegates module
	delegatesCode, err := g.generateDelegatesFile(m)
	if err != nil {
		return nil, fmt.Errorf("generating delegates file: %w", err)
	}
	files[fmt.Sprintf("%s/delegates.ixx", folder)] = delegatesCode

	// Generate a module file for each group
	for groupName := range groups {
		groupCode, err := g.generateGroupFile(m, groupName, opts)
		if err != nil {
			return nil, fmt.Errorf("failed to generate group %s: %w", groupName, err)
		}
		files[fmt.Sprintf("%s/%s.ixx", folder, groupName)] = groupCode
	}

	// Generate main module interface that re-exports all pieces
	mainModule, err := g.generateMainHeader(m, groups)
	if err != nil {
		return nil, fmt.Errorf("generating main module: %w", err)
	}
	files[fmt.Sprintf("%s/package.ixx", folder)] = mainModule

	result := &GeneratorResult{
		Files: files,
	}

	return result, nil
}

func (g *CxxGenerator) generateEnums(m *manifest.Manifest) (string, error) {
	// Use the base generator's CollectEnums helper
	return g.CollectEnums(m, g.generateEnum)
}

func (g *CxxGenerator) generateEnum(enum *manifest.Enum, underlyingType string) (string, error) {
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

func (g *CxxGenerator) generateAliases(m *manifest.Manifest) (string, error) {
	// Use the base generator's CollectAliases helper
	return g.CollectAliases(m, g.generateAlias)
}

func (g *CxxGenerator) generateAlias(alias *manifest.Alias, underlyingType string) (string, error) {
	var sb strings.Builder

	if alias.Description != "" {
		sb.WriteString(fmt.Sprintf("  // %s\n", alias.Description))
	}

	sb.WriteString(fmt.Sprintf("  using %s = %s;\n", alias.Name, underlyingType))

	return sb.String(), nil
}

func (g *CxxGenerator) generateDelegates(m *manifest.Manifest) (string, error) {
	return g.CollectDelegates(m, g.generateDelegate)
}

func (g *CxxGenerator) generateDelegate(proto *manifest.Prototype) (string, error) {
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

// CxxDocOptions configures Doxygen-style documentation generation for C++
type CxxDocOptions struct {
	Description  string
	Params       []manifest.ParamType
	ParamAliases []*manifest.ParamAlias
	RetType      *manifest.RetType
	RetAlias     *manifest.RetAlias
	Indent       string // "  " for module-level, "    " for class members
}

// generateCxxDocumentation generates Doxygen-style comments for C++ methods
func (g *CxxGenerator) generateCxxDocumentation(opts CxxDocOptions) string {
	var sb strings.Builder

	sb.WriteString(opts.Indent + "/**\n")
	if opts.Description != "" {
		sb.WriteString(fmt.Sprintf("%s * @brief %s\n", opts.Indent, opts.Description))
	}

	// Parameters section
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
	if opts.RetType != nil && opts.RetType.Type != "void" {
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

	sb.WriteString(opts.Indent + " */\n")
	return sb.String()
}

func (g *CxxGenerator) generateMethod(pluginName string, method *manifest.Method) (string, error) {
	var sb strings.Builder

	// Generate function signature
	retType, err := g.typeMapper.MapReturnType(&method.RetType)
	if err != nil {
		return "", err
	}

	formattedParams, err := FormatParameters(method.ParamTypes, ParamFormatTypesAndNames, g.typeMapper)
	if err != nil {
		return "", err
	}

	// Generate type alias for function pointer
	funcTypeParams, err := FormatParameters(method.ParamTypes, ParamFormatTypes, g.typeMapper)
	if err != nil {
		return "", err
	}
	sb.WriteString(fmt.Sprintf("  using _%s = %s (*)(%s);\n", method.Name, retType, funcTypeParams))

	// Generate global exported function pointer
	sb.WriteString(fmt.Sprintf("  PLUGIFY_EXPORT _%s __%s_%s = nullptr;\n", method.Name, pluginName, method.Name))

	// Generate exported wrapper function
	paramNames, err := FormatParameters(method.ParamTypes, ParamFormatNames, g.typeMapper)
	if err != nil {
		return "", err
	}

	// Generate documentation comment
	sb.WriteString(g.generateCxxDocumentation(CxxDocOptions{
		Description: method.Description,
		Params:      method.ParamTypes,
		RetType:     &method.RetType,
		Indent:      "  ",
	}))

	// Add deprecation attribute if present
	if method.Deprecated != "" {
		sb.WriteString(fmt.Sprintf("  [[deprecated(\"%s\")]]\n", method.Deprecated))
	}

	sb.WriteString(fmt.Sprintf("\n  %s %s(%s) {\n", retType, method.Name, formattedParams))
	if method.RetType.Type == "void" {
		sb.WriteString(fmt.Sprintf("    return __%s_%s(%s);\n", pluginName, method.Name, paramNames))
	} else {
		sb.WriteString(fmt.Sprintf("    return __%s_%s(%s);\n", pluginName, method.Name, paramNames))
	}
	sb.WriteString("  }\n")

	return sb.String(), nil
}

func (g *CxxGenerator) generateClasses(m *manifest.Manifest) (string, error) {
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

func (g *CxxGenerator) generateClass(m *manifest.Manifest, class *manifest.Class) (string, error) {
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
	sb.WriteString("  /**\n")
	if class.Description != "" {
		sb.WriteString(fmt.Sprintf("   * @brief %s\n", class.Description))
	}
	sb.WriteString("   */\n")

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
			sb.WriteString(g.generatePrivateHelpers(class, invalidValue))
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

// generatePrivateHelpers generates private helper methods for classes with destructors
func (g *CxxGenerator) generatePrivateHelpers(class *manifest.Class, invalidValue string) string {
	return fmt.Sprintf(`    void destroy() const noexcept {
      if (_handle != %s && _ownership == %s::Owned) {
        %s(_handle);
      }
    }

    void nullify() noexcept {
      _handle = %s;
      _ownership = %s::Borrowed;
    }

`, invalidValue, OwnershipEnumName, *class.Destructor, invalidValue, OwnershipEnumName)
}

// generateUtilityMethods generates get, release, reset, and swap methods
func (g *CxxGenerator) generateUtilityMethods(className, invalidValue, handleType string, hasDtor bool) string {
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
func (g *CxxGenerator) generateComparisonOperators(className, invalidValue string) string {
	return fmt.Sprintf(`    explicit operator bool() const noexcept { return _handle != %s; }
    [[nodiscard]] auto operator<=>(const %s& other) const noexcept { return _handle <=> other._handle; }
    [[nodiscard]] bool operator==(const %s& other) const noexcept { return _handle == other._handle; }

`, invalidValue, className, className)
}

// generateCopyMoveSemantics generates copy/move constructors and assignment operators
func (g *CxxGenerator) generateCopyMoveSemantics(className string, hasDtor bool) string {
	if hasDtor {
		// Custom move semantics for classes with destructor
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

`, className, className, className, className, className, className, className, className, className)
	}

	// Default semantics for classes without destructor
	return fmt.Sprintf(`    %s(const %s&) = default;
    %s& operator=(const %s&) = default;
    %s(%s&&) noexcept = default;
    %s& operator=(%s&&) noexcept = default;
    ~%s() = default;

`, className, className, className, className, className, className, className, className, className)
}

func (g *CxxGenerator) generateConstructor(m *manifest.Manifest, class *manifest.Class, methodName string) (string, error) {
	// Find the method in the manifest
	method := FindMethod(m, methodName)
	if method == nil {
		return "", fmt.Errorf("constructor method %s not found", methodName)
	}

	var sb strings.Builder

	// Generate documentation
	sb.WriteString(g.generateCxxDocumentation(CxxDocOptions{
		Description: method.Description,
		Params:      method.ParamTypes,
		Indent:      "    ",
	}))

	// Add deprecation attribute if present
	if method.Deprecated != "" {
		sb.WriteString(fmt.Sprintf("    [[deprecated(\"%s\")]]\n", method.Deprecated))
	}

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

func (g *CxxGenerator) generateBinding(m *manifest.Manifest, class *manifest.Class, binding *manifest.Binding) (string, error) {
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
	sb.WriteString(g.generateCxxDocumentation(CxxDocOptions{
		Description:  method.Description,
		Params:       methodParams,
		ParamAliases: binding.ParamAliases,
		RetType:      &method.RetType,
		RetAlias:     binding.RetAlias,
		Indent:       "    ",
	}))

	// Add deprecation attribute if present (check both binding and underlying method)
	deprecationReason := binding.Deprecated
	if deprecationReason == "" {
		deprecationReason = method.Deprecated
	}
	if deprecationReason != "" {
		sb.WriteString(fmt.Sprintf("    [[deprecated(\"%s\")]]\n", deprecationReason))
	}

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
	callArgs := ""
	if binding.BindSelf {
		callArgs = "_handle"
	}

	for i, param := range methodParams {
		if callArgs != "" {
			callArgs += ", "
		}

		paramName := param.Name

		// Check if parameter has alias and needs .release() or .get()
		if i < len(binding.ParamAliases) && binding.ParamAliases[i] != nil {
			if binding.ParamAliases[i].Owner {
				callArgs += paramName + ".release()"
			} else {
				callArgs += paramName + ".get()"
			}
		} else {
			callArgs += paramName
		}
	}

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

func (g *CxxGenerator) applyParamAliases(formattedParams string, params []manifest.ParamType, aliases []*manifest.ParamAlias) string {
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
func (g *CxxGenerator) generateEnumsFile(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	// Module declaration
	sb.WriteString(fmt.Sprintf("// Generated from %s.pplugin\n\n", m.Name))
	sb.WriteString(fmt.Sprintf("export module %s.enums;\n\n", m.Name))

	// Global module fragment for standard library includes
	sb.WriteString("import <cstdint>;\n\n")

	// Namespace
	sb.WriteString(fmt.Sprintf("export namespace %s {\n\n", m.Name))

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

// generateEnumsFile generates a file containing all aliases
func (g *CxxGenerator) generateAliasesFile(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	// Module declaration
	sb.WriteString(fmt.Sprintf("// Generated from %s.pplugin\n\n", m.Name))
	sb.WriteString(fmt.Sprintf("export module %s.aliases;\n\n", m.Name))

	// Global module fragment for standard library includes
	sb.WriteString("import <cstdint>;\n\n")

	// Namespace
	sb.WriteString(fmt.Sprintf("export namespace %s {\n\n", m.Name))

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
func (g *CxxGenerator) generateDelegatesFile(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	// Module declaration
	sb.WriteString(fmt.Sprintf("// Generated from %s.pplugin\n\n", m.Name))
	sb.WriteString(fmt.Sprintf("export module %s.delegates;\n\n", m.Name))

	// Import enums module
	sb.WriteString(fmt.Sprintf("export import %s.enums;\n", m.Name))
	sb.WriteString(fmt.Sprintf("export import %s.aliases;\n\n", m.Name))

	sb.WriteString("import <cstdint>;\n")
	sb.WriteString("import <plg>;\n\n")

	// Namespace
	sb.WriteString(fmt.Sprintf("export namespace %s {\n\n", m.Name))

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
func (g *CxxGenerator) generateGroupFile(m *manifest.Manifest, groupName string, opts *GeneratorOptions) (string, error) {
	var sb strings.Builder

	// Module declaration
	sb.WriteString(fmt.Sprintf("// Generated from %s.pplugin (group: %s)\n\n", m.Name, groupName))
	sb.WriteString(fmt.Sprintf("export module %s.%s;\n", m.Name, groupName))

	// Import delegates module
	sb.WriteString(fmt.Sprintf("export import %s.enums;\n", m.Name))
	sb.WriteString(fmt.Sprintf("export import %s.aliases;\n", m.Name))
	sb.WriteString(fmt.Sprintf("export import %s.delegates;\n\n", m.Name))

	sb.WriteString("import <cstdint>;\n")
	sb.WriteString("import <plg>;\n\n")

	// Find which other groups this group depends on (for method calls from classes)
	if len(m.Classes) > 0 {
		dependentGroups := g.FindDependentGroups(m, groupName)
		if len(dependentGroups) > 0 {
			for depGroup := range dependentGroups {
				sb.WriteString(fmt.Sprintf("import %s.%s;\n", m.Name, depGroup))
			}
			sb.WriteString("\n")
		}
	}

	// Export
	sb.WriteString("#if defined(_WIN32)\n")
	sb.WriteString("#define PLUGIFY_EXPORT __declspec(dllexport)\n")
	sb.WriteString("#else\n")
	sb.WriteString("#define PLUGIFY_EXPORT __attribute__((visibility(\"default\")))\n")
	sb.WriteString("#endif\n\n")

	// Namespace
	sb.WriteString(fmt.Sprintf("export namespace %s {\n\n", m.Name))

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

// generateMainHeader generates the main module interface file that re-exports all submodules
func (g *CxxGenerator) generateMainHeader(m *manifest.Manifest, groups map[string]struct{}) (string, error) {
	var sb strings.Builder

	// Module declaration
	sb.WriteString(fmt.Sprintf("// Generated from %s.pplugin\n", m.Name))
	sb.WriteString("// This is the primary module interface that re-exports all components\n\n")
	sb.WriteString(fmt.Sprintf("export module %s;\n\n", m.Name))

	// Re-export enums and delegates
	sb.WriteString(fmt.Sprintf("export import %s.enums;\n", m.Name))
	sb.WriteString(fmt.Sprintf("export import %s.aliases;\n", m.Name))
	sb.WriteString(fmt.Sprintf("export import %s.delegates;\n", m.Name))

	// Re-export all group modules
	for groupName := range groups {
		sb.WriteString(fmt.Sprintf("export import %s.%s;\n", m.Name, groupName))
	}

	return sb.String(), nil
}
