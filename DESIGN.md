# Design Document: Plugify Unified Generator

## Overview

The Plugify Unified Generator is a rewrite of the original Python-based header generators into a single, high-performance Go application. It replaces 6+ separate Python scripts with one extensible binary.

## Architecture

### 1. Core Components

```
┌─────────────────────────────────────────────────┐
│              CLI (main.go)                      │
│  - Flag parsing                                 │
│  - File I/O                                     │
│  - Error handling                               │
└─────────────────┬───────────────────────────────┘
                  │
                  ▼
┌─────────────────────────────────────────────────┐
│          Manifest Parser                        │
│  - JSON deserialization                         │
│  - Schema validation                            │
│  - Type definitions                             │
└─────────────────┬───────────────────────────────┘
                  │
                  ▼
┌─────────────────────────────────────────────────┐
│        Generator Registry                       │
│  - Plugin system for generators                 │
│  - Dynamic dispatch                             │
└─────────────────┬───────────────────────────────┘
                  │
      ┌───────────┴────────────┬─────────────┐
      ▼                        ▼             ▼
┌──────────────┐      ┌──────────────┐   ┌──────────────┐
│ C++ Generator│      │ V8 Generator │   │ Go Generator │
└──────────────┘      └──────────────┘   └──────────────┘
      │                        │                  │
      └────────────┬───────────┴──────────────────┘
                   │
                   ▼
          ┌─────────────────┐
          │  Type Mapper    │
          │  Interface      │
          └─────────────────┘
```

### 2. Key Design Patterns

#### **Strategy Pattern: Type Mapping**

Each language has different type representations. The `TypeMapper` interface abstracts this:

```go
type TypeMapper interface {
    MapType(baseType string, context TypeContext, isArray bool) (string, error)
    MapParamType(param *ParamType, context TypeContext) (string, error)
    MapReturnType(retType *TypeInfo) (string, error)
}
```

**Example:**
- `string` in manifest →
  - C++: `const plg::string&` (ref context) or `plg::string` (value)
  - V8: `string`
  - Go: `string` (public API) + `String192*` (C FFI)

#### **Registry Pattern: Generator Management**

Generators self-register on initialization:

```go
func init() {
    Register(NewCppGenerator())
    Register(NewV8Generator())
    // ... more generators
}
```

This allows:
- **Zero coupling**: Main doesn't know about specific generators
- **Easy extension**: Add new generator = drop in new file
- **Dynamic loading**: Could support plugins in the future

#### **Template Method Pattern: BaseGenerator**

Common logic (name sanitization, caching, utilities) lives in `BaseGenerator`:

```go
type BaseGenerator struct {
    name            string
    typeMapper      TypeMapper
    invalidNames    map[string]struct{}
    enumCache       map[string]struct{}
    delegateCache   map[string]struct{}
    aliasCache      map[string]struct{}
}
```

Language-specific generators embed and extend this base.

### 3. Type System

#### **Type Contexts**

Types are rendered differently based on context:

```go
type TypeContext int

const (
    TypeContextValue    // Function parameter by value
    TypeContextRef      // Function parameter by reference
    TypeContextReturn   // Return type
    TypeContextCast     // Cast/marshaling expression
)
```

**C++ Example:**
```cpp
// TypeContextValue
void Foo(plg::string str);

// TypeContextRef
void Foo(const plg::string& str);

// TypeContextReturn
plg::string Foo();
```

#### **Array Handling**

Arrays are detected via `[]` suffix in manifest:
- Manifest: `"int32[]"`
- C++: `plg::vector<int32_t>`
- V8: `number[]`
- Go: `[]int32`

Helper methods on types:
```go
func (t *TypeInfo) IsArray() bool
func (t *TypeInfo) BaseType() string
```

### 4. Code Generation Pipeline

```
Manifest (.pplugin)
    ↓
Parse JSON
    ↓
Validate Schema
    ↓
Extract Enums/Delegates (deduplicate)
    ↓
Generate Type Definitions
    ↓
For Each Method:
    ├─ Generate Documentation
    ├─ Map Parameter Types
    ├─ Map Return Type
    ├─ Generate Implementation
    └─ Sanitize Names
    ↓
Format & Write Output
```

### 5. Generator Lifecycle

Each generator implements:

```go
type Generator interface {
    Name() string
    Generate(m *Manifest) (*GeneratorResult, error)
}
```

**Typical Generate() flow:**

1. **Reset caches** (enums, delegates)
2. **Write header/preamble**
3. **First pass: collect enums/delegates** from all methods
4. **Generate type definitions** (deduplicated)
5. **Generate methods** one by one
6. **Write footer**
7. **Return** map of filename → content

### 6. Performance Optimizations

