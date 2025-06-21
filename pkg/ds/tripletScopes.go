package ds

import "fmt"

/*

Definir "bloques de contexto" o "secciones" en NexusL para diferentes tipos de construcciones (functions, facts, assertions, algebra, etc.) aporta varios beneficios clave:

Beneficios de los Bloques de Contexto:
Organización y Legibilidad:

Hace el código de NexusL mucho más fácil de leer y entender, especialmente a medida que la base de conocimiento crece. Un desarrollador puede ver rápidamente qué tipo de declaraciones se encuentran en cada sección.
facts { ... } deja claro que esa sección contiene solo declaraciones de hechos base.
algebra { ... } indica que se van a definir expresiones matemáticas.
Claridad Semántica y Propósito:

Los bloques sirven como una señal explícita para el transpilador (y para el programador) sobre el "propósito" de las declaraciones dentro de ellos.
Esto permite al transpilador aplicar diferentes reglas de procesamiento y generación de código Go según el tipo de bloque.
Extensibilidad y Modularidad del Transpilador:

Es mucho más fácil añadir nuevos tipos de funcionalidad (como álgebra) si ya tienes un mecanismo de bloques. Simplemente defines un nuevo tipo de bloque y le asignas su propio módulo de manejo en el transpilador.
Tu parser.go puede tener una función parseFactsBlock(), otra parseAlgebraBlock(), etc.
Posibilidad de Scoping o Namespacing (a futuro):

Aunque Datalog en sí no tiene "scopes" en el sentido de lenguajes imperativos, podrías usar estos bloques para implementar un tipo de "namespacing" o "módulos".
Por ejemplo, context MySystem { ... } podría implicar que todos los predicados o símbolos definidos dentro de ese bloque llevan un prefijo implícito (MySystem_some_predicate) en el Datalog generado, o que pertenecen a un "grafo con nombre" si tu motor Datalog lo soportara (aunque markkurossi/datalog no lo hace de forma nativa).
*/

type tripletScope string

const (
	definition tripletScope = "def"    // NOT EVALUATED - Establish shared understanding
	facts      tripletScope = "fact"   // LOGICAL EVALUATED - a statement that can be objectively verified and is accepted as true
	assertions tripletScope = "assert" // HUMAN / LLM - EVALUATED a viewpoint and can be debated.
	functions  tripletScope = "func"   // expression to be evaluated as method
	expression tripletScope = "expr"   // expression to be evaluated as symbols
)

func (s tripletScope) scope() string {
	switch s {
	case definition:
		return "def"
	case facts:
		return "fact"
	case assertions:
		return "assert"
	case functions:
		return "func"
	case expression:
		return "expr"
	default:
		return fmt.Sprintf("UnknownState(%s)", string(s))
	}
}
