package binary_search_tree

import "testing"

func TestAVLKeepBalance(t *testing.T) {
	avl := NewAVLHelper()
	avl.AddNode(1)
	avl.AddNode(2)
	avl.AddNode(3)
	avl.AddNode(4)
	avl.AddNode(5)
	t.Log(levelTravel(avl.GetAVLTree()))

}

func levelTravel(root *AVLTree) []int {
	result := make([]int, 0)

	if root == nil {
		return result
	}

	testHelper(&result, root, 0)

	return result
}

func testHelper(result *[]int, root *AVLTree, index int) {
	if root == nil {
		return
	}

	if len(*result) < index+1 {
		*result = append(*result, make([]int, index+1-len(*result))...)
	}

	(*result)[index] = root.data

	testHelper(result, root.left, index*2+1)
	testHelper(result, root.right, index*2+2)

}
