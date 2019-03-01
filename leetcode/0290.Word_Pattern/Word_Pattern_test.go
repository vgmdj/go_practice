package Word_Pattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordPattern(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(true, wordPattern("abba", "dog cat cat dog"))
	ast.Equal(false, wordPattern("abba", "dog cat cat fish"))
	ast.Equal(false, wordPattern("aaaa", "dog cat fish dog"))
	ast.Equal(false, wordPattern("abba", "dog dog dog dog"))
}
