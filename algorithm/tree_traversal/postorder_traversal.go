package tree_traversal

import "container/list"

func PostOrder(root *Tree) []int {
	var result []int
	if root == nil {
		return nil
	}

	result = append(result, PreOrder(root.Left)...)
	result = append(result, PreOrder(root.Right)...)
	result = append(result, root.Val)

	return result

}

func PostOrderNonRecursive(root *Tree) []int {
	var result []int
	if root == nil {
		return result
	}

	stack := list.New()
	p := root
	lastVisit := new(Tree)

	for p != nil || stack.Len() != 0 {
		for p != nil {
			stack.PushBack(p)
			p = p.Left

		}

		if stack.Len() != 0 {
			node := stack.Back()
			if node.Value.(*Tree).Right != nil && lastVisit != node.Value.(*Tree).Right {
				p = node.Value.(*Tree).Right
				continue
			}

			result = append(result, node.Value.(*Tree).Val)
			stack.Remove(node)
			lastVisit = node.Value.(*Tree)

		}

	}

	return result
}
