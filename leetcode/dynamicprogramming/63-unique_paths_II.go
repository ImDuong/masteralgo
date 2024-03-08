package dynamicprogramming

func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	if obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	for i := range obstacleGrid {
		if obstacleGrid[i][0] == 1 {
			break
		}
		obstacleGrid[i][0] = -1
	}

	for j := range obstacleGrid[0] {
		if obstacleGrid[0][j] == 1 {
			break
		}
		obstacleGrid[0][j] = -1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			a, b := 0, 0
			if obstacleGrid[i][j-1] != 1 {
				b = obstacleGrid[i][j-1]
			}
			if obstacleGrid[i-1][j] != 1 {
				a = obstacleGrid[i-1][j]
			}
			obstacleGrid[i][j] = a + b
		}
	}

	return obstacleGrid[m-1][n-1] * -1
}
