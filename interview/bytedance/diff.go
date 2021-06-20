package bytedance

/*
数组a,先递增再递减，输出数组中不同元素个数。要求：O(1)空间复杂度，尽可能小的时间复杂度，不能改变原数组。 如arr=1,2,4,6,3,2,2 输出5


*/
func Diff(data []int) int {
	if len(data) <= 1 {
		return len(data)
	}

	result := 0
	left, right := 0, len(data)-1
	for left <= right {
		if data[left] < data[right] {
			if left < len(data)-1 && data[left] != data[left+1] {
				result++
			}

			left++

		} else if data[left] > data[right] {
			if right > 0 && data[right] != data[right-1] {
				result++
			}

			right--

		} else {
			value := data[left]
			for left < len(data) && data[left] == value {
				left++
			}

			for right >= 0 && data[right] == value {
				right--
			}
			result++

		}

	}

	return result

}
