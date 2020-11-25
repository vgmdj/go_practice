package Ways_to_Make_a_Fair_Array

func waysToMakeFair(nums []int) int {
	result, odd, even := 0, 0, 0
	for i, n := range nums {
		if i&1 == 0 {
			even += n
		} else {
			odd += n
		}
	}

	newOdd, newEven := 0, 0
	for i, n := range nums {
		if i&1 == 0 {
			even -= n
		} else {
			odd -= n
		}

		if odd+newOdd == even+newEven {
			result++
		}
		if i&1 == 0 {
			newOdd += n
		} else {
			newEven += n
		}
	}
	return result
}
