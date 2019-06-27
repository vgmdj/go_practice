package Palindrome_Partitioning

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	ast := assert.New(t)

	ast.Equal([][]string{
		{"b","b"},
		{"bb"},
	},partition("bb"))

	ast.Equal([][]string{
		{"a", "a", "b"},
		{"aa", "b"},
	}, partition("aab"))
}
