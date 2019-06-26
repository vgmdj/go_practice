package Construct_Binary_Tree_from_Inorder_and_Postorder_Traversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstruct(t *testing.T) {
	ast := assert.New(t)

	//ast.Equal(nil, buildTree([]int{1, 2, 3, 4}, []int{1, 2, 3, 4}))

	ast.Equal(nil, buildTree(
		[]int{9,3,15,20,7},
		[]int{9,15,7,20,3}))
}
