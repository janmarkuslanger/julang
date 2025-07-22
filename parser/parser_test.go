package parser_test

import (
	"testing"

	"github.com/janmarkuslanger/julang/ast"
	"github.com/janmarkuslanger/julang/lexer"
	"github.com/janmarkuslanger/julang/parser"
)

func TestLetStatementExpressionValue(t *testing.T) {
	input := "let x = 5;"

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(program.Statements) != 1 {
		t.Fatalf("expected 1 statement, got %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.LetStatement)
	if !ok {
		t.Fatalf("statement not LetStatement, got %T", program.Statements[0])
	}

	lit, ok := stmt.Value.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("value not IntegerLiteral, got %T", stmt.Value)
	}

	if lit.Value != 5 {
		t.Errorf("expected value to be 5, got %d", lit.Value)
	}
}
