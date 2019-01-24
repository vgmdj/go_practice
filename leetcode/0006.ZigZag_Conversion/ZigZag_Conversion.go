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
