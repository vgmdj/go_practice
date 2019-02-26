package _090_Subsets_II

import "sort"

func subsetsWithDup(nums []int) [][]int {
	var result [][]int
	var visited = make([]bool, len(nums))
	sort.Ints(nums)
	backTrack(&result, nums, []int{}, 0, visited)
	return result

}

func backTrack(result *[][]int, nums, subset []int, start int, visited []bool) {
	*result = append(*result, append([]int{}, subset[:]...))

	for i := start; i < len(nums); i++ {
		if visited[i] || (i > 0 && !visited[i-1] && nums[i] == nums[i-1]) {
			continue
		}

		visited[i] = true
		backTrack(result, nums, append(subset, nums[i]), i+1, visited)
		visited[i] = false
	}

}
