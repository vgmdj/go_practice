package sort

//插入排序
func insertSort(nums []int) {
	for i := 1; i < len(nums)-1; i++ {
		for j := i; j > 0 && nums[j] < nums[j-1]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}

}
