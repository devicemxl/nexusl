// ast.go
package main

// Node es la interfaz base para todos los nodos del AST.
// Cada nodo del AST debe implementar un método TokenLiteral() string para depuración.
type Node interface {
	TokenLiteral() string // Devuelve el literal del token asociado al nodo.
	String() string       // Devuelve una representación en cadena del nodo (para depuración).
}

// Statement es la interfaz para todos los nodos que representan sentencias.
// Las sentencias no producen un valor directamente en la evaluación.
type Statement interface {
	Node
	statementNode() // Método dummy para marcar que es un Statement.
}

// Expression es la interfaz para todos los nodos que representan expresiones.
// Las expresiones producen un valor en la evaluación.

type Expression interface {
	Node
	expressionNode() // Método dummy para marcar que es una Expression.
}

// Program es el nodo raíz de nuestro AST.
// Contiene una lista de sentencias.
type Program struct {
	Statements []Statement
}

// TokenLiteral para Program (normalmente el del primer token, pero puede ser vacío).
// Esto es útil para depuración, ya que el programa puede no tener un token literal específico.
// En un programa real, podrías querer manejar esto de otra manera.
// Aquí simplemente devolvemos el literal del primer token de la primera sentencia.
// Si no hay sentencias, devolvemos una cadena vacía.
// Esto es útil para saber qué tipo de programa estamos tratando de analizar.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// String para Program (útil para depuración, imprime el programa completo).
// Recorre todas las sentencias y las convierte a cadena.
// Cada sentencia se separa por un salto de línea.
// Esto es útil para ver el programa completo de una vez.
// En un programa real, podrías querer formatear esto de otra manera.
// Aquí simplemente concatenamos todas las cadenas de las sentencias.
func (p *Program) String() string {
	var out []byte
	for _, s := range p.Statements {
		out = append(out, s.String()...)
	}
	return string(out)
}

// --- Nodos Específicos para la Tripleta Aplanada ---

