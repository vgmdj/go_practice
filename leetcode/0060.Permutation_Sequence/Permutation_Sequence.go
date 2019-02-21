package Permutation_Sequence

func getPermutation(n int, k int) string {
	var result string
	var visited = make([]bool, n+1)
	var getN = make(map[int]int)
	for i := -1; i <= n; i++ {
		if i < 2 {
			getN[i] = 1
			continue
		}

		getN[i] = i * getN[i-1]
	}

	backTrack(&result, "", n, k, 1, visited, getN)

	return result
}

func backTrack(result *string, permutation string, n, k, deep int, visited []bool, getN map[int]int) {
	series := getN[n-deep]
	if k == 1 && deep == n+1 {
		*result = permutation
		return
	}

	var count = 1
	for i := 1; i <= n; i++ {
		if visited[i] {
			continue
		}

		if series*count < k {
			count++
			continue
		}

		visited[i] = true
		backTrack(result, permutation+string(i+'0'), n, k-series*(count-1), deep+1, visited, getN)
	}

}
