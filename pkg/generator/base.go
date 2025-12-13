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
	NullPtrError     = "null pointer"
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

// BaseGenerator provides common functionality for all generators
type BaseGenerator struct {
	name          string
	typeMapper    TypeMapper
	invalidNames  map[string]bool
	enumCache     map[string]bool
	delegateCache map[string]bool
}

// NewBaseGenerator creates a new base generator
func NewBaseGenerator(name string, typeMapper TypeMapper, invalidNames []string) *BaseGenerator {
	invalidMap := make(map[string]bool)
	for _, name := range invalidNames {
		invalidMap[name] = true
	}

	return &BaseGenerator{
		name:          name,
		typeMapper:    typeMapper,
		invalidNames:  invalidMap,
		enumCache:     make(map[string]bool),
		delegateCache: make(map[string]bool),
	}
}

// Name returns the generator name
func (g *BaseGenerator) Name() string {
	return g.name
}

// Sanitizer handles reserved keywords by appending underscore
func (g *BaseGenerator) Sanitizer(name string) string {
	if g.invalidNames[name] {
		return name + "_"
	}
	return name
}

// GetGroups gets the list of all groups from manifest
func (g *BaseGenerator) GetGroups(m *manifest.Manifest) map[string]bool {
	// Collect all unique groups from both methods and classes
	groups := make(map[string]bool)
	hasUngroupedMethods := false
	hasUngroupedClasses := false

	for _, method := range m.Methods {
		groupName := strings.ToLower(method.Group)
		if groupName != "" {
			groups[groupName] = true
		} else {
			hasUngroupedMethods = true
		}
	}

	for _, class := range m.Classes {
		groupName := strings.ToLower(class.Group)
		if groupName != "" {
			groups[groupName] = true
		} else {
			hasUngroupedClasses = true
		}
	}

	// Add default group for methods/classes without a group
	if hasUngroupedMethods || hasUngroupedClasses {
		groups[DefaultGroupName] = true
	}

	return groups
}

// IsEnumCached checks if an enum has already been generated
func (g *BaseGenerator) IsEnumCached(name string) bool {
	return g.enumCache[name]
}

// CacheEnum marks an enum as generated
func (g *BaseGenerator) CacheEnum(name string) {
	g.enumCache[name] = true
}

// IsDelegateCached checks if a delegate has already been generated
func (g *BaseGenerator) IsDelegateCached(name string) bool {
	return g.delegateCache[name]
}

// CacheDelegate marks a delegate as generated
func (g *BaseGenerator) CacheDelegate(name string) {
	g.delegateCache[name] = true
}

// ResetCaches clears enum and delegate caches (call before generation)
func (g *BaseGenerator) ResetCaches() {
	g.enumCache = make(map[string]bool)
	g.delegateCache = make(map[string]bool)
}

// TypeMapper is the interface for type mapping strategies
type TypeMapper interface {
	// MapType converts a manifest type to a target language type
	MapType(baseType string, context TypeContext, isArray bool) (string, error)

	// MapParamType converts a parameter type (handles ref, default values)
	MapParamType(param *manifest.ParamType, context TypeContext) (string, error)

	// MapReturnType converts a return type
	MapReturnType(retType *manifest.TypeInfo) (string, error)

	// MapHandleType converts a invalid type
	MapHandleType(class *manifest.Class) (string, string, error)
}

// TypeContext represents the context in which a type appears
type TypeContext int

const (
	TypeContextValue  TypeContext = iota // Value parameter
	TypeContextRef                       // Reference parameter
	TypeContextReturn                    // Return type
	TypeContextCast                      // Cast expression
)

// EnumGenerator is a callback function that generates code for an enum type
// It takes an enum and returns the generated code as a string
type EnumGenerator func(*manifest.EnumType, string) (string, error)

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
			typeName, err := mapper.MapParamType(&param, TypeContextValue)
			if err != nil {
				return "", err
			}
			result += typeName

		case ParamFormatNames:
			paramName := param.Name
			result += paramName

		case ParamFormatTypesAndNames:
			typeName, err := mapper.MapParamType(&param, TypeContextValue)
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
func (g *BaseGenerator) FindDependentGroups(m *manifest.Manifest, currentGroupName string) map[string]bool {
	groupName := currentGroupName
	// Pre-index methods for O(1) lookup.
	methodToGroup := make(map[string]string, len(m.Methods))
	for i := range m.Methods {
		method := &m.Methods[i]
		methodGroup := method.Group

		methodToGroup[method.Name] = methodGroup
		//methodToGroup[method.FuncName] = methodGroup
	}

	referencedGroups := make(map[string]bool)

	for _, class := range m.Classes {
		classGroup := class.Group
		if classGroup != groupName {
			continue
		}

		for _, ctorName := range class.Constructors {
			if refGroup, ok := methodToGroup[ctorName]; ok && refGroup != classGroup {
				referencedGroups[refGroup] = true
			}
		}

		if class.Destructor != nil {
			if refGroup, ok := methodToGroup[*class.Destructor]; ok && refGroup != classGroup {
				referencedGroups[refGroup] = true
			}
		}

		for _, binding := range class.Bindings {
			if refGroup, ok := methodToGroup[binding.Method]; ok && refGroup != classGroup {
				referencedGroups[refGroup] = true
			}
		}
	}

	return referencedGroups
}

