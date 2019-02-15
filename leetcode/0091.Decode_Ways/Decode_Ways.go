package _091_Decode_Ways

func numDecodings(s string) int {
	if len(s) < 1 || s[0] == '0' {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	if s[1] == '0' && s[0] != '1' && s[0] != '2' {
		return 0
	}

	var dp = make([]int, len(s))
	dp[0] = 1

	if s[:2] > "26" || s[:2] == "10" || s[:2] == "20" {
		dp[1] = 1
	} else {
		dp[1] = 2
	}

	for i := 2; i < len(s); i++ {
		if s[i] == '0' {
			if s[i-1] != '1' && s[i-1] != '2' {
				return 0
			}

			dp[i] = dp[i-2]
			continue

		}

		if s[i-1] == '0' || s[i-1:i+1] > "26"{
			dp[i] = dp[i-1]
			continue
		}

		if s[i-1:i+1] <= "26" {
			dp[i] = dp[i-1] + dp[i-2]
			continue
		}
	}

	return dp[len(s)-1]

}
