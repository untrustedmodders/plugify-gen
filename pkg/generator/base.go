package generator

import (
	"fmt"
	"strings"

	"github.com/untrustedmodders/plugify-gen/pkg/manifest"
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

// SanitizeName handles reserved keywords by appending underscore
func (g *BaseGenerator) SanitizeName(name string) string {
	if g.invalidNames[name] {
		return name + "_"
	}
	return name
}

// GetGroupName handles reserved keywords by appending underscore
func (g *BaseGenerator) GetGroupName(name string) string {
	sanitized := g.SanitizeName(name)
	if sanitized != "" {
		return strings.ToLower(sanitized)
	}
	return "core"
}

// GetGroups gets the list of all groups from manifest
func (g *BaseGenerator) GetGroups(m *manifest.Manifest) map[string]bool {
	// Collect all unique groups from both methods and classes
	groups := make(map[string]bool)
	hasUngroupedMethods := false
	hasUngroupedClasses := false

	for _, method := range m.Methods {
		groupName := strings.ToLower(g.SanitizeName(method.Group))
		if groupName != "" {
			groups[groupName] = true
		} else {
			hasUngroupedMethods = true
		}
	}

	for _, class := range m.Classes {
		groupName := strings.ToLower(g.SanitizeName(class.Group))
		if groupName != "" {
			groups[groupName] = true
		} else {
			hasUngroupedClasses = true
		}
	}

	// Add default group for methods/classes without a group
	if hasUngroupedMethods || hasUngroupedClasses {
		groups["core"] = true
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

// ParameterFormat specifies what format parameters should be generated in
type ParameterFormat int

const (
	ParamFormatTypes         ParameterFormat = iota // Just types: "int, string"
	ParamFormatNames                                // Just names: "x, y"
	ParamFormatTypesAndNames                        // Types and names: "int x, string y"
)

// FormatParameters is a helper to format method parameters
func FormatParameters(params []manifest.ParamType, format ParameterFormat, mapper TypeMapper, sanitize func(string) string) (string, error) {
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
			if paramName == "" {
				paramName = fmt.Sprintf("p%d", i)
			}
			result += sanitize(paramName)

		case ParamFormatTypesAndNames:
			typeName, err := mapper.MapParamType(&param, TypeContextValue)
			if err != nil {
				return "", err
			}
			paramName := param.Name
			if paramName == "" {
				paramName = fmt.Sprintf("p%d", i)
			}
			result += typeName + " " + sanitize(paramName)
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
