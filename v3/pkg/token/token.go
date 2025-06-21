package token

// TokenType es el tipo de un token léxico.
type TokenType string

const (
	//
	// ======================================================== #
	// "System" Keywords
	// ======================================================== #
	// Parser, token, ast, lang...
	//
	// System Special Tokens
	//----------------------
	ILLEGAL      TokenType = "ILLEGAL"    // ILLEGAL es un token no reconocido.
	EOF          TokenType = "EOF"        // EOF es el fin de archivo.
	WHITESPACE   TokenType = "WHITESPACE" // WHITESPACE es un espacio en blanco (puede ser ignorado por el parser).
	STRING_QUOTE TokenType = `"`          // String literal opening quote
	//
	// ======================================================== #
	// "Agentic entity" Keywords
	// ======================================================== #
	// seran la base de https://www.uv.mx/personal/cavalerio/2011/09/01/modulo-2-habilidades-basicas-de-pensamiento/
	//
	SYMBOL_KW TokenType = "SYMBOL" // Declaration of a symbolic entity (e.g., SYMBOL x)
	//
	// Core Verb Actions
	//----------------------
	// These keywords are used to define the core actions or operations
	// and to define the semantic roles of expressions in the language,
	// indicating their purpose or function.
	HAS TokenType = "HAS" // Semantic possession or attribute binding (e.g., x HAS value)
	DO  TokenType = "DO"  // Imperative or procedural action (e.g., DO { ... }) - Omitida en verbosDO se convirtio en palabra clave para facilitar el rec del programa
	IS  TokenType = "IS"  // Semantic identity or classification (e.g., x IS type)\
	// Core attribute conditions
	HOW   TokenType = "HOW"   // Semantic manner or method (e.g., x HOW action)
	WHEN  TokenType = "WHEN"  // Semantic temporal condition (e.g., x WHEN condition)
	WHERE TokenType = "WHERE" // Semantic spatial condition (e.g., x WHERE location)
	//
	// modal verbs
	//----------------------
	// (future after prolog implementation)
	//
	// Ability or permission
	// These keywords indicate the ability or permission to perform an action.
	// They are used to express capability or potential.
	COULD      TokenType = "COULD"      // Past
	CAN        TokenType = "CAN"        // present
	BE_ABLE_TO TokenType = "BE_ABLE_TO" // Future
	// Permission or possibility
	// These keywords indicate the permission or possibility of an action.
	MAYBE      TokenType = "MAYBE"      // Past
	PERMISSION TokenType = "PERMISSION" // Present
	ALLOWED_TO TokenType = "ALLOWED_TO" // Future
	// Possibility
	// These keywords indicate the possibility of an action.
	// They are used to express uncertainty or potential outcomes.
	MIGHT TokenType = "MIGHT" // Past
	MAY   TokenType = "MAY"   // Present - future
	// Necessity or obligation
	// These keywords indicate the necessity or obligation of an action.
	// They are used to express requirements or recommendations.
	HAD_TO       TokenType = "HAD_TO"       // Past
	MUST         TokenType = "MUST"         // Present
	WILL_HAVE_TO TokenType = "WILL_HAVE_TO" // Future
	// Suggestion or recommendation
	// These keywords indicate a suggestion or recommendation for an action.
	// They are used to express advice or guidance.
	SHOULD_HAVE TokenType = "SHOULD_HAVE" // Past
	SHOULD      TokenType = "SHOULD"      // Present - future
	// Requirement or necessity
	// These keywords indicate a requirement or necessity for an action.
	// They are used to express obligations or needs.
	NEED_TO   TokenType = "NEED_TO"   // Past
	NEED      TokenType = "NEED"      // Present - future
	WILL_NEED TokenType = "WILL_NEED" //
	//
	// ======================================================== #
	// "Language" tokens
	// ======================================================== #
	//
	// Literals
	//----------------------
	IDENTIFIER TokenType = "IDENTIFIER" // IDENTIFIER es un identificador (sujeto, verbo, variable, etc.).
	STRING     TokenType = "STRING"     // STRING es una cadena de texto entre comillas dobles.
	INTEGER    TokenType = "INTEGER"    // Para números como 30
	BOOLEAN    TokenType = "BOOLEAN"    // true, false
	NIL        TokenType = "NIL"        // Null-like literal
	// //=====================================
	// Palabras clave (Keywords)
	DEF    TokenType = "DEF"    // Nueva palabra clave
	ASSIGN TokenType = "ASSIGN" // Si decides tener un token explícito para ASSIGN
	FACT   TokenType = "FACT"   // Si decides tener un token explícito para FACT
	//
	// Delimiters and Grouping
	//----------------------
	SYMBOL_SIGN TokenType = "$" // Sign to indicate a symbol
	COLON       TokenType = ":" // Colon (e.g., key:value) -- see TERNARY_COLON ambiguity need to be resolved
	SEMICOLON   TokenType = ";" // Semicolon for statement separation
	COMMA       TokenType = "," // Comma for separating elements
	DOT         TokenType = "." // Dot for member access
	//
	// Data Structures Containers
	//----------------------
	LPAREN   TokenType = "(" // Left parenthesis
	RPAREN   TokenType = ")" // Right parenthesis
	LBRACKET TokenType = "[" // Left bracket
	RBRACKET TokenType = "]" // Right bracket
	LCURLY   TokenType = "{" // Left curly brace
	RCURLY   TokenType = "}" // Right curly brace
	// Function Definition Keywords
	FUNC   TokenType = "FUNC"   // Function declaration (e.g., FUNC myFunc param:[...] code:[...])
	PARAM  TokenType = "PARAM"  // Parameter block (e.g., PARAM: [x has:Value is:int;])
	CODE   TokenType = "CODE"   // Function code block (treated as pseudo-code)
	EXPORT TokenType = "EXPORT" // Export/return a value from within a function scope
	// Program Structure and Metadata Blocks
	PROGRAM  TokenType = "PROGRAM"  // Root-level declaration of a program
	HEADER   TokenType = "HEADER"   // Metadata or descriptive header for the program
	CONTENT  TokenType = "CONTENT"  // Main content or logic block
	OUTPUTS  TokenType = "OUTPUTS"  // Declaration of expected outputs
	METADATA TokenType = "METADATA" // Additional metadata, e.g., author, version
	INPUT    TokenType = "INPUT"    // Declaration of expected inputs
	IMPORT   TokenType = "IMPORT"   // External dependencies or modules
	// State & Data Definitions
	STATE TokenType = "STATE" // Declaration of internal mutable or reactive state
	DATA  TokenType = "DATA"  // Raw or structured data definitions
	// Control Flow Keywords
	ELSE     TokenType = "ELSE"     // Alternative branch in conditional execution
	IF       TokenType = "IF"       // Conditional execution (e.g., IF condition { ... })
	WHILE    TokenType = "WHILE"    // Looping construct (e.g., WHILE condition { ... })
	FOR      TokenType = "FOR"      // Iteration construct (e.g., FOR i IN range { ... })
	BREAK    TokenType = "BREAK"    // Exit from a loop or block
	CONTINUE TokenType = "CONTINUE" // Skip to the next iteration of a loop
	RETURN   TokenType = "RETURN"   // Return a value from a function or block
	SWITCH   TokenType = "SWITCH"   // Multi-way branch (e.g., SWITCH expression { ... })
	CASE     TokenType = "CASE"     // Case branch in a switch statement
	DEFAULT  TokenType = "DEFAULT"  // Default branch in a switch statement
	TRY      TokenType = "TRY"      // Exception handling block
	CATCH    TokenType = "CATCH"    // Exception handling block
	FINALLY  TokenType = "FINALLY"  // Final block in exception handling
	THROW    TokenType = "THROW"    // Raise an exception or error
	ASSERT   TokenType = "ASSERT"   // Assertion for debugging or validation
	// Type Annotations
	TYPE_ANNOTATION TokenType = "TYPE_ANNOTATION" // Type annotations (e.g., x: int)
	TYPE_CAST       TokenType = "TYPE_CAST"       // Type casting (e.g., x as int)
	TYPE_DEFINITION TokenType = "TYPE_DEFINITION" // Type definition (e.g., TYPE MyType { ... })
	//
	// ======================================================== #
	// "Mathematical" tokens
	// ======================================================== #
	//
	// Assignment Operator
	//----------------------
	// Arithmetic Operators
	EQUAL    TokenType = "=" // Assignment (e.g., x = 10)
	PLUS     TokenType = "+" // Addition
	MINUS    TokenType = "-" // Subtraction or unary negation
	MULTIPLY TokenType = "*" // Multiplication
	DIVIDE   TokenType = "/" // Division
	MODULO   TokenType = "%" // Modulus (remainder)
	// Comparison Operators
	//----------------------
	ARROW          TokenType = "=>" // Arrow operator (e.g., x -> y)
	EQUALITY       TokenType = "==" // Equality
	NOT_EQUALS     TokenType = "!=" // Inequality
	GREATER        TokenType = ">"  // Greater than
	LESS           TokenType = "<"  // Less than
	GREATER_EQUALS TokenType = ">=" // Greater than or equal to
	LESS_EQUALS    TokenType = "<=" // Less than or equal to
	// Logical Operators (boolean logic)
	//----------------------
	AND TokenType = "&&" // Logical AND
	OR  TokenType = "||" // Logical OR
	NOT TokenType = "!"  // Logical NOT (negation)
	// Bitwise Operators (operate on binary representations)
	//----------------------
	BIT_AND TokenType = "&"  // Bitwise AND
	BIT_OR  TokenType = "|"  // Bitwise OR
	BIT_XOR TokenType = "^"  // Bitwise XOR (exclusive OR)
	BIT_NOT TokenType = "~"  // Bitwise NOT (inversion)
	BIT_SHL TokenType = "<<" // Bitwise shift left
	BIT_SHR TokenType = ">>" // Bitwise shift right
	// Compound Assignment Operators
	//----------------------
	ASSIGN_PLUS     TokenType = "+=" // Addition assignment
	ASSIGN_MINUS    TokenType = "-=" // Subtraction assignment
	ASSIGN_MULTIPLY TokenType = "*=" // Multiplication assignment
	ASSIGN_DIVIDE   TokenType = "/=" // Division assignment
	ASSIGN_MODULO   TokenType = "%=" // Modulus assignment
	// Increment and Decrement Operators
	//----------------------
	INCREMENT TokenType = "++" // Increment (e.g., i++)
	DECREMENT TokenType = "--" // Decrement (e.g., i--)
	// Ternary Operator
	TERNARY_CONDITION TokenType = "?" // Ternary condition (e.g., cond ? a : b) -  inline if (abbreviated iif)
	// SEE COLORN - AMBIGUITY NEED TO BE RESOLVED In the future
	//TERNARY_COLON     TokenType = ":" // Ternary colon separator - if-else statement
	// Nullish Coalescing Operator
	NULLISH_COALESCE TokenType = "??" // Nullish coalescing (e.g., x ?? default)
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

// ======================================================== #
// "Reserved Words" in nexusL
// ======================================================== #
//
// keywords maps reserved keywords to their TokenType
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

	//
	// ======================================================== #
	// "Agentic entity" Keywords
	// ======================================================== #
	//
	"symbol": SYMBOL_KW, // Declaration of a symbolic entity (e.g., SYMBOL x)
	"has":    HAS,       // Semantic possession or attribute binding (e.g., x HAS value)
	"do":     DO,        // Imperative or procedural action (e.g., DO { ... })
	"is":     IS,        // Semantic identity or classification (e.g., x IS type)\
	"how":    HOW,       // Semantic manner or method (e.g., x HOW action)
	"when":   WHEN,      // Semantic temporal condition (e.g., x WHEN condition)
	"where":  WHERE,     // Semantic spatial condition (e.g., x WHERE location)
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
