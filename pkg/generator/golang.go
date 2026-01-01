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
	usedNames  map[string]struct{}
	usedLocals map[string]int
}

// NewGolangGenerator creates a new Go generator
func NewGolangGenerator() *GolangGenerator {
	mapper := NewGolangTypeMapper()
	return &GolangGenerator{
		BaseGenerator: NewBaseGenerator("golang", mapper, GoReservedWords),
		typeMapper:    mapper,
		usedNames:     make(map[string]struct{}),
		usedLocals:    make(map[string]int),
	}
}

// Generate generates Go bindings (.go and .h files)
func (g *GolangGenerator) Generate(m *manifest.Manifest, opts *GeneratorOptions) (*GeneratorResult, error) {
	g.ResetCaches()
	m.Sanitize(g.Sanitizer)
	g.usedNames = make(map[string]struct{})
	g.usedLocals = make(map[string]int)
	opts = EnsureOptions(opts)

	files := make(map[string]string)

	// Collect all unique groups from methods and classes
	groups := g.GetGroups(m)

	// Generate separate enums file
	enumsGoCode, err := g.generateEnumsGoFile(m)
	if err != nil {
		return nil, fmt.Errorf("generating enums file: %w", err)
	}
	files[fmt.Sprintf("%s/enums.go", m.Name)] = enumsGoCode

	// Generate separate aliases file
	aliasesGoCode, err := g.generateAliasesGoFile(m)
	if err != nil {
		return nil, fmt.Errorf("generating aliases file: %w", err)
	}
	files[fmt.Sprintf("%s/aliases.go", m.Name)] = aliasesGoCode

	// Generate separate delegates file
	delegatesGoCode, err := g.generateDelegatesGoFile(m)
	if err != nil {
		return nil, fmt.Errorf("generating delegates file: %w", err)
	}
	files[fmt.Sprintf("%s/delegates.go", m.Name)] = delegatesGoCode

	// Generate .h file for enums
	sharedHCode, err := g.generateSharedHFile(m)
	if err != nil {
		return nil, fmt.Errorf("generating shared header: %w", err)
	}
	files[fmt.Sprintf("%s/shared.h", m.Name)] = sharedHCode

	// Generate group-specific files
	for groupName := range groups {
		goCode, err := g.generateGroupGoFile(m, groupName, opts)
		if err != nil {
			return nil, fmt.Errorf("failed to generate group %s: %w", groupName, err)
		}
		files[fmt.Sprintf("%s/%s.go", m.Name, groupName)] = goCode

		hCode, err := g.generateGroupHFile(m, groupName, opts)
		if err != nil {
			return nil, fmt.Errorf("failed to generate group %s header: %w", groupName, err)
		}
		files[fmt.Sprintf("%s/%s.h", m.Name, groupName)] = hCode

		cCode, err := g.generateGroupCFile(m, groupName, opts)
		if err != nil {
			return nil, fmt.Errorf("failed to generate group %s impl: %w", groupName, err)
		}
		files[fmt.Sprintf("%s/%s.c", m.Name, groupName)] = cCode
	}

	result := &GeneratorResult{
		Files: files,
	}

	return result, nil
}

// generateEnums generates enum definitions
func (g *GolangGenerator) generateEnums(m *manifest.Manifest) (string, error) {
	return g.CollectEnums(m, g.generateEnum)
}

