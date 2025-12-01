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
	invalidNames := []string{
		"alignas", "alignof", "and", "and_eq", "asm", "auto", "bitand", "bitor",
		"bool", "break", "case", "catch", "char", "char8_t", "char16_t", "char32_t",
		"class", "compl", "concept", "const", "consteval", "constexpr", "constinit",
		"const_cast", "continue", "co_await", "co_return", "co_yield", "decltype",
		"default", "delete", "do", "double", "dynamic_cast", "else", "enum", "explicit",
		"export", "extern", "false", "float", "for", "friend", "goto", "if", "inline",
		"int", "long", "mutable", "namespace", "new", "noexcept", "not", "not_eq",
		"nullptr", "operator", "or", "or_eq", "private", "protected", "public",
		"register", "reinterpret_cast", "requires", "return", "short", "signed",
		"sizeof", "static", "static_assert", "static_cast", "struct", "switch",
		"template", "this", "thread_local", "throw", "true", "try", "typedef",
		"typeid", "typename", "union", "unsigned", "using", "virtual", "void",
		"volatile", "wchar_t", "while", "xor", "xor_eq",
	}

	return &CppGenerator{
		BaseGenerator: NewBaseGenerator("cpp", NewCppTypeMapper(), invalidNames),
	}
}

// Generate generates C++ bindings
func (g *CppGenerator) Generate(m *manifest.Manifest) (*GeneratorResult, error) {
	g.ResetCaches()

	// Collect all unique groups from both methods and classes
	groups := g.GetGroups(m)

	files := make(map[string]string)
	folder := fmt.Sprintf("external/plugify/include/%s", m.Name)

	// Generate separate enums file
	enumsCode, err := g.generateEnumsFile(m)
	if err != nil {
		return nil, err
	}
	files[fmt.Sprintf("%s/%s/enums.hpp", folder, m.Name)] = enumsCode

	// Generate separate delegates file
	delegatesCode, err := g.generateDelegatesFile(m)
	if err != nil {
		return nil, err
	}
	files[fmt.Sprintf("%s/%s/delegates.hpp", folder, m.Name)] = delegatesCode

	// Generate a file for each group
	for groupName := range groups {
		groupCode, err := g.generateGroupFile(m, groupName)
		if err != nil {
			return nil, fmt.Errorf("failed to generate group %s: %w", groupName, err)
		}
		files[fmt.Sprintf("%s/%s/%s.hpp", folder, m.Name, groupName)] = groupCode
	}

	// Generate main header that includes all pieces
	mainHeader, err := g.generateMainHeader(m, groups)
	if err != nil {
		return nil, err
	}
	files[fmt.Sprintf("%s/%s.hpp", folder, m.Name)] = mainHeader

	result := &GeneratorResult{
		Files: files,
	}

	return result, nil
}

