package XOR_Queries_of_a_Subarray

func xorQueries(arr []int, queries [][]int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	result := make([]int,len(queries))
	dp := make([]int,len(arr))
	dp[0] = arr[0]
	for i := 1 ;i < len(arr); i++{
		dp[i] = dp[i-1] ^ arr[i]
	}

	for i :=0 ;i < len(queries) ; i++ {
		if queries[i][0] == 0 {
			result[i] = dp[queries[i][1]]
			continue
		}

		result[i] = dp[queries[i][0]-1] ^ dp[queries[i][1]]

	}

	return result

}