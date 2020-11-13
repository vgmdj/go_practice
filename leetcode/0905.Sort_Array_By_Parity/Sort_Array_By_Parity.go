package Sort_Array_By_Parity

func sortArrayByParity(A []int) []int {
	left, right := 0, len(A)-1
	for ; left < len(A) && left < right; left++ {
		if A[left]%2 == 0 {
			continue
		}

		for right > left && A[right]%2 == 1 {
			right--
		}

		A[left], A[right] = A[right], A[left]

	}

	return A

}
