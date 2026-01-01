package generator

import (
	"fmt"
	"strings"

	"github.com/untrustedmodders/plugify-gen/pkg/manifest"
)

// V8Generator generates TypeScript definition files for V8/JavaScript
type V8Generator struct {
	*BaseGenerator
}

// NewV8Generator creates a new V8/JavaScript generator
func NewV8Generator() *V8Generator {
	return &V8Generator{
		BaseGenerator: NewBaseGenerator("v8", NewV8TypeMapper(), V8ReservedWords),
	}
}

// Generate generates V8/JavaScript TypeScript definitions
func (g *V8Generator) Generate(m *manifest.Manifest, opts *GeneratorOptions) (*GeneratorResult, error) {
	g.ResetCaches()
	m.Sanitize(g.Sanitizer)
	opts = EnsureOptions(opts)

	var sb strings.Builder

	// Header comment
	sb.WriteString(fmt.Sprintf("// Generated from %s.pplugin\n\n", m.Name))

	// Import plugify base types
	// Add static methods and namespaces for plugify types
	sb.WriteString("declare module \"plugify\" {\n")

	sb.WriteString(g.generatePlugify())

	// Module declaration
	sb.WriteString(fmt.Sprintf("declare module \":%s\" {\n", m.Name))
	sb.WriteString("  import { Vector2, Vector3, Vector4, Matrix4x4 } from \"plugify\";\n\n")

	// Generate enums
	enumsCode, err := g.generateEnums(m)
	if err != nil {
		return nil, fmt.Errorf("generating enums: %w", err)
	}
	if enumsCode != "" {
		sb.WriteString(enumsCode)
		sb.WriteString("\n")
	}

	// Generate aliases
	aliasesCode, err := g.generateAliases(m)
	if err != nil {
		return nil, fmt.Errorf("generating aliases: %w", err)
	}
	if aliasesCode != "" {
		sb.WriteString(aliasesCode)
		sb.WriteString("\n")
	}

	// Generate delegates
	delegatesCode, err := g.generateDelegates(m)
	if err != nil {
		return nil, fmt.Errorf("generating delegates: %w", err)
	}
	if delegatesCode != "" {
		sb.WriteString(delegatesCode)
		sb.WriteString("\n")
	}

	// Generate methods
	for _, method := range m.Methods {
		methodCode, err := g.generateMethod(&method)
		if err != nil {
			return nil, fmt.Errorf("failed to generate method %s: %w", method.Name, err)
		}
		sb.WriteString(methodCode)
	}

	// Generate classes (if enabled)
	if opts.GenerateClasses && len(m.Classes) > 0 {
		classesCode, err := g.generateClasses(m)
		if err != nil {
			return nil, fmt.Errorf("generating classes: %w", err)
		}
		sb.WriteString(classesCode)
	}

	// Close module
	sb.WriteString("}\n")

	result := &GeneratorResult{
		Files: map[string]string{
			fmt.Sprintf("pps/%s.d.ts", m.Name): sb.String(),
		},
	}

	return result, nil
}

