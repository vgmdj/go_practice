package Subsets

func subsets(nums []int) [][]int {
	var result [][]int
	var visited = make([]bool, len(nums))
	backTrack(&result, nums, []int{}, 0, visited)
	return result

}

func backTrack(result *[][]int, nums, subset []int, start int, visited []bool) {
	*result = append(*result, append([]int{}, subset[:]...))

	for i := start; i < len(nums); i++ {
		if visited[i] {
			continue
		}

		visited[i] = true
		backTrack(result, nums, append(subset, nums[i]), i+1, visited)
		visited[i] = false
	}

}
