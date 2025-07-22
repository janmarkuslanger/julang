package utils_test

import (
	"testing"

	"github.com/janmarkuslanger/julang/utils"
	"github.com/stretchr/testify/assert"
)

func TestIsDigit(t *testing.T) {
	assert.Equal(t, utils.IsDigit('a'), false)
	assert.Equal(t, utils.IsDigit('A'), false)
	assert.Equal(t, utils.IsDigit('N'), false)
	assert.Equal(t, utils.IsDigit('Z'), false)
	assert.Equal(t, utils.IsDigit('0'), true)
	assert.Equal(t, utils.IsDigit('1'), true)
	assert.Equal(t, utils.IsDigit('9'), true)
}
