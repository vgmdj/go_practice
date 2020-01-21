package Minimum_Flips_to_Make_a_OR_b_Equal_to_c

func minFlips(a int, b int, c int) int {
	result := 0
	for c != 0 || a != 0 || b != 0 {
		ca, cb, cc := a&1, b&1, c&1
		a, b, c = a>>1, b>>1, c>>1
		if cc == 0 {
			result += ca + cb

		} else if ca != 1 && cb != 1 {
			result++

		}

	}

	return result

}
