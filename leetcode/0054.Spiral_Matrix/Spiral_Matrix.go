package Spiral_Matrix

func spiralOrder(matrix [][]int) []int {
	var result []int
	if len(matrix) == 0 {
		return nil
	}

	var row, column = len(matrix) - 1, len(matrix[0]) - 1
	var count = (row + 1) * (column + 1)

	for i := 0; i <= row && column >= 0 && row >= 0; i++ {
		for j := i; j <= column; j++ {
			result = append(result, matrix[i][j])
		}

		for j := i + 1; j <= row; j++ {
			result = append(result, matrix[j][column])
		}

		for j := column - 1; j >= i; j-- {
			result = append(result, matrix[row][j])
		}

		for j := row - 1; j > i; j-- {
			result = append(result, matrix[j][i])
		}

		row--
		column--

	}

	return result[:count]

}
