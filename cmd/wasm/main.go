package main

import (
	"fmt"
	"syscall/js"

	"github.com/untrustedmodders/plugify-gen/pkg/generator"
	"github.com/untrustedmodders/plugify-gen/pkg/manifest"
)

var version = "dev" // Version is set via -ldflags during build

// ConvertResult represents the result of a conversion
type ConvertResult struct {
	Success bool              `json:"success"`
	Files   map[string]string `json:"files,omitempty"`
	Error   string            `json:"error,omitempty"`
}

// convertManifest converts a manifest file content to language bindings
func convertManifest(this js.Value, args []js.Value) interface{} {
	if len(args) < 2 || len(args) > 3 {
		return map[string]interface{}{
			"success": false,
			"error":   "Expected 2-3 arguments: manifestContent (string), language (string), and optional options (object)",
		}
	}

	manifestContent := args[0].String()
	language := args[1].String()

	// Parse options if provided
	opts := &generator.GeneratorOptions{
		GenerateClasses: true, // Default to true
	}
	if len(args) >= 3 && args[2].Type() == js.TypeObject {
		if generateClasses := args[2].Get("generateClasses"); generateClasses.Type() == js.TypeBoolean {
			opts.GenerateClasses = generateClasses.Bool()
		}
	}

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
	result, err := gen.Generate(m, opts)
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
	languages := []interface{}{"cpp", "v8", "golang", "dotnet", "python", "lua", "dlang"}
	return languages
}

// getVersion returns the version
func getVersion(this js.Value, args []js.Value) interface{} {
	return version + "-wasm"
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
