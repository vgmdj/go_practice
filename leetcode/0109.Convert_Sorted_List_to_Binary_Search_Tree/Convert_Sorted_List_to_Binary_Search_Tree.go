package Convert_Sorted_List_to_Binary_Search_Tree

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func sortedListToBST(head *ListNode) *TreeNode {
	nums := convertListToArray(head)

	return sortedArrayToBST(0, len(nums)-1, nums)
}

func convertListToArray(head *ListNode) []int {
	var result []int

	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}

	return result

}

func sortedArrayToBST(left, right int, nums []int) *TreeNode {
	if right-left < 0 || left < 0 {
		return nil
	}

	mid := (left + right) / 2
	tree := &TreeNode{
		Val: nums[mid],
	}

	tree.Left = sortedArrayToBST(left, mid-1, nums)
	tree.Right = sortedArrayToBST(mid+1, right, nums)

	return tree

}
