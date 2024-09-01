package array

func Challenge807(grid [][]int) int {
	return maxIncreaseKeepingSkyline(grid)
}

// this problem turns to change a cell (i, j) to min(maxRow(i), maxColumn(j))
// -> the procedure is nearly the same with Challenge1380
func maxIncreaseKeepingSkyline(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	maxRow := make([]luckyCell, m)
	maxCol := make([]luckyCell, n)
	for i := 0; i < m; i++ {
		maxRow[i] = luckyCell{}
		for j := 0; j < n; j++ {
			if maxRow[i].value < grid[i][j] {
				maxRow[i].value = grid[i][j]
				maxRow[i].pos = j
			}
		}
	}
	for i := 0; i < n; i++ {
		maxCol[i] = luckyCell{}
		for j := 0; j < m; j++ {
			if maxCol[i].value < grid[j][i] {
				maxCol[i].value = grid[j][i]
				maxCol[i].pos = j
			}
		}
	}
	maxIncr := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			maxIncr += min(maxRow[i].value, maxCol[j].value) - grid[i][j]
		}
	}
	return maxIncr
}
