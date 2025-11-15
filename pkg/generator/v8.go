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
	invalidNames := []string{
		"abstract", "arguments", "await", "boolean", "break", "byte", "case",
		"catch", "char", "class", "const", "continue", "debugger", "default",
		"delete", "do", "double", "else", "enum", "eval", "export", "extends",
		"false", "final", "finally", "float", "for", "function", "goto", "if",
		"implements", "import", "in", "instanceof", "int", "interface", "let",
		"long", "native", "new", "null", "package", "private", "protected",
		"public", "return", "short", "static", "super", "switch", "synchronized",
		"this", "throw", "throws", "transient", "true", "try", "typeof", "var",
		"void", "volatile", "while", "with", "yield",
	}

	return &V8Generator{
		BaseGenerator: NewBaseGenerator("v8", NewV8TypeMapper(), invalidNames),
	}
}

// Generate generates V8/JavaScript TypeScript definitions
func (g *V8Generator) Generate(m *manifest.Manifest) (*GeneratorResult, error) {
	g.ResetCaches()

	var sb strings.Builder

	// Header comment
	sb.WriteString(fmt.Sprintf("// Generated from %s.pplugin\n\n", m.Name))

	// Import plugify base types
	// Add static methods and namespaces for plugify types
	sb.WriteString("declare module \"plugify\" {\n")

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

	// Module declaration
	sb.WriteString(fmt.Sprintf("declare module \":%s\" {\n", m.Name))
	sb.WriteString("  import { Vector2, Vector3, Vector4, Matrix4x4 } from \"plugify\";\n\n")

	// Generate enums and delegates
	enumsCode, err := g.generateEnumsAndDelegates(m)
	if err != nil {
		return nil, err
	}
	sb.WriteString(enumsCode)

	// Generate methods
	for _, method := range m.Methods {
		methodCode, err := g.generateMethod(&method)
		if err != nil {
			return nil, fmt.Errorf("failed to generate method %s: %w", method.Name, err)
		}
		sb.WriteString(methodCode)
	}

	// Close module
	sb.WriteString("}\n")

	result := &GeneratorResult{
		Files: map[string]string{
			fmt.Sprintf("%s.d.ts", m.Name): sb.String(),
		},
	}

	return result, nil
}

func (g *V8Generator) generateEnumsAndDelegates(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	for _, method := range m.Methods {
		// Check return type
		if method.RetType.Enum != nil && !g.IsEnumCached(method.RetType.Enum.Name) {
			enumCode := g.generateEnum(method.RetType.Enum)
			sb.WriteString(enumCode)
			g.CacheEnum(method.RetType.Enum.Name)
		}
		if method.RetType.Prototype != nil && !g.IsDelegateCached(method.RetType.Prototype.Name) {
			delegateCode, err := g.generateDelegate(method.RetType.Prototype)
			if err != nil {
				return "", err
			}
			sb.WriteString(delegateCode)
			g.CacheDelegate(method.RetType.Prototype.Name)
		}

		// Check parameters
		for _, param := range method.ParamTypes {
			if param.Enum != nil && !g.IsEnumCached(param.Enum.Name) {
				enumCode := g.generateEnum(param.Enum)
				sb.WriteString(enumCode)
				g.CacheEnum(param.Enum.Name)
			}
			if param.Prototype != nil && !g.IsDelegateCached(param.Prototype.Name) {
				delegateCode, err := g.generateDelegate(param.Prototype)
				if err != nil {
					return "", err
				}
				sb.WriteString(delegateCode)
				g.CacheDelegate(param.Prototype.Name)
			}
		}
	}

	if sb.Len() > 0 {
		sb.WriteString("\n")
	}

	return sb.String(), nil
}

func (g *V8Generator) generateEnum(enum *manifest.EnumType) string {
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

	sb.WriteString("  }\n\n")
	return sb.String()
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

	sb.WriteString(fmt.Sprintf("  export type %s = (%s) => %s;\n\n", proto.Name, params, retType))
	return sb.String(), nil
}

func (g *V8Generator) generateReturnType(retType *manifest.TypeInfo, params []manifest.ParamType) (string, error) {
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
			paramType, err := g.typeMapper.MapParamType(&param, TypeContextValue)
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

		typeName, err := g.typeMapper.MapParamType(&param, TypeContextValue)
		if err != nil {
			return "", err
		}

		paramName := param.Name
		if paramName == "" {
			paramName = fmt.Sprintf("p%d", i)
		}

		// TypeScript style: name: type
		result += g.SanitizeName(paramName) + ": " + typeName
	}

	return result, nil
}

func (g *V8Generator) generateMethod(method *manifest.Method) (string, error) {
	var sb strings.Builder

	// JSDoc comment
	sb.WriteString("  /**\n")
	if method.Description != "" {
		sb.WriteString(fmt.Sprintf("   * %s\n", method.Description))
	}
	for _, param := range method.ParamTypes {
		sb.WriteString(fmt.Sprintf("   * @param %s", g.SanitizeName(param.Name)))
		if param.Description != "" {
			sb.WriteString(fmt.Sprintf(" - %s", param.Description))
		}
		sb.WriteString("\n")
	}
	if method.RetType.Type != "void" && method.RetType.Description != "" {
		sb.WriteString(fmt.Sprintf("   * @returns %s\n", method.RetType.Description))
	}
	sb.WriteString("   */\n")

	// Generate return type (with tuple for ref parameters)
	retType, err := g.generateReturnType(&method.RetType, method.ParamTypes)
	if err != nil {
		return "", err
	}

	params, err := g.formatTSParameters(method.ParamTypes)
	if err != nil {
		return "", err
	}

	sb.WriteString(fmt.Sprintf("  export function %s(%s): %s;\n\n", g.SanitizeName(method.Name), params, retType))

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

func (m *V8TypeMapper) MapParamType(param *manifest.ParamType, context TypeContext) (string, error) {
	// Check for enum type first
	if param.Enum != nil {
		typeName := param.Enum.Name
		if param.IsArray() {
			typeName = typeName + "[]"
		}
		return typeName, nil
	}

	// Check for function/delegate type
	if param.Prototype != nil {
		return param.Prototype.Name, nil
	}

	// Regular type mapping
	return m.MapType(param.BaseType(), context, param.IsArray())
}

func (m *V8TypeMapper) MapReturnType(retType *manifest.TypeInfo) (string, error) {
	// Check for enum type
	if retType.Enum != nil {
		typeName := retType.Enum.Name
		if retType.IsArray() {
			typeName = typeName + "[]"
		}
		return typeName, nil
	}

	// Check for function/delegate type
	if retType.Prototype != nil {
		return retType.Prototype.Name, nil
	}

	// Regular type mapping
	return m.MapType(retType.BaseType(), TypeContextReturn, retType.IsArray())
}
