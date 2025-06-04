// ast.go
package main

import (
	"bytes"
	"fmt"
	"strings"
)

// Node es la interfaz base para todos los nodos del AST.
// Cada nodo del AST debe implementar un método TokenLiteral() string para depuración.
type Node interface {
	TokenLiteral() string // Devuelve el literal del token asociado al nodo.
	String() string       // Devuelve una representación en cadena del nodo (para depuración).
}

// Statement es la interfaz para todos los nodos que representan sentencias.
// Las sentencias no producen un valor directamente en la evaluación.
type Statement interface {
	Node
	statementNode() // Método dummy para marcar que es un Statement.
}

// Expression es la interfaz para todos los nodos que representan expresiones.
// Las expresiones producen un valor en la evaluación.
type Expression interface {
	Node
	expressionNode() // Método dummy para marcar que es una Expression.
}

// Program es el nodo raíz de nuestro AST.
// Contiene una lista de sentencias.
type Program struct {
	Statements []Statement
}

// TokenLiteral para Program (normalmente el del primer token, pero puede ser vacío).
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// String para Program (útil para depuración, imprime el programa completo).
func (p *Program) String() string {
	var out []byte
	for _, s := range p.Statements {
		out = append(out, s.String()...)
	}
	return string(out)
}

// --- Nodos Específicos para la Tripleta Aplanada ---

