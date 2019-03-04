package Reverse_Vowels_of_a_String

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	ast := assert.New(t)

	ast.Equal("holle", reverseVowels("hello"))
	ast.Equal("sos", reverseVowels("sos"))
	ast.Equal("leotcede", reverseVowels("leetcode"))
}
