package main

import (
	"fmt"
	"syscall/js"

	"github.com/untrustedmodders/plugify-generator/pkg/generator"
	"github.com/untrustedmodders/plugify-generator/pkg/manifest"
)

// ConvertResult represents the result of a conversion
type ConvertResult struct {
	Success bool              `json:"success"`
	Files   map[string]string `json:"files,omitempty"`
	Error   string            `json:"error,omitempty"`
}

// convertManifest converts a manifest file content to language bindings
func convertManifest(this js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		return map[string]interface{}{
			"success": false,
			"error":   "Expected 2 arguments: manifestContent (string) and language (string)",
		}
	}

	manifestContent := args[0].String()
	language := args[1].String()

	// Parse manifest
	m, err := manifest.Parse([]byte(manifestContent))
	if err != nil {
		return map[string]interface{}{
			"success": false,
			"error":   fmt.Sprintf("Error parsing manifest: %v", err),
		}
	}

	// Get generator for target language
	gen, err := generator.GetGenerator(language)
	if err != nil {
		return map[string]interface{}{
			"success": false,
			"error":   fmt.Sprintf("Error getting generator: %v. Supported languages: %s", err, generator.SupportedLanguages()),
		}
	}

	// Generate code
	result, err := gen.Generate(m)
	if err != nil {
		return map[string]interface{}{
			"success": false,
			"error":   fmt.Sprintf("Error generating code: %v", err),
		}
	}

	files := make(map[string]interface{}, len(result.Files))
	for k, v := range result.Files {
		files[k] = v
	}

	// Return result
	return map[string]interface{}{
		"success": true,
		"files":   files,
	}
}

// getSupportedLanguages returns list of supported languages
func getSupportedLanguages(this js.Value, args []js.Value) interface{} {
	languages := []interface{}{"cpp", "v8", "golang", "dotnet", "python", "lua"}
	return languages
}

// getVersion returns the version
func getVersion(this js.Value, args []js.Value) interface{} {
	return "1.0.0-wasm"
}

func main() {
	c := make(chan struct{}, 0)

	// Register JavaScript functions
	js.Global().Set("convertManifest", js.FuncOf(convertManifest))
	js.Global().Set("getSupportedLanguages", js.FuncOf(getSupportedLanguages))
	js.Global().Set("getVersion", js.FuncOf(getVersion))

	fmt.Println("Plugify Generator WebAssembly module loaded")
	<-c
}