// Identifier representa un identificador (como 'david', 'corre', 'rapido').
// En muchos lenguajes, esto se representa como 'IDENT'.
// En nuestro caso, usamos 'IDENT' tambien.
// El campo Token contendrá el tipo de token (IDENT).
// El campo Value contendrá el valor real (e.g., "david").
// El método String devolverá el literal del token (e.g., "david").
type Identifier struct {
	Token Token // El token IDENT.
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// StringLiteral representa una cadena literal (ej. "hello").
// En muchos lenguajes, esto se representa como '"hello"'.
// En nuestro caso, usamos 'STRING'.
// El campo Token contendrá el tipo de token (STRING).
// El campo Value contendrá el valor real (e.g., "hello").
// El método String devolverá el literal del token (e.g., "\"hello\"").
type StringLiteral struct {
	Token Token // El token STRING.
	Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return "\"" + sl.Value + "\"" } // Para que se imprima con comillas.

// IntegerLiteral representa un número entero (ej. 42).
// En muchos lenguajes, esto se representa como '42'.
// En nuestro caso, usamos 'INT'.
// El campo Token contendrá el tipo de token (INT).
// El campo Value contendrá el valor real (e.g., 42).
// El método String devolverá el literal del token (e.g., "42").
type IntegerLiteral struct {
	Token Token // El token INT.
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal } // El literal ya es la representación.

// FloatLiteral representa un número flotante (ej. 3.14).
// En muchos lenguajes, esto se representa como '3.14'.
// En nuestro caso, usamos 'FLOAT'.
// El campo Token contendrá el tipo de token (FLOAT).
// El campo Value contendrá el valor real (e.g., 3.14).
// El método String devolverá el literal del token (e.g., "3.14").
type FloatLiteral struct {
	Token Token // El token FLOAT.
	Value float64
}

func (fl *FloatLiteral) expressionNode()      {}
func (fl *FloatLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FloatLiteral) String() string       { return fl.Token.Literal }

// BooleanLiteral represents a boolean value (true/false).
// In many languages, this is represented as 'true' or 'false'.
// In our case, we use 'TRUE' and 'FALSE'.
// The Token field will hold the token type (TRUE or FALSE).
// The Value field will hold the actual boolean value (true or false).
type BooleanLiteral struct {
	Token Token // The token TRUE or FALSE.
	Value bool
}

func (b *BooleanLiteral) expressionNode()      {}
func (b *BooleanLiteral) TokenLiteral() string { return b.Token.Literal }
func (b *BooleanLiteral) String() string       { return b.Token.Literal }

// NilLiteral represents the null-like value.
// In many languages, this is represented as 'null', 'nil', or 'None'.
// In our case, we use 'nil'.
// The Token field will hold the token type (NIL).
// The Value field is not needed as nil has no value.
// The String method will return "nil".
type NilLiteral struct {
	Token Token // The token NIL.
}

func (n *NilLiteral) expressionNode()      {}
func (n *NilLiteral) TokenLiteral() string { return n.Token.Literal }
func (n *NilLiteral) String() string       { return "nil" }

// FlatTripletStatement representa una sentencia de tripleta aplanada: sujeto verbo atributo;
// Ejemplo: david corre rapido;
// En este caso, el sujeto es un identificador (e.g., 'david'), el verbo es otro identificador (e.g., 'corre'),
// y el objeto puede ser un identificador, una cadena, un número entero, un número flotante, un booleano o nil.
// La estructura de la tripleta es: (sujeto, verbo, objeto).
type FlatTripletStatement struct {
	Token   Token       // El token del sujeto (e.g., IDENT de "david")
	Subject *Identifier // El sujeto de la tripleta (e.g., david)
	Verb    *Identifier // El verbo/predicado (e.g., corre)
	Object  Expression  // El objeto (e.g., rapido, "rápido", 123)
}

func (fts *FlatTripletStatement) statementNode()       {}
func (fts *FlatTripletStatement) TokenLiteral() string { return fts.Token.Literal }
func (fts *FlatTripletStatement) String() string {
	return fts.Subject.String() + " " + fts.Verb.String() + " " + fts.Object.String() + ";"
}

// --- Nuevas Estructuras para la Tripleta Cualificada ---

// SymbolIdentifier representa un identificador que comienza con '$'.
// Será la expresión que el parser construya cuando vea "$nombre".
type SymbolIdentifier struct {
	Token Token  // El token SYMSIGN
	Value string // El nombre del símbolo (e.g., "david")
}

func (si *SymbolIdentifier) expressionNode()      {}
func (si *SymbolIdentifier) TokenLiteral() string { return si.Token.Literal }
func (si *SymbolIdentifier) String() string       { return "$" + si.Value } // Para que se imprima como $david

// QualifiedProperty representa un par clave-valor como 'do:"corre"' o 'how:rapido'.
// Es como una "sub-tripleta" o un calificador de un sujeto.
type QualifiedProperty struct {
	Key    *Identifier // El calificador (e.g., "do", "how", "when")
	Object Expression  // El valor asociado (e.g., "corre" (StringLiteral), rapido (Identifier))
}

func (qp *QualifiedProperty) String() string {
	return qp.Key.String() + ":" + qp.Object.String()
}

// QualifiedTripletStatement representa la sentencia: $sujeto Propiedad:Valor Propiedad:Valor;
// Ejemplo: $david do:"corre" how:"rapido";
type QualifiedTripletStatement struct {
	Token      Token                // El token '$' que inicia la sentencia
	Subject    *SymbolIdentifier    // El identificador del sujeto (e.g., "$david")
	Qualifiers []*QualifiedProperty // Lista de calificadores (do:"corre", how:"rapido")
}

func (qts *QualifiedTripletStatement) statementNode()       {}
func (qts *QualifiedTripletStatement) TokenLiteral() string { return qts.Token.Literal }
func (qts *QualifiedTripletStatement) String() string {
	var out string
	out += qts.Subject.String() // Print $subject

	for _, qp := range qts.Qualifiers {
		out += " " + qp.String() // Print each qualified property
	}
	out += ";" // Add the terminator
	return out
}
