package Container_With_Most_Water

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}

	left, right := 0, len(height)-1
	max := 0
	for left < right {
		max = Max(Min(height[left], height[right])*(right-left), max)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}

	}

	return max

}

func Max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
