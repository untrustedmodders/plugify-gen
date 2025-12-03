package generator

import (
	"fmt"
	"strings"

	"github.com/untrustedmodders/plugify-gen/pkg/manifest"
)

// GolangGenerator generates Go bindings
type GolangGenerator struct {
	*BaseGenerator
	typeMapper *GolangTypeMapper
	usedNames  map[string]bool
}

// NewGolangGenerator creates a new Go generator
func NewGolangGenerator() *GolangGenerator {
	invalidNames := []string{
		"break", "case", "chan", "const", "continue", "default", "defer", "else",
		"fallthrough", "for", "func", "go", "goto", "if", "import", "interface",
		"map", "package", "range", "return", "select", "struct", "switch", "type",
		"var", "append", "bool", "byte", "cap", "close", "complex", "complex64",
		"complex128", "copy", "delete", "error", "false", "float32", "float64",
		"imag", "int", "int8", "int16", "int32", "int64", "iota", "len", "make",
		"new", "nil", "panic", "print", "println", "real", "recover", "rune",
		"string", "true", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr",
	}

	mapper := NewGolangTypeMapper()
	return &GolangGenerator{
		BaseGenerator: NewBaseGenerator("golang", mapper, invalidNames),
		typeMapper:    mapper,
		usedNames:     make(map[string]bool),
	}
}

// Generate generates Go bindings (.go and .h files)
func (g *GolangGenerator) Generate(m *manifest.Manifest) (*GeneratorResult, error) {
	g.ResetCaches()
	g.usedNames = make(map[string]bool)

	files := make(map[string]string)

	// Collect all unique groups from methods and classes
	groups := g.GetGroups(m)

	// Generate separate enums file
	enumsGoCode, err := g.generateEnumsGoFile(m)
	if err != nil {
		return nil, err
	}
	files[fmt.Sprintf("%s/enums.go", m.Name)] = enumsGoCode

	// Generate separate delegates file
	delegatesGoCode, err := g.generateDelegatesGoFile(m)
	if err != nil {
		return nil, err
	}
	files[fmt.Sprintf("%s/delegates.go", m.Name)] = delegatesGoCode

	// Generate .h file for enums
	sharedHCode, err := g.generateSharedHFile(m)
	if err != nil {
		return nil, err
	}
	files[fmt.Sprintf("%s/shared.h", m.Name)] = sharedHCode

	// Generate group-specific files
	for groupName := range groups {
		goCode, err := g.generateGroupGoFile(m, groupName)
		if err != nil {
			return nil, fmt.Errorf("failed to generate group %s: %w", groupName, err)
		}
		files[fmt.Sprintf("%s/%s.go", m.Name, groupName)] = goCode

		hCode, err := g.generateGroupHFile(m, groupName)
		if err != nil {
			return nil, fmt.Errorf("failed to generate group %s header: %w", groupName, err)
		}
		files[fmt.Sprintf("%s/%s.h", m.Name, groupName)] = hCode
	}

	result := &GeneratorResult{
		Files: files,
	}

	return result, nil
}

// generateEnums generates enum definitions
func (g *GolangGenerator) generateEnums(m *manifest.Manifest) (string, error) {
	var sb strings.Builder
	localUsed := make(map[string]int)

	var processEnum func(enum *manifest.EnumType, enumType string) error
	processEnum = func(enum *manifest.EnumType, enumType string) error {
		if g.IsEnumCached(enum.Name) {
			return nil
		}
		g.CacheEnum(enum.Name)

		goType, err := g.typeMapper.MapType(enumType, TypeContextValue, false)
		if err != nil {
			return err
		}

		// Add enum description
		if enum.Description != "" {
			sb.WriteString(fmt.Sprintf("// %s - %s\n", enum.Name, enum.Description))
		}

		sb.WriteString(fmt.Sprintf("type %s = %s\n\n", enum.Name, goType))
		sb.WriteString("const (\n")

		for i, value := range enum.Values {
			rawName := value.Name
			if rawName == "" {
				rawName = fmt.Sprintf("Value%d", i)
			}

			baseName := g.SanitizeName(rawName)

			// Track local duplicates
			localUsed[baseName]++
			isLocalDup := localUsed[baseName] > 1

			candidate := enum.Name + "_" + baseName

			// Resolve conflicts
			if g.usedNames[candidate] || isLocalDup {
				candidate = g.resolveConflict(candidate)
			}

			g.usedNames[candidate] = true

			// Add value description
			if value.Description != "" {
				sb.WriteString(fmt.Sprintf("\t// %s - %s\n", rawName, value.Description))
			}

			sb.WriteString(fmt.Sprintf("\t%s %s = %d\n", candidate, enum.Name, value.Value))
		}

		sb.WriteString(")\n")
		return nil
	}

	var processPrototype func(proto *manifest.Prototype) error
	processPrototype = func(proto *manifest.Prototype) error {
		// Check return type
		if proto.RetType.Type != "void" && proto.RetType.Enum != nil {
			if err := processEnum(proto.RetType.Enum, proto.RetType.Type); err != nil {
				return err
			}
		}
		// Check parameters
		for _, param := range proto.ParamTypes {
			if param.Enum != nil {
				if err := processEnum(param.Enum, param.Type); err != nil {
					return err
				}
			}
			if param.Prototype != nil {
				if err := processPrototype(param.Prototype); err != nil {
					return err
				}
			}
		}
		return nil
	}

	// Process all methods
	for _, method := range m.Methods {
		// Check return type
		if method.RetType.Type != "void" && method.RetType.Enum != nil {
			if err := processEnum(method.RetType.Enum, method.RetType.Type); err != nil {
				return "", err
			}
		}
		// Check return type prototype
		if method.RetType.Prototype != nil {
			if err := processPrototype(method.RetType.Prototype); err != nil {
				return "", err
			}
		}
		// Check parameters
		for _, param := range method.ParamTypes {
			if param.Enum != nil {
				if err := processEnum(param.Enum, param.Type); err != nil {
					return "", err
				}
			}
			if param.Prototype != nil {
				if err := processPrototype(param.Prototype); err != nil {
					return "", err
				}
			}
		}
	}

	return sb.String(), nil
}

// resolveConflict generates a unique name by appending a suffix
func (g *GolangGenerator) resolveConflict(name string) string {
	if !g.usedNames[name] {
		return name
	}

	suffix := 2
	for {
		trial := fmt.Sprintf("%s_%d", name, suffix)
		if !g.usedNames[trial] {
			return trial
		}
		suffix++
	}
}