func (g *GolangGenerator) generateEnum(enum *manifest.Enum, underlyingType string) (string, error) {
	var sb strings.Builder

	// Add enum description
	if enum.Description != "" {
		sb.WriteString(fmt.Sprintf("// %s - %s\n", enum.Name, enum.Description))
	}

	sb.WriteString(fmt.Sprintf("type %s = %s\n\n", enum.Name, underlyingType))
	sb.WriteString("const (\n")

	for _, value := range enum.Values {
		rawName := value.Name
		baseName := rawName

		// Track local duplicates
		g.usedLocals[baseName]++
		isLocalDup := g.usedLocals[baseName] > 1

		candidate := enum.Name + "_" + baseName

		// Resolve conflicts
		_, ok := g.usedNames[candidate]
		if ok || isLocalDup {
			candidate = g.resolveConflict(candidate)
		}

		g.usedNames[candidate] = struct{}{}

		// Add value description
		if value.Description != "" {
			sb.WriteString(fmt.Sprintf("\t// %s - %s\n", rawName, value.Description))
		}

		sb.WriteString(fmt.Sprintf("\t%s %s = %d\n", candidate, enum.Name, value.Value))
	}

	sb.WriteString(")\n")
	return sb.String(), nil
}

// generateAliases generates aliases definitions
func (g *GolangGenerator) generateAliases(m *manifest.Manifest) (string, error) {
	return g.CollectAliases(m, g.generateAlias)
}

func (g *GolangGenerator) generateAlias(alias *manifest.Alias, underlyingType string) (string, error) {
	var sb strings.Builder

	// Add alias description
	if alias.Description != "" {
		sb.WriteString(fmt.Sprintf("// %s - %s\n", alias.Name, alias.Description))
	}

	sb.WriteString(fmt.Sprintf("type %s = %s\n", alias.Name, underlyingType))

	return sb.String(), nil
}

// resolveConflict generates a unique name by appending a suffix
func (g *GolangGenerator) resolveConflict(name string) string {
	_, ok := g.usedNames[name]
	if !ok {
		return name
	}

	suffix := 2
	for {
		trial := fmt.Sprintf("%s_%d", name, suffix)
		_, ok := g.usedNames[trial]
		if !ok {
			return trial
		}
		suffix++
	}
}

// generateOwnershipTypes generates noCopy and ownership type definitions
func (g *GolangGenerator) generateOwnershipTypes() string {
	ownership := strings.ToLower(OwnershipEnumName)
	var sb strings.Builder
	sb.WriteString("// noCopy prevents copying via go vet\n")
	sb.WriteString("type noCopy struct{}\n\n")
	sb.WriteString("func (*noCopy) Lock()   {}\n")
	sb.WriteString("func (*noCopy) Unlock() {}\n\n")
	sb.WriteString("// ownership indicates whether the instance owns the underlying handle\n")
	sb.WriteString(fmt.Sprintf("type %s bool\n\n", ownership))
	sb.WriteString("const (\n")
	sb.WriteString(fmt.Sprintf("\tOwned    %s = true\n", ownership))
	sb.WriteString(fmt.Sprintf("\tBorrowed %s = false\n", ownership))
	sb.WriteString(")\n\n")
	return sb.String()
}

// generateDocumentation generates documentation comments for functions/methods
// If useBriefFormat is true, uses "// name \n//  @brief description" format, otherwise "// name - description"
func (g *GolangGenerator) generateDocumentation(name string, description string, params []manifest.ParamType, retType *manifest.RetType, useBriefFormat bool) string {
	var sb strings.Builder

	if description != "" {
		if useBriefFormat {
			sb.WriteString(fmt.Sprintf("// %s \n//  @brief %s\n//\n", name, description))
		} else {
			sb.WriteString(fmt.Sprintf("// %s - %s\n", name, description))
		}
	}

	for _, param := range params {
		if param.Description != "" {
			sb.WriteString(fmt.Sprintf("//  @param %s: %s\n", param.Name, param.Description))
		}
	}

	if retType != nil && retType.Type != "void" && retType.Description != "" {
		sb.WriteString(fmt.Sprintf("//\n//  @return %s\n", retType.Description))
	}

	return sb.String()
}

// processParameters is a generic helper that iterates over parameters and formats them
func (g *GolangGenerator) processParameters(params []manifest.ParamType, processor func(int, *manifest.ParamType) (string, error)) ([]string, error) {
	var result []string
	for i, param := range params {
		part, err := processor(i, &param)
		if err != nil {
			return nil, err
		}
		if part != "" {
			result = append(result, part)
		}
	}
	return result, nil
}

