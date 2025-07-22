// /nexusl/internal/proloGo/unify.go
// .
// Unificación de símbolos en ProloGo
// .
// Este archivo contiene la lógica para unificar símbolos en ProloGo,
// incluyendo la gestión de ligaduras de variables y la comparación de
// estructuras complejas.
// .
// La unificación es un proceso clave en la resolución de consultas
// y la ejecución de programas ProloGo.
// .
package prologo

import (
	"fmt"

	"github.com/devicemxl/nexusl/ds"
)

// UnificationBinding representa una única ligadura de una variable a un valor dentro de un entorno.
type UnificationBinding struct {
	Variable *ds.Symbol // La variable que ha sido ligada
	OldValue *ds.Symbol // El valor al que estaba ligada ANTES de esta unificación (para deshacer)
	WasBound bool       // true si la variable ya estaba ligada antes de esta unificación
}

// Environment representa el entorno de ligaduras para una rama de unificación.
// Es una colección de ligaduras temporales que pueden ser aplicadas o deshechas.
type Environment struct {
	// 'Bindings' almacena las ligaduras actuales de las variables en este entorno.
	// La clave es el ID de la variable, el valor es el símbolo al que está ligada.
	Bindings map[ds.SymbolID]*ds.Symbol
	// 'trail' es una pila de operaciones de deshacer.
	// Cada elemento registra el estado anterior de una variable antes de ser ligada.
	trail []UnificationBinding
}

// --- Métodos de Environment ---

// NewEnvironment crea un nuevo entorno de unificación vacío.
func NewEnvironment() *Environment {
	return &Environment{
		Bindings: make(map[ds.SymbolID]*ds.Symbol),
		trail:    []UnificationBinding{},
	}
}

// GetBinding obtiene la ligadura de una variable en este entorno.
func (env *Environment) GetBinding(varID ds.SymbolID) (*ds.Symbol, bool) {
	val, ok := env.Bindings[varID]
	return val, ok
}

// AddBinding añade una ligadura al entorno y registra el estado anterior en el trail.
func (env *Environment) AddBinding(variable *ds.Symbol, value *ds.Symbol) {
	// Antes de ligar, guarda el estado actual de la variable en el trail
	oldBinding, wasBound := env.Bindings[variable.ID]
	env.trail = append(env.trail, UnificationBinding{
		Variable: variable,
		OldValue: oldBinding,
		WasBound: wasBound,
	})

	// Aplica la nueva ligadura en el entorno
	env.Bindings[variable.ID] = value
}

// Backtrack deshace las ligaduras registradas en el trail desde un punto de control.
// Esto se llamaría si una unificación o una rama de resolución falla.
// Por simplicidad, esta función es un placeholder. En un Prolog real, se pasaría
// un "punto de elección" o un "estado del trail" para deshacer solo hasta allí.
func (env *Environment) Backtrack() {
	fmt.Println("INFO: Backtracking environment (simplified for example).")
	for _, bind := range env.trail {
		if bind.WasBound {
			env.Bindings[bind.Variable.ID] = bind.OldValue
		} else {
			delete(env.Bindings, bind.Variable.ID)
		}
	}
	env.trail = []UnificationBinding{}              // Limpiar el trail
	env.Bindings = make(map[ds.SymbolID]*ds.Symbol) // Resetear Bindings (simplificado para ejemplo)
}

// ApplyBindingsToSymbols recorre las ligaduras en el entorno y las "commit" a los Símbolos originales.
// Esto se llamaría si una rama de unificación tiene éxito y queremos que las ligaduras persistan globalmente.
func (env *Environment) ApplyBindingsToSymbols() {
	fmt.Println("INFO: Applying Bindings to global symbols.")
	for varID, val := range env.Bindings {
		if variable, ok := ds.SymbolsByID[varID]; ok { // Asume que ds.SymbolsByID es donde están las variables originales
			variable.IsBound = true
			variable.Binding = val
			variable.State = ds.Embodied
		}
	}
}

// --- Funciones para el Motor de Unificación ---

// Deref sigue las ligaduras de un Símbolo en un entorno dado.
// Primero consulta el entorno de unificación, luego el Symbol.Binding si existe.
func Deref(s *ds.Symbol, env *Environment) *ds.Symbol {
	fmt.Printf("Deref called: s=%s\n", s.String()) // Add this
	// Si es una variable, primero consulta el entorno actual de unificación
	if s.LogicalType == ds.LT_Variable {
		if boundVal, ok := env.GetBinding(s.ID); ok {
			// Si está ligada en este entorno, sigue desreferenciando recursivamente en este entorno
			return Deref(boundVal, env)
		}
		// Si no está ligada en este entorno, pero tiene un Binding "global" (desde una unificación anterior commitada)
		if s.IsBound && s.Binding != nil {
			return Deref(s.Binding, env) // Sigue el binding "permanente" del símbolo
		}
	}
	// Si no es variable, o no está ligada en el entorno/globalmente, devuelve el propio símbolo.
	return s
}

