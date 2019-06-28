package Kth_Smallest_Element_in_a_BST

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
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

			k--
			if k == 0 {
				return node.Value.(*TreeNode).Val
			}

			p = node.Value.(*TreeNode).Right

		}

	}

	return -1

}

func kthSmallest1(root *TreeNode, k int) int {
	return helper1(root, &k)
}

func helper1(root *TreeNode, k *int) int {
	if root == nil {
		return -1
	}

	result := helper1(root.Left, k)
	if *k == 0 {
		return result
	}

	*k--
	if *k == 0 {
		return root.Val
	}

	return helper1(root.Right, k)

}

func kthSmallest2(root *TreeNode, k int) int {
	result := make([]int, k)
	helper2(root, k, &result)
	return result[k-1]
}

func helper2(root *TreeNode, k int, result *[]int) {
	if root == nil || len(*result) == k {
		return
	}

	helper2(root.Left, k, result)

	*result = append(*result, root.Val)

	helper2(root.Right, k, result)

}
