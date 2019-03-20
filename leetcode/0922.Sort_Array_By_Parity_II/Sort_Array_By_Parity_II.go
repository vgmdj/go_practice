package Sort_Array_By_Parity_II

//time O(n) space O(1)
func sortArrayByParityII(A []int) []int {
	even, odd := 1, 0
	for ; odd < len(A); odd += 2 {
		if A[odd]%2 == 0 {
			continue
		}

		for A[even]%2 == 1 {
			even += 2
		}

		A[even], A[odd] = A[odd], A[even]

	}

	return A
}

//time O(n) space O(n)
func sortArrayByParityII2(A []int) []int {
	size := len(A)
	res := make([]int, size)
	even, odd := 0, 1

	for _, a := range A {
		if a%2 == 0 {
			res[even] = a
			even += 2
		} else {
			res[odd] = a
			odd += 2
		}
	}

	return res
}