// generateDelegates generates delegate (function type) definitions
func (g *GolangGenerator) generateDelegates(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	var processPrototype func(proto *manifest.Prototype) error
	processPrototype = func(proto *manifest.Prototype) error {
		if g.IsDelegateCached(proto.Name) {
			return nil
		}
		g.CacheDelegate(proto.Name)

		// Add delegate description
		if proto.Description != "" {
			sb.WriteString(fmt.Sprintf("// %s - %s\n", proto.Name, proto.Description))
		}

		// Generate parameter list
		params, err := g.formatParams(proto.ParamTypes, true)
		if err != nil {
			return err
		}

		// Generate return type
		returnType := ""
		if proto.RetType.Type != "void" {
			returnType, err = g.typeMapper.MapReturnType(&proto.RetType)
			if err != nil {
				return err
			}
		}

		// Generate function type
		sb.WriteString(fmt.Sprintf("type %s func(%s)", proto.Name, params))
		if returnType != "" {
			sb.WriteString(fmt.Sprintf(" %s", returnType))
		}
		sb.WriteString("\n\n")

		return nil
	}

	// Process all methods
	for _, method := range m.Methods {
		// Check return type
		if method.RetType.Type != "void" && method.RetType.Prototype != nil {
			if err := processPrototype(method.RetType.Prototype); err != nil {
				return "", err
			}
		}
		// Check parameters
		for _, param := range method.ParamTypes {
			if param.Prototype != nil {
				if err := processPrototype(param.Prototype); err != nil {
					return "", err
				}
			}
		}
	}

	return sb.String(), nil
}

// generateMethod generates a single method implementation
func (g *GolangGenerator) generateMethod(method *manifest.Method) (string, error) {
	var sb strings.Builder

	// Generate documentation
	if method.Description != "" {
		sb.WriteString(fmt.Sprintf("// %s \n//  @brief %s\n//\n", method.Name, method.Description))
	}

	// Add parameter documentation
	for _, param := range method.ParamTypes {
		if param.Description != "" {
			sb.WriteString(fmt.Sprintf("//  @param %s: %s\n", param.Name, param.Description))
		}
	}

	// Add return type documentation
	if method.RetType.Type != "void" && method.RetType.Description != "" {
		sb.WriteString(fmt.Sprintf("//\n//  @return %s\n", method.RetType.Description))
	}

	// Generate function signature
	params, err := g.formatParams(method.ParamTypes, true)
	if err != nil {
		return "", err
	}

	returnType := ""
	if method.RetType.Type != "void" {
		returnType, err = g.typeMapper.MapReturnType(&method.RetType)
		if err != nil {
			return "", err
		}
	}

	sb.WriteString(fmt.Sprintf("func %s(%s)", method.Name, params))
	if returnType != "" {
		sb.WriteString(fmt.Sprintf(" %s", returnType))
	}
	sb.WriteString(" {\n")

	// Generate method body
	body, err := g.generateMethodBody(method)
	if err != nil {
		return "", err
	}
	sb.WriteString(body)

	sb.WriteString("}\n")

	return sb.String(), nil
}

// generateMethodBody generates the method body with marshaling
func (g *GolangGenerator) generateMethodBody(method *manifest.Method) (string, error) {
	var sb strings.Builder
	indent := "\t"
	innerIndent := indent

	isObjRet := method.RetType.Type != "void" && g.typeMapper.IsObjType(method.RetType.Type)
	isPodRet := method.RetType.Type != "void" && g.typeMapper.IsPodType(method.RetType.Type)
	hasRet := method.RetType.Type != "void" && !isObjRet

	// Generate parameter casting
	paramsCast, err := g.generateParamsCast(method, indent)
	if err != nil {
		return "", err
	}

	hasCast := len(paramsCast) > 0 && method.RetType.Type != "void"

	// Declare return value if needed
	if isObjRet || hasCast {
		returnType, _ := g.typeMapper.MapReturnType(&method.RetType)
		sb.WriteString(fmt.Sprintf("%svar __retVal %s\n", indent, returnType))
	}

	// Add parameter casting code
	hasTry := false
	insertIndexStart := 0
	insertIndexEnd := 0
	if len(paramsCast) > 0 {
		for _, cast := range paramsCast {
			sb.WriteString(cast)
		}
		insertIndexStart = sb.Len()
		sb.WriteString(fmt.Sprintf("%splugify.Block {\n", indent))
		sb.WriteString(fmt.Sprintf("%s\tTry: func() {\n", indent))
		insertIndexEnd = sb.Len()
		innerIndent = "\t\t\t"
		hasTry = hasCast
	}

	// Generate function call
	functionCall, err := g.generateFunctionCall(method)
	if err != nil {
		return "", err
	}

	if isObjRet {
		retTypeCast := g.typeMapper.retTypeCastMap[method.RetType.Type]
		sb.WriteString(fmt.Sprintf("%s__native := %s\n", innerIndent, functionCall))
		sb.WriteString(fmt.Sprintf("%s__retVal_native = *(*%s)(unsafe.Pointer(&__native))\n", innerIndent, retTypeCast))
	} else if hasTry {
		if isPodRet {
			retTypeCast := g.typeMapper.retTypeCastMap[method.RetType.Type]
			sb.WriteString(fmt.Sprintf("%s__native := %s\n", innerIndent, functionCall))
			sb.WriteString(fmt.Sprintf("%s__retVal = *(*%s)(unsafe.Pointer(&__native))\n", innerIndent, retTypeCast))
		} else {
			retTypeCast := g.typeMapper.retTypeCastMap[method.RetType.Type]
			if retTypeCast != "" {
				sb.WriteString(fmt.Sprintf("%s__retVal = %s(%s)\n", innerIndent, retTypeCast, functionCall))
			} else {
				sb.WriteString(fmt.Sprintf("%s__retVal = %s\n", innerIndent, functionCall))
			}
		}
	} else if hasRet {
		if isPodRet {
			retTypeCast := g.typeMapper.retTypeCastMap[method.RetType.Type]
			sb.WriteString(fmt.Sprintf("%s__native := %s\n", innerIndent, functionCall))
			sb.WriteString(fmt.Sprintf("%s__retVal := *(*%s)(unsafe.Pointer(&__native))\n", innerIndent, retTypeCast))
		} else if method.RetType.Type == "function" {
			returnType, _ := g.typeMapper.MapReturnType(&method.RetType)
			sb.WriteString(fmt.Sprintf("%s__retVal := plugify.GetDelegateForFunctionPointer(%s, reflect.TypeOf(%s(nil))).(%s)\n",
				innerIndent, functionCall, returnType, returnType))
		} else {
			returnType, _ := g.typeMapper.MapReturnType(&method.RetType)
			sb.WriteString(fmt.Sprintf("%s__retVal := %s(%s)\n", innerIndent, returnType, functionCall))
		}
	} else {
		sb.WriteString(fmt.Sprintf("%s%s\n", innerIndent, functionCall))
	}

	// Generate assignment casting (unmarshal)
	assignCast, err := g.generateParamsCastAssign(method, innerIndent)
	if err != nil {
		return "", err
	}

	if len(assignCast) > 0 {
		sb.WriteString(fmt.Sprintf("%s// Unmarshal - Convert native data to managed data.\n", innerIndent))
		for _, assign := range assignCast {
			sb.WriteString(assign)
		}
	}

	// Generate cleanup code
	cleanupCast, err := g.generateParamsCastCleanup(method, innerIndent)
	if err != nil {
		return "", err
	}

	if len(cleanupCast) > 0 {
		sb.WriteString(fmt.Sprintf("%s\t},\n%s\tFinally: func() {\n%s// Perform cleanup.\n", indent, indent, innerIndent))
		for _, cleanup := range cleanupCast {
			sb.WriteString(cleanup)
		}
		sb.WriteString(fmt.Sprintf("%s\t},\n%s}.Do()\n", indent, indent))
	} else if len(paramsCast) > 0 {
		RemoveFromBuilder(&sb, insertIndexStart, insertIndexEnd)
		RemoveLeadingTabs(&sb, 2, insertIndexStart, sb.Len())
	}

	// Return value
	if method.RetType.Type != "void" {
		sb.WriteString(fmt.Sprintf("%sreturn __retVal\n", indent))
	}

	return sb.String(), nil
}

