package token_test

import (
	"testing"

	"github.com/janmarkuslanger/julang/token"
	"github.com/stretchr/testify/assert"
)

func TestLookupIdent(t *testing.T) {
	assert.Equal(t, token.LookupIdent("var"), token.TokenType(token.VAR))
	assert.Equal(t, token.LookupIdent("letsgo"), token.TokenType(token.ID))
	assert.Equal(t, token.LookupIdent("const"), token.TokenType(token.ID))
}
