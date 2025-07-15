package parser

import (
	"fmt"
	"strconv"

	"github.com/devicemxl/nexusl/Gothic/ast"
	"github.com/devicemxl/nexusl/Gothic/ds" // Necesario para ds.Symbol
	"github.com/devicemxl/nexusl/Gothic/lexer"
	"github.com/devicemxl/nexusl/Gothic/metamodel" // Para el metamodelo líquido
	"github.com/devicemxl/nexusl/Gothic/token"
)

// Parser representa la instancia del analizador sintáctico
type Parser struct {
	l *lexer.Lexer // El lexer del que obtiene tokens

	errors []string // Almacena los errores de parseo

	curToken  token.Token // Token actual que estamos examinando
	peekToken token.Token // Siguiente token después del actual (para "mirar adelante")

	metamodel *metamodel.MetamodelDefinitions // Las definiciones cargadas del metamodelo
}

// New crea e inicializa un nuevo Parser
func New(l *lexer.Lexer, mm *metamodel.MetamodelDefinitions) *Parser {
	p := &Parser{
		l:         l,
		errors:    []string{},
		metamodel: mm,
	}

	// Leer dos tokens para que curToken y peekToken estén inicializados
	p.nextToken()
	p.nextToken()

	return p
}

// nextToken avanza los punteros de tokens
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.Ensambladora()
}

// curTokenIs verifica si el token actual es del tipo esperado
func (p *Parser) curTokenIs(t token.TokenClass) bool {
	return p.curToken.Type == t
}

// peekTokenIs verifica si el PRÓXIMO token es del tipo esperado
func (p *Parser) peekTokenIs(t token.TokenClass) bool {
	return p.peekToken.Type == t
}

// expectPeek espera que el próximo token sea del tipo esperado y lo consume.
// Si no lo es, añade un error.
func (p *Parser) expectPeek(t token.TokenClass) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

// Errors devuelve la lista de errores encontrados
func (p *Parser) Errors() []string {
	return p.errors
}

// peekError añade un error si el token de anticipación no es el esperado
func (p *Parser) peekError(t token.TokenClass) {
	msg := fmt.Sprintf("Line %d, Column %d: Expected next token to be %s, got %s instead (%q)",
		p.peekToken.Line, p.peekToken.Column, t, p.peekToken.Type, p.peekToken.Word)
	p.errors = append(p.errors, msg)
}

// noCurTokenError añade un error para el token actual si no es el esperado
func (p *Parser) noCurTokenError(t token.TokenClass) {
	msg := fmt.Sprintf("Line %d, Column %d: Expected current token to be %s, got %s instead (%q)",
		p.curToken.Line, p.curToken.Column, t, p.curToken.Type, p.curToken.Word)
	p.errors = append(p.errors, msg)
}

/*
// --- Métodos de Ayuda para el Metamodelo ---
// getScopeSymbol busca la definición del scope en el metamodelo y devuelve un *ds.Symbol
func (p *Parser) getScopeSymbol(scopeName string) (*ds.Symbol, bool) {
	def, exists := p.metamodel.Scopes[scopeName]
	if !exists {
		return nil, false
	}
	// Aquí creas o recuperas el ds.Symbol real basándote en la definición.
	// Por simplicidad, asumimos que ds.Symbol tiene un constructor o que ya existe.
	// En un sistema real, podrías tener un SymbolTable que gestiona estas instancias.
	return &ds.Symbol{Name: def.Name, ThingType: def.Type}, true // Simplificado
}
*/
// --- Funciones de Parseo Principales ---

// ParseProgram es el método principal que parsea todo el programa NexusL
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	// Bucle principal de parseo, consume sentencias hasta EOF
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		// Si parseStatement no pudo parsear nada (retornó nil), avanzamos el token
		// para evitar bucles infinitos en caso de errores de sintaxis no recuperables.
		if stmt == nil && p.curToken.Type != token.EOF {
			p.nextToken()
		}
	}

	return program
}

// parseStatement intenta parsear una única sentencia.
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.FACT: // Usar el nuevo FACT_KEYWORD
		return p.parseFactStatement()
	default:
		p.noCurTokenError(token.FACT) // Reporta que esperaba 'fact' keyword
		return nil
	}
}

