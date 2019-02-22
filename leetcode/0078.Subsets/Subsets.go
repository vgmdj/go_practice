package Subsets

func subsets(nums []int) [][]int {
	var result [][]int
	backTrack(&result, nums, []int{}, -1, 0)
	return result

}

func backTrack(result *[][]int, nums, subset []int, preNode, index int) {
	if preNode == index-1 {
		*result = append(*result, append([]int{}, subset[:]...))
	}

	for i := index; i < len(nums); i++ {
		backTrack(result, nums, append(subset, nums[index]), index, i+1)
		index++
	}

}
