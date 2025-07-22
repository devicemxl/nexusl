package ast

import (
	"fmt"
	"strings"

	"github.com/devicemxl/nexusl/ds" // Asumimos que ds ya define Symbol y ThingType
	"github.com/devicemxl/nexusl/internal/Gothic/token"
)

// Node es la interfaz base para todos los nodos del AST
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement representa una sentencia (ej. una tripleta, una declaración de variable)
type Statement interface {
	Node
	statementNode() // Método dummy para marcar que es una sentencia
}

// Expression representa una expresión (ej. un literal, una operación)
type Expression interface {
	Node
	expressionNode() // Método dummy para marcar que es una expresión
}

// Program es el nodo raíz del AST
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
	var out strings.Builder
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// --- Nodos Específicos para `fact Car is symbol;` ---

// FactStatement representa una declaración de `fact` atómica.
type FactStatement struct {
	Token     token.Token // El token 'fact'
	Scope     *ds.Symbol  // Referencia al Symbol del scope "fact"
	Subject   Expression
	Predicate Expression // <--- ¡CAMBIO CLAVE AQUÍ! Ahora es Expression
	Object    Expression // <--- ¡CAMBIO CLAVE AQUÍ! Ahora es Expression
}

func (fs *FactStatement) statementNode()       {}
func (fs *FactStatement) TokenLiteral() string { return fs.Token.Word } // Debería ser "fact"
func (fs *FactStatement) String() string {
	return fmt.Sprintf("%s %s %s %s;", fs.TokenLiteral(), fs.Subject.String(), fs.Predicate.String(), fs.Object.String())
}

// Identifier representa un identificador (como "Car" o "symbol")
type Identifier struct {
	Token token.Token // El token IDENTIFIER
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Word }
func (i *Identifier) String() string       { return i.Value }

// Simplemente para que puedas ver los valores en el AST
// Más adelante, "symbol" no será solo un Identifier, sino un Symbol Type.
// Esto es para que el parser pueda manejarlo por ahora.
// En un sistema real, 'symbol' podría ser una referencia a un tipo de Thing.

// --- NUEVAS DEFINICIONES DE LITERALES ---

// StringLiteral representa un valor de cadena literal.
// Ej: "red", "kitchen", "david"
type StringLiteral struct {
	Token token.Token // El token STRING
	Value string      // El valor de la cadena (ej. "red")
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Word }
func (sl *StringLiteral) String() string       { return fmt.Sprintf("%q", sl.Value) } // Formatea con comillas para visualización

// IntegerLiteral representa un valor numérico entero literal.
// Ej: 123, 42
type IntegerLiteral struct {
	Token token.Token // El token INT
	Value int64       // El valor entero (ej. 123)
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Word }
func (il *IntegerLiteral) String() string       { return il.Token.Word } // Devuelve la representación en cadena del número

// FloatLiteral representa un valor numérico de punto flotante literal.
// Ej: 1.23, 3.14159
type FloatLiteral struct {
	Token token.Token // El token FLOAT
	Value float64     // El valor flotante (ej. 1.23)
}

func (fl *FloatLiteral) expressionNode()      {}
func (fl *FloatLiteral) TokenLiteral() string { return fl.Token.Word }
func (fl *FloatLiteral) String() string       { return fl.Token.Word } // Devuelve la representación en cadena del número

// BooleanLiteral representa un valor booleano literal.
// Ej: true, false
type BooleanLiteral struct {
	Token token.Token // El token BOOLEAN
	Value bool        // El valor booleano (true o false)
}

func (bl *BooleanLiteral) expressionNode()      {}
func (bl *BooleanLiteral) TokenLiteral() string { return bl.Token.Word }
func (bl *BooleanLiteral) String() string       { return bl.Token.Word } // Devuelve "true" o "false"
