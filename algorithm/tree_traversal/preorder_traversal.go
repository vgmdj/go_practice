package tree_traversal

import "container/list"

func PreOrder(root *Tree) []int {
	var result []int
	if root == nil {
		return result
	}

	result = append(result, root.Val)
	result = append(result, PreOrder(root.Left)...)
	result = append(result, PreOrder(root.Right)...)

	return result

}

func PreOrderNonRecursive(root *Tree) []int {
	var result []int
	if root == nil {
		return result
	}

	stack := list.New()
	p := root

	for p != nil || stack.Len() != 0 {
		for p != nil {
			result = append(result, p.Val)
			stack.PushBack(p)
			p = p.Left
		}

		if stack.Len() != 0 {
			node := stack.Back()
			stack.Remove(node)

			p = node.Value.(*Tree).Right
		}
	}

	return result
}
