package lexer_test

import (
	"testing"

	"github.com/janmarkuslanger/julang/lexer"
	"github.com/janmarkuslanger/julang/token"
)

func TestNewLexerInitializesCorrectly(t *testing.T) {
	input := "let"
	l := lexer.New(input)

	if l.Input() != input {
		t.Errorf("expected input %q, got %q", input, l.Input())
	}

	if l.CurrentChar() != 'l' {
		t.Errorf("expected first character to be 'l', got %q", l.CurrentChar())
	}
}

func TestNextTokenReturnsCorrectTokens(t *testing.T) {
	input := "let x = 5;"

	expected := []token.Token{
		{Type: token.LET, Literal: "let"},
		{Type: token.ID, Literal: "x"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.INT, Literal: "5"},
		{Type: token.SEMI, Literal: ";"},
		{Type: token.EOF, Literal: ""},
	}

	l := lexer.New(input)

	for i, exp := range expected {
		tok := l.NextToken()

		if tok.Type != exp.Type {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q", i, exp.Type, tok.Type)
		}
		if tok.Literal != exp.Literal {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, exp.Literal, tok.Literal)
		}
	}
}
