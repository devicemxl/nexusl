// lexer/0preProcess.go
package lexer

import (
	"strings"
)

// StepOne limpia el código de comentarios de línea única (//) y multi-línea (/* */).
// Devuelve el código limpio como un solo string, conservando los saltos de línea y el espaciado original.
func StepOne(fullProgramText string) string {
	var cleanedTextBuilder strings.Builder
	inMultiLineComment := false

	lines := strings.Split(fullProgramText, "\n")

	for _, line := range lines {
		currentLine := line

		// --- Manejo de comentarios multi-línea (/* ... */) ---
		if inMultiLineComment {
			// Estamos dentro de un comentario multi-línea, buscamos el final.
			endIndex := strings.Index(currentLine, "*/")
			if endIndex != -1 {
				// Encontramos el final del comentario. El resto de la línea es código.
				currentLine = currentLine[endIndex+2:] // +2 para saltar "*/"
				inMultiLineComment = false
			} else {
				// Aún estamos dentro del comentario multi-línea. Toda la línea se descarta.
				currentLine = ""
			}
		}

		if !inMultiLineComment {
			// Buscamos el inicio de un comentario multi-línea
			startIndex := strings.Index(currentLine, "/*")
			if startIndex != -1 {
				endIndex := strings.Index(currentLine, "*/")

				if endIndex != -1 && endIndex > startIndex {
					// Comentario multi-línea en una sola línea.
					currentLine = currentLine[:startIndex] + currentLine[endIndex+2:]
				} else {
					// Inicio de un comentario multi-línea que abarca varias líneas.
					currentLine = currentLine[:startIndex] // Conserva la parte antes del /*
					inMultiLineComment = true
				}
			}
		}

		// --- Manejo de comentarios de línea única (//) ---
		// Si no estamos en un comentario multi-línea activo, o si la línea se liberó de uno
		if !inMultiLineComment {
			singleLineCommentIndex := strings.Index(currentLine, "//")
			if singleLineCommentIndex != -1 {
				currentLine = currentLine[:singleLineCommentIndex] // Recorta hasta el //
			}
		}

		// --- Manejo de comentarios de línea única (#) ---
		// Asumiendo que # también es de línea única y se elimina como //
		if !inMultiLineComment {
			hashCommentIndex := strings.Index(currentLine, "#")
			if hashCommentIndex != -1 {
				currentLine = currentLine[:hashCommentIndex] // Recorta hasta el #
			}
		}

		// Agrega la línea limpia (o vacía) al builder, manteniendo el salto de línea.
		cleanedTextBuilder.WriteString(currentLine)
		if len(lines) > 1 || strings.HasSuffix(fullProgramText, "\n") { // Add \n if original had multiple lines or ended with \n
			cleanedTextBuilder.WriteByte('\n')
		}
	}

	return strings.TrimSuffix(cleanedTextBuilder.String(), "\n") // Elimina el último \n si no era intencional
}
