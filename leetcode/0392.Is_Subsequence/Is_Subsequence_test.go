package Is_Subsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubsequence(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(true, isSubsequence("abc", "asdbsldc"))
	ast.Equal(true, isSubsequence("", "abc"))
	ast.Equal(false, isSubsequence("axc", "ahbgdc"))
}
