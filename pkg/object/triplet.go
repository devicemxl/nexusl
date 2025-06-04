// pkg/object/triplet.go

package object

import "fmt"

// TermType representa el tipo de un componente (sujeto, predicado, objeto) dentro de una Tripleta.
type TermType int

const (
	TermLiteral   TermType = iota // Un valor literal (string, int, bool)
	TermSymbol                    // Un símbolo (identificador sin valor concreto)
	TermAction                    // Una acción (cuando el objeto de la tripleta es una acción invocable) - Considera si esto es un SymbolTerm o BuiltinObject
	TermReference                 // Una referencia a otra tripleta o entidad
	TermVariable                  // Una variable lógica (para la unificación/inferencia)
)

// Term representa un componente (sujeto, predicado, objeto) de una Tripleta.
// Es una envoltura para los diferentes tipos de valores que pueden aparecer en una tripleta.
type Term struct {
	Type  TermType
	Value interface{} // Puede contener string (para símbolos/variables), int64, bool, o un Object (para literales complejos)
}

// NewSymbolTerm crea un Term de tipo Symbol.
func NewSymbolTerm(value string) *Term {
	return &Term{Type: TermSymbol, Value: value}
}

// NewLiteralTerm crea un Term de tipo Literal a partir de un Object de NexusL.
// Por ejemplo, un StringObject se convierte en un TermLiteral con el valor string.
func NewLiteralTerm(obj Object) *Term {
	switch o := obj.(type) {
	case *StringObject:
		return &Term{Type: TermLiteral, Value: o.Value}
	case *IntegerObject:
		return &Term{Type: TermLiteral, Value: o.Value}
	case *BooleanObject:
		return &Term{Type: TermLiteral, Value: o.Value}
	// Añadir otros tipos de objetos que puedan ser literales (ej. FloatObject si lo añades)
	default:
		// Si es un tipo de objeto que no se mapea directamente a un tipo primitivo para el Term,
		// podrías querer representar su String() o devolver un error.
		return &Term{Type: TermLiteral, Value: o.String()} // Fallback a la representación string
	}
}

// NewVariableTerm crea un Term de tipo Variable.
func NewVariableTerm(name string) *Term {
	return &Term{Type: TermVariable, Value: name}
}

// Stringer para Term para una buena representación textual.
func (t *Term) String() string {
	switch t.Type {
	case TermLiteral:
		// Formatea el literal. Si es un string, lo encierra en comillas simples.
		if s, ok := t.Value.(string); ok {
			return fmt.Sprintf("'%s'", s)
		}
		return fmt.Sprintf("%v", t.Value) // Para números, booleanos, etc.
	case TermSymbol:
		return t.Value.(string)
	case TermVariable:
		return fmt.Sprintf("?%s", t.Value.(string)) // Convención de Prolog para variables
	case TermAction: // Si el Value es un *ActionObject o similar
		return fmt.Sprintf("<action:%v>", t.Value)
	case TermReference: // Si el Value es una referencia a otro TripletaObject o ID
		return fmt.Sprintf("<ref:%v>", t.Value)
	default:
		return fmt.Sprintf("<TermType:%d Value:%v>", t.Type, t.Value)
	}
}

// TripletaObject representa un hecho o un patrón de conocimiento en la base de conocimiento de NexusL.
type TripletaObject struct {
	Subject   *Term
	Predicate *Term
	Object    *Term // El objeto en el conocimiento también es un Term
}

func (t *TripletaObject) Type() ObjectType { return TRIPLETA_OBJ }
func (t *TripletaObject) String() string {
	return fmt.Sprintf("(%s %s %s)",
		t.Subject.String(), // Usa String() del Term para una representación correcta
		t.Predicate.String(),
		t.Object.String(),
	)
}
