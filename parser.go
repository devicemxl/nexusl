// parser.go
package main

import (
	"fmt"
	"strconv"
)

// Parser holds the state of the parser.
type Parser struct {
	l *Lexer // The lexer from which to get tokens

	errors []string // List of parsing errors

	curToken  Token // The current token under examination
	peekToken Token // The next token (look-ahead)

	// Maps for parsing functions based on token type (prefix/infix parsers)
	prefixParseFns map[TokenType]prefixParseFn
	infixParseFns  map[TokenType]infixParseFn
}

// Function types for parsing expressions (will be used later for full expression parsing)
type (
	prefixParseFn func() Expression
	infixParseFn  func(Expression) Expression
)

// NewParser creates a new Parser instance.
func NewParser(l *Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// Initialize prefixParseFns for literals and identifiers
	p.prefixParseFns = make(map[TokenType]prefixParseFn)
	p.registerPrefix(IDENT, p.parseIdentifier)
	p.registerPrefix(INT, p.parseIntegerLiteral)
	p.registerPrefix(FLOAT, p.parseFloatLiteral)
	p.registerPrefix(STRING, p.parseStringLiteral)
	p.registerPrefix(TRUE, p.parseBooleanLiteral)
	p.registerPrefix(FALSE, p.parseBooleanLiteral)
	p.registerPrefix(NIL, p.parseNilLiteral)
	p.registerPrefix(SYMSIGN, p.parseSymbolIdentifier) // NEW: Register prefix for '$'

	// Advance twice to set both curToken and peekToken
	p.nextToken()
	p.nextToken()

	return p
}

// Errors returns the list of parsing errors.
func (p *Parser) Errors() []string {
	return p.errors
}

