package generator

import (
	"fmt"
	"strings"

	"github.com/untrustedmodders/plugify-generator/pkg/manifest"
)

// LuaGenerator generates Lua stub files
type LuaGenerator struct {
	*BaseGenerator
}

// NewLuaGenerator creates a new Lua generator
func NewLuaGenerator() *LuaGenerator {
	invalidNames := []string{
		"and", "break", "do", "else", "elseif", "end", "false",
		"for", "function", "goto", "if", "in", "local", "nil",
		"not", "or", "repeat", "return", "then", "true", "until", "while",
	}

	return &LuaGenerator{
		BaseGenerator: NewBaseGenerator("lua", NewLuaTypeMapper(), invalidNames),
	}
}

// Generate generates Lua bindings
func (g *LuaGenerator) Generate(m *manifest.Manifest) (*GeneratorResult, error) {
	g.ResetCaches()

	var sb strings.Builder

	// Header comment
	sb.WriteString(fmt.Sprintf("-- Generated from %s.pplugin\n\n", m.Name))

	// Generate enums
	enumsCode, err := g.generateEnums(m)
	if err != nil {
		return nil, err
	}
	if enumsCode != "" {
		sb.WriteString(enumsCode)
		sb.WriteString("\n")
	}

	// Generate methods
	for _, method := range m.Methods {
		methodCode, err := g.generateMethod(&method)
		if err != nil {
			return nil, fmt.Errorf("failed to generate method %s: %w", method.Name, err)
		}
		sb.WriteString(methodCode)
		sb.WriteString("\n")
	}

	result := &GeneratorResult{
		Files: map[string]string{
			fmt.Sprintf("pps/%s.lua", m.Name): sb.String(),
		},
	}

	return result, nil
}

func (g *LuaGenerator) generateEnums(m *manifest.Manifest) (string, error) {
	var sb strings.Builder

	for _, method := range m.Methods {
		// Check return type
		if method.RetType.Enum != nil && !g.IsEnumCached(method.RetType.Enum.Name) {
			enumCode := g.generateEnum(method.RetType.Enum)
			sb.WriteString(enumCode)
			sb.WriteString("\n")
			g.CacheEnum(method.RetType.Enum.Name)
		}

		// Check parameters
		for _, param := range method.ParamTypes {
			if param.Enum != nil && !g.IsEnumCached(param.Enum.Name) {
				enumCode := g.generateEnum(param.Enum)
				sb.WriteString(enumCode)
				sb.WriteString("\n")
				g.CacheEnum(param.Enum.Name)
			}

			// Check nested prototypes
			if param.Prototype != nil {
				g.processPrototypeEnums(param.Prototype, &sb)
			}
		}

		// Check return type prototype
		if method.RetType.Prototype != nil {
			g.processPrototypeEnums(method.RetType.Prototype, &sb)
		}
	}

	return sb.String(), nil
}

func (g *LuaGenerator) processPrototypeEnums(proto *manifest.Prototype, sb *strings.Builder) {
	if proto.RetType.Enum != nil && !g.IsEnumCached(proto.RetType.Enum.Name) {
		sb.WriteString(g.generateEnum(proto.RetType.Enum))
		sb.WriteString("\n")
		g.CacheEnum(proto.RetType.Enum.Name)
	}

	for _, param := range proto.ParamTypes {
		if param.Enum != nil && !g.IsEnumCached(param.Enum.Name) {
			sb.WriteString(g.generateEnum(param.Enum))
			sb.WriteString("\n")
			g.CacheEnum(param.Enum.Name)
		}
		if param.Prototype != nil {
			g.processPrototypeEnums(param.Prototype, sb)
		}
	}
}

