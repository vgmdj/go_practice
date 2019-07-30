package Maximal_Square

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	maxLength := byte(0)

	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == '1' {
			maxLength = '1'
			break
		}
	}

	for i := 0; i < len(matrix) && maxLength == 0; i++ {
		if matrix[i][0] == '1' {
			maxLength = '1'
			break
		}
	}

	for row := 1; row < len(matrix); row++ {
		for column := 1; column < len(matrix[row]); column++ {
			if matrix[row][column] == '0' {
				continue
			}

			matrix[row][column] = min(matrix[row-1][column],
				matrix[row][column-1],
				matrix[row-1][column-1]) + 1

			if maxLength < matrix[row][column] {
				maxLength = matrix[row][column]
			}
		}
	}

	maxLength -= '0'

	return int(maxLength * maxLength)

}

func min(a, b, c byte) byte {
	if a > b {
		a, b = b, a
	}

	if a > c {
		a, c = c, a
	}

	return a
}