// generateDelegates generates delegate (function type) definitions
func (g *GolangGenerator) generateDelegates(m *manifest.Manifest) (string, error) {
	return g.CollectDelegates(m, g.generateDelegate)
}

func (g *GolangGenerator) generateDelegate(proto *manifest.Prototype) (string, error) {
	var sb strings.Builder

	// Add delegate description
	if proto.Description != "" {
		sb.WriteString(fmt.Sprintf("// %s - %s\n", proto.Name, proto.Description))
	}

	// Generate parameter list
	params, err := g.formatParams(proto.ParamTypes, true)
	if err != nil {
		return "", err
	}

	// Generate return type
	returnType := ""
	if proto.RetType.Type != "void" {
		returnType, err = g.typeMapper.MapReturnType(&proto.RetType)
		if err != nil {
			return "", err
		}
	}

	// Generate function type
	sb.WriteString(fmt.Sprintf("type %s func(%s)", proto.Name, params))
	if returnType != "" {
		sb.WriteString(fmt.Sprintf(" %s", returnType))
	}
	sb.WriteString("\n\n")

	return sb.String(), nil
}

// generateMethod generates a single method implementation
func (g *GolangGenerator) generateMethod(method *manifest.Method) (string, error) {
	var sb strings.Builder

	// Add deprecation comment if present
	if method.Deprecated != "" {
		sb.WriteString(fmt.Sprintf("// Deprecated: %s\n", method.Deprecated))
	}

	// Generate documentation
	sb.WriteString(g.generateDocumentation(method.Name, method.Description, method.ParamTypes, &method.RetType, true))

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
// goMethodBodyContext holds the context for generating a Go method body
type goMethodBodyContext struct {
	indent           string
	innerIndent      string
	isObjRet         bool
	isPodRet         bool
	hasRet           bool
	hasCast          bool
	hasTry           bool
	insertIndexStart int
	insertIndexEnd   int
	paramsCast       []string
}

func (g *GolangGenerator) generateMethodBody(method *manifest.Method) (string, error) {
	var sb strings.Builder

	// Set up context
	ctx, err := g.initializeGoMethodContext(method)
	if err != nil {
		return "", err
	}

	// Declare return variables
	g.declareGoMethodVariables(method, &sb, ctx)

	// Add parameter casting and try block
	g.setupGoCastingAndTry(&sb, ctx)

	// Generate function call
	if err := g.generateGoMethodCall(method, &sb, ctx); err != nil {
		return "", err
	}

	// Generate unmarshaling
	if err := g.generateGoUnmarshalingSection(method, &sb, ctx); err != nil {
		return "", err
	}

	// Generate cleanup
	if err := g.generateGoCleanupSection(method, &sb, ctx); err != nil {
		return "", err
	}

	// Generate return statement
	g.generateGoReturnStatement(method, &sb, ctx)

	return sb.String(), nil
}

// initializeGoMethodContext sets up the context for Go method body generation
func (g *GolangGenerator) initializeGoMethodContext(method *manifest.Method) (*goMethodBodyContext, error) {
	ctx := &goMethodBodyContext{
		indent:      "\t",
		innerIndent: "\t",
		isObjRet:    method.RetType.Type != "void" && g.typeMapper.IsObjType(method.RetType.Type),
		isPodRet:    method.RetType.Type != "void" && g.typeMapper.IsPodType(method.RetType.Type),
	}

	ctx.hasRet = method.RetType.Type != "void" && !ctx.isObjRet

	// Generate parameter casting
	var err error
	ctx.paramsCast, err = g.generateParamsCast(method, ctx.indent)
	if err != nil {
		return nil, err
	}

	ctx.hasCast = len(ctx.paramsCast) > 0 && method.RetType.Type != "void"

	return ctx, nil
}

// declareGoMethodVariables declares the necessary variables for the Go method
func (g *GolangGenerator) declareGoMethodVariables(method *manifest.Method, sb *strings.Builder, ctx *goMethodBodyContext) {
	// Declare return value if needed
	if ctx.isObjRet || ctx.hasCast {
		returnType, _ := g.typeMapper.MapReturnType(&method.RetType)
		sb.WriteString(fmt.Sprintf("%svar __retVal %s\n", ctx.indent, returnType))
	}
}

// setupGoCastingAndTry adds parameter casting code and sets up try block if needed
func (g *GolangGenerator) setupGoCastingAndTry(sb *strings.Builder, ctx *goMethodBodyContext) {
	if len(ctx.paramsCast) > 0 {
		for _, cast := range ctx.paramsCast {
			sb.WriteString(cast)
		}
		ctx.insertIndexStart = sb.Len()
		sb.WriteString(fmt.Sprintf("%splugify.Block {\n", ctx.indent))
		sb.WriteString(fmt.Sprintf("%s\tTry: func() {\n", ctx.indent))
		ctx.insertIndexEnd = sb.Len()
		ctx.innerIndent = "\t\t\t"
		ctx.hasTry = ctx.hasCast
	}
}

// generateGoMethodCall generates the actual function call with proper type handling
func (g *GolangGenerator) generateGoMethodCall(method *manifest.Method, sb *strings.Builder, ctx *goMethodBodyContext) error {
	functionCall, err := g.generateFunctionCall(method)
	if err != nil {
		return err
	}

	if ctx.isObjRet {
		retTypeCast := g.typeMapper.retTypeCastMap[method.RetType.Type]
		sb.WriteString(fmt.Sprintf("%s__native := %s\n", ctx.innerIndent, functionCall))
		sb.WriteString(fmt.Sprintf("%s__retVal_native = *(*%s)(unsafe.Pointer(&__native))\n", ctx.innerIndent, retTypeCast))
	} else if ctx.hasTry {
		if ctx.isPodRet {
			retTypeCast := g.typeMapper.retTypeCastMap[method.RetType.Type]
			sb.WriteString(fmt.Sprintf("%s__native := %s\n", ctx.innerIndent, functionCall))
			sb.WriteString(fmt.Sprintf("%s__retVal = *(*%s)(unsafe.Pointer(&__native))\n", ctx.innerIndent, retTypeCast))
		} else {
			retTypeCast := g.typeMapper.retTypeCastMap[method.RetType.Type]
			if retTypeCast != "" {
				sb.WriteString(fmt.Sprintf("%s__retVal = %s(%s)\n", ctx.innerIndent, retTypeCast, functionCall))
			} else {
				sb.WriteString(fmt.Sprintf("%s__retVal = %s\n", ctx.innerIndent, functionCall))
			}
		}
	} else if ctx.hasRet {
		if ctx.isPodRet {
			retTypeCast := g.typeMapper.retTypeCastMap[method.RetType.Type]
			sb.WriteString(fmt.Sprintf("%s__native := %s\n", ctx.innerIndent, functionCall))
			sb.WriteString(fmt.Sprintf("%s__retVal := *(*%s)(unsafe.Pointer(&__native))\n", ctx.innerIndent, retTypeCast))
		} else if method.RetType.Type == "function" {
			returnType, _ := g.typeMapper.MapReturnType(&method.RetType)
			sb.WriteString(fmt.Sprintf("%s__retVal := plugify.GetDelegateForFunctionPointer(%s, reflect.TypeOf(%s(nil))).(%s)\n",
				ctx.innerIndent, functionCall, returnType, returnType))
		} else {
			returnType, _ := g.typeMapper.MapReturnType(&method.RetType)
			sb.WriteString(fmt.Sprintf("%s__retVal := %s(%s)\n", ctx.innerIndent, returnType, functionCall))
		}
	} else {
		sb.WriteString(fmt.Sprintf("%s%s\n", ctx.innerIndent, functionCall))
	}

	return nil
}

// generateGoUnmarshalingSection generates unmarshaling code for ref parameters
func (g *GolangGenerator) generateGoUnmarshalingSection(method *manifest.Method, sb *strings.Builder, ctx *goMethodBodyContext) error {
	assignCast, err := g.generateParamsCastAssign(method, ctx.innerIndent)
	if err != nil {
		return err
	}

	if len(assignCast) > 0 {
		sb.WriteString(fmt.Sprintf("%s// Unmarshal - Convert native data to managed data.\n", ctx.innerIndent))
		for _, assign := range assignCast {
			sb.WriteString(assign)
		}
	}

	return nil
}

// generateGoCleanupSection generates cleanup code and closes try block if needed
func (g *GolangGenerator) generateGoCleanupSection(method *manifest.Method, sb *strings.Builder, ctx *goMethodBodyContext) error {
	cleanupCast, err := g.generateParamsCastCleanup(method, ctx.innerIndent)
	if err != nil {
		return err
	}

	if len(cleanupCast) > 0 {
		sb.WriteString(fmt.Sprintf("%s\t},\n%s\tFinally: func() {\n%s// Perform cleanup.\n", ctx.indent, ctx.indent, ctx.innerIndent))
		for _, cleanup := range cleanupCast {
			sb.WriteString(cleanup)
		}
		sb.WriteString(fmt.Sprintf("%s\t},\n%s}.Do()\n", ctx.indent, ctx.indent))
	} else if len(ctx.paramsCast) > 0 {
		RemoveFromBuilder(sb, ctx.insertIndexStart, ctx.insertIndexEnd)
		RemoveLeadingTabs(sb, 2, ctx.insertIndexStart, sb.Len())
	}

	return nil
}

// generateGoReturnStatement generates the return statement
func (g *GolangGenerator) generateGoReturnStatement(method *manifest.Method, sb *strings.Builder, ctx *goMethodBodyContext) {
	if method.RetType.Type != "void" {
		sb.WriteString(fmt.Sprintf("%sreturn __retVal\n", ctx.indent))
	}
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

	// Handle parameters using generic processor
	paramCasts, err := g.processParameters(method.ParamTypes, func(_ int, param *manifest.ParamType) (string, error) {
		return g.generateParamCast(param, indent)
	})
	if err != nil {
		return nil, err
	}
	result = append(result, paramCasts...)

	return result, nil
}

// generateParamCast generates casting code for a single parameter
func (g *GolangGenerator) generateParamCast(param *manifest.ParamType, indent string) (string, error) {
	paramType := g.typeMapper.valTypeCastMap[param.Type]
	name := param.Name

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

	// Handle parameters using generic processor
	paramAssigns, err := g.processParameters(method.ParamTypes, func(_ int, param *manifest.ParamType) (string, error) {
		if param.Ref {
			return g.generateParamAssign(param, indent)
		}
		return "", nil
	})
	if err != nil {
		return nil, err
	}
	result = append(result, paramAssigns...)

	return result, nil
}

// generateParamAssign generates assignment code for a ref parameter
func (g *GolangGenerator) generateParamAssign(param *manifest.ParamType, indent string) (string, error) {
	paramType := g.typeMapper.assTypeCastMap[param.Type]
	name := param.Name

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
func (g *GolangGenerator) generateReturnAssign(retType *manifest.RetType, indent string) (string, error) {
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

	// Handle parameters using generic processor
	paramCleanups, err := g.processParameters(method.ParamTypes, func(_ int, param *manifest.ParamType) (string, error) {
		cleanup := g.generateParamCleanup(param)
		if cleanup != "" {
			return indent + cleanup, nil
		}
		return "", nil
	})
	if err != nil {
		return nil, err
	}
	result = append(result, paramCleanups...)

	return result, nil
}

// generateParamCleanup generates cleanup code for a parameter
func (g *GolangGenerator) generateParamCleanup(param *manifest.ParamType) string {
	paramType := g.typeMapper.delTypeCastMap[param.Type]
	if paramType == "" {
		return ""
	}

	name := param.Name
	return fmt.Sprintf("%s(&__%s)\n", paramType, name)
}

// generateReturnCleanup generates cleanup code for return value
func (g *GolangGenerator) generateReturnCleanup(retType *manifest.RetType, indent string) string {
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
	for _, param := range params {
		name := param.Name

		if withTypes {
			typeName, err := g.typeMapper.MapParamType(&param)
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
		name := param.Name

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

	/// Generate extern pointer
	sb.WriteString(fmt.Sprintf("extern %s (*__%s_%s)(%s);\n\n",
		retType, pluginName, method.Name, paramTypes))

	// Generate wrapper function
	sb.WriteString(fmt.Sprintf("static %s %s(%s) {\n", retType, method.Name, paramList))

	if method.RetType.Type != "void" {
		sb.WriteString(fmt.Sprintf("\treturn __%s_%s(%s);\n", pluginName, method.Name, paramNames))
	} else {
		sb.WriteString(fmt.Sprintf("\t__%s_%s(%s);\n", pluginName, method.Name, paramNames))
	}

	sb.WriteString("}\n")

	return sb.String(), nil
}

// generateCMethod generates a single C method in the .c file
func (g *GolangGenerator) generateCMethod(method *manifest.Method, pluginName string) (string, error) {
	var sb strings.Builder

	// Get return type
	retType := "void"
	if method.RetType.Type != "void" {
		retType = g.typeMapper.GetCType(method.RetType.Type, false, true)
	}

	paramTypes, err := g.formatCParams(method.ParamTypes, false)
	if err != nil {
		return "", err
	}

	/// Generate impl pointer
	sb.WriteString(fmt.Sprintf("PLUGIFY_EXPORT %s (*__%s_%s)(%s) = NULL;\n\n",
		retType, pluginName, method.Name, paramTypes))

	return sb.String(), nil
}

// formatCParams formats C parameters
func (g *GolangGenerator) formatCParams(params []manifest.ParamType, withNames bool) (string, error) {
	if len(params) == 0 {
		return "", nil
	}

	var parts []string
	for _, param := range params {
		typeName := g.typeMapper.GetCType(param.Type, param.Ref, false)

		if withNames {
			name := param.Name
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
	for _, param := range params {
		name := param.Name
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
		sb.WriteString(g.generateOwnershipTypes())
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
		sb.WriteString(fmt.Sprintf("var (\n\t%s = errors.New(\"%s: %s\")\n)\n\n", errVarName, class.Name, EmptyHandleError))
	}

	// Class documentation
	if class.Deprecated != "" {
		sb.WriteString(fmt.Sprintf("// Deprecated: %s\n", class.Deprecated))
	}
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
	method := FindMethod(m, methodName)
	if method == nil {
		return "", fmt.Errorf("constructor method %s not found", methodName)
	}

	var sb strings.Builder

	// Add deprecation comment if present
	if method.Deprecated != "" {
		sb.WriteString(fmt.Sprintf("// Deprecated: %s\n", method.Deprecated))
	}

	// Generate documentation
	funcName := fmt.Sprintf("New%s%s", class.Name, method.Name)
	sb.WriteString(g.generateDocumentation(funcName, method.Description, method.ParamTypes, nil, false))

	// Generate function signature
	params, err := g.formatParams(method.ParamTypes, true)
	if err != nil {
		return "", err
	}
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

	// Add deprecation comment if present (check both binding and underlying method)
	deprecationReason := binding.Deprecated
	if deprecationReason == "" {
		deprecationReason = method.Deprecated
	}
	if deprecationReason != "" {
		sb.WriteString(fmt.Sprintf("// Deprecated: %s\n", deprecationReason))
	}

	// Generate documentation
	sb.WriteString(g.generateDocumentation(binding.Name, method.Description, methodParams, &method.RetType, false))

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

	//hasCtor := len(class.Constructors) > 0
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
			if hasDtor /*|| hasCtor*/ {
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
		name := param.Name

		if withTypes {
			typeName := ""
			var err error

			// Check if this parameter has an alias
			if i < len(aliases) && aliases[i] != nil {
				typeName = fmt.Sprintf("*%s", aliases[i].Name)
			} else {
				typeName, err = g.typeMapper.MapParamType(&param)
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
		name := param.Name

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
		sb.WriteString(g.generateOwnershipTypes())
	}

	return sb.String(), nil
}

// generateAliasesGoFile generates a file containing all aliases
func (g *GolangGenerator) generateAliasesGoFile(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	// Package declaration
	sb.WriteString(fmt.Sprintf("package %s\n\n", m.Name))

	// Add comment header
	sb.WriteString(fmt.Sprintf("// Generated from %s\n\n", m.Name))

	// Generate aliases
	aliasesCode, err := g.generateAliases(m)
	if err != nil {
		return "", err
	}
	if aliasesCode != "" {
		sb.WriteString(aliasesCode)
		sb.WriteString("\n")
	}

	return sb.String(), nil
}

// generateDelegatesGoFile generates a file containing all delegate types
func (g *GolangGenerator) generateDelegatesGoFile(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	// Package declaration
	sb.WriteString(fmt.Sprintf("package %s\n\n", m.Name))

	sb.WriteString("import \"github.com/untrustedmodders/go-plugify\"\n\n")

	sb.WriteString("var _ = plugify.Plugin.Loaded\n\n")

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

	sb.WriteString("#if defined(_WIN32)\n")
	sb.WriteString("#define PLUGIFY_EXPORT __declspec(dllexport)\n")
	sb.WriteString("#else\n")
	sb.WriteString("#define PLUGIFY_EXPORT __attribute__((visibility(\"default\")))\n")
	sb.WriteString("#endif\n\n")

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

	return sb.String(), nil
}

// generateGroupGoFile generates a group-specific Go file
func (g *GolangGenerator) generateGroupGoFile(m *manifest.Manifest, groupName string, opts *GeneratorOptions) (string, error) {
	var sb strings.Builder

	// Package declaration
	sb.WriteString(fmt.Sprintf("package %s\n\n", m.Name))

	// CGo comment block
	sb.WriteString("/*\n")
	sb.WriteString(fmt.Sprintf("#include \"%s.h\"\n", groupName))

	// Add noescape directives for methods in this group
	for _, method := range m.Methods {
		methodGroup := method.Group
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
		methodGroup := method.Group
		if methodGroup == groupName {
			methodCode, err := g.generateMethod(&method)
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
			}
		}
	}

	return sb.String(), nil
}

// generateGroupHFile generates a group-specific C header file
func (g *GolangGenerator) generateGroupHFile(m *manifest.Manifest, groupName string, opts *GeneratorOptions) (string, error) {
	var sb strings.Builder

	// Header guard and includes
	sb.WriteString("#pragma once\n\n")
	sb.WriteString("#include \"shared.h\"\n\n")

	// Method implementations for this group
	for _, method := range m.Methods {
		methodGroup := method.Group
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

// generateGroupCFile generates a group-specific C impl file
func (g *GolangGenerator) generateGroupCFile(m *manifest.Manifest, groupName string, opts *GeneratorOptions) (string, error) {
	var sb strings.Builder

	// Header guard and includes
	sb.WriteString("#include \"shared.h\"\n\n")

	// Method implementations for this group
	for _, method := range m.Methods {
		methodGroup := method.Group
		if methodGroup == groupName {
			methodCode, err := g.generateCMethod(&method, m.Name)
			if err != nil {
				return "", err
			}
			sb.WriteString(methodCode)
			sb.WriteString("\n")
		}
	}

	return sb.String(), nil
}
