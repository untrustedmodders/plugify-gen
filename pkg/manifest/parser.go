package manifest

import (
	"encoding/json"
	"fmt"
	"os"
)

// ParseFile parses a .pplugin manifest file
func ParseFile(path string) (*Manifest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read manifest file: %w", err)
	}

	return Parse(data)
}

// Parse parses manifest JSON data
func Parse(data []byte) (*Manifest, error) {
	var m Manifest
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, fmt.Errorf("failed to parse manifest JSON: %w", err)
	}

	if err := validate(&m); err != nil {
		return nil, fmt.Errorf("manifest validation failed: %w", err)
	}

	return &m, nil
}

// validate performs basic validation on the manifest
func validate(m *Manifest) error {
	if m.Name == "" {
		return fmt.Errorf("manifest name is required")
	}
	if m.Version == "" {
		return fmt.Errorf("manifest version is required")
	}
	if m.Language == "" {
		return fmt.Errorf("manifest language is required")
	}

	// Validate methods
	for i, method := range m.Methods {
		if method.Name == "" {
			return fmt.Errorf("method[%d]: name is required", i)
		}
		if method.FuncName == "" {
			return fmt.Errorf("method[%d]: funcName is required", i)
		}

		// Validate parameters
		for j, param := range method.ParamTypes {
			if param.Name == "" {
				return fmt.Errorf("method[%d].param[%d]: name is required", i, j)
			}
			if param.Type == "" {
				return fmt.Errorf("method[%d].param[%d]: type is required", i, j)
			}
		}

		// Validate return type
		if method.RetType.Type == "" {
			return fmt.Errorf("method[%d]: retType.type is required", i)
		}
	}

	return nil
}
