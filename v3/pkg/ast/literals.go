// pkg/ast/literals.go

package ast

import (
	"strconv"

	"github.com/devicemxl/nexusl/pkg/token"
)

// StringLiteral representa un literal de cadena.
type StringLiteral struct {
	Token token.Token // El token STRING
	Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return "\"" + sl.Value + "\"" }

// IntegerLiteral representa un literal entero.
type IntegerLiteral struct {
	Token token.Token // El token NUMBER
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return strconv.FormatInt(il.Value, 10) }

// BooleanLiteral representa un literal booleano (si los añades)
// type BooleanLiteral struct { ... }

// NullLiteral representa el valor nulo (si lo añades como un nodo AST)
// type NullLiteral struct { ... }
