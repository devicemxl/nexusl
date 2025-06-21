// pkg/object/object.go
/*
archivo principal de la interfaz y tipos base)

Este archivo debería contener principalmente:

La interfaz Object.
La interfaz Term y TermType (si es que quieres separarla de triplet.go para reusabilidad, aunque está bien en triplet.go por ahora).
Tipos básicos como NullObject y ErrorObject.
Constantes ObjectType y BuiltinFunction.
*/

package object

import "fmt"

// ObjectType representa el tipo de un objeto de NexusL.
type ObjectType string

const (
	NULL_OBJ     = "NULL"
	ERROR_OBJ    = "ERROR"
	NUMBER_OBJ   = "NUMBER" // Para IntegerObject o FloatObject
	STRING_OBJ   = "STRING"
	BOOLEAN_OBJ  = "BOOLEAN"
	SYMBOL_OBJ   = "SYMBOL"   // Para SymbolObject
	BUILTIN_OBJ  = "BUILTIN"  // Para BuiltinObject
	TRIPLETA_OBJ = "TRIPLETA" // Para TripletaObject

	// Agrega aquí otros tipos de objetos a medida que los desarrolles
	// FUNCTION_OBJ, LIST_OBJ, etc.
)

// Object es la interfaz que todos los valores en NexusL deben implementar.
type Object interface {
	Type() ObjectType
	String() string
}

// NullObject representa el valor nulo en NexusL.
var NULL = &NullObject{} // Instancia única de Null

type NullObject struct{}

func (n *NullObject) Type() ObjectType { return NULL_OBJ }
func (n *NullObject) String() string   { return "null" }

// ErrorObject representa un error de ejecución.
type ErrorObject struct {
	Message string
}

func (e *ErrorObject) Type() ObjectType { return ERROR_OBJ }
func (e *ErrorObject) String() string   { return "ERROR: " + e.Message }

// NewError es una función de ayuda para crear un nuevo ErrorObject.
func NewError(format string, a ...interface{}) *ErrorObject {
	return &ErrorObject{Message: fmt.Sprintf(format, a...)}
}

/*


// BuiltinFunction representa una función nativa del sistema.
type BuiltinFunction func(args ...Object) Object // Puede tomar múltiples argumentos, o un solo argumento Object

// BuiltinObject representa una función builtin de NexusL.
// Definido aquí para que BuiltinFunction pueda ser usada en la interfaz.
// Podría moverse a builtin.go también.
type BuiltinObject struct {
	Fn BuiltinFunction
}

func (b *BuiltinObject) Type() ObjectType { return BUILTIN_OBJ }
func (b *BuiltinObject) String() string   { return "builtin function" }
*/
