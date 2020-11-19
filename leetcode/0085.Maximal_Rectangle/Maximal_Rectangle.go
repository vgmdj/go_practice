package maximal_rectangle

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	result := 0
	dp := make([]int, len(matrix[0]))

	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] == '1' {
				dp[col] = dp[col] + 1

			} else {
				dp[col] = 0

			}
		}
		result = max(result, maxRowAera(dp))
	}

	return result

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxRowAera(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	result := 0
	stack := []int{-1}
	for i := range heights {
		for len(stack) > 1 && heights[i] < heights[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			area := heights[top] * (i - stack[len(stack)-1] - 1)
			result = max(area, result)

		}

		stack = append(stack, i)
	}

	for len(stack) > 1 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		area := heights[top] * (len(heights) - stack[len(stack)-1] - 1)
		result = max(area, result)
	}

	return result

}
