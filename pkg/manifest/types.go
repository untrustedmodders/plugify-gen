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
	FuncName    string      `json:"funcName"`
	ParamTypes  []ParamType `json:"paramTypes"`
	RetType     RetType     `json:"retType"`
}

// ParamType represents a function parameter
type ParamType struct {
	Name        string     `json:"name"`
	Type        string     `json:"type"`
	Ref         bool       `json:"ref,omitempty"`
	Description string     `json:"description,omitempty"`
	Default     *int64     `json:"default,omitempty"`
	Alias       *string    `json:"alias,omitempty"`
	Enum        *EnumType  `json:"enum,omitempty"`
	Prototype   *Prototype `json:"prototype,omitempty"`
}

// RetType represents a type
type RetType struct {
	Type        string     `json:"type"`
	Description string     `json:"description,omitempty"`
	Ref         bool       `json:"ref,omitempty"`
	Alias       *string    `json:"alias,omitempty"`
	Enum        *EnumType  `json:"enum,omitempty"`
	Prototype   *Prototype `json:"prototype,omitempty"`
}

// EnumType represents an enum definition
type EnumType struct {
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
	ParamAliases []*ParamAlias `json:"paramAliases,omitempty"`
	RetAlias     *RetAlias     `json:"retAlias,omitempty"`
}

// ParamAlias represents a parameter that should be treated as a class type
type ParamAlias struct {
	Name  string `json:"name"`
	Owner bool   `json:"owner,omitempty"`
}

// RetAlias represents a return value that should be treated as a class type
type RetAlias struct {
	Name  string `json:"name"`
	Owner bool   `json:"owner,omitempty"`
}

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

// IsArray returns true if the parameter type is an array
func (p *ParamType) IsArray() bool {
	return len(p.Type) > 2 && p.Type[len(p.Type)-2:] == "[]"
}

// BaseType returns the parameter type without array suffix
func (p *ParamType) BaseType() string {
	if p.IsArray() {
		return p.Type[:len(p.Type)-2]
	}
	return p.Type
}
