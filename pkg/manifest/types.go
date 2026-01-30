package manifest

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
	Default     *int64     `json:"default,omitempty"`
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
}

// Enum represents an enum definition
type Enum struct {
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	Description string      `json:"description,omitempty"`
	Values      []EnumValue `json:"values"`
}

// EnumValue represents a single enum value
type EnumValue struct {
	Name        string `json:"name"`
	Value       int    `json:"value"`
	Description string `json:"description,omitempty"`
}

// Prototype represents a function pointer/delegate type
type Prototype struct {
	Name        string      `json:"name"`
	Description string      `json:"description,omitempty"`
	ParamTypes  []ParamType `json:"paramTypes"`
	RetType     RetType     `json:"retType"`
}

// Class represents an RAII wrapper class for handle-based APIs
type Class struct {
	Name         string    `json:"name"`
	Group        string    `json:"group,omitempty"`
	Description  string    `json:"description,omitempty"`
	Deprecated   string    `json:"deprecated,omitempty"`
	HandleType   string    `json:"handleType,omitempty"`
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
	Name  string `json:"name"`
	Owner bool   `json:"owner,omitempty"`
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
