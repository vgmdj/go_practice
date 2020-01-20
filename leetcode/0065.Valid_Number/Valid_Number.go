package Valid_Number

func isNumber(s string) bool {
	var (
		finals = []bool{false, false, false, true, false, true, true, false, true}
		status = [][]int{{0, 1, 6, 2, -1, -1},
			{-1, -1, 6, 2, -1, -1},
			{-1, -1, 3, -1, -1, -1},
			{8, -1, 3, -1, 4, -1},
			{-1, 7, 5, -1, -1, -1},
			{8, -1, 5, -1, -1, -1},
			{8, -1, 6, 3, 4, -1},
			{-1, -1, 5, -1, -1, -1},
			{8, -1, -1, -1, -1, -1},
		}
	)
	row := 0
	for i := 0; i < len(s); i++ {
		row = status[row][column(s[i])]
		if row == -1 {
			return false
		}
	}

	return finals[row]
}

func column(bt byte) int {
	switch bt {
	case ' ':
		return 0
	case '+', '-':
		return 1
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return 2
	case '.':
		return 3
	case 'e':
		return 4
	default:
		return 5
	}
}
