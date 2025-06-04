// lexer.go
package main

import (
	"strings" // Necesario para strings.Builder
)

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
func NewLexer(input string) *Lexer { // Renombrado a New para consistencia con NewParser
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
func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace() // Ignore whitespace characters

	startLine := l.line
	startColumn := l.column

	switch l.ch {
	case '(':
		tok = newToken(LPAREN, string(l.ch), startLine, startColumn)
	case ')':
		tok = newToken(RPAREN, string(l.ch), startLine, startColumn)
	case '[':
		tok = newToken(LBRACKET, string(l.ch), startLine, startColumn)
	case ']':
		tok = newToken(RBRACKET, string(l.ch), startLine, startColumn)
	case '{':
		tok = newToken(LCURLY, string(l.ch), startLine, startColumn)
	case '}':
		tok = newToken(RCURLY, string(l.ch), startLine, startColumn)
	case ':':
		tok = newToken(COLON, string(l.ch), startLine, startColumn)
	case ';':
		tok = newToken(SEMICOLON, string(l.ch), startLine, startColumn)
	case ',':
		tok = newToken(COMMA, string(l.ch), startLine, startColumn)
	case '$': // Handle SYMSIGN explicitly
		tok = newToken(SYMSIGN, string(l.ch), startLine, startColumn)
	case '?': // Handle QUESTION explicitly
		tok = newToken(TERNARY_CONDITION, string(l.ch), startLine, startColumn)
	case '"': // Handle String Literals
		tok.Type = STRING
		tok.Line = startLine
		tok.Column = startColumn
		tok.Literal = l.readString() // readString will consume the closing quote
		return tok                   // Return early as readString already advanced position
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()                                                            // Consume first '='
			tok = newToken(EQUALS, string(ch)+string(l.ch), startLine, startColumn) // Consume second '='
		} else if l.peekChar() == '>' { // Handle ARROW "=>"
			ch := l.ch
			l.readChar()                                                           // Consume '='
			tok = newToken(ARROW, string(ch)+string(l.ch), startLine, startColumn) // Consume '>'
		} else {
			tok = newToken(ASSIGN, string(l.ch), startLine, startColumn)
		}
	case '+':
		tok = newToken(PLUS, string(l.ch), startLine, startColumn)
	case '-':
		tok = newToken(MINUS, string(l.ch), startLine, startColumn)
	case '*':
		tok = newToken(MULTIPLY, string(l.ch), startLine, startColumn)
	case '/':
		tok = newToken(DIVIDE, string(l.ch), startLine, startColumn)
	case '%':
		tok = newToken(MODULO, string(l.ch), startLine, startColumn)
	case '&':
		if l.peekChar() == '&' {
			ch := l.ch
			l.readChar()
			tok = newToken(AND, string(ch)+string(l.ch), startLine, startColumn)
		} else {
			tok = newToken(BIT_AND, string(l.ch), startLine, startColumn) // Bitwise AND if single '&'
		}
	case '|':
		if l.peekChar() == '|' {
			ch := l.ch
			l.readChar()
			tok = newToken(OR, string(ch)+string(l.ch), startLine, startColumn)
		} else {
			tok = newToken(BIT_OR, string(l.ch), startLine, startColumn) // Bitwise OR if single '|'
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = newToken(NOT_EQUALS, string(ch)+string(l.ch), startLine, startColumn)
		} else {
			tok = newToken(NOT, string(l.ch), startLine, startColumn)
		}
	case 0: // EOF
		tok = newToken(EOF, "", startLine, startColumn)
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = LookupIdent(tok.Literal) // Check if it's a keyword (FUNC, PARAM, HAS, IS, etc.) or IDENT
			return tok                          // Return early as readIdentifier already advanced position
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			if containsDot(tok.Literal) {
				tok.Type = FLOAT
			} else {
				tok.Type = INT
			}
			return tok // Return early as readNumber already advanced position
		} else {
			tok = newToken(ILLEGAL, string(l.ch), startLine, startColumn)
		}
	}

	l.readChar() // Advance for the next token (if not already advanced by specific read functions)
	return tok
}

// readIdentifier reads an identifier (letters, numbers, underscores)
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) || l.ch == '_' {
		l.readChar()
	}
	return l.input[position:l.position]
}

// skipWhitespace advances the lexer past whitespace characters.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// newToken is a helper function to create a new Token.
func newToken(tokenType TokenType, literal string, line int, column int) Token {
	return Token{Type: tokenType, Literal: literal, Line: line, Column: column}
}

// isLetter checks if the character is a letter (a-z, A-Z) or an underscore.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// isDigit checks if the character is a digit (0-9).
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

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

// readString reads characters until the closing quote is found, handling basic escape sequences.
func (l *Lexer) readString() string {
	// No longer need 'position' here as we're building the string with strings.Builder
	var out strings.Builder

	// Consume the opening quote (l.ch is currently '"')
	l.readChar()

	for {
		if l.ch == '"' || l.ch == 0 { // End of string or EOF
			break
		}
		if l.ch == '\\' { // Handle escape sequences
			l.readChar() // Consume the backslash
			switch l.ch {
			case 'n':
				out.WriteByte('\n')
			case 't':
				out.WriteByte('\t')
			case '"':
				out.WriteByte('"')
			case '\\':
				out.WriteByte('\\')
			default:
				out.WriteByte(l.ch) // For unknown escapes, just write the char
			}
		} else {
			out.WriteByte(l.ch)
		}
		l.readChar() // Consume the character (or escaped character)
	}
	l.readChar() // Consume the closing quote
	return out.String()
}

// containsDot is a helper to check if a string contains a dot for float detection
func containsDot(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			return true
		}
	}
	return false
}
