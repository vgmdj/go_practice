package Word_Search

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(word) == 0 {
		return false
	}

	rlen := len(board)
	clen := len(board[0])

	visited := make([]bool, rlen*clen)

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			if helper(board, word, 0, row, col, visited) {
				return true
			}

		}
	}

	return false

}

func helper(board [][]byte, word string, wp int, px, py int, visited []bool) bool {
	if wp >= len(word) {
		return true
	}

	if px < 0 || px >= len(board) || py < 0 || py >= len(board[0]) || visited[px*len(board[px])+py] || board[px][py] != word[wp] {
		return false
	}

	visited[px*len(board[px])+py] = true

	try := helper(board, word, wp+1, px-1, py, visited) ||
		helper(board, word, wp+1, px+1, py, visited) ||
		helper(board, word, wp+1, px, py-1, visited) ||
		helper(board, word, wp+1, px, py+1, visited)
	visited[px*len(board[px])+py] = try

	return try

}