// generateParamsCast generates parameter casting code
func (g *GolangGenerator) generateParamsCast(method *manifest.Method, indent string) ([]string, error) {
	var result []string

	// Handle return type
	if method.RetType.Type != "void" && g.typeMapper.IsObjType(method.RetType.Type) {
		retTypeCast := g.typeMapper.retTypeCastMap[method.RetType.Type]
		if retTypeCast != "" {
			result = append(result, fmt.Sprintf("%svar __retVal_native %s\n", indent, retTypeCast))
		}
	}

	// Handle parameters
	for _, param := range method.ParamTypes {
		cast, err := g.generateParamCast(&param, indent)
		if err != nil {
			return nil, err
		}
		if cast != "" {
			result = append(result, cast)
		}
	}

	return result, nil
}

// generateParamCast generates casting code for a single parameter
func (g *GolangGenerator) generateParamCast(param *manifest.ParamType, indent string) (string, error) {
	paramType := g.typeMapper.valTypeCastMap[param.Type]
	name := g.SanitizeName(param.Name)

	if paramType == "" {
		return "", nil
	}

	// Handle vector/matrix types
	if strings.HasPrefix(paramType, "C.Vector") || strings.HasPrefix(paramType, "C.Matrix") {
		if param.Ref {
			return fmt.Sprintf("%s__%s := *(*%s)(unsafe.Pointer(%s))\n", indent, name, paramType, name), nil
		}
		return fmt.Sprintf("%s__%s := *(*%s)(unsafe.Pointer(&%s))\n", indent, name, paramType, name), nil
	}

	// Handle other types
	if param.Ref {
		return fmt.Sprintf("%s__%s := %s(*%s)\n", indent, name, paramType, name), nil
	}
	return fmt.Sprintf("%s__%s := %s(%s)\n", indent, name, paramType, name), nil
}

// generateParamsCastAssign generates assignment casting code (unmarshal)
func (g *GolangGenerator) generateParamsCastAssign(method *manifest.Method, indent string) ([]string, error) {
	var result []string

	// Handle return type
	if method.RetType.Type != "void" && g.typeMapper.IsObjType(method.RetType.Type) {
		assign, err := g.generateReturnAssign(&method.RetType, indent)
		if err != nil {
			return nil, err
		}
		if assign != "" {
			result = append(result, assign)
		}
	}

	// Handle parameters
	for _, param := range method.ParamTypes {
		if param.Ref {
			assign, err := g.generateParamAssign(&param, indent)
			if err != nil {
				return nil, err
			}
			if assign != "" {
				result = append(result, assign)
			}
		}
	}

	return result, nil
}

// generateParamAssign generates assignment code for a ref parameter
func (g *GolangGenerator) generateParamAssign(param *manifest.ParamType, indent string) (string, error) {
	paramType := g.typeMapper.assTypeCastMap[param.Type]
	name := g.SanitizeName(param.Name)

	if paramType == "" {
		return "", nil
	}

	// Handle vector/matrix types
	if strings.HasPrefix(paramType, "plugify.Vector") || strings.HasPrefix(paramType, "plugify.Matrix") {
		return fmt.Sprintf("%s*%s = *(*%s)(unsafe.Pointer(&__%s))\n", indent, name, paramType, name), nil
	}

	// Handle VectorData types
	if strings.Contains(paramType, "VectorData") {
		return fmt.Sprintf("%s%sTo(&__%s, %s)\n", indent, paramType, name, name), nil
	}

	// Handle other types
	if strings.HasPrefix(paramType, "plugify.") {
		return fmt.Sprintf("%s*%s = %s(&__%s)\n", indent, name, paramType, name), nil
	}

	if param.Enum != nil {
		return fmt.Sprintf("%s*%s = %s(__%s)\n", indent, name, param.Enum.Name, name), nil
	}

	return fmt.Sprintf("%s*%s = %s(__%s)\n", indent, name, paramType, name), nil
}

// generateReturnAssign generates assignment code for return value
func (g *GolangGenerator) generateReturnAssign(retType *manifest.TypeInfo, indent string) (string, error) {
	paramType := g.typeMapper.assTypeCastMap[retType.Type]

	if paramType == "" {
		return "", nil
	}

	if retType.Enum != nil {
		return fmt.Sprintf("%s__retVal = %sT[%s](&__retVal_native)\n", indent, paramType, retType.Enum.Name), nil
	}

	return fmt.Sprintf("%s__retVal = %s(&__retVal_native)\n", indent, paramType), nil
}

