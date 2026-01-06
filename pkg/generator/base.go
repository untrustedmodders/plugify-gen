package generator

import (
	"fmt"
	"strings"

	"github.com/untrustedmodders/plugify-gen/pkg/manifest"
)

// Common constants used across generators
const (
	// DefaultGroupName is used when a method or class has no group specified
	DefaultGroupName = "core"

	// OwnershipEnumName is the standard name for the ownership enum
	OwnershipEnumName = "Ownership"

	// Standard error messages
	EmptyHandleError = "empty handle"
)

// GeneratorOptions contains options for code generation
type GeneratorOptions struct {
	// GenerateClasses controls whether to generate class wrappers
	GenerateClasses bool
}

// EnsureOptions returns valid options, using defaults if nil
func EnsureOptions(opts *GeneratorOptions) *GeneratorOptions {
	if opts == nil {
		return &GeneratorOptions{GenerateClasses: true}
	}
	return opts
}

// Generator is the interface that all language-specific generators must implement
type Generator interface {
	// Name returns the generator name (e.g., "cpp", "golang")
	Name() string

	// Generate produces code from a manifest
	Generate(m *manifest.Manifest, opts *GeneratorOptions) (*GeneratorResult, error)
}

// GeneratorResult contains the generated files
type GeneratorResult struct {
	// Files maps filename to file content
	Files map[string]string
}

// DocOptions holds common documentation generation options
// This provides a unified interface for documentation across different language generators
type DocOptions struct {
	Indent       string                 // Indentation string (e.g., "\t", "  ")
	Summary      string                 // Summary/brief description
	Description  string                 // Detailed description
	Deprecated   string                 // Detailed deprecation message
	Params       []manifest.ParamType   // Method parameters
	ParamAliases []*manifest.ParamAlias // Parameter aliases for type substitution
	//Returns          string                 // Return value description
	RetType          manifest.RetType   // Return type information
	RetAlias         *manifest.RetAlias // Return value alias
	IncludeCallbacks bool
	AddThrows        bool
}

// DocFormatter is a callback function that formats documentation in language-specific style
// Each generator can provide its own formatter to generate language-appropriate doc comments
type DocFormatter func(opts DocOptions) string

// BaseGenerator provides common functionality for all generators
type BaseGenerator struct {
	name          string
	typeMapper    TypeMapper
	invalidNames  map[string]struct{}
	enumCache     map[string]struct{}
	delegateCache map[string]struct{}
	aliasCache    map[string]struct{}
}

// NewBaseGenerator creates a new base generator
func NewBaseGenerator(name string, typeMapper TypeMapper, invalidNames []string) *BaseGenerator {
	invalidMap := make(map[string]struct{})
	for _, name := range invalidNames {
		invalidMap[name] = struct{}{}
	}

	return &BaseGenerator{
		name:          name,
		typeMapper:    typeMapper,
		invalidNames:  invalidMap,
		enumCache:     make(map[string]struct{}),
		delegateCache: make(map[string]struct{}),
		aliasCache:    make(map[string]struct{}),
	}
}

// Name returns the generator name
func (g *BaseGenerator) Name() string {
	return g.name
}

// Sanitizer handles reserved keywords by appending underscore
func (g *BaseGenerator) Sanitizer(name string) string {
	_, ok := g.invalidNames[name]
	if ok {
		return name + "_"
	}
	return name
}

// GetGroups gets the list of all groups from manifest
func (g *BaseGenerator) GetGroups(m *manifest.Manifest) map[string]struct{} {
	// Collect all unique groups from both methods and classes
	groups := make(map[string]struct{})
	hasUngroupedMethods := false
	hasUngroupedClasses := false

	for _, method := range m.Methods {
		groupName := strings.ToLower(method.Group)
		if groupName != "" {
			groups[groupName] = struct{}{}
		} else {
			hasUngroupedMethods = true
		}
	}

	for _, class := range m.Classes {
		groupName := strings.ToLower(class.Group)
		if groupName != "" {
			groups[groupName] = struct{}{}
		} else {
			hasUngroupedClasses = true
		}
	}

	// Add default group for methods/classes without a group
	if hasUngroupedMethods || hasUngroupedClasses {
		groups[DefaultGroupName] = struct{}{}
	}

	return groups
}

