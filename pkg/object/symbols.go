// pkg/object/symbols.go

package object

// Útil si en el futuro quieres String() más detallado

// SymbolObject representa un símbolo en NexusL,
// que puede tener propiedades dinámicas. (e.g., "cli", "print", "Juan").
// Los símbolos son identificadores que no son literales directos ni variables.
type SymbolObject struct {
	Value string
	// Members es un mapa para almacenar propiedades/métodos asociados a este símbolo.
	// Ej: "cli" puede tener un miembro "print" que es un BuiltinObject.
	// "David" puede tener un miembro "rightLeg" que es un StringObject("healthy").
	Members map[string]Object
}

func (s *SymbolObject) Type() ObjectType { return SYMBOL_OBJ }
func (s *SymbolObject) String() string {
	// Puedes extender esta función para mostrar los miembros para depuración.
	// Por ahora, solo devuelve el nombre del símbolo.
	// Si quieres algo más detallado:
	// var b bytes.Buffer
	// b.WriteString(s.Value)
	// if len(s.Members) > 0 {
	//     b.WriteString(" {")
	//     first := true
	//     for k, v := range s.Members {
	//         if !first {
	//             b.WriteString(", ")
	//         }
	//         b.WriteString(k)
	//         b.WriteString(": ")
	//         b.WriteString(v.String())
	//         first = false
	//     }
	//     b.WriteString("}")
	// }
	// return b.String()
	return s.Value
}

// NewSymbolObject crea e inicializa un nuevo SymbolObject.
func NewSymbolObject(value string) *SymbolObject {
	return &SymbolObject{
		Value:   value,
		Members: make(map[string]Object), // ¡Importante inicializar el mapa!
	}
}

// GetMember obtiene un miembro del SymbolObject por su nombre.
func (s *SymbolObject) GetMember(name string) (Object, bool) {
	member, ok := s.Members[name]
	return member, ok
}

// SetMember establece un miembro en el SymbolObject.
func (s *SymbolObject) SetMember(name string, val Object) {
	s.Members[name] = val
}
