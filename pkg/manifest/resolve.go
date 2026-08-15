package manifest

import (
	"fmt"
	"sort"
)

// resolve links every by-name prototype/enum reference to its definition and
// hoists inline definitions into the manifest's shared tables, so that after it
// runs every Property.Enum and Property.Prototype points at a fully populated
// definition and each distinct type appears exactly once in Enums/Prototypes.
//
// This mirrors Manifest::Resolve in plugify core; the two must accept the same
// manifests, so the duplicate-definition and cycle rules are kept in step.
func resolve(m *Manifest) error {
	t := typeTable{
		prototypes: map[string]*Prototype{},
		enums:      map[string]*Enum{},
	}

	// Declared definitions go in first, so that a clash between two of them is
	// reported against the manifest's own tables rather than against whichever
	// inline definition happened to be walked first.
	declared := make([]*Prototype, 0, len(m.Prototypes))
	for i := range m.Prototypes {
		if _, err := t.registerPrototype(&m.Prototypes[i], scope("manifest", "")); err != nil {
			return err
		}
		declared = append(declared, m.Prototypes[i])
	}
	for i := range m.Enums {
		if _, err := t.registerEnum(&m.Enums[i], scope("manifest", "")); err != nil {
			return err
		}
	}

	// Pass one: hoist inline definitions into the tables. collect descends into
	// each definition it introduces, so walking the methods and the declared
	// prototypes reaches everything.
	for i := range m.Methods {
		if err := t.collectMethod(&m.Methods[i], scope("method", m.Methods[i].Name)); err != nil {
			return err
		}
	}
	for _, prototype := range declared {
		if err := t.collectPrototype(prototype, scope("prototype", prototype.Name)); err != nil {
			return err
		}
	}

	// Pass two: swap each by-name reference for the definition it names. Every
	// definition is in the tables by now, so references may point forwards.
	for i := range m.Methods {
		if err := t.linkMethod(&m.Methods[i], scope("method", m.Methods[i].Name)); err != nil {
			return err
		}
	}
	for _, prototype := range t.prototypes {
		if err := t.linkPrototype(prototype, scope("prototype", prototype.Name)); err != nil {
			return err
		}
	}

	if err := t.detectCycles(); err != nil {
		return err
	}

	m.Prototypes = t.sortedPrototypes()
	m.Enums = t.sortedEnums()
	return nil
}

// where names the property being checked, for an error message. Formatting it
// eagerly for every parameter of every method cost several times more than all
// the checking put together, so the pieces are carried around instead and only
// joined on the path that actually reports a failure.
type where struct {
	parent *where // the property this one sits inside, if any
	kind   string // "method" or "prototype", when there is no parent
	name   string // the method or prototype it belongs to
	index  int    // parameter position, or idxReturn / idxScope
}

const (
	idxScope  = -1
	idxReturn = -2
)

func scope(kind, name string) where {
	return where{kind: kind, name: name, index: idxScope}
}

func (w where) param(i int) where { return where{parent: &w, index: i} }
func (w where) returnType() where { return where{parent: &w, index: idxReturn} }

func (w where) String() string {
	if w.parent == nil {
		if w.name == "" {
			return w.kind
		}
		return fmt.Sprintf("%s %q", w.kind, w.name)
	}
	prefix := w.parent.String()
	if w.index == idxReturn {
		return prefix + " return type"
	}
	return fmt.Sprintf("%s param[%d]", prefix, w.index)
}

// typeTable gathers every prototype and enum in a manifest into one table keyed
// by name, so a definition declared up front and the same definition written
// inline collapse onto a single shared value.
type typeTable struct {
	prototypes map[string]*Prototype
	enums      map[string]*Enum
}

// registerPrototype reports whether this definition is the first seen under its
// name. A later duplicate is repointed at the first, so generators can treat
// pointer identity as type identity.
func (t *typeTable) registerPrototype(slot **Prototype, context where) (bool, error) {
	prototype := *slot
	if prototype == nil || prototype.ref {
		return false, nil
	}
	if prototype.Name == "" {
		return false, fmt.Errorf("%s: prototype definition must have a name", context)
	}

	existing, found := t.prototypes[prototype.Name]
	if !found {
		t.prototypes[prototype.Name] = prototype
		return true, nil
	}
	if existing != prototype && !samePrototype(existing, prototype) {
		return false, fmt.Errorf("%s: conflicting definitions for prototype %q", context, prototype.Name)
	}
	*slot = existing
	return false, nil
}

func (t *typeTable) registerEnum(slot **Enum, context where) (bool, error) {
	enum := *slot
	if enum == nil || enum.ref {
		return false, nil
	}
	// Manifests predating named references wrote a value-less enum object to mean
	// "the enum of this name, defined elsewhere". That spelling is gone, so say
	// what to write instead rather than letting it collide with the real one.
	if len(enum.Values) == 0 {
		return false, fmt.Errorf(
			"%s: enum %q has no values; write it as \"enum\": %q to refer to a definition in the manifest's 'enums' table",
			context, enum.Name, enum.Name)
	}
	if enum.Name == "" {
		return false, fmt.Errorf("%s: enum definition must have a name", context)
	}

	existing, found := t.enums[enum.Name]
	if !found {
		t.enums[enum.Name] = enum
		return true, nil
	}
	if existing != enum && !sameEnum(existing, enum) {
		return false, fmt.Errorf("%s: conflicting definitions for enum %q", context, enum.Name)
	}
	*slot = existing
	return false, nil
}

