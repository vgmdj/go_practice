package Unique_Binary_Search_Trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUniqueBinarySearchTrees(t *testing.T){
	ast := assert.New(t)

	ast.Equal(1,numTrees(1))
	ast.Equal(2,numTrees(2))
	ast.Equal(5,numTrees(3))
	ast.Equal(14,numTrees(4))
}
