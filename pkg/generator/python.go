package generator

import (
	"fmt"
	"strings"

	"github.com/untrustedmodders/plugify-gen/pkg/manifest"
)

// PythonGenerator generates Python stub (.pyi) files
type PythonGenerator struct {
	*BaseGenerator
}

// NewPythonGenerator creates a new Python generator
func NewPythonGenerator() *PythonGenerator {
	invalidNames := []string{
		"False", "await", "else", "import", "pass",
		"None", "break", "except", "in", "raise",
		"True", "class", "finally", "is", "return",
		"and", "continue", "for", "lambda", "try",
		"as", "def", "from", "nonlocal", "while",
		"assert", "del", "global", "not", "with",
		"async", "elif", "if", "or", "yield",
	}

	return &PythonGenerator{
		BaseGenerator: NewBaseGenerator("python", NewPythonTypeMapper(), invalidNames),
	}
}

// Generate generates Python bindings
func (g *PythonGenerator) Generate(m *manifest.Manifest) (*GeneratorResult, error) {
	g.ResetCaches()

	var sb strings.Builder

	// Check if we need Callable import
	needsCallable := g.needsCallable(m)

	// Header imports
	if needsCallable {
		sb.WriteString("from collections.abc import Callable\n")
	}
	sb.WriteString("from enum import IntEnum\n")
	sb.WriteString("from plugify.plugin import Vector2, Vector3, Vector4, Matrix4x4\n\n")
	sb.WriteString(fmt.Sprintf("# Generated from %s.pplugin\n\n", m.Name))

	// Generate enums
	enumsCode, err := g.generateEnums(m)
	if err != nil {
		return nil, err
	}
	if enumsCode != "" {
		sb.WriteString(enumsCode)
		sb.WriteString("\n")
	}

	// Generate methods
	for _, method := range m.Methods {
		methodCode, err := g.generateMethod(&method)
		if err != nil {
			return nil, fmt.Errorf("failed to generate method %s: %w", method.Name, err)
		}
		sb.WriteString(methodCode)
		sb.WriteString("\n")
	}

	// Generate classes
	if len(m.Classes) > 0 {
		classesCode, err := g.generateClasses(m)
		if err != nil {
			return nil, err
		}
		sb.WriteString(classesCode)
		sb.WriteString("\n")
	}

	result := &GeneratorResult{
		Files: map[string]string{
			fmt.Sprintf("pps/%s.pyi", m.Name): sb.String(),
		},
	}

	return result, nil
}

// needsCallable checks if any method uses function/delegate types
func (g *PythonGenerator) needsCallable(m *manifest.Manifest) bool {
	for _, method := range m.Methods {
		if method.RetType.Prototype != nil {
			return true
		}
		for _, param := range method.ParamTypes {
			if param.Prototype != nil {
				return true
			}
		}
	}
	return false
}

func (g *PythonGenerator) generateEnums(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

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

func (g *PythonGenerator) processPrototypeEnums(proto *manifest.Prototype, sb *strings.Builder) {
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

func (g *PythonGenerator) generateEnum(enum *manifest.EnumType) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("class %s(IntEnum):\n", enum.Name))

	if enum.Description != "" {
		sb.WriteString(fmt.Sprintf("    \"\"\"\n    %s\n    \"\"\"\n", enum.Description))
	}

	for _, val := range enum.Values {
		if val.Description != "" {
			sb.WriteString(fmt.Sprintf("    # %s\n", val.Description))
		}
		sb.WriteString(fmt.Sprintf("    %s = %d\n", val.Name, val.Value))
	}

	return sb.String()
}

