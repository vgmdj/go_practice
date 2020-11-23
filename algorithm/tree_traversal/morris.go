package tree_traversal

// time 2N
// space 1
func morris(root *Tree) []int {
	result := make([]int, 0)
	var predecessor *Tree
	for root != nil {
		if root.Left != nil {
			predecessor = root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				predecessor = predecessor.Right
			}

			if predecessor.Right == nil {
				predecessor.Right = root
				root = root.Left

			} else {
				predecessor.Right = nil
				result = append(result, root.Val)
				root = root.Right

			}

		} else {
			result = append(result, root.Val)
			root = root.Right

		}

	}

	return result
}
