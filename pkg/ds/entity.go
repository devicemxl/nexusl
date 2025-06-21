package ds

import (
	"errors" // Importa el paquete errors para manejar errores
	"fmt"    // Importa el paquete fmt para formateo de salida
	"sync"   // Importa el paquete sync para sincronización (manejo de concurrencia)
)

// EntityState representa el estado de existencia o realización de un Entidad.
// Sirve para distinguir entre un concepto abstracto (exists) y una entidad concreta (embodied).
type EntityState int

const (
	// exists indica que un Entidad es un concepto abstracto o un placeholder.
	// Ej: La idea de "un testigo", sin saber aún quién es.
	exists EntityState = iota
	// embodied indica que un Entidad ha sido instanciado con un valor concreto.
	// Ej: "Juan Pérez" es el testigo específico.
	embodied
)

// String devuelve la representación en cadena del EntityState.
// Esto mejora la legibilidad al imprimir el estado del Entidad.
func (s EntityState) String() string {
	switch s {
	case exists:
		return "exists"
	case embodied:
		return "embodied"
	default:
		return fmt.Sprintf("UnknownState(%d)", s)
	}
}

// EntityID es un tipo para identificar un Entidad de forma única.
// Se usa un uint32 para IDs positivos y eficientes.
type EntityID uint32

// Entity representa la unidad atómica de referencia en el sistema.
// Cada Entity denota una entidad individual en el dominio de discurso.
type Entity struct {
	ID         EntityID    // ID único del Entidad.
	PublicName string      // Nombre legible o identificador externo del Entidad (ej: "Juan Pérez").
	State      EntityState // Estado del Entidad: ¿existe como concepto o está instanciado?
	Thing      string      // Tipo general de cosa que el Entidad representa (ej: "Persona", "Función", "Vehículo").
	// Funciona como un esquema o categoría conceptual para el Entidad, similar a Schema.org.
	Value      interface{}            // Valor concreto si el Entidad ha sido instanciado (ej: el nombre "Juan Pérez", el número 21).
	Properties map[string]interface{} // Mapa para almacenar propiedades arbitrarias del Entidad (ej: "color": "red").
	// Permite adjuntar atributos clave-valor adicionales al Entidad, útil para modelar esquemas flexibles.
	Proc func(args ...interface{}) (interface{}, error) // Procedimiento o función asociada al Entidad.
	// Esto permite que un Entidad represente una operación ejecutable (como una "fórmula" o "verbo").
}

var (
	// EntitysByID es un mapa que almacena todos los Entidades creados, accesibles por su ID único.
	EntitysByID = make(map[EntityID]*Entity)
	// nextID mantiene el siguiente ID disponible para un nuevo Entidad.
	nextID EntityID = 100 // Empieza en 100 para dejar IDs bajos libres si se quieren para algo especial.
	// mu es un Mutex para proteger el acceso concurrente a nextID y EntitysByID.
	// Garantiza la seguridad de hilos al crear nuevos Entidades o acceder al mapa global.
	mu sync.Mutex
)

// NewEntity crea un nuevo Entidad anónimo con un ID único y un estado inicial 'exists'.
// Este Entidad es un placeholder abstracto hasta que se le asigna un nombre o valor concreto.
func NewEntity() *Entity {
	mu.Lock()         // Bloquea el Mutex para acceso exclusivo a variables globales.
	defer mu.Unlock() // Desbloquea el Mutex al salir de la función.

	id := nextID // Asigna el siguiente ID disponible.
	nextID++     // Incrementa el contador para el siguiente Entidad.

	s := &Entity{
		ID:         id,
		State:      exists,                       // Por defecto, un Entidad nuevo existe como concepto.
		Thing:      "Thing",                      // Tipo genérico por defecto, a ser redefinido.
		Properties: make(map[string]interface{}), // Inicializa el mapa de propiedades.
	}

	EntitysByID[id] = s // Almacena el nuevo Entidad en el mapa global.

	return s // Devuelve el puntero al nuevo Entidad.
}

// AssignPublicName asigna un nombre legible a un Entidad.
// Esto permite referenciar el Entidad de forma más amigable que por su ID.
func (s *Entity) AssignPublicName(name string) {
	s.PublicName = name
	// Aquí podrías agregar lógica para persistir el Entidad en una KB (Base de Conocimiento)
	// Por ejemplo, enviar un triple "Entity_named(s.ID, name)" a un store de Datalog.
}

// SetThing establece el tipo o categoría de cosa que el Entidad representa.
// Esto es clave para el modelado semántico, similar a un concepto de Schema.org o una clase OWL.
func (s *Entity) SetThing(thing string) {
	s.Thing = thing
}

// InstantiateAs asigna un valor concreto a un Entidad, cambiando su estado a 'embodied'.
// Esto transforma el Entidad de un placeholder a una entidad con un referente específico.
func (s *Entity) InstantiateAs(val interface{}) {
	s.State = embodied // El Entidad ahora está "materializado".
	s.Value = val      // Se le asigna su valor concreto (ej: "Juan Pérez", 42, un struct).
}

// AddProperty añade una propiedad clave-valor al mapa de propiedades del Entidad.
// Permite añadir metadatos o atributos adicionales que no son parte de los campos fijos del Entity.
func (s *Entity) AddProperty(key string, value interface{}) {
	if s.Properties == nil {
		s.Properties = make(map[string]interface{}) // Inicializa el mapa si es nil.
	}
	s.Properties[key] = value // Agrega o actualiza la propiedad.
}

// GetProperty obtiene el valor de una propiedad específica del Entidad.
// Devuelve el valor y un booleano indicando si la propiedad existe.
func (s *Entity) GetProperty(key string) (interface{}, bool) {
	val, ok := s.Properties[key]
	return val, ok
}