// Bind establece una ligadura para una variable DENTRO DEL ENTORNO actual.
// No modifica directamente el *ds.Symbol a nivel global, lo hace a través del entorno.
func Bind(variable *ds.Symbol, value *ds.Symbol, env *Environment) error {
	fmt.Printf("Bind called: variable=%s, value=%s\n", variable.String(), value.String()) // Add this

	if variable.LogicalType != ds.LT_Variable {
		return fmt.Errorf("attempted to bind a non-variable Symbol: %s", variable.PublicName)
	}

	// Obtener el valor desreferenciado para evitar ciclos infinitos o ligaduras triviales
	valDeref := Deref(value, env)

	// Comprobación de ocurrencia (occurs check): crucial para evitar ligar X a f(X) o listas cíclicas.
	// Esto es una implementación simplificada. Un occurs check completo es más complejo y recursivo.
	if valDeref.LogicalType == ds.LT_Variable && variable.ID == valDeref.ID {
		return nil // Ligando X a X, no-op
	}
	// Para un occurs check más robusto, se debería recorrer valDeref (si es estructura o lista)
	// para ver si 'variable' aparece dentro de ella. Esto es a menudo omitido en Prologs
	// por rendimiento (por ejemplo, Prolog estándar lo omite por defecto con el "occurs check off").

	env.AddBinding(variable, valDeref) // Añade la ligadura al entorno, que la registra en el trail
	return nil
}

// Unify intenta hacer que dos símbolos 'x' e 'y' sean lógicamente equivalentes
// realizando ligaduras de variables si es necesario, dentro de un entorno dado.
// Retorna true si la unificación es exitosa, false en caso contrario.
func Unify(x, y *ds.Symbol, env *Environment) bool {
	// 1. Desreferenciar ambos símbolos en el contexto del entorno actual.

	fmt.Printf("Unify called: x=%s, y=%s\n", x.String(), y.String()) // Add this
	x = Deref(x, env)
	y = Deref(y, env)
	fmt.Printf("Unify (dereferenced): x=%s, y=%s\n", x.String(), y.String()) // Add this

	// 2. Casos base de unificación
	if x == y { // Si los punteros son idénticos después de desreferenciar
		return true
	}

	// 3. Unificación con variables
	if x.LogicalType == ds.LT_Variable {
		if err := Bind(x, y, env); err != nil { // Llama a Bind con el entorno
			fmt.Printf("Error during variable bind for x: %v\n", err)
			return false
		}
		return true
	}

	if y.LogicalType == ds.LT_Variable {
		if err := Bind(y, x, env); err != nil { // Llama a Bind con el entorno
			fmt.Printf("Error during variable bind for y: %v\n", err)
			return false
		}
		return true
	}

	// 4. Casos especiales para símbolos predefinidos
	if x.LogicalType == ds.LT_Anonymous || y.LogicalType == ds.LT_Anonymous {
		return true
	}
	if x.LogicalType == ds.LT_Null && y.LogicalType == ds.LT_Null {
		return true
	}
	// Si uno es nulo y el otro no, no unifican.
	if (x.LogicalType == ds.LT_Null && y.LogicalType != ds.LT_Null) ||
		(y.LogicalType == ds.LT_Null && x.LogicalType != ds.LT_Null) {
		return false
	}

	// 5. Unificación de constantes
	if x.LogicalType == ds.LT_Constant && y.LogicalType == ds.LT_Constant {
		// Las constantes unifican si sus valores concretos son iguales.
		return x.Value == y.Value // Asegúrate de que Value sea comparable.
	}

	// 6. Unificación de listas
	if x.LogicalType == ds.LT_List && y.LogicalType == ds.LT_List {
		listX := x.Value.(*ds.ListPair)
		listY := y.Value.(*ds.ListPair)
		// Recursivamente unificar la cabeza y luego la cola, pasando el mismo entorno.
		return Unify(listX.Head, listY.Head, env) && Unify(listX.Tail, listY.Tail, env)
	}

	// 7. Unificación de estructuras (términos compuestos)
	if x.LogicalType == ds.LT_Structure && y.LogicalType == ds.LT_Structure {
		structX := x.Value.(*ds.StructureTerm)
		structY := y.Value.(*ds.StructureTerm)

		// El functor debe unificar (podría ser una variable también, o solo igual por ID)
		if !Unify(structX.Functor, structY.Functor, env) { // Unifica los functores
			return false
		}

		if len(structX.Args) != len(structY.Args) {
			return false
		}

		for i := 0; i < len(structX.Args); i++ {
			if !Unify(structX.Args[i], structY.Args[i], env) {
				return false
			}
		}
		return true
	}

	// 8. Tipos lógicos incompatibles que no se cubrieron.
	return false
}
