package token

// TokenType es el tipo de un token léxico.
type TokenType string

const (
	// //=====================================
	// tokens "de sistema"
	ILLEGAL    TokenType = "ILLEGAL"    // ILLEGAL es un token no reconocido.
	EOF        TokenType = "EOF"        // EOF es el fin de archivo.
	WHITESPACE TokenType = "WHITESPACE" // WHITESPACE es un espacio en blanco (puede ser ignorado por el parser).
	// //=====================================
	// Identificadores y Literales
	IDENTIFIER TokenType = "IDENTIFIER" // IDENTIFIER es un identificador (sujeto, verbo, variable, etc.).
	STRING     TokenType = "STRING"     // STRING es una cadena de texto entre comillas dobles.
	NUMBER     TokenType = "NUMBER"     // Para números como 30
	INTEGER    TokenType = "INTEGER"    // 123
	BOOLEAN    TokenType = "BOOLEAN"    // true, false
	// //=====================================
	// Palabras clave (Keywords)
	DEF    TokenType = "DEF"    // Nueva palabra clave
	ASSIGN TokenType = "ASSIGN" // Si decides tener un token explícito para ASSIGN
	FACT   TokenType = "FACT"   // Si decides tener un token explícito para FACT
	// //=====================================
	// Operadores / Delimitadores (ya existentes o futuros)
	SEMICOLON TokenType = ";" // SEMICOLON es el delimitador de sentencia.
	LPAREN    TokenType = "("
	RPAREN    TokenType = ")"

	// ... aquí irían otros operadores o delimitadores como `?`, `=`, `(`, `)`, etc.
)

// Precedence constants for expression parsing
const (
	_ int = iota // Assigns 0 to _ and then increments
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	CALL        // myFunction(X)
)

// Token es un token léxico con su tipo y valor literal.
type Token struct {
	// Type es el tipo del token.
	Type TokenType
	// Literal es el valor literal del token.
	Literal string
}

// keywords mapea palabras clave reservadas
// //=====================================
/*
	var keywords = map[string]TokenType{...}: Declara un mapa global
	llamado keywords. Este mapa asocia las cadenas de texto de las
	palabras clave ("def", "fact") con sus respectivos TokenType
	(DEF, FACT).
*/
var keywords = map[string]TokenType{
	"def":    DEF,
	"assign": ASSIGN, // Si lo agregas
	"fact":   FACT,   // Si lo agregas
	"true":   BOOLEAN,
	"false":  BOOLEAN,
	// Puedes añadir más palabras clave aquí a medida que el lenguaje crezca,
	// por ejemplo: "true", "false", "null", "if", "else", "func", "return", etc.
}

// LookupIdent verifica si la cadena de entrada es una palabra clave reservada.
// Si es una palabra clave, devuelve su TokenType. De lo contrario, devuelve IDENTIFIER.
// ESTA FUNCIÓN ES CRUCIAL Y LA USARÁ EL LEXER
/*
	func LookupIdent(ident string) TokenType: Esta es la función central.
	Cuando el lexer ha leído un identificador (una secuencia de letras,
	números y guiones), lo pasa a LookupIdent. Esta función primero busca
	en el mapa keywords. Si encuentra una coincidencia (ej., si ident es
	"def"), devuelve token.DEF. Si no la encuentra, significa que es un
	identificador general (como "David" o "cli"), y devuelve token.IDENTIFIER.
*/
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENTIFIER
}

// New Function to map token types to their precedence
// This will be useful when you implement full expression parsing
// (e.g., arithmetic, boolean operations).
var precedences = map[TokenType]int{
	// Add actual operator tokens here as you define them (e.g., PLUS, MINUS)
	// tk.EQ: EQUALS,
	// tk.LT: LESSGREATER,
	// tk.GT: LESSGREATER,
	// tk.PLUS: SUM,
	// tk.MUL: PRODUCT,
}

func GetPrecedence(t TokenType) int {
	if p, ok := precedences[t]; ok {
		return p
	}
	return LOWEST // Default for non-operators or unknown operators
}
