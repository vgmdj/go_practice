package Combination_Sum

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	sort.Ints(candidates)
	backTrack(&result, []int{}, candidates, 0, 0, target)
	return result
}

func backTrack(result *[][]int, subset, candidates []int, start, sum, target int) {
	if sum == target {
		*result = append(*result, append([]int{}, subset[:]...))
		return
	}

	for i := start; i < len(candidates) && sum+candidates[i] <= target; i++ {
		if i != start && candidates[i] == candidates[i-1] {
			continue
		}
		backTrack(result, append(subset, candidates[i]), candidates, i+1, sum+candidates[i], target)
	}

}