// String devuelve una representación en cadena del Entidad para facilitar la depuración y visualización.
// Muestra el nombre público (o ID anónimo), estado, tipo y valor.
func (s *Entity) String() string {
	name := s.PublicName // Usa el nombre público si está asignado.
	if name == "" {
		name = fmt.Sprintf("anon:%d", s.ID) // Si no, usa un nombre anónimo basado en el ID.
	}
	return fmt.Sprintf("[%s | %s | %s | %v]", name, s.State, s.Thing, s.Value)
}

// CallProc ejecuta el procedimiento (función Go) asociado a un Entidad.
// Esto es útil para Entidades que representan operaciones, funciones o acciones ejecutables.
func (s *Entity) CallProc(args ...interface{}) (interface{}, error) {
	if s.Proc == nil {
		return nil, errors.New("no procedure attached") // Error si no hay procedimiento asignado.
	}
	return s.Proc(args...) // Ejecuta el procedimiento con los argumentos dados.
}

/*
func main() {
	// Punto de entrada principal para la demostración de Entity.


		ESTA ES LA DEFINICION PARA LA ASIGNACION DE FORMULAS
		Este bloque demuestra cómo un Entidad puede representar una "fórmula"
		o una "función ejecutable" en el sistema.
		El Entidad 'doble' se asocia con un procedimiento Go que duplica un número.

	fmt.Printf("\nESTA ES LA DEFINICION PARA LA ASIGNACION DE VERBOS (FORMULAS)\n")
	// Crear un Entidad anónimo para representar la operación "doble".
	s := NewEntity()

	// Asignar un nombre público al Entidad, actuando como el nombre de la "fórmula" o "verbo".
	s.AssignPublicName("doble")

	// Establecer el tipo de cosa que el Entidad representa: en este caso, una "Función".
	s.SetThing("Function")

	// Asignar un procedimiento (una función anónima de Go) al campo Proc del Entidad.
	// Este procedimiento define la lógica de la "fórmula": duplicar el número recibido.
	s.Proc = func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("expected 1 argument, got %d", len(args))
		}
		switch v := args[0].(type) {
		case int:
			return v * 2, nil
		default:
			return nil, fmt.Errorf("unsupported type: %T", v)
		}
	}

	// Llamar al procedimiento asociado al Entidad 'doble' con el argumento 21.
	res, err := s.CallProc(21)
	if err != nil {
		fmt.Println("Error executing proc:", err)
	} else {
		fmt.Printf("Result of %s(21) = %v\n", s.PublicName, res)
	}


		ESTA ES LA DEFINICION PARA MANEJO DE SUJETOS
		Este bloque demuestra cómo los Entidades pueden representar "sujetos" o entidades
		concretas en el sistema, a las que se les pueden asignar propiedades y valores.

	fmt.Printf("\nESTA ES LA DEFINICION PARA LA ASIGNACION DE SUJETOS\n")
	fmt.Println("\n--- Otro Entidad ---")

	// Crear un Entidad anónimo para un testigo abstracto.
	// Inicialmente, solo se afirma su existencia como un placeholder.
	witnessEntity := NewEntity()
	fmt.Printf("Entidad creado: %s\n", witnessEntity)
	// Añadir propiedades descriptivas al Entidad.
	witnessEntity.AddProperty("role", "witness")
	witnessEntity.AddProperty("saw_red_car", true)
	fmt.Printf("Propiedades del Entidad: %v\n", witnessEntity.Properties)

	// Asignar un nombre público y un tipo más específico al Entidad 'witnessEntity'.
	witnessEntity.AssignPublicName("Juan Pérez")
	witnessEntity.SetThing("Human")
	// Instanciar el Entidad con un valor concreto, haciéndolo 'embodied'.
	witnessEntity.InstantiateAs("Juan Pérez")
	fmt.Printf("Entidad instanciado: %s\n", witnessEntity)

	fmt.Println("\n--- Otro Entidad ---")

	// Crear otro Entidad anónimo para un coche abstracto.
	carEntity := NewEntity()
	// Añadir propiedades al Entidad del coche.
	carEntity.AddProperty("color", "red")
	carEntity.AddProperty("unknown_brand", true) // Indica que la marca es desconocida.
	fmt.Printf("Entidad de coche: %s\n", carEntity)

	// Asignar un nombre público y un tipo más específico al Entidad 'carEntity'.
	carEntity.AssignPublicName("Nissan Sentra 2015")
	carEntity.SetThing("Vehicle")
	// Instanciar el Entidad del coche con un valor concreto.
	carEntity.InstantiateAs("Nissan Sentra 2015")
	fmt.Printf("Entidad de coche instanciado: %s\n", carEntity)
	fmt.Printf("\n\n")

	// Crear una nueva entidad
	entity := NewEntity()
	entity.AssignPublicName("David")
	entity.SetThing("Person")

	// Definir una función para MainVerb
	run := func(e *Entity) (interface{}, error) {
		return fmt.Sprintf("%s do %s", e.PublicName, e.Properties["action"]), nil
	}

	// Asignar la propiedad "action" a la entidad
	entity.AddProperty("action", "run")

	// Crear una tripleta con la función MainVerb
	triplet := FlatTripletStatement{
		Scope:   "def",
		Subject: entity.ID,
		Verb: VerbStatement{
			ModalVerb: "DO",
			MainVerb:  run,
		},
		Object: ObjectStatement{
			Condition: "HOW",
			Statement: "fast",
		},
	}
	// Ejecutar la función MainVerb
	result, err := triplet.Verb.MainVerb(entity)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
	fmt.Printf("\n\n\n")

}
*/
