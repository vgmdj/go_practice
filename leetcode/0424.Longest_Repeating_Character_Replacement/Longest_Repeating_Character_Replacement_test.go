package Longest_Repeating_Character_Replacement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplacement(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(7,characterReplacement("AABABABBABBCAA", 2))
	ast.Equal(4,characterReplacement("ABAABB", 1))

}
