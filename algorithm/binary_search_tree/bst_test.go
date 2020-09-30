package binary_search_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBstSearch(t *testing.T) {
	ast := assert.New(t)

	bst := testBstTree()
	res1 := bst.Search(63)
	ast.Equal([]int{63, 0, 70}, res1.ToArray())

	res2 := bst.Search(22)
	ast.Equal([]int{}, res2.ToArray())

	res3 := bst.Search(25)
	ast.Equal([]int{25, 12, 37, 6, 20, 30}, res3.ToArray())

	ast.Equal(true, bst.IsBst())

}

func TestIsBst(t *testing.T) {
	ast := assert.New(t)
	bst := testBstTree()
	bst.left.data = 51

	ast.Equal(false, bst.IsBst())

}

func TestBstAddNode(t *testing.T) {
	ast := assert.New(t)

	bst := NewBinarySearchTree(50)
	bst.AddNode(25)
	bst.AddNode(75)
	bst.AddNode(12)
	bst.AddNode(37)
	bst.AddNode(63)
	bst.AddNode(100)
	bst.AddNode(6)
	bst.AddNode(20)
	bst.AddNode(30)
	bst.AddNode(70)

	ast.Equal(testBstTree().ToArray(), bst.ToArray())

	ast.Equal(true, bst.IsBst())

}

func TestBstRemove(t *testing.T) {
	ast := assert.New(t)
	bst := testBstTree()
	bst.RemoveNode(63)
	ast.Equal([]int{}, bst.ToArray())

	bst.RemoveNode(48)
	bst.RemoveNode(25)
	ast.Equal([]int{50, 30, 75, 12, 37, 63, 100, 6, 20, 0, 0, 0, 70}, bst.ToArray())
	ast.Equal(true, bst.IsBst())

	bst.RemoveNode(50)
	ast.Equal([]int{63, 30, 70, 12, 37, 0, 75, 6, 20, 0, 0, 0, 0, 0, 100}, bst.ToArray())
	ast.Equal(true, bst.IsBst())

}

func TestMinNode(t *testing.T) {
	ast := assert.New(t)
	bst := testBstTree()
	min := bst.MinNode()
	ast.Equal(6, min.data)

}

func TestMaxNode(t *testing.T) {
	ast := assert.New(t)
	bst := testBstTree()
	max := bst.MaxNode()
	ast.Equal(100, max.data)
}

/*
					50
		25							75
	12        37				63				100
6	   20	30                      70

*/

func testBstTree() *BinarySearchTree {
	bst := NewBinarySearchTree(50)
	addNewNode(bst, 25, left)
	addNewNode(bst, 75, right)
	addNewNode(bst.left, 12, left)
	addNewNode(bst.left, 37, right)
	addNewNode(bst.right, 63, left)
	addNewNode(bst.right, 100, right)
	addNewNode(bst.left.left, 6, left)
	addNewNode(bst.left.left, 20, right)
	//addNewNode(bst.left.right, 30, left)
	addNewNode(bst.right.left, 70, right)
	return bst
}

type position int

const (
	left  position = 1
	right position = 2
)

func addNewNode(parent *BinarySearchTree, val int, p position) {
	node := &BinarySearchTree{
		data:  val,
		left:  nil,
		right: nil,
	}

	if p == left {
		parent.left = node
		return
	}

	if p == right {
		parent.right = node
		return
	}

}
