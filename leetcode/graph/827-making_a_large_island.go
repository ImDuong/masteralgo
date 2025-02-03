package graph

func Chalenge827(grid [][]int) int {
	return largestIsland(grid)
}

// Break challenge to 2 subproblems:
// 1. Find islands
// 2. Try to flip each zero to re-calculate area & compare each new area
// Use coloring method to track island size for each island
// In details, we reuse grid to store the island number, where an individual array to store size for each island number
func largestIsland(grid [][]int) int {
	// get island number
	get := func(i, j int) int {
		if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) {
			return grid[i][j]
		}
		return 0
	}

	dirs := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	// color resembles the island number
	// output: the size of island containing cell[i][j]
	var paint func(i, j int, color int) int
	paint = func(i, j int, color int) int {
		// avoid re-painting visited cells or zero-cells
		if get(i, j) != 1 {
			return 0
		}

		// paint the cell with island number
		grid[i][j] = color

		// init with size of 1 for the current cell
		size := 1
		for _, dir := range dirs {
			size += paint(i+dir[0], j+dir[1], color)
		}
		return size
	}

	// step 1: track size of each island, where:
	// index: island number, value: island size
	// the first island number is started from index 2, to differentiate with 0 and 1 cells
	sizes := []int{0, 0}
	hasZero := false
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 0 {
				hasZero = true
				continue
			}

			// new island number is the current length of sizes
			sizes = append(sizes, paint(i, j, len(sizes)))
		}
	}

	if !hasZero {
		return len(grid) * len(grid[0])
	}

	// step 2: flip each zero to re-calculate & compare area
	rs := 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] != 0 {
				continue
			}

			// flip the current cell from 0 to 1
			curArea := 1

			// use map to avoid re-calculate area for a same island
			uniqueIslands := make(map[int]bool)
			for _, dir := range dirs {
				islandNumber := get(i+dir[0], j+dir[1])
				if _, ok := uniqueIslands[islandNumber]; ok {
					continue
				}
				uniqueIslands[islandNumber] = true
				curArea += sizes[islandNumber]
			}

			rs = max(rs, curArea)
		}
	}
	return rs
}
