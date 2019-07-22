package Combination_Sum_III

func combinationSum3(k int, n int) [][]int {
	results := make([][]int, 0)
	backTracking(k, n, 0, 0, []int{}, &results)
	return results

}

func backTracking(k, n, last, sum int, subsets []int, results *[][]int) {
	if (n-sum) > (19-k)*k/2 || (k == 1 && n-sum <= last) || k <= 0 {
		return
	}

	if n-sum > last && k == 1 {
		subsets = append(subsets, n-sum)
		*results = append(*results, append([]int{}, subsets...))
		return
	}

	max := (n - sum - k + 1) / k

	for i := last + 1; i <= max && i <= 9; i++ {
		backTracking(k-1, n, i, sum+i, append(subsets, i), results)
	}

}
