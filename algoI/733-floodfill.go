package algoI

import "masteralgo/pkg/helpers"

func FloodFill(image [][]int, sr int, sc int, color int) [][]int {
	nbRows := len(image)
	nbCols := len(image[0])
	visited := make([]bool, nbRows*nbCols)
	var stack []int
	stack = append(stack, sr*nbCols+sc)
	visited[sr*nbCols+sc] = true
	for len(stack) > 0 {
		cellIdx := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cellRowIdx, cellColIdx := helpers.GetRowColIndexes(cellIdx, nbCols)
		oldColor := image[cellRowIdx][cellColIdx]
		image[cellRowIdx][cellColIdx] = color
		adjCells := helpers.GetAdjacentCellIdx(cellIdx, nbRows, nbCols)
		for i := range adjCells {
			adjRowIdx, adjColIdx := helpers.GetRowColIndexes(adjCells[i], nbCols)
			if image[adjRowIdx][adjColIdx] == oldColor && !visited[adjCells[i]] {
				stack = append(stack, adjCells[i])
				visited[adjCells[i]] = true
			}
		}
	}
	return image
}
