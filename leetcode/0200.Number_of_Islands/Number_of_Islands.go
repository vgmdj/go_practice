package Number_of_Islands

func numIslands(grid [][]byte) int {
	islands := 0

	for i, row := range grid {
		for j, col := range row {
			if col == '0' {
				continue
			}

			islands++
			markIsland(grid, i, j)
		}
	}

	return islands
}

func markIsland(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) ||
		grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'

	markIsland(grid, i-1, j) // up
	markIsland(grid, i+1, j) // down
	markIsland(grid, i, j-1) // left
	markIsland(grid, i, j+1) // right
}