func (g *LuaGenerator) generateEnum(enum *manifest.EnumType) string {
	var sb strings.Builder

	// Enum comment
	if enum.Description != "" {
		sb.WriteString(fmt.Sprintf("-- %s: %s\n", enum.Name, enum.Description))
	} else {
		sb.WriteString(fmt.Sprintf("-- Enum: %s\n", enum.Name))
	}

	// Lua table
	sb.WriteString(fmt.Sprintf("%s = {\n", enum.Name))

	for _, val := range enum.Values {
		if val.Description != "" {
			sb.WriteString(fmt.Sprintf("  -- %s\n", val.Description))
		}
		sb.WriteString(fmt.Sprintf("  %s = %d,\n", val.Name, val.Value))
	}

	sb.WriteString("}\n")

	return sb.String()
}

func (g *LuaGenerator) generateMethod(method *manifest.Method) (string, error) {
	var sb strings.Builder

	// Generate LDoc-style documentation
	sb.WriteString(g.generateDocumentation(method))

	// Generate function signature
	params := g.formatParameters(method.ParamTypes)
	sb.WriteString(fmt.Sprintf("function %s(%s) end\n", g.SanitizeName(method.Name), params))

	return sb.String(), nil
}

func (g *LuaGenerator) formatParameters(params []manifest.ParamType) string {
	if len(params) == 0 {
		return ""
	}

	result := ""
	for i, param := range params {
		if i > 0 {
			result += ", "
		}

		paramName := param.Name
		if paramName == "" {
			paramName = fmt.Sprintf("p%d", i)
		}

		result += g.SanitizeName(paramName)
	}

	return result
}

func (g *LuaGenerator) generateDocumentation(method *manifest.Method) string {
	var sb strings.Builder

	// Main description
	if method.Description != "" {
		sb.WriteString(fmt.Sprintf("--- %s\n", method.Description))
	} else {
		sb.WriteString(fmt.Sprintf("--- %s\n", method.Name))
	}

	// Parameters
	for _, param := range method.ParamTypes {
		paramName := param.Name
		if paramName == "" {
			paramName = "unnamed"
		}
		paramType := param.Type
		paramDesc := param.Description

		sb.WriteString(fmt.Sprintf("-- @param %s %s %s\n",
			g.SanitizeName(paramName), paramType, paramDesc))
	}

	// Return type
	if method.RetType.Type != "void" {
		retDesc := method.RetType.Description
		sb.WriteString(fmt.Sprintf("-- @return %s %s\n", method.RetType.Type, retDesc))
	}

	// Callback prototypes
	for _, param := range method.ParamTypes {
		if param.Prototype != nil {
			proto := param.Prototype
			protoName := proto.Name
			protoDesc := proto.Description
			if protoDesc == "" {
				protoDesc = "Callback function"
			}

			sb.WriteString(fmt.Sprintf("-- @callback %s %s - %s\n",
				protoName, g.SanitizeName(param.Name), protoDesc))

			// Callback parameters
			for _, protoParam := range proto.ParamTypes {
				sb.WriteString(fmt.Sprintf("-- @param %s %s %s\n",
					g.SanitizeName(protoParam.Name), protoParam.Type, protoParam.Description))
			}

			// Callback return
			if proto.RetType.Type != "void" {
				sb.WriteString(fmt.Sprintf("-- @return %s %s\n",
					proto.RetType.Type, proto.RetType.Description))
			}
		}
	}

	return sb.String()
}

// LuaTypeMapper is a placeholder - Lua doesn't need explicit type mapping in stubs
type LuaTypeMapper struct{}

func NewLuaTypeMapper() *LuaTypeMapper {
	return &LuaTypeMapper{}
}

func (m *LuaTypeMapper) MapType(baseType string, context TypeContext, isArray bool) (string, error) {
	// Lua is dynamically typed, return empty string (not used)
	return "", nil
}

func (m *LuaTypeMapper) MapParamType(param *manifest.ParamType, context TypeContext) (string, error) {
	return "", nil
}

func (m *LuaTypeMapper) MapReturnType(retType *manifest.TypeInfo) (string, error) {
	return "", nil
}
