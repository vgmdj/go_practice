package tree_traversal

import "container/list"

func InOrder(root *Tree) []int {
	var result []int
	if root == nil {
		return nil
	}

	result = append(result, InOrder(root.Left)...)
	result = append(result, root.Val)
	result = append(result, InOrder(root.Right)...)

	return result

}

func InOrderTravel(root *Tree) []int {
	result := make([]int, 0)
	InOrderHelper(root, &result)
	return result
}

func InOrderHelper(root *Tree, result *[]int) {
	if root == nil {
		return
	}

	InOrderHelper(root.Left, result)
	*result = append(*result, root.Val)
	InOrderHelper(root.Right, result)

}

func InOrderNonRecursive(root *Tree) []int {
	var result []int
	if root == nil {
		return result
	}

	stack := list.New()
	p := root

	for p != nil || stack.Len() != 0 {
		for p != nil {
			stack.PushBack(p)
			p = p.Left

		}

		if stack.Len() != 0 {
			node := stack.Back()
			stack.Remove(node)

			result = append(result, node.Value.(*Tree).Val)

			p = node.Value.(*Tree).Right

		}

	}

	return result
}
