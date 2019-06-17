package First_Unique_Character_in_a_String

func firstUniqChar(s string) int {
	counter := make([]int, 26)

	for _, v := range s {
		counter[v-'a']++
	}

	for k, v := range s {
		if counter[v-'a'] == 1 {
			return k
		}

	}

	return -1

}
