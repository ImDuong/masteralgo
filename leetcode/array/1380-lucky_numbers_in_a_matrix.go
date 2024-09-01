package array

func Challenge1380(matrix [][]int) []int {
	return luckyNumbers(matrix)
}

type luckyCell struct {
	value int
	pos   int
}

func luckyNumbers(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	rs := make([]int, 0)
	minRow := make([]luckyCell, m)
	maxCol := make([]luckyCell, n)
	for i := 0; i < m; i++ {
		minRow[i] = luckyCell{
			value: 100000,
		}
		for j := 0; j < n; j++ {
			if minRow[i].value > matrix[i][j] {
				minRow[i].value = matrix[i][j]
				minRow[i].pos = j
			}
		}
	}
	for i := 0; i < n; i++ {
		maxCol[i] = luckyCell{}
		for j := 0; j < m; j++ {
			if maxCol[i].value < matrix[j][i] {
				maxCol[i].value = matrix[j][i]
				maxCol[i].pos = j
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if minRow[i].pos == j && maxCol[j].pos == i {
				rs = append(rs, minRow[i].value)
			}
		}
	}
	return rs
}
