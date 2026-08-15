package manifest

import "encoding/json"

// Manifest represents a .pplugin file structure
type Manifest struct {
	Schema       string       `json:"$schema"`
	Version      string       `json:"version"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Author       string       `json:"author"`
	Website      string       `json:"website"`
	License      string       `json:"license"`
	Entry        string       `json:"entry"`
	Platforms    []string     `json:"platforms"`
	Language     string       `json:"language"`
	Dependencies []Dependency `json:"dependencies"`
	Methods      []Method     `json:"methods"`
	Classes      []Class      `json:"classes,omitempty"`
	Prototypes   []*Prototype `json:"prototypes,omitempty"`
	Enums        []*Enum      `json:"enums,omitempty"`
}

// Dependency represents a plugin dependency
type Dependency struct {
	Name     string `json:"name"`
	Optional bool   `json:"optional,omitempty"`
}

// Method represents an exported method/function
type Method struct {
	Name        string      `json:"name"`
	Group       string      `json:"group,omitempty"`
	Description string      `json:"description,omitempty"`
	Deprecated  string      `json:"deprecated,omitempty"`
	FuncName    string      `json:"funcName"`
	ParamTypes  []ParamType `json:"paramTypes"`
	RetType     RetType     `json:"retType"`
}

// Property represents a parameter/return type
type Property struct {
	Name        string     `json:"name"`
	Type        string     `json:"type"`
	Ref         bool       `json:"ref,omitempty"`
	Description string     `json:"description,omitempty"`
	Default     *any       `json:"default,omitempty"`
	Alias       *Alias     `json:"alias,omitempty"`
	Enum        *Enum      `json:"enum,omitempty"`
	Prototype   *Prototype `json:"prototype,omitempty"`
}

// ParamType represents a function parameter
type ParamType = Property

// RetType represents a type
type RetType = Property

// Alias represents an alias definition
type Alias struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Deprecated  string `json:"deprecated,omitempty"`
}

// Enum represents an enum definition
type Enum struct {
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Deprecated  string  `json:"deprecated,omitempty"`
	Values      []Value `json:"values"`

	// ref records that the manifest wrote this as a bare name rather than a
	// definition. resolve fills in the rest and clears it, so generators never
	// see it set.
	ref bool
}

// UnmarshalJSON accepts either a full definition or the name of an entry in the
// manifest's top-level "enums" table.
func (e *Enum) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '"' {
		var name string
		if err := json.Unmarshal(data, &name); err != nil {
			return err
		}
		*e = Enum{Name: name, ref: true}
		return nil
	}

	type plain Enum // shed the method set so this does not recurse
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	*e = Enum(decoded)
	return nil
}

// Value represents a single enum value
type Value struct {
	Name        string `json:"name"`
	Value       int    `json:"value"`
	Description string `json:"description,omitempty"`
}

// Prototype represents a function pointer/delegate type
type Prototype struct {
	Name        string      `json:"name"`
	Description string      `json:"description,omitempty"`
	Deprecated  string      `json:"deprecated,omitempty"`
	ParamTypes  []ParamType `json:"paramTypes"`
	RetType     RetType     `json:"retType"`

	// ref records that the manifest wrote this as a bare name rather than a
	// definition. resolve fills in the rest and clears it, so generators never
	// see it set.
	ref bool
}

// UnmarshalJSON accepts either a full definition or the name of an entry in the
// manifest's top-level "prototypes" table.
func (p *Prototype) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '"' {
		var name string
		if err := json.Unmarshal(data, &name); err != nil {
			return err
		}
		*p = Prototype{Name: name, ref: true}
		return nil
	}

	type plain Prototype // shed the method set so this does not recurse
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	*p = Prototype(decoded)
	return nil
}

// Class represents an RAII wrapper class for handle-based APIs
type Class struct {
	Name         string    `json:"name"`
	Group        string    `json:"group,omitempty"`
	Description  string    `json:"description,omitempty"`
	Deprecated   string    `json:"deprecated,omitempty"`
	HandleType   string    `json:"handleType,omitempty"`
	HandleAlias  string    `json:"handleAlias,omitempty"`
	InvalidValue string    `json:"invalidValue,omitempty"`
	NullPolicy   string    `json:"nullPolicy,omitempty"`
	Constructors []string  `json:"constructors,omitempty"`
	Destructor   *string   `json:"destructor,omitempty"`
	Bindings     []Binding `json:"bindings"`
}

// Binding represents a method in a wrapper class
type Binding struct {
	Name         string        `json:"name"`
	Method       string        `json:"method"`
	BindSelf     bool          `json:"bindSelf,omitempty"`
	Deprecated   string        `json:"deprecated,omitempty"`
	ParamAliases []*ParamAlias `json:"paramAliases,omitempty"`
	RetAlias     *RetAlias     `json:"retAlias,omitempty"`
}

// Bind represents a value that should be treated as a class type
type Bind struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Deprecated  string `json:"deprecated,omitempty"`
	Owner       bool   `json:"owner,omitempty"`
}

// ParamAlias represents a parameter that should be treated as a class type
type ParamAlias = Bind

// RetAlias represents a return value that should be treated as a class type
type RetAlias = Bind

// IsArray returns true if the type is an array (ends with [])
func (t *RetType) IsArray() bool {
	return len(t.Type) > 2 && t.Type[len(t.Type)-2:] == "[]"
}

// BaseType returns the type without array suffix
func (t *RetType) BaseType() string {
	if t.IsArray() {
		return t.Type[:len(t.Type)-2]
	}
	return t.Type
}
