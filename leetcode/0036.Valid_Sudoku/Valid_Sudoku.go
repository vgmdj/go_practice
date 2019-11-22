package Valid_Sudoku

func isValidSudoku(board [][]byte) bool {
	rows := [9][9]bool{}
	columns := [9][9]bool{}
	boxes := [9][9]bool{}

	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			boxIndex := r/3*3 + c/3
			num := board[r][c]
			if num == '.' {
				continue
			}

			num -= '1'

			if rows[r][num] || columns[c][num] || boxes[boxIndex][num] {
				return false
			}

			rows[r][num] = true
			columns[c][num] = true
			boxes[boxIndex][num] = true

		}

	}

	return true

}
