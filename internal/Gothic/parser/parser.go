// Gothic/parser/parser.go
package parser

import (
	"fmt"
	"strconv"

	"github.com/devicemxl/nexusl/ds"
	"github.com/devicemxl/nexusl/internal/Gothic/ast"
	"github.com/devicemxl/nexusl/internal/Gothic/lexer"
	"github.com/devicemxl/nexusl/internal/Gothic/metamodel"
	"github.com/devicemxl/nexusl/internal/Gothic/token" // Ensure this is imported correctly
)

// Parser holds the lexer, current token, peek token, and errors.
type Parser struct {
	l         *lexer.Lexer
	metamodel *metamodel.MetamodelDefinitions // The metamodel facade

	curToken  token.Token
	peekToken token.Token
	errors    []string
}

// New creates a new Parser instance.
func New(l *lexer.Lexer, mm *metamodel.MetamodelDefinitions) *Parser {
	p := &Parser{
		l:         l,
		metamodel: mm,
		errors:    []string{},
	}
	// Read two tokens, so curToken and peekToken are both set.

	p.nextToken()
	fmt.Printf("inside-DEBUG: parseExpression called. Current Token: Type=%s, Word=%q, Line=%d, Col=%d\n", p.curToken.Type, p.curToken.Word, p.curToken.Line, p.curToken.Column) // DEPURAR

	p.nextToken()
	fmt.Printf("inside-DEBUG: parseExpression called. Current Token: Type=%s, Word=%q, Line=%d, Col=%d\n", p.curToken.Type, p.curToken.Word, p.curToken.Line, p.curToken.Column) // DEPURAR
	// <--- This will initialize both curToken (FACT) and peekToken (Car)
	return p
}

// nextToken advances the parser's current and peek tokens.
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.Ensambladora() // Get the next token from the lexer
}

// ParseProgram parses the entire program source code.
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		// The statement parsing functions are responsible for advancing `p.curToken`
		// to the token *after* the statement's end (e.g., after the semicolon).
		// We don't call
		p.nextToken()
		fmt.Printf("inside-DEBUG: parseExpression called. Current Token: Type=%s, Word=%q, Line=%d, Col=%d\n", p.curToken.Type, p.curToken.Word, p.curToken.Line, p.curToken.Column) // DEPURAR
		//here unconditionally, as parseStatement should handle it.
		// However, if parseStatement returns nil (due to error), we need to advance to avoid infinite loop.
		if stmt == nil {
			// If statement parsing failed, we need to advance to prevent infinite loop
			// by skipping the current problematic token.

			p.nextToken()
			fmt.Printf("inside-DEBUG: parseExpression called. Current Token: Type=%s, Word=%q, Line=%d, Col=%d\n", p.curToken.Type, p.curToken.Word, p.curToken.Line, p.curToken.Column) // DEPURAR
			// Skip the token that caused the error
		}
	}
	return program
}

// parseStatement tries to parse a single statement.
func (p *Parser) parseStatement() ast.Statement {
	fmt.Printf("inside-DEBUG: parseStatement called. Current Token: Type=%s, Word=%q, Line=%d, Col=%d\n",
		p.curToken.Type, p.curToken.Word, p.curToken.Line, p.curToken.Column) // <-- Added DEBUG here

	switch p.curToken.Type {
	case token.FACT:
		return p.parseFactStatement()
	default:
		p.noCurTokenError(token.FACT) // Report that we expected 'fact' keyword
		return nil
	}
}

