# WebAssembly Integration Guide

This guide explains how to integrate the Plugify Generator WebAssembly module into your web application (Nuxt, React, Vue, vanilla JS, etc.).

## Building the WebAssembly Module

### Prerequisites

- Go 1.16 or later installed
- Set `GOROOT` environment variable (usually set automatically with Go installation)

### Build Steps

**Windows:**
```bash
build-wasm.bat
```

**Linux/Mac:**
```bash
chmod +x build-wasm.sh
./build-wasm.sh
```

This will create three files in the `dist` directory:
- `plugify-gen.wasm` - The WebAssembly binary
- `wasm_exec.js` - Go's WebAssembly runtime (needed to run the WASM module)
- `version.js` - Version information for cache busting

## Integration

### 1. Copy Files to Your Project

Copy `plugify-gen.wasm`, `wasm_exec.js`, and `version.js` to your web project's public directory:

**Nuxt 3:**
```
your-nuxt-app/
  public/
    plugify-gen.wasm
    wasm_exec.js
    version.js
```

**Next.js:**
```
your-nextjs-app/
  public/
    plugify-gen.wasm
    wasm_exec.js
    version.js
```

**Vite/Vue:**
```
your-vite-app/
  public/
    plugify-gen.wasm
    wasm_exec.js
    version.js
```

### 2. Load the WebAssembly Module

#### Nuxt 3 Example

Create a composable `composables/usePlugifyGen.ts`:

```typescript
export const usePlugifyGen = () => {
  const isReady = ref(false)
  const error = ref<string | null>(null)

  const loadWasm = async () => {
    try {
      // Load version info for cache busting
      const versionScript = document.createElement('script')
      versionScript.src = '/version.js'
      await new Promise((resolve, reject) => {
        versionScript.onload = resolve
        versionScript.onerror = reject
        document.head.appendChild(versionScript)
      })

      // Load the Go WASM runtime
      const script = document.createElement('script')
      script.src = '/wasm_exec.js'

      await new Promise((resolve, reject) => {
        script.onload = resolve
        script.onerror = reject
        document.head.appendChild(script)
      })

      // Initialize Go and load WASM with cache busting
      const go = new (window as any).Go()
      const version = (window as any).PLUGIFY_GEN_VERSION || 'dev'
      const response = await fetch(`/plugify-gen.wasm?v=${version}`)
      const buffer = await response.arrayBuffer()
      const result = await WebAssembly.instantiate(buffer, go.importObject)

      go.run(result.instance)

      // Wait for module to initialize
      await new Promise(resolve => setTimeout(resolve, 100))

      isReady.value = true
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to load WASM'
      console.error('WASM load error:', err)
    }
  }

  const convertManifest = (manifestContent: string, language: string) => {
    if (!isReady.value) {
      throw new Error('WASM module not ready')
    }

    const result = (window as any).convertManifest(manifestContent, language)
    return result
  }

  const getSupportedLanguages = (): string[] => {
    if (!isReady.value) return []
    return (window as any).getSupportedLanguages()
  }

  return {
    isReady,
    error,
    loadWasm,
    convertManifest,
    getSupportedLanguages,
  }
}
```

Use in your component:

```vue
<script setup lang="ts">
const { isReady, error, loadWasm, convertManifest } = usePlugifyGen()

onMounted(() => {
  loadWasm()
})

const handleConvert = async (file: File, language: string) => {
  const content = await file.text()
  const result = convertManifest(content, language)

  if (result.success) {
    console.log('Generated files:', result.files)
    // Handle the generated files
  } else {
    console.error('Error:', result.error)
  }
}
</script>
```

#### Vanilla JavaScript Example

```html
<!DOCTYPE html>
<html>
<head>
    <script src="version.js"></script>
    <script src="wasm_exec.js"></script>
</head>
<body>
    <input type="file" id="fileInput" accept=".pplugin">
    <select id="language">
        <option value="cpp">C++</option>
        <option value="v8">JavaScript/V8</option>
        <option value="python">Python</option>
        <option value="lua">Lua</option>
        <option value="dotnet">C#/.NET</option>
        <option value="golang">Go</option>
    </select>
    <button id="convert">Convert</button>

    <script>
        let wasmReady = false;

        async function loadWasm() {
            const go = new Go();
            // Use version for cache busting
            const version = window.PLUGIFY_GEN_VERSION || 'dev';
            const result = await WebAssembly.instantiateStreaming(
                fetch(`plugify-gen.wasm?v=${version}`),
                go.importObject
            );
            go.run(result.instance);
            await new Promise(resolve => setTimeout(resolve, 100));
            wasmReady = true;
            console.log('WASM ready (v' + version + ')');
        }

        document.getElementById('convert').addEventListener('click', async () => {
            if (!wasmReady) {
                alert('WASM not ready yet');
                return;
            }

            const file = document.getElementById('fileInput').files[0];
            const language = document.getElementById('language').value;

            const content = await file.text();
            const result = convertManifest(content, language);

            if (result.success) {
                console.log('Generated files:', result.files);
                // Download or display files
                for (const [filename, content] of Object.entries(result.files)) {
                    console.log(`${filename}:\n${content}`);
                }
            } else {
                alert('Error: ' + result.error);
            }
        });

        loadWasm();
    </script>
</body>
</html>
```

