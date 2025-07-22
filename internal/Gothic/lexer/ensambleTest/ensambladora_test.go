package lexer_test // Nota: Cambiado a package GothicLexer_test

import (
	"testing"

	"github.com/devicemxl/nexusl/internal/Gothic/lexer" // Ajusta la ruta de importación si es diferente
	tk "github.com/devicemxl/nexusl/internal/Gothic/token"
)

// Asegúrate de que esta ruta sea correcta

// Mock de la estructura Lexer y sus métodos auxiliares
// Esto es para que el test pueda instanciar un lexer de forma simplificada
// si no quieres importar directamente el paquete 'lexer' en el test package,
// o si necesitas simular su comportamiento.
// Si ya tienes tu struct Lexer en el paquete lexer, puedes omitir este mock
// y usar `lexer.NewLexer(input string)` si tienes esa función, o construirlo directamente.
type TestLexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
	line         int  // current line number
	column       int  // current column number
}

func (l *TestLexer) ReadChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // EOF
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++

	// Actualiza línea y columna
	if l.ch == '\n' {
		l.line++
		l.column = 1
	} else {
		l.column++
	}
}

func (l *TestLexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *TestLexer) NewToken(tokenType tk.TokenClass, literal string) tk.Token {
	return tk.Token{Type: tokenType, Word: literal, Line: l.line, Column: l.column}
}

func (l *TestLexer) SkipWhitespace() {
	// Esto es un mock simplificado. Tu lexer real debe tener la lógica de saltar
	// comentarios de una línea y espacios aquí si no se preprocesan.
	// Para el test, asumiremos que el input ya está preprocesado o que el lexer real lo maneja.
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.ReadChar()
	}
}

func (l *TestLexer) ReadIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) || l.ch == '_' { // Identificadores pueden contener _, letras y números
		l.ReadChar()
	}
	return l.input[position:l.position]
}

func (l *TestLexer) ReadNumber() string {
	position := l.position
	// Simplificado: lee enteros. Para floats se necesita más lógica.
	for isDigit(l.ch) {
		l.ReadChar()
	}
	return l.input[position:l.position]
}