func (g *V8Generator) generatePlugify() string {
	var sb strings.Builder

	// Plugin
	sb.WriteString("  /** Represents a plugin with metadata information. */\n")
	sb.WriteString("  type Plugin = {\n")
	sb.WriteString("    /** Unique identifier for the plugin */\n")
	sb.WriteString("    id: bigint;\n")
	sb.WriteString("    /** Name of the plugin */\n")
	sb.WriteString("    name: string;\n")
	sb.WriteString("    /** Description of the plugin */\n")
	sb.WriteString("    description: string;\n")
	sb.WriteString("    /** Version of the plugin */\n")
	sb.WriteString("    version: string;\n")
	sb.WriteString("    /** Author of the plugin */\n")
	sb.WriteString("    author: string;\n")
	sb.WriteString("    /** Website of the plugin */\n")
	sb.WriteString("    website: string;\n")
	sb.WriteString("    /** Software license of the plugin */\n")
	sb.WriteString("    license: string;\n")
	sb.WriteString("    /** Installation location of the plugin */\n")
	sb.WriteString("    location: string;\n")
	sb.WriteString("    /** List of plugin dependencies */\n")
	sb.WriteString("    dependencies: string[];\n")
	sb.WriteString("    /** Base directory where plugin files reside */\n")
	sb.WriteString("    base_dir: string;\n")
	sb.WriteString("    /** Directory for plugin extensions */\n")
	sb.WriteString("    extensions_dir: string;\n")
	sb.WriteString("    /** Directory for configuration files */\n")
	sb.WriteString("    configs_dir: string;\n")
	sb.WriteString("    /** Directory for plugin data files */\n")
	sb.WriteString("    data_dir: string;\n")
	sb.WriteString("    /** Directory for log files */\n")
	sb.WriteString("    logs_dir: string;\n")
	sb.WriteString("    /** Directory for cached files */\n")
	sb.WriteString("    cache_dir: string;\n")
	sb.WriteString("  };\n\n")

	// Vector2
	sb.WriteString("  /** Represents a 2D vector with mathematical operations. */\n")
	sb.WriteString("  export type Vector2 = {\n")
	sb.WriteString("    /** X-coordinate of the vector */\n")
	sb.WriteString("    x: number;\n")
	sb.WriteString("    /** Y-coordinate of the vector */\n")
	sb.WriteString("    y: number;\n")
	sb.WriteString("    /** Adds another Vector2 to this vector */\n")
	sb.WriteString("    add(vector: Vector2): Vector2;\n")
	sb.WriteString("    /** Subtracts another Vector2 from this vector */\n")
	sb.WriteString("    subtract(vector: Vector2): Vector2;\n")
	sb.WriteString("    /** Scales this vector by a scalar */\n")
	sb.WriteString("    scale(scalar: number): Vector2;\n")
	sb.WriteString("    /** Returns the magnitude (length) of the vector */\n")
	sb.WriteString("    magnitude(): number;\n")
	sb.WriteString("    /** Returns a normalized (unit length) version of this vector */\n")
	sb.WriteString("    normalize(): Vector2;\n")
	sb.WriteString("    /** Returns the dot product with another Vector2 */\n")
	sb.WriteString("    dot(vector: Vector2): number;\n")
	sb.WriteString("    /** Computes the distance between this vector and another Vector2 */\n")
	sb.WriteString("    distanceTo(vector: Vector2): number;\n")
	sb.WriteString("    /** Converts vector to string */\n")
	sb.WriteString("    toString(): string;\n")
	sb.WriteString("  };\n\n")

	sb.WriteString("  export namespace Vector2 {\n")
	sb.WriteString("    /** Returns a zero vector (0, 0). */\n")
	sb.WriteString("    function zero(): Vector2;\n")
	sb.WriteString("    /** Returns a unit vector (1, 1). */\n")
	sb.WriteString("    function unit(): Vector2;\n")
	sb.WriteString("  }\n\n")

	// Vector3
	sb.WriteString("  /** Represents a 3D vector with mathematical operations. */\n")
	sb.WriteString("  export type Vector3 = {\n")
	sb.WriteString("    /** X-coordinate of the vector */\n")
	sb.WriteString("    x: number;\n")
	sb.WriteString("    /** Y-coordinate of the vector */\n")
	sb.WriteString("    y: number;\n")
	sb.WriteString("    /** Z-coordinate of the vector */\n")
	sb.WriteString("    z: number;\n")
	sb.WriteString("    /** Adds another Vector3 to this vector */\n")
	sb.WriteString("    add(vector: Vector3): Vector3;\n")
	sb.WriteString("    /** Subtracts another Vector3 from this vector */\n")
	sb.WriteString("    subtract(vector: Vector3): Vector3;\n")
	sb.WriteString("    /** Scales this vector by a scalar */\n")
	sb.WriteString("    scale(scalar: number): Vector3;\n")
	sb.WriteString("    /** Returns the magnitude (length) of the vector */\n")
	sb.WriteString("    magnitude(): number;\n")
	sb.WriteString("    /** Returns a normalized (unit length) version of this vector */\n")
	sb.WriteString("    normalize(): Vector3;\n")
	sb.WriteString("    /** Returns the dot product with another Vector3 */\n")
	sb.WriteString("    dot(vector: Vector3): number;\n")
	sb.WriteString("    /** Computes the cross product with another Vector3 */\n")
	sb.WriteString("    cross(vector: Vector3): Vector3;\n")
	sb.WriteString("    /** Computes the distance between this vector and another Vector3 */\n")
	sb.WriteString("    distanceTo(vector: Vector3): number;\n")
	sb.WriteString("    /** Converts vector to string */\n")
	sb.WriteString("    toString(): string;\n")
	sb.WriteString("  };\n\n")

	sb.WriteString("  export namespace Vector3 {\n")
	sb.WriteString("    /** Returns a zero vector (0, 0, 0). */\n")
	sb.WriteString("    function zero(): Vector3;\n")
	sb.WriteString("    /** Returns a unit vector (1, 1, 1). */\n")
	sb.WriteString("    function unit(): Vector3;\n")
	sb.WriteString("  }\n\n")

	// Vector4
	sb.WriteString("  /** Represents a 4D vector with mathematical operations. */\n")
	sb.WriteString("  export type Vector4 = {\n")
	sb.WriteString("    /** X-coordinate of the vector */\n")
	sb.WriteString("    x: number;\n")
	sb.WriteString("    /** Y-coordinate of the vector */\n")
	sb.WriteString("    y: number;\n")
	sb.WriteString("    /** Z-coordinate of the vector */\n")
	sb.WriteString("    z: number;\n")
	sb.WriteString("    /** W-coordinate of the vector */\n")
	sb.WriteString("    w: number;\n")
	sb.WriteString("    /** Adds another Vector4 to this vector */\n")
	sb.WriteString("    add(vector: Vector4): Vector4;\n")
	sb.WriteString("    /** Subtracts another Vector4 from this vector */\n")
	sb.WriteString("    subtract(vector: Vector4): Vector4;\n")
	sb.WriteString("    /** Scales this vector by a scalar */\n")
	sb.WriteString("    scale(scalar: number): Vector4;\n")
	sb.WriteString("    /** Returns the magnitude (length) of the vector */\n")
	sb.WriteString("    magnitude(): number;\n")
	sb.WriteString("    /** Returns a normalized (unit length) version of this vector */\n")
	sb.WriteString("    normalize(): Vector4;\n")
	sb.WriteString("    /** Returns the dot product with another Vector4 */\n")
	sb.WriteString("    dot(vector: Vector4): number;\n")
	sb.WriteString("    /** Computes the distance between this vector and another Vector4 */\n")
	sb.WriteString("    distanceTo(vector: Vector4): number;\n")
	sb.WriteString("    /** Converts vector to string */\n")
	sb.WriteString("    toString(): string;\n")
	sb.WriteString("  };\n\n")

	sb.WriteString("  export namespace Vector4 {\n")
	sb.WriteString("    /** Returns a zero vector (0, 0, 0, 0). */\n")
	sb.WriteString("    function zero(): Vector4;\n")
	sb.WriteString("    /** Returns a unit vector (1, 1, 1, 1). */\n")
	sb.WriteString("    function unit(): Vector4;\n")
	sb.WriteString("  }\n\n")

	// Matrix4x4
	sb.WriteString("  /** Represents a 4x4 matrix with transformation operations. */\n")
	sb.WriteString("  export type Matrix4x4 = {\n")
	sb.WriteString("    /** Matrix elements stored as a 2D array */\n")
	sb.WriteString("    m: number[][];\n")
	sb.WriteString("    /** Adds another matrix to this matrix */\n")
	sb.WriteString("    add(matrix: Matrix4x4): Matrix4x4;\n")
	sb.WriteString("    /** Subtracts another matrix from this matrix */\n")
	sb.WriteString("    subtract(matrix: Matrix4x4): Matrix4x4;\n")
	sb.WriteString("    /** Multiplies this matrix with another matrix */\n")
	sb.WriteString("    multiply(matrix: Matrix4x4): Matrix4x4;\n")
	sb.WriteString("    /** Multiplies this matrix with a Vector4 */\n")
	sb.WriteString("    multiplyVector(vector: Vector4): Vector4;\n")
	sb.WriteString("    /** Returns the transpose of this matrix */\n")
	sb.WriteString("    transpose(): Matrix4x4;\n")
	sb.WriteString("    /** Returns a string representation of the matrix */\n")
	sb.WriteString("    toString(): string;\n")
	sb.WriteString("  };\n\n")

	sb.WriteString("  export namespace Matrix4x4 {\n")
	sb.WriteString("    /** Returns an identity matrix. */\n")
	sb.WriteString("    function identity(): Matrix4x4;\n")
	sb.WriteString("    /** Returns a zero matrix (all values set to 0). */\n")
	sb.WriteString("    function zero(): Matrix4x4;\n")
	sb.WriteString("    /** Creates a scaling matrix with scale factors for each axis. */\n")
	sb.WriteString("    function scaling(sx: number, sy: number, sz: number): Matrix4x4;\n")
	sb.WriteString("    /** Creates a translation matrix using given translation values. */\n")
	sb.WriteString("    function translation(tx: number, ty: number, tz: number): Matrix4x4;\n")
	sb.WriteString("    /** Creates a rotation matrix around the X-axis. */\n")
	sb.WriteString("    function rotationX(angle: number): Matrix4x4;\n")
	sb.WriteString("    /** Creates a rotation matrix around the Y-axis. */\n")
	sb.WriteString("    function rotationY(angle: number): Matrix4x4;\n")
	sb.WriteString("    /** Creates a rotation matrix around the Z-axis. */\n")
	sb.WriteString("    function rotationZ(angle: number): Matrix4x4;\n")
	sb.WriteString("  }\n")
	sb.WriteString("}\n\n")

	return sb.String()
}

