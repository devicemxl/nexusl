// lexer/theLexer.go
package lexer

import (
	tk "github.com/devicemxl/nexusl/Gothic/token"
)

// Lexer representa la instancia del analizador léxico
type Lexer struct {
	Input        string
	Position     int  // current position in input (points to current char)
	ReadPosition int  // current reading position in input (after current char)
	Ch           byte // current char under examination
	Line         int  // current line number (1-based)
	Column       int  // current column number (1-based)
	LineStartPos int  // Position of the start of the current line in Input string (0-indexed)
}

// SkipWhitespace avanza el lexer sobre los caracteres de espacio en blanco.
// Es importante que *actualice* l.Line y l.Column correctamente al pasar por '\n'.
func (l *Lexer) SkipWhitespace() {
	for l.Ch == ' ' || l.Ch == '\t' || l.Ch == '\n' || l.Ch == '\r' {
		l.ReadChar()
	}
}

// LookupIdent verifica si la palabra es una palabra clave o un identificador.
func LookupIdent(ident string) tk.TokenClass {
	if tok, ok := tk.Keywords[ident]; ok {
		return tok
	}
	return tk.IDENTIFIER // Cambié de TK.TOKEN a TK.IDENTIFIER si no es una palabra clave.
}

// New crea e inicializa un nuevo lexer
func New(input string) *Lexer {
	l := &Lexer{
		Input:        input,
		Position:     0,
		ReadPosition: 0,
		Ch:           0,
		Line:         1,
		Column:       0, // Inicializa en 0 para que la primera ReadChar lo ponga en 1 para el primer carácter
		LineStartPos: 0,
	}
	l.ReadChar() // Carga el primer carácter y actualiza Line/Column
	return l
}

// ReadChar es el método que avanza el lexer un carácter
func (l *Lexer) ReadChar() {
	if l.ReadPosition >= len(l.Input) {
		l.Ch = 0 // EOF
	} else {
		l.Ch = l.Input[l.ReadPosition]
	}

	// Actualiza Position para que apunte al carácter que acabamos de cargar en l.Ch
	l.Position = l.ReadPosition
	l.ReadPosition++

	// Lógica de actualización de línea y columna
	if l.Ch == '\n' {
		l.Line++
		l.Column = 0                    // Reinicia la columna para la nueva línea (se hará 1 en la próxima lectura de carácter)
		l.LineStartPos = l.ReadPosition // La nueva línea comienza justo DESPUÉS del '\n'
	} else {
		// Solo incrementa la columna si no es un salto de línea.
		l.Column++ // l.Column siempre debería reflejar la columna del carácter actual en l.Ch
	}
}

// NewToken crea un nuevo token con el tipo, literal y la posición de inicio del token.
// La columna se calcula basándose en tokenStartPos (la Position del lexer cuando el token comenzó a ser leído).
func (l *Lexer) NewToken(tokenType tk.TokenClass, literal string, tokenStartPos int) tk.Token {
	// La columna del token se calcula usando la posición de inicio del token (tokenStartPos)
	// y la posición de inicio de la línea (l.LineStartPos).
	// Se suma 1 porque las columnas suelen ser 1-based.
	columnForToken := (tokenStartPos - l.LineStartPos) + 1
	return tk.Token{Type: tokenType, Word: literal, Line: l.Line, Column: columnForToken}
}

// peekChar devuelve el próximo carácter sin avanzar la posición
func (l *Lexer) peekChar() byte {
	if l.ReadPosition >= len(l.Input) {
		return 0
	}
	return l.Input[l.ReadPosition]
}

