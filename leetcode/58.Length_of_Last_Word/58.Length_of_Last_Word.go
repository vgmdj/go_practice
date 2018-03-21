package Length_of_Last_Word

func lengthOfLastWord(s string) int {
	var count int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && count == 0 {
			continue
		}

		if s[i] == ' ' {
			break
		}

		count++

	}

	return count
}
