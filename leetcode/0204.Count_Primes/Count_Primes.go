package Count_Primes

func countPrimes(n int) int {
	if n < 3 {
		return 0
	}

	count := 1
	notPrimes := make([]bool, n)
	for i := 3; i < n; i += 2 {
		if notPrimes[i] {
			continue
		}

		count++
		for j := i * i; j < n; j += i {
			notPrimes[j] = true
		}

	}

	return count
}