// generateParamsCastCleanup generates cleanup code
func (g *GolangGenerator) generateParamsCastCleanup(method *manifest.Method, indent string) ([]string, error) {
	var result []string

	// Handle return type
	if method.RetType.Type != "void" && g.typeMapper.IsObjType(method.RetType.Type) {
		cleanup := g.generateReturnCleanup(&method.RetType, indent)
		if cleanup != "" {
			result = append(result, cleanup)
		}
	}

	// Handle parameters
	for _, param := range method.ParamTypes {
		cleanup := g.generateParamCleanup(&param)
		if cleanup != "" {
			result = append(result, indent+cleanup)
		}
	}

	return result, nil
}

// generateParamCleanup generates cleanup code for a parameter
func (g *GolangGenerator) generateParamCleanup(param *manifest.ParamType) string {
	paramType := g.typeMapper.delTypeCastMap[param.Type]
	if paramType == "" {
		return ""
	}

	name := g.SanitizeName(param.Name)
	return fmt.Sprintf("%s(&__%s)\n", paramType, name)
}

// generateReturnCleanup generates cleanup code for return value
func (g *GolangGenerator) generateReturnCleanup(retType *manifest.TypeInfo, indent string) string {
	returnType := g.typeMapper.delTypeCastMap[retType.Type]
	if returnType == "" {
		return ""
	}

	return fmt.Sprintf("%s%s(&__retVal_native)\n", indent, returnType)
}

// generateFunctionCall generates the C function call
func (g *GolangGenerator) generateFunctionCall(method *manifest.Method) (string, error) {
	params, err := g.formatCastParams(method.ParamTypes)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("C.%s(%s)", method.Name, params), nil
}

// formatParams formats parameters with types and names
func (g *GolangGenerator) formatParams(params []manifest.ParamType, withTypes bool) (string, error) {
	if len(params) == 0 {
		return "", nil
	}

	var parts []string
	for i, param := range params {
		name := param.Name
		if name == "" {
			name = fmt.Sprintf("p%d", i)
		}
		name = g.SanitizeName(name)

		if withTypes {
			typeName, err := g.typeMapper.MapParamType(&param, TypeContextValue)
			if err != nil {
				return "", err
			}
			parts = append(parts, fmt.Sprintf("%s %s", name, typeName))
		} else {
			parts = append(parts, name)
		}
	}

	return strings.Join(parts, ", "), nil
}

// formatCastParams formats cast parameter names for C function call
func (g *GolangGenerator) formatCastParams(params []manifest.ParamType) (string, error) {
	if len(params) == 0 {
		return "", nil
	}

	var parts []string
	for _, param := range params {
		name := g.SanitizeName(param.Name)

		if g.typeMapper.IsObjType(param.Type) {
			ctype := g.typeMapper.ctypesMap[param.Type]
			parts = append(parts, fmt.Sprintf("(*%s)(unsafe.Pointer(&__%s))", ctype, name))
		} else if g.typeMapper.IsPodType(param.Type) || (param.Type[:3] == "vec" || param.Type[:3] == "mat") {
			parts = append(parts, fmt.Sprintf("&__%s", name))
		} else if param.Ref {
			parts = append(parts, fmt.Sprintf("&__%s", name))
		} else {
			parts = append(parts, fmt.Sprintf("__%s", name))
		}
	}

	return strings.Join(parts, ", "), nil
}

// generateHMethod generates a single C method in the .h file
func (g *GolangGenerator) generateHMethod(method *manifest.Method, pluginName string) (string, error) {
	var sb strings.Builder

	// Get return type
	retType := "void"
	if method.RetType.Type != "void" {
		retType = g.typeMapper.GetCType(method.RetType.Type, false, true)
	}

	// Get parameters
	paramList, err := g.formatCParams(method.ParamTypes, true)
	if err != nil {
		return "", err
	}

	paramTypes, err := g.formatCParams(method.ParamTypes, false)
	if err != nil {
		return "", err
	}

	paramNames, err := g.formatCParamNames(method.ParamTypes)
	if err != nil {
		return "", err
	}

	// Generate function
	sb.WriteString(fmt.Sprintf("static %s %s(%s) {\n", retType, method.Name, paramList))
	sb.WriteString(fmt.Sprintf("\ttypedef %s (*%sFn)(%s);\n", retType, method.Name, paramTypes))
	sb.WriteString(fmt.Sprintf("\tstatic %sFn __func = NULL;\n", method.Name))
	sb.WriteString(fmt.Sprintf("\tif (__func == NULL) Plugify_GetMethodPtr2(\"%s.%s\", (void**)&__func);\n",
		pluginName, method.Name))

	if method.RetType.Type != "void" {
		sb.WriteString(fmt.Sprintf("\treturn __func(%s);\n", paramNames))
	} else {
		sb.WriteString(fmt.Sprintf("\t__func(%s);\n", paramNames))
	}

	sb.WriteString("}\n")

	return sb.String(), nil
}

// formatCParams formats C parameters
func (g *GolangGenerator) formatCParams(params []manifest.ParamType, withNames bool) (string, error) {
	if len(params) == 0 {
		return "", nil
	}

	var parts []string
	for i, param := range params {
		typeName := g.typeMapper.GetCType(param.Type, param.Ref, false)

		if withNames {
			name := param.Name
			if name == "" {
				name = fmt.Sprintf("p%d", i)
			}
			parts = append(parts, fmt.Sprintf("%s %s", typeName, name))
		} else {
			parts = append(parts, typeName)
		}
	}

	return strings.Join(parts, ", "), nil
}

// formatCParamNames formats C parameter names only
func (g *GolangGenerator) formatCParamNames(params []manifest.ParamType) (string, error) {
	if len(params) == 0 {
		return "", nil
	}

	var parts []string
	for i, param := range params {
		name := param.Name
		if name == "" {
			name = fmt.Sprintf("p%d", i)
		}
		parts = append(parts, name)
	}

	return strings.Join(parts, ", "), nil
}