func (g *V8Generator) generateEnums(m *manifest.Manifest) (string, error) {
	return g.CollectEnums(m, g.generateEnum)
}

func (g *V8Generator) generateEnum(enum *manifest.Enum, underlyingType string) (string, error) {
	var sb strings.Builder

	if enum.Description != "" {
		sb.WriteString(fmt.Sprintf("  /** %s */\n", enum.Description))
	}

	sb.WriteString(fmt.Sprintf("  export const enum %s {\n", enum.Name))

	for i, val := range enum.Values {
		if val.Description != "" {
			sb.WriteString(fmt.Sprintf("    /** %s */\n", val.Description))
		}
		sb.WriteString(fmt.Sprintf("    %s = %d", val.Name, val.Value))
		if i < len(enum.Values)-1 {
			sb.WriteString(",")
		}
		sb.WriteString("\n")
	}

	sb.WriteString("  }\n")
	return sb.String(), nil
}

func (g *V8Generator) generateAliases(m *manifest.Manifest) (string, error) {
	return g.CollectAliases(m, g.generateAlias)
}

func (g *V8Generator) generateAlias(alias *manifest.Alias, underlyingType string) (string, error) {
	var sb strings.Builder

	if alias.Description != "" {
		sb.WriteString(fmt.Sprintf("  /** %s */\n", alias.Description))
	}

	sb.WriteString(fmt.Sprintf("    %s = %s;\n", alias.Name, underlyingType))

	return sb.String(), nil
}