## API Reference

### Global Functions

The WebAssembly module exposes these functions to the global `window` object:

#### `convertManifest(manifestContent: string, language: string)`

Converts a manifest to language bindings.

**Parameters:**
- `manifestContent` (string): The content of the `.pplugin` manifest file
- `language` (string): Target language (`cpp`, `v8`, `python`, `lua`, `dotnet`, `golang`)

**Returns:**
```typescript
{
  success: boolean
  files?: { [filename: string]: string }  // Only present if success is true
  error?: string                          // Only present if success is false
}
```

**Example:**
```javascript
const result = convertManifest(manifestContent, 'cpp')

if (result.success) {
  // result.files = { "plugin.hpp": "...", ... }
  for (const [filename, content] of Object.entries(result.files)) {
    console.log(`${filename}: ${content.length} bytes`)
  }
} else {
  console.error(result.error)
}
```

#### `getSupportedLanguages()`

Returns list of supported languages.

**Returns:** `string[]`

**Example:**
```javascript
const languages = getSupportedLanguages()
// ['cpp', 'v8', 'python', 'lua', 'dotnet', 'golang']
```

#### `getVersion()`

Returns the version string.

**Returns:** `string`

**Example:**
```javascript
const version = getVersion()
// '1.0.0-wasm'
```

## TypeScript Types

```typescript
interface ConvertResult {
  success: boolean
  files?: Record<string, string>
  error?: string
}

declare global {
  interface Window {
    Go: any
    convertManifest: (manifestContent: string, language: string) => ConvertResult
    getSupportedLanguages: () => string[]
    getVersion: () => string
  }
}
```

## Deployment Considerations

### Cache Busting

The build process automatically generates a `version.js` file that contains the current version. This file is used to append a version query parameter to the WASM URL (e.g., `plugify-gen.wasm?v=1.0.0`), ensuring clients always download the latest version.

**How it works:**
1. Build script writes version to `dist/version.js`
2. HTML includes `<script src="version.js"></script>` before loading WASM
3. JavaScript appends `?v=${version}` when fetching the WASM file
4. When version changes, browsers bypass cache and fetch new WASM

**Important:** Always deploy all three files together:
- `plugify-gen.wasm`
- `wasm_exec.js`
- `version.js`

### MIME Types

Ensure your web server serves `.wasm` files with the correct MIME type:
```
application/wasm
```

### Netlify

Add to `netlify.toml`:
```toml
[[headers]]
  for = "/*.wasm"
  [headers.values]
    Content-Type = "application/wasm"
```

### Vercel

Add to `vercel.json`:
```json
{
  "headers": [
    {
      "source": "/(.*).wasm",
      "headers": [
        {
          "key": "Content-Type",
          "value": "application/wasm"
        }
      ]
    }
  ]
}
```

### Nginx

Add to nginx config:
```nginx
types {
    application/wasm wasm;
}
```

## File Size Optimization

The WASM file can be large. Consider:

1. **Compression**: Enable gzip/brotli compression on your server
2. **Lazy Loading**: Only load WASM when needed
3. **Build with TinyGo**: For smaller binaries (requires code adjustments)

## Browser Compatibility

WebAssembly is supported in:
- Chrome 57+
- Firefox 52+
- Safari 11+
- Edge 16+

## Troubleshooting

### "Failed to instantiate WebAssembly module"

- Ensure `wasm_exec.js` is loaded before trying to instantiate
- Check browser console for specific errors
- Verify MIME type is set correctly

### "convertManifest is not defined"

- Wait for WASM to initialize (add delay after `go.run()`)
- Check that `wasm_exec.js` is loaded

### Large file size

- The Go WASM runtime adds significant overhead (~2-3MB)
- This is normal for Go WASM applications
- Consider using compression

## Example Nuxt Project Structure

```
nuxt-app/
├── public/
│   ├── plugify-gen.wasm
│   └── wasm_exec.js
├── composables/
│   └── usePlugifyGen.ts
├── components/
│   └── ManifestConverter.vue
└── pages/
    └── index.vue
```

## Performance Notes

- Initial WASM load time: ~100-500ms (depends on file size and network)
- Conversion time: <10ms for typical manifests
- All processing happens client-side, no server needed

## License

Same license as Plugify project (GPLv3)