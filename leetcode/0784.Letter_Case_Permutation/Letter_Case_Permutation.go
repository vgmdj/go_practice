package Letter_Case_Permutation

func letterCasePermutation(S string) []string {
	result := []string{S}
	backTracking(S, &result, 0)
	return result

}

func backTracking(S string, result *[]string, pos int) {
	subset := []byte(S)
	for i := pos; i < len(S); i++ {
		if (subset[i] < 'a' || subset[i] > 'z') && (subset[i] < 'A' || subset[i] > 'Z') {
			continue
		}

		if subset[i] >= 'a' && subset[i] <= 'z' {
			subset[i] = subset[i] + 'A' - 'a'

		} else if subset[i] >= 'A' && subset[i] <= 'Z' {
			subset[i] = subset[i] - 'A' + 'a'

		}

		*result = append(*result, string(subset))
		backTracking(string(subset), result, i+1)
		subset[i] = S[i]

	}

}
