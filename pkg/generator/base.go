package generator

import (
	"fmt"

	"github.com/untrustedmodders/plugify-generator/pkg/manifest"
)

// Generator is the interface that all language-specific generators must implement
type Generator interface {
	// Name returns the generator name (e.g., "cpp", "golang")
	Name() string

	// Generate produces code from a manifest
	Generate(m *manifest.Manifest) (*GeneratorResult, error)
}

// GeneratorResult contains the generated files
type GeneratorResult struct {
	// Files maps filename to file content
	Files map[string]string
}

// BaseGenerator provides common functionality for all generators
type BaseGenerator struct {
	name            string
	typeMapper      TypeMapper
	invalidNames    map[string]bool
	enumCache       map[string]bool
	delegateCache   map[string]bool
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
}

// TypeContext represents the context in which a type appears
type TypeContext int

const (
	TypeContextValue TypeContext = iota  // Value parameter
	TypeContextRef                       // Reference parameter
	TypeContextReturn                    // Return type
	TypeContextCast                      // Cast expression
)

// ParameterFormat specifies what format parameters should be generated in
type ParameterFormat int

const (
	ParamFormatTypes ParameterFormat = iota  // Just types: "int, string"
	ParamFormatNames                         // Just names: "x, y"
	ParamFormatTypesAndNames                 // Types and names: "int x, string y"
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
		}
	}

	return result, nil
}
