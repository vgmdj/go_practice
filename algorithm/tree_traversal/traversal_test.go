package tree_traversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeTraversal(t *testing.T) {
	ast := assert.New(t)

	//pre-order traversal
	preResult := []int{1, 2, 4, 6, 5, 7, 3}
	testArray(ast, preResult, PreOrder(TestTree()))
	testArray(ast, preResult, PreOrderNonRecursive(TestTree()))

	//in-order traversal
	inResult := []int{4, 6, 2, 7, 5, 1, 3}
	testArray(ast, inResult, InOrder(TestTree()))
	testArray(ast, inResult, InOrderNonRecursive(TestTree()))

	//post-order traversal
	postResult := []int{6, 4, 7, 5, 2, 3, 1}
	testArray(ast, postResult, PostOrder(TestTree()))
	testArray(ast, postResult, PostOrderNonRecursive(TestTree()))

	//BFS
	bfsResult := []int{1, 2, 3, 4, 5, 6, 7}
	testArray(ast, bfsResult, LevelTraversal(TestTree()))

	//dfs pre-order traversal but return the bfs [][]int result
	bfsResults := [][]int{
		{1},
		{2, 3},
		{4, 5},
		{6, 7},
	}
	testArrays(ast, bfsResults, LevelTraversal2(TestTree()))

}

func testArray(ast *assert.Assertions, array1 []int, array2 []int) {
	if !ast.Equal(len(array1), len(array2)) {
		return
	}

	for i := 0; i < len(array1); i++ {
		ast.Equal(array1[i], array2[i])
	}

}

func testArrays(ast *assert.Assertions, array1 [][]int, array2 [][]int) {
	if !ast.Equal(len(array1), len(array2)) {
		return
	}

	for i := 0; i < len(array1); i++ {
		testArray(ast, array1[i], array2[i])
	}

}
