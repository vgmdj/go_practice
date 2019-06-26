package Construct_Binary_Tree_from_Inorder_and_Postorder_Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	position := make(map[int]int, len(postorder))
	for k, v := range inorder {
		position[v] = k
	}

	index := len(postorder) - 1

	return buildTreeHelper(postorder, &index, 0, len(postorder)-1, position)

}

func buildTreeHelper(postorder []int, index *int, start, end int, position map[int]int) *TreeNode {
	if start > end {
		return nil
	}

	root := &TreeNode{Val: postorder[*index]}
	p := position[root.Val]
	*index = *index - 1

	root.Right = buildTreeHelper(postorder, index, p+1, end, position)
	root.Left = buildTreeHelper(postorder, index, start, p-1, position)

	return root

}
