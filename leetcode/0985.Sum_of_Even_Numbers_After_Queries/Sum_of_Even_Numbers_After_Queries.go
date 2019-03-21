package Sum_of_Even_Numbers_After_Queries

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	var res = make([]int, len(A))
	sum := 0
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			sum += A[i]
		}
	}

	for i := 0; i < len(A); i++ {
		val, index := queries[i][0], queries[i][1]
		if A[index]&1 == 0 {
			sum -= A[index]
		}

		A[index] += val
		if A[index]&1 == 0 {
			sum += A[index]
		}

		res[i] = sum

	}

	return res

}