func (l *TestLexer) ReadString() (string, error) {
	position := l.position + 1 // Start after the opening quote
	for {
		l.ReadChar()
		if l.ch == '"' { // Found closing quote
			break
		}
		if l.ch == 0 { // EOF before closing quote
			return "", nil // Or return an error: errors.New("Unterminated string")
		}
		// TODO: Handle escape characters like \" or \n
	}
	return l.input[position:l.position], nil
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// La función de prueba real para Ensambladora
func TestEnsambladora(t *testing.T) {
	tests := []struct {
		name           string // ¡Agrega este campo!
		input          string
		expectedTokens []tk.Token
	}{
		{
			input: "let x := 10 + y - 5; // Comentario de línea única (será preprocesado)\n rule myRule @(s p o) :: someFact { body }",
			expectedTokens: []tk.Token{
				{Type: tk.LET, Word: "let", Line: 1, Column: 1},
				{Type: tk.IDENTIFIER, Word: "x", Line: 1, Column: 5},
				{Type: tk.ASSIGN, Word: ":=", Line: 1, Column: 7},
				{Type: tk.INTEGER, Word: "10", Line: 1, Column: 10},
				{Type: tk.PLUS, Word: "+", Line: 1, Column: 13},
				{Type: tk.IDENTIFIER, Word: "y", Line: 1, Column: 15},
				{Type: tk.MINUS, Word: "-", Line: 1, Column: 17},
				{Type: tk.INTEGER, Word: "5", Line: 1, Column: 19},
				{Type: tk.SEMICOLON, Word: ";", Line: 1, Column: 20},
				{Type: tk.RULE, Word: "rule", Line: 2, Column: 2}, // Columna 2 si la línea anterior tuvo \n
				{Type: tk.IDENTIFIER, Word: "myRule", Line: 2, Column: 7},
				{Type: tk.LIST_BUILDER, Word: "@(", Line: 2, Column: 14}, // Inicio del List Builder
				{Type: tk.IDENTIFIER, Word: "s", Line: 2, Column: 16},
				{Type: tk.IDENTIFIER, Word: "p", Line: 2, Column: 18},
				{Type: tk.IDENTIFIER, Word: "o", Line: 2, Column: 20},
				{Type: tk.RPAREN, Word: ")", Line: 2, Column: 21}, // Cierre del List Builder
				{Type: tk.RESOLUTION, Word: "::", Line: 2, Column: 23},
				{Type: tk.IDENTIFIER, Word: "someFact", Line: 2, Column: 26},
				{Type: tk.LCURLY, Word: "{", Line: 2, Column: 35},
				{Type: tk.IDENTIFIER, Word: "body", Line: 2, Column: 37},
				{Type: tk.RCURLY, Word: "}", Line: 2, Column: 42},
				{Type: tk.EOF, Word: "", Line: 2, Column: 43},
			},
		},
		{
			input: `= == += -= *= /= %= ** -> :: != |> |* || < <= > >= << >> <[ ]> @[ @{ @<`,
			expectedTokens: []tk.Token{
				{Type: tk.ASSIGN_EQUAL, Word: "=", Line: 1, Column: 1},
				{Type: tk.EQUALITY, Word: "==", Line: 1, Column: 3},
				{Type: tk.ASSIGN_PLUS, Word: "+=", Line: 1, Column: 6},
				{Type: tk.ASSIGN_MINUS, Word: "-=", Line: 1, Column: 9},
				{Type: tk.ASSIGN_MULTIPLY, Word: "*=", Line: 1, Column: 12},
				{Type: tk.ASSIGN_DIVIDE, Word: "/=", Line: 1, Column: 15},
				{Type: tk.ASSIGN_MODULO, Word: "%=", Line: 1, Column: 18},
				{Type: tk.POWER, Word: "**", Line: 1, Column: 21},
				{Type: tk.ARROW, Word: "->", Line: 1, Column: 24},
				{Type: tk.RESOLUTION, Word: "::", Line: 1, Column: 27},
				{Type: tk.NOT_EQUALS, Word: "!=", Line: 1, Column: 30},
				{Type: tk.PIPE, Word: "|>", Line: 1, Column: 33},
				{Type: tk.MONADIC_PIPE, Word: "|*", Line: 1, Column: 36},
				{Type: tk.OR_GATE, Word: "||", Line: 1, Column: 39},
				{Type: tk.LESS, Word: "<", Line: 1, Column: 42},
				{Type: tk.LESS_EQUALS, Word: "<=", Line: 1, Column: 44},
				{Type: tk.GREATER, Word: ">", Line: 1, Column: 47},
				{Type: tk.GREATER_EQUALS, Word: ">=", Line: 1, Column: 49},
				{Type: tk.BIT_SHL, Word: "<<", Line: 1, Column: 52},
				{Type: tk.BIT_SHR, Word: ">>", Line: 1, Column: 55},
				{Type: tk.LVECT, Word: "<[", Line: 1, Column: 58},
				{Type: tk.RVECT, Word: "]>", Line: 1, Column: 61},
				{Type: tk.ARRAY_BUILDER, Word: "@[", Line: 1, Column: 64},
				{Type: tk.SET_BUILDER, Word: "@{", Line: 1, Column: 67},
				{Type: tk.VECTOR_BUILDER, Word: "@<", Line: 1, Column: 70},
				{Type: tk.EOF, Word: "", Line: 1, Column: 72},
			},
		},
		{
			input: `"hello world" 'c' . , ; : ! & ^ | ~ { } [ ]`,
			expectedTokens: []tk.Token{
				{Type: tk.STRING, Word: "hello world", Line: 1, Column: 1}, // Pasa correctamente
				{Type: tk.CHAR, Word: "c", Line: 1, Column: 15},            // Pasa correctamente ahora con tk.CHAR
				// --- ELIMINA ESTAS DOS LÍNEAS YA QUE NO ESTÁN EN EL INPUT LITERAL ---
				// {Type: tk.IDENTIFIER, Word: "c", Line: 1, Column: 17},
				// {Type: tk.SINGLE_QUOTE, Word: "'", Line: 1, Column: 18},
				// ------------------------------------------------------------------
				{Type: tk.DOT, Word: ".", Line: 1, Column: 19},       // Ajusta la columna si es necesario
				{Type: tk.COMMA, Word: ",", Line: 1, Column: 21},     // Ajusta la columna
				{Type: tk.SEMICOLON, Word: ";", Line: 1, Column: 23}, // Ajusta la columna
				{Type: tk.COLON, Word: ":", Line: 1, Column: 25},     // Ajusta la columna
				{Type: tk.NOT_GATE, Word: "!", Line: 1, Column: 27},  // Ajusta la columna
				{Type: tk.BIT_AND, Word: "&", Line: 1, Column: 29},   // Ajusta la columna
				{Type: tk.BIT_XOR, Word: "^", Line: 1, Column: 31},   // Ajusta la columna
				{Type: tk.BIT_OR, Word: "|", Line: 1, Column: 33},    // Ajusta la columna
				{Type: tk.BIT_NOT, Word: "~", Line: 1, Column: 35},   // Ajusta la columna
				{Type: tk.LCURLY, Word: "{", Line: 1, Column: 37},    // Ajusta la columna
				{Type: tk.RCURLY, Word: "}", Line: 1, Column: 39},    // Ajusta la columna
				{Type: tk.LBRACKET, Word: "[", Line: 1, Column: 41},  // Ajusta la columna
				{Type: tk.RBRACKET, Word: "]", Line: 1, Column: 43},  // Ajusta la columna
				{Type: tk.EOF, Word: "", Line: 1, Column: 44},        // Ajusta la columna
			},
		},
		{
			input: `/* This is a multi-line
                     comment that should be skipped */
                     func myFunc()`,
			expectedTokens: []tk.Token{
				{Type: tk.FUNC, Word: "func", Line: 3, Column: 22}, // Asume que la línea 1 y 2 son del comentario
				{Type: tk.IDENTIFIER, Word: "myFunc", Line: 3, Column: 27},
				{Type: tk.LPAREN, Word: "(", Line: 3, Column: 33},
				{Type: tk.RPAREN, Word: ")", Line: 3, Column: 34},
				{Type: tk.EOF, Word: "", Line: 3, Column: 35},
			},
		},
		{
			input: `// This is a single-line comment (will be removed by CleanSingleLineComments)
                     let value = 42 # Another single-line comment (removed)
                     `,
			expectedTokens: []tk.Token{
				{Type: tk.LET, Word: "let", Line: 2, Column: 22}, // Asume que la primera línea está vacía después del preproceso
				{Type: tk.IDENTIFIER, Word: "value", Line: 2, Column: 26},
				{Type: tk.ASSIGN_EQUAL, Word: "=", Line: 2, Column: 32},
				{Type: tk.INTEGER, Word: "42", Line: 2, Column: 34},
				{Type: tk.EOF, Word: "", Line: 3, Column: 22}, // EOF al final de la última línea limpia
			},
		},
		{
			input: `a b c`, // Prueba para identificadores simples
			expectedTokens: []tk.Token{
				{Type: tk.IDENTIFIER, Word: "a", Line: 1, Column: 1},
				{Type: tk.IDENTIFIER, Word: "b", Line: 1, Column: 3},
				{Type: tk.IDENTIFIER, Word: "c", Line: 1, Column: 5},
				{Type: tk.EOF, Word: "", Line: 1, Column: 6},
			},
		},
		{
			input: `123 456`, // Prueba para números simples
			expectedTokens: []tk.Token{
				{Type: tk.INTEGER, Word: "123", Line: 1, Column: 1},
				{Type: tk.INTEGER, Word: "456", Line: 1, Column: 5},
				{Type: tk.EOF, Word: "", Line: 1, Column: 8},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cleanedInput := lexer.StepOne(tt.input)
			t.Logf("DEBUG Test Case: %s", tt.name)             // Log del nombre del test
			t.Logf("DEBUG Cleaned Input:\n%q\n", cleanedInput) // Log del input limpio

			// Une el slice de líneas en un solo string si StepOne devuelve []string
			joinedInput := ""
			if lines, ok := interface{}(cleanedInput).([]string); ok {
				joinedInput = ""
				for i, line := range lines {
					if i > 0 {
						joinedInput += "\n"
					}
					joinedInput += line
				}
			} else if s, ok := interface{}(cleanedInput).(string); ok {
				joinedInput = s
			} else {
				t.Fatalf("StepOne returned unexpected type: %T", cleanedInput)
			}

			// Instancia tu lexer real
			l := &lexer.Lexer{Input: joinedInput, Position: 0, ReadPosition: 0, Ch: 0, Line: 1, Column: 0, LineStartPos: 0}

			t.Logf("DEBUG Before first ReadChar: Ch=%q (0x%x), Pos=%d, ReadPos=%d, Line=%d, Col=%d", l.Ch, l.Ch, l.Position, l.ReadPosition, l.Line, l.Column)
			l.ReadChar() // Llama a ReadChar() una vez al inicio para inicializar l.Ch
			t.Logf("DEBUG After first ReadChar: Ch=%q (0x%x), Pos=%d, ReadPos=%d, Line=%d, Col=%d", l.Ch, l.Ch, l.Position, l.ReadPosition, l.Line, l.Column)

			for i, expectedToken := range tt.expectedTokens {
				actualToken := l.Ensambladora()                                                // Llama a tu función NextToken
				t.Logf("DEBUG Token %d: Expected %+v, Got %+v", i, expectedToken, actualToken) // Log de cada token

				// Compara el tipo de token
				if actualToken.Type != expectedToken.Type {
					t.Errorf("Test '%s' token %d - tipo de token incorrecto. esperado=%q, obtenido=%q",
						tt.name, i, expectedToken.Type, actualToken.Type)
				}
				// Compara el literal del token (Word)
				if actualToken.Word != expectedToken.Word {
					t.Errorf("Test '%s' token %d - literal del token incorrecto. esperado=%q, obtenido=%q",
						tt.name, i, expectedToken.Word, actualToken.Word)
				}
				// Compara la línea y columna (¡muy importante!)
				if actualToken.Line != expectedToken.Line {
					t.Errorf("Test '%s' token %d - línea incorrecta. esperado=%d, obtenido=%d (token: %q)",
						tt.name, i, expectedToken.Line, actualToken.Line, actualToken.Word)
				}
				if actualToken.Column != expectedToken.Column {
					t.Errorf("Test '%s' token %d - columna incorrecta. esperado=%d, obtenido=%d (token: %q)",
						tt.name, i, expectedToken.Column, actualToken.Column, actualToken.Word)
				}
			}
		})
	}
}