// parseFactStatement parsea una declaración 'fact' atómica: 'fact Subject Predicate Object;'
func (p *Parser) parseFactStatement() *ast.FactStatement {
	factToken := p.curToken // Capturamos el token 'fact' (Type=FACT_KEYWORD)

	factScopeSymbol, ok := p.metamodel.LookupScope(factToken.Word)
	if !ok || factScopeSymbol.Thing != ds.TripletScopeType {
		p.errors = append(p.errors, fmt.Sprintf("Line %d, Column %d: Unknown or invalid scope '%s'", factToken.Line, factToken.Column, factToken.Word))
		return nil
	}

	p.nextToken()
	fmt.Printf("inside-DEBUG: parseExpression called. Current Token: Type=%s, Word=%q, Line=%d, Col=%d\n", p.curToken.Type, p.curToken.Word, p.curToken.Line, p.curToken.Column) // DEPURAR
	// Consume 'fact'. curToken ahora es 'Car'

	// Sujeto
	subject := p.parseExpression() // Parses 'Car' (Type=IDENTIFIER)
	if subject == nil {
		return nil
	}

	p.nextToken()
	fmt.Printf("inside-DEBUG: parseExpression called. Current Token: Type=%s, Word=%q, Line=%d, Col=%d\n", p.curToken.Type, p.curToken.Word, p.curToken.Line, p.curToken.Column) // DEPURAR
	// Consumes 'Car'. curToken ahora es 'is'

	// Predicado
	predicate := p.parseExpression() // Parses 'is' (Type=IS)
	if predicate == nil {
		return nil
	}

	p.nextToken()
	fmt.Printf("inside-DEBUG: parseExpression called. Current Token: Type=%s, Word=%q, Line=%d, Col=%d\n", p.curToken.Type, p.curToken.Word, p.curToken.Line, p.curToken.Column) // DEPURAR
	// Consumes 'is'. curToken ahora es 'symbol'

	// Objeto
	object := p.parseExpression() // Parses 'symbol' (Type=SYMBOL)
	if object == nil {
		return nil
	}

	p.nextToken()
	fmt.Printf("inside-DEBUG: parseExpression called. Current Token: Type=%s, Word=%q, Line=%d, Col=%d\n", p.curToken.Type, p.curToken.Word, p.curToken.Line, p.curToken.Column) // DEPURAR
	// Consumes 'symbol'. curToken ahora es ';'

	// Esperar y consumir el punto y coma final
	if !p.curTokenIs(token.SEMICOLON) {
		p.noCurTokenError(token.SEMICOLON)
		return nil
	}

	p.nextToken()
	fmt.Printf("inside-DEBUG: parseExpression called. Current Token: Type=%s, Word=%q, Line=%d, Col=%d\n", p.curToken.Type, p.curToken.Word, p.curToken.Line, p.curToken.Column) // DEPURAR
	// Consume ';'. curToken ahora es EOF (or the start of next statement)

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
	case token.IS: // This should be "IS" (uppercase)
		// If 'is' is a keyword token, treat it as an identifier for the AST
		return &ast.Identifier{Token: p.curToken, Value: p.curToken.Word}
	case token.SYMBOL: // This should be "SYMBOL" (uppercase)
		// Treat "symbol" as an identifier in the AST for now.
		return &ast.Identifier{Token: p.curToken, Value: p.curToken.Word}
	default:
		p.errors = append(p.errors, fmt.Sprintf("Line %d, Column %d: Unexpected token %s (%q) when expecting an expression.",
			p.curToken.Line, p.curToken.Column, p.curToken.Type, p.curToken.Word))
		return nil
	}
}

// parseIdentifier parsea un token IDENTIFIER y lo convierte en un nodo *ast.Identifier.
func (p *Parser) parseIdentifier() *ast.Identifier {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Word}
}

// parseStringLiteral parsea un token STRING y lo convierte en un nodo *ast.StringLiteral.
func (p *Parser) parseStringLiteral() *ast.StringLiteral {
	return &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Word}
}

// parseIntegerLiteral parsea un token INT y lo convierte en un nodo *ast.IntegerLiteral.
func (p *Parser) parseIntegerLiteral() *ast.IntegerLiteral {
	val, err := strconv.ParseInt(p.curToken.Word, 0, 64)
	if err != nil {
		p.errors = append(p.errors, fmt.Sprintf("Line %d, Column %d: Could not parse %q as integer: %v",
			p.curToken.Line, p.curToken.Column, p.curToken.Word, err))
		return nil
	}
	return &ast.IntegerLiteral{Token: p.curToken, Value: val}
}

// parseFloatLiteral parsea un token FLOAT y lo convierte en un nodo *ast.FloatLiteral.
func (p *Parser) parseFloatLiteral() *ast.FloatLiteral {
	val, err := strconv.ParseFloat(p.curToken.Word, 64)
	if err != nil {
		p.errors = append(p.errors, fmt.Sprintf("Line %d, Column %d: Could not parse %q as float: %v",
			p.curToken.Line, p.curToken.Column, p.curToken.Word, err))
		return nil
	}
	return &ast.FloatLiteral{Token: p.curToken, Value: val}
}

// parseBooleanLiteral parsea un token BOOLEAN y lo convierte en un nodo *ast.BooleanLiteral.
func (p *Parser) parseBooleanLiteral() *ast.BooleanLiteral {
	val := (p.curToken.Word == "true")
	return &ast.BooleanLiteral{Token: p.curToken, Value: val}
}

// Helper methods for token checking and error reporting (no changes needed here)
func (p *Parser) curTokenIs(t token.TokenClass) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenClass) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenClass) bool {
	if p.peekTokenIs(t) {

		p.nextToken()
		fmt.Printf("inside-DEBUG: parseExpression called. Current Token: Type=%s, Word=%q, Line=%d, Col=%d\n", p.curToken.Type, p.curToken.Word, p.curToken.Line, p.curToken.Column) // DEPURAR

		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenClass) {
	msg := fmt.Sprintf("Expected next token to be %s, got %s instead (%q)",
		t, p.peekToken.Type, p.peekToken.Word)
	p.errors = append(p.errors, msg)
}

func (p *Parser) noCurTokenError(t token.TokenClass) {
	msg := fmt.Sprintf("Expected current token to be %s, got %s instead (%q)",
		t, p.curToken.Type, p.curToken.Word)
	p.errors = append(p.errors, msg)
}
