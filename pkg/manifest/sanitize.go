package manifest

import (
	"strings"
)

// SanitizeNameFunc is a function that sanitizes a name for language-specific reserved keywords
type SanitizeNameFunc func(name string) string

// Sanitize sanitizes all names in the manifest
// It auto-generates parameter names if empty and applies language-specific sanitization
func (m *Manifest) Sanitize(sanitizeName SanitizeNameFunc) {
	// Sanitize methods
	for i := range m.Methods {
		method := &m.Methods[i]

		// Sanitize method name
		if method.Name != "" && sanitizeName != nil {
			method.Name = sanitizeName(method.Name)
		}

		// Sanitize function name
		if method.FuncName != "" && sanitizeName != nil {
			method.FuncName = sanitizeName(method.FuncName)
		}

		// Sanitize group name
		if method.Group != "" && sanitizeName != nil {
			method.Group = sanitizeName(method.Group)
		}

		// Make groups lower case
		if method.Group == "" {
			method.Group = "core"
		} else {
			method.Group = strings.ToLower(method.Group)
		}

		// Sanitize parameters
		sanitizeParamTypes(method.ParamTypes, sanitizeName)

		// Sanitize return type enum/prototype names
		sanitizeTypeInfo(&method.RetType, sanitizeName)
	}

	// Sanitize classes
	for i := range m.Classes {
		class := &m.Classes[i]

		// Sanitize class name
		if class.Name != "" && sanitizeName != nil {
			class.Name = sanitizeName(class.Name)
		}

		// Sanitize group name
		if class.Group != "" && sanitizeName != nil {
			class.Group = sanitizeName(class.Group)
		}

		// Make groups lower case
		if class.Group == "" {
			class.Group = "core"
		} else {
			class.Group = strings.ToLower(class.Group)
		}

		// Sanitize bindings
		for j := range class.Bindings {
			binding := &class.Bindings[j]

			if binding.Name != "" && sanitizeName != nil {
				binding.Name = sanitizeName(binding.Name)
			}

			if binding.Method != "" && sanitizeName != nil {
				binding.Method = sanitizeName(binding.Method)
			}

			// Sanitize param aliases
			for k := range binding.ParamAliases {
				if binding.ParamAliases[k] != nil && binding.ParamAliases[k].Name != "" && sanitizeName != nil {
					binding.ParamAliases[k].Name = sanitizeName(binding.ParamAliases[k].Name)
				}
			}

			// Sanitize return alias
			if binding.RetAlias != nil && binding.RetAlias.Name != "" && sanitizeName != nil {
				binding.RetAlias.Name = sanitizeName(binding.RetAlias.Name)
			}
		}
	}
}

// sanitizeParamTypes sanitizes parameter names and types
func sanitizeParamTypes(params []ParamType, sanitizeName SanitizeNameFunc) {
	for i := range params {
		param := &params[i]

		// Sanitize parameter name
		if sanitizeName != nil {
			param.Name = sanitizeName(param.Name)
		}

		// Sanitize enum name if present
		if param.Enum != nil {
			sanitizeEnum(param.Enum, sanitizeName)
		}

		// Sanitize prototype if present
		if param.Prototype != nil {
			sanitizePrototype(param.Prototype, sanitizeName)
		}
	}
}

// sanitizeTypeInfo sanitizes type information (for return types)
func sanitizeTypeInfo(typeInfo *TypeInfo, sanitizeName SanitizeNameFunc) {
	if typeInfo.Enum != nil {
		sanitizeEnum(typeInfo.Enum, sanitizeName)
	}

	if typeInfo.Prototype != nil {
		sanitizePrototype(typeInfo.Prototype, sanitizeName)
	}
}

// sanitizeEnum sanitizes enum names and values
func sanitizeEnum(enum *EnumType, sanitizeName SanitizeNameFunc) {
	if enum.Name != "" && sanitizeName != nil {
		enum.Name = sanitizeName(enum.Name)
	}

	for i := range enum.Values {
		value := &enum.Values[i]
		if value.Name != "" && sanitizeName != nil {
			value.Name = sanitizeName(value.Name)
		}
	}
}

// sanitizePrototype sanitizes prototype/delegate names and parameters
func sanitizePrototype(proto *Prototype, sanitizeName SanitizeNameFunc) {
	if proto.Name != "" && sanitizeName != nil {
		proto.Name = sanitizeName(proto.Name)
	}

	// Sanitize prototype parameters
	sanitizeParamTypes(proto.ParamTypes, sanitizeName)

	// Sanitize prototype return type
	sanitizeTypeInfo(&proto.RetType, sanitizeName)
}
