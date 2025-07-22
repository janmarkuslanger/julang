package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ID     = "ID"
	INT    = "INT"
	STRING = "STRING"
	ASSIGN = "="
	PLUS   = "+"
	SEMI   = ";"
	VAR    = "var"
	EOF    = ""
)

var keywords = map[string]TokenType{
	"var": VAR,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return ID
}

func New(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}
