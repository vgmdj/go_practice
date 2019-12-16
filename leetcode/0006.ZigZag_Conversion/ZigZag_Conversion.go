package ZigZag_Conversion

func convert(s string, numRows int) string {
	if numRows < 2 || len(s) <= numRows {
		return s
	}

	bytes := []byte(s)
	var result []byte
	dis, add := (numRows-1)*2, 0

	for i := 0; i < numRows; i++ {
		add = (numRows - i - 1) * 2
		for j := i; j < len(s); {
			result = append(result, bytes[j])

			j += add

			if add == 0 {
				j += dis
			}

			if dis-add != 0 {
				add = dis - add
			}

		}
	}

	return string(result)
}

func convert2(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	rows := Min(len(s), numRows)
	contents := make([][]byte, rows)

	row, step := 0, -1
	for i := 0; i < len(s); i++ {
		contents[row] = append(contents[row], s[i])
		if row == 0 || row == numRows-1 {
			step = -step
		}
		row += step
	}

	result := ""
	for _, v := range contents {
		result += string(v)
	}

	return result

}

func Min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
