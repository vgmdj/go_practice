package main

import "log"

// 归并排序
//先递归分治
func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	nums1 := mergeSort(nums[:mid])
	nums2 := mergeSort(nums[mid:])

	return merge(nums1, nums2)
}

//每个分组里再按从小到大的顺序组成nums3
//master公式 T(N) = aT(n/b)+O(n^d)
//T(N)=2T(n/2)+(n)
//a=b=2 d=1
func merge(nums1 []int, nums2 []int) []int {
	var nums3 []int
	var position1, position2 int

	for position1 < len(nums1) && position2 < len(nums2) {
		if nums1[position1] < nums2[position2] {
			nums3 = append(nums3, nums1[position1])
			position1++
		} else {
			nums3 = append(nums3, nums2[position2])
			position2++
		}
	}

	nums3 = append(nums3, nums1[position1:]...)
	nums3 = append(nums3, nums2[position2:]...)

	return nums3
}

func main() {
	nums := []int{1, 3, 5, 7, 9, 2, 6, 0, 12}

	nums = mergeSort(nums)
	log.Println(nums)

}