func (g *PythonGenerator) generateClasses(m *manifest.Manifest) (string, error) {
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

func (g *PythonGenerator) generateClass(m *manifest.Manifest, class *manifest.Class) (string, error) {
	var sb strings.Builder

	hasDtor := class.Destructor != ""

	// Class declaration with docstring
	sb.WriteString(fmt.Sprintf("class %s:\n", class.Name))
	if class.Description != "" {
		sb.WriteString(fmt.Sprintf("    \"\"\"\n    %s\n    \"\"\"\n", class.Description))
	}

	// Generate constructors
	if len(class.Constructors) > 0 {
		for _, ctorName := range class.Constructors {
			ctorCode, err := g.generateConstructor(m, class, ctorName)
			if err != nil {
				return "", err
			}
			sb.WriteString(ctorCode)
			sb.WriteString("\n")
		}
	} else {
		// Default constructor if no constructors specified
		sb.WriteString("    def __init__(self) -> None:\n")
		sb.WriteString("        ...\n\n")
	}

	// Generate destructor and context manager methods for classes with destructors
	if hasDtor {
		sb.WriteString(g.generateDestructorMethods(class))
	}

	// Generate utility methods (get, release, reset, valid) for all classes
	sb.WriteString(g.generateUtilityMethods(class))

	// Generate bindings (methods)
	for _, binding := range class.Bindings {
		methodCode, err := g.generateBinding(m, class, &binding)
		if err != nil {
			return "", err
		}
		sb.WriteString(methodCode)
		sb.WriteString("\n")
	}

	return sb.String(), nil
}

func (g *PythonGenerator) generateDestructorMethods(class *manifest.Class) string {
	var sb strings.Builder

	// __del__ method
	sb.WriteString("    def __del__(self) -> None:\n")
	sb.WriteString("        \"\"\"\n")
	sb.WriteString("        Destructor. Releases the underlying handle if owned.\n")
	sb.WriteString("        \"\"\"\n")
	sb.WriteString("        ...\n\n")

	// close method
	sb.WriteString("    def close(self) -> None:\n")
	sb.WriteString("        \"\"\"\n")
	sb.WriteString("        Close/destroy the handle if owned.\n")
	sb.WriteString("        \"\"\"\n")
	sb.WriteString("        ...\n\n")

	// __enter__ method for context manager
	sb.WriteString(fmt.Sprintf("    def __enter__(self) -> \"%s\":\n", class.Name))
	sb.WriteString("        \"\"\"\n")
	sb.WriteString("        Enter the runtime context for this object.\n\n")
	sb.WriteString("        Returns:\n")
	sb.WriteString(fmt.Sprintf("            %s: self\n", class.Name))
	sb.WriteString("        \"\"\"\n")
	sb.WriteString("        ...\n\n")

	// __exit__ method for context manager
	sb.WriteString("    def __exit__(self, exc_type: type[BaseException] | None, exc_val: BaseException | None, exc_tb: object) -> None:\n")
	sb.WriteString("        \"\"\"\n")
	sb.WriteString("        Exit the runtime context and release resources.\n\n")
	sb.WriteString("        Args:\n")
	sb.WriteString("            exc_type (type[BaseException] | None): Exception type if an exception was raised\n")
	sb.WriteString("            exc_val (BaseException | None): Exception value if an exception was raised\n")
	sb.WriteString("            exc_tb (object): Traceback if an exception was raised\n")
	sb.WriteString("        \"\"\"\n")
	sb.WriteString("        ...\n\n")

	return sb.String()
}

func (g *PythonGenerator) generateUtilityMethods(class *manifest.Class) string {
	var sb strings.Builder

	// Get the handle type mapped to Python
	handleType, _ := g.typeMapper.MapType(class.HandleType, TypeContextReturn, false)

	// get() method
	sb.WriteString(fmt.Sprintf("    def get(self) -> %s:\n", handleType))
	sb.WriteString("        \"\"\"\n")
	sb.WriteString("        Get the raw handle value without transferring ownership.\n\n")
	sb.WriteString("        Returns:\n")
	sb.WriteString(fmt.Sprintf("            %s: The underlying handle value\n", handleType))
	sb.WriteString("        \"\"\"\n")
	sb.WriteString("        ...\n\n")

	// release() method
	sb.WriteString(fmt.Sprintf("    def release(self) -> %s:\n", handleType))
	sb.WriteString("        \"\"\"\n")
	sb.WriteString("        Release ownership of the handle and return it.\n\n")
	sb.WriteString("        Returns:\n")
	sb.WriteString(fmt.Sprintf("            %s: The released handle value\n", handleType))
	sb.WriteString("        \"\"\"\n")
	sb.WriteString("        ...\n\n")

	// reset() method
	sb.WriteString("    def reset(self) -> None:\n")
	sb.WriteString("        \"\"\"\n")
	sb.WriteString("        Reset the handle by closing it.\n")
	sb.WriteString("        \"\"\"\n")
	sb.WriteString("        ...\n\n")

	// valid() method
	sb.WriteString("    def valid(self) -> bool:\n")
	sb.WriteString("        \"\"\"\n")
	sb.WriteString("        Check if the handle is valid.\n\n")
	sb.WriteString("        Returns:\n")
	sb.WriteString("            bool: True if the handle is valid, False otherwise\n")
	sb.WriteString("        \"\"\"\n")
	sb.WriteString("        ...\n\n")

	return sb.String()
}

func (g *PythonGenerator) generateConstructor(m *manifest.Manifest, class *manifest.Class, methodName string) (string, error) {
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

	// Generate __init__ signature
	params, err := g.formatParameters(method.ParamTypes)
	if err != nil {
		return "", err
	}

	if params != "" {
		sb.WriteString(fmt.Sprintf("    def __init__(self, %s) -> None:\n", params))
	} else {
		sb.WriteString("    def __init__(self) -> None:\n")
	}

	// Generate docstring
	if method.Description != "" {
		sb.WriteString("        \"\"\"\n")
		sb.WriteString(fmt.Sprintf("        %s\n", method.Description))
		if len(method.ParamTypes) > 0 {
			sb.WriteString("\n        Args:\n")
			for _, param := range method.ParamTypes {
				paramName := param.Name
				if paramName == "" {
					paramName = "unnamed"
				}
				sb.WriteString(fmt.Sprintf("            %s (%s): %s\n",
					g.SanitizeName(paramName), param.Type, param.Description))
			}
		}
		sb.WriteString("        \"\"\"\n")
	}

	// Stub body
	sb.WriteString("        ...\n")

	return sb.String(), nil
}

func (g *PythonGenerator) generateBinding(m *manifest.Manifest, class *manifest.Class, binding *manifest.Binding) (string, error) {
	// Find the method in the manifest
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

	// Format parameters
	formattedParams, err := g.formatParameters(methodParams)
	if err != nil {
		return "", err
	}

	// Apply parameter aliases (replace types with class names)
	if len(binding.ParamAliases) > 0 {
		formattedParams = g.applyParamAliases(formattedParams, methodParams, binding.ParamAliases)
	}

	// Generate return type
	retType, err := g.typeMapper.MapReturnType(&method.RetType)
	if err != nil {
		return "", err
	}

	// Handle return alias
	if binding.RetAlias != nil && binding.RetAlias.Name != "" {
		retType = binding.RetAlias.Name
	}

	// Determine if method is static
	isStatic := !binding.BindSelf
	decorator := ""
	if isStatic {
		decorator = "    @staticmethod\n"
	}

	// Generate method signature
	if decorator != "" {
		sb.WriteString(decorator)
	}

	if isStatic {
		if formattedParams != "" {
			sb.WriteString(fmt.Sprintf("    def %s(%s) -> %s:\n", binding.Name, formattedParams, retType))
		} else {
			sb.WriteString(fmt.Sprintf("    def %s() -> %s:\n", binding.Name, retType))
		}
	} else {
		if formattedParams != "" {
			sb.WriteString(fmt.Sprintf("    def %s(self, %s) -> %s:\n", binding.Name, formattedParams, retType))
		} else {
			sb.WriteString(fmt.Sprintf("    def %s(self) -> %s:\n", binding.Name, retType))
		}
	}

	// Generate docstring
	if method.Description != "" {
		sb.WriteString("        \"\"\"\n")
		sb.WriteString(fmt.Sprintf("        %s\n", method.Description))

		if len(methodParams) > 0 {
			sb.WriteString("\n        Args:\n")
			for i, param := range methodParams {
				paramName := param.Name
				if paramName == "" {
					paramName = "unnamed"
				}

				// Check if this parameter has an alias
				paramType := param.Type
				if i < len(binding.ParamAliases) && binding.ParamAliases[i].Name != "" {
					paramType = binding.ParamAliases[i].Name
				}

				sb.WriteString(fmt.Sprintf("            %s (%s): %s\n",
					g.SanitizeName(paramName), paramType, param.Description))
			}
		}

		if method.RetType.Type != "void" {
			returnType := method.RetType.Type
			if binding.RetAlias != nil && binding.RetAlias.Name != "" {
				returnType = binding.RetAlias.Name
			}
			sb.WriteString(fmt.Sprintf("\n        Returns:\n            %s: %s\n",
				returnType, method.RetType.Description))
		}

		sb.WriteString("        \"\"\"\n")
	}

	// Stub body
	sb.WriteString("        ...\n")

	return sb.String(), nil
}

func (g *PythonGenerator) applyParamAliases(formattedParams string, params []manifest.ParamType, aliases []manifest.ParamAlias) string {
	result := formattedParams
	for i, alias := range aliases {
		if i < len(params) && alias.Name != "" {
			paramName := params[i].Name
			if paramName == "" {
				paramName = fmt.Sprintf("p%d", i)
			}
			paramName = g.SanitizeName(paramName)

			// Find and replace the type for this parameter
			// Look for "paramName: OldType" and replace with "paramName: NewType"
			oldPattern := paramName + ": "
			idx := strings.Index(result, oldPattern)
			if idx != -1 {
				// Find the end of the type (either comma or end of string)
				start := idx + len(oldPattern)
				end := start
				for end < len(result) && result[end] != ',' {
					end++
				}
				// Replace the type
				result = result[:start] + alias.Name + result[end:]
			}
		}
	}
	return result
}

func (g *PythonGenerator) generateMethod(method *manifest.Method) (string, error) {
	var sb strings.Builder

	// Generate function signature
	params, err := g.formatParameters(method.ParamTypes)
	if err != nil {
		return "", err
	}

	// Generate return type (with tuple for ref parameters)
	retType, err := g.generateReturnType(&method.RetType, method.ParamTypes)
	if err != nil {
		return "", err
	}

	sb.WriteString(fmt.Sprintf("def %s(%s) -> %s:\n", g.SanitizeName(method.Name), params, retType))

	// Generate docstring
	sb.WriteString(g.generateDocstring(method))

	// Stub body
	sb.WriteString("    ...\n")

	return sb.String(), nil
}

func (g *PythonGenerator) formatParameters(params []manifest.ParamType) (string, error) {
	if len(params) == 0 {
		return "", nil
	}

	result := ""
	for i, param := range params {
		if i > 0 {
			result += ", "
		}

		typeName, err := g.typeMapper.MapParamType(&param, TypeContextValue)
		if err != nil {
			return "", err
		}

		paramName := param.Name
		if paramName == "" {
			paramName = fmt.Sprintf("p%d", i)
		}

		result += g.SanitizeName(paramName) + ": " + typeName
	}

	return result, nil
}

func (g *PythonGenerator) generateReturnType(retType *manifest.TypeInfo, params []manifest.ParamType) (string, error) {
	// Check if any parameters are ref (out parameters)
	hasRef := false
	for _, param := range params {
		if param.Ref {
			hasRef = true
			break
		}
	}

	baseRetType, err := g.typeMapper.MapReturnType(retType)
	if err != nil {
		return "", err
	}

	if !hasRef {
		return baseRetType, nil
	}

	// Build tuple with return value + ref parameters
	types := []string{baseRetType}
	for _, param := range params {
		if param.Ref {
			paramType, err := g.typeMapper.MapParamType(&param, TypeContextValue)
			if err != nil {
				return "", err
			}
			types = append(types, paramType)
		}
	}

	return fmt.Sprintf("tuple[%s]", strings.Join(types, ", ")), nil
}

func (g *PythonGenerator) generateDocstring(method *manifest.Method) string {
	var sb strings.Builder

	sb.WriteString("    \"\"\"\n")
	if method.Description != "" {
		sb.WriteString(fmt.Sprintf("    %s\n\n", method.Description))
	}

	if len(method.ParamTypes) > 0 {
		sb.WriteString("    Args:\n")
		for _, param := range method.ParamTypes {
			paramName := param.Name
			if paramName == "" {
				paramName = "unnamed"
			}
			sb.WriteString(fmt.Sprintf("        %s (%s): %s\n",
				g.SanitizeName(paramName), param.Type, param.Description))
		}
	}

	if method.RetType.Type != "void" {
		sb.WriteString(fmt.Sprintf("\n    Returns:\n        %s: %s\n",
			method.RetType.Type, method.RetType.Description))
	}

	// Add callback prototypes
	for _, param := range method.ParamTypes {
		if param.Prototype != nil {
			sb.WriteString(fmt.Sprintf("\n    Callback Prototype (%s):\n", param.Prototype.Name))
			if param.Prototype.Description != "" {
				sb.WriteString(fmt.Sprintf("        %s\n\n", param.Prototype.Description))
			}
			if len(param.Prototype.ParamTypes) > 0 {
				sb.WriteString("        Args:\n")
				for _, protoParam := range param.Prototype.ParamTypes {
					sb.WriteString(fmt.Sprintf("            %s (%s): %s\n",
						protoParam.Name, protoParam.Type, protoParam.Description))
				}
			}
			if param.Prototype.RetType.Type != "void" {
				sb.WriteString(fmt.Sprintf("\n        Returns:\n            %s: %s\n",
					param.Prototype.RetType.Type, param.Prototype.RetType.Description))
			}
		}
	}

	sb.WriteString("    \"\"\"\n")
	return sb.String()
}

// PythonTypeMapper implements type mapping for Python
type PythonTypeMapper struct{}

func NewPythonTypeMapper() *PythonTypeMapper {
	return &PythonTypeMapper{}
}

func (m *PythonTypeMapper) MapType(baseType string, context TypeContext, isArray bool) (string, error) {
	typeMap := map[string]string{
		"void":   "None",
		"bool":   "bool",
		"char8":  "str",
		"char16": "str",
		"int8":   "int",
		"int16":  "int",
		"int32":  "int",
		"int64":  "int",
		"uint8":  "int",
		"uint16": "int",
		"uint32": "int",
		"uint64": "int",
		"ptr64":  "int",
		"float":  "float",
		"double": "float",
		"string": "str",
		"any":    "object",
		"vec2":   "Vector2",
		"vec3":   "Vector3",
		"vec4":   "Vector4",
		"mat4x4": "Matrix4x4",
	}

	mapped, ok := typeMap[baseType]
	if !ok {
		// Assume it's a custom type
		mapped = baseType
	}

	if isArray {
		mapped = fmt.Sprintf("list[%s]", mapped)
	}

	return mapped, nil
}

func (m *PythonTypeMapper) MapParamType(param *manifest.ParamType, context TypeContext) (string, error) {
	// Check for enum
	if param.Enum != nil {
		typeName := param.Enum.Name
		if param.IsArray() {
			typeName = fmt.Sprintf("list[%s]", typeName)
		}
		return typeName, nil
	}

	// Check for function/delegate
	if param.Prototype != nil {
		return m.generateCallableType(param.Prototype)
	}

	return m.MapType(param.BaseType(), context, param.IsArray())
}

func (m *PythonTypeMapper) MapReturnType(retType *manifest.TypeInfo) (string, error) {
	// Check for enum
	if retType.Enum != nil {
		typeName := retType.Enum.Name
		if retType.IsArray() {
			typeName = fmt.Sprintf("list[%s]", typeName)
		}
		return typeName, nil
	}

	// Check for function/delegate
	if retType.Prototype != nil {
		return m.generateCallableType(retType.Prototype)
	}

	return m.MapType(retType.BaseType(), TypeContextReturn, retType.IsArray())
}

func (m *PythonTypeMapper) generateCallableType(proto *manifest.Prototype) (string, error) {
	// Generate parameter types
	paramTypes := []string{}
	for _, param := range proto.ParamTypes {
		pType, err := m.MapParamType(&param, TypeContextValue)
		if err != nil {
			return "", err
		}
		paramTypes = append(paramTypes, pType)
	}

	// Generate return type
	retType, err := m.MapReturnType(&proto.RetType)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Callable[[%s], %s]", strings.Join(paramTypes, ", "), retType), nil
}
