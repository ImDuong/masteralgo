package dynamicprogramming

func MinPathSum(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 && j > 0 {
				grid[i][j] += grid[i][j-1]
				continue
			}
			if j == 0 && i > 0 {
				grid[i][j] += grid[i-1][j]
				continue
			}
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}
