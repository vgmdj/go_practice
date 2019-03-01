package Spiral_Matrix_II

func generateMatrix(n int) [][]int {
	var row, column = n-1, n-1
	var matrix = make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	var result = 1
	for i := 0; i <= row && column >= 0 && row >= 0 && result <= n*n; i++ {
		for j := i; j <= column; j++ {
			matrix[i][j] = result
			result++
		}

		for j := i + 1; j <= row; j++ {
			matrix[j][column] = result
			result++
		}

		for j := column - 1; j >= i; j-- {
			matrix[row][j] = result
			result++
		}

		for j := row - 1; j > i; j-- {
			matrix[j][i] = result
			result++
		}

		row--
		column--

	}

	return matrix

}
