package lexer

import (
	"github.com/janmarkuslanger/julang/token"
	"github.com/janmarkuslanger/julang/utils"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) readIdentifier() string {
	start := l.position
	for utils.IsLetter(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	start := l.position
	for utils.IsDigit(l.ch) {
		l.readChar()
	}

	return l.input[start:l.position]
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = token.New(token.ASSIGN, l.ch)
	case '+':
		tok = token.New(token.PLUS, l.ch)
	case ';':
		tok = token.New(token.SEMI, l.ch)
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}
	default:
		if utils.IsLetter(l.ch) {
			literal := l.readIdentifier()
			tok.Type = token.LookupIdent(literal)
			tok.Literal = literal
			return tok
		}

		if utils.IsDigit(l.ch) {
			literal := l.readNumber()
			tok = token.Token{Type: token.INT, Literal: literal}
			return tok
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) Input() string {
	return l.input
}

func (l *Lexer) CurrentChar() byte {
	return l.ch
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}
