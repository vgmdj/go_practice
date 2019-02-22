package Combinations

func combine(n int, k int) [][]int {
	var result [][]int
	backTrack(&result, []int{}, n, k, 1)
	return result
}

func backTrack(result *[][]int, subset []int, n, k, start int) {
	if k == 0 {
		*result = append(*result, append([]int{}, subset[:]...))
		return
	}

	for i := start; i <= n && i <= n-k+1; i++ {
		backTrack(result, append(subset, i), n, k-1, i+1)

	}

}
