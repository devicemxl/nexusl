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
	p.registerPrefix(SYMSIGN, p.parseSymbolIdentifier)
	p.registerPrefix(LBRACKET, p.parseListLiteral) // Register prefix for '[' to parse lists

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

// parseExpression is the main entry point for parsing expressions.
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

// parseSymbolIdentifier handles the '$' prefix for symbols.
func (p *Parser) parseSymbolIdentifier() Expression {
	symsignToken := p.curToken

	if !p.expectPeek(IDENT) { // Expect the identifier name after '$'
		return nil
	}

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

// parseListLiteral handles parsing of list expressions like [elem1, elem2, ...]
func (p *Parser) parseListLiteral() Expression {
	list := &ListLiteral{Token: p.curToken}         // Current token is LBRACKET '['
	list.Elements = p.parseExpressionList(RBRACKET) // Parse elements until RBRACKET ']'
	return list
}

// parseExpressionList is a helper function to parse a comma-separated list of expressions
// until a given end token (e.g., RPAREN, RBRACKET, RCURLY).
func (p *Parser) parseExpressionList(endToken TokenType) []Expression {
	var list []Expression

	// If the next token is the end token, it's an empty list
	if p.peekTokenIs(endToken) {
		p.nextToken() // Consume the end token
		return list
	}

	p.nextToken() // Advance to the first element

	// Parse the first element
	list = append(list, p.parseExpression(0)) // Precedence 0 for top-level expressions

	// Loop to parse subsequent elements
	for p.peekTokenIs(COMMA) || p.peekTokenIs(SEMICOLON) { // Allow both comma and semicolon as separators
		p.nextToken() // Consume COMMA or SEMICOLON
		p.nextToken() // Advance to the next element
		list = append(list, p.parseExpression(0))
	}

	// Expect the end token
	if !p.expectPeek(endToken) {
		return nil // Error already recorded by expectPeek
	}

	return list
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
		// Only advance if the current token is not the EOF, and if the statement
		// parsing didn't already consume the next statement's start token.
		// For now, assume statements consume their own terminators (like SEMICOLON).
		// This `p.nextToken()` here might over-advance if a statement parser
		// already moved past the terminator. It's safer to let each statement
		// parser handle its own token consumption until its logical end.
		// However, for simplicity with current grammar, we'll keep it.
		if p.curToken.Type != EOF {
			// If the statement didn't consume the semicolon, consume it here
			// This part needs careful design based on exact grammar.
			// For now, assuming statements end with SEMICOLON and their parsers consume it.
			// If not, this nextToken() could skip valid tokens.
			// A more robust approach is to check if the current token is a terminator
			// before advancing to the next statement.
		}
	}

	return program
}

// parseStatement determines the type of statement to parse.
func (p *Parser) parseStatement() Statement {
	switch p.curToken.Type {
	case SYMSIGN: // Handle qualified triplet statements starting with '$'
		return p.parseQualifiedTripletStatement()
	case IDENT: // Handle FlatTripletStatement (like `david corre rapido;`) or function definition (if "func" is IDENT)
		// Check if it's "func" keyword
		if p.curToken.Literal == "func" && p.peekTokenIs(COLON) {
			return p.parseFunctionDefinitionStatement()
		}
		return p.parseFlatTripletStatement()
	case FUNC: // If FUNC is a distinct TokenType (not just IDENT "func")
		return p.parseFunctionDefinitionStatement()
	default:
		p.errors = append(p.errors, fmt.Sprintf("expected start of statement to be SYMSIGN, IDENT or FUNC, got %s instead", p.curToken.Type))
		return nil
	}
}

