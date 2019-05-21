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

func LevelTraversal2(root *Tree) [][]int {
	var result [][]int

	if root == nil {
		return result
	}

	dfsHelper(&result, root, 0)

	return result

}

func dfsHelper(result *[][]int, root *Tree, level int) {
	if root == nil {
		return
	}

	if len(*result) < level+1 {
		*result = append(*result, []int{root.Val})

	} else {
		(*result)[level] = append((*result)[level], root.Val)

	}

	dfsHelper(result, root.Left, level+1)
	dfsHelper(result, root.Right, level+1)

}
