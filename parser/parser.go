package parser

import (
	"strconv"

	"github.com/janmarkuslanger/julang/ast"
	"github.com/janmarkuslanger/julang/lexer"
	"github.com/janmarkuslanger/julang/token"
)

type Parser struct {
	l       *lexer.Lexer
	tok     token.Token
	peekTok token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.tok = p.peekTok
	p.peekTok = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.tok.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTok.Type == t {
		p.nextToken()
		return true
	}
	return false
}

func (p *Parser) parseExpression() ast.Expression {
	switch p.tok.Type {
	case token.INT:
		return p.parseIntegerLiteral()
	default:
		return nil
	}
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.tok.Type {
	case token.VAR:
		return p.parseVarStatement()
	default:
		return nil
	}
}

func (p *Parser) parseVarStatement() *ast.VarStatement {
	stmt := &ast.VarStatement{Token: p.tok}

	if !p.expectPeek(token.ID) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: p.tok,
		Value: p.tok.Literal,
	}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()
	stmt.Value = p.parseExpression()

	if !p.expectPeek(token.SEMI) {
		return nil
	}

	return stmt
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.tok}

	value, err := strconv.ParseInt(p.tok.Literal, 0, 64)
	if err != nil {
		return nil
	}

	lit.Value = value
	return lit
}
