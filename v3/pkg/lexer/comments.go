// Package lexer proporciona un analizador léxico para el lenguaje NexusL.
package lexer

/*
Los comentarios son, en esencia, una forma de "espacio en blanco" que el lexer debe ignorar completamente, al igual que los espacios, tabulaciones y nuevas líneas.
*/
// skipWhitespace consume espacios en blanco y comentarios de una sola línea.
func (l *Lexer) skipWhitespace() {
	for {
		// Consume espacios en blanco normales
		for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
			l.readChar()
		}

		// Verifica si el siguiente par de caracteres es el inicio de un comentario de una sola línea `//`
		if l.ch == '/' && l.peekChar() == '/' {
			// Es un comentario de una sola línea
			l.skipLineComment()
		} else {
			// No hay más espacios en blanco ni comentarios, salimos del bucle
			break
		}
	}
}

// skipLineComment consume todos los caracteres hasta el final de la línea o EOF.
func (l *Lexer) skipLineComment() {
	// Ya sabemos que l.ch es '/' y l.peekChar() es '/'.
	// Consumimos ambos caracteres iniciales del comentario.
	l.readChar() // Consume el primer '/'
	l.readChar() // Consume el segundo '/'

	// Ahora consumimos el resto de la línea hasta un '\n' o EOF.
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
	// Opcional: Si quieres que el lexer salte también el '\n' al final del comentario.
	// l.readChar() // Consume el '\n' para que la próxima llamada a NextToken no lo procese
}
