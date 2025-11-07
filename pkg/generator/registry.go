package generator

import (
	"fmt"
	"sort"
	"strings"
)

var generatorRegistry = make(map[string]Generator)

// Register registers a generator for a language
func Register(gen Generator) {
	generatorRegistry[gen.Name()] = gen
}

// GetGenerator returns a generator for the specified language
func GetGenerator(language string) (Generator, error) {
	gen, ok := generatorRegistry[language]
	if !ok {
		return nil, fmt.Errorf("unsupported language: %s", language)
	}
	return gen, nil
}

// SupportedLanguages returns a comma-separated list of supported languages
func SupportedLanguages() string {
	languages := make([]string, 0, len(generatorRegistry))
	for lang := range generatorRegistry {
		languages = append(languages, lang)
	}
	sort.Strings(languages)
	return strings.Join(languages, ", ")
}

func init() {
	// Register all generators
	Register(NewCppGenerator())
	Register(NewV8Generator())
	Register(NewPythonGenerator())
	Register(NewLuaGenerator())
	Register(NewDotnetGenerator())
	Register(NewGolangGenerator())
}
