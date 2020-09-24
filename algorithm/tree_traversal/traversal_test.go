package tree_traversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeTraversal(t *testing.T) {
	ast := assert.New(t)

	//pre-order traversal
	preResult := []int{1, 2, 4, 6, 5, 7, 3}
	ast.Equal(preResult, PreOrder(TestTree()))
	ast.Equal(preResult, PreOrderNonRecursive(TestTree()))

	//in-order traversal
	inResult := []int{4, 6, 2, 7, 5, 1, 3}
	ast.Equal(inResult, InOrder(TestTree()))
	ast.Equal(inResult, InOrderNonRecursive(TestTree()))
	ast.Equal(inResult, InOrderTravel(TestTree()))

	//post-order traversal
	postResult := []int{6, 4, 7, 5, 2, 3, 1}
	ast.Equal(postResult, PostOrder(TestTree()))
	ast.Equal(postResult, PostOrderNonRecursive(TestTree()))

	//BFS
	bfsResult := []int{1, 2, 3, 4, 5, 6, 7}
	ast.Equal(bfsResult, LevelTraversal(TestTree()))

	//use dfs pre-order traversal and return the bfs []int result
	bfsResultWithNull := []int{1, 2, 3, 4, 5, 0, 0, 0, 6, 7}
	ast.Equal(bfsResultWithNull, LevelTraversalWithNull(TestTree()))

	// use dfs pre-order traversal and return the bfs [][]int result without null
	bfsResultWithoutNull := [][]int{
		{1},
		{2, 3},
		{4, 5},
		{6, 7},
	}
	ast.Equal(bfsResultWithoutNull, LevelTraversalWithoutNull(TestTree()))

}
