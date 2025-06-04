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

	// Leer dos tk. para que curToken y peekToken estén inicializados.
	p.nextToken()
	p.nextToken()

	// Registrar funciones de parseo de prefijo para expresiones.
	p.prefixParseFns = make(map[tk.TokenType]prefixParseFn)
	p.registerPrefix(tk.IDENTIFIER, p.parseIdentifier)
	p.registerPrefix(tk.STRING, p.parseStringLiteral)
	p.registerPrefix(tk.NUMBER, p.parseIntegerLiteral) // <--- NUEVO: Registrar función para números

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
	case tk.DEF:
		return p.parseFlatTripletaStatement()
	case tk.FACT:
		return p.parseFlatTripletaStatement()
	case tk.IDENTIFIER: // Para invocaciones como 'cli print "hello";'
		return p.parseFlatTripletaStatement()
	// ... otros tipos de sentencias si los tienes (ej. return, let)
	default:
		// Si el token actual no inicia una sentencia conocida, podría ser un error
		// o un tipo de expresión que puede ser una sentencia por sí misma.
		// Para tripletas, usualmente se requiere un terminador.
		return nil // O manejar error
	}
}

// parseFlatTripletaStatement parsea una tripleta plana y devuelve un nodo del AST.
func (p *Parser) parseFlatTripletaStatement() *ast.FlatTripletaStatement {
	// 1. Capturar el token inicial de la sentencia (DEF, FACT, o IDENTIFIER como 'cli')
	stmt := &ast.FlatTripletaStatement{Token: p.curToken}

	// 2. Determinar el Sujeto (fts.Subject)
	// El sujeto puede ser el 'p.curToken.Literal' si es una invocación directa,
	// o el siguiente token si es 'def' o 'fact'.

	// Si el token actual es 'def' o 'fact', el sujeto real de la tripleta viene DESPUÉS.
	if p.curToken.Type == tk.DEF || p.curToken.Type == tk.FACT {
		p.nextToken() // Avanza de 'def' o 'fact'
		// Ahora p.curToken debe ser el sujeto (ej. "David")
		subjectIdent, ok := p.parseIdentifier().(*ast.Identifier) // Parse 'David'
		if !ok {
			p.errors = append(p.errors, fmt.Sprintf("expected identifier as subject after %s, got %s", stmt.Token.Literal, p.curToken.Literal))
			return nil
		}
		stmt.Subject = subjectIdent // Asigna el Identifier "David" como Sujeto
	} else if p.curToken.Type == tk.IDENTIFIER {
		// Si el token actual es un IDENTIFIER (ej. "cli", "David"),
		// entonces este token es *tanto* el token inicial de la sentencia *como* el sujeto.
		// No necesitamos `p.nextToken()` aquí, ni `p.parseIdentifier()`.
		// El sujeto es simplemente el `Identifier` que corresponde a `p.curToken`.
		stmt.Subject = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	} else {
		// Esto no debería ocurrir si `parseStatement` ya filtró los tipos de token.
		// Pero es buena práctica tener un fallback.
		p.errors = append(p.errors, fmt.Sprintf("unexpected token type for flat tripleta statement: %s", p.curToken.Literal))
		return nil
	}

	// Avanzar al Verbo (ya sea que hayamos avanzado de 'def'/'fact' o terminado con el sujeto IDENTIFIER)
	p.nextToken()

	// 3. Parsear el Verbo (predicado)
	verbIdent, ok := p.parseIdentifier().(*ast.Identifier)
	if !ok {
		p.errors = append(p.errors, fmt.Sprintf("expected identifier as verb, got %s", p.curToken.Literal))
		return nil
	}
	stmt.Verb = verbIdent

	p.nextToken() // Avanzar del Verbo

	// 4. Parsear el Objeto/Atributo
	stmt.Object = p.parseExpression(tk.LOWEST) // Asume que LOWEST es una constante de precedencia, ej. 0
	if stmt.Object == nil {
		p.errors = append(p.errors, "expected object expression")
		return nil
	}

	// 5. Esperar el delimitador de sentencia (;)
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