// parseFlatTripletStatement (kept from previous iteration, if you want to support both)
func (p *Parser) parseFlatTripletStatement() Statement {
	// Expect current token to be IDENT (the subject)
	subjectToken := p.curToken
	subject := &Identifier{Token: subjectToken, Value: subjectToken.Literal}

	// Expect and parse the Verb (next token must be an IDENT or keyword)
	// The verb itself is an IDENT token, but its literal might be a keyword like DO, HAS, IS etc.
	if !p.expectPeek(IDENT) { // Expect IDENT for the verb
		p.errors = append(p.errors, fmt.Sprintf("expected verb (IDENT) after subject, got %s instead", p.peekToken.Type))
		return nil
	}
	verbToken := p.curToken
	verb := &Identifier{Token: verbToken, Value: verbToken.Literal}

	// 3. Expect and parse the Object (can be any Expression type)
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

func (qp *QualifiedProperty) expressionNode()      {}                              // Asegura que implementa Expression
func (qp *QualifiedProperty) TokenLiteral() string { return qp.Key.Token.Literal } // Usa el literal de la clave

// parseQualifiedTripletStatement handles $subject property:value property:value;
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
	// A qualified property starts with an IDENT (key) or LBRACKET (for lists of properties)
	for p.peekTokenIs(IDENT) || p.peekTokenIs(LBRACKET) { // Allow IDENT for key or LBRACKET for list of properties
		if p.peekTokenIs(LBRACKET) {
			// Case: $david [do:"corre", do:"baila"] how:"rapido";
			// This means the 'Object' of the triplet is a ListLiteral of QualifiedProperties.
			// The current logic for Qualifiers assumes a flat list of QualifiedProperty.
			// To handle this, you'd need to parse the list here and then potentially
			// expand it into multiple QualifiedTripletStatements later in the interpreter,
			// or change the AST to allow a ListLiteral directly as a qualifier.
			// For now, let's assume the list of properties is *inside* a single qualifier.
			// Example: $david props:[do:"corre", how:"rapido"];
			// Or, if it's like $david [do:"corre"] how:"rapido";
			// The parser needs to decide if [do:"corre"] is ONE qualifier or a special syntax.
			// Based on your example: $david [do:"corre" do:"baila" do:"vuela"] how:"rapido";
			// This implies the list is a *collection of qualifiers*.

			// Let's adjust this to handle the list of qualified properties.
			// This will be a ListLiteral whose elements are QualifiedProperty expressions.
			p.nextToken() // Consume LBRACKET
			listToken := p.curToken
			listElements := p.parseExpressionList(RBRACKET) // Parse elements until RBRACKET
			listLiteral := &ListLiteral{Token: listToken, Elements: listElements}

			// Now, how to integrate this list into Qualifiers?
			// The current Qualifiers is []*QualifiedProperty.
			// If the list contains QualifiedProperty nodes, we can append them.
			for _, elem := range listLiteral.Elements {
				qp, ok := elem.(*QualifiedProperty)
				if !ok {
					p.errors = append(p.errors, fmt.Sprintf("expected QualifiedProperty inside list, got %T instead", elem))
					return nil
				}
				stmt.Qualifiers = append(stmt.Qualifiers, qp)
			}
			// After parsing the list, continue the loop to check for more qualifiers (like 'how:"rapido"')
			continue // Continue the loop to check for next qualifier
		}

		// Parse a single QualifiedProperty (key:value)
		p.nextToken() // Advance curToken to the key (IDENT)
		keyToken := p.curToken
		key := &Identifier{Token: keyToken, Value: keyToken.Literal}

		if !p.expectPeek(COLON) {
			return nil
		}

		p.nextToken() // Advance curToken to the object's token
		object := p.parseExpression(0)
		if object == nil {
			return nil
		}

		stmt.Qualifiers = append(stmt.Qualifiers, &QualifiedProperty{Key: key, Object: object})
	}

	if !p.expectPeek(SEMICOLON) {
		return nil
	}

	return stmt
}

// parseFunctionDefinitionStatement handles parsing of function definitions.
func (p *Parser) parseFunctionDefinitionStatement() Statement {
	// Current token is FUNC (or IDENT with literal "func")
	funcToken := p.curToken
	stmt := &FunctionDefinitionStatement{Token: funcToken}

	// Expect COLON after FUNC
	if !p.expectPeek(COLON) {
		return nil
	}

	// Expect IDENT for function name
	if !p.expectPeek(IDENT) {
		return nil
	}
	stmt.Name = &Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// Expect COLON again (for func:name: param:[...])
	if !p.expectPeek(COLON) {
		return nil
	}

	// Expect PARAM keyword
	if !p.expectPeek(PARAM) {
		return nil
	}

	// Expect LBRACKET for parameters list
	if !p.expectPeek(LBRACKET) {
		return nil
	}
	stmt.Parameters = p.parseFunctionParameters()

	// Expect RBRACKET for parameters list
	if !p.expectPeek(RBRACKET) {
		return nil
	}

	// Expect CODE keyword
	if !p.expectPeek(CODE) {
		return nil
	}

	// Expect LBRACKET for code block
	if !p.expectPeek(LBRACKET) {
		return nil
	}
	stmt.Body = p.parseFunctionBody()

	// Expect RBRACKET for code block
	if !p.expectPeek(RBRACKET) {
		return nil
	}

	// Expect EXPORT keyword
	if !p.expectPeek(EXPORT) {
		return nil
	}

	// Expect LBRACKET for export list
	if !p.expectPeek(LBRACKET) {
		return nil
	}
	stmt.ExportedValue = p.parseExportedValue()

	// Expect RBRACKET for export list
	if !p.expectPeek(RBRACKET) {
		return nil
	}

	// Expect SEMICOLON at the end of function definition
	if !p.expectPeek(SEMICOLON) {
		return nil
	}

	return stmt
}

