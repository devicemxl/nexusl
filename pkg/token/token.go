package token

// TokenClass representa el tipo de un token.
type TokenClass string

// Tipos de tokens
type Token struct {
	Type TokenClass // Tipo de token (IDENTIFIER, KEYWORD, etc.)
	Word string
}

const (
	// System Special Tokens
	//----------------------
	EMPTY        TokenClass = "EMPTY"        // TOKEN es un token genérico que puede ser usado para cualquier tipo de token.
	TOKEN        TokenClass = "TOKEN"        // TOKEN es un token genérico que puede ser usado para cualquier tipo de token.
	ILLEGAL      TokenClass = "ILLEGAL"      // ILLEGAL es un token no reconocido.
	EOF          TokenClass = "EOF"          // EOF es el fin de archivo.
	WHITESPACE   TokenClass = "WHITESPACE"   // WHITESPACE es un espacio en blanco (puede ser ignorado por el parser).
	STRING_QUOTE TokenClass = "STRING_QUOTE" // String literal opening quote
	KEYWORD      TokenClass = "KEYWORD"
	IDENTIFIER   TokenClass = "IDENTIFIER"
	STRING       TokenClass = "STRING"
	INTEGER      TokenClass = "INT"
	SEMICOLON    TokenClass = "SEMICOLON"
)

// classifyToken toma el literal de un token y lo clasifica.
// Esta función es el corazón de la clasificación léxica.
func ClassifyToken(word string) TokenClass {
	var tokenClass TokenClass

	switch word {
	/*
		case "=":
			TokenClass = ASSIGN
		case "+":
			TokenClass = PLUS
		case "-":
			TokenClass = MINUS
		case "*":
			TokenClass = ASTERISK
		case "/":
			TokenClass = SLASH
		case "<":
			TokenClass = LT
		case ">":
			TokenClass = GT
		case "(":
			TokenClass = LPAREN
		case ")":
			TokenClass = RPAREN
		case "{":
			TokenClass = LBRACE
		case "}":
			TokenClass = RBRACE
		case ",":
			TokenClass = COMMA
		case "==": // Manejar operadores de dos caracteres
			TokenClass = EQ
		case "!=":
			TokenClass = NOT_EQ
	*/
	case `"`: // Manejar comillas para cadenas
		tokenClass = STRING_QUOTE
	case ";":
		tokenClass = SEMICOLON
	case "":
		tokenClass = EMPTY
	default:
		tokenClass = KEYWORD
		return tokenClass
	}
	return tokenClass
}
