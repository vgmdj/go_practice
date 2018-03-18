package main

import "log"

func insertSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	for i := 1; i < len(nums)-1; i++ {
		for j := i; j > 0 && nums[j] < nums[j-1]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}

}

func main() {
	nums := []int{1, 3, 5, 7, 9, 2, 6, 0, 12}

	insertSort(nums)
	log.Println(nums)

}
