package utils_test

import (
	"testing"

	"github.com/janmarkuslanger/julang/utils"
	"github.com/stretchr/testify/assert"
)

func TestIsLetter(t *testing.T) {
	assert.Equal(t, utils.IsLetter('c'), true)
	assert.Equal(t, utils.IsLetter('a'), true)
	assert.Equal(t, utils.IsLetter(1), false)
	assert.Equal(t, utils.IsLetter(' '), false)
}
