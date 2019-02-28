package Isomorphic_Strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsOmorphic(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(false, isIsomorphic("ab", "aa"))
	ast.Equal(true, isIsomorphic("add", "egg"))
	ast.Equal(true, isIsomorphic("title", "paper"))
	ast.Equal(false, isIsomorphic("foo", "bar"))
}
