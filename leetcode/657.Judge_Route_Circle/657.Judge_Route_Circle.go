package Judge_Route_Circle

func judgeCircle(moves string) bool {
	var x, y int

	for _, v := range moves {
		switch v {
		default:
			return false
		case 'U':
			y++
		case 'D':
			y--
		case 'L':
			x++
		case 'R':
			x--
		}
	}

	return x == 0 && y == 0

}
