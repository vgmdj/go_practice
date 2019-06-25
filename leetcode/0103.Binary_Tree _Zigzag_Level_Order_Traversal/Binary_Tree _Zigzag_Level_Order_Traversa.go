package Binary_Tree__Zigzag_Level_Order_Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder2(root *TreeNode) [][]int {
	result := make([][]int, 0)

	dfsHelper(root, &result, 0)

	for i := 1; i < len(result); i += 2 {
		reverse(result[i])
	}

	return result

}

func dfsHelper(root *TreeNode, result *[][]int, level int) {
	if root == nil {
		return
	}

	if len(*result) <= level {
		*result = append(*result, []int{root.Val})

	} else {
		(*result)[level] = append((*result)[level], root.Val)

	}

	dfsHelper(root.Left, result, level+1)
	dfsHelper(root.Right, result, level+1)

}

func reverse(nums []int) {
	head, tail := 0, len(nums)-1

	for head < tail {
		nums[head], nums[tail] = nums[tail], nums[head]
		head++
		tail--

	}

}
