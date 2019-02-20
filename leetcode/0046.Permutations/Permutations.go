package Permutations

func permute(nums []int) [][]int {
	var result [][]int
	var visited = make([]bool, len(nums))
	backTrack(&result, nums, []int{}, visited)
	return result

}

func backTrack(result *[][]int, nums, subset []int, visited []bool) {
	if len(subset) == len(nums) {
		*result = append(*result, append([]int{}, subset[:]...))
		return
	}

	for k := range nums {
		if visited[k] {
			continue
		}

		visited[k] = true
		backTrack(result, nums, append(subset, nums[k]), visited)
		visited[k] = false

	}

}
