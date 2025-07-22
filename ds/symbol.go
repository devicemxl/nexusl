// Gothic/ds/symbol.go  (Enriquecido para motor de Unificación)
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

	// Símbolos predefinidos fundamentales para el motor de unificación
	NullSymbol      *Symbol // Representa la lista vacía o el término nulo
	AnonymousSymbol *Symbol // Representa la variable anónima (_)
)

func init() {
	fmt.Println("INFO: ds/symbol.go init() started.")
	// --- IMPORTANTE: ELIMINAR ESTAS LÍNEAS DE LOCK/UNLOCK AQUÍ ---
	// mu.Lock()
	// defer mu.Unlock()
	// -------------------------------------------------------------

	fmt.Println("INFO: ds/symbol.go init() - Initializing nextID, SymbolsByID, SymbolsByPublicName.")
	nextID = 1000 // Starting ID for symbols
	SymbolsByID = make(map[SymbolID]*Symbol)
	SymbolsByPublicName = make(map[string]*Symbol)

	// --- Debugging NullSymbol creation ---
	fmt.Println("INFO: ds/symbol.go init() - Creating NullSymbol.")
	// NewSymbol() ya maneja su propio bloqueo, no necesitamos el bloqueo externo de init.
	NullSymbol = NewSymbol()
	fmt.Println("INFO: ds/symbol.go init() - NullSymbol created. Assigning properties.")
	NullSymbol.PublicName = "nil"
	NullSymbol.Thing = LiteralType
	NullSymbol.LogicalType = LT_Null
	NullSymbol.State = Embodied
	// La asignación a SymbolsByPublicName debería ser segura porque NewSymbol ya bloqueó y desbloqueó
	// Sin embargo, si quieres que *esta asignación* esté protegida, necesitarías un LOCK/UNLOCK alrededor de ELLA,
	// o mover esta lógica a NewSymbol. Pero usualmente init se considera de un solo hilo.
	// Para simplificar, asumiremos que init es de un solo hilo y las llamadas a NewSymbol son suficientes.
	SymbolsByPublicName["nil"] = NullSymbol
	fmt.Println("INFO: ds/symbol.go init() - NullSymbol fully initialized.")

	// --- Debugging AnonymousSymbol creation ---
	fmt.Println("INFO: ds/symbol.go init() - Creating AnonymousSymbol.")
	AnonymousSymbol = NewSymbol()
	fmt.Println("INFO: ds/symbol.go init() - AnonymousSymbol created. Assigning properties.")
	AnonymousSymbol.PublicName = "_"
	AnonymousSymbol.Thing = IdentifierType
	AnonymousSymbol.LogicalType = LT_Anonymous
	AnonymousSymbol.State = Embodied
	SymbolsByPublicName["_"] = AnonymousSymbol
	fmt.Println("INFO: ds/symbol.go init() - AnonymousSymbol fully initialized.")

	fmt.Println("INFO: ds/symbol.go init() finished.")
}

// SymbolID es el identificador único para un Símbolo.
type SymbolID int

// SymbolState representa el estado del ciclo de vida de un Símbolo.
type SymbolState int

const (
	Exists   SymbolState = iota // El Símbolo existe (declarado como concepto).
	Embodied                    // El Símbolo tiene un valor concreto o está ligado.
)

// String devuelve la representación en cadena de SymbolState.
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
// Esto es para la clasificación del dominio (qué 'tipo de cosa' es).
type ThingType string

const (
	LiteralType      ThingType = "Literal"      // Para valores concretos (números, strings, bools).
	IdentifierType   ThingType = "Identifier"   // Para nombres de entidades, variables, etc.
	PredicateType    ThingType = "Predicate"    // Para predicados (has:, is:, do:).
	TripletScopeType ThingType = "TripletScope" // Para scopes de alto nivel (fact, program, func).
	MacroType        ThingType = "Macro"        // Para macros de lenguaje (que se expanden en AST).
	// Puedes añadir más ThingTypes aquí a medida que los necesites:
	// AgentType        ThingType = "Agent"
	// LocationType     ThingType = "Location"
	// EventType        ThingType = "Event"
)

// LogicalType define cómo se interpreta un Symbol en el contexto del motor de unificación.
// Esto es para la lógica interna de unificación.
type LogicalType int

const (
	LT_Undefined LogicalType = iota // Tipo lógico no especificado (por defecto al crear).
	LT_Variable                     // Una variable lógica que puede ser ligada (ej. X en Prolog).
	LT_Constant                     // Un valor atómico que no puede ser descompuesto (ej. 42, "hello", true).
	LT_List                         // Un símbolo que representa una lista (necesita Head/Tail).
	LT_Structure                    // Un símbolo que representa un término compuesto (necesita Functor/Args).
	LT_Anonymous                    // El símbolo de variable anónima (_).
	LT_Null                         // El símbolo que representa la lista vacía o el término nulo.
)