// parseFactStatement parsea una declaración 'fact' atómica: 'fact Subject Predicate Object;'
func (p *Parser) parseFactStatement() *ast.FactStatement {
	factToken := p.curToken

	factScopeSymbol, ok := p.metamodel.LookupScope(factToken.Word)
	if !ok || factScopeSymbol.Thing != ds.TripletScopeType {
		p.errors = append(p.errors, fmt.Sprintf("Line %d, Column %d: Unknown or invalid scope '%s'", factToken.Line, factToken.Column, factToken.Word))
		return nil
	}
	p.nextToken() // Consume 'fact'

	// Parsear el Sujeto usando la nueva función general parseExpression
	subject := p.parseExpression()
	if subject == nil {
		return nil // Error ya reportado por parseExpression
	}
	p.nextToken() // Consume el token del Sujeto

	// Parsear el Predicado usando la nueva función general parseExpression
	predicate := p.parseExpression()
	if predicate == nil {
		return nil // Error ya reportado por parseExpression
	}
	p.nextToken() // Consume el token del Predicado

	// Parsear el Objeto usando la nueva función general parseExpression
	object := p.parseExpression()
	if object == nil {
		return nil // Error ya reportado por parseExpression
	}
	p.nextToken() // Consume el token del Objeto

	// Esperar y consumir el punto y coma final
	if !p.curTokenIs(token.SEMICOLON) {
		p.noCurTokenError(token.SEMICOLON)
		return nil
	}
	p.nextToken() // Consume ';' para la próxima iteración

	return &ast.FactStatement{
		Token:     factToken,
		Scope:     factScopeSymbol,
		Subject:   subject,
		Predicate: predicate,
		Object:    object,
	}
}

// parseExpression es la función principal que decide qué tipo de expresión parsear
func (p *Parser) parseExpression() ast.Expression {
	fmt.Printf("DEBUG: parseExpression called. Current Token: Type=%s, Word=%q, Line=%d, Col=%d\n",
		p.curToken.Type, p.curToken.Word, p.curToken.Line, p.curToken.Column)

	switch p.curToken.Type {
	case token.IDENTIFIER:
		return p.parseIdentifier()
	case token.STRING:
		return p.parseStringLiteral()
	case token.INTEGER:
		return p.parseIntegerLiteral()
	case token.FLOAT:
		return p.parseFloatLiteral()
	case token.BOOLEAN:
		return p.parseBooleanLiteral()
	case token.IS: // Si "is" es un token de palabra clave, también puede ser una expresión predicado
		return &ast.Identifier{Token: p.curToken, Value: p.curToken.Word} // Trátalo como Identifier por ahora
		// Agrega más casos aquí a medida que implementes el parseo de otras expresiones:
		// case token.LPAREN: return p.parseCallExpression() // Para (robot move ...)
		// case token.LBRACE: return p.parseMapLiteral() // Para @{ ... } o how::"fast"
		// case token.LBRACKET: return p.parseListLiteral() // Para @[ ... ]
	default:
		p.errors = append(p.errors, fmt.Sprintf("Line %d, Column %d: Unexpected token %s (%q) when expecting an expression.",
			p.curToken.Line, p.curToken.Column, p.curToken.Type, p.curToken.Word))
		return nil
	}
}

// parseIdentifier parsea un token IDENTIFIER y lo convierte en un nodo *ast.Identifier.
// Ahora solo se llama desde parseExpression.
func (p *Parser) parseIdentifier() *ast.Identifier {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Word}
}

// Nuevas funciones para parsear literales
func (p *Parser) parseStringLiteral() *ast.StringLiteral {
	return &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Word}
}

func (p *Parser) parseIntegerLiteral() *ast.IntegerLiteral {
	val, err := strconv.ParseInt(p.curToken.Word, 0, 64)
	if err != nil {
		p.errors = append(p.errors, fmt.Sprintf("Line %d, Column %d: Could not parse %q as integer: %v",
			p.curToken.Line, p.curToken.Column, p.curToken.Word, err))
		return nil
	}
	return &ast.IntegerLiteral{Token: p.curToken, Value: val}
}

func (p *Parser) parseFloatLiteral() *ast.FloatLiteral {
	val, err := strconv.ParseFloat(p.curToken.Word, 64)
	if err != nil {
		p.errors = append(p.errors, fmt.Sprintf("Line %d, Column %d: Could not parse %q as float: %v",
			p.curToken.Line, p.curToken.Column, p.curToken.Word, err))
		return nil
	}
	return &ast.FloatLiteral{Token: p.curToken, Value: val}
}

func (p *Parser) parseBooleanLiteral() *ast.BooleanLiteral {
	val := (p.curToken.Word == "true")
	return &ast.BooleanLiteral{Token: p.curToken, Value: val}
}