func (g *V8Generator) generateDelegates(m *manifest.Manifest) (string, error) {
	return g.CollectDelegates(m, g.generateDelegate)
}

func (g *V8Generator) generateDelegate(proto *manifest.Prototype) (string, error) {
	var sb strings.Builder

	if proto.Description != "" {
		sb.WriteString(fmt.Sprintf("  /** %s */\n", proto.Description))
	}

	// Generate return type (with tuple for ref parameters)
	retType, err := g.generateReturnType(&proto.RetType, proto.ParamTypes)
	if err != nil {
		return "", err
	}

	// Generate parameters (TypeScript style: name: type)
	params, err := g.formatTSParameters(proto.ParamTypes)
	if err != nil {
		return "", err
	}

	sb.WriteString(fmt.Sprintf("  export type %s = (%s) => %s;\n", proto.Name, params, retType))
	return sb.String(), nil
}

func (g *V8Generator) generateClasses(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	for i, class := range m.Classes {
		if i > 0 {
			sb.WriteString("\n")
		}
		classCode, err := g.generateClass(m, &class)
		if err != nil {
			return "", fmt.Errorf("failed to generate class %s: %w", class.Name, err)
		}
		sb.WriteString(classCode)
	}

	return sb.String(), nil
}

func (g *V8Generator) generateClass(m *manifest.Manifest, class *manifest.Class) (string, error) {
	var sb strings.Builder

	hasCtor := len(class.Constructors) > 0
	hasDtor := class.Destructor != nil

	// Class JSDoc comment
	sb.WriteString("  /**\n")
	if class.Description != "" {
		sb.WriteString(fmt.Sprintf("   * %s\n", class.Description))
	}
	// Add deprecation comment if present
	if class.Deprecated != "" {
		sb.WriteString(fmt.Sprintf("   * @deprecated %s\n", class.Deprecated))
	}
	sb.WriteString("   */\n")

	// Class declaration
	sb.WriteString(fmt.Sprintf("  export class %s {\n", class.Name))

	// Generate constructors
	if hasCtor {
		for _, ctorName := range class.Constructors {
			ctorCode, err := g.generateConstructor(m, class, ctorName)
			if err != nil {
				return "", err
			}
			sb.WriteString(ctorCode)
		}
	} else {
		// Default constructor if no constructors specified
		hasDefaultConstructor := g.HasConstructorWithNoParam(m, class)
		if !hasDefaultConstructor {
			sb.WriteString("    constructor();\n\n")
		}

		// Main constructor if no constructors and destructors specified
		if !hasDtor {
			sb.WriteString("    constructor(handle);\n\n")
		}
	}

	// Generate utility methods (valid, get, release, and close if destructor exists)
	utilCode, err := g.generateUtilityMethods(class)
	if err != nil {
		return "", err
	}
	sb.WriteString(utilCode)

	// Generate bindings (methods)
	for _, binding := range class.Bindings {
		methodCode, err := g.generateBinding(m, class, &binding)
		if err != nil {
			return "", err
		}
		sb.WriteString(methodCode)
	}

	sb.WriteString("  }\n\n")

	return sb.String(), nil
}

