// Gothic/ds/symbol.go  (Actualizado)
package ds

import (
	"fmt"
	"sync"
)

// Global maps for symbol management
var (
	mu                  sync.Mutex
	nextID              SymbolID
	SymbolsByID         map[SymbolID]*Symbol
	SymbolsByPublicName map[string]*Symbol
)

func init() {
	mu.Lock()
	defer mu.Unlock()
	nextID = 1000 // Starting ID for symbols
	SymbolsByID = make(map[SymbolID]*Symbol)
	SymbolsByPublicName = make(map[string]*Symbol)

	// Pre-populate some fundamental ThingTypes as Symbols themselves, if needed,
	// or just ensure their string values are consistent with the `ThingType` consts.
	// For now, the consts are sufficient.
}

// SymbolID is the unique identifier for a Symbol
type SymbolID int

// SymbolState represents the lifecycle state of a Symbol
type SymbolState int

const (
	Exists   SymbolState = iota // Symbol exists (declared)
	Embodied                    // Symbol has a concrete value
)

// String representation for SymbolState
func (s SymbolState) String() string {
	switch s {
	case Exists:
		return "Exists"
	case Embodied:
		return "Embodied"
	default:
		return fmt.Sprintf("UnknownState(%d)", s)
	}
}

// ThingType define las categorías semánticas de los símbolos.
// Esto es nuevo aquí (movido desde el antiguo ds/ds.go).
type ThingType string

const (
	// Tipos fundamentales de "Thing" que un Symbol puede representar.
	// Exportamos estas constantes con mayúscula inicial.
	LiteralType      ThingType = "Literal"      // Para valores concretos (números, strings, bools)
	IdentifierType   ThingType = "Identifier"   // Para nombres de entidades, variables, etc.
	PredicateType    ThingType = "Predicate"    // Para predicados (has:, is:, do:)
	TripletScopeType ThingType = "TripletScope" // Para scopes de alto nivel (fact, program, func)
	MacroType        ThingType = "Macro"        // Para macros de lenguaje (que se expanden en AST)
	// Puedes añadir más ThingTypes aquí a medida que los necesites:
	// AgentType        ThingType = "Agent"
	// LocationType     ThingType = "Location"
	// EventType        ThingType = "Event"
)

// SymbolProc is the type for a procedure attached to a Symbol
type SymbolProc func(args ...interface{}) (interface{}, error)

// Symbol represents a node in the NexusMesh (conceptually a node in an RDF graph)
type Symbol struct {
	ID         SymbolID
	State      SymbolState            // Exists, Embodied
	Thing      ThingType              // CAMBIO AQUÍ: Ahora es de tipo ThingType, no string
	PublicName string                 // Human-readable, globally unique name (optional, but used for resolution)
	Value      interface{}            // The concrete value the symbol might embody (e.g., string "robot-1", int 42, custom struct)
	Properties map[string]interface{} // Key-value store for additional attributes of the symbol
	Proc       SymbolProc             // If the symbol represents a procedure (action, function, macro)
	Embedding  []float32              // El embedding del símbolo
}

// NewSymbol creates a new unique Symbol instance.
func NewSymbol() *Symbol {
	mu.Lock()
	defer mu.Unlock()

	s := &Symbol{
		ID:         nextID,
		State:      Exists,
		Thing:      IdentifierType, // Default a IdentifierType, no "Thing" string
		Properties: make(map[string]interface{}),
	}
	SymbolsByID[s.ID] = s
	nextID++
	return s
}

// NewSymbolWithPublicName es un helper para crear y asignar nombre de una vez
func NewSymbolWithPublicName(name string, thingType ThingType) *Symbol {
	s := NewSymbol()
	s.AssignPublicName(name)
	s.SetThing(thingType) // Usar el nuevo método SetThing con ThingType
	return s
}

// AssignPublicName assigns a human-readable name to a Symbol.
// It ensures the symbol is discoverable by this name and handles updates in the global map.
func (s *Symbol) AssignPublicName(name string) {
	mu.Lock()
	defer mu.Unlock()

	if s.PublicName != "" && s.PublicName != name {
		delete(SymbolsByPublicName, s.PublicName)
	}

	s.PublicName = name
	SymbolsByPublicName[name] = s
}

// LookupSymbolByPublicName retrieves a Symbol by its public name.
func LookupSymbolByPublicName(name string) (*Symbol, bool) {
	mu.Lock()
	defer mu.Unlock()
	sym, ok := SymbolsByPublicName[name]
	return sym, ok
}

// SetThing sets the "type" or "concept" that this symbol represents.
// CAMBIO AQUÍ: Ahora acepta ThingType directamente.
func (s *Symbol) SetThing(thingType ThingType) {
	s.Thing = thingType
}

// InstantiateAs sets the concrete value for the Symbol, changing its state to Embodied.
func (s *Symbol) InstantiateAs(value interface{}) {
	s.Value = value
	s.State = Embodied
}

// AddProperty adds or updates a key-value property for the Symbol.
func (s *Symbol) AddProperty(key string, value interface{}) {
	s.Properties[key] = value
}

// GetProperty retrieves a property by its key.
func (s *Symbol) GetProperty(key string) (interface{}, bool) {
	val, ok := s.Properties[key]
	return val, ok
}

// CallProc attempts to execute the procedure attached to the Symbol.
func (s *Symbol) CallProc(args ...interface{}) (interface{}, error) {
	if s.Proc == nil {
		return nil, fmt.Errorf("no procedure attached to symbol %s (ID: %d)", s.PublicName, s.ID)
	}
	return s.Proc(args...)
}

// String provides a human-readable representation of the Symbol.
func (s *Symbol) String() string {
	valStr := "<nil>"
	if s.Value != nil {
		valStr = fmt.Sprintf("%v", s.Value)
	}
	namePart := fmt.Sprintf("anon:%d", s.ID)
	if s.PublicName != "" {
		namePart = s.PublicName
	}
	// CAMBIO AQUÍ: Incluimos s.Thing directamente.
	return fmt.Sprintf("[%s | %s | %s | %s]", namePart, s.State.String(), s.Thing, valStr)
}