#### **String Building**
Uses `strings.Builder` for efficient concatenation (no allocations per append).

#### **Single-Pass Generation**
One iteration through methods, caching enums/delegates on the fly.

#### **No Reflection**
Direct struct access, no runtime reflection overhead.

#### **Compiled Binary**
Go compiles to native code → 10-50x faster than Python interpreter.

## Adding a New Language Generator

### Step-by-Step Guide

**1. Create generator file:** `pkg/generator/mylang.go`

```go
package generator

type MyLangGenerator struct {
    *BaseGenerator
}

func NewMyLangGenerator() *MyLangGenerator {
    invalidNames := []string{"class", "def", "return", ...}
    return &MyLangGenerator{
        BaseGenerator: NewBaseGenerator("mylang", NewMyLangTypeMapper(), invalidNames),
    }
}

func (g *MyLangGenerator) Generate(m *manifest.Manifest) (*GeneratorResult, error) {
    g.ResetCaches()

    // Your generation logic here

    return &GeneratorResult{
        Files: map[string]string{
            "output.ext": "generated code",
        },
    }, nil
}
```

**2. Implement TypeMapper:**

```go
type MyLangTypeMapper struct{}

func NewMyLangTypeMapper() *MyLangTypeMapper {
    return &MyLangTypeMapper{}
}

func (m *MyLangTypeMapper) MapType(baseType string, context TypeContext, isArray bool) (string, error) {
    typeMap := map[string]string{
        "int32":  "i32",
        "string": "str",
        // ... map all types
    }

    mapped := typeMap[baseType]
    if isArray && context&TypeContextAlias == 0 {
        mapped = "[]" + mapped  // Adjust for your language
    }
    return mapped, nil
}

func (m *MyLangTypeMapper) MapParamType(param *manifest.ParamType) (string, error) {
    return m.MapType(param.BaseType(), context, param.IsArray())
}

func (m *MyLangTypeMapper) MapReturnType(retType *manifest.TypeInfo) (string, error) {
    return m.MapType(retType.BaseType(), TypeContextReturn, retType.IsArray())
}
```

**3. Register in `registry.go`:**

```go
func init() {
    Register(NewCppGenerator())
    Register(NewV8Generator())
    Register(NewMyLangGenerator())  // Add this
}
```

**4. Test:**

```bash
make build
./plugify-gen -manifest test.pplugin -output ./out -lang mylang -verbose
```

## Comparison with Python Implementation

### Python (generator._cpppy)
- **Pros:**
  - Quick prototyping
  - Dynamic typing for flexibility
- **Cons:**
  - Slow (interpreted)
  - No type safety
  - 6+ separate scripts with duplicated logic
  - Hard to maintain

### Go (Unified Generator)
- **Pros:**
  - 10-50x faster execution
  - Type-safe at compile time
  - Single binary for all languages
  - Shared common logic
  - Easy to extend
  - Cross-platform
- **Cons:**
  - More verbose (explicit type declarations)
  - Requires compilation step

## Future Enhancements

### 1. Template-Based Generation
Move from hardcoded string building to `text/template`:

```go
const cppTemplate = `
#pragma once
{{range .Enums}}
enum class {{.Name}} : {{.Type}} {
  {{range .Values}}{{.Name}} = {{.Value}},{{end}}
};
{{end}}
`
```

### 2. Configuration File
Support `.plugify-gen.yaml` for project-specific settings:

```yaml
generators:
  cpp:
    output: include/
    namespace_prefix: "plg::"
  v8:
    output: typings/
```

### 3. Incremental Generation
Only regenerate if manifest changed (check hash/mtime).

### 4. Plugin System
Load generators from external shared libraries.

### 5. IDE Integration
- LSP server for `.pplugin` files
- Syntax highlighting
- Auto-completion

## Testing Strategy

### Unit Tests
- Manifest parsing
- Type mapping
- Name sanitization

### Integration Tests
- Full generation from sample manifests
- Output comparison with Python generators

### Regression Tests
- Generate from all existing `.pplugin` files
- Compare with Python output (byte-for-byte if possible)

### Performance Tests
- Benchmark against Python
- Memory profiling

## Migration Path

### Phase 1: Parallel Deployment (Current)
- Keep Python generators
- Deploy Go generator alongside
- Compare outputs

### Phase 2: Switch to Go
- Update CI/CD to use Go generator
- Keep Python as fallback

### Phase 3: Deprecation
- Remove Python generators
- Archive for reference

## Conclusion

The unified generator provides:
1. **Better performance** (10-50x speedup)
2. **Easier maintenance** (shared logic)
3. **Type safety** (catch errors at compile time)
4. **Extensibility** (plugin-based architecture)

While requiring more upfront design, the long-term benefits outweigh the initial investment.