// String devuelve la representación en cadena de LogicalType.
func (lt LogicalType) String() string {
	switch lt {
	case LT_Undefined:
		return "Undefined"
	case LT_Variable:
		return "Variable"
	case LT_Constant:
		return "Constant"
	case LT_List:
		return "List"
	case LT_Structure:
		return "Structure"
	case LT_Anonymous:
		return "Anonymous"
	case LT_Null:
		return "Null"
	default:
		return fmt.Sprintf("UnknownLogicalType(%d)", lt)
	}
}

// SymbolProc es el tipo para un procedimiento (función Go) asociado a un Símbolo.
type SymbolProc func(args ...interface{}) (interface{}, error)

// Symbol representa la unidad atómica de referencia en el sistema.
// Cada Symbol denota una entidad individual y puede ser interpretado
// de diferentes maneras según su LogicalType y ThingType.
type Symbol struct {
	ID          SymbolID               // ID único del Símbolo.
	State       SymbolState            // Exists, Embodied.
	Thing       ThingType              // Categoría semántica del símbolo (ej. "Persona", "Función").
	LogicalType LogicalType            // Tipo lógico para el motor de unificación (Variable, Constant, List, etc.).
	PublicName  string                 // Nombre legible y único globalmente (opcional, usado para resolución).
	Value       interface{}            // El valor concreto que el símbolo podría encarnar (ej. string "robot-1", int 42, struct).
	Properties  map[string]interface{} // Almacén clave-valor para atributos adicionales.
	Proc        SymbolProc             // Si el símbolo representa un procedimiento ejecutable.
	Embedding   []float32              // El embedding vectorial del símbolo para búsquedas semánticas, etc.

	// Campos específicos para el manejo de variables lógicas en el motor de unificación:
	IsBound bool    // Indica si esta variable lógica ya está ligada a otro Símbolo.
	Binding *Symbol // Puntero al Símbolo al que está ligada esta variable.
	// Nota: Si 'Binding' está seteado, 'Value' de la variable NO contiene su valor ligado.
	// El valor ligado se obtiene desreferenciando 'Binding'.
}

// ListPair representa un par cons para construir listas en el motor de unificación.
// Será el 'Value' de un Symbol cuando su LogicalType sea LT_List.
type ListPair struct {
	Head *Symbol // El primer elemento de la lista.
	Tail *Symbol // El resto de la lista (otro *Symbol de tipo LT_List o LT_Null).
}

// StructureTerm representa una estructura o término compuesto para el motor de unificación.
// Será el 'Value' de un Symbol cuando su LogicalType sea LT_Structure.
type StructureTerm struct {
	Functor *Symbol   // El "nombre" o predicado de la estructura (ej. el símbolo para "f" en f(A,B)).
	Args    []*Symbol // Los argumentos de la estructura.
}

// NewSymbol crea una nueva instancia de Symbol única.
// Por defecto, se inicializa como un identificador no definido lógicamente.
func NewSymbol() *Symbol {
	mu.Lock()
	defer mu.Unlock()

	s := &Symbol{
		ID:          nextID,
		State:       Exists,
		Thing:       IdentifierType, // Default semántico
		LogicalType: LT_Undefined,   // Default lógico
		Properties:  make(map[string]interface{}),
	}
	SymbolsByID[s.ID] = s
	nextID++
	return s
}

// NewSymbolWithPublicName es un helper para crear y asignar nombre de una vez.
// Se inicializa con ThingType y LogicalType Undefined, a ser especificados.
func NewSymbolWithPublicName(name string, thingType ThingType) *Symbol {
	s := NewSymbol()
	s.AssignPublicName(name)
	s.SetThing(thingType)
	return s
}

// --- Constructores Específicos para el Motor de Unificación ---

// NewVariableSymbol crea un nuevo Symbol de tipo Variable Lógica.
func NewVariableSymbol(name string) *Symbol {
	s := NewSymbol()
	s.PublicName = name
	s.Thing = IdentifierType // Una variable es un tipo de identificador semántico
	s.LogicalType = LT_Variable
	s.IsBound = false
	s.State = Exists // Una variable existe como concepto, pero no está "Embodied" hasta que se liga
	return s
}

