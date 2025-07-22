package token_test // Nota: Cambiado a package token_test para testing externo

import (
	"testing"

	. "github.com/devicemxl/nexusl/internal/Gothic/token" // Importa tu paquete token
)

func TestLookupIdent(t *testing.T) {
	tests := []struct {
		input    string
		expected TokenClass
	}{
		// Palabras clave
		{"if", IF},
		{"else", ELSE},
		{"and", AND_GATE}, // Asumiendo que "and" es la palabra clave para AND_GATE
		{"or", OR_GATE},
		{"not", NOT_GATE},
		{"mil", NIL},     // Aquí "mil" es el literal para NIL
		{"maybe", MAYBE}, // Aquí "maybe" es el literal para MAYBE
		{"func", FUNC},
		{"let", LET},
		{"return", RETURN},
		{"try", TRY},
		{"catch", CATCH},
		{"finally", FINALLY},
		{"throw", THROW},
		{"read_only", READ_ONLY}, // Si "read_only" es la forma esperada
		{"require", REQUIRE},     // Si "requires" es la forma esperada
		{"enable", ENABLE},       // Si "enables" es la forma esperada

		// Identificadores (no palabras clave)
		{"robot", IDENTIFIER},
		{"myVariable", IDENTIFIER},
		{"_agentName", IDENTIFIER},
		{"x", IDENTIFIER},
		{"_123", IDENTIFIER}, // Los identificadores pueden empezar con _

		// Símbolos que actúan como palabras clave (si están en Keywords)
		/*
			{"->", ARROW},
			{"|>", PIPE},
			{"|*", MONADIC_PIPE},
			{"++", INCREMENT},
		*/
	}

	for _, tt := range tests {
		actual := LookupIdent(tt.input)
		if actual != tt.expected {
			t.Errorf("LookupIdent(%q) ERROR: esperado %q, obtenido %q",
				tt.input, tt.expected, actual)
		}
	}
}

// Puedes añadir más tests aquí si es necesario para otros aspectos de tus tokens,
// como la creación de tokens, si creas una función NewToken.
/*
func TestNewToken(t *testing.T) {
	tok := NewToken(IDENTIFIER, "testVar", 1, 10)
	if tok.Type != IDENTIFIER {
		t.Errorf("NewToken type wrong. expected=%q, got=%q", IDENTIFIER, tok.Type)
	}
	if tok.Word != "testVar" {
		t.Errorf("NewToken literal wrong. expected=%q, got=%q", "testVar", tok.Word)
	}
	if tok.Line != 1 {
		t.Errorf("NewToken line wrong. expected=%d, got=%d", 1, tok.Line)
	}
	if tok.Column != 10 {
		t.Errorf("NewToken column wrong. expected=%d, got=%d", 10, tok.Column)
	}
}
*/
