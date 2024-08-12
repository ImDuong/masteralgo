package dynamicprogramming

import (
	"fmt"
	"strconv"
)

// when matrix[i][j] == '1', f[i][j] = min(f[i - 1][j - 1], f[i][j - 1], f[i - 1][j]) + 1
// else, f[i][j] = 0
func MaximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])
	f := make([][]int, m)

	globalMaxArea := 0
	for i := 0; i < m; i++ {
		f[i] = make([]int, n)
		for j := 0; j < n; j++ {
			byteToInt, _ := strconv.Atoi(string(matrix[i][j]))
			f[i][j] = byteToInt
			if f[i][j] == 0 {
				continue
			}
			if i == 0 || j == 0 {
				globalMaxArea = max(globalMaxArea, f[i][j]*f[i][j])
				continue
			}
			localMinArea := f[i-1][j-1]
			localMinArea = min(localMinArea, f[i-1][j])
			localMinArea = min(localMinArea, f[i][j-1])
			f[i][j] = f[i][j] + localMinArea
			globalMaxArea = max(globalMaxArea, f[i][j]*f[i][j])
			if globalMaxArea == 9 {
				fmt.Println("here")
			}
		}
	}
	return globalMaxArea
}
