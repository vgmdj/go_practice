package Surrounded_Regions

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}

	visited := make([][]bool, len(board))
	for i, rows := range board {
		visited[i] = make([]bool, len(rows))
	}

	var mark func(row, col int)
	mark = func(row, col int) {
		if row < 0 || col < 0 || row >= len(board) || col >= len(board[row]) ||
			visited[row][col] || board[row][col] == 'X' {
			return
		}

		visited[row][col] = true
		mark(row-1, col)
		mark(row+1, col)
		mark(row, col-1)
		mark(row, col+1)

	}

	for i := 0; i < len(board); i++ {
		mark(i, 0)
		mark(i, len(board[0])-1)
	}

	for i := 1; i < len(board[0])-1; i++ {
		mark(0, i)
		mark(len(board)-1, i)
	}

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			if board[row][col] == 'X' || visited[row][col] {
				continue
			}

			board[row][col] = 'X'
		}

	}

}
