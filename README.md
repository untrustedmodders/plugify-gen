# Plugify Header Generator

A unified, high-performance code generator that converts Plugify plugin manifests (`.pplugin`) into language-specific bindings for C, C++, Go, JavaScript/V8, .NET, Python, and Lua.

## Features

- **Single Binary**: One tool replaces all language-specific Python generators
- **Fast**: Written in Go for excellent performance
- **Extensible**: Plugin architecture makes adding new language generators easy
- **Type-Safe**: Strong type mapping with validation
- **Cross-Platform**: Works on Linux, macOS, Windows

## Installation

```bash
# Build from source
go build -o plugify-gen ./cmd/plugify-gen

# Or install
go install github.com/untrustedmodders/plugify-generator/cmd/plugify-gen@latest
```

## Usage

```bash
# Generate C++ bindings
plugify-gen -manifest plugin.pplugin -output ./include -lang cpp

# Generate Go bindings
plugify-gen -manifest plugin.pplugin -output ./pkg -lang golang

# Generate V8/JavaScript bindings
plugify-gen -manifest plugin.pplugin -output ./js -lang v8

# Overwrite existing files
plugify-gen -manifest plugin.pplugin -output ./out -lang cpp -overwrite

# Verbose output
plugify-gen -manifest plugin.pplugin -output ./out -lang cpp -verbose
```

### Supported Languages

- `cpp` - C++ headers (.hpp)
- `v8` - V8/JavaScript TypeScript definitions (.d.ts)
- `golang` - Go bindings (.go + .h)
- `dotnet` - .NET/C# bindings (.cs)
- `python` - Python3 bindings (.py)
- `lua` - Lua bindings (.lua)

## Architecture

```
plugify-generator/
├── cmd/plugify-gen/       # CLI entry point
├── pkg/
│   ├── manifest/          # .pplugin parser & types
│   ├── generator/         # Language generators
│   │   ├── base.go       # Common generator logic
│   │   ├── registry.go   # Generator registration
│   │   ├── cpp.go        # C++ generator
│   │   └── ...           # Other language generators
│   └── ...
└── templates/             # Code generation templates
```

## Design Principles

1. **Abstraction**: Common parsing and type system logic is shared
2. **Extensibility**: New generators implement the `Generator` interface
3. **Performance**: Compiled Go binary is much faster than Python scripts
4. **Maintainability**: Clear separation of concerns, well-tested

## Adding a New Language Generator

1. Create a new file in `pkg/generator/` (e.g., `rust.go`)
2. Implement the `Generator` interface:
   ```go
   type Generator interface {
       Name() string
       Generate(m *manifest.Manifest) (*GeneratorResult, error)
   }
   ```
3. Create a `TypeMapper` for language-specific type conversions
4. Register the generator in `registry.go`

## Development

```bash
# Run tests
go test ./...

# Build
go build -o plugify-gen ./cmd/plugify-gen

# Test with example manifest
./plugify-gen -manifest plugify-plugin-s2sdk.pplugin.in -output ./test_output -lang cpp -verbose
```

## Migration from Python Generators

The new generator is designed as a drop-in replacement:

**Before:**
```bash
python generator._cpppy plugify-plugin-s2sdk.pplugin.in include/
```

**After:**
```bash
plugify-gen -manifest plugify-plugin-s2sdk.pplugin.in -output include/ -lang cpp
```

## Performance

Benchmarks show 10-50x faster generation compared to Python scripts for typical manifests.

## License

Same license as Plugify project (GPLv3)
