// lexer/1ensambladora.go
package lexer

import (
	tk "github.com/devicemxl/nexusl/internal/Gothic/token"
)

// Ensambladora es el método principal que devuelve el próximo tk.Token
func (l *Lexer) Ensambladora() tk.Token {
	l.SkipWhitespace()

	// 1. CAPTURA LA POSICIÓN DE INICIO DEL TOKEN AQUÍ.
	// Esto es l.Position cuando el token comienza.
	startTokenPosition := l.Position

	var tok tk.Token

	switch l.Ch {
	// In "alphabetic order"
	case '-':
		if l.peekChar() == '=' { // ASSIGN_MINUS -=
			l.ReadChar() // Consume '-'
			l.ReadChar() // Consume '='
			tok = l.NewToken(tk.ASSIGN_MINUS, l.Input[startTokenPosition:l.Position], startTokenPosition)
		} else if l.peekChar() == '>' { // ARROW ->
			l.ReadChar() // Consume '-'
			l.ReadChar() // Consume '>'
			tok = l.NewToken(tk.ARROW, l.Input[startTokenPosition:l.Position], startTokenPosition)
		} else { // MINUS -
			tok = l.NewToken(tk.MINUS, string(l.Ch), startTokenPosition)
			l.ReadChar() // Consume el carácter actual
		}
		return tok // Agregar retorno explícito aquí y en todos los demás casos de uno o dos chars

	case ',':
		tok = l.NewToken(tk.COMMA, string(l.Ch), startTokenPosition)
		l.ReadChar() // Consume el carácter actual
		return tok

	case ';':
		tok = l.NewToken(tk.SEMICOLON, string(l.Ch), startTokenPosition)
		l.ReadChar() // Consume el carácter actual
		return tok

	case ':':
		if l.peekChar() == ':' { // RESOLUTION case ::
			l.ReadChar() // Consume ':'
			l.ReadChar() // Consume ':'
			tok = l.NewToken(tk.RESOLUTION, l.Input[startTokenPosition:l.Position], startTokenPosition)
		} else if l.peekChar() == '=' { // ASSIGN :=
			l.ReadChar() // Consume ':'
			l.ReadChar() // Consume '='
			tok = l.NewToken(tk.ASSIGN, l.Input[startTokenPosition:l.Position], startTokenPosition)
		} else { // COLON -
			tok = l.NewToken(tk.COLON, string(l.Ch), startTokenPosition)
			l.ReadChar() // Consume el carácter actual
		}
		return tok

	case '!':
		if l.peekChar() == '=' { // NOT_EQUALS case !=
			l.ReadChar() // Consume '!'
			l.ReadChar() // Consume '='
			tok = l.NewToken(tk.NOT_EQUALS, l.Input[startTokenPosition:l.Position], startTokenPosition)
		} else { // NOT_GATE !
			tok = l.NewToken(tk.NOT_GATE, string(l.Ch), startTokenPosition)
			l.ReadChar() // Consume el carácter actual
		}
		return tok

	case '.':
		tok = l.NewToken(tk.DOT, string(l.Ch), startTokenPosition)
		l.ReadChar() // Consume el carácter actual
		return tok

	case '\'':
		l.ReadChar()
		literal := l.readCharLiteral()
		tok = l.NewToken(tk.CHAR, literal, startTokenPosition)
		return tok // readCharLiteral ya avanzó el puntero

	case '"':
		l.ReadChar()
		literal := l.readStringLiteral()
		tok = l.NewToken(tk.STRING, literal, startTokenPosition)
		return tok // readStringLiteral ya avanzó el puntero

	case '@': // Manejo de Builders
		if l.peekChar() == '(' { // Es el inicio de un List Builder "@("
			l.ReadChar() // Consume '@'
			l.ReadChar() // Consume '('
			tok = l.NewToken(tk.LIST_BUILDER, l.Input[startTokenPosition:l.Position], startTokenPosition)
		} else if l.peekChar() == '[' { // Es el inicio de un Array Builder "@["
			l.ReadChar() // Consume '@'
			l.ReadChar() // Consume '['
			tok = l.NewToken(tk.ARRAY_BUILDER, l.Input[startTokenPosition:l.Position], startTokenPosition)
		} else if l.peekChar() == '{' { // Es el inicio de un SET Builder "@{"
			l.ReadChar() // Consume '@'
			l.ReadChar() // Consume '{'
			tok = l.NewToken(tk.SET_BUILDER, l.Input[startTokenPosition:l.Position], startTokenPosition)
		} else if l.peekChar() == '<' { // Es el inicio de un VECTOR Builder "@<"
			l.ReadChar() // Consume '@'
			l.ReadChar() // Consume '<'
			tok = l.NewToken(tk.VECTOR_BUILDER, l.Input[startTokenPosition:l.Position], startTokenPosition)
		} else { // No es un builder, es solo el símbolo '@' (AT_EACH)
			tok = l.NewToken(tk.AT_EACH, string(l.Ch), startTokenPosition)
			l.ReadChar() // Consume el carácter actual
		}
		return tok

	case '(':
		tok = l.NewToken(tk.LPAREN, string(l.Ch), startTokenPosition)
		l.ReadChar() // Consume el carácter actual
		return tok
	case ')':
		tok = l.NewToken(tk.RPAREN, string(l.Ch), startTokenPosition)
		l.ReadChar() // Consume el carácter actual
		return tok
	case '{':
		tok = l.NewToken(tk.LCURLY, string(l.Ch), startTokenPosition)
		l.ReadChar() // Consume el carácter actual
		return tok
	case '}':
		tok = l.NewToken(tk.RCURLY, string(l.Ch), startTokenPosition)
		l.ReadChar() // Consume el carácter actual
		return tok
	case '[':
		tok = l.NewToken(tk.LBRACKET, string(l.Ch), startTokenPosition)
		l.ReadChar() // Consume el carácter actual
		return tok

	case ']':
		if l.peekChar() == '>' { // RVECT ]>
			l.ReadChar() // Consume ']'
			l.ReadChar() // Consume '>'
			tok = l.NewToken(tk.RVECT, l.Input[startTokenPosition:l.Position], startTokenPosition)
		} else { // RBRACKET ]
			tok = l.NewToken(tk.RBRACKET, string(l.Ch), startTokenPosition)
			l.ReadChar() // Consume el carácter actual
		}
		return tok

	case '*':
		if l.peekChar() == '*' { // POWER **
			l.ReadChar() // Consume '*'
			l.ReadChar() // Consume '*'
			tok = l.NewToken(tk.POWER, l.Input[startTokenPosition:l.Position], startTokenPosition)
		} else if l.peekChar() == '=' { // ASSIGN_MULTIPLY  *=
			l.ReadChar() // Consume '*'
			l.ReadChar() // Consume '='
			tok = l.NewToken(tk.ASSIGN_MULTIPLY, l.Input[startTokenPosition:l.Position], startTokenPosition)
		} else { // MULTIPLY *
			tok = l.NewToken(tk.MULTIPLY, string(l.Ch), startTokenPosition)
			l.ReadChar() // Consume el carácter actual
		}
		return tok

	case '/':
		if l.peekChar() == '=' { // ASSIGN_DIVIDE /=
			l.ReadChar() // Consume '/'
			l.ReadChar() // Consume '='
			tok = l.NewToken(tk.ASSIGN_DIVIDE, l.Input[startTokenPosition:l.Position], startTokenPosition)
		} else { // División normal /
			tok = l.NewToken(tk.DIVIDE, string(l.Ch), startTokenPosition)
			l.ReadChar() // Consume el carácter actual
		}
		return tok

	case '&':
		tok = l.NewToken(tk.BIT_AND, string(l.Ch), startTokenPosition)
		l.ReadChar() // Consume el carácter actual
		return tok

	case '#':
		tok = l.NewToken(tk.ILLEGAL, string(l.Ch), startTokenPosition) // Asumo que es ilegal si no es comentario
		l.ReadChar()                                                   // Consume el carácter actual
		return tok

	case '%':
		if l.peekChar() == '=' { // ASSIGN_MODULO %=
			l.ReadChar() // Consume '%'
			l.ReadChar() // Consume '='
			tok = l.NewToken(tk.ASSIGN_MODULO, l.Input[startTokenPosition:l.Position], startTokenPosition)
		} else { // MODULO %
			tok = l.NewToken(tk.MODULO, string(l.Ch), startTokenPosition)
			l.ReadChar() // Consume el carácter actual
		}
		return tok

	case '`':
		tok = l.NewToken(tk.BACKTICK, string(l.Ch), startTokenPosition)
		l.ReadChar() // Consume el carácter actual
		return tok

	case '^':
		tok = l.NewToken(tk.BIT_XOR, string(l.Ch), startTokenPosition)
		l.ReadChar() // Consume el carácter actual
		return tok

	case '+':
		if l.peekChar() == '=' { // ASSIGN_PLUS +=
			l.ReadChar() // Consume '+'
			l.ReadChar() // Consume '='
			tok = l.NewToken(tk.ASSIGN_PLUS, l.Input[startTokenPosition:l.Position], startTokenPosition)
		} else { // PLUS +
			tok = l.NewToken(tk.PLUS, string(l.Ch), startTokenPosition)
			l.ReadChar() // Consume el carácter actual
		}
		return tok

	case '<':
		if l.peekChar() == '[' { // LVECT  <[
			l.ReadChar() // Consume '<'
			l.ReadChar() // Consume '['
			tok = l.NewToken(tk.LVECT, l.Input[startTokenPosition:l.Position], startTokenPosition)
		} else if l.peekChar() == '<' { // BIT_SHL  <<
			l.ReadChar() // Consume '<'
			l.ReadChar() // Consume '<'
			tok = l.NewToken(tk.BIT_SHL, l.Input[startTokenPosition:l.Position], startTokenPosition)
		} else if l.peekChar() == '=' { // LESS_EQUALS <=
			l.ReadChar() // Consume '<'
			l.ReadChar() // Consume '='
			tok = l.NewToken(tk.LESS_EQUALS, l.Input[startTokenPosition:l.Position], startTokenPosition)
		} else { // LESS <
			tok = l.NewToken(tk.LESS, string(l.Ch), startTokenPosition)
			l.ReadChar() // Consume el carácter actual
		}
		return tok

	case '=':
		if l.peekChar() == '=' { // EQUALITY    ==
			l.ReadChar() // Consume '='
			l.ReadChar() // Consume '='
			tok = l.NewToken(tk.EQUALITY, l.Input[startTokenPosition:l.Position], startTokenPosition)
		} else { // ASSIGN_EQUAL =
			tok = l.NewToken(tk.ASSIGN_EQUAL, string(l.Ch), startTokenPosition)
			l.ReadChar() // Consume el carácter actual
		}
		return tok

	case '>':
		if l.peekChar() == '=' { // GREATER_EQUALS  >=
			l.ReadChar() // Consume '>'
			l.ReadChar() // Consume '='
			tok = l.NewToken(tk.GREATER_EQUALS, l.Input[startTokenPosition:l.Position], startTokenPosition)
		} else if l.peekChar() == '>' { // BIT_SHR  >>
			l.ReadChar() // Consume '>'
			l.ReadChar() // Consume '>'
			tok = l.NewToken(tk.BIT_SHR, l.Input[startTokenPosition:l.Position], startTokenPosition)
		} else { // GREATER >
			tok = l.NewToken(tk.GREATER, string(l.Ch), startTokenPosition)
			l.ReadChar() // Consume el carácter actual
		}
		return tok

	case '|':
		if l.peekChar() == '>' { // Pipe case |>
			l.ReadChar() // Consume '|'
			l.ReadChar() // Consume '>'
			tok = l.NewToken(tk.PIPE, l.Input[startTokenPosition:l.Position], startTokenPosition)
		} else if l.peekChar() == '*' { // Nomadic |*
			l.ReadChar() // Consume '|'
			l.ReadChar() // Consume '*'
			tok = l.NewToken(tk.MONADIC_PIPE, l.Input[startTokenPosition:l.Position], startTokenPosition)
		} else if l.peekChar() == '|' { // OR Gate boolean ||
			l.ReadChar() // Consume '|'
			l.ReadChar() // Consume '|'
			tok = l.NewToken(tk.OR_GATE, l.Input[startTokenPosition:l.Position], startTokenPosition)
		} else { // Bitwise OR  |
			tok = l.NewToken(tk.BIT_OR, string(l.Ch), startTokenPosition)
			l.ReadChar() // Consume el carácter actual
		}
		return tok

	case '~':
		tok = l.NewToken(tk.BIT_NOT, string(l.Ch), startTokenPosition)
		l.ReadChar() // Consume el carácter actual
		return tok

	case 0: // EOF
		tok = l.NewToken(tk.EOF, "", startTokenPosition) // Literal vacío para EOF
		return tok                                       // Retorna directamente para EOF

	default: // Identificadores, números
		if IsLetter(l.Ch) {
			tok.Line = l.Line
			tok.Column = l.Column
			literal := l.ReadIdentifier() // ReadIdentifier ya avanza l.Ch
			// --- ¡EL CAMBIO CLAVE ESTÁ AQUÍ! ---
			// Usar LookupIdent para determinar si es una palabra clave o un IDENTIFIER
			tok.Type = LookupIdent(tok.Word)
			//
			tok = l.NewToken(tk.LookupIdent(literal), literal, startTokenPosition)
			return tok // ReadIdentifier ya avanzó el puntero
		} else if isDigit(l.Ch) {
			literal, tokenType := l.ReadNumber() // ReadNumber ya avanza l.Ch
			tok = l.NewToken(tokenType, literal, startTokenPosition)
			return tok // ReadNumber ya avanzó el puntero
		} else {
			tok = l.NewToken(tk.ILLEGAL, string(l.Ch), startTokenPosition)
			l.ReadChar() // Consume el carácter ilegal
			return tok
		}
	}
}
