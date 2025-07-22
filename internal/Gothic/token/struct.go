// token/struct.go
package token

import "fmt"

// TokenClass representa el tipo de un token.
//
// Define un nuevo tipo llamado TokenClass que es esencialmente una cadena. Esto es
// ideal para representar los tipos de tokens como ILLEGAL, EOF, IDENTIFIER, KEYWORD,
// HOPE, RULE, etc. Proporciona una fuerte tipificación sin la necesidad de un enum
// explícito, y permite usar las constantes de cadena que ya has estado definiendo.
type TokenClass string

// Tipos de tokens
// Esta es la definición fundamental de lo que constituye un token una vez que el
// lexer lo ha identificado.
type Token struct {
	// Aquí es se usan las constantes TokenClass (como "ILLEGAL", "IDENTIFIER", "MAY",
	// etc.). Es el identificador principal de qué tipo de elemento sintáctico es.
	Type TokenClass
	// Este campo es crucial. Contendrá el "lexema" o la secuencia de caracteres real
	// que el lexer encontró en el código fuente que corresponde a este token. Por
	// ejemplo, si el Type es IDENTIFIER, Word podría ser "robot". Si el Type es
	// NUMBER, Word podría ser "0.80".
	Word string
	// Importantísimo para el manejo de errores. Cuando el parser encuentre un error,
	// se podra decirle al usuario exactamente en qué línea del código ocurrió.
	Line int
	// Complemento de Line. Indica la posición precisa (columna) dentro de esa línea
	// donde comienza el token. Esto hace que los mensajes de error sean muy específicos
	// y útiles para la depuración.
	Column int
}

// String devuelve una representación en cadena del Token.
// Útil para la depuración y mensajes de error.
func (t Token) String() string {
	return fmt.Sprintf("Type: %s, Word: '%s', Line: %d, Column: %d",
		t.Type, t.Word, t.Line, t.Column)
}

// LookupIdent determines whether a given identifier is a keyword or a generic identifier.
func LookupIdent(ident string) TokenClass {
	if tok, ok := Keywords[ident]; ok {
		return tok
	}
	return IDENTIFIER
}
