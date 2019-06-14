package Kth_Largest_Element_in_an_Array

import (
	"math"
	"sort"
)

//直接排序
func findKthLargest1(nums []int, k int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	return nums[k-1]

}

//冒泡排序
func findKthLargest2(nums []int, k int) int {

	for i := 0; i < len(nums)-1 && i < k; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				nums[j], nums[i] = nums[i], nums[j]
			}
		}

	}

	return nums[k-1]

}

//快排
func findKthLargest3(nums []int, k int) int {
	if len(nums) < 1 {
		return math.MinInt32
	}

	head, tail := 0, len(nums)-1
	mid, next := nums[0], 1

	for head < tail {
		if nums[next] > mid {
			nums[next], nums[tail] = nums[tail], nums[next]
			tail--

		} else {
			nums[next], nums[head] = nums[head], nums[next]
			head++
			next++

		}

	}

	if head < len(nums)-k {
		return findKthLargest3(nums[next:], k)
	}

	if head == len(nums)-k {
		return nums[next-1]
	}

	return findKthLargest3(nums[:next-1], k+head-len(nums))

}

//大根堆排
func findKthLargest(nums []int, k int) int {

	heapSort(nums, k)

	return nums[len(nums)-k]
}

func heapSort(nums []int, k int) {

	for i := len(nums)/2 - 1; i >= 0; i-- {
		maxHeapify(nums, i, len(nums)-1)
	}

	for i := len(nums) - 1; i >= len(nums)-k; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		maxHeapify(nums, 0, i-1)
	}
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
