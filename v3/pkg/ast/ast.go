// pkg/ast/ast.go

package ast

import (
	"bytes"
	// No necesitas importar token aquí si no lo usas directamente,
	// solo las structs que lo usan (que ahora están en otros archivos)
)

// Node es la interfaz base para todos los nodos del AST.
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement representa una sentencia en el AST.
type Statement interface {
	Node
	statementNode() // Método dummy para identificar como Statement
}

// Expression representa una expresión en el AST.
type Expression interface {
	Node
	expressionNode() // Método dummy para identificar como Expression
}

// Program es el nodo raíz de todo el AST.
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}
