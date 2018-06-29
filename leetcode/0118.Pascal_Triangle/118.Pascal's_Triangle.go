package Pascal_s_Triangle

func generate(numRows int) [][]int {
	result := [][]int{}
	rows := []int{}

	for i := 0; i < numRows; i++ {
		rows = generateRow(rows)

		result = append(result, rows)
	}

	return result

}

func generateRow(rows []int) (result []int) {

	for i := 0; i < len(rows); i++ {
		if i == 0 {
			result = append(result, rows[0])
			continue
		}

		result = append(result, rows[i]+rows[i-1])
	}

	result = append(result, 1)

	return result

}
