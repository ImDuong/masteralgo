package main

import (
	"fmt"
	"os"
)

type Cell struct {
	Row, Col int
}

var directions = []Cell{
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

func getNbSurvivors(grid [][]int) int {
	nbMinutes := -1
	nbZombies := 0
	nbEmptyCells := 0
	queue := []Cell{}

	// find all rotten oranges first
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 2 {
				queue = append(queue, Cell{
					Row: i,
					Col: j,
				})
				nbZombies++
			} else if grid[i][j] == 0 {
				nbEmptyCells++
			}
		}
	}
	if nbZombies+nbEmptyCells == len(grid)*len(grid[0]) {
		return 0
	}

	// process layers by layers
	for len(queue) > 0 {
		// pop all rotten oranges in once
		poppedCells := make([]Cell, len(queue))
		i := 0
		for len(queue) > 0 {
			var poppedCell Cell
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
					queue = append(queue, Cell{
						Row: adjCells[i].Row,
						Col: adjCells[i].Col,
					})
					nbZombies++
				}
			}
		}
		nbMinutes++
	}
	if nbZombies+nbEmptyCells != len(grid)*len(grid[0]) {
		return -1
	}
	return nbMinutes
}

func getAdjCells(curCell Cell, nbRows, nbCols int) []Cell {
	cells := []Cell{}
	for i := range directions {
		row := curCell.Row + directions[i].Row
		col := curCell.Col + directions[i].Col
		if !(row < 0 || row >= nbRows || col < 0 || col >= nbCols) {
			cells = append(cells, Cell{
				Row: row,
				Col: col,
			})
		}
	}
	return cells
}

func handleError(err error) {
	fmt.Println("error: ", err)
	os.Exit(0)
}

func main() {
	var nbRows, nbCols int
	_, err := fmt.Scanf("%d", &nbRows)
	if err != nil {
		handleError(err)
	}
	_, err = fmt.Scanf("%d", &nbCols)
	if err != nil {
		handleError(err)
	}
	grid := make([][]int, nbRows)
	for i := range grid {
		grid[i] = make([]int, nbCols)
		for j := range grid[i] {
			_, err = fmt.Scanf("%d", &grid[i][j])
			if err != nil {
				handleError(err)
			}
		}
	}
	fmt.Println(getNbSurvivors(grid))
}
