package Pascal_Triangle_II

func getRow(rowIndex int) []int {
	rows := []int{}

	for i := 0; i <= rowIndex; i++ {
		rows = generateRow(rows)

	}

	return rows

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
