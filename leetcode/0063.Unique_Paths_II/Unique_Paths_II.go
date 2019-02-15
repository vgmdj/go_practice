package Unique_Paths_II

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) < 1 {
		return 0
	}

	var row, column int
	for row = 0; row < len(obstacleGrid); row++ {
		for column = 0; column < len(obstacleGrid[row]); column++ {
			if obstacleGrid[row][column] == 1 {
				obstacleGrid[row][column] = 0
				continue
			}

			if (row == 0 && column > 0 && obstacleGrid[row][column-1] == 0) ||
				(column == 0 && row > 0 && obstacleGrid[row-1][column] == 0) {
				obstacleGrid[row][column] = 0
				continue
			}

			if row == 0 || column == 0 {
				obstacleGrid[row][column] = 1
				continue
			}

			obstacleGrid[row][column] = obstacleGrid[row-1][column] + obstacleGrid[row][column-1]
		}
	}

	return obstacleGrid[row-1][column-1]

}
