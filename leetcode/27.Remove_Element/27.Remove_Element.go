package Remove_Element

func removeElement(nums []int, val int) int {
	var head, tail int
	for tail = len(nums) - 1; head <= tail; {
		if nums[head] == val {
			nums[head], nums[tail] = nums[tail], nums[head]
			tail--
		} else {
			head++
		}

	}

	return head
}
