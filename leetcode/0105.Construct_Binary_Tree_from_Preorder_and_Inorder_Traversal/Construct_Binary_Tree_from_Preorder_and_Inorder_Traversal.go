package Construct_Binary_Tree_from_Preorder_and_Inorder_Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	position := make(map[int]int, len(inorder))
	for k, v := range inorder {
		position[v] = k
	}

	index := 0

	return buildTreeHelper(preorder, &index, 0, len(preorder)-1, position)

}

func buildTreeHelper(preorder []int, index *int, start, end int, position map[int]int) *TreeNode {
	if start > end {
		return nil
	}

	root := &TreeNode{Val: preorder[*index]}
	p := position[root.Val]
	*index = *index + 1

	root.Left = buildTreeHelper(preorder, index, start, p-1, position)
	root.Right = buildTreeHelper(preorder, index, p+1, end, position)

	return root

}
