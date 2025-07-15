package ds

import (
	"fmt"
	"strings" // Necesario para strings.Join en el método String()
)

// Triplet representa una relación en el NexusMesh.
// Es la unidad fundamental para representar hechos, acciones, definiciones
// y la estructura interna del "sea of nodes".
type Triplet struct {
	// Scope: Un Symbol que define el contexto o la semántica de la tripleta (ej. 'program', 'fact', 'func', 'def').
	// Indica el "tipo" de afirmación o acción que representa la tripleta.
	Scope *Symbol
	// Subject (S): Siempre es un Symbol. La entidad principal de la tripleta.
	Subject *Symbol
	// Predicate (P): Puede ser un Symbol (para predicados simples) o una estructura compleja.
	// Ejemplos: Symbol (para 'is'), []interface{} (para listas de acciones),
	// map[string]interface{} (para atributos complejos como how::).
	Predicate interface{}
	// Object (O): Puede ser un Symbol (para objetos simples) o una estructura compleja.
	// Ejemplos: Symbol (para 'david'), []interface{} (para colecciones),
	// Triplet (para tripletas anidadas como objetos).
	Object interface{}
}

// NewTriplet crea una nueva instancia de Triplet con los componentes dados.
// Permite gran flexibilidad al aceptar cualquier tipo para Predicate y Object
// gracias a su definición como interface{}.
func NewTriplet(s *Symbol, p interface{}, o interface{}, scope *Symbol) *Triplet {
	return &Triplet{
		Subject:   s,
		Predicate: p,
		Object:    o,
		Scope:     scope,
	}
}

// formatInterfaceValue es una función auxiliar para formatear los valores de interface{}
// de manera recursiva y legible para el método String().
func formatInterfaceValue(val interface{}) string {
	if val == nil {
		return "<nil>"
	}

	switch v := val.(type) {
	case *Symbol:
		return v.String()
	case *Triplet:
		return v.String()
	case []interface{}:
		elements := make([]string, len(v))
		for i, elem := range v {
			elements[i] = formatInterfaceValue(elem)
		}
		return fmt.Sprintf("[%s]", strings.Join(elements, " "))
	case map[string]interface{}:
		pairs := make([]string, 0, len(v))
		for key, elem := range v {
			pairs = append(pairs, fmt.Sprintf("%s::%s", key, formatInterfaceValue(elem)))
		}
		return fmt.Sprintf("{%s}", strings.Join(pairs, " "))
	default:
		return fmt.Sprintf("%v", v)
	}
}

// String devuelve una representación en cadena de la Triplet para facilitar la depuración y visualización.
func (t *Triplet) String() string {
	subjectStr := formatInterfaceValue(t.Subject)
	predicateStr := formatInterfaceValue(t.Predicate)
	objectStr := formatInterfaceValue(t.Object)

	scopeStr := "nil"
	if t.Scope != nil {
		scopeStr = t.Scope.PublicName // Assuming Symbol has PublicName, or use Name if it's the public one
	}

	return fmt.Sprintf("(%s %s %s) [%s]", subjectStr, predicateStr, objectStr, scopeStr)
}
