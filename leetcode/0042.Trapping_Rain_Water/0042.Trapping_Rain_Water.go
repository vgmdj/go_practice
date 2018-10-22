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

//dynamic programming
func trap2(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	result := 0
	leftMax, rightMax := make([]int, len(height)), make([]int, len(height))

	leftMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		if height[i] > leftMax[i-1] {
			leftMax[i] = height[i]
		} else {
			leftMax[i] = leftMax[i-1]
		}
	}

	rightMax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		if height[i] > rightMax[i+1] {
			rightMax[i] = height[i]
		} else {
			rightMax[i] = rightMax[i+1]
		}
	}

	for i := 0; i < len(height)-1; i++ {
		if leftMax[i] < rightMax[i] {
			result += leftMax[i] - height[i]
		} else {
			result += rightMax[i] - height[i]
		}
	}

	return result
}
