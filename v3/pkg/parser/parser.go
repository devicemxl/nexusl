// Package parser proporciona un parser para el lenguaje NexusL.
package parser

import (
	"fmt"
	"strconv"

	"github.com/devicemxl/nexusl/pkg/ast"
	lx "github.com/devicemxl/nexusl/pkg/lexer"
	tk "github.com/devicemxl/nexusl/pkg/token"
)

// Parser es una estructura que representa un parser para el lenguaje NexusL.
type Parser struct {
	// l es el lexer que proporciona los tk..
	l *lx.Lexer

	// curToken es el tk.actual bajo inspección.
	curToken tk.Token
	// peekToken es el próximo tk.
	peekToken tk.Token
	// errors es una lista de errores encontrados durante el parseo.
	errors []string

	// prefixParseFns es un mapa de funciones de parseo de prefijo para expresiones.
	prefixParseFns map[tk.TokenType]prefixParseFn
	// infixParseFns es un mapa de funciones de parseo de infix para expresiones.
	infixParseFns map[tk.TokenType]infixParseFn
}

type (
	// prefixParseFn es una función que parsea una expresión de prefijo.
	prefixParseFn func() ast.Expression
	// infixParseFn es una función que parsea una expresión de infix.
	infixParseFn func(ast.Expression) ast.Expression
)

// NewParser devuelve un nuevo parser para el lenguaje NexusL.
func NewParser(l *lx.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	// Registrar funciones de parseo de prefijo para expresiones.
	p.prefixParseFns = make(map[tk.TokenType]prefixParseFn)
	p.registerPrefix(tk.IDENTIFIER, p.parseIdentifier)
	p.registerPrefix(tk.STRING, p.parseStringLiteral)
	p.registerPrefix(tk.INTEGER, p.parseIntegerLiteral) // <--- NUEVO: Registrar función para números

	// Leer dos tk. para que curToken y peekToken estén inicializados.
	p.nextToken()
	p.nextToken()
	return p
}

// Errors devuelve la lista de errores encontrados durante el parseo.
func (p *Parser) Errors() []string {
	return p.errors
}

// peekError agrega un error a la lista de errores si el próximo tk.no es del tipo esperado.
func (p *Parser) peekError(t tk.TokenType) {
	msg := fmt.Sprintf("expected next tk.to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// nextToken avanza al siguiente tk.
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram parsea un programa completo y devuelve un AST.
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != tk.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken() // Avanzar al siguiente tk.para la próxima iteración.
	}
	return program
}

// parseStatement parsea una sentencia y devuelve un nodo del AST.
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case tk.DEF, tk.FACT, tk.ASSIGN: // Ahora todos son manejados por la misma función
		return p.parseFlatTripletaStatement()
	case tk.IDENTIFIER: // Si empieza con un identificador, también puede ser una FlatTripletaStatement (para invocaciones)
		// Esta es la parte crucial que faltaba o estaba mal manejada.
		// Queremos que "cli print "hello";" sea una FlatTripletaStatement.
		// La lógica en `parseFlatTripletaStatement` debe diferenciar.
		// Podrías tener una función `parseInvocationStatement` si las invocaciones son fundamentalmente diferentes.
		// Por simplicidad, intentemos que `parseFlatTripletaStatement` maneje ambos.
		return p.parseFlatTripletaStatement() // Asume que una tripleta sin prefijo es también una FlatTripleta
	default:
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
}

// parseDefStatement es una nueva función para parsear la sentencia 'def'.
// Su estructura es idéntica a parseFlatTripletaStatement, pero asegura
// que el Token.Type sea tk.DEF.
func (p *Parser) parseDefStatement() *ast.FlatTripletaStatement {
	stmt := &ast.FlatTripletaStatement{Token: p.curToken} // El token será tk.DEF

	// Sujeto
	p.nextToken()                               // Avanzar al sujeto
	stmt.Subject = p.parseExpression(tk.LOWEST) // LOWEST precedence

	// Predicado (Verb)
	p.nextToken()                     // Avanzar al predicado
	if !p.curTokenIs(tk.IDENTIFIER) { // Los predicados son identificadores
		p.errors = append(p.errors, fmt.Sprintf("expected identifier as verb, got %s", p.curToken.Literal))
		return nil
	}
	stmt.Verb = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// Objeto
	p.nextToken()                              // Avanzar al objeto
	stmt.Object = p.parseExpression(tk.LOWEST) // LOWEST precedence

	// Consumir el punto y coma
	if !p.expectPeek(tk.SEMICOLON) {
		return nil
	}

	return stmt
}