// ProcessPrototypeEnums recursively processes a prototype to find and generate all enum types.
// This helper eliminates duplicated enum processing logic across all generators.
//
// Parameters:
//   - proto: The prototype to process
//   - sb: String builder to write generated enum code to
//   - enumGen: Callback function that generates code for an enum type
//
// The function:
//  1. Checks the return type for enums
//  2. Recursively processes nested prototypes in the return type
//  3. Checks all parameters for enums
//  4. Recursively processes nested prototypes in parameters
//
// Enums are automatically cached to prevent duplicates.
func (g *BaseGenerator) ProcessPrototypeEnums(proto *manifest.Prototype, sb *strings.Builder, enumGen EnumGenerator) error {
	// Check return type for enum
	if proto.RetType.Enum != nil && !g.IsEnumCached(proto.RetType.Enum.Name) {
		enumType, err := g.typeMapper.MapType(proto.RetType.Type, TypeContextReturn, false)
		if err != nil {
			return err
		}
		enumCode, err := enumGen(proto.RetType.Enum, enumType)
		if err != nil {
			return err
		}
		sb.WriteString(enumCode)
		sb.WriteString("\n")
		g.CacheEnum(proto.RetType.Enum.Name)
	}

	// Recursively process return type prototype
	if proto.RetType.Prototype != nil {
		err := g.ProcessPrototypeEnums(proto.RetType.Prototype, sb, enumGen)
		if err != nil {
			return err
		}
	}

	// Check parameters for enums and nested prototypes
	for _, param := range proto.ParamTypes {
		if param.Enum != nil && !g.IsEnumCached(param.Enum.Name) {
			enumType, err := g.typeMapper.MapType(param.Type, TypeContextReturn, false)
			if err != nil {
				return err
			}
			enumCode, err := enumGen(param.Enum, enumType)
			if err != nil {
				return err
			}
			sb.WriteString(enumCode)
			sb.WriteString("\n")
			g.CacheEnum(param.Enum.Name)
		}
		if param.Prototype != nil {
			err := g.ProcessPrototypeEnums(param.Prototype, sb, enumGen)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// CollectEnums collects and generates all enum types from methods in a manifest.
// This helper eliminates duplicated enum collection logic across all generators.
//
// Parameters:
//   - m: The manifest containing methods to process
//   - enumGen: Callback function that generates code for an enum type
//
// Returns:
//   - Generated code for all enums as a string
//
// The function processes:
//  1. Method return types for enums
//  2. Method return type prototypes (recursively)
//  3. Method parameters for enums
//  4. Method parameter prototypes (recursively)
//
// Enums are automatically cached to prevent duplicates.
func (g *BaseGenerator) CollectEnums(m *manifest.Manifest, enumGen EnumGenerator) (string, error) {
	var sb strings.Builder

	for _, method := range m.Methods {
		// Check return type for enum
		if method.RetType.Enum != nil && !g.IsEnumCached(method.RetType.Enum.Name) {
			enumType, err := g.typeMapper.MapType(method.RetType.Type, TypeContextReturn, false)
			if err != nil {
				return "", err
			}
			enumCode, err := enumGen(method.RetType.Enum, enumType)
			if err != nil {
				return "", err
			}
			sb.WriteString(enumCode)
			sb.WriteString("\n")
			g.CacheEnum(method.RetType.Enum.Name)
		}

		// Process return type prototype
		if method.RetType.Prototype != nil {
			err := g.ProcessPrototypeEnums(method.RetType.Prototype, &sb, enumGen)
			if err != nil {
				return "", err
			}
		}

		// Check each parameter for enum
		for _, param := range method.ParamTypes {
			if param.Enum != nil && !g.IsEnumCached(param.Enum.Name) {
				enumType, err := g.typeMapper.MapType(param.Type, TypeContextReturn, false)
				if err != nil {
					return "", err
				}
				enumCode, err := enumGen(param.Enum, enumType)
				if err != nil {
					return "", err
				}
				sb.WriteString(enumCode)
				sb.WriteString("\n")
				g.CacheEnum(param.Enum.Name)
			}

			// Process nested prototypes
			if param.Prototype != nil {
				err := g.ProcessPrototypeEnums(param.Prototype, &sb, enumGen)
				if err != nil {
					return "", err
				}
			}
		}
	}

	return sb.String(), nil
}

func (g *BaseGenerator) CollectDelegates(m *manifest.Manifest, delegateGenerator DelegateGenerator) (string, error) {
	var sb strings.Builder

	for _, method := range m.Methods {
		// Check return type
		if method.RetType.Prototype != nil && !g.IsDelegateCached(method.RetType.Prototype.Name) {
			delegateCode, err := delegateGenerator(method.RetType.Prototype)
			if err != nil {
				return "", err
			}
			sb.WriteString(delegateCode)
			g.CacheDelegate(method.RetType.Prototype.Name)
		}

		// Check parameters
		for _, param := range method.ParamTypes {
			if param.Prototype != nil && !g.IsDelegateCached(param.Prototype.Name) {
				delegateCode, err := delegateGenerator(param.Prototype)
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
