package algoI

import "masteralgo/internal/core/domain"

var directions = []domain.Cell{
	{
		Row: 1,
		Col: 0,
	},
	{
		Row: -1,
		Col: 0,
	},
	{
		Row: 0,
		Col: 1,
	},
	{
		Row: 0,
		Col: -1,
	},
}

func UpdateMatrix(mat [][]int) [][]int {
	queue := []domain.Cell{}

	// find all zero-cells first
	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == 0 {
				queue = append(queue, domain.Cell{
					Row: i,
					Col: j,
				})
			} else {
				mat[i][j] = -1
			}
		}
	}

	// process layers by layers
	for len(queue) > 0 {
		var poppedCell domain.Cell
		poppedCell, queue = queue[0], queue[1:]
		adjCells := getAdjCells(poppedCell, len(mat), len(mat[0]))
		for i := range adjCells {
			if mat[adjCells[i].Row][adjCells[i].Col] == -1 {
				mat[adjCells[i].Row][adjCells[i].Col] = mat[poppedCell.Row][poppedCell.Col] + 1
				queue = append(queue, domain.Cell{
					Row: adjCells[i].Row,
					Col: adjCells[i].Col,
				})
			}
		}
	}
	return mat
}

func getAdjCells(curCell domain.Cell, nbRows, nbCols int) []domain.Cell {
	cells := []domain.Cell{}
	for i := range directions {
		row := curCell.Row + directions[i].Row
		col := curCell.Col + directions[i].Col
		if !(row < 0 || row >= nbRows || col < 0 || col >= nbCols) {
			cells = append(cells, domain.Cell{
				Row: row,
				Col: col,
			})
		}
	}
	return cells
}
