// statements.go
package ast

import (
	"bytes"

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

func (fts *FlatTripletaStatement) String() string {
	var out bytes.Buffer

	// Imprime la palabra clave DEF o FACT si está presente
	if fts.Token.Type == token.DEF || fts.Token.Type == token.FACT {
		out.WriteString(fts.Token.Literal)
		out.WriteString(" ")
	}
	out.WriteString(fts.Subject.String())
	out.WriteString(" ")
	out.WriteString(fts.Verb.String())
	out.WriteString(" ")
	if fts.Object != nil {
		out.WriteString(fts.Object.String())
	}
	out.WriteString(";") // Cierre de sentencia

	return out.String()
}