func (t *typeTable) collectProperty(prop *Property, context where) error {
	introduced, err := t.registerPrototype(&prop.Prototype, context)
	if err != nil {
		return err
	}
	if _, err := t.registerEnum(&prop.Enum, context); err != nil {
		return err
	}

	// Only descend into a definition this call introduced; one that collapsed
	// onto an earlier definition has already been walked.
	if introduced {
		return t.collectPrototype(prop.Prototype, context)
	}
	return nil
}

func (t *typeTable) collectPrototype(prototype *Prototype, context where) error {
	for i := range prototype.ParamTypes {
		if err := t.collectProperty(&prototype.ParamTypes[i], context.param(i)); err != nil {
			return err
		}
	}
	return t.collectProperty(&prototype.RetType, context.returnType())
}

func (t *typeTable) collectMethod(method *Method, context where) error {
	for i := range method.ParamTypes {
		if err := t.collectProperty(&method.ParamTypes[i], context.param(i)); err != nil {
			return err
		}
	}
	return t.collectProperty(&method.RetType, context.returnType())
}

func (t *typeTable) linkProperty(prop *Property, context where) error {
	if prop.Prototype != nil && prop.Prototype.ref {
		definition, found := t.prototypes[prop.Prototype.Name]
		if !found {
			return fmt.Errorf("%s: unknown prototype %q", context, prop.Prototype.Name)
		}
		prop.Prototype = definition
	}
	if prop.Enum != nil && prop.Enum.ref {
		definition, found := t.enums[prop.Enum.Name]
		if !found {
			return fmt.Errorf("%s: unknown enum %q", context, prop.Enum.Name)
		}
		prop.Enum = definition
	}
	return nil
}

func (t *typeTable) linkPrototype(prototype *Prototype, context where) error {
	for i := range prototype.ParamTypes {
		if err := t.linkProperty(&prototype.ParamTypes[i], context.param(i)); err != nil {
			return err
		}
	}
	return t.linkProperty(&prototype.RetType, context.returnType())
}

func (t *typeTable) linkMethod(method *Method, context where) error {
	for i := range method.ParamTypes {
		if err := t.linkProperty(&method.ParamTypes[i], context.param(i)); err != nil {
			return err
		}
	}
	return t.linkProperty(&method.RetType, context.returnType())
}

// detectCycles rejects a prototype that can reach itself. Such a prototype would
// send anything walking a signature recursively into unbounded recursion, and
// plugify core rejects it too.
func (t *typeTable) detectCycles() error {
	const (
		onStack = 1
		done    = 2
	)
	marks := map[*Prototype]int{}

	var visit func(prototype *Prototype) error
	visit = func(prototype *Prototype) error {
		switch marks[prototype] {
		case done:
			return nil
		case onStack:
			return fmt.Errorf("prototype %q is part of a reference cycle", prototype.Name)
		}
		marks[prototype] = onStack

		descend := func(prop *Property) error {
			if prop.Prototype != nil {
				return visit(prop.Prototype)
			}
			return nil
		}
		for i := range prototype.ParamTypes {
			if err := descend(&prototype.ParamTypes[i]); err != nil {
				return err
			}
		}
		if err := descend(&prototype.RetType); err != nil {
			return err
		}

		marks[prototype] = done
		return nil
	}

	for _, name := range sortedKeys(t.prototypes) {
		if err := visit(t.prototypes[name]); err != nil {
			return err
		}
	}
	return nil
}

func (t *typeTable) sortedPrototypes() []*Prototype {
	if len(t.prototypes) == 0 {
		return nil
	}
	out := make([]*Prototype, 0, len(t.prototypes))
	for _, name := range sortedKeys(t.prototypes) {
		out = append(out, t.prototypes[name])
	}
	return out
}

func (t *typeTable) sortedEnums() []*Enum {
	if len(t.enums) == 0 {
		return nil
	}
	out := make([]*Enum, 0, len(t.enums))
	for _, name := range sortedKeys(t.enums) {
		out = append(out, t.enums[name])
	}
	return out
}

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// Two definitions written under the same name are only allowed if they say the
// same thing, so that repeating an inline enum across several methods keeps
// working. Nested prototypes and enums compare by name rather than by value:
// within a manifest a name denotes exactly one type, which also keeps this
// terminating for recursive types.

func sameProperty(a, b *Property) bool {
	if a.Type != b.Type || a.Ref != b.Ref {
		return false
	}
	aProto, bProto := "", ""
	if a.Prototype != nil {
		aProto = a.Prototype.Name
	}
	if b.Prototype != nil {
		bProto = b.Prototype.Name
	}
	aEnum, bEnum := "", ""
	if a.Enum != nil {
		aEnum = a.Enum.Name
	}
	if b.Enum != nil {
		bEnum = b.Enum.Name
	}
	return aProto == bProto && aEnum == bEnum
}

func sameEnum(a, b *Enum) bool {
	if len(a.Values) != len(b.Values) {
		return false
	}
	for i := range a.Values {
		if a.Values[i].Name != b.Values[i].Name || a.Values[i].Value != b.Values[i].Value {
			return false
		}
	}
	return true
}

func samePrototype(a, b *Prototype) bool {
	if len(a.ParamTypes) != len(b.ParamTypes) {
		return false
	}
	if !sameProperty(&a.RetType, &b.RetType) {
		return false
	}
	for i := range a.ParamTypes {
		if !sameProperty(&a.ParamTypes[i], &b.ParamTypes[i]) {
			return false
		}
	}
	return true
}