// NewConstantSymbol crea un nuevo Symbol de tipo Constante.
// El 'name' puede ser el valor mismo o un nombre representativo (ej. "42", "Juan").
func NewConstantSymbol(name string, value interface{}) *Symbol {
	s := NewSymbol()
	s.AssignPublicName(name) // El nombre público puede ser el valor string, o un alias
	s.Thing = LiteralType    // Una constante es un tipo de literal semántico
	s.LogicalType = LT_Constant
	s.Value = value
	s.State = Embodied // Una constante siempre tiene un valor concreto
	return s
}

// NewListSymbol crea un nuevo Symbol que representa un par cons de lista.
// `head` y `tail` deben ser *Symbol. Para una lista vacía, `tail` debe ser `NullSymbol`.
func NewListSymbol(head, tail *Symbol) *Symbol {
	s := NewSymbol()
	s.Thing = LiteralType // O podrías definir un ThingType específico como "ListThing"
	s.LogicalType = LT_List
	s.Value = &ListPair{Head: head, Tail: tail}
	s.State = Embodied
	return s
}

// NewStructureSymbol crea un nuevo Symbol que representa una estructura (término compuesto).
// `functor` debe ser un *Symbol (generalmente una constante/identificador), y `args` una slice de *Symbol.
func NewStructureSymbol(functor *Symbol, args []*Symbol) *Symbol {
	s := NewSymbol()
	s.PublicName = functor.PublicName // El nombre público de la estructura suele ser el del functor
	s.Thing = PredicateType           // Las estructuras a menudo representan predicados o relaciones
	s.LogicalType = LT_Structure
	s.Value = &StructureTerm{Functor: functor, Args: args}
	s.State = Embodied
	return s
}

// --- Métodos del Símbolo ---

// AssignPublicName asigna un nombre legible a un Símbolo.
// Asegura que el símbolo sea descubrible por este nombre y maneja las actualizaciones en el mapa global.
func (s *Symbol) AssignPublicName(name string) {
	mu.Lock()
	defer mu.Unlock()

	// Elimina la entrada antigua si el nombre cambia
	if s.PublicName != "" && SymbolsByPublicName[s.PublicName] == s {
		delete(SymbolsByPublicName, s.PublicName)
	}

	s.PublicName = name
	SymbolsByPublicName[name] = s
}

// LookupSymbolByPublicName recupera un Símbolo por su nombre público.
func LookupSymbolByPublicName(name string) (*Symbol, bool) {
	mu.Lock()
	defer mu.Unlock()
	sym, ok := SymbolsByPublicName[name]
	return sym, ok
}

// SetThing establece el "tipo" o "concepto" semántico que este símbolo representa.
func (s *Symbol) SetThing(thingType ThingType) {
	s.Thing = thingType
}

// InstantiateAs establece el valor concreto para el Símbolo, cambiando su estado a Embodied.
// Esto es para símbolos que encarnan un valor directo (constantes).
// Para variables lógicas, usa la función Bind.
func (s *Symbol) InstantiateAs(value interface{}) {
	s.Value = value
	s.State = Embodied
	// Asegúrate de que LogicalType sea apropiado (ej. LT_Constant)
	if s.LogicalType == LT_Undefined {
		s.LogicalType = LT_Constant // Por defecto, si no se ha especificado, asume que es constante
	}
}

// AddProperty añade o actualiza una propiedad clave-valor para el Símbolo.
func (s *Symbol) AddProperty(key string, value interface{}) {
	if s.Properties == nil { // Asegurarse de que el mapa esté inicializado
		s.Properties = make(map[string]interface{})
	}
	s.Properties[key] = value
}

// GetProperty recupera una propiedad por su clave.
func (s *Symbol) GetProperty(key string) (interface{}, bool) {
	val, ok := s.Properties[key]
	return val, ok
}

// CallProc intenta ejecutar el procedimiento (función Go) adjunto al Símbolo.
func (s *Symbol) CallProc(args ...interface{}) (interface{}, error) {
	if s.Proc == nil {
		return nil, fmt.Errorf("no procedure attached to symbol %s (ID: %d)", s.PublicName, s.ID)
	}
	return s.Proc(args...)
}

// String proporciona una representación legible del Símbolo.

