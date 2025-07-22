package ds

import "fmt"

// UnificationBinding representa una única ligadura de una variable a un valor dentro de un entorno.
type UnificationBinding struct {
	Variable *Symbol // La variable que ha sido ligada
	OldValue *Symbol // El valor al que estaba ligada ANTES de esta unificación (para deshacer)
	WasBound bool    // true si la variable ya estaba ligada antes de esta unificación
}

// Environment representa el entorno de ligaduras para una rama de unificación.
// Es una colección de ligaduras temporales que pueden ser aplicadas o deshechas.
type Environment struct {
	// 'Bindings' almacena las ligaduras actuales de las variables en este entorno.
	// La clave es el ID de la variable, el valor es el símbolo al que está ligada.
	Bindings map[SymbolID]*Symbol
	// 'Trail' es una pila de operaciones de deshacer.
	// Cada elemento registra el estado anterior de una variable antes de ser ligada.
	Trail []UnificationBinding
}

// NewEnvironment crea un nuevo entorno de unificación vacío.
func NewEnvironment() *Environment {
	return &Environment{
		Bindings: make(map[SymbolID]*Symbol),
		Trail:    []UnificationBinding{},
	}
}

// --- Modificadores del Environment (para uso interno de Unify, Deref, Bind) ---

// GetBinding obtiene la ligadura de una variable en este entorno.
func (env *Environment) GetBinding(varID SymbolID) (*Symbol, bool) {
	val, ok := env.Bindings[varID]
	return val, ok
}

// AddBinding añade una ligadura al entorno y registra el estado anterior en el Trail.
func (env *Environment) AddBinding(variable *Symbol, value *Symbol) {
	// Antes de ligar, guarda el estado actual de la variable en el Trail
	oldBinding, wasBound := env.Bindings[variable.ID]
	env.Trail = append(env.Trail, UnificationBinding{
		Variable: variable,
		OldValue: oldBinding,
		WasBound: wasBound,
	})

	// Aplica la nueva ligadura en el entorno
	env.Bindings[variable.ID] = value
}

// Backtrack deshace las ligaduras registradas en el Trail desde un punto de control.
// Esto se llamaría si una unificación o una rama de resolución falla.
// Por simplicidad, esta función es un placeholder. En un Prolog real, se pasaría
// un "punto de elección" o un "estado del Trail" para deshacer solo hasta allí.
func (env *Environment) Backtrack() {
	// En un sistema Prolog real, esta función sería más sofisticada,
	// desapilando solo las ligaduras hechas desde el último punto de elección.
	// Aquí, simplemente vacía el Trail y el mapa para un ejemplo didáctico.
	// Para un backtracking granular, necesitas un "checkpoint" en el Trail.
	fmt.Println("INFO: Backtracking environment.")
	for _, bind := range env.Trail {
		if bind.WasBound {
			env.Bindings[bind.Variable.ID] = bind.OldValue
		} else {
			delete(env.Bindings, bind.Variable.ID)
		}
		// Es importante también "desligar" el Symbol.IsBound y Symbol.Binding si fueron modificados directamente
		// Pero la idea de este Environment es evitar modificar el Symbol globalmente hasta el commit.
	}
	env.Trail = []UnificationBinding{}        // Limpiar el Trail
	env.Bindings = make(map[SymbolID]*Symbol) // Resetear Bindings (simplificado para ejemplo)
}

// ApplyBindingsToSymbols recorre las ligaduras en el entorno y las "commit" a los Símbolos originales.
// Esto se llamaría si una rama de unificación tiene éxito y queremos que las ligaduras persistan globalmente.
func (env *Environment) ApplyBindingsToSymbols() {
	for varID, val := range env.Bindings {
		if variable, ok := SymbolsByID[varID]; ok { // Asume que SymbolsByID es donde están las variables originales
			variable.IsBound = true
			variable.Binding = val
			variable.State = Embodied
		}
	}
}
