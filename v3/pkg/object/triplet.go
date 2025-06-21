// triplet.go
package object

import "fmt"

// TermType representa el tipo de un componente (sujeto, predicado, objeto) dentro de una Tripleta.
type TermType int

const (
	TermLiteral   TermType = iota // Un valor literal (string, int, bool)
	TermSymbol                    // Un símbolo (identificador sin valor concreto)
	TermAction                    // Una acción (cuando el objeto de la tripleta es una acción invocable)
	TermReference                 // Una referencia a otra tripleta o entidad
	TermVariable                  // Una variable lógica (para la unificación/inferencia)
)

// Term representa un componente (sujeto, predicado, objeto) de una Tripleta.
type Term struct {
	Type  TermType
	Value interface{} // Puede contener string (para símbolos/variables), int64, bool, o un Object
}

func NewSymbolTerm(value string) *Term {
	return &Term{Type: TermSymbol, Value: value}
}

func NewLiteralTerm(obj Object) *Term {
	switch o := obj.(type) {
	case *StringObject:
		return &Term{Type: TermLiteral, Value: o.Value}
	case *IntegerObject:
		return &Term{Type: TermLiteral, Value: o.Value}
	case *BooleanObject:
		return &Term{Type: TermLiteral, Value: o.Value}
	case *NullObject:
		return &Term{Type: TermLiteral, Value: nil} // Representar null como nil o un valor sentinel
	// Agrega otros tipos de literales si los tienes (ej. FloatObject)
	default:
		return &Term{Type: TermLiteral, Value: obj.String()} // Fallback para cualquier otro objeto
	}
}

func NewVariableTerm(name string) *Term {
	return &Term{Type: TermVariable, Value: name}
}

// Stringer para Term para una buena representación textual.
func (t *Term) String() string {
	switch t.Type {
	case TermLiteral:
		if s, ok := t.Value.(string); ok {
			return fmt.Sprintf("'%s'", s) // Citas para strings
		}
		return fmt.Sprintf("%v", t.Value)
	case TermSymbol:
		return t.Value.(string)
	case TermVariable:
		return fmt.Sprintf("?%s", t.Value.(string))
	case TermAction:
		return fmt.Sprintf("<action:%v>", t.Value)
	case TermReference:
		return fmt.Sprintf("<ref:%v>", t.Value)
	default:
		return fmt.Sprintf("<TermType:%d Value:%v>", t.Type, t.Value)
	}
}

// TripletaObject representa un hecho o un patrón de conocimiento.
type TripletaObject struct {
	Subject   *Term
	Predicate *Term
	Object    *Term
}

func (t *TripletaObject) Type() ObjectType { return TRIPLETA_OBJ }
func (t *TripletaObject) String() string {
	return fmt.Sprintf("(%s %s %s)", t.Subject.String(), t.Predicate.String(), t.Object.String())
}
