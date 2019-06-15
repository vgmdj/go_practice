package Search_a_2D_Matrix_II

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row := len(matrix) - 1
	col := 0

	for row >= 0 && col < len(matrix[0]) {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			row--
		} else {
			col++
		}
	}
	return false
}

func searchMatrix2(matrix [][]int, target int) bool {
	if len(matrix) < 1 {
		return false
	}

	return helper(matrix, target, 0, 0, len(matrix)-1, len(matrix[0])-1)

}

func helper(matrix [][]int, target int, xs, ys, xe, ye int) bool {
	if xs < 0 || ys < 0 || xs > xe || ys > ye || xe > len(matrix)-1 || ye > len(matrix[xe])-1 ||
		matrix[xs][ys] > target || matrix[xe][ye] < target {
		return false
	}

	xmid, ymid := (xs+xe)/2, (ys+ye)/2
	if matrix[xmid][ymid] == target {
		return true
	}

	if matrix[xmid][ymid] > target {
		return helper(matrix, target, xs, ys, xmid-1, ye) ||
			helper(matrix, target, xmid, ys, xmid, ymid-1)
	}

	return helper(matrix, target, xs, ymid+1, xmid, ye) ||
		helper(matrix, target, xmid+1, ys, xe, ye)

}