// parseFlatTripletaStatement parsea una tripleta plana y devuelve un nodo del AST.
// parseFlatTripletaStatement es ahora más genérica y se usa para DEF, FACT, ASSIGN y tripletas implícitas.
// parseFlatTripletaStatement parsea una sentencia FlatTripletaStatement.
// Ejemplos:
//
//	def David is:symbol;
//	fact David has:rightLeg "healthy";
//	cli print "hello";
//	David has:age 30;
func (p *Parser) parseFlatTripletaStatement() *ast.FlatTripletaStatement {
	stmt := &ast.FlatTripletaStatement{Token: p.curToken}

	// El primer token es el token principal (DEF, FACT, o el Sujeto si es una invocación directa)
	// Si es DEF o FACT, el sujeto real viene después.
	if stmt.Token.Type == tk.DEF || stmt.Token.Type == tk.FACT {
		if !p.expectPeek(tk.IDENTIFIER) {
			return nil
		}
		stmt.Subject = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		p.nextToken() // Avanzar al siguiente token (el verbo/predicado)
	} else {
		// Si no es DEF o FACT, asumimos que curToken es el Sujeto (ej. "cli" en "cli print").
		// Asegúrate de que el sujeto sea un identificador.
		subjectExpr := p.parseIdentifier() // Esto consume el token actual (el sujeto)
		if subjectExpr == nil {
			return nil // Error ya reportado por parseIdentifier si no es un identificador
		}
		stmt.Subject = subjectExpr
		// El nextToken() se hará implícitamente por el parseIdentifier que ya avanzó.
		// No, parseIdentifier no avanza, solo crea el nodo. Necesitamos avanzar.
		p.nextToken() // Avanzar al siguiente token (el verbo/predicado)
	}

	// Parsear el Verbo/Predicado (siempre debe ser un IDENTIFIER)
	if !p.curTokenIs(tk.IDENTIFIER) {
		p.peekError(tk.IDENTIFIER)
		return nil
	}
	stmt.Verb = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	p.nextToken() // Avanzar al siguiente token (el Objeto o ;)

	// Parsear el Objeto (puede ser un IDENTIFIER, STRING, NUMBER, BOOLEAN, o nulo)
	// Aquí es donde cambia: ya no esperamos un IDENTIFIER forzosamente,
	// sino cualquier EXPRESIÓN válida que pueda ser un objeto.
	if !p.curTokenIs(tk.SEMICOLON) { // Si no es un punto y coma, entonces hay un objeto
		objExp := p.parseExpression(tk.LOWEST) // Intentar parsear una expresión
		if objExp == nil {
			return nil // Error ya reportado
		}
		stmt.Object = objExp
	} else {
		// No hay objeto explícito, se asigna un objeto nulo o un valor por defecto.
		// En NexusL, los objetos pueden ser omitidos si el predicado implica una acción sin complemento.
		// Podrías tener un ast.NullLiteral si lo defines, o simplemente dejarlo como nil
		// y el evaluador lo manejará. Por ahora, lo dejaremos en nil y el evaluador lo manejará.
		stmt.Object = nil // Representa la ausencia de un objeto explícito
	}

	// Asegurarse de que termine con un punto y coma si el objeto no lo consumió (como un StringLiteral no lo hace)
	if !p.expectPeek(tk.SEMICOLON) {
		return nil
	}

	return stmt
}

// curTokenIs devuelve true si el tk.actual es del tipo esperado.
func (p *Parser) curTokenIs(t tk.TokenType) bool {
	return p.curToken.Type == t
}

// peekTokenIs devuelve true si el próximo tk.es del tipo esperado.
func (p *Parser) peekTokenIs(t tk.TokenType) bool {
	return p.peekToken.Type == t
}

// expectPeek devuelve true si el próximo tk.es del tipo esperado y avanza al siguiente tk.
func (p *Parser) expectPeek(t tk.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

// parseExpression maneja la lógica de parseo de expresiones.
func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExp := prefix()

	// Por ahora no hay operadores infix, así que solo devolvemos la expresión prefijo.
	return leftExp
}

// registerPrefix registra una función de parseo de prefijo para un tipo de tk.
func (p *Parser) registerPrefix(tokenType tk.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

// registerInfix registra una función de parseo de infix para un tipo de tk.
func (p *Parser) registerInfix(tokenType tk.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

// parseIdentifier parsea un identificador y devuelve un nodo del AST.
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

// parseStringLiteral parsea un literal de cadena y devuelve un nodo del AST.
func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
}

// NUEVO: parseIntegerLiteral convierte el token NUMBER en un IntegerLiteral.
func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value
	return lit
}

// noPrefixParseFnError agrega un error a la lista de errores si no hay una función de parseo de prefijo para un tipo de tk.
func (p *Parser) noPrefixParseFnError(t tk.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}
