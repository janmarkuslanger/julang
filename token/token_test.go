package token_test

import (
	"testing"

	"github.com/janmarkuslanger/julang/token"
	"github.com/stretchr/testify/assert"
)

func TestLookupIdent(t *testing.T) {
	assert.Equal(t, token.LookupIdent("let"), token.TokenType(token.LET))
	assert.Equal(t, token.LookupIdent("letsgo"), token.TokenType(token.ID))
	assert.Equal(t, token.LookupIdent("const"), token.TokenType(token.CONST))
}