// IsEnumCached checks if an enum has already been generated
func (g *BaseGenerator) IsEnumCached(name string) bool {
	_, ok := g.enumCache[name]
	return ok
}

// CacheEnum marks an enum as generated
func (g *BaseGenerator) CacheEnum(name string) {
	g.enumCache[name] = struct{}{}
}

// IsDelegateCached checks if a delegate has already been generated
func (g *BaseGenerator) IsDelegateCached(name string) bool {
	_, ok := g.delegateCache[name]
	return ok
}

// CacheDelegate marks a delegate as generated
func (g *BaseGenerator) CacheDelegate(name string) {
	g.delegateCache[name] = struct{}{}
}

// IsAliasCached checks if an alias has already been generated
func (g *BaseGenerator) IsAliasCached(name string) bool {
	_, ok := g.aliasCache[name]
	return ok
}

// CacheAlias marks an alias as generated
func (g *BaseGenerator) CacheAlias(name string) {
	g.aliasCache[name] = struct{}{}
}

// ResetCaches clears enum and delegate caches (call before generation)
func (g *BaseGenerator) ResetCaches() {
	g.enumCache = make(map[string]struct{})
	g.delegateCache = make(map[string]struct{})
	g.aliasCache = make(map[string]struct{})
}

// TypeMapper is the interface for type mapping strategies
type TypeMapper interface {
	// MapType converts a manifest type to a target language type
	MapType(baseType string, context TypeContext, isArray bool) (string, error)

	// MapParamType converts a parameter type (handles ref, default values)
	MapParamType(param *manifest.ParamType) (string, error)

	// MapReturnType converts a return type
	MapReturnType(retType *manifest.RetType) (string, error)

	// MapHandleType converts a invalid type
	MapHandleType(class *manifest.Class) (string, string, error)
}

// TypeContext represents the context in which a type appears
type TypeContext int

const (
	TypeContextValue  TypeContext = iota // Value parameter
	TypeContextRef                       // Reference parameter
	TypeContextReturn                    // Return type
)

// EnumGenerator is a callback function that generates code for an enum type
// It takes an enum and returns the generated code as a string
type EnumGenerator func(*manifest.Enum, string) (string, error)

type AliasGenerator func(*manifest.Alias, string) (string, error)

type DelegateGenerator func(*manifest.Prototype) (string, error)

// ParameterFormat specifies what format parameters should be generated in
type ParameterFormat int

const (
	ParamFormatTypes         ParameterFormat = iota // Just types: "int, string"
	ParamFormatNames                                // Just names: "x, y"
	ParamFormatTypesAndNames                        // Types and names: "int x, string y"
)

// FormatParameters is a helper to format method parameters
func FormatParameters(params []manifest.ParamType, format ParameterFormat, mapper TypeMapper) (string, error) {
	if len(params) == 0 {
		return "", nil
	}

	result := ""
	for i, param := range params {
		if i > 0 {
			result += ", "
		}

		switch format {
		case ParamFormatTypes:
			typeName, err := mapper.MapParamType(&param)
			if err != nil {
				return "", err
			}
			result += typeName

		case ParamFormatNames:
			paramName := param.Name
			result += paramName

		case ParamFormatTypesAndNames:
			typeName, err := mapper.MapParamType(&param)
			if err != nil {
				return "", err
			}
			paramName := param.Name
			result += typeName + " " + paramName
			if param.Default != nil {
				result += fmt.Sprintf(" = %d", *param.Default)
			}
		}
	}

	return result, nil
}

// RemoveFromBuilderRange removes the substring from index n1 up to (but not including) n2
// from the builder’s current string. If indices are out of bounds, they are clamped safely.
func RemoveFromBuilder(b *strings.Builder, n1, n2 int) {
	s := b.String()

	// Clamp indices to valid range
	if n1 < 0 {
		n1 = 0
	}
	if n2 > len(s) {
		n2 = len(s)
	}
	if n1 >= n2 {
		return // nothing to remove
	}

	s = s[:n1] + s[n2:]
	b.Reset()
	b.WriteString(s)
}

