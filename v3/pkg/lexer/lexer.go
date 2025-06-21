// Package lexer proporciona un analizador léxico para el lenguaje NexusL.
package lexer

import (
	tk "github.com/devicemxl/nexusl/pkg/token" // Renombrado para evitar conflictos si hay un 'token' local
)

// Lexer es un analizador léxico que tokeniza la entrada de código.
type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// NewLexer devuelve un nuevo analizador léxico para la entrada dada.
func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar avanza la posición del lexer en la entrada.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// peekChar devuelve el próximo caracter sin avanzar la posición.
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

// NextToken devuelve el siguiente token del input.
func (l *Lexer) NextToken() tk.Token { // Usamos el alias tk.Token
	var tok tk.Token

	l.skipWhitespace()

	switch l.ch {
	//
	// Data Structures Containers
	//----------------------
	case '(':
		tok = newToken(tk.LPAREN, l.ch)
	case ')':
		tok = newToken(tk.RPAREN, l.ch)
	case '[':
		tok = newToken(tk.LBRACKET, l.ch)
	case ']':
		tok = newToken(tk.RBRACKET, l.ch)
	case '{':
		tok = newToken(tk.LCURLY, l.ch)
	case '}':
		tok = newToken(tk.RCURLY, l.ch)
	// Comparison Operators
	//----------------------
	case '=':
		if l.peekChar() == '>' {
			ch := l.ch
			l.readChar() // Consume '>'
			tok = tk.Token{Type: tk.ARROW, Literal: string(ch) + string(l.ch)}
		}
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar() // Consume '='
			tok = tk.Token{Type: tk.EQUALITY, Literal: string(ch) + string(l.ch)}
		} else {
			// Decide qué hacer con un '=' solitario.
			// Si no es un operador válido por sí mismo, podría ser ILLEGAL.
			// Por ahora, para evitar errores, puedes dejarlo como ILLEGAL si no es ARROW.
			//tok = newToken(tk.ILLEGAL, l.ch) // O definir un token ASSIGN si lo necesitas en el futuro
			tok = newToken(tk.EQUAL, l.ch)
		}
	//
	// Delimiters and Grouping
	//----------------------s
	case ';':
		tok = newToken(tk.SEMICOLON, l.ch)
	//
	// Literals
	//----------------------
	case '"':
		tok.Type = tk.STRING
		tok.Literal = l.readString()
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		tok.Type = tk.INTEGER // O tk.INTEGER si solo manejas enteros
		tok.Literal = l.readNumber()
	case 0:
		tok = newToken(tk.EOF, l.ch)
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			// AQUÍ ES DONDE LLAMAMOS A LookupIdent PARA CLASIFICAR EL IDENTIFICADOR
			/*
				Después de que readIdentifier() ha recolectado una secuencia de caracteres
				que parecen un identificador (ej., "def", "David", "is:symbol"), en lugar
				de simplemente asignar tk.IDENTIFIER, llama a tk.LookupIdent para que
				determine si es una palabra clave o un identificador regular.
			*/
			tok.Type = tk.LookupIdent(tok.Literal) // <-- CAMBIO CRUCIAL AQUÍ
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = tk.INTEGER
			return tok
		} else {
			tok = newToken(tk.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// newToken devuelve un nuevo token con el tipo y valor literal dados.
func newToken(tokenType tk.TokenType, ch byte) tk.Token {
	return tk.Token{Type: tokenType, Literal: string(ch)}
}

// readIdentifier lee un identificador completo (letras, números, guiones bajos, guiones).
func (l *Lexer) readIdentifier() string {
	position := l.position
	// isLetter ahora debe manejar guiones también para predicados como 'is:symbol' o 'has:rightLeg'
	for isLetter(l.ch) || isDigit(l.ch) || l.ch == '_' || l.ch == ':' { // Añadido ':' para is:symbol, has:rightLeg
		l.readChar()
	}
	return l.input[position:l.position]
}

// readString permanece igual
func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

/*
// skipWhitespace permanece igual
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
*/
// readNumber permanece igual
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// isDigit permanece igual
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// isLetter (asegúrate de que permite los caracteres para los predicados como 'is:symbol')
// Permitir letras, guiones bajos, guiones y DOS PUNTOS para identificadores/predicados
// como 'is:symbol', 'has:rightLeg', 'to:kitchen'
/*
	Hemos modificado isLetter para que también incluya el carácter : (dos puntos). Esto
	es crucial para que identificadores como is:symbol, has:rightLeg y to:kitchen sean
	leídos como un único token (IDENTIFIER, o una palabra clave si se decide que is: es
	una palabra clave en el futuro). Si no incluyes :, el lexer tokenizará is:symbol
	como [IDENTIFIER "is"], [ILLEGAL ":"], [IDENTIFIER "symbol"], lo cual no es lo que
	queremos.
*/
// isLetter (asegúrate de que permite los caracteres para los predicados como 'is:symbol')
// Permitir letras, guiones bajos, guiones y DOS PUNTOS para identificadores/predicados
// como 'is:symbol', 'has:rightLeg', 'to:kitchen'
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '-' || ch == ':'
}
