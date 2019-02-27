package Repeated_DNA_Sequences

func findRepeatedDnaSequences(s string) []string {
	if len(s) < 11 {
		return []string{}
	}

	var result = make([]string, 0, len(s)-9)
	var ref = make(map[string]int, len(s)-9)
	for i := 0; i < len(s)-9; i++ {
		str := s[i : i+10]
		if count, ok := ref[str]; ok && count == 1 {
			result = append(result, str)
		}

		ref[str]++

	}

	return result

}
