package Search_in_a_Binary_Search_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if val == root.Val {
		return root
	}

	if val < root.Val {
		return searchBST(root.Left, val)
	}

	return searchBST(root.Right, val)

}
