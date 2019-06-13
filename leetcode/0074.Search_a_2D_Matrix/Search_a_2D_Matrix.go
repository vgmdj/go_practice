package Search_a_2D_Matrix

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) < 1 {
		return false
	}

	row, column := len(matrix), len(matrix[0])
	pMin, pMax := 0, row*column-1

	for pMin <= pMax {
		mid := (pMin + pMax) / 2

		x := mid / column
		y := mid % column

		if matrix[x][y] == target {
			return true
		}

		if matrix[x][y] > target {
			pMax = mid - 1

		} else {
			pMin = mid + 1

		}

	}

	return false

}
