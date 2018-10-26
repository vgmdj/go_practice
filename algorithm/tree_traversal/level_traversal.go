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
