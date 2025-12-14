/**
 * TypeScript type definitions for Plugify Generator WebAssembly module
 *
 * Copy this file to your Nuxt/TypeScript project for type safety.
 *
 * Usage in Nuxt:
 * 1. Copy to your project (e.g., `types/plugify-gen.d.ts`)
 * 2. Types will be automatically picked up by TypeScript
 */

/**
 * Result from converting a manifest to language bindings
 */
export interface ConvertResult {
  /** Whether the conversion was successful */
  success: boolean

  /** Generated files (only present when success is true) */
  files?: Record<string, string>

  /** Error message (only present when success is false) */
  error?: string
}

/**
 * Supported target languages
 */
export type SupportedLanguage = 'cpp' | 'cxx' | 'v8' | 'python' | 'lua' | 'dotnet' | 'golang' | 'dlang' | 'rust'

/**
 * Global functions exposed by the Plugify Generator WASM module
 */
declare global {
  interface Window {
    /**
     * Go WebAssembly runtime
     */
    Go: any

    /**
     * Convert a manifest file to language bindings
     *
     * @param manifestContent - The content of the .pplugin manifest file
     * @param language - Target language (cpp, cxx, v8, python, lua, dotnet, golang, dlang, rust)
     * @returns Conversion result with generated files or error message
     *
     * @example
     * ```typescript
     * const result = convertManifest(manifestContent, 'cpp')
     * if (result.success) {
     *   for (const [filename, content] of Object.entries(result.files)) {
     *     console.log(`${filename}: ${content.length} bytes`)
     *   }
     * } else {
     *   console.error(result.error)
     * }
     * ```
     */
    convertManifest(manifestContent: string, language: SupportedLanguage): ConvertResult

    /**
     * Get list of supported target languages
     *
     * @returns Array of supported language identifiers
     *
     * @example
     * ```typescript
     * const languages = getSupportedLanguages()
     * // ['cpp', 'cxx', 'v8', 'python', 'lua', 'dotnet', 'golang', 'rust'. 'dlang']
     * ```
     */
    getSupportedLanguages(): SupportedLanguage[]

    /**
     * Get the version of the WASM module
     *
     * @returns Version string
     *
     * @example
     * ```typescript
     * const version = getVersion()
     * // '1.0.0-wasm'
     * ```
     */
    getVersion(): string
  }
}

export {}