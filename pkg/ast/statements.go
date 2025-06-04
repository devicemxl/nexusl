// pkg/ast/statements.go
package ast

import (
	"bytes"
	// "fmt" // No es necesario si usas bytes.Buffer para la cadena
	"github.com/devicemxl/nexusl/pkg/token"
)

// FlatTripletaStatement representa una tripleta plana como "sujeto verbo objeto;"
// O una sentencia "def sujeto predicado objeto;" o "fact sujeto predicado objeto;"
type FlatTripletaStatement struct {
	Token   token.Token // El primer token de la sentencia (DEF, FACT, IDENTIFIER)
	Subject Expression  // Normalmente *Identifier (o una expresión más compleja si la permites)
	Verb    *Identifier // El "predicado" o nombre del método/atributo
	Object  Expression  // El "complemento" o argumento/valor
}

func (fts *FlatTripletaStatement) statementNode()       {}
func (fts *FlatTripletaStatement) TokenLiteral() string { return fts.Token.Literal }

/*
Ejemplos de salida con la String() recomendada:

Input: def David is:symbol;

fts.Token.Type = token.DEF, fts.Token.Literal = "def"
fts.Subject.String() = "David"
fts.Verb.String() = "is:symbol"
fts.Object.String() = "" (asumiendo objeto nulo para is:symbol)
Output: def David is:symbol ;
Input: fact David has:rightLeg;

fts.Token.Type = token.FACT, fts.Token.Literal = "fact"
fts.Subject.String() = "David"
fts.Verb.String() = "has:rightLeg"
fts.Object.String() = ""
Output: fact David has:rightLeg ;
Input: cli print "hello";

fts.Token.Type = token.IDENTIFIER, fts.Token.Literal = "cli"
fts.Subject.String() = "cli"
fts.Verb.String() = "print"
fts.Object.String() = ""hello""
Output: cli print "hello"; (¡ya no hay cli cli!)

Esta versión de String() es más robusta y refleja mejor la intención para
diferentes tipos de sentencias.
*/
func (fts *FlatTripletaStatement) String() string {
	var out bytes.Buffer

	// Si la sentencia es "def" o "fact", imprime la palabra clave inicial.
	// Si es una invocación directa (ej. "cli print"), fts.Token.Literal ya es el sujeto.
	// En este caso, fts.Subject.String() y fts.Token.Literal son el mismo,
	// por lo que solo imprimimos el sujeto una vez.
	if fts.Token.Type == token.DEF || fts.Token.Type == token.FACT {
		out.WriteString(fts.Token.Literal) // Ej: "def", "fact"
		out.WriteString(" ")
		out.WriteString(fts.Subject.String()) // Ej: "David"
		out.WriteString(" ")
	} else { // Implícitamente, fts.Token.Type == token.IDENTIFIER (para invocaciones)
		out.WriteString(fts.Subject.String()) // Ej: "cli" o "David"
		out.WriteString(" ")
	}

	out.WriteString(fts.Verb.String()) // Ej: "is:symbol", "print", "has:rightLeg", "walk"
	out.WriteString(" ")

	if fts.Object != nil {
		out.WriteString(fts.Object.String()) // Ej: "true", "\"hello\"", "" (si es nulo)
	}

	out.WriteString(";") // El terminador
	return out.String()
}