func (g *GolangGenerator) generateClasses(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	// Generate shared types if any class has a destructor
	hasDestructor := false
	for _, class := range m.Classes {
		if class.Destructor != nil {
			hasDestructor = true
			break
		}
	}

	if hasDestructor {
		sb.WriteString("// noCopy prevents copying via go vet\n")
		sb.WriteString("type noCopy struct{}\n\n")
		sb.WriteString("func (*noCopy) Lock()   {}\n")
		sb.WriteString("func (*noCopy) Unlock() {}\n\n")
		sb.WriteString("// ownership indicates whether the instance owns the underlying handle\n")
		sb.WriteString("type ownership bool\n\n")
		sb.WriteString("const (\n")
		sb.WriteString("\tOwned    ownership = true\n")
		sb.WriteString("\tBorrowed ownership = false\n")
		sb.WriteString(")\n\n")
	}

	// Generate each class
	for _, class := range m.Classes {
		classCode, err := g.generateClass(m, &class)
		if err != nil {
			return "", fmt.Errorf("failed to generate class %s: %w", class.Name, err)
		}
		sb.WriteString(classCode)
		sb.WriteString("\n")
	}

	return sb.String(), nil
}

func (g *GolangGenerator) generateClass(m *manifest.Manifest, class *manifest.Class) (string, error) {
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
	_, handleType, err := g.typeMapper.MapHandleType(class)
	if err != nil {
		return "", err
	}

	// Only generate error variable if class has a handle
	if hasHandle {
		// Generate error variable for this class
		errVarName := fmt.Sprintf("%sErrEmptyHandle", class.Name)
		sb.WriteString(fmt.Sprintf("var (\n\t%s = errors.New(\"%s: empty handle\")\n)\n\n", errVarName, class.Name))
	}

	// Class documentation
	if class.Description != "" {
		sb.WriteString(fmt.Sprintf("// %s - %s\n", class.Name, class.Description))
	}

	// Struct definition
	sb.WriteString(fmt.Sprintf("type %s struct {\n", class.Name))
	if hasHandle {
		sb.WriteString(fmt.Sprintf("\thandle    %s\n", handleType))
		if hasDtor {
			sb.WriteString("\tcleanup   runtime.Cleanup\n")
			sb.WriteString("\townership ownership\n")
			sb.WriteString("\tnoCopy    noCopy\n")
		}
	}
	sb.WriteString("}\n\n")

	// Only generate handle-related code if class has a handle
	if hasHandle {
		// Generate constructors
		for _, ctorName := range class.Constructors {
			ctorCode, err := g.generateConstructor(m, class, ctorName)
			if err != nil {
				return "", err
			}
			sb.WriteString(ctorCode)
		}

		// Generate helper constructors (borrowed/owned)
		if hasDtor {
			helperCode, err := g.generateHelperConstructors(class)
			if err != nil {
				return "", err
			}
			sb.WriteString(helperCode)
		} else {
			ctorCode, err := g.generateDefaultConstructor(class)
			if err != nil {
				return "", err
			}
			sb.WriteString(ctorCode)
		}

		// Generate finalizer
		if hasDtor {
			finalizerCode, err := g.generateFinalizer(class)
			if err != nil {
				return "", err
			}
			sb.WriteString(finalizerCode)
		}

		// Generate utility methods
		utilityCode, err := g.generateUtilityMethods(class)
		if err != nil {
			return "", err
		}
		sb.WriteString(utilityCode)
	}

	// Generate class bindings
	errVarName := fmt.Sprintf("%sErrEmptyHandle", class.Name)
	for _, binding := range class.Bindings {
		methodCode, err := g.generateBinding(m, class, &binding, errVarName)
		if err != nil {
			return "", err
		}
		sb.WriteString(methodCode)
	}

	return sb.String(), nil
}

func (g *GolangGenerator) generateConstructor(m *manifest.Manifest, class *manifest.Class, methodName string) (string, error) {
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
	if method.Description != "" {
		sb.WriteString(fmt.Sprintf("// New%s%s - %s\n", class.Name, method.Name, method.Description))
	}

	// Generate parameters documentation
	for _, param := range method.ParamTypes {
		if param.Description != "" {
			sb.WriteString(fmt.Sprintf("//  @param %s: %s\n", param.Name, param.Description))
		}
	}

	// Generate function signature
	params, err := g.formatParams(method.ParamTypes, true)
	if err != nil {
		return "", err
	}

	funcName := fmt.Sprintf("New%s%s", class.Name, method.Name)
	sb.WriteString(fmt.Sprintf("func %s(%s) *%s {\n", funcName, params, class.Name))

	// Generate call to underlying C function
	callParams, err := g.formatParams(method.ParamTypes, false)
	if err != nil {
		return "", err
	}

	if class.Destructor != nil {
		sb.WriteString(fmt.Sprintf("\treturn New%sOwned(%s(%s))\n", class.Name, method.Name, callParams))
	} else {
		sb.WriteString(fmt.Sprintf("\treturn &%s{\n", class.Name))
		sb.WriteString(fmt.Sprintf("\t\thandle: %s(%s),\n", method.Name, callParams))
		sb.WriteString("\t}\n")
	}

	sb.WriteString("}\n\n")

	return sb.String(), nil
}

func (g *GolangGenerator) generateHelperConstructors(class *manifest.Class) (string, error) {
	var sb strings.Builder

	invalidValue, handleType, err := g.typeMapper.MapHandleType(class)
	if err != nil {
		return "", err
	}

	// newBorrowed helper
	sb.WriteString(fmt.Sprintf("// New%sBorrowed creates a %s from a borrowed handle\n", class.Name, class.Name))
	sb.WriteString(fmt.Sprintf("func New%sBorrowed(handle %s) *%s {\n", class.Name, handleType, class.Name))
	sb.WriteString(fmt.Sprintf("\tif handle == %s {\n", invalidValue))
	sb.WriteString(fmt.Sprintf("\t\treturn &%s{}\n", class.Name))
	sb.WriteString("\t}\n")
	sb.WriteString(fmt.Sprintf("\treturn &%s{\n", class.Name))
	sb.WriteString("\t\thandle:    handle,\n")
	sb.WriteString("\t\townership: Borrowed,\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n\n")

	// newOwned helper
	sb.WriteString(fmt.Sprintf("// New%sOwned creates a %s from an owned handle\n", class.Name, class.Name))
	sb.WriteString(fmt.Sprintf("func New%sOwned(handle %s) *%s {\n", class.Name, handleType, class.Name))
	sb.WriteString(fmt.Sprintf("\tif handle == %s {\n", invalidValue))
	sb.WriteString(fmt.Sprintf("\t\treturn &%s{}\n", class.Name))
	sb.WriteString("\t}\n")
	sb.WriteString(fmt.Sprintf("\tw := &%s{\n", class.Name))
	sb.WriteString("\t\thandle:    handle,\n")
	sb.WriteString("\t\townership: Owned,\n")
	sb.WriteString("\t}\n")
	sb.WriteString("\tw.cleanup = runtime.AddCleanup(w, w.finalize, struct{}{})\n")
	sb.WriteString("\treturn w\n")
	sb.WriteString("}\n\n")

	return sb.String(), nil
}