// RemoveLeadingTabsCharRange removes up to n leading tabs from each line
// whose characters overlap the range [startChar, endChar) in the builder's string.
func RemoveLeadingTabs(b *strings.Builder, n, startChar, endChar int) {
	s := b.String()
	if startChar < 0 {
		startChar = 0
	}
	if endChar > len(s) {
		endChar = len(s)
	}
	if startChar >= endChar {
		return
	}

	lines := strings.Split(s, "\n")
	pos := 0 // track start index of each line in the full string

	for i, line := range lines {
		lineStart := pos
		lineEnd := pos + len(line) // exclusive, \n handled later

		// Skip lines completely outside the range
		if lineEnd <= startChar || lineStart >= endChar {
			pos += len(line) + 1 // +1 for the \n
			continue
		}

		// Only consider characters in the overlap of line and [startChar, endChar)
		localStart := 0
		if startChar > lineStart {
			localStart = startChar - lineStart
		}
		localEnd := len(line)
		if endChar < lineEnd {
			localEnd = endChar - lineStart
		}

		// Remove leading tabs inside the local range
		substr := line[localStart:localEnd]
		count := 0
		for strings.HasPrefix(substr, "\t") && count < n {
			substr = substr[1:]
			count++
		}
		line = line[:localStart] + substr + line[localEnd:]
		lines[i] = line

		pos += len(line) + 1
	}

	b.Reset()
	b.WriteString(strings.Join(lines, "\n"))
}

// FindMethod returns the method with the given name or nil
func FindMethod(m *manifest.Manifest, name string) *manifest.Method {
	for i := range m.Methods {
		if m.Methods[i].Name == name || m.Methods[i].FuncName == name {
			return &m.Methods[i]
		}
	}
	return nil
}

// HasConstructorWithNoParam checks if any constructor has exactly 0 parameter
func (g *BaseGenerator) HasConstructorWithNoParam(m *manifest.Manifest, class *manifest.Class) bool {
	for _, ctorName := range class.Constructors {
		method := FindMethod(m, ctorName)
		if method != nil && len(method.ParamTypes) == 0 {
			return true
		}
	}
	return false
}

// HasConstructorWithSingleHandleParam checks if any constructor has exactly 1 parameter of the handle type
func (g *BaseGenerator) HasConstructorWithSingleHandleParam(m *manifest.Manifest, class *manifest.Class) bool {
	for _, ctorName := range class.Constructors {
		method := FindMethod(m, ctorName)
		if method != nil && len(method.ParamTypes) == 1 && method.ParamTypes[0].Type == class.HandleType {
			return true
		}
	}
	return false
}

// FindDependentGroups finds other groups that this group depends on (methods used by classes)
func (g *BaseGenerator) FindDependentGroups(m *manifest.Manifest, currentGroupName string) map[string]struct{} {
	groupName := currentGroupName
	// Pre-index methods for O(1) lookup.
	methodToGroup := make(map[string]string, len(m.Methods))
	for i := range m.Methods {
		method := &m.Methods[i]
		methodGroup := method.Group

		methodToGroup[method.Name] = methodGroup
		//methodToGroup[method.FuncName] = methodGroup
	}

	referencedGroups := make(map[string]struct{})

	for _, class := range m.Classes {
		classGroup := class.Group
		if classGroup != groupName {
			continue
		}

		for _, ctorName := range class.Constructors {
			if refGroup, ok := methodToGroup[ctorName]; ok && refGroup != classGroup {
				referencedGroups[refGroup] = struct{}{}
			}
		}

		if class.Destructor != nil {
			if refGroup, ok := methodToGroup[*class.Destructor]; ok && refGroup != classGroup {
				referencedGroups[refGroup] = struct{}{}
			}
		}

		for _, binding := range class.Bindings {
			if refGroup, ok := methodToGroup[binding.Method]; ok && refGroup != classGroup {
				referencedGroups[refGroup] = struct{}{}
			}
		}
	}

	return referencedGroups
}

