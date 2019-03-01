package Rotate_Image

func rotate(matrix [][]int) {
	var boundary = len(matrix) - 1
	if boundary < 1 {
		return
	}

	for i := 0; i < boundary; i, boundary = i+1, boundary-1 {
		for j := i; j < boundary; j++ {
			matrix[i][j], matrix[boundary-j+i][i] = matrix[boundary-j+i][i], matrix[i][j]
		}

		for j := boundary; j > i; j-- {
			matrix[j][i], matrix[boundary][j] = matrix[boundary][j], matrix[j][i]
		}

		for j := i; j < boundary; j++ {
			matrix[boundary][boundary-j+i], matrix[j][boundary] = matrix[j][boundary], matrix[boundary][boundary-j+i]
		}
	}

}