func (g *V8Generator) generateUtilityMethods(class *manifest.Class) (string, error) {
	// Get the handle type mapped to TypeScript
	_, handleType, err := g.typeMapper.MapHandleType(class)
	if err != nil {
		return "", err
	}

	methods := []struct {
		name        string
		returnType  string
		description string
		returnDesc  string
	}{
		{"valid", "boolean", "Check if the handle is valid.", "True if the handle is valid, false otherwise"},
		{"get", handleType, "Get the raw handle value without transferring ownership.", "The underlying handle value"},
		{"release", handleType, "Release ownership of the handle and return it.", "The released handle value"},
		{"reset", "void", "Reset the handle by closing it.", ""},
	}

	var sb strings.Builder
	for _, method := range methods {
		sb.WriteString("    /**\n")
		sb.WriteString(fmt.Sprintf("     * %s\n", method.description))
		if method.returnDesc != "" {
			sb.WriteString(fmt.Sprintf("     * @returns %s\n", method.returnDesc))
		}
		sb.WriteString("     */\n")
		sb.WriteString(fmt.Sprintf("    %s(): %s;\n\n", method.name, method.returnType))
	}

	// close() method - only if destructor exists
	if class.Destructor != nil {
		sb.WriteString("    /**\n")
		sb.WriteString("     * Close and destroy the handle if owned.\n")
		sb.WriteString("     */\n")
		sb.WriteString("    close(): void;\n\n")
	}

	return sb.String(), nil
}

func (g *V8Generator) generateConstructor(m *manifest.Manifest, class *manifest.Class, methodName string) (string, error) {
	// Find the method in the manifest
	method := FindMethod(m, methodName)
	if method == nil {
		return "", fmt.Errorf("constructor method %s not found", methodName)
	}

	var sb strings.Builder

	// JSDoc comment
	sb.WriteString(g.generateJSDoc(JSDocOptions{
		Description: method.Description,
		Params:      method.ParamTypes,
		Indent:      "    ",
	}))

	// Constructor signature
	params, err := g.formatTSParameters(method.ParamTypes)
	if err != nil {
		return "", err
	}

	sb.WriteString(fmt.Sprintf("    constructor(%s);\n\n", params))

	return sb.String(), nil
}