func (g *GolangGenerator) generateDefaultConstructor(class *manifest.Class) (string, error) {
	var sb strings.Builder

	_, handleType, err := g.typeMapper.MapHandleType(class)
	if err != nil {
		return "", err
	}

	// New helper
	sb.WriteString(fmt.Sprintf("// New%s creates a %s from a handle\n", class.Name, class.Name))
	sb.WriteString(fmt.Sprintf("func New%s(handle %s) *%s {\n", class.Name, handleType, class.Name))
	sb.WriteString(fmt.Sprintf("\treturn &%s{\n", class.Name))
	sb.WriteString("\t\thandle:    handle,\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n\n")

	return sb.String(), nil
}

func (g *GolangGenerator) generateFinalizer(class *manifest.Class) (string, error) {
	var sb strings.Builder

	invalidValue, _, err := g.typeMapper.MapHandleType(class)
	if err != nil {
		return "", err
	}

	// finalize function
	sb.WriteString("// finalize is the finalizer function (like C++ destructor)\n")
	sb.WriteString(fmt.Sprintf("func (w *%s) finalize(_ struct{}) {\n", class.Name))
	sb.WriteString("\tif plugify.Plugin.Loaded {\n")
	sb.WriteString("\t\tw.destroy()\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n\n")

	// destroy function
	sb.WriteString("// destroy cleans up owned handles\n")
	sb.WriteString(fmt.Sprintf("func (w *%s) destroy() {\n", class.Name))
	sb.WriteString(fmt.Sprintf("\tif w.handle != %s && w.ownership == Owned {\n", invalidValue))
	sb.WriteString(fmt.Sprintf("\t\t%s(w.handle)\n", *class.Destructor))
	sb.WriteString("\t}\n")
	sb.WriteString("}\n\n")

	// nullify function
	sb.WriteString("// nullify resets the handle\n")
	sb.WriteString(fmt.Sprintf("func (w *%s) nullify() {\n", class.Name))
	sb.WriteString(fmt.Sprintf("\tw.handle = %s\n", invalidValue))
	sb.WriteString("\tw.ownership = Borrowed\n")
	sb.WriteString("}\n\n")

	return sb.String(), nil
}

func (g *GolangGenerator) generateUtilityMethods(class *manifest.Class) (string, error) {
	var sb strings.Builder

	invalidValue, handleType, err := g.typeMapper.MapHandleType(class)
	if err != nil {
		return "", err
	}

	hasDtor := class.Destructor != nil

	// Close method (only if destructor exists)
	if hasDtor {
		sb.WriteString("// Close explicitly destroys the handle (like C++ destructor, but manual)\n")
		sb.WriteString(fmt.Sprintf("func (w *%s) Close() {\n", class.Name))
		sb.WriteString("\tw.Reset()\n")
		sb.WriteString("}\n\n")
	}

	// Get method
	sb.WriteString("// Get returns the underlying handle\n")
	sb.WriteString(fmt.Sprintf("func (w *%s) Get() %s {\n", class.Name, handleType))
	sb.WriteString("\treturn w.handle\n")
	sb.WriteString("}\n\n")

	// Release method
	sb.WriteString("// Release releases ownership and returns the handle\n")
	sb.WriteString(fmt.Sprintf("func (w *%s) Release() %s {\n", class.Name, handleType))
	if hasDtor {
		sb.WriteString("\tif w.ownership == Owned {\n")
		sb.WriteString("\t\tw.cleanup.Stop()\n")
		sb.WriteString("\t}\n")
	}
	sb.WriteString("\thandle := w.handle\n")
	if hasDtor {
		sb.WriteString("\tw.nullify()\n")
	} else {
		sb.WriteString(fmt.Sprintf("\tw.handle = %s\n", invalidValue))
	}
	sb.WriteString("\treturn handle\n")
	sb.WriteString("}\n\n")

	// Reset method
	sb.WriteString("// Reset destroys and resets the handle\n")
	sb.WriteString(fmt.Sprintf("func (w *%s) Reset() {\n", class.Name))
	if hasDtor {
		sb.WriteString("\tif w.ownership == Owned {\n")
		sb.WriteString("\t\tw.cleanup.Stop()\n")
		sb.WriteString("\t}\n")
		sb.WriteString("\tw.destroy()\n")
		sb.WriteString("\tw.nullify()\n")
	} else {
		sb.WriteString(fmt.Sprintf("\tw.handle = %s\n", invalidValue))
	}
	sb.WriteString("}\n\n")

	// IsValid method
	sb.WriteString("// IsValid returns true if handle is not nil\n")
	sb.WriteString(fmt.Sprintf("func (w *%s) IsValid() bool {\n", class.Name))
	sb.WriteString(fmt.Sprintf("\treturn w.handle != %s\n", invalidValue))
	sb.WriteString("}\n\n")

	return sb.String(), nil
}

