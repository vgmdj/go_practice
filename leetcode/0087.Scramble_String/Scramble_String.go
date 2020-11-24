package Scramble_String

func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}

	if !hasSameChars(s1, s2) {
		return false
	}

	for i := 1; i < len(s1); i++ {
		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
			return true
		}

		if isScramble(s1[:i], s2[len(s2)-i:]) && isScramble(s1[i:], s2[:len(s2)-i]) {
			return true
		}

	}

	return false
}

func hasSameChars(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	m := [26]int{}
	for _, v := range str1 {
		m[v-'a']++
	}

	for _, v := range str2 {
		m[v-'a']--
	}

	for i := range m {
		if m[i] != 0 {
			return false
		}
	}

	return true

}

func isScrambleDP(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	length := len(s1)
	dp := make([][][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([][]bool, length)
		for j := 0; j < length; j++ {
			dp[i][j] = make([]bool, length+1)
			if s1[i] == s2[j] {
				dp[i][j][1] = true
			}
		}
	}

	for l := 2; l <= length; l++ {
		for i := 0; i <= length-l; i++ {
			for j := 0; j <= length-l; j++ {
				for k := 1; k <= l-1; k++ {
					if dp[i][j][k] && dp[i+k][j+k][l-k] {
						dp[i][j][l] = true
						break
					}
					if dp[i][j+l-k][k] && dp[i+k][j][l-k] {
						dp[i][j][l] = true
						break
					}
				}
			}
		}
	}

	return dp[0][0][length]
}
