// pkg/object/literals.go
/*
(Para los tipos de datos primitivos)

Este archivo debería contener:

StringObject
IntegerObject (o NumberObject si lo unificas).
BooleanObject.
*/
package object

import (
	"fmt"
	"strconv"
)

// StringObject representa una cadena de texto.
type StringObject struct {
	Value string
}

func (s *StringObject) Type() ObjectType { return STRING_OBJ }
func (s *StringObject) String() string   { return fmt.Sprintf("%q", s.Value) } // Mejor para depuración

// IntegerObject representa un número entero.
type IntegerObject struct {
	Value int64
}

func (i *IntegerObject) Type() ObjectType { return NUMBER_OBJ } // Usamos NUMBER_OBJ para IntegerObject
func (i *IntegerObject) String() string   { return strconv.FormatInt(i.Value, 10) }

// BooleanObject representa un valor booleano (true/false).
type BooleanObject struct {
	Value bool
}

func (b *BooleanObject) Type() ObjectType { return BOOLEAN_OBJ }
func (b *BooleanObject) String() string   { return strconv.FormatBool(b.Value) }

// Instancias globales para True y False (similar a NULL)
var (
	TRUE  = &BooleanObject{Value: true}
	FALSE = &BooleanObject{Value: false}
)
