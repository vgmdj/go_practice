package Increasing_Decreasing_String

func sortString(s string) string {
	result := make([]byte, len(s))
	alphabet := [26]int{}
	for i := range s {
		c := s[i] - 'a'
		alphabet[c]++
	}

	for p := 0; p < len(s); {
		for i := 0; i < len(alphabet); i++ {
			if alphabet[i] > 0 {
				result[p] = byte(i) + 'a'
				alphabet[i]--
				p++
			}
		}

		for i := len(alphabet)-1; i >= 0; i-- {
			if alphabet[i] > 0 {
				result[p] = byte(i) + 'a'
				alphabet[i]--
				p++
			}
		}
	}

	return string(result)

}
