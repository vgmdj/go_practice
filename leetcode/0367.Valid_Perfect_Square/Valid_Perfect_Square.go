package Valid_Perfect_Square

func isPerfectSquare(num int) bool {
	if num < 1 {
		return false
	}

	var left, right = 0, num
	for left <= right {
		mid := (right + left) / 2
		square := mid * mid
		if square == num {
			return true

		} else if square > num {
			right = mid - 1

		} else {
			left = mid + 1

		}

	}

	return false

}
