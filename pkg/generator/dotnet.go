package generator

import (
	"fmt"
	"strings"

	"github.com/untrustedmodders/plugify-gen/pkg/manifest"
)

// DotnetGenerator generates C# (.NET) bindings
type DotnetGenerator struct {
	*BaseGenerator
	typeMapper *DotnetTypeMapper
}

// NewDotnetGenerator creates a new .NET/C# generator
func NewDotnetGenerator() *DotnetGenerator {
	mapper := NewDotnetTypeMapper()
	return &DotnetGenerator{
		BaseGenerator: NewBaseGenerator("dotnet", mapper, CSharpReservedWords),
		typeMapper:    mapper,
	}
}

// Generate generates .NET bindings
func (g *DotnetGenerator) Generate(m *manifest.Manifest, opts *GeneratorOptions) (*GeneratorResult, error) {
	g.ResetCaches()
	m.Sanitize(g.Sanitizer)
	opts = EnsureOptions(opts)

	files := make(map[string]string)

	// Collect all unique groups from methods and classes
	groups := g.GetGroups(m)

	// Generate separate enums file
	enumsCode, err := g.generateEnumsFile(m)
	if err != nil {
		return nil, fmt.Errorf("generating enums file: %w", err)
	}
	files[fmt.Sprintf("imported/%s/enums.cs", m.Name)] = enumsCode

	// Generate separate delegates file
	delegatesCode, err := g.generateDelegatesFile(m)
	if err != nil {
		return nil, fmt.Errorf("generating delegates file: %w", err)
	}
	files[fmt.Sprintf("imported/%s/delegates.cs", m.Name)] = delegatesCode

	// Generate group-specific files (methods and classes)
	for groupName := range groups {
		groupCode, err := g.generateGroupFile(m, groupName, opts)
		if err != nil {
			return nil, fmt.Errorf("failed to generate group %s: %w", groupName, err)
		}
		files[fmt.Sprintf("imported/%s/%s.cs", m.Name, groupName)] = groupCode
	}

	result := &GeneratorResult{
		Files: files,
	}

	return result, nil
}

// XmlDocOptions configures XML documentation generation
type XmlDocOptions struct {
	Indent       string
	Summary      string
	Params       []manifest.ParamType
	Returns      string
	ParamAliases []*manifest.ParamAlias
}

// generateXmlDocumentation generates XML documentation comments
func (g *DotnetGenerator) generateXmlDocumentation(opts XmlDocOptions) string {
	var sb strings.Builder

	// Summary
	if opts.Summary != "" {
		sb.WriteString(fmt.Sprintf("%s/// <summary>\n", opts.Indent))
		sb.WriteString(fmt.Sprintf("%s/// %s\n", opts.Indent, opts.Summary))
		sb.WriteString(fmt.Sprintf("%s/// </summary>\n", opts.Indent))
	}

	// Parameters
	for i, param := range opts.Params {
		desc := param.Description
		if desc == "" {
			// Check for alias description
			if i < len(opts.ParamAliases) && opts.ParamAliases[i] != nil {
				desc = opts.ParamAliases[i].Name + " parameter"
			} else {
				desc = param.Name
			}
		}
		sb.WriteString(fmt.Sprintf("%s/// <param name=\"%s\">%s</param>\n",
			opts.Indent, param.Name, desc))
	}

	// Returns
	if opts.Returns != "" {
		sb.WriteString(fmt.Sprintf("%s/// <returns>%s</returns>\n",
			opts.Indent, opts.Returns))
	}

	return sb.String()
}