func (g *GolangGenerator) generateBinding(m *manifest.Manifest, class *manifest.Class, binding *manifest.Binding, errVarName string) (string, error) {
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
	if method.Description != "" {
		sb.WriteString(fmt.Sprintf("// %s - %s\n", binding.Name, method.Description))
	}

	// Document parameters
	for _, param := range methodParams {
		if param.Description != "" {
			sb.WriteString(fmt.Sprintf("//  @param %s: %s\n", param.Name, param.Description))
		}
	}

	// Document return type
	if method.RetType.Type != "void" && method.RetType.Description != "" {
		sb.WriteString(fmt.Sprintf("//  @return %s\n", method.RetType.Description))
	}

	// Generate method signature
	formattedParams, err := g.formatClassParams(methodParams, binding.ParamAliases, true)
	if err != nil {
		return "", err
	}

	// Build return type
	returnSignature := ""

	// Determine if method is static
	if !binding.BindSelf {
		if method.RetType.Type != "void" {
			if binding.RetAlias != nil && binding.RetAlias.Name != "" {
				returnSignature = fmt.Sprintf(" *%s", binding.RetAlias.Name)
			} else {
				retType, err := g.typeMapper.MapReturnType(&method.RetType)
				if err != nil {
					return "", err
				}
				returnSignature = fmt.Sprintf(" %s", retType)
			}
		}
	} else {
		if method.RetType.Type != "void" {
			if binding.RetAlias != nil && binding.RetAlias.Name != "" {
				returnSignature = fmt.Sprintf(" (*%s, error)", binding.RetAlias.Name)
			} else {
				retType, err := g.typeMapper.MapReturnType(&method.RetType)
				if err != nil {
					return "", err
				}
				returnSignature = fmt.Sprintf(" (%s, error)", retType)
			}
		} else {
			returnSignature = " error"
		}
	}

	sb.WriteString(fmt.Sprintf("func (w *%s) %s(%s)%s {\n", class.Name, binding.Name, formattedParams, returnSignature))

	// Generate null check
	invalidValue, _, err := g.typeMapper.MapHandleType(class)
	if err != nil {
		return "", err
	}

	nullPolicy := class.NullPolicy
	if nullPolicy == "" {
		nullPolicy = "throw"
	}

	if binding.BindSelf {
		if nullPolicy == "throw" {
			sb.WriteString(fmt.Sprintf("\tif w.handle == %s {\n", invalidValue))
			if method.RetType.Type != "void" {
				sb.WriteString("\t\tvar zero ")
				if binding.RetAlias != nil && binding.RetAlias.Name != "" {
					sb.WriteString(fmt.Sprintf("*%s\n", binding.RetAlias.Name))
				} else {
					retType, _ := g.typeMapper.MapReturnType(&method.RetType)
					sb.WriteString(fmt.Sprintf("%s\n", retType))
				}
				sb.WriteString(fmt.Sprintf("\t\treturn zero, %s\n", errVarName))
			} else {
				sb.WriteString(fmt.Sprintf("\t\treturn %s\n", errVarName))
			}
			sb.WriteString("\t}\n")
		}
	}

	// Build call arguments
	callArgs, err := g.formatClassCallArgs(methodParams, binding)
	if err != nil {
		return "", err
	}

	hasCtor := len(class.Constructors) > 0
	hasDtor := class.Destructor != nil

	// Generate call
	if method.RetType.Type == "void" {
		sb.WriteString(fmt.Sprintf("\t%s(%s)\n", method.FuncName, callArgs))
		if binding.BindSelf {
			sb.WriteString("\treturn nil\n")
		}
	} else {
		errorTag := ""
		if binding.BindSelf {
			errorTag = ", nil"
		}
		if binding.RetAlias != nil && binding.RetAlias.Name != "" {
			ownership := ""
			if hasDtor || hasCtor {
				if binding.RetAlias.Owner {
					ownership = fmt.Sprintf("New%sOwned", binding.RetAlias.Name)
				} else {
					ownership = fmt.Sprintf("New%sBorrowed", binding.RetAlias.Name)
				}
			} else {
				ownership = fmt.Sprintf("New%s", binding.RetAlias.Name)
			}
			sb.WriteString(fmt.Sprintf("\treturn %s(%s(%s))%s\n", ownership, method.FuncName, callArgs, errorTag))
		} else {
			sb.WriteString(fmt.Sprintf("\treturn %s(%s)%s\n", method.FuncName, callArgs, errorTag))
		}
	}

	sb.WriteString("}\n\n")

	return sb.String(), nil
}

func (g *GolangGenerator) formatClassParams(params []manifest.ParamType, aliases []*manifest.ParamAlias, withTypes bool) (string, error) {
	if len(params) == 0 {
		return "", nil
	}

	var parts []string
	for i, param := range params {
		name := g.SanitizeName(param.Name)

		if withTypes {
			typeName := ""
			var err error

			// Check if this parameter has an alias
			if i < len(aliases) && aliases[i] != nil {
				typeName = fmt.Sprintf("*%s", aliases[i].Name)
			} else {
				typeName, err = g.typeMapper.MapParamType(&param, TypeContextValue)
				if err != nil {
					return "", err
				}
			}
			parts = append(parts, fmt.Sprintf("%s %s", name, typeName))
		} else {
			parts = append(parts, name)
		}
	}

	return strings.Join(parts, ", "), nil
}

func (g *GolangGenerator) formatClassCallArgs(params []manifest.ParamType, binding *manifest.Binding) (string, error) {
	var parts []string

	// Add self if bindSelf
	if binding.BindSelf {
		parts = append(parts, "w.handle")
	}

	// Add other parameters
	for i, param := range params {
		name := g.SanitizeName(param.Name)

		// Check if parameter has alias
		if i < len(binding.ParamAliases) && binding.ParamAliases[i] != nil {
			if binding.ParamAliases[i].Owner {
				parts = append(parts, fmt.Sprintf("%s.Release()", name))
			} else {
				parts = append(parts, fmt.Sprintf("%s.Get()", name))
			}
		} else {
			parts = append(parts, name)
		}
	}

	return strings.Join(parts, ", "), nil
}

// generateEnumsGoFile generates a file containing all enums
func (g *GolangGenerator) generateEnumsGoFile(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	// Package declaration
	sb.WriteString(fmt.Sprintf("package %s\n\n", m.Name))

	// Add comment header
	sb.WriteString(fmt.Sprintf("// Generated from %s\n\n", m.Name))

	// Generate enums
	enumsCode, err := g.generateEnums(m)
	if err != nil {
		return "", err
	}
	if enumsCode != "" {
		sb.WriteString(enumsCode)
		sb.WriteString("\n")
	}

	// Generate ownership types if any class has a destructor
	hasDestructor := false
	for _, class := range m.Classes {
		if class.Destructor != nil {
			hasDestructor = true
			break
		}
	}

	if hasDestructor {
		sb.WriteString("// noCopy prevents copying via go vet\n")
		sb.WriteString("type noCopy struct{}\n\n")
		sb.WriteString("func (*noCopy) Lock()   {}\n")
		sb.WriteString("func (*noCopy) Unlock() {}\n\n")
		sb.WriteString("// ownership indicates whether the instance owns the underlying handle\n")
		sb.WriteString("type ownership bool\n\n")
		sb.WriteString("const (\n")
		sb.WriteString("\tOwned    ownership = true\n")
		sb.WriteString("\tBorrowed ownership = false\n")
		sb.WriteString(")\n\n")
	}

	return sb.String(), nil
}

