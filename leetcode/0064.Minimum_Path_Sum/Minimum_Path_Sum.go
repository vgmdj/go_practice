package _064_Minimum_Path_Sum

func minPathSum(grid [][]int) int {
	if len(grid) < 1 {
		return 0
	}

	var row, column int
	for row = 0; row < len(grid); row++ {
		for column = 0; column < len(grid[row]); column++ {
			if row == 0 && column == 0 {
				continue
			}

			if row == 0 {
				grid[0][column] += grid[0][column-1]
				continue
			}

			if column == 0 {
				grid[row][0] += grid[row-1][0]
				continue
			}

			grid[row][column] = min(grid[row][column]+grid[row-1][column],
				grid[row][column]+grid[row][column-1])

		}

	}

	return grid[row-1][column-1]

}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
