package sort

func selectSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[minIndex] > nums[j] {
				minIndex = j
			}
		}

		nums[minIndex], nums[i] = nums[i], nums[minIndex]
	}

}