type enumHandler func(enum *manifest.Enum, typeName string) error
type aliasHandler func(alias *manifest.Alias, typeName string) error
type protoHandler func(proto *manifest.Prototype) error

// walkPrototype traverses proto graph and calls handlers when enums/prototypes are found.
// visitedProtos prevents infinite recursion and duplicate work.
func (g *BaseGenerator) walkPrototype(proto *manifest.Prototype, onEnum enumHandler, onAlias aliasHandler, onProto protoHandler, visitedProtos map[string]struct{}) error {
	if visitedProtos == nil {
		visitedProtos = make(map[string]struct{})
	}
	if proto.Name != "" {
		_, ok := visitedProtos[proto.Name]
		if ok {
			return nil
		}
		visitedProtos[proto.Name] = struct{}{}
	}

	// Return type
	if proto.RetType.Enum != nil {
		if err := onEnum(proto.RetType.Enum, proto.RetType.Type); err != nil {
			return err
		}
	}
	if proto.RetType.Alias != nil {
		if err := onAlias(proto.RetType.Alias, proto.RetType.Type); err != nil {
			return err
		}
	}
	if proto.RetType.Prototype != nil {
		if err := onProto(proto.RetType.Prototype); err != nil {
			return err
		}
		if err := g.walkPrototype(proto.RetType.Prototype, onEnum, onAlias, onProto, visitedProtos); err != nil {
			return err
		}
	}

	// Parameters
	for _, param := range proto.ParamTypes {
		if param.Enum != nil {
			if err := onEnum(param.Enum, param.Type); err != nil {
				return err
			}
		}
		if param.Alias != nil {
			if err := onAlias(param.Alias, param.Type); err != nil {
				return err
			}
		}
		if param.Prototype != nil {
			if err := onProto(param.Prototype); err != nil {
				return err
			}
			if err := g.walkPrototype(param.Prototype, onEnum, onAlias, onProto, visitedProtos); err != nil {
				return err
			}
		}
	}

	return nil
}

// ensureEnumGenerated centralizes mapType -> enumGen -> write -> cache
func (g *BaseGenerator) ensureEnumGenerated(enum *manifest.Enum, typeName string, context TypeContext, sb *strings.Builder, enumGen EnumGenerator) error {
	if g.IsEnumCached(enum.Name) {
		return nil
	}
	mapped, err := g.typeMapper.MapType(typeName, context, false)
	if err != nil {
		return err
	}
	enumCode, err := enumGen(enum, mapped)
	if err != nil {
		return err
	}
	sb.WriteString(enumCode)
	sb.WriteString("\n")
	g.CacheEnum(enum.Name)
	return nil
}

// ensureEnumGenerated centralizes mapType -> enumGen -> write -> cache
func (g *BaseGenerator) ensureAliasGenerated(alias *manifest.Alias, typeName string, context TypeContext, sb *strings.Builder, aliasGen AliasGenerator) error {
	if g.IsAliasCached(alias.Name) {
		return nil
	}
	mapped, err := g.typeMapper.MapType(typeName, context, false)
	if err != nil {
		return err
	}
	aliasCode, err := aliasGen(alias, mapped)
	if err != nil {
		return err
	}
	sb.WriteString(aliasCode)
	sb.WriteString("\n")
	g.CacheAlias(alias.Name)
	return nil
}

// ensureDelegateGenerated centralizes delegate generation -> write -> cache
func (g *BaseGenerator) ensureDelegateGenerated(proto *manifest.Prototype, sb *strings.Builder, delegateGen DelegateGenerator) error {
	if g.IsDelegateCached(proto.Name) {
		return nil
	}
	code, err := delegateGen(proto)
	if err != nil {
		return err
	}
	sb.WriteString(code)
	sb.WriteString("\n")
	g.CacheDelegate(proto.Name)
	return nil
}

