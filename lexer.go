// lexer.go
package main

// Lexer holds the state of the scanner.
type Lexer struct {
	input        string // The string to be tokenized
	position     int    // Current reading position in input (current char)
	readPosition int    // Next reading position in input (after current char)
	ch           byte   // Current character under examination
	line         int    // Current line number
	column       int    // Current column number
}

// New creates a new Lexer instance.
func NewLexer(input string) *Lexer {
	l := &Lexer{input: input, line: 1, column: 0}
	l.readChar() // Initialize position and ch
	return l
}

// readChar reads the next character in the input and advances the position.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII for "NUL" character, indicates EOF
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++

	// Update line and column numbers
	if l.ch == '\n' {
		l.line++
		l.column = 0
	} else {
		l.column++
	}
}

// peekChar returns the character at readPosition without advancing.
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

// NextToken reads characters until it identifies the next complete token.
// NextToken reads characters until it identifies the next complete token.
func (l *Lexer) NextToken() Token { // [cite: 29]
	var tok Token

	l.skipWhitespace() // Ignore whitespace characters [cite: 32]

	tok.Line = l.line
	tok.Column = l.column

	// Save original column for multicharacter tokens like numbers/strings
	startLine := l.line     // Make sure you define startLine here if it's not defined
	startColumn := l.column // Make sure you define startColumn here if it's not defined

	switch l.ch {
	// Add other single-character tokens here in the future, e.g.:
	// triplets & verbs
	case '(':
		tok = newToken(LPAREN, string(l.ch), startLine, startColumn) // CAMBIO AQUÍ
	case ')':
		tok = newToken(RPAREN, string(l.ch), startLine, startColumn) // CAMBIO AQUÍ
	case ':':
		tok = newToken(COLON, string(l.ch), startLine, startColumn) // CAMBIO AQUÍ
	case ';':
		tok = newToken(SEMICOLON, string(l.ch), startLine, startColumn) // CAMBIO AQUÍ
	case '$':
		tok = newToken(SYMSIGN, string(l.ch), startLine, startColumn) // CAMBIO AQUÍ
		// datastrctures separators
	case '{':
		tok = newToken(LCURLY, string(l.ch), startLine, startColumn) // CAMBIO AQUÍ
	case '}':
		tok = newToken(RCURLY, string(l.ch), startLine, startColumn) // CAMBIO AQUÍ
	case '[':
		tok = newToken(LBRACKET, string(l.ch), startLine, startColumn) // CAMBIO AQUÍ
	case ']':
		tok = newToken(RBRACKET, string(l.ch), startLine, startColumn) // CAMBIO AQUÍ
		// data type
	case '"':
		// No newToken(STRING_QUOTE, ...) here.
		// Instead, read the string content directly.
		tok.Type = STRING // The type of the token will be STRING
		tok.Line = startLine
		tok.Column = startColumn
		tok.Literal = l.readString() // This reads the actual content of the string
		return tok                   // readString already advanced past the closing quote
	case 0:
		tok = newToken(EOF, "", startLine, startColumn) // Literal para EOF es ""
	default:
		// Check for identifiers (which includes keywords) [cite: 30]
		if isLetter(l.ch) { // [cite: 35]
			tok.Literal = l.readIdentifier()    // [cite: 31]
			tok.Type = LookupIdent(tok.Literal) // Check if it's a keyword [cite: 23]
			return tok                          // Return early as readIdentifier already advanced position
		} else if isDigit(l.ch) { // NEW: Handle numbers
			tok.Literal = l.readNumber()
			// Determine if it's INT or FLOAT based on the literal
			if containsDot(tok.Literal) {
				tok.Type = FLOAT
			} else {
				tok.Type = INT
			}
			return tok
		} else {
			tok = newToken(ILLEGAL, string(l.ch), startLine, startColumn) // CAMBIO AQUÍ
		}
	}

	l.readChar() // Advance for the next token (if not already advanced by readIdentifier/readNumber/readString) [cite: 27]
	return tok
}

// readIdentifier reads an identifier (letters, numbers, underscores, hyphens) [cite: 30]
func (l *Lexer) readIdentifier() string {
	position := l.position                                              // [cite: 31]
	for isLetter(l.ch) || isDigit(l.ch) || l.ch == '-' || l.ch == '_' { // [cite: 31]
		l.readChar() // [cite: 31]
	}
	return l.input[position:l.position] // [cite: 31]
}

// skipWhitespace advances the lexer past whitespace characters. [cite: 32]
func (l *Lexer) skipWhitespace() { // [cite: 32]
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' { // [cite: 33]
		l.readChar() // [cite: 33]
	}
}

// newToken is a helper function to create a new Token. [cite: 34]
// IMPORTANT: This needs to accept a string for literal, not byte for multicharacter tokens
// newToken is a helper function to create a new Token.
// IMPORTANT: This needs to accept a string for literal, not byte for multicharacter tokens
func newToken(tokenType TokenType, literal string, line, column int) Token { // [cite: 34]
	return Token{Type: tokenType, Literal: literal, Line: line, Column: column}
}

// isLetter checks if the character is a letter (a-z, A-Z) or an underscore. [cite: 35]
func isLetter(ch byte) bool { // [cite: 35]
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' // [cite: 36]
}

// isDigit checks if the character is a digit (0-9). [cite: 37]
func isDigit(ch byte) bool { // [cite: 37]
	return '0' <= ch && ch <= '9' // [cite: 37]
}

// UNCOMMENT AND ACTIVATE THESE FUNCTIONS:
// readNumber reads a sequence of digits, potentially including a single decimal point.
func (l *Lexer) readNumber() string {
	position := l.position
	hasDecimal := false
	for isDigit(l.ch) || (l.ch == '.' && !hasDecimal) {
		if l.ch == '.' {
			hasDecimal = true
		}
		l.readChar()
	}
	return l.input[position:l.position]
}

// readString reads characters until the closing quote is found.
func (l *Lexer) readString() string {
	position := l.position + 1 // +1 to skip the opening quote
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 { // End of string or EOF
			break
		}
	}
	literal := l.input[position:l.position] // Get the content without quotes
	l.readChar()                            // Consume the closing quote
	return literal
}

// NEW: Helper to check if a string contains a dot for float detection
func containsDot(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			return true
		}
	}
	return false
}
