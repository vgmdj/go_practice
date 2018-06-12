package Excel_Sheet_Column_Number

func titleToNumber(s string) int {
	sum := 0

	for _, ch := range s {
		current := int(ch-'A') + 1
		sum = sum*26 + current
	}

	return sum

}
