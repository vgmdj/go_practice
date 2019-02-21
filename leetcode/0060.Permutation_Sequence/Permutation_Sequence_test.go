package Permutation_Sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
123
132
213
231
312
321
*/
func TestGetPermutation(t *testing.T) {
	ast := assert.New(t)

	ast.Equal("2314",getPermutation(4,9))
	ast.Equal("213", getPermutation(3, 3))
	ast.Equal("231", getPermutation(3, 4))

}
