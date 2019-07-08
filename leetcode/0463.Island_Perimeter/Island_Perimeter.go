package Island_Perimeter

func islandPerimeter(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	result := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == 1 {
				result += perimeter(grid, r, c)
			}
		}
	}

	return result

}

func perimeter(grid [][]int, x, y int) int {
	row := len(grid)
	column := len(grid[0])
	result := 0

	if x == 0 || grid[x-1][y] == 0 {
		result += 1
	}

	if y == 0 || grid[x][y-1] == 0 {
		result += 1
	}

	if y == column-1 || grid[x][y+1] == 0 {
		result += 1
	}

	if x == row-1 || grid[x+1][y] == 0 {
		result += 1
	}

	return result

}
