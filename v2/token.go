// token.go
package main

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

const (
	//
	// ======================================================== #
	// "System" Keywords
	// ======================================================== #
	// Parser, token, ast, lang...
	//
	// System Special Tokens
	//----------------------
	ILLEGAL      TokenType = "ILLEGAL" // Unrecognized or invalid token
	EOF          TokenType = "EOF"     // End of input
	STRING_QUOTE TokenType = `"`       // String literal opening quote
	ARROW        TokenType = "=>"      // Arrow operator (e.g., x -> y)
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
	IDENT  TokenType = "IDENT"  // Identifiers (e.g., myVar, calculaArea)
	STRING TokenType = "STRING" // String literals (e.g., "hello")
	INT    TokenType = "INT"    // Integer literals (e.g., 42)
	FLOAT  TokenType = "FLOAT"  // Float literals (e.g., 3.14)
	TRUE   TokenType = "TRUE"   // Boolean literal true
	FALSE  TokenType = "FALSE"  // Boolean literal false
	NIL    TokenType = "NIL"    // Null-like literal
	//
	// Delimiters and Grouping
	//----------------------
	SYMSIGN   TokenType = "$" // Sign to indicate a symbol
	COLON     TokenType = ":" // Colon (e.g., key:value) -- see TERNARY_COLON ambiguity need to be resolved
	SEMICOLON TokenType = ";" // Semicolon for statement separation
	COMMA     TokenType = "," // Comma for separating elements
	DOT       TokenType = "." // Dot for member access
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
	ASSIGN TokenType = "=" // Assignment (e.g., x = 10)
	// Arithmetic Operators
	PLUS     TokenType = "+" // Addition
	MINUS    TokenType = "-" // Subtraction or unary negation
	MULTIPLY TokenType = "*" // Multiplication
	DIVIDE   TokenType = "/" // Division
	MODULO   TokenType = "%" // Modulus (remainder)
	// Comparison Operators
	//----------------------
	EQUALS         TokenType = "==" // Equality
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

// ======================================================== #
// "Reserved Words" in nexusL
// ======================================================== #
//
// keywords maps reserved keywords to their TokenType
var keywords = map[string]TokenType{

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

// LookupIdent determines whether a given identifier is a keyword or a generic identifier.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
