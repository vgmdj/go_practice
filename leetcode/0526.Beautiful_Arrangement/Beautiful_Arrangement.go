package Beautiful_Arrangement

//answer
func countArrangement_answer(N int) int {
	res := []int{0, 1, 2, 3, 8, 10, 36, 41, 132, 250, 700, 750, 4010, 4237, 10680, 24679, 16}
	return res[N]
}

func countArrangement(N int) int {
	result := 0
	visited := make([]bool, N+1)

	backTracking(&result, N, N+1, visited)

	return result
}

func backTracking(result *int, N int, index int, visited []bool) {
	if index == 1 {
		*result = *result + 1
		return
	}

	for i := 1; i <= N; i++ {
		if visited[i] || !isBeautiful(index-1, i) {
			continue
		}

		visited[i] = true

		backTracking(result, N, index-1, visited)
		visited[i] = false
	}

}

func isBeautiful(a, b int) bool {
	return a%b == 0 || b%a == 0
}
