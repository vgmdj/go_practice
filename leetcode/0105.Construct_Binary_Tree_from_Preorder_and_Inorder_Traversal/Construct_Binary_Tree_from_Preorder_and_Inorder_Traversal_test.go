package Construct_Binary_Tree_from_Preorder_and_Inorder_Traversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstruct(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(nil, buildTree([]int{1, 2, 3, 4}, []int{1, 2, 3, 4}))

	ast.Equal(nil, buildTree(
		[]int{1, 2, 4, 8, 9, 5, 3, 6, 7},
		[]int{8, 4, 9, 2, 5, 1, 6, 3, 7}))
}
