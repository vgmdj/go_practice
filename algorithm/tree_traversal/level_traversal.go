package tree_traversal

import "container/list"

func LevelTraversal(root *Tree) []int {
	var result []int
	if root == nil {
		return result
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() != 0 {
		node := queue.Front()
		nTree := node.Value.(*Tree)
		queue.Remove(node)
		if nTree == nil {
			continue
		}

		result = append(result, nTree.Val)
		queue.PushBack(nTree.Left)
		queue.PushBack(nTree.Right)

	}

	return result

}

func LevelTraversalWithoutNull(root *Tree) [][]int {
	result := make([][]int, 0)

	if root == nil {
		return result
	}

	dfsWithoutNullHelper(&result, root, 0)

	return result

}

func dfsWithoutNullHelper(result *[][]int, root *Tree, index int) {
	if root == nil {
		return
	}

	if len(*result) < index+1 {
		*result = append(*result, make([]int, 0))
	}

	(*result)[index] = append((*result)[index], root.Val)

	dfsWithoutNullHelper(result, root.Left, index+1)
	dfsWithoutNullHelper(result, root.Right, index+1)

}

func LevelTraversalWithNull(root *Tree) []int {
	result := make([]int, 0)

	if root == nil {
		return result
	}

	dfsWithNullHelper(&result, root, 0)

	return result

}

func dfsWithNullHelper(result *[]int, root *Tree, index int) {
	if root == nil {
		return
	}

	if len(*result) < index+1 {
		*result = append(*result, make([]int, index+1-len(*result))...)
	}

	(*result)[index] = root.Val

	dfsWithNullHelper(result, root.Left, index*2+1)
	dfsWithNullHelper(result, root.Right, index*2+2)

}
