package Trapping_Rain_Water

//brute force
func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	var result = 0
	for i := 1; i < len(height)-1; i++ {
		maxLeft, maxRight := 0, 0
		for left := i; left >= 0; left-- {
			if height[left] > maxLeft {
				maxLeft = height[left]
			}
		}

		for right := i; right < len(height); right++ {
			if height[right] > maxRight {
				maxRight = height[right]
			}
		}

		if maxLeft > maxRight {
			result += maxRight - height[i]
			continue
		}

		result += maxLeft - height[i]

	}

	return result

}
