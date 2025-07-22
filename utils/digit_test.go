package utils_test

import (
	"testing"

	"github.com/janmarkuslanger/julang/utils"
	"github.com/stretchr/testify/assert"
)

func TestIsDigit(t *testing.T) {
	cases := []struct {
		input    byte
		expected bool
		name     string
	}{
		{'a', false, "lowercase letter"},
		{'A', false, "uppercase letter"},
		{'N', false, "middle uppercase letter"},
		{'Z', false, "last uppercase letter"},
		{'0', true, "first digit"},
		{'1', true, "middle digit"},
		{'9', true, "last digit"},
		{'/', false, "character before '0'"},
		{':', false, "character after '9'"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, utils.IsDigit(tc.input), tc.expected)
		})
	}
}
