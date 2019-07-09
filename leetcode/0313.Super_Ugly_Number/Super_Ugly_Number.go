package Super_Ugly_Number

import "math"

func nthSuperUglyNumber(n int, primes []int) int {
	dp := make([]int, n)
	multiple := make([]int, len(primes))
	dp[0] = 1

	for i := 1; i < n; i++ {
		minimum := math.MaxInt32
		for j := 0; j < len(multiple); j++ {
			minimum = min(minimum, primes[j]*dp[multiple[j]])
		}
		dp[i] = minimum

		for j := 0; j < len(multiple); j++ {
			if dp[multiple[j]]*primes[j] == minimum {
				multiple[j]++
			}

		}

	}

	return dp[n-1]

}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