// generateDelegatesGoFile generates a file containing all delegate types
func (g *GolangGenerator) generateDelegatesGoFile(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	// Package declaration
	sb.WriteString(fmt.Sprintf("package %s\n\n", m.Name))

	// Add comment header
	sb.WriteString(fmt.Sprintf("// Generated from %s\n\n", m.Name))

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

// generateSharedHFile generates the .h file for types
func (g *GolangGenerator) generateSharedHFile(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	// Header guard and includes
	sb.WriteString("#pragma once\n\n")
	sb.WriteString("#include <stdlib.h>\n")
	sb.WriteString("#include <stdint.h>\n")
	sb.WriteString("#include <stdbool.h>\n\n")

	// Type definitions
	sb.WriteString("typedef struct String { char* data; size_t size; size_t cap; } String;\n")
	sb.WriteString("typedef struct Vector { void* begin; void* end; void* capacity; } Vector;\n")
	sb.WriteString("typedef struct Vector2 { float x, y; } Vector2;\n")
	sb.WriteString("typedef struct Vector3 { float x, y, z; } Vector3;\n")
	sb.WriteString("typedef struct Vector4 { float x, y, z, w; } Vector4;\n")
	sb.WriteString("typedef struct Matrix4x4 { float m[4][4]; } Matrix4x4;\n")
	sb.WriteString("typedef struct Variant {\n")
	sb.WriteString("\tunion {\n")
	sb.WriteString("\tbool boolean;\n")
	sb.WriteString("\tchar char8;\n")
	sb.WriteString("\twchar_t char16;\n")
	sb.WriteString("\tint8_t int8;\n")
	sb.WriteString("\tint16_t int16;\n")
	sb.WriteString("\tint32_t int32;\n")
	sb.WriteString("\tint64_t int64;\n")
	sb.WriteString("\tuint8_t uint8;\n")
	sb.WriteString("\tuint16_t uint16;\n")
	sb.WriteString("\tuint32_t uint32;\n")
	sb.WriteString("\tuint64_t uint64;\n")
	sb.WriteString("\tvoid* ptr;\n")
	sb.WriteString("\tfloat flt;\n")
	sb.WriteString("\tdouble dbl;\n")
	sb.WriteString("\tString str;\n")
	sb.WriteString("\tVector vec;\n")
	sb.WriteString("\tVector2 vec2;\n")
	sb.WriteString("\tVector3 vec3;\n")
	sb.WriteString("\tVector4 vec4;\n")
	sb.WriteString("\t};\n")
	sb.WriteString("#if INTPTR_MAX == INT32_MAX\n")
	sb.WriteString("\tvolatile char pad[8];\n")
	sb.WriteString("#endif\n")
	sb.WriteString("\tuint8_t current;\n")
	sb.WriteString("} Variant;\n\n")

	// External declarations
	sb.WriteString("extern void* Plugify_GetMethodPtr(const char* methodName);\n")
	sb.WriteString("extern void Plugify_GetMethodPtr2(const char* methodName, void** addressDest);\n\n")

	return sb.String(), nil
}

// generateGroupGoFile generates a group-specific Go file
func (g *GolangGenerator) generateGroupGoFile(m *manifest.Manifest, groupName string) (string, error) {
	var sb strings.Builder

	// Package declaration
	sb.WriteString(fmt.Sprintf("package %s\n\n", m.Name))

	// CGo comment block
	sb.WriteString("/*\n")
	sb.WriteString(fmt.Sprintf("#include \"%s.h\"\n", groupName))

	// Add noescape directives for methods in this group
	for _, method := range m.Methods {
		methodGroup := g.GetGroupName(method.Group)
		if methodGroup == groupName {
			sb.WriteString(fmt.Sprintf("#cgo noescape %s\n", method.Name))
		}
	}

	sb.WriteString("*/\n")
	sb.WriteString("import \"C\"\n")
	sb.WriteString("import (\n")
	sb.WriteString("\t\"errors\"\n")
	sb.WriteString("\t\"reflect\"\n")
	sb.WriteString("\t\"runtime\"\n")
	sb.WriteString("\t\"unsafe\"\n")
	sb.WriteString("\t\"github.com/untrustedmodders/go-plugify\"\n")
	sb.WriteString(")\n\n")

	sb.WriteString("var _ = errors.New(\"\")\n")
	sb.WriteString("var _ = reflect.TypeOf(0)\n")
	sb.WriteString("var _ = runtime.GOOS\n")
	sb.WriteString("var _ = unsafe.Sizeof(0)\n")
	sb.WriteString("var _ = plugify.Plugin.Loaded\n\n")

	// Add comment header
	sb.WriteString(fmt.Sprintf("// Generated from %s (group: %s)\n\n", m.Name, groupName))

	// Generate methods for this group
	for _, method := range m.Methods {
		methodGroup := g.GetGroupName(method.Group)
		if methodGroup == groupName {
			methodCode, err := g.generateMethod(&method)
			if err != nil {
				return "", fmt.Errorf("failed to generate method %s: %w", method.Name, err)
			}
			sb.WriteString(methodCode)
			sb.WriteString("\n")
		}
	}

	// Generate classes for this group
	for _, class := range m.Classes {
		classGroup := g.GetGroupName(class.Group)
		if classGroup == groupName {
			classCode, err := g.generateClass(m, &class)
			if err != nil {
				return "", fmt.Errorf("failed to generate class %s: %w", class.Name, err)
			}
			sb.WriteString(classCode)
		}
	}

	return sb.String(), nil
}

// generateGroupHFile generates a group-specific C header file
func (g *GolangGenerator) generateGroupHFile(m *manifest.Manifest, groupName string) (string, error) {
	var sb strings.Builder

	// Header guard and includes
	sb.WriteString("#pragma once\n\n")
	sb.WriteString("#include \"shared.h\"\n\n")

	// Method implementations for this group
	for _, method := range m.Methods {
		methodGroup := g.GetGroupName(method.Group)
		if methodGroup == groupName {
			methodCode, err := g.generateHMethod(&method, m.Name)
			if err != nil {
				return "", err
			}
			sb.WriteString(methodCode)
			sb.WriteString("\n")
		}
	}

	return sb.String(), nil
}
