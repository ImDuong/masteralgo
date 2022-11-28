package algoI

import "masteralgo/internal/core/domain"

func OrangesRotting(grid [][]int) int {
	nbMinutes := -1
	nbRottenOranges := 0
	nbEmptyCells := 0
	queue := []domain.Cell{}

	// find all rotten oranges first
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 2 {
				queue = append(queue, domain.Cell{
					Row: i,
					Col: j,
				})
				nbRottenOranges++
			} else if grid[i][j] == 0 {
				nbEmptyCells++
			}
		}
	}
	if nbRottenOranges+nbEmptyCells == len(grid)*len(grid[0]) {
		return 0
	}

	// process layers by layers
	for len(queue) > 0 {
		// pop all rotten oranges in once
		poppedCells := make([]domain.Cell, len(queue))
		i := 0
		for len(queue) > 0 {
			var poppedCell domain.Cell
			poppedCell, queue = queue[0], queue[1:]

			// avoid append to make code run faster
			poppedCells[i] = poppedCell
			i++
		}
		for idx := range poppedCells {
			adjCells := getAdjCells(poppedCells[idx], len(grid), len(grid[0]))
			for i := range adjCells {
				if grid[adjCells[i].Row][adjCells[i].Col] == 1 {
					grid[adjCells[i].Row][adjCells[i].Col] = 2
					queue = append(queue, domain.Cell{
						Row: adjCells[i].Row,
						Col: adjCells[i].Col,
					})
					nbRottenOranges++
				}
			}
		}
		nbMinutes++
	}
	if nbRottenOranges+nbEmptyCells != len(grid)*len(grid[0]) {
		return -1
	}
	return nbMinutes
}
