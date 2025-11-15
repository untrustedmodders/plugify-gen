# Plugify Header Generator

A unified, high-performance code generator that converts Plugify plugin manifests (`.pplugin`) into language-specific bindings for C, C++, Go, JavaScript/V8, .NET, Python, and Lua.

## Features

- **Single Binary**: One tool replaces all language-specific Python generators
- **Fast**: Written in Go for excellent performance
- **Extensible**: Plugin architecture makes adding new language generators easy
- **Type-Safe**: Strong type mapping with validation
- **Cross-Platform**: Works on Linux, macOS, Windows
- **WebAssembly**: Run in the browser - perfect for web-based tools

## Installation

```bash
# Build from source
go build -o plugify-gen ./cmd/plugify-gen

# Or install
go install github.com/untrustedmodders/plugify-gen/cmd/plugify-gen@latest
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

- ✅ `cpp` - C++ headers (.hpp) - **COMPLETE**
- ✅ `v8` - V8/JavaScript TypeScript definitions (.d.ts) - **COMPLETE**
- ✅ `python` - Python3 type stubs (.pyi) - **COMPLETE**
- ✅ `lua` - Lua stubs (.lua) - **COMPLETE**
- ✅ `dotnet` - .NET/C# bindings (.cs) - **COMPLETE**
- ✅ `golang` - Go bindings (.go + .h) - **COMPLETE**

## Architecture

```
plugify-gen/
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
./plugify-gen -manifest plugify-plugin-s2sdk.pplugin -output ./test_output -lang cpp -verbose
```

## Migration from Python Generators

The new generator is designed as a drop-in replacement:

**Before:**
```bash
python generator.py plugify-plugin-s2sdk.pplugin include/
```

**After:**
```bash
plugify-gen -manifest plugify-plugin-s2sdk.pplugin -output include/ -lang cpp
```

## Performance

Benchmarks show 10-50x faster generation compared to Python scripts for typical manifests.

## WebAssembly Support

### 🌐 Try it Online

**[Live Demo on GitHub Pages](https://untrustedmodders.github.io/plugify-gen/)** - Convert manifests directly in your browser!

### Build Locally

Plugify Generator can be compiled to WebAssembly for use in web applications:

```bash
# Build WebAssembly version
./build-wasm.sh      # Linux/Mac
build-wasm.bat       # Windows
```

This creates `plugify-gen.wasm` and `wasm_exec.js` in the `dist/` directory.

**Integration Example (Nuxt, React, Vue, etc.):**

```javascript
// Load WASM module
const go = new Go();
const result = await WebAssembly.instantiateStreaming(
  fetch('/plugify-gen.wasm'),
  go.importObject
);
go.run(result.instance);

// Convert manifest
const result = convertManifest(manifestContent, 'cpp');
if (result.success) {
  console.log(result.files); // { "plugin.hpp": "...", ... }
}
```

See [WASM.md](./WASM.md) for complete integration guide with Nuxt, React, and vanilla JS examples.

## License

Same license as Plugify project (GPLv3)
