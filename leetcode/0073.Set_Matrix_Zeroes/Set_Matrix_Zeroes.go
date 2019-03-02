package Set_Matrix_Zeroes

func setZeroes(matrix [][]int) {
	if len(matrix) < 1 {
		return
	}

	var row = make([]bool, len(matrix))
	var column = make([]bool, len(matrix[0]))
	for r, rv := range matrix {
		for c, cv := range rv {
			if cv == 0 {
				row[r] = true
				column[c] = true
			}

		}
	}

	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			if row[r] || column[c] {
				matrix[r][c] = 0
			}
		}
	}

	return
}
