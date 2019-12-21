package Squares_of_a_Sorted_Array

func sortedSquares(A []int) []int {
	result := make([]int, len(A))
	index := len(result) - 1
	left, right := 0, len(A)-1

	for left <= right {
		if left == right {
			result[index] = A[left] * A[right]
			return result[index:]
		}

		lv, rv := A[left], A[right]
		if lv < 0 {
			lv = -lv
		}

		if rv < 0 {
			rv = -rv
		}

		if rv >= lv {
			result[index] = rv * rv
			right--
			index--

		} else {
			result[index] = lv * lv
			left++
			index--

		}

	}

	return result[index:]

}
