package Third_Maximum_Number

//大根堆排
func thirdMax2(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	for i := len(nums)/2 - 1; i >= 0; i-- {
		maxHeapify(nums, i, len(nums)-1)
	}

	i, k := len(nums)-1, 3
	for ; i >= 0 && k > 1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		maxHeapify(nums, 0, i-1)

		if i != len(nums)-1 && nums[i] != nums[i+1] {
			k--
		}
	}

	if k > 1 {
		return nums[len(nums)-1]
	}

	return nums[i+1]
}

func maxHeapify(nums []int, start, end int) {
	var father = start
	var child = 2*father + 1

	for child <= end {
		if child+1 <= end && nums[child] < nums[child+1] {
			child++
		}

		if nums[father] > nums[child] {
			return
		}

		nums[father], nums[child] = nums[child], nums[father]
		father = child
		child = 2*father + 1

	}
}
