// pkg/ast/expressions.go

package ast

import (
	"github.com/devicemxl/nexusl/pkg/token"
)

// Identifier representa un identificador en el AST.
type Identifier struct {
	Token token.Token // El token IDENTIFIER
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// CallExpression (futura)
// type CallExpression struct { ... }

// RuleExpression (futura)
// type RuleExpression struct { ... }