func (g *V8Generator) generateBinding(m *manifest.Manifest, class *manifest.Class, binding *manifest.Binding) (string, error) {
	// Find the method in the manifest
	method := FindMethod(m, binding.Method)
	if method == nil {
		return "", fmt.Errorf("method %s not found", binding.Method)
	}

	var sb strings.Builder

	// Determine parameters (skip first if bindSelf)
	params := method.ParamTypes
	startIdx := 0
	if binding.BindSelf && len(params) > 0 {
		startIdx = 1
	}
	methodParams := params[startIdx:]

	// Format parameters
	formattedParams, err := g.formatTSParameters(methodParams)
	if err != nil {
		return "", err
	}

	// Apply parameter aliases (replace types with class names)
	if len(binding.ParamAliases) > 0 {
		formattedParams = g.applyParamAliases(formattedParams, methodParams, binding.ParamAliases)
	}

	// Generate return type
	retType, err := g.typeMapper.MapReturnType(&method.RetType)
	if err != nil {
		return "", err
	}

	// Handle return alias
	if binding.RetAlias != nil && binding.RetAlias.Name != "" {
		retType = binding.RetAlias.Name
	}

	// JSDoc comment
	sb.WriteString(g.generateJSDoc(JSDocOptions{
		Description: method.Description,
		Params:      methodParams,
		RetType:     &method.RetType,
		Indent:      "    ",
	}))

	// Method signature
	if !binding.BindSelf {
		sb.WriteString(fmt.Sprintf("    static %s(%s): %s;\n\n", binding.Name, formattedParams, retType))
	} else {
		sb.WriteString(fmt.Sprintf("    %s(%s): %s;\n\n", binding.Name, formattedParams, retType))
	}

	return sb.String(), nil
}

func (g *V8Generator) applyParamAliases(formattedParams string, params []manifest.ParamType, aliases []*manifest.ParamAlias) string {
	result := formattedParams
	for i, alias := range aliases {
		if i < len(params) && alias.Name != "" {
			paramName := params[i].Name

			// Find and replace the type for this parameter
			// Look for "paramName: OldType" and replace with "paramName: NewType"
			oldPattern := paramName + ": "
			idx := strings.Index(result, oldPattern)
			if idx != -1 {
				// Find the end of the type (either comma or end of string)
				start := idx + len(oldPattern)
				end := start
				for end < len(result) && result[end] != ',' {
					end++
				}
				// Replace the type
				result = result[:start] + alias.Name + result[end:]
			}
		}
	}
	return result
}

func (g *V8Generator) generateReturnType(retType *manifest.RetType, params []manifest.ParamType) (string, error) {
	// Check if any parameters are ref (out parameters)
	hasRef := false
	for _, param := range params {
		if param.Ref {
			hasRef = true
			break
		}
	}

	baseRetType, err := g.typeMapper.MapReturnType(retType)
	if err != nil {
		return "", err
	}

	if !hasRef {
		return baseRetType, nil
	}

	// Build tuple with return value + ref parameters
	types := []string{baseRetType}
	for _, param := range params {
		if param.Ref {
			paramType, err := g.typeMapper.MapParamType(&param)
			if err != nil {
				return "", err
			}
			types = append(types, paramType)
		}
	}

	return fmt.Sprintf("[%s]", strings.Join(types, ", ")), nil
}

// formatTSParameters formats parameters in TypeScript style (name: type)
func (g *V8Generator) formatTSParameters(params []manifest.ParamType) (string, error) {
	if len(params) == 0 {
		return "", nil
	}

	result := ""
	for i, param := range params {
		if i > 0 {
			result += ", "
		}

		typeName, err := g.typeMapper.MapParamType(&param)
		if err != nil {
			return "", err
		}

		paramName := param.Name

		// TypeScript style: name: type
		result += paramName + ": " + typeName
	}

	return result, nil
}

// JSDocOptions configures JSDoc comment generation for TypeScript
type JSDocOptions struct {
	Description  string
	Params       []manifest.ParamType
	ParamAliases []*manifest.ParamAlias
	RetType      *manifest.RetType
	Indent       string // "  " for top-level, "    " for class methods
}

