package Number_of_Segments_in_a_String

func countSegments(s string) int {
	if len(s) == 0 {
		return 0
	}

	count := 0
	if s[0] != ' ' {
		count++
	}

	for i := 0; i < len(s)-1; i++ {
		if s[i] == ' ' && s[i+1] != ' ' {
			count++
		}
	}

	return count

}