func (g *CppGenerator) generateEnums(m *manifest.Manifest) (string, error) {
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

func (g *CppGenerator) processPrototypeEnums(proto *manifest.Prototype, sb *strings.Builder) {
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

func (g *CppGenerator) generateEnum(enum *manifest.EnumType, underlyingType string) string {
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
	return sb.String()
}

func (g *CppGenerator) generateDelegates(m *manifest.Manifest) (string, error) {
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
	params, err := FormatParameters(proto.ParamTypes, ParamFormatTypes, g.typeMapper, g.SanitizeName)
	if err != nil {
		return "", err
	}

	sb.WriteString(fmt.Sprintf("  using %s = %s (*)(%s);\n\n", proto.Name, retType, params))
	return sb.String(), nil
}

func (g *CppGenerator) generateMethod(pluginName string, method *manifest.Method) (string, error) {
	var sb strings.Builder

	// Generate documentation comment
	sb.WriteString("  /**\n")
	if method.Description != "" {
		sb.WriteString(fmt.Sprintf("   * @brief %s\n", method.Description))
	}
	for _, param := range method.ParamTypes {
		paramType := param.Type
		if param.Ref {
			paramType += "&"
		}
		sb.WriteString(fmt.Sprintf("   * @param %s (%s)", g.SanitizeName(param.Name), paramType))
		if param.Description != "" {
			sb.WriteString(fmt.Sprintf(": %s", param.Description))
		}
		sb.WriteString("\n")
	}
	if method.RetType.Type != "void" {
		sb.WriteString(fmt.Sprintf("   * @return %s", method.RetType.Type))
		if method.RetType.Description != "" {
			sb.WriteString(fmt.Sprintf(": %s", method.RetType.Description))
		}
		sb.WriteString("\n")
	}
	sb.WriteString("   */\n")

	// Generate function signature
	retType, err := g.typeMapper.MapReturnType(&method.RetType)
	if err != nil {
		return "", err
	}

	params, err := FormatParameters(method.ParamTypes, ParamFormatTypesAndNames, g.typeMapper, g.SanitizeName)
	if err != nil {
		return "", err
	}

	sb.WriteString(fmt.Sprintf("  inline %s %s(%s) {\n", retType, g.SanitizeName(method.Name), params))

	// Generate function body
	// Type alias for function pointer
	funcTypeParams, err := FormatParameters(method.ParamTypes, ParamFormatTypes, g.typeMapper, g.SanitizeName)
	if err != nil {
		return "", err
	}
	sb.WriteString(fmt.Sprintf("    using %sFn = %s (*)(%s);\n", method.Name, retType, funcTypeParams))
	sb.WriteString(fmt.Sprintf("    static %sFn __func = nullptr;\n", method.Name))
	sb.WriteString(fmt.Sprintf("    if (__func == nullptr) plg::GetMethodPtr2(\"%s.%s\", reinterpret_cast<void**>(&__func));\n", pluginName, method.Name))

	// Call function
	paramNames, err := FormatParameters(method.ParamTypes, ParamFormatNames, g.typeMapper, g.SanitizeName)
	if err != nil {
		return "", err
	}

	if method.RetType.Type == "void" {
		sb.WriteString(fmt.Sprintf("    __func(%s);\n", paramNames))
	} else {
		sb.WriteString(fmt.Sprintf("    return __func(%s);\n", paramNames))
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
	handleType, err := g.typeMapper.MapType(class.HandleType, TypeContextReturn, false)
	if err != nil {
		return "", err
	}

	// Determine invalid value representation
	invalidValue := class.InvalidValue
	if class.HandleType == "ptr64" || strings.Contains(handleType, "*") {
		if invalidValue == "0" || invalidValue == "" {
			invalidValue = "nullptr"
		}
	}

	hasCtor := len(class.Constructors) > 0
	hasDtor := class.Destructor != ""

	// Class documentation
	sb.WriteString("  /**\n")
	if class.Description != "" {
		sb.WriteString(fmt.Sprintf("   * @brief %s\n", class.Description))
	}
	sb.WriteString("   */\n")

	// Class declaration
	sb.WriteString(fmt.Sprintf("  class %s final {\n", class.Name))

	sb.WriteString("  public:\n")

	// Default constructor
	sb.WriteString(fmt.Sprintf("    %s() = default;\n\n", class.Name))

	// Generate constructors from methods
	for _, ctorName := range class.Constructors {
		ctorCode, err := g.generateConstructor(m, class, ctorName)
		if err != nil {
			return "", err
		}
		sb.WriteString(ctorCode)
	}

	// Destructor and copy/move semantics
	if hasDtor {
		sb.WriteString(fmt.Sprintf("    ~%s() {\n", class.Name))
		sb.WriteString("      destroy();\n")
		sb.WriteString("    }\n\n")

		sb.WriteString(fmt.Sprintf("    %s(const %s&) = delete;\n", class.Name, class.Name))
		sb.WriteString(fmt.Sprintf("    %s& operator=(const %s&) = delete;\n\n", class.Name, class.Name))

		sb.WriteString(fmt.Sprintf("    %s(%s&& other) noexcept\n", class.Name, class.Name))
		sb.WriteString("      : _handle(other._handle)\n")
		sb.WriteString("      , _ownership(other._ownership) {\n")
		sb.WriteString("      other.nullify();\n")
		sb.WriteString("    }\n\n")

		sb.WriteString(fmt.Sprintf("    %s& operator=(%s&& other) noexcept {\n", class.Name, class.Name))
		sb.WriteString("      if (this != &other) {\n")
		sb.WriteString("        destroy();\n")
		sb.WriteString("        _handle = other._handle;\n")
		sb.WriteString("        _ownership = other._ownership;\n")
		sb.WriteString("        other.nullify();\n")
		sb.WriteString("      }\n")
		sb.WriteString("      return *this;\n")
		sb.WriteString("    }\n\n")
	} else {
		sb.WriteString(fmt.Sprintf("    %s(const %s&) = default;\n", class.Name, class.Name))
		sb.WriteString(fmt.Sprintf("    %s& operator=(const %s&) = default;\n", class.Name, class.Name))
		sb.WriteString(fmt.Sprintf("    %s(%s&&) noexcept = default;\n", class.Name, class.Name))
		sb.WriteString(fmt.Sprintf("    %s& operator=(%s&&) noexcept = default;\n", class.Name, class.Name))
		sb.WriteString(fmt.Sprintf("    ~%s() = default;\n\n", class.Name))
	}

	// Constructor from handle
	if hasDtor {
		sb.WriteString(fmt.Sprintf("    %s(%s handle, Ownership ownership) : _handle(handle), _ownership(ownership) {}\n\n", class.Name, handleType))
	} else {
		ctorTag := ""
		if hasCtor {
			ctorTag = ", Ownership"
		}
		sb.WriteString(fmt.Sprintf("    explicit %s(%s handle%s) : _handle(handle) {}\n\n", class.Name, handleType, ctorTag))
	}

	// Utility methods
	sb.WriteString(fmt.Sprintf("    [[nodiscard]] auto get() const noexcept { return _handle; }\n\n"))

	sb.WriteString("    [[nodiscard]] auto release() noexcept {\n")
	sb.WriteString("      auto handle = _handle;\n")
	if hasDtor {
		sb.WriteString("      nullify();\n")
	} else {
		sb.WriteString(fmt.Sprintf("      _handle = %s;\n", invalidValue))
	}
	sb.WriteString("      return handle;\n")
	sb.WriteString("    }\n\n")

	sb.WriteString("    void reset() noexcept {\n")
	if hasDtor {
		sb.WriteString("      destroy();\n")
		sb.WriteString("      nullify();\n")
	} else {
		sb.WriteString(fmt.Sprintf("      _handle = %s;\n", invalidValue))
	}
	sb.WriteString("    }\n\n")

	sb.WriteString(fmt.Sprintf("    void swap(%s& other) noexcept {\n", class.Name))
	sb.WriteString("      using std::swap;\n")
	sb.WriteString("      swap(_handle, other._handle);\n")
	if hasDtor {
		sb.WriteString("      swap(_ownership, other._ownership);\n")
	}
	sb.WriteString("    }\n\n")

	sb.WriteString(fmt.Sprintf("    friend void swap(%s& lhs, %s& rhs) noexcept { lhs.swap(rhs); }\n\n", class.Name, class.Name))

	// Operators
	sb.WriteString(fmt.Sprintf("    explicit operator bool() const noexcept { return _handle != %s; }\n", invalidValue))
	sb.WriteString(fmt.Sprintf("    [[nodiscard]] auto operator<=>(const %s& other) const noexcept { return _handle <=> other._handle; }\n", class.Name))
	sb.WriteString(fmt.Sprintf("    [[nodiscard]] bool operator==(const %s& other) const noexcept { return _handle == other._handle; }\n\n", class.Name))

	// Generate class bindings
	for _, binding := range class.Bindings {
		methodCode, err := g.generateBinding(m, class, &binding)
		if err != nil {
			return "", err
		}
		sb.WriteString(methodCode)
		sb.WriteString("\n")
	}

	// Private section
	sb.WriteString("  private:\n")

	if hasDtor {
		// Destroy helper
		sb.WriteString("    void destroy() const noexcept {\n")
		sb.WriteString(fmt.Sprintf("      if (_handle != %s && _ownership == Ownership::Owned) {\n", invalidValue))
		sb.WriteString(fmt.Sprintf("        %s(_handle);\n", class.Destructor))
		sb.WriteString("      }\n")
		sb.WriteString("    }\n\n")

		// Nullify helper
		sb.WriteString("    void nullify() noexcept {\n")
		sb.WriteString(fmt.Sprintf("      _handle = %s;\n", invalidValue))
		sb.WriteString("      _ownership = Ownership::Borrowed;\n")
		sb.WriteString("    }\n\n")
	}

	// Member variables
	sb.WriteString(fmt.Sprintf("    %s _handle{%s};\n", handleType, invalidValue))
	if hasDtor {
		sb.WriteString("    Ownership _ownership{Ownership::Borrowed};\n")
	}

	sb.WriteString("  };\n\n")

	return sb.String(), nil
}

func (g *CppGenerator) generateConstructor(m *manifest.Manifest, class *manifest.Class, methodName string) (string, error) {
	// Find the method in the manifest
	var method *manifest.Method
	for i := range m.Methods {
		if m.Methods[i].Name == methodName || m.Methods[i].FuncName == methodName {
			method = &m.Methods[i]
			break
		}
	}
	if method == nil {
		return "", fmt.Errorf("constructor method %s not found", methodName)
	}

	var sb strings.Builder

	// Generate documentation
	sb.WriteString("    /**\n")
	if method.Description != "" {
		sb.WriteString(fmt.Sprintf("     * @brief %s\n", method.Description))
	}

	// Document parameters (skip first param if it's the return handle in C API)
	for _, param := range method.ParamTypes {
		paramType := param.Type
		if param.Ref {
			paramType += "&"
		}
		sb.WriteString(fmt.Sprintf("     * @param %s (%s)", g.SanitizeName(param.Name), paramType))
		if param.Description != "" {
			sb.WriteString(fmt.Sprintf(": %s", param.Description))
		}
		sb.WriteString("\n")
	}
	sb.WriteString("     */\n")

	// Generate constructor signature
	params, err := FormatParameters(method.ParamTypes, ParamFormatTypesAndNames, g.typeMapper, g.SanitizeName)
	if err != nil {
		return "", err
	}

	sb.WriteString(fmt.Sprintf("    explicit %s(%s)\n", class.Name, params))

	// Generate initialization list
	paramNames, err := FormatParameters(method.ParamTypes, ParamFormatNames, g.typeMapper, g.SanitizeName)
	if err != nil {
		return "", err
	}

	sb.WriteString(fmt.Sprintf("      : %s(%s(%s), Ownership::Owned) {}\n\n", class.Name, method.FuncName, paramNames))

	return sb.String(), nil
}

func (g *CppGenerator) generateBinding(m *manifest.Manifest, class *manifest.Class, binding *manifest.Binding) (string, error) {
	// Find the underlying method
	var method *manifest.Method
	for i := range m.Methods {
		if m.Methods[i].Name == binding.Method || m.Methods[i].FuncName == binding.Method {
			method = &m.Methods[i]
			break
		}
	}
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
	sb.WriteString("    /**\n")
	if method.Description != "" {
		sb.WriteString(fmt.Sprintf("     * @brief %s\n", method.Description))
	}

	// Document parameters (excluding self if bindSelf)
	for i, param := range methodParams {
		paramType := param.Type
		if param.Ref {
			paramType += "&"
		}

		// Check if this parameter has an alias
		var aliasName string
		if i < len(binding.ParamAliases) && binding.ParamAliases[i] != nil {
			aliasName = binding.ParamAliases[i].Name
		}

		if aliasName != "" {
			sb.WriteString(fmt.Sprintf("     * @param %s (%s)", g.SanitizeName(param.Name), aliasName))
		} else {
			sb.WriteString(fmt.Sprintf("     * @param %s (%s)", g.SanitizeName(param.Name), paramType))
		}

		if param.Description != "" {
			sb.WriteString(fmt.Sprintf(": %s", param.Description))
		}
		sb.WriteString("\n")
	}

	// Document return type
	if method.RetType.Type != "void" {
		returnType := method.RetType.Type
		if binding.RetAlias != nil && binding.RetAlias.Name != "" {
			returnType = binding.RetAlias.Name
		}
		sb.WriteString(fmt.Sprintf("     * @return %s", returnType))
		if method.RetType.Description != "" {
			sb.WriteString(fmt.Sprintf(": %s", method.RetType.Description))
		}
		sb.WriteString("\n")
	}
	sb.WriteString("     */\n")

	// Determine if method is static
	isStatic := !binding.BindSelf

	// Generate method signature
	retType, err := g.typeMapper.MapReturnType(&method.RetType)
	if err != nil {
		return "", err
	}

	// Handle return type alias
	if binding.RetAlias != nil && binding.RetAlias.Name != "" {
		retType = binding.RetAlias.Name
	}

	formattedParams, err := FormatParameters(methodParams, ParamFormatTypesAndNames, g.typeMapper, g.SanitizeName)
	if err != nil {
		return "", err
	}

	// Apply parameter aliases to signature
	if len(binding.ParamAliases) > 0 {
		formattedParams = g.applyParamAliases(formattedParams, methodParams, binding.ParamAliases)
	}

	staticKeyword := ""
	if isStatic {
		staticKeyword = "static "
	}

	sb.WriteString(fmt.Sprintf("    %s%s %s(%s) {\n", staticKeyword, retType, binding.Name, formattedParams))

	// Generate null check if needed
	nullPolicy := class.NullPolicy
	if nullPolicy == "" {
		nullPolicy = "throw"
	}

	if !isStatic && nullPolicy == "throw" {
		invalidValue := class.InvalidValue
		if class.HandleType == "ptr64" {
			invalidValue = "nullptr"
		}
		sb.WriteString(fmt.Sprintf("      if (_handle == %s) throw std::runtime_error(\"%s: Empty handle\");\n", invalidValue, class.Name))
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

		paramName := g.SanitizeName(param.Name)

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

	// Generate return statement
	if method.RetType.Type == "void" {
		sb.WriteString(fmt.Sprintf("      %s(%s);\n", method.FuncName, callArgs))
	} else {
		// Handle return alias
		if binding.RetAlias != nil && binding.RetAlias.Name != "" {
			ownership := "Ownership::Borrowed"
			if binding.RetAlias.Owner {
				ownership = "Ownership::Owned"
			}
			sb.WriteString(fmt.Sprintf("      return %s(%s(%s), %s);\n", binding.RetAlias.Name, method.FuncName, callArgs, ownership)) // always pass ownership just as a tag
		} else {
			sb.WriteString(fmt.Sprintf("      return %s(%s);\n", method.FuncName, callArgs))
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
			paramType, _ := g.typeMapper.MapParamType(&param, TypeContextValue)

			// Determine the replacement type
			replacementType := aliases[i].Name
			if aliases[i].Owner {
				replacementType = aliases[i].Name + "&&"
			} else {
				replacementType = "const " + aliases[i].Name + "&"
			}

			// Replace in the formatted params
			oldPattern := fmt.Sprintf("%s %s", paramType, g.SanitizeName(param.Name))
			newPattern := fmt.Sprintf("%s %s", replacementType, g.SanitizeName(param.Name))
			result = strings.ReplaceAll(result, oldPattern, newPattern)
		}
	}

	return result
}

func (g *CppGenerator) needsOwnershipEnum(m *manifest.Manifest) bool {
	for _, class := range m.Classes {
		if class.Destructor != "" {
			return true
		}
	}
	return false
}

// generateEnumsFile generates a file containing all enums
func (g *CppGenerator) generateEnumsFile(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	// Header guard and includes
	sb.WriteString("#pragma once\n\n")
	sb.WriteString("#include <cstdint>\n\n")
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
	if g.needsOwnershipEnum(m) {
		sb.WriteString("  /// Ownership type for RAII wrappers\n")
		sb.WriteString("  enum class Ownership : bool { Borrowed, Owned };\n\n")
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
func (g *CppGenerator) generateGroupFile(m *manifest.Manifest, groupName string) (string, error) {
	var sb strings.Builder

	// Header guard and includes
	sb.WriteString("#pragma once\n\n")
	sb.WriteString(fmt.Sprintf("#include \"%s.hpp\"\n", m.Name))

	sb.WriteString(fmt.Sprintf("// Generated from %s.pplugin (group: %s)\n\n", m.Name, groupName))

	// Namespace
	sb.WriteString(fmt.Sprintf("namespace %s {\n\n", m.Name))

	// Generate methods for this group
	for _, method := range m.Methods {
		methodGroup := g.SanitizeNameLower(method.Group)
		if methodGroup == "" {
			methodGroup = "main"
		}
		if methodGroup == groupName {
			methodCode, err := g.generateMethod(m.Name, &method)
			if err != nil {
				return "", fmt.Errorf("failed to generate method %s: %w", method.Name, err)
			}
			sb.WriteString(methodCode)
			sb.WriteString("\n")
		}
	}

	// Generate classes for this group
	for _, class := range m.Classes {
		classGroup := g.SanitizeNameLower(class.Group)
		if classGroup == "" {
			classGroup = "main"
		}
		if classGroup == groupName {
			classCode, err := g.generateClass(m, &class)
			if err != nil {
				return "", fmt.Errorf("failed to generate class %s: %w", class.Name, err)
			}
			sb.WriteString(classCode)
			sb.WriteString("\n")
		}
	}

	// Close namespace
	sb.WriteString(fmt.Sprintf("} // namespace %s\n", m.Name))

	return sb.String(), nil
}

// generateMainHeader generates the main header file that includes all group files
func (g *CppGenerator) generateMainHeader(m *manifest.Manifest, groups map[string]bool) (string, error) {
	var sb strings.Builder

	// Header guard
	sb.WriteString("#pragma once\n\n")
	sb.WriteString(fmt.Sprintf("// Generated from %s.pplugin\n", m.Name))
	sb.WriteString("// This header includes all generated components\n\n")

	// Include enums and delegates first
	sb.WriteString(fmt.Sprintf("#include \"%s/enums.hpp\"\n", m.Name))
	sb.WriteString(fmt.Sprintf("#include \"%s/delegates.hpp\"\n", m.Name))

	// Include all group files
	// Convert map to sorted slice for deterministic output
	var groupNames []string
	for groupName := range groups {
		groupNames = append(groupNames, groupName)
	}
	// Sort the group names for consistent order
	for i := 0; i < len(groupNames); i++ {
		for j := i + 1; j < len(groupNames); j++ {
			if groupNames[i] > groupNames[j] {
				groupNames[i], groupNames[j] = groupNames[j], groupNames[i]
			}
		}
	}

	for _, groupName := range groupNames {
		sb.WriteString(fmt.Sprintf("#include \"%s/%s.hpp\"\n", m.Name, groupName))
	}

	return sb.String(), nil
}

// CppTypeMapper implements type mapping for C++
type CppTypeMapper struct{}

func NewCppTypeMapper() *CppTypeMapper {
	return &CppTypeMapper{}
}

func (m *CppTypeMapper) MapType(baseType string, context TypeContext, isArray bool) (string, error) {
	// Base type mapping
	typeMap := map[string]string{
		"void":   "void",
		"bool":   "bool",
		"char8":  "char",
		"char16": "char16_t",
		"int8":   "int8_t",
		"int16":  "int16_t",
		"int32":  "int32_t",
		"int64":  "int64_t",
		"uint8":  "uint8_t",
		"uint16": "uint16_t",
		"uint32": "uint32_t",
		"uint64": "uint64_t",
		"ptr64":  "void*",
		"float":  "float",
		"double": "double",
		"string": "plg::string",
		"any":    "plg::any",
		"vec2":   "plg::vec2",
		"vec3":   "plg::vec3",
		"vec4":   "plg::vec4",
		"mat4x4": "plg::mat4x4",
	}

	mapped, ok := typeMap[baseType]
	if !ok {
		// Assume it's a custom type (enum or delegate)
		mapped = baseType
	}

	// Handle arrays
	if isArray {
		mapped = fmt.Sprintf("plg::vector<%s>", mapped)
	}

	// Handle parameter context (value parameters)
	// Object-like types pass by const& even when not ref=true
	if context == TypeContextValue && baseType != "void" {
		if m.isObjectLikeType(baseType) || isArray {
			mapped = fmt.Sprintf("const %s&", mapped)
		}
	}

	// Handle reference context (ref=true parameters)
	if context == TypeContextRef && baseType != "void" {
		mapped = mapped + "&"
	}

	return mapped, nil
}

// isObjectLikeType returns true for types that should be passed by const& in parameters
func (m *CppTypeMapper) isObjectLikeType(baseType string) bool {
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

func (m *CppTypeMapper) MapParamType(param *manifest.ParamType, context TypeContext) (string, error) {
	// Check for enum type first
	if param.Enum != nil {
		typeName := param.Enum.Name
		if param.IsArray() {
			typeName = fmt.Sprintf("plg::vector<%s>", typeName)
		}
		// Enums are primitive types, pass by value or reference
		if param.Ref && !param.IsArray() {
			return typeName + "&", nil
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

func (m *CppTypeMapper) MapReturnType(retType *manifest.TypeInfo) (string, error) {
	// Check for enum type
	if retType.Enum != nil {
		typeName := retType.Enum.Name
		if retType.IsArray() {
			typeName = fmt.Sprintf("plg::vector<%s>", typeName)
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
