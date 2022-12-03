package algoI

import "masteralgo/pkg/helpers"

func MaxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	nbRows := len(grid)
	nbCols := len(grid[0])
	nbCells := nbRows * nbCols
	visited := make([]bool, nbCells)
	var stack []int
	for i := 0; i < nbCells; i++ {
		if visited[i] {
			continue
		}
		stack = append(stack, i)
		visited[i] = true
		curArea := 0
		for len(stack) > 0 {
			cellIdx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			cellRowIdx, cellColIdx := helpers.GetRowColIndexes(cellIdx, nbCols)
			curArea += grid[cellRowIdx][cellColIdx]
			if grid[cellRowIdx][cellColIdx] != 0 {
				adjCells := helpers.GetAdjacentCellIdx(cellIdx, nbRows, nbCols)
				for j := range adjCells {
					adjRowIdx, adjColIdx := helpers.GetRowColIndexes(adjCells[j], nbCols)
					if grid[adjRowIdx][adjColIdx] > 0 && !visited[adjCells[j]] {
						stack = append(stack, adjCells[j])
						visited[adjCells[j]] = true
					}
				}
			}
		}
		if maxArea < curArea {
			maxArea = curArea
		}
	}
	return maxArea
}
