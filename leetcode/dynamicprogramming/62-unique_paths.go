package dynamicprogramming

func UniquePaths(m int, n int) int {
	paths := make([][]int, m)
	for i := range paths {
		paths[i] = make([]int, n)
		for j := range paths[i] {
			paths[0][j] = 1
		}
		paths[i][0] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			paths[i][j] = paths[i][j-1] + paths[i-1][j]
		}
	}
	return paths[m-1][n-1]
}

func UniquePathsV2(m int, n int) int {
	paths := make([]int, n)
	for i := range paths {
		paths[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			paths[j] += paths[j-1]
		}
	}
	return paths[n-1]
}