// peekError records an error if the peekToken is not of the expected type.
func (p *Parser) peekError(t TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// nextToken advances the parser's tokens.
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// curTokenIs checks if the current token is of a given type.
func (p *Parser) curTokenIs(t TokenType) bool {
	return p.curToken.Type == t
}

// peekTokenIs checks if the next token is of a given type.
func (p *Parser) peekTokenIs(t TokenType) bool {
	return p.peekToken.Type == t
}

// expectPeek checks if the next token is of the expected type and advances if true.
// If not, it records an error.
func (p *Parser) expectPeek(t TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

// registerPrefix registers a prefix parsing function for a given token type.
func (p *Parser) registerPrefix(tokenType TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

// registerInfix registers an infix parsing function for a given token type.
func (p *Parser) registerInfix(tokenType TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

// parseExpression is the main entry point for parsing expressions (will be expanded later).
// For now, it just calls the prefix parser for simple literals/identifiers.
func (p *Parser) parseExpression(precedence int) Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExp := prefix()

	// Later, this loop will handle infix operators based on precedence
	// for now, we only parse simple literals/identifiers
	return leftExp
}

// noPrefixParseFnError records an error when no prefix parsing function is found.
func (p *Parser) noPrefixParseFnError(t TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found (token literal: %s)", t, p.curToken.Literal)
	p.errors = append(p.errors, msg)
}

// --- Parsing Functions for Literals and Identifiers (Prefix Parsers) ---

func (p *Parser) parseIdentifier() Expression {
	return &Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

// NEW: parseSymbolIdentifier handles the '$' prefix for symbols.
func (p *Parser) parseSymbolIdentifier() Expression {
	// The current token is '$' (SYMSIGN). Store it.
	symsignToken := p.curToken

	// The next token must be an IDENT, which is the actual name of the symbol.
	if !p.expectPeek(IDENT) {
		// expectPeek will record an error if it's not an IDENT.
		return nil
	}

	// The current token is now the IDENT (e.g., "david").
	// Create the SymbolIdentifier AST node.
	return &SymbolIdentifier{Token: symsignToken, Value: p.curToken.Literal}
}

func (p *Parser) parseStringLiteral() Expression {
	return &StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseIntegerLiteral() Expression {
	val, err := strconv.ParseInt(p.curToken.Literal, 10, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	return &IntegerLiteral{Token: p.curToken, Value: val}
}

func (p *Parser) parseFloatLiteral() Expression {
	val, err := strconv.ParseFloat(p.curToken.Literal, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as float", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	return &FloatLiteral{Token: p.curToken, Value: val}
}

func (p *Parser) parseBooleanLiteral() Expression {
	return &BooleanLiteral{Token: p.curToken, Value: p.curTokenIs(TRUE)}
}

func (p *Parser) parseNilLiteral() Expression {
	return &NilLiteral{Token: p.curToken}
}

// --- Main Parsing Loop ---

// ParseProgram parses the entire program.
func (p *Parser) ParseProgram() *Program {
	program := &Program{}
	program.Statements = []Statement{}

	for p.curToken.Type != EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken() // Move to the next token, effectively to the start of the next statement
	}

	return program
}

// parseStatement determines the type of statement to parse.
// REEMPLAZA TU parseStatement() ACTUAL CON ESTA FUNCIÓN:
func (p *Parser) parseStatement() Statement {
	switch p.curToken.Type {
	case SYMSIGN: // Handle qualified triplet statements starting with '$'
		return p.parseQualifiedTripletStatement()
	case IDENT: // Handle FlatTripletStatement (like `david corre rapido;`)
		return p.parseFlatTripletStatement()
	default:
		p.errors = append(p.errors, fmt.Sprintf("expected start of statement to be SYMSIGN or IDENT, got %s instead", p.curToken.Type))
		return nil
	}
}

// parseFlatTripletStatement (kept from previous iteration, if you want to support both)
func (p *Parser) parseFlatTripletStatement() Statement {
	// Expect current token to be IDENT (the subject)
	subjectToken := p.curToken
	subject := &Identifier{Token: subjectToken, Value: subjectToken.Literal}

	// Expect and parse the Verb (next token must be an IDENT or keyword)
	if !p.expectPeek(IDENT) && !p.peekTokenIs(DO) && !p.peekTokenIs(HAS) && !p.peekTokenIs(IS) &&
		!p.peekTokenIs(HOW) && !p.peekTokenIs(WHEN) && !p.peekTokenIs(WHERE) {
		p.errors = append(p.errors, fmt.Sprintf("expected verb (IDENT or keyword) after subject, got %s instead", p.peekToken.Type))
		return nil
	}
	verbToken := p.curToken
	verb := &Identifier{Token: verbToken, Value: verbToken.Literal} // Verb is an Identifier, whether keyword or not

	// 3. Expect and parse the Object (can be any Expression type)
	// We use the parseExpression function here, it will automatically
	// call the correct prefix parser for IDENT, STRING, INT, etc.
	p.nextToken()                  // Move from Verb to Object's token
	object := p.parseExpression(0) // Precedence doesn't matter for simple literals

	if object == nil {
		return nil // Error already recorded by parseExpression or its prefix fn
	}

	// 4. Expect the SEMICOLON terminator
	if !p.expectPeek(SEMICOLON) {
		return nil // Error already recorded
	}

	// Construct the FlatTripletStatement AST node
	stmt := &FlatTripletStatement{
		Token:   subjectToken,
		Subject: subject,
		Verb:    verb,
		Object:  object,
	}
	return stmt
}

// NEW: parseQualifiedTripletStatement handles $subject property:value property:value;
func (p *Parser) parseQualifiedTripletStatement() Statement {
	stmt := &QualifiedTripletStatement{}
	stmt.Token = p.curToken // The SYMSIGN token '$'

	// 1. Parse the Subject (which must be a SymbolIdentifier)
	subjectExp := p.parseExpression(0) // This will handle '$' and then the IDENT
	if subjectExp == nil {
		return nil // Error already recorded by parseSymbolIdentifier
	}
	subject, ok := subjectExp.(*SymbolIdentifier)
	if !ok {
		p.errors = append(p.errors, fmt.Sprintf("expected SymbolIdentifier after '$', got %T instead", subjectExp))
		return nil
	}
	stmt.Subject = subject

	stmt.Qualifiers = []*QualifiedProperty{}

	// 2. Parse Qualified Properties (loop until SEMICOLON or EOF)
	for p.peekToken.Type != SEMICOLON && p.peekToken.Type != EOF {
		// --- INICIO DE LA LÓGICA REVISADA PARA PARSEAR LA CLAVE ---
		var keyToken Token
		// Primero, verificamos si el próximo token es un IDENT o una de las palabras clave permitidas.
		if p.peekTokenIs(IDENT) || p.peekTokenIs(DO) || p.peekTokenIs(HAS) || p.peekTokenIs(IS) ||
			p.peekTokenIs(HOW) || p.peekTokenIs(WHEN) || p.peekTokenIs(WHERE) {
			p.nextToken()         // Avanza curToken al token de la clave (e.g., 'do', 'how', o cualquier IDENT)
			keyToken = p.curToken // Asigna el curToken actual como la clave
		} else {
			// Si no es ninguno de los tipos esperados, es un error.
			p.errors = append(p.errors, fmt.Sprintf("expected property key (IDENT or keyword) after subject or previous property, got %s instead", p.peekToken.Type))
			return nil
		}
		key := &Identifier{Token: keyToken, Value: keyToken.Literal}
		// --- FIN DE LA LÓGICA REVISADA ---

		// Expect the COLON
		if !p.expectPeek(COLON) {
			return nil // Error recorded
		}

		// Parse the Object (value). The lexer is currently at COLON, need to move to the value.
		p.nextToken()                  // Move to the object's token (e.g., "corre", "rapido", `corre`, `rapido`)
		object := p.parseExpression(0) // Parse "corre", "rapido", `corre`, `rapido` (could be STRING or IDENT or other literal)
		if object == nil {
			return nil // Error recorded
		}

		// Add the qualified property to the list
		stmt.Qualifiers = append(stmt.Qualifiers, &QualifiedProperty{Key: key, Object: object})
	}

	// 3. Expect the SEMICOLON terminator
	if !p.expectPeek(SEMICOLON) {
		return nil // Error recorded
	}

	return stmt
}