// IsLetter verifica si un byte es una letra (para identificadores).
func IsLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// isDigit verifica si un byte es un dígito.
// isDigit verifica si un byte es un dígito.
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// ReadNumber lee un número completo (entero, flotante o complejo).
// Determina el tipo de número y lo devuelve como string.
func (l *Lexer) ReadNumber() (string, tk.TokenClass) {
	position := l.Position
	tokenType := tk.INTEGER // Asumimos entero por defecto

	// Leer parte entera
	for isDigit(l.Ch) {
		l.ReadChar()
	}

	// Comprobar si es flotante (contiene un punto decimal)
	if l.Ch == '.' && isDigit(l.peekChar()) {
		tokenType = tk.FLOAT // Es un flotante
		l.ReadChar()         // Consume el '.'
		for isDigit(l.Ch) {
			l.ReadChar() // Lee los dígitos después del punto
		}
	}

	// Comprobar si es un número complejo (contiene 'i' al final, asumiendo formato Go/Python)
	// Opcional: Podrías querer soportar 2+3i, 3i, etc. Por ahora, solo detecta 'i' al final de un número.
	// Para 2+3i, necesitarías una fase de parsing que combine INTEGER/FLOAT + PLUS + INTEGER/FLOAT + I
	if l.Ch == 'i' { // Asumiendo 'i' como sufijo para números imaginarios puros como '3i' o '3.14i'
		// Esto clasifica números como '10i' o '3.14i' como COMPLEX.
		// Si tu gramática soporta '2 + 3i' como un solo token complejo, esto sería más complejo.
		tokenType = tk.COMPLEX // Es un número complejo (imaginario puro en esta implementación)
		l.ReadChar()           // Consume la 'i'
	} else if l.Ch == 'E' || l.Ch == 'e' { // Notación científica (ej. 1e5, 1.2e-3)
		// Solo si no es ya un número complejo puro. Un número complejo como 3e+2i es más avanzado.
		// Aquí detectamos 'e' o 'E' para números flotantes en notación científica.
		// Asegúrate de que no es la 'e' de un identificador.
		l.ReadChar() // Consume 'E' o 'e'
		if l.Ch == '+' || l.Ch == '-' {
			l.ReadChar() // Consume '+' o '-'
		}
		for isDigit(l.Ch) {
			l.ReadChar()
		}
		tokenType = tk.FLOAT // Un número con notación científica es un flotante
	}

	return l.Input[position:l.Position], tokenType
}

// ReadIdentifier lee un identificador completo.
func (l *Lexer) ReadIdentifier() string {
	position := l.Position
	for IsLetter(l.Ch) || isDigit(l.Ch) {
		l.ReadChar()
	}
	return l.Input[position:l.Position]
}

// readStringLiteral lee un literal de cadena (entre comillas dobles).
// Se asume que l.Ch ya está en el carácter *después* de la comilla de apertura.
// Esta función lee hasta la comilla de cierre y avanza el lexer hasta después de ella.
func (l *Lexer) readStringLiteral() string {
	start := l.Position // La posición actual es donde empieza el contenido del string

	for {
		l.ReadChar() // Avanza al siguiente carácter

		// Si es la comilla de cierre o EOF
		if l.Ch == '"' || l.Ch == 0 {
			break
		}
		// TODO: Implementar lógica para caracteres de escape si los necesitas
		// Ejemplo simple para '\\' y '\"':
		// if l.Ch == '\\' {
		// 	l.ReadChar() // Consume el backslash
		// 	// Podrías añadir un switch aquí para diferentes caracteres de escape
		// }
	}

	// El literal es desde 'start' hasta l.Position (que es la comilla de cierre)
	literal := l.Input[start:l.Position]

	// Si no es EOF, consume la comilla de cierre
	if l.Ch == '"' {
		l.ReadChar() // Consume la comilla de cierre
	}

	return literal
}

// readCharLiteral lee un literal de caracter (entre comillas simples).
// Se asume que l.Ch ya está en el carácter *después* de la comilla de apertura.
// Esta función lee hasta la comilla de cierre y avanza el lexer hasta después de ella.
func (l *Lexer) readCharLiteral() string {
	start := l.Position // La posición actual es donde empieza el contenido del caracter

	for {
		l.ReadChar() // Avanza al siguiente carácter

		// Si es la comilla simple de cierre o EOF
		if l.Ch == '\'' || l.Ch == 0 {
			break
		}
		// TODO: Implementar lógica para caracteres de escape si los necesitas
	}

	// El literal es desde 'start' hasta l.Position (que es la comilla de cierre)
	literal := l.Input[start:l.Position]

	// Si no es EOF, consume la comilla de cierre
	if l.Ch == '\'' {
		l.ReadChar() // Consume la comilla de cierre
	}

	return literal
}