// formatParameters is a generic helper that formats parameters using a custom formatter function
func (g *DotnetGenerator) formatParameters(params []manifest.ParamType, formatter func(int, *manifest.ParamType) (string, error)) (string, error) {
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

// processParameters is a generic helper that processes parameters and returns a list of code strings
func (g *DotnetGenerator) processParameters(params []manifest.ParamType, processor func(*manifest.ParamType) string) []string {
	var result []string
	for _, param := range params {
		code := processor(&param)
		if code != "" {
			result = append(result, code)
		}
	}
	return result
}

func (g *DotnetGenerator) generateEnums(m *manifest.Manifest) (string, error) {
	return g.CollectEnums(m, g.generateEnum)
}

func (g *DotnetGenerator) generateEnum(enum *manifest.EnumType, underlyingType string) (string, error) {
	var sb strings.Builder

	// XML documentation
	if enum.Description != "" {
		sb.WriteString(g.generateXmlDocumentation(XmlDocOptions{
			Indent:  "\t",
			Summary: enum.Description,
		}))
	}

	sb.WriteString(fmt.Sprintf("\n\t[Flags]\tpublic enum %s : %s\n\t{\n", enum.Name, underlyingType))

	for i, val := range enum.Values {
		if val.Description != "" {
			sb.WriteString(g.generateXmlDocumentation(XmlDocOptions{
				Indent:  "\t\t",
				Summary: val.Description,
			}))
		}
		sb.WriteString(fmt.Sprintf("\t\t%s = %d", val.Name, val.Value))
		if i < len(enum.Values)-1 {
			sb.WriteString(",")
		}
		sb.WriteString("\n")
	}

	sb.WriteString("\t}\n")
	return sb.String(), nil
}

func (g *DotnetGenerator) generateDelegates(m *manifest.Manifest) (string, error) {
	return g.CollectDelegates(m, g.generateDelegate)
}

func (g *DotnetGenerator) generateDelegate(proto *manifest.Prototype) (string, error) {
	var sb strings.Builder

	// XML documentation
	if proto.Description != "" {
		sb.WriteString(g.generateXmlDocumentation(XmlDocOptions{
			Indent:  "\t",
			Summary: proto.Description,
		}))
	}

	// Return type
	retType, err := g.typeMapper.MapDelegateReturnType(&proto.RetType)
	if err != nil {
		return "", err
	}

	// Parameters
	params, err := g.formatDelegateParameters(proto.ParamTypes)
	if err != nil {
		return "", err
	}

	sb.WriteString(fmt.Sprintf("\tpublic delegate %s %s(%s);\n", retType, proto.Name, params))

	return sb.String(), nil
}

func (g *DotnetGenerator) formatDelegateParameters(params []manifest.ParamType) (string, error) {
	return g.formatParameters(params, func(i int, param *manifest.ParamType) (string, error) {
		typeName, err := g.typeMapper.MapDelegateParamType(param)
		if err != nil {
			return "", err
		}
		return typeName + " " + param.Name, nil
	})
}

func (g *DotnetGenerator) generateMethod(method *manifest.Method, pluginName string) (string, error) {
	var sb strings.Builder

	// XML documentation
	sb.WriteString(g.generateDocumentation(method))

	// Generate function pointer declarations
	managedTypes, err := g.formatManagedTypes(method)
	if err != nil {
		return "", err
	}

	unmanagedTypes, err := g.formatUnmanagedTypes(method)
	if err != nil {
		return "", err
	}

	methodName := method.Name

	// Managed delegate pointer
	sb.WriteString(fmt.Sprintf("\t\tinternal static delegate*<%s> %s = &___%s;\n", managedTypes, methodName, methodName))

	// Unmanaged function pointer
	sb.WriteString(fmt.Sprintf("\t\tinternal static delegate* unmanaged[Cdecl]<%s> __%s;\n", unmanagedTypes, methodName))

	// Wrapper method signature
	params, err := g.formatMethodParameters(method.ParamTypes)
	if err != nil {
		return "", err
	}

	retType, err := g.typeMapper.MapReturnType(&method.RetType)
	if err != nil {
		return "", err
	}

	sb.WriteString(fmt.Sprintf("\t\tprivate static %s ___%s(%s)\n", retType, methodName, params))
	sb.WriteString("\t\t{\n")

	// Method body
	body, err := g.generateMethodBody(method)
	if err != nil {
		return "", err
	}
	sb.WriteString(body)

	sb.WriteString("\t\t}\n")

	return sb.String(), nil
}

func (g *DotnetGenerator) generateDocumentation(method *manifest.Method) string {
	summary := method.Description
	if summary == "" {
		summary = method.Name
	}

	returns := ""
	if method.RetType.Type != "void" && method.RetType.Description != "" {
		returns = method.RetType.Description
	}

	return g.generateXmlDocumentation(XmlDocOptions{
		Indent:  "\t\t",
		Summary: summary,
		Params:  method.ParamTypes,
		Returns: returns,
	})
}

func (g *DotnetGenerator) formatManagedTypes(method *manifest.Method) (string, error) {
	typesStr, err := g.formatParameters(method.ParamTypes, func(_ int, param *manifest.ParamType) (string, error) {
		return g.typeMapper.MapParamType(param, TypeContextValue)
	})
	if err != nil {
		return "", err
	}

	// Add return type
	retType, err := g.typeMapper.MapReturnType(&method.RetType)
	if err != nil {
		return "", err
	}

	if typesStr != "" {
		return typesStr + ", " + retType, nil
	}
	return retType, nil
}

func (g *DotnetGenerator) formatUnmanagedTypes(method *manifest.Method) (string, error) {
	typesStr, err := g.formatParameters(method.ParamTypes, func(_ int, param *manifest.ParamType) (string, error) {
		return g.typeMapper.MapUnmanagedParamType(param)
	})
	if err != nil {
		return "", err
	}

	// Add return type
	retType, err := g.typeMapper.MapUnmanagedReturnType(&method.RetType)
	if err != nil {
		return "", err
	}

	if typesStr != "" {
		return typesStr + ", " + retType, nil
	}
	return retType, nil
}

func (g *DotnetGenerator) formatMethodParameters(params []manifest.ParamType) (string, error) {
	return g.formatParameters(params, func(_ int, param *manifest.ParamType) (string, error) {
		typeName, err := g.typeMapper.MapParamType(param, TypeContextValue)
		if err != nil {
			return "", err
		}

		result := typeName + " " + param.Name
		if param.Default != nil {
			result += fmt.Sprintf(" = %d", *param.Default)
		}
		return result, nil
	})
}

func (g *DotnetGenerator) generateMethodBody(method *manifest.Method) (string, error) {
	var sb strings.Builder
	indent := "\t\t\t"

	needsMarshaling := g.needsMarshaling(method)
	hasReturn := method.RetType.Type != "void"
	isObjectReturn := g.typeMapper.isObjectReturn(method.RetType.Type)
	isFunctionReturn := g.typeMapper.isFunction(method.RetType.Type)

	// Check if we'll have fixed blocks
	hasFixedBlocks := false
	for _, param := range method.ParamTypes {
		if param.Ref && !g.typeMapper.isObjectReturn(param.Type) {
			hasFixedBlocks = true
			break
		}
	}

	// Declare return variable if needed (before fixed blocks to avoid scope issues)
	declareRetVal := hasReturn && (isObjectReturn || needsMarshaling || hasFixedBlocks)
	if declareRetVal {
		retTypeName, _ := g.typeMapper.MapReturnType(&method.RetType)
		sb.WriteString(fmt.Sprintf("%s%s __retVal;\n", indent, retTypeName))
	}

	// Declare native return variable for object returns
	if isObjectReturn {
		nativeRetType := g.typeMapper.getNativeReturnType(method.RetType.Type)
		sb.WriteString(fmt.Sprintf("%s%s __retVal_native;\n", indent, nativeRetType))
	}

	// Generate fixed blocks for ref primitive/POD parameters
	fixedBlocks := g.generateFixedBlocks(method, indent)
	for _, fixedBlock := range fixedBlocks {
		sb.WriteString(fixedBlock)
	}

	// Parameter marshaling (constructor calls)
	marshalCode := g.generateParameterMarshaling(method, indent)
	if marshalCode != "" {
		sb.WriteString(marshalCode)
	}

	hasTry := (needsMarshaling && hasReturn)
	innerIndent := indent
	insertIndexStart := 0
	insertIndexEnd := 0
	if hasTry || marshalCode != "" {
		insertIndexStart = sb.Len()
		sb.WriteString(fmt.Sprintf("%stry {\n", indent))
		insertIndexEnd = sb.Len()
		innerIndent = indent + "\t"
	}

	methodName := method.Name

	// Function call
	callParams := g.generateCallParameters(method)
	if isObjectReturn {
		sb.WriteString(fmt.Sprintf("%s__retVal_native = __%s(%s);\n", innerIndent, methodName, callParams))
	} else if hasTry || declareRetVal {
		sb.WriteString(fmt.Sprintf("%s__retVal = __%s(%s);\n", innerIndent, methodName, callParams))
	} else if hasReturn {
		retTypeName, _ := g.typeMapper.MapReturnType(&method.RetType)
		sb.WriteString(fmt.Sprintf("%s%s __retVal = __%s(%s);\n", innerIndent, retTypeName, methodName, callParams))
	} else {
		sb.WriteString(fmt.Sprintf("%s__%s(%s);\n", innerIndent, methodName, callParams))
	}

	// Unmarshal return value and ref parameters
	unmarshalCode := g.generateUnmarshaling(method, innerIndent)
	if unmarshalCode != "" {
		sb.WriteString(fmt.Sprintf("%s// Unmarshal - Convert native data to managed data.\n", innerIndent))
		sb.WriteString(unmarshalCode)
	}

	// Cleanup
	cleanupCode := g.generateCleanup(method, innerIndent)
	if cleanupCode != "" {
		sb.WriteString(fmt.Sprintf("%s}\n%sfinally {\n", indent, indent))
		sb.WriteString(fmt.Sprintf("%s// Perform cleanup.\n", innerIndent))
		sb.WriteString(cleanupCode)
		sb.WriteString(fmt.Sprintf("%s}\n", indent))
	} else if hasTry {
		// Remove try block if no cleanup needed
		RemoveFromBuilder(&sb, insertIndexStart, insertIndexEnd)
		RemoveLeadingTabs(&sb, 1, insertIndexStart, sb.Len())
	}

	// Close fixed blocks
	for i := len(fixedBlocks) - 1; i >= 0; i-- {
		sb.WriteString(fmt.Sprintf("%s}\n", indent))
	}

	// Return statement
	if hasReturn {
		if isFunctionReturn {
			retTypeName, _ := g.typeMapper.MapReturnType(&method.RetType)
			sb.WriteString(fmt.Sprintf("%sreturn Marshalling.GetDelegateForFunctionPointer<%s>(__retVal);\n", indent, retTypeName))
		} else {
			sb.WriteString(fmt.Sprintf("%sreturn __retVal;\n", indent))
		}
	}

	return sb.String(), nil
}

func (g *DotnetGenerator) needsMarshaling(method *manifest.Method) bool {
	if g.typeMapper.isObjectReturn(method.RetType.Type) {
		return true
	}

	for _, param := range method.ParamTypes {
		if g.typeMapper.isObjectReturn(param.Type) {
			return true
		}
	}

	return false
}

func (g *DotnetGenerator) generateFixedBlocks(method *manifest.Method, indent string) []string {
	var blocks []string

	for _, param := range method.ParamTypes {
		// Only create fixed blocks for ref parameters that don't need marshaling (primitives, POD, enums)
		if !param.Ref {
			continue
		}

		// Skip object types (they need marshaling, not fixed blocks)
		if g.typeMapper.isObjectReturn(param.Type) {
			continue
		}

		paramName := param.Name
		var typeName string

		if param.Enum != nil {
			typeName = param.Enum.Name
		} else {
			typeName, _ = g.typeMapper.MapType(param.BaseType(), TypeContextValue, false)
		}

		block := fmt.Sprintf("%sfixed(%s* __%s = &%s) {\n", indent, typeName, paramName, paramName)
		blocks = append(blocks, block)
	}

	return blocks
}

func (g *DotnetGenerator) generateParameterMarshaling(method *manifest.Method, indent string) string {
	codes := g.processParameters(method.ParamTypes, func(param *manifest.ParamType) string {
		constructor := g.typeMapper.getConstructor(param.Type)
		if constructor == "" {
			return ""
		}

		paramName := param.Name

		// Use NativeMethodsT for enum parameters
		if param.Enum != nil {
			constructor = strings.Replace(constructor, "NativeMethods.", "NativeMethodsT.", 1)
		}

		if strings.Contains(constructor, "ConstructVector") {
			return fmt.Sprintf("%svar __%s = %s(%s, %s.Length);\n", indent, paramName, constructor, paramName, paramName)
		} else if constructor == "Marshalling.GetFunctionPointerForDelegate" {
			// Function pointer - handled differently
			return ""
		} else {
			return fmt.Sprintf("%svar __%s = %s(%s);\n", indent, paramName, constructor, paramName)
		}
	})

	return strings.Join(codes, "")
}

func (g *DotnetGenerator) generateCallParameters(method *manifest.Method) string {
	params := g.processParameters(method.ParamTypes, func(param *manifest.ParamType) string {
		paramName := param.Name

		if g.typeMapper.isObjectReturn(param.Type) {
			return fmt.Sprintf("&__%s", paramName)
		} else if g.typeMapper.isPODType(param.Type) && param.Ref {
			return fmt.Sprintf("__%s", paramName)
		} else if g.typeMapper.isPODType(param.Type) {
			return fmt.Sprintf("&%s", paramName)
		} else if g.typeMapper.isFunction(param.Type) {
			return fmt.Sprintf("Marshalling.GetFunctionPointerForDelegate(%s)", paramName)
		} else if param.Ref {
			return fmt.Sprintf("__%s", paramName)
		}
		return paramName
	})

	return strings.Join(params, ", ")
}

func (g *DotnetGenerator) generateUnmarshaling(method *manifest.Method, indent string) string {
	var sb strings.Builder

	// Unmarshal return value
	if g.typeMapper.isObjectReturn(method.RetType.Type) {
		converter := g.typeMapper.getDataConverter(method.RetType.Type)
		if converter != "" {
			// Use NativeMethodsT for enum parameters
			if method.RetType.Enum != nil {
				converter = strings.Replace(converter, "NativeMethods.", "NativeMethodsT.", 1)
			}

			if strings.Contains(converter, "VectorData") {
				sizeFunc := g.typeMapper.getSizeFunction(method.RetType.Type)
				retTypeName, _ := g.typeMapper.MapReturnType(&method.RetType)
				// Remove [] suffix for array initialization
				baseType := strings.TrimSuffix(retTypeName, "[]")
				sb.WriteString(fmt.Sprintf("%s__retVal = new %s[%s(&__retVal_native)];\n", indent, baseType, sizeFunc))
				sb.WriteString(fmt.Sprintf("%s%s(&__retVal_native, __retVal);\n", indent, converter))
			} else {
				sb.WriteString(fmt.Sprintf("%s__retVal = %s(&__retVal_native);\n", indent, converter))
			}
		}
	}

	// Unmarshal ref parameters
	codes := g.processParameters(method.ParamTypes, func(param *manifest.ParamType) string {
		if !param.Ref {
			return ""
		}

		converter := g.typeMapper.getDataConverter(param.Type)
		if converter == "" {
			return ""
		}

		paramName := param.Name

		// Use NativeMethodsT for enum parameters
		if param.Enum != nil {
			converter = strings.Replace(converter, "NativeMethods.", "NativeMethodsT.", 1)
		}

		if strings.Contains(converter, "VectorData") {
			sizeFunc := g.typeMapper.getSizeFunction(param.Type)
			return fmt.Sprintf("%sArray.Resize(ref %s, %s(&__%s));\n%s%s(&__%s, %s);\n",
				indent, paramName, sizeFunc, paramName,
				indent, converter, paramName, paramName)
		} else if converter != "" {
			return fmt.Sprintf("%s%s = %s(&__%s);\n", indent, paramName, converter, paramName)
		}
		return ""
	})
	sb.WriteString(strings.Join(codes, ""))

	return sb.String()
}

func (g *DotnetGenerator) generateCleanup(method *manifest.Method, indent string) string {
	var sb strings.Builder

	// Cleanup return value
	if g.typeMapper.isObjectReturn(method.RetType.Type) {
		destructor := g.typeMapper.getDestructor(method.RetType.Type)
		if destructor != "" {
			sb.WriteString(fmt.Sprintf("%s%s(&__retVal_native);\n", indent, destructor))
		}
	}

	// Cleanup parameters
	codes := g.processParameters(method.ParamTypes, func(param *manifest.ParamType) string {
		destructor := g.typeMapper.getDestructor(param.Type)
		if destructor == "" {
			return ""
		}
		return fmt.Sprintf("%s%s(&__%s);\n", indent, destructor, param.Name)
	})
	sb.WriteString(strings.Join(codes, ""))

	return sb.String()
}

func (g *DotnetGenerator) generateClasses(m *manifest.Manifest) (string, error) {
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

func (g *DotnetGenerator) generateClass(m *manifest.Manifest, class *manifest.Class) (string, error) {
	var sb strings.Builder

	// Check if this is a handleless class (static methods only)
	hasHandle := class.HandleType != "" && class.HandleType != "void"

	// Map handle type
	invalidValue, handleType, err := g.typeMapper.MapHandleType(class)
	if err != nil {
		return "", err
	}

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
	summary := class.Description
	if summary == "" {
		if hasHandle {
			summary = class.Name + " wrapper"
		} else {
			summary = "Static utility class for " + class.Name
		}
	}
	sb.WriteString(g.generateXmlDocumentation(XmlDocOptions{
		Indent:  "\t",
		Summary: summary,
	}))

	// Class declaration
	if hasHandle {
		if hasDtor {
			sb.WriteString(fmt.Sprintf("\tinternal sealed unsafe class %s : SafeHandle\n\t{\n", class.Name))
		} else {
			sb.WriteString(fmt.Sprintf("\tinternal sealed unsafe class %s\n\t{\n", class.Name))
			sb.WriteString(fmt.Sprintf("\t\tprivate %s handle;\n\n", handleType))
		}
	} else {
		sb.WriteString(fmt.Sprintf("\tinternal static unsafe class %s\n\t{\n", class.Name))
	}

	// Only generate handle-related code if class has a handle
	if hasHandle {
		// Default constructor
		hasDefaultConstructor := g.HasConstructorWithNoParam(m, class)
		if !hasDefaultConstructor {
			sb.WriteString(fmt.Sprintf("\t\tpublic %s() {}\n", class.Name))
		}

		// Generate constructors
		for _, ctorName := range class.Constructors {
			ctorCode, err := g.generateClassConstructor(m, class, ctorName)
			if err != nil {
				return "", err
			}
			sb.WriteString(ctorCode)
		}

		// Internal constructor for handle wrapping
		if hasDtor {
			// Check if there's a constructor with exactly 1 param of handle type to avoid ambiguity
			hasHandleOnlyConstructor := g.HasConstructorWithSingleHandleParam(m, class)
			ownershipDefault := ""
			if !hasHandleOnlyConstructor {
				ownershipDefault = " = Ownership.Borrowed"
			}

			sb.WriteString("\t\t/// <summary>\n")
			sb.WriteString(fmt.Sprintf("\t\t/// Internal constructor for creating %s from existing handle\n", class.Name))
			sb.WriteString("\t\t/// </summary>\n")
			sb.WriteString(fmt.Sprintf("\t\tpublic %s(%s handle, Ownership ownership%s) : base((nint)handle, ownsHandle: ownership == Ownership.Owned)\n", class.Name, handleType, ownershipDefault))
			sb.WriteString("\t\t{\n\t\t}\n\n")

			// ReleaseHandle override
			sb.WriteString("\t\t/// <summary>\n")
			sb.WriteString("\t\t/// Releases the handle (called automatically by SafeHandle)\n")
			sb.WriteString("\t\t/// </summary>\n")
			sb.WriteString("\t\tprotected override bool ReleaseHandle()\n")
			sb.WriteString("\t\t{\n")
			sb.WriteString(fmt.Sprintf("\t\t\t%s.%s((%s)handle);\n", m.Name, *class.Destructor, handleType))
			sb.WriteString("\t\t\treturn true;\n")
			sb.WriteString("\t\t}\n\n")

			// IsInvalid override
			sb.WriteString("\t\t/// <summary>\n")
			sb.WriteString(fmt.Sprintf("\t\t/// Checks if the %s has a valid handle\n", class.Name))
			sb.WriteString("\t\t/// </summary>\n")
			sb.WriteString(fmt.Sprintf("\t\tpublic override bool IsInvalid => handle == %s;\n\n", invalidValue))
		} else {
			ownershipTag := ""
			if hasCtor {
				ownershipTag = ", Ownership ownership = Ownership.Borrowed"
			}
			sb.WriteString("\t\t/// <summary>\n")
			sb.WriteString(fmt.Sprintf("\t\t/// Internal constructor for creating %s from existing handle\n", class.Name))
			sb.WriteString("\t\t/// </summary>\n")
			sb.WriteString(fmt.Sprintf("\t\tpublic %s(%s handle%s)\n", class.Name, handleType, ownershipTag))
			sb.WriteString("\t\t{\n")
			sb.WriteString("\t\t\tthis.handle = handle;\n")
			sb.WriteString("\t\t}\n\n")
		}

		// Utility properties and methods
		if hasDtor {
			sb.WriteString("\t\t/// <summary>\n")
			sb.WriteString("\t\t/// Gets the underlying handle\n")
			sb.WriteString("\t\t/// </summary>\n")
			sb.WriteString(fmt.Sprintf("\t\tpublic %s Handle => (%s)handle;\n\n", handleType, handleType))

			sb.WriteString("\t\t/// <summary>\n")
			sb.WriteString("\t\t/// Checks if the handle is valid\n")
			sb.WriteString("\t\t/// </summary>\n")
			sb.WriteString(fmt.Sprintf("\t\tpublic bool IsValid => handle != %s;\n\n", invalidValue))

			sb.WriteString("\t\t/// <summary>\n")
			sb.WriteString("\t\t/// Gets the underlying handle\n")
			sb.WriteString("\t\t/// </summary>\n")
			sb.WriteString(fmt.Sprintf("\t\tpublic %s Get() => (%s)handle;\n\n", handleType, handleType))

			sb.WriteString("\t\t/// <summary>\n")
			sb.WriteString("\t\t/// Releases ownership of the handle and returns it\n")
			sb.WriteString("\t\t/// </summary>\n")
			sb.WriteString(fmt.Sprintf("\t\tpublic %s Release()\n", handleType))
			sb.WriteString("\t\t{\n")
			sb.WriteString("\t\t\tvar h = handle;\n")
			sb.WriteString("\t\t\tSetHandleAsInvalid();\n")
			sb.WriteString(fmt.Sprintf("\t\t\treturn (%s)h;\n", handleType))
			sb.WriteString("\t\t}\n\n")

			sb.WriteString("\t\t/// <summary>\n")
			sb.WriteString("\t\t/// Disposes the handle\n")
			sb.WriteString("\t\t/// </summary>\n")
			sb.WriteString("\t\tpublic void Reset()\n")
			sb.WriteString("\t\t{\n")
			sb.WriteString("\t\t\tDispose();\n")
			sb.WriteString("\t\t}\n\n")
		} else {
			sb.WriteString("\t\t/// <summary>\n")
			sb.WriteString("\t\t/// Gets the underlying handle\n")
			sb.WriteString("\t\t/// </summary>\n")
			sb.WriteString(fmt.Sprintf("\t\tpublic %s Handle => handle;\n\n", handleType))

			sb.WriteString("\t\t/// <summary>\n")
			sb.WriteString("\t\t/// Checks if the handle is valid\n")
			sb.WriteString("\t\t/// </summary>\n")
			sb.WriteString(fmt.Sprintf("\t\tpublic bool IsValid => handle != %s;\n\n", invalidValue))

			sb.WriteString("\t\t/// <summary>\n")
			sb.WriteString("\t\t/// Gets the underlying handle\n")
			sb.WriteString("\t\t/// </summary>\n")
			sb.WriteString(fmt.Sprintf("\t\tpublic %s Get() => handle;\n\n", handleType))

			sb.WriteString("\t\t/// <summary>\n")
			sb.WriteString("\t\t/// Releases ownership of the handle and returns it\n")
			sb.WriteString("\t\t/// </summary>\n")
			sb.WriteString(fmt.Sprintf("\t\tpublic %s Release()\n", handleType))
			sb.WriteString("\t\t{\n")
			sb.WriteString("\t\t\tvar h = handle;\n")
			sb.WriteString(fmt.Sprintf("\t\t\thandle = %s;\n", invalidValue))
			sb.WriteString("\t\t\treturn h;\n")
			sb.WriteString("\t\t}\n\n")

			sb.WriteString("\t\t/// <summary>\n")
			sb.WriteString("\t\t/// Resets the handle to invalid\n")
			sb.WriteString("\t\t/// </summary>\n")
			sb.WriteString("\t\tpublic void Reset()\n")
			sb.WriteString("\t\t{\n")
			sb.WriteString(fmt.Sprintf("\t\t\thandle = %s;\n", invalidValue))
			sb.WriteString("\t\t}\n\n")
		}
	}

	// Generate class bindings
	for _, binding := range class.Bindings {
		bindingCode, err := g.generateClassBinding(m, class, &binding)
		if err != nil {
			return "", err
		}
		sb.WriteString(bindingCode)
	}

	sb.WriteString("\t}\n\n")

	return sb.String(), nil
}

func (g *DotnetGenerator) generateClassConstructor(m *manifest.Manifest, class *manifest.Class, methodName string) (string, error) {
	// Find the method in the manifest
	method := FindMethod(m, methodName)
	if method == nil {
		return "", fmt.Errorf("constructor method %s not found", methodName)
	}

	var sb strings.Builder

	// Generate documentation
	summary := method.Description
	if summary == "" {
		summary = "Creates a new " + class.Name + " instance"
	}
	sb.WriteString(g.generateXmlDocumentation(XmlDocOptions{
		Indent:  "\t\t",
		Summary: summary,
		Params:  method.ParamTypes,
	}))

	// Generate constructor signature
	params, err := g.formatMethodParameters(method.ParamTypes)
	if err != nil {
		return "", err
	}

	sb.WriteString(fmt.Sprintf("\t\tpublic %s(%s)", class.Name, params))

	// Generate initialization
	paramNames := []string{}
	for _, param := range method.ParamTypes {
		paramNames = append(paramNames, param.Name)
	}

	hasDtor := class.Destructor != nil

	if hasDtor {
		sb.WriteString(fmt.Sprintf(" : this(%s.%s(%s), Ownership.Owned)\n",
			m.Name, methodName, strings.Join(paramNames, ", ")))
		sb.WriteString("\t\t{\n\t\t}\n\n")
	} else {
		sb.WriteString("\n\t\t{\n")
		sb.WriteString(fmt.Sprintf("\t\t\tthis.handle = %s.%s(%s);\n",
			m.Name, methodName, strings.Join(paramNames, ", ")))
		sb.WriteString("\t\t}\n\n")
	}

	return sb.String(), nil
}

func (g *DotnetGenerator) generateClassBinding(m *manifest.Manifest, class *manifest.Class, binding *manifest.Binding) (string, error) {
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
	summary := method.Description
	if summary == "" {
		summary = binding.Name
	}

	returns := ""
	if method.RetType.Type != "void" {
		returns = method.RetType.Description
		if returns == "" {
			if binding.RetAlias != nil && binding.RetAlias.Name != "" {
				returns = binding.RetAlias.Name + " instance"
			} else {
				returns = "Return value"
			}
		}
	}

	sb.WriteString(g.generateXmlDocumentation(XmlDocOptions{
		Indent:       "\t\t",
		Summary:      summary,
		Params:       methodParams,
		Returns:      returns,
		ParamAliases: binding.ParamAliases,
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

	// Format parameters with aliases
	formattedParams := ""
	for i, param := range methodParams {
		if i > 0 {
			formattedParams += ", "
		}

		paramType, err := g.typeMapper.MapParamType(&param, TypeContextValue)
		if err != nil {
			return "", err
		}

		// Apply parameter alias if exists
		if i < len(binding.ParamAliases) && binding.ParamAliases[i] != nil {
			paramType = binding.ParamAliases[i].Name
		}

		formattedParams += paramType + " " + param.Name

		// Add default value if present
		if param.Default != nil {
			formattedParams += fmt.Sprintf(" = %d", *param.Default)
		}
	}

	// Determine if method is static
	if !binding.BindSelf {
		sb.WriteString(fmt.Sprintf("\t\tpublic static %s %s(%s)\n", retType, binding.Name, formattedParams))
	} else {
		sb.WriteString(fmt.Sprintf("\t\tpublic %s %s(%s)\n", retType, binding.Name, formattedParams))
	}

	sb.WriteString("\t\t{\n")

	// Add validity check for instance methods
	if binding.BindSelf {
		sb.WriteString("\t\t\tObjectDisposedException.ThrowIf(!IsValid, this);\n")
	}

	hasCtor := len(class.Constructors) > 0
	hasDtor := class.Destructor != nil

	methodName := method.Name

	// Generate method body for SafeHandle classes (need ref counting)
	if hasDtor && binding.BindSelf {
		sb.WriteString("\t\t\tbool success = false;\n")
		sb.WriteString("\t\t\tDangerousAddRef(ref success);\n")
		sb.WriteString("\t\t\ttry\n")
		sb.WriteString("\t\t\t{\n")

		// Build call arguments
		callArgs := g.buildCallArguments(binding, methodParams, "Handle")

		// Generate call
		if method.RetType.Type == "void" {
			sb.WriteString(fmt.Sprintf("\t\t\t\t%s.%s(%s);\n",
				m.Name, methodName, callArgs))
		} else {
			if binding.RetAlias != nil && binding.RetAlias.Name != "" {
				ownership := ""
				if hasDtor || hasCtor {
					if binding.RetAlias.Owner {
						ownership = ", Ownership.Owned"
					} else {
						ownership = ", Ownership.Borrowed"
					}
				}
				sb.WriteString(fmt.Sprintf("\t\t\t\treturn new %s(%s.%s(%s)%s);\n",
					binding.RetAlias.Name, m.Name, methodName, callArgs, ownership))
			} else {
				sb.WriteString(fmt.Sprintf("\t\t\t\treturn %s.%s(%s);\n",
					m.Name, methodName, callArgs))
			}
		}

		sb.WriteString("\t\t\t}\n")
		sb.WriteString("\t\t\tfinally\n")
		sb.WriteString("\t\t\t{\n")
		sb.WriteString("\t\t\t\tif (success) DangerousRelease();\n")
		sb.WriteString("\t\t\t}\n")
	} else {
		// Build call arguments for non-SafeHandle or static methods
		callArgs := g.buildCallArguments(binding, methodParams, "handle")

		// Generate call
		if method.RetType.Type == "void" {
			sb.WriteString(fmt.Sprintf("\t\t\t%s.%s(%s);\n",
				m.Name, methodName, callArgs))
		} else {
			if binding.RetAlias != nil && binding.RetAlias.Name != "" {
				ownership := ""
				if hasDtor || hasCtor {
					if binding.RetAlias.Owner {
						ownership = ", Ownership.Owned"
					} else {
						ownership = ", Ownership.Borrowed"
					}
				}
				sb.WriteString(fmt.Sprintf("\t\t\treturn new %s(%s.%s(%s)%s);\n",
					binding.RetAlias.Name, m.Name, methodName, callArgs, ownership))
			} else {
				sb.WriteString(fmt.Sprintf("\t\t\treturn %s.%s(%s);\n",
					m.Name, methodName, callArgs))
			}
		}
	}

	sb.WriteString("\t\t}\n\n")

	return sb.String(), nil
}

func (g *DotnetGenerator) buildCallArguments(binding *manifest.Binding, methodParams []manifest.ParamType, handleParam string) string {
	callArgs := ""

	// Add self handle if bindSelf
	if binding.BindSelf {
		callArgs = handleParam
	}

	// Add other parameters
	for i, param := range methodParams {
		if callArgs != "" {
			callArgs += ", "
		}

		paramName := ""
		if param.Ref {
			paramName += "ref "
		}
		paramName += param.Name

		// Check if parameter has alias and needs .Release() or .Get()
		if i < len(binding.ParamAliases) && binding.ParamAliases[i] != nil {
			if binding.ParamAliases[i].Owner {
				callArgs += paramName + ".Release()"
			} else {
				callArgs += paramName + ".Get()"
			}
		} else {
			callArgs += paramName
		}
	}

	return callArgs
}

// generateEnumsFile generates a file containing all enums
func (g *DotnetGenerator) generateEnumsFile(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	// Using statements
	sb.WriteString("using System;\n\n")
	sb.WriteString(fmt.Sprintf("// Generated from %s.pplugin\n\n", m.Name))

	// Namespace
	sb.WriteString(fmt.Sprintf("namespace %s {\n", m.Name))
	sb.WriteString("#pragma warning disable CS0649\n\n")

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
	sb.WriteString("\t/// <summary>\n")
	sb.WriteString("\t/// Ownership type for RAII wrappers\n")
	sb.WriteString("\t/// </summary>\n")
	sb.WriteString("\tinternal enum Ownership { Borrowed, Owned }\n\n")

	sb.WriteString("#pragma warning restore CS0649\n")
	sb.WriteString("}\n")

	return sb.String(), nil
}

// generateDelegatesFile generates a file containing all delegate definitions
func (g *DotnetGenerator) generateDelegatesFile(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	// Using statements
	sb.WriteString("using System;\n")
	sb.WriteString("using System.Numerics;\n\n")
	sb.WriteString("using Plugify;\n\n")
	sb.WriteString(fmt.Sprintf("// Generated from %s.pplugin\n\n", m.Name))

	// Namespace
	sb.WriteString(fmt.Sprintf("namespace %s {\n", m.Name))
	sb.WriteString("#pragma warning disable CS0649\n\n")

	// Generate delegates
	delegatesCode, err := g.generateDelegates(m)
	if err != nil {
		return "", err
	}
	if delegatesCode != "" {
		sb.WriteString(delegatesCode)
	}

	sb.WriteString("#pragma warning restore CS0649\n")
	sb.WriteString("}\n")

	return sb.String(), nil
}

// generateGroupFile generates a group-specific file containing methods for that group
func (g *DotnetGenerator) generateGroupFile(m *manifest.Manifest, groupName string, opts *GeneratorOptions) (string, error) {
	var sb strings.Builder

	// Using statements
	sb.WriteString("using System;\n")
	sb.WriteString("using System.Numerics;\n")
	sb.WriteString("using System.Runtime.CompilerServices;\n")
	sb.WriteString("using System.Runtime.InteropServices;\n")
	sb.WriteString("using Plugify;\n\n")
	sb.WriteString(fmt.Sprintf("// Generated from %s.pplugin (group: %s)\n\n", m.Name, groupName))

	// Namespace
	sb.WriteString(fmt.Sprintf("namespace %s {\n", m.Name))
	sb.WriteString("#pragma warning disable CS0649\n\n")

	// Class containing methods for this group
	sb.WriteString(fmt.Sprintf("\tinternal static unsafe partial class %s {\n\n", m.Name))

	// Generate methods for this group
	for _, method := range m.Methods {
		methodGroup := method.Group
		if methodGroup == groupName {
			methodCode, err := g.generateMethod(&method, m.Name)
			if err != nil {
				return "", fmt.Errorf("failed to generate method %s: %w", method.Name, err)
			}
			sb.WriteString(methodCode)
			sb.WriteString("\n")
		}
	}

	sb.WriteString("\t}\n\n")

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
			}
		}
	}

	sb.WriteString("#pragma warning restore CS0649\n")
	sb.WriteString("}\n")

	return sb.String(), nil
}
