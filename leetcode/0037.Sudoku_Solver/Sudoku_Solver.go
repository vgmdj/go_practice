package Sudoku_Solver

const SudokuNumber = 9

func solveSudoku(board [][]byte) {

	rows := [SudokuNumber][SudokuNumber]bool{}
	columns := [SudokuNumber][SudokuNumber]bool{}
	subs := [SudokuNumber][SudokuNumber]bool{}

	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			if board[r][c] == '.' {
				continue
			}

			num := board[r][c] - '1'
			subIndex := r/3*3 + c/3

			rows[r][num] = true
			columns[c][num] = true
			subs[subIndex][num] = true

		}
	}

	backTrack(0, board, rows, columns, subs)

}

func backTrack(index int, board [][]byte, rows, columns, subs [SudokuNumber][SudokuNumber]bool) bool {
	if index == SudokuNumber*SudokuNumber {
		return true
	}

	row := index / SudokuNumber
	column := index % SudokuNumber
	subIndex := row/3*3 + column/3

	if board[row][column] != '.' {
		return backTrack(index+1, board, rows, columns, subs)
	}

	for i := 0; i < SudokuNumber; i++ {
		if rows[row][i] || columns[column][i] || subs[subIndex][i] {
			continue
		}

		num := byte(i) + '1'
		board[row][column] = num
		rows[row][i] = true
		columns[column][i] = true
		subs[subIndex][i] = true

		if backTrack(index+1, board, rows, columns, subs) {
			return true
		}

		rows[row][i] = false
		columns[column][i] = false
		subs[subIndex][i] = false
		board[row][column] = '.'

	}

	return false

}
