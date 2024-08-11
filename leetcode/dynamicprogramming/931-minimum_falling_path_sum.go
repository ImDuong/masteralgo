package dynamicprogramming

func MinFallingPathSum(matrix [][]int) int {
	// f[i][j] = min(f[i + 1][j - 1], f[i + 1][j], f[i + 1][j + 1])
	n := len(matrix)
	if n == 1 {
		return matrix[0][0]
	}
	f := make([][]int, n)
	f[n-1] = make([]int, n)
	for j := 0; j < n; j++ {
		f[n-1][j] = matrix[n-1][j]
	}
	globalMinS := 100 * n * n
	for i := n - 2; i >= 0; i-- {
		f[i] = make([]int, n)
		for j := 0; j < n; j++ {
			localMinS := f[i+1][j]
			if j-1 >= 0 {
				localMinS = min(localMinS, f[i+1][j-1])
			}
			if j+1 < n {
				localMinS = min(localMinS, f[i+1][j+1])
			}
			f[i][j] = localMinS + matrix[i][j]
			if i == 0 {
				globalMinS = min(globalMinS, f[i][j])
			}
		}
	}
	return globalMinS
}
