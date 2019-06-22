package Find_Smallest_Letter_Greater_Than_Target

func nextGreatestLetter(letters []byte, target byte) byte {
	if len(letters) < 1 {
		return 0
	}

	if target >= letters[len(letters)-1] {
		return letters[0]
	}

	left, right := 0, len(letters)-1
	for left < right {
		mid := (left + right) / 2

		if letters[mid] > target {
			right = mid - 1

		} else {
			left = mid + 1

		}

	}

	return letters[left]
}