// String proporciona una representación legible del Símbolo.
func (s *Symbol) String() string {
	valStr := "<nil>"
	if s.Value != nil {
		switch s.LogicalType {
		case LT_List:
			lp := s.Value.(*ListPair)
			// Usamos Deref de prologo.Deref aquí, lo cual es problemático.
			// Lo mejor es que String no haga Deref completo recursivo.
			// O hacer un Deref muy limitado que solo siga el .Binding
			// sin el concepto de Environment.

			// Para evitar circularidad en String, podemos usar una comprobación de recursión.
			// O, más simple, que String() muestre la estructura de la ligadura tal cual.
			// Implementación actual (puede causar loop si Deref lo hace):
			// valStr = fmt.Sprintf("[%s|%s]", lp.Head.String(), lp.Tail.String())

			// Versión segura para String():
			headStr := lp.Head.PublicName // O solo ID si no tiene nombre
			if lp.Head.PublicName == "" {
				headStr = fmt.Sprintf("anon:%d", lp.Head.ID)
			}
			tailStr := lp.Tail.PublicName // O solo ID
			if lp.Tail.PublicName == "" {
				tailStr = fmt.Sprintf("anon:%d", lp.Tail.ID)
			}
			// Esto no desreferencia, solo muestra los punteros inmediatos
			valStr = fmt.Sprintf("[%s|%s]", headStr, tailStr)

		case LT_Structure:
			st := s.Value.(*StructureTerm)
			argsStr := ""
			for i, arg := range st.Args {
				if i > 0 {
					argsStr += ", "
				}
				// arg.String() también puede causar recursión infinita
				// Para evitarlo, si un argumento es una variable ligada,
				// quizás solo mostrar su nombre o ID, no su contenido dereferenciado.
				argName := arg.PublicName
				if argName == "" {
					argName = fmt.Sprintf("anon:%d", arg.ID)
				}
				argsStr += argName
			}
			functorName := st.Functor.PublicName
			if functorName == "" {
				functorName = fmt.Sprintf("anon:%d", st.Functor.ID)
			}
			valStr = fmt.Sprintf("%s(%s)", functorName, argsStr)

		default:
			valStr = fmt.Sprintf("%v", s.Value)
		}
	} else if s.LogicalType == LT_Variable && s.IsBound && s.Binding != nil {
		// ¡ESTA ES LA LÍNEA CRÍTICA!
		// No podemos llamar a Deref de prologo.Deref aquí porque requiere un *Environment.
		// Y no podemos llamar a una Deref "global" sin env si no existe o es problemática.
		// La solución más simple para String(): simplemente muestra a qué está ligada la variable.
		// Esto *no* hará una desreferenciación completa recursiva, solo el siguiente paso.
		valStr = fmt.Sprintf("-> %s", s.Binding.PublicName) // O s.Binding.String() si el String() del Binding es seguro
		// Si quieres que String() muestre el valor final, tendrías que pasarle un Environment
		// o implementar un Deref solo para impresión que no requiera el entorno de unificación.
		// Por ahora, mostrar solo el nombre del binding es más seguro para evitar loops aquí.

		// Una opción más robusta para el String() de una variable ligada para evitar loops:
		// Usar un marcador para recursión.
		// Por ejemplo, puedes pasar un mapa de Symbols visitados.
		// Pero para empezar, solo imprimir el nombre del Binding es lo más seguro.
	}

	namePart := s.PublicName
	if namePart == "" {
		namePart = fmt.Sprintf("anon:%d", s.ID)
	}
	return fmt.Sprintf("[%s | ID:%d | %s | %s | %s | %s]", namePart, s.ID, s.LogicalType, s.State.String(), s.Thing, valStr)
}

// --- Funciones para el motor de Unificación ---

// Deref sigue las ligaduras de un Símbolo variable en el entorno.
// Si el Símbolo no es una variable, no está ligado, o ya es un valor concreto,
// devuelve el propio Símbolo. Esta es la función clave para obtener el valor "real" de una variable.
func Deref(s *Symbol) *Symbol {
	// Si es una variable y está ligada, sigue el Binding.
	if s.LogicalType == LT_Variable && s.IsBound && s.Binding != nil {
		// Llamada recursiva para seguir la cadena de ligaduras hasta el valor final
		// o la variable no ligada más profunda.
		return Deref(s.Binding)
	}
	// Si no es una variable, o es una variable no ligada, o su Binding es nulo,
	// o ya es un valor concreto (como una constante o estructura),
	// devuelve el Símbolo actual.
	return s
}

// Bind establece una ligadura para una variable.
// La variable 'variable' se liga al Símbolo 'value'.
// Retorna un error si se intenta ligar algo que no es una variable lógica.
func Bind(variable *Symbol, value *Symbol) error {
	if variable.LogicalType != LT_Variable {
		return fmt.Errorf("attempted to bind a non-variable Symbol: %s", variable.PublicName)
	}
	// Si la variable ya está ligada al mismo valor, no hacer nada.
	if variable.IsBound && variable.Binding == value {
		return nil
	}
	// Si la variable está intentando ligarse a sí misma, es una no-op.
	if variable.ID == value.ID {
		return nil
	}

	variable.IsBound = true
	variable.Binding = value
	variable.State = Embodied // Una variable ligada se considera "materializada"
	return nil
}