// CollectEnums uses the generic walker and the helper above
func (g *BaseGenerator) CollectEnums(m *manifest.Manifest, enumGen EnumGenerator) (string, error) {
	var sb strings.Builder

	ctx := TypeContextReturn

	onEnum := func(enum *manifest.Enum, typeName string) error {
		return g.ensureEnumGenerated(enum, typeName, ctx, &sb, enumGen)
	}
	onAlias := func(alias *manifest.Alias, typeName string) error {
		return nil
	}
	onProto := func(proto *manifest.Prototype) error {
		return nil
	}

	for _, method := range m.Methods {
		// method return
		if method.RetType.Enum != nil {
			if err := g.ensureEnumGenerated(method.RetType.Enum, method.RetType.Type, ctx, &sb, enumGen); err != nil {
				return "", err
			}
		}
		// walk return prototype
		if method.RetType.Prototype != nil {
			if err := g.walkPrototype(method.RetType.Prototype, onEnum, onAlias, onProto, nil); err != nil {
				return "", err
			}
		}

		// parameters
		for _, param := range method.ParamTypes {
			if param.Enum != nil {
				if err := g.ensureEnumGenerated(param.Enum, param.Type, ctx, &sb, enumGen); err != nil {
					return "", err
				}
			}
			if param.Prototype != nil {
				if err := g.walkPrototype(param.Prototype, onEnum, onAlias, onProto, nil); err != nil {
					return "", err
				}
			}
		}
	}
	return sb.String(), nil
}

// CollectAliases mirrors CollectEnums but uses delegates
func (g *BaseGenerator) CollectAliases(m *manifest.Manifest, aliasGen AliasGenerator) (string, error) {
	var sb strings.Builder

	ctx := TypeContextReturn

	onEnum := func(enum *manifest.Enum, typeName string) error {
		return nil
	}
	onAlias := func(alias *manifest.Alias, typeName string) error {
		return g.ensureAliasGenerated(alias, typeName, ctx, &sb, aliasGen)
	}
	onProto := func(proto *manifest.Prototype) error {
		return nil
	}

	for _, method := range m.Methods {
		// method return
		if method.RetType.Alias != nil {
			if err := g.ensureAliasGenerated(method.RetType.Alias, method.RetType.Type, ctx, &sb, aliasGen); err != nil {
				return "", err
			}
		}
		// walk return prototype
		if method.RetType.Prototype != nil {
			if err := g.walkPrototype(method.RetType.Prototype, onEnum, onAlias, onProto, nil); err != nil {
				return "", err
			}
		}

		// parameters
		for _, param := range method.ParamTypes {
			if param.Alias != nil {
				if err := g.ensureAliasGenerated(param.Alias, param.Type, ctx, &sb, aliasGen); err != nil {
					return "", err
				}
			}
			if param.Prototype != nil {
				if err := g.walkPrototype(param.Prototype, onEnum, onAlias, onProto, nil); err != nil {
					return "", err
				}
			}
		}
	}
	return sb.String(), nil
}

// CollectDelegates mirrors CollectEnums but uses delegates
func (g *BaseGenerator) CollectDelegates(m *manifest.Manifest, delegateGen DelegateGenerator) (string, error) {
	var sb strings.Builder

	onEnum := func(enum *manifest.Enum, typeName string) error {
		return nil
	}
	onAlias := func(alias *manifest.Alias, typeName string) error {
		return nil
	}
	onProto := func(proto *manifest.Prototype) error {
		return g.ensureDelegateGenerated(proto, &sb, delegateGen)
	}

	for _, method := range m.Methods {
		// return prototype
		if method.RetType.Prototype != nil {
			if err := g.ensureDelegateGenerated(method.RetType.Prototype, &sb, delegateGen); err != nil {
				return "", err
			}
			if err := g.walkPrototype(method.RetType.Prototype, onEnum, onAlias, onProto, nil); err != nil {
				return "", err
			}
		}

		// params
		for _, param := range method.ParamTypes {
			if param.Prototype != nil {
				if err := g.ensureDelegateGenerated(param.Prototype, &sb, delegateGen); err != nil {
					return "", err
				}
				if err := g.walkPrototype(param.Prototype, onEnum, onAlias, onProto, nil); err != nil {
					return "", err
				}
			}
		}
	}
	return sb.String(), nil
}