// parseFunctionParameters parses the parameter list for a function.
// Format: IDENT HAS IDENT IS IDENT;
func (p *Parser) parseFunctionParameters() []*Parameter {
	var params []*Parameter

	// If the next token is RBRACKET, it's an empty parameter list
	if p.peekTokenIs(RBRACKET) {
		p.nextToken() // Consume RBRACKET
		return params
	}

	p.nextToken() // Advance to the first parameter's IDENT

	for {
		param := &Parameter{}
		param.Token = p.curToken // IDENT token for parameter name
		param.Name = &Identifier{Token: p.curToken, Value: p.curToken.Literal}

		if !p.expectPeek(HAS) {
			return nil
		}
		if !p.expectPeek(IDENT) || p.curToken.Literal != "Value" { // "Value" keyword
			p.errors = append(p.errors, fmt.Sprintf("expected 'Value' keyword, got %s instead", p.curToken.Literal))
			return nil
		}

		if !p.expectPeek(IS) {
			return nil
		}

		// Type of the parameter (int, string, etc.)
		if !p.expectPeek(IDENT) && !p.curTokenIs(INT) && !p.curTokenIs(FLOAT) && !p.curTokenIs(STRING) && !p.curTokenIs(TRUE) && !p.curTokenIs(FALSE) {
			p.errors = append(p.errors, fmt.Sprintf("expected parameter type (IDENT, INT, FLOAT, STRING, BOOLEAN), got %s instead", p.curToken.Type))
			return nil
		}
		param.Type = &Identifier{Token: p.curToken, Value: p.curToken.Literal}

		params = append(params, param)

		if p.peekTokenIs(SEMICOLON) {
			p.nextToken()                // Consume SEMICOLON
			if p.peekTokenIs(RBRACKET) { // Allow trailing semicolon
				break
			}
			p.nextToken() // Advance to the next parameter's IDENT
		} else if p.peekTokenIs(RBRACKET) {
			break // End of parameters list
		} else {
			p.errors = append(p.errors, fmt.Sprintf("expected ';' or ']' after parameter, got %s instead", p.peekToken.Type))
			return nil
		}
	}
	return params
}

// parseFunctionBody parses the code block for a function.
// Currently, it's simplified to just parse assignment statements.
func (p *Parser) parseFunctionBody() []*AssignmentStatement {
	var statements []*AssignmentStatement

	// If the next token is RBRACKET, it's an empty body
	if p.peekTokenIs(RBRACKET) {
		p.nextToken() // Consume RBRACKET
		return statements
	}

	p.nextToken() // Advance to the first statement's IDENT (for assignment)

	for {
		stmt := p.parseAssignmentStatement() // Assuming only assignment statements for now
		if stmt == nil {
			return nil
		}
		statements = append(statements, stmt)

		if p.peekTokenIs(SEMICOLON) {
			p.nextToken()                // Consume SEMICOLON
			if p.peekTokenIs(RBRACKET) { // Allow trailing semicolon
				break
			}
			p.nextToken() // Advance to the next statement's IDENT
		} else if p.peekTokenIs(RBRACKET) {
			break // End of body
		} else {
			p.errors = append(p.errors, fmt.Sprintf("expected ';' or ']' after statement, got %s instead", p.peekToken.Type))
			return nil
		}
	}
	return statements
}

// parseAssignmentStatement parses a simple assignment statement like `variable = expression;`
func (p *Parser) parseAssignmentStatement() *AssignmentStatement {
	stmt := &AssignmentStatement{Token: p.curToken} // Current token is IDENT for variable name

	// Expect IDENT for variable name
	if p.curToken.Type != IDENT {
		p.errors = append(p.errors, fmt.Sprintf("expected IDENT for variable name, got %s instead", p.curToken.Type))
		return nil
	}
	stmt.Name = &Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// Expect ASSIGN operator
	if !p.expectPeek(ASSIGN) {
		return nil
	}

	// Parse the value/expression being assigned
	p.nextToken() // Advance to the expression's token
	stmt.Value = p.parseExpression(0)
	if stmt.Value == nil {
		return nil
	}

	return stmt
}

// parseExportedValue parses the exported value list for a function.
func (p *Parser) parseExportedValue() Expression {
	// If the next token is RBRACKET, it's an empty export list
	if p.peekTokenIs(RBRACKET) {
		p.nextToken() // Consume RBRACKET
		return nil    // Or an empty list literal, depending on desired AST
	}

	p.nextToken() // Advance to the first exported value

	// For simplicity, assuming only one exported value for now, which can be any expression
	exportedExp := p.parseExpression(0)

	// No need to loop for multiple exports if only one is allowed
	// if p.peekTokenIs(COMMA) { ... }

	return exportedExp
}
