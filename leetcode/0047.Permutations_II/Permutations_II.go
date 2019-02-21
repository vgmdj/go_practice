package Permutations

import "sort"

func permuteUnique(nums []int) [][]int {
	var result [][]int
	var visited = make([]bool, len(nums))
	sort.Ints(nums)
	backTrack(&result, nums, []int{}, visited)
	return result

}

//如果不连续,且与上一数相同,则说明已处理过这种情况
func backTrack(result *[][]int, nums, subset []int, visited []bool) {
	if len(subset) == len(nums) {
		*result = append(*result, append([]int{}, subset[:]...))
		return
	}

	for k := range nums {
		if visited[k] || (k > 0 && nums[k] == nums[k-1] && !visited[k-1]) {
			continue
		}

		visited[k] = true
		backTrack(result, nums, append(subset, nums[k]), visited)
		visited[k] = false

	}

}
