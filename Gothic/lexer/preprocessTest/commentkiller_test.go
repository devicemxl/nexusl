package lexer_test // Nota: Cambiado a package GothicLexer_test

import (
	"strings"
	"testing"

	"github.com/devicemxl/nexusl/Gothic/lexer" // Ajusta la ruta de importación si es diferente
)

func TestCleanSingleLineComments(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Comentario de una línea al final",
			input:    "let x := 10 // Esto es un comentario",
			expected: "let x := 10 ", // O "let x := 10" si recortas espacios finales
		},
		{
			name:     "Comentario de una línea al inicio",
			input:    "// Un comentario inicial\nfunc foo()",
			expected: "\nfunc foo()", // La línea del comentario se vacía
		},
		{
			name:     "Comentario de una línea con #",
			input:    "rule A is B # Esto es otro comentario",
			expected: "rule A is B ",
		},
		{
			name:     "Múltiples comentarios en una misma línea (solo el primero afecta)",
			input:    "value = 1 // Comentario 1 // Comentario 2",
			expected: "value = 1 ",
		},
		{
			name:     "Código sin comentarios",
			input:    "let result := (2 + 3) * 4",
			expected: "let result := (2 + 3) * 4",
		},
		{
			name:     "Línea solo con comentario",
			input:    "// Solo comentarios",
			expected: "",
		},
		{
			name:     "Múltiples líneas con y sin comentarios",
			input:    "func main()\n    let x := 1 // var x\n    let y := 2 # var y\n    return x + y\n",
			expected: "func main()\n    let x := 1 \n    let y := 2 \n    return x + y\n",
		},
		{
			name:     "Comentarios multilínea (deben ser ignorados por esta función)",
			input:    "/* Bloque de\n   comentario multilínea */\nlet z := 10;",
			expected: "/* Bloque de\n   comentario multilínea */\nlet z := 10;",
		},
		{
			name:     "Comentarios multilínea en la misma línea",
			input:    "value = /* inline */ 10;",
			expected: "value = /* inline */ 10;",
		},
		{
			name: "Varias líneas sin comentarios",
			input: `
rule is_active (room) {
    if (room.state == "active")
}
`,
			expected: `
rule is_active (room) {
    if (room.state == "active")
}
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := lexer.StepOne(tt.input)

			// Usamos strings.TrimSuffix para quitar saltos de línea extra si es necesario,
			// y strings.TrimSpace para la línea individual.
			// Esto es para manejar las diferencias sutiles en los espacios finales después del recorte.
			processedExpectedLines := []string{}
			for _, line := range strings.Split(tt.expected, "\n") {
				processedExpectedLines = append(processedExpectedLines, strings.TrimRight(line, " ")) // Quitamos espacios al final de la línea
			}
			processedExpected := strings.Join(processedExpectedLines, "\n")

			processedActualLines := []string{}
			for _, line := range strings.Split(actual, "\n") {
				processedActualLines = append(processedActualLines, strings.TrimRight(line, " "))
			}
			processedActual := strings.Join(processedActualLines, "\n")

			if processedActual != processedExpected {
				t.Errorf("Test '%s' ERROR:\nEsperado:\n%q\nObtenido:\n%q",
					tt.name, processedExpected, processedActual)
			}
		})
	}
}
