package lexer

import (
	"fmt"
	"strings"

	tk "github.com/devicemxl/nexusl/pkg/token"
)

// StepOne toma una línea de código y la limpia de comentarios,
// devolviendo un slice de palabras.
func StepOne(line string) []string {
	/*
	* ============================================= * ==
	* Clean the code from (one line) comments
	* Before Continue
	* ============================================= * ==
	 */
	words := strings.Fields(line)
	// Elimina los comentarios que empiezan con "//"
	for i, word := range words {
		if strings.HasPrefix(word, "//") {
			words = words[:i] // Elimina el comentario y todo lo que sigue
			break
		}
	}
	/*
		LOS COMENTARIOS MULTILINEA SON UNA ESTRUCTURA
		NO PUEDEN ANALIZARSE LINEA POR LINEA
		ENTONCES DEBEN MARCARSE LOS TOKEN Y DAR TRATAMIENTO
		DE CULQUIER OTRA DS
		OMITIR YA EN EL LEXER
	*/
	// Elimina los comentarios que empiezan con "/*" y terminan con "*/"
	if strings.Contains(line, "/*") && strings.Contains(line, "*/") {
		// IF "/*" INDEX < "*/" INDEX
		index1 := strings.Index(line, "/*")
		index2 := strings.Index(line, "*/")
		//
		if index1 < index2 {
			words = strings.Fields(line[:index1])
		}
	}

	// Elimina los comentarios que empiezan con "#"
	for i, word := range words {
		if strings.HasPrefix(word, "#") {
			words = words[:i] // Elimina el comentario y todo lo que sigue
			break
		}
	}

	// Elimina los comentarios dentro de lineas que contienen "//"
	// Esto es para casos donde "//" está en medio de una palabra o línea
	// prtimos la linea en caracteres
	if strings.Contains(line, "//") {
		// Elimina el comentario y todo lo que sigue
		// Busca el índice del comentario "//"
		index := strings.Index(line, "//")
		if index != -1 {
			// Recorta la línea hasta antes del comentario
			line = line[:index]
			words = strings.Fields(line)
		}
	}

	return words
}

// ingresa una Line y regresa un slice de tokens
func StepTwo(words []string) []tk.Token {
	tokenized := []tk.Token{}
	// si el slice de palabras es vacío, devuelve un token vacío
	if len(words) < 1 {
		var ThisWord tk.Token
		ThisWord.Type = tk.EMPTY
		ThisWord.Word = ""
		tokenized = append(tokenized, ThisWord)
		return tokenized

	} else {
		// Si hay palabras, las tokeniza
		// y las clasifica
		//
		// Para eso usamos la misma logica que cuando
		// eliminamos los comentarios dentro de lineas que contienen "//"
		// .
		for _, word := range words {
			var ThisWord tk.Token
			ThisWord.Word = word
			clf := tk.ClassifyToken(word) // Aquí se puede agregar la lógica de clasificación
			fmt.Println(clf)
			tokenized = append(tokenized, ThisWord)
		}
		//fmt.Println(tokenized)
		//globalLexer[lineNumber] = tokenized
		return tokenized

	}
	//
}