// Identifier representa un identificador (como 'david', 'corre', 'rapido').
type Identifier struct {
	Token Token // El token IDENT.
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// StringLiteral representa una cadena literal (ej. "hello").
type StringLiteral struct {
	Token Token // El token STRING.
	Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return "\"" + sl.Value + "\"" } // Para que se imprima con comillas.

// IntegerLiteral representa un número entero (ej. 42).
type IntegerLiteral struct {
	Token Token // El token INT.
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal } // El literal ya es la representación.

// FloatLiteral representa un número flotante (ej. 3.14).
type FloatLiteral struct {
	Token Token // El token FLOAT.
	Value float64
}

func (fl *FloatLiteral) expressionNode()      {}
func (fl *FloatLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FloatLiteral) String() string       { return fl.Token.Literal }

// BooleanLiteral represents a boolean value (true/false).
type BooleanLiteral struct {
	Token Token // The token TRUE or FALSE.
	Value bool
}

func (b *BooleanLiteral) expressionNode()      {}
func (b *BooleanLiteral) TokenLiteral() string { return b.Token.Literal }
func (b *BooleanLiteral) String() string       { return b.Token.Literal }

// NilLiteral represents the null-like value.
type NilLiteral struct {
	Token Token // The token NIL.
}

func (n *NilLiteral) expressionNode()      {}
func (n *NilLiteral) TokenLiteral() string { return n.Token.Literal }
func (n *NilLiteral) String() string       { return "nil" }

// FlatTripletStatement representa una sentencia de tripleta aplanada: sujeto verbo atributo;
type FlatTripletStatement struct {
	Token   Token       // El token del sujeto (e.g., IDENT de "david")
	Subject *Identifier // El sujeto de la tripleta (e.g., david)
	Verb    *Identifier // El verbo/predicado (e.g., corre)
	Object  Expression  // El objeto (e.g., rapido, "rápido", 123)
}

func (fts *FlatTripletStatement) statementNode()       {}
func (fts *FlatTripletStatement) TokenLiteral() string { return fts.Token.Literal }
func (fts *FlatTripletStatement) String() string {
	return fts.Subject.String() + " " + fts.Verb.String() + " " + fts.Object.String() + ";"
}

// --- Nuevas Estructuras para la Tripleta Cualificada ---

// SymbolIdentifier representa un identificador que comienza con '$'.
type SymbolIdentifier struct {
	Token Token  // El token SYMSIGN
	Value string // El nombre del símbolo (e.g., "david")
}

func (si *SymbolIdentifier) expressionNode()      {}
func (si *SymbolIdentifier) TokenLiteral() string { return si.Token.Literal }
func (si *SymbolIdentifier) String() string       { return "$" + si.Value } // Para que se imprima como $david

// QualifiedProperty representa un par clave-valor como 'do:"corre"' o 'how:rapido'.
type QualifiedProperty struct {
	Key    *Identifier // El calificador (e.g., "do", "how", "when")
	Object Expression  // El valor asociado (e.g., "corre" (StringLiteral), rapido (Identifier))
}

func (qp *QualifiedProperty) String() string {
	return qp.Key.String() + ":" + qp.Object.String()
}

// QualifiedTripletStatement representa la sentencia: $sujeto Propiedad:Valor Propiedad:Valor;
type QualifiedTripletStatement struct {
	Token      Token                // El token '$' que inicia la sentencia
	Subject    *SymbolIdentifier    // El identificador del sujeto (e.g., "$david")
	Qualifiers []*QualifiedProperty // Lista de calificadores (do:"corre", how:"rapido")
}

func (qts *QualifiedTripletStatement) statementNode()       {}
func (qts *QualifiedTripletStatement) TokenLiteral() string { return qts.Token.Literal }
func (qts *QualifiedTripletStatement) String() string {
	var out string
	out += qts.Subject.String() // Print $subject

	for _, qp := range qts.Qualifiers {
		out += " " + qp.String() // Print each qualified property
	}
	out += ";" // Add the terminator
	return out
}

// --- Nueva Estructura para Listas ---

// ListLiteral representa una lista de expresiones (e.g., [1, "two", $symbol, (a b c)]).
type ListLiteral struct {
	Token    Token        // El token LBRACKET '['
	Elements []Expression // Las expresiones dentro de la lista
}

func (ll *ListLiteral) expressionNode()      {}
func (ll *ListLiteral) TokenLiteral() string { return ll.Token.Literal }
func (ll *ListLiteral) String() string {
	var out bytes.Buffer
	elements := []string{}
	for _, el := range ll.Elements {
		elements = append(elements, el.String())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

// --- Nuevas Estructuras para Definiciones de Función ---

// FunctionDefinitionStatement representa una declaración de definición de función.
// Ejemplo: func:calculaArea param:[base has Value is int;] code:[c = a * b;] export:[D];
type FunctionDefinitionStatement struct {
	Token         Token                  // El token FUNC
	Name          *Identifier            // Nombre de la función (e.g., "calculaArea")
	Parameters    []*Parameter           // Lista de parámetros
	Body          []*AssignmentStatement // Cuerpo de la función (por ahora, asignaciones)
	ExportedValue Expression             // Valor a exportar (puede ser un identificador o una expresión)
}

func (fds *FunctionDefinitionStatement) statementNode()       {}
func (fds *FunctionDefinitionStatement) TokenLiteral() string { return fds.Token.Literal }
func (fds *FunctionDefinitionStatement) String() string {
	var out bytes.Buffer
	out.WriteString(fds.TokenLiteral()) // "func"
	out.WriteString(":")
	out.WriteString(fds.Name.String()) // "calculaArea"
	out.WriteString(" param:[")
	for i, p := range fds.Parameters {
		out.WriteString(p.String())
		if i < len(fds.Parameters)-1 {
			out.WriteString("; ")
		}
	}
	out.WriteString("] code:[")
	for i, s := range fds.Body {
		out.WriteString(s.String())
		if i < len(fds.Body)-1 {
			out.WriteString("; ")
		}
	}
	out.WriteString("] export:[")
	if fds.ExportedValue != nil {
		out.WriteString(fds.ExportedValue.String())
	}
	out.WriteString("];")
	return out.String()
}

// Parameter representa un parámetro de función.
// Ejemplo: base has Value is int
type Parameter struct {
	Token Token       // El token IDENT del nombre del parámetro
	Name  *Identifier // Nombre del parámetro (e.g., "base")
	Type  *Identifier // Tipo del parámetro (e.g., "int")
	// Podrías añadir un campo para el valor por defecto si lo necesitas
}

func (p *Parameter) statementNode()       {} // Los parámetros son parte de la declaración de función, no sentencias independientes
func (p *Parameter) expressionNode()      {} // Los parámetros no son expresiones en sí mismos
func (p *Parameter) TokenLiteral() string { return p.Token.Literal }
func (p *Parameter) String() string {
	return fmt.Sprintf("%s has Value is %s", p.Name.String(), p.Type.String())
}

// AssignmentStatement representa una sentencia de asignación simple.
// Ejemplo: c = a * b;
type AssignmentStatement struct {
	Token Token       // El token IDENT del nombre de la variable
	Name  *Identifier // Nombre de la variable asignada
	Value Expression  // La expresión asignada a la variable
}

func (as *AssignmentStatement) statementNode()       {}
func (as *AssignmentStatement) TokenLiteral() string { return as.Token.Literal }
func (as *AssignmentStatement) String() string {
	var out bytes.Buffer
	out.WriteString(as.Name.String())
	out.WriteString(" = ")
	if as.Value != nil {
		out.WriteString(as.Value.String())
	}
	return out.String()
}
