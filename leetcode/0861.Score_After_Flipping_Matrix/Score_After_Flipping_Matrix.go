package Score_After_Flipping_Matrix

func matrixScore(A [][]int) int {
	rowLength, colLength := len(A), len(A[0])
	result := 1 << (colLength - 1) * rowLength
	for col := 1; col < colLength; col++ {
		count := 0
		for row := 1; row < rowLength; row++ {
			if A[row][col]^A[row][0] == 1 {
				count++
			}
		}

		if count < rowLength-count {
			count = rowLength - count
		}

		result += 1 << (colLength - col - 1) * count

	}

	return result

}