// generateJSDoc generates JSDoc-style comments for methods
func (g *V8Generator) generateJSDoc(opts JSDocOptions) string {
	var sb strings.Builder

	sb.WriteString(opts.Indent + "/**\n")
	if opts.Description != "" {
		sb.WriteString(fmt.Sprintf("%s * %s\n", opts.Indent, opts.Description))
	}

	// Parameters section
	for i, param := range opts.Params {
		sb.WriteString(fmt.Sprintf("%s * @param %s", opts.Indent, param.Name))

		// Always append description if available (alias check is redundant)
		if param.Description != "" {
			sb.WriteString(fmt.Sprintf(" - %s", param.Description))
		}
		sb.WriteString("\n")

		// Note: Type information is in the signature, not JSDoc for TypeScript
		_ = i // Keep for potential future use with ParamAliases
	}

	// Return type section
	if opts.RetType != nil && opts.RetType.Type != "void" && opts.RetType.Description != "" {
		sb.WriteString(fmt.Sprintf("%s * @returns %s\n", opts.Indent, opts.RetType.Description))
	}
	// Add deprecation comment if present
	if method.Deprecated != "" {
		sb.WriteString(fmt.Sprintf("   * @deprecated %s\n", method.Deprecated))
	}
	sb.WriteString("   */\n")

	sb.WriteString(opts.Indent + " */\n")
	return sb.String()
}

func (g *V8Generator) generateMethod(method *manifest.Method) (string, error) {
	var sb strings.Builder

	// JSDoc comment
	sb.WriteString(g.generateJSDoc(JSDocOptions{
		Description: method.Description,
		Params:      method.ParamTypes,
		RetType:     &method.RetType,
		Indent:      "  ",
	}))

	// Generate return type (with tuple for ref parameters)
	retType, err := g.generateReturnType(&method.RetType, method.ParamTypes)
	if err != nil {
		return "", err
	}

	params, err := g.formatTSParameters(method.ParamTypes)
	if err != nil {
		return "", err
	}

	sb.WriteString(fmt.Sprintf("  export function %s(%s): %s;\n\n", method.Name, params, retType))

	return sb.String(), nil
}

// V8TypeMapper implements type mapping for V8/JavaScript
type V8TypeMapper struct{}

func NewV8TypeMapper() *V8TypeMapper {
	return &V8TypeMapper{}
}

func (m *V8TypeMapper) MapType(baseType string, context TypeContext, isArray bool) (string, error) {
	typeMap := map[string]string{
		"void":   "void",
		"bool":   "boolean",
		"char8":  "number",
		"char16": "number",
		"int8":   "number",
		"int16":  "number",
		"int32":  "number",
		"int64":  "number",
		"uint8":  "number",
		"uint16": "number",
		"uint32": "number",
		"uint64": "bigint",
		"ptr64":  "bigint",
		"float":  "number",
		"double": "number",
		"string": "string",
		"any":    "any",
		"vec2":   "Vector2",
		"vec3":   "Vector3",
		"vec4":   "Vector4",
		"mat4x4": "Matrix4x4",
	}

	mapped, ok := typeMap[baseType]
	if !ok {
		// Assume it's a custom type (enum or delegate)
		mapped = baseType
	}

	// Handle arrays
	if isArray {
		mapped = mapped + "[]"
	}

	return mapped, nil
}

func (m *V8TypeMapper) MapParamType(param *manifest.ParamType) (string, error) {
	var typeName string
	switch {
	case param.Enum != nil:
		typeName = param.Enum.Name

	case param.Alias != nil:
		typeName = param.Alias.Name

	case param.Prototype != nil:
		return param.Prototype.Name, nil

	default:
		typeName = param.BaseType()
	}

	// Regular type mapping
	return m.MapType(typeName, TypeContextValue, param.IsArray())
}

func (m *V8TypeMapper) MapReturnType(retType *manifest.RetType) (string, error) {
	var typeName string
	switch {
	case retType.Enum != nil:
		typeName = retType.Enum.Name

	case retType.Alias != nil:
		typeName = retType.Alias.Name

	case retType.Prototype != nil:
		return retType.Prototype.Name, nil

	default:
		typeName = retType.BaseType()
	}

	// Regular type mapping
	return m.MapType(typeName, TypeContextReturn, retType.IsArray())
}

func (m *V8TypeMapper) MapHandleType(class *manifest.Class) (string, string, error) {
	handleType, err := m.MapType(class.HandleType, TypeContextReturn, false)
	if err != nil {
		return "", "", err
	}
	return "null", handleType, err
}
