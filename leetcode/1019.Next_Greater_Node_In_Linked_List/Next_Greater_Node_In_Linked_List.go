package Next_Greater_Node_In_Linked_List

type ListNode struct {
	Val  int
	Next *ListNode
}


func nextLargerNodes(head *ListNode) []int {

	values := make([]int, 0)
	cur := head
	for cur != nil {
		values = append(values, cur.Val)
		cur = cur.Next
	}

	result := make([]int, len(values))
	stack := make([]int, 0)

	for i := 0; i < len(values); {
		if len(stack) == 0 || values[i] <= values[stack[len(stack)-1]] {
			stack = append(stack, i)
			i++

		} else {
			result[stack[len(stack)-1]] = values[i]
			stack = stack[:len(stack)-1]

		}

	}

	return result

}