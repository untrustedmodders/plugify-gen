package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	"path/filepath"

	"github.com/untrustedmodders/plugify-generator/pkg/generator"
	"github.com/untrustedmodders/plugify-generator/pkg/manifest"
)

// version is set via ldflags during build
var version = "dev"

func main() {
	start := time.Now()

	var (
		manifestPath = flag.String("manifest", "", "Path to .pplugin manifest file (required)")
		outputDir    = flag.String("output", "", "Output directory (required)")
		language     = flag.String("lang", "", "Target language: cpp, v8, golang, dotnet, python, lua (required)")
		overwrite    = flag.Bool("overwrite", false, "Overwrite existing files")
		verbose      = flag.Bool("verbose", false, "Enable verbose output")
		showVersion  = flag.Bool("version", false, "Show version")
	)

	flag.Parse()

	if *showVersion {
		fmt.Printf("plugify-generator v%s\n", version)
		os.Exit(0)
	}

	if *manifestPath == "" || *outputDir == "" || *language == "" {
		fmt.Fprintf(os.Stderr, "Error: manifest, output, and lang are required\n\n")
		flag.Usage()
		os.Exit(1)
	}

	// Parse manifest
	if *verbose {
		fmt.Printf("Parsing manifest: %s\n", *manifestPath)
	}

	m, err := manifest.ParseFile(*manifestPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing manifest: %v\n", err)
		os.Exit(1)
	}

	if *verbose {
		fmt.Printf("Loaded plugin: %s (version %s)\n", m.Name, m.Version)
		fmt.Printf("Found %d methods\n", len(m.Methods))
	}

	// Get generator for target language
	gen, err := generator.GetGenerator(*language)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		fmt.Fprintf(os.Stderr, "Supported languages: %s\n", generator.SupportedLanguages())
		os.Exit(1)
	}

	// Create output directory
	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	// Generate code
	if *verbose {
		fmt.Printf("Generating %s bindings...\n", *language)
	}

	result, err := gen.Generate(m)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating code: %v\n", err)
		os.Exit(1)
	}

	// Write output files
	for filename, content := range result.Files {
		outputPath := filepath.Join(*outputDir, filename)

		// Check if file exists and overwrite flag
		if _, err := os.Stat(outputPath); err == nil && !*overwrite {
			fmt.Fprintf(os.Stderr, "Error: file %s already exists (use -overwrite to replace)\n", outputPath)
			os.Exit(1)
		}

		// Create parent directories if they don't exist
		if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating directory for %s: %v\n", outputPath, err)
			os.Exit(1)
		}

		if err := os.WriteFile(outputPath, []byte(content), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing file %s: %v\n", outputPath, err)
			os.Exit(1)
		}

		if *verbose {
			fmt.Printf("Generated: %s (%d bytes)\n", outputPath, len(content))
		}
	}

    elapsed := time.Since(start)
	fmt.Printf("✓ Successfully generated %s bindings in %s\n", *language, *outputDir)
	if *verbose {
	    fmt.Printf("Execution time: %s\n", elapsed)
	}
}
