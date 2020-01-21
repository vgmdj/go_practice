package Convert_Integer_to_the_Sum_of_Two_No_Zero_Integers

func getNoZeroIntegers(n int) []int {
	a, b := 1, n-1
	mul := 1
	for b/10 != 0 {
		c := b % 10
		if c != 1 && c != 0 {
			mul *= 10
			a += mul
			b /= 10
			continue
		}

		a += mul
		b--
	}

	return []int{a, n - a}
}
