package dsu

func Chalenge827(grid [][]int) int {
	return largestIsland(grid)
}

// Break challenge to 2 subproblems:
// 1. Find islands
// 2. Try to flip each zero to re-calculate area & compare each new area
func largestIsland(grid [][]int) int {
	n := len(grid)
	dsu := NewIslandDSU(n * n)

	dirs := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	hasZero := false
	// Step 1: Union adjacent 1s cells -> construct islands with calculated area
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 0 {
				hasZero = true
				continue
			}

			// flatten 2D -> 1D
			cur := i*n + j
			for _, dir := range dirs {
				nextx, nexty := i+dir[0], j+dir[1]
				if !(nextx >= 0 && nextx < n && nexty >= 0 && nexty < n && grid[nextx][nexty] == 1) {
					continue
				}
				dsu.Union(cur, nextx*n+nexty)
			}
		}
	}

	// Edge case: if no zeroes found -> all cells are land
	if !hasZero {
		return n * n
	}

	// Step 2: try to flip each zero and re-calculate the new area of merging islands
	rs := 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] != 0 {
				continue
			}
			curArea := 1
			uniqueRoot := make(map[int]bool)
			for _, dir := range dirs {
				nextx, nexty := i+dir[0], j+dir[1]
				if !(nextx >= 0 && nextx < n && nexty >= 0 && nexty < n && grid[nextx][nexty] == 1) {
					continue
				}
				uniqueRoot[dsu.Find(nextx*n+nexty)] = true
			}

			for r := range uniqueRoot {
				curArea += dsu.IslandSize[r]
			}

			rs = max(rs, curArea)
		}
	}
	return rs
}

type IslandDSU struct {
	NbNodes    int
	IslandSize []int
	Parent     []int
}

func NewIslandDSU(nbNodes int) *IslandDSU {
	dsu := IslandDSU{
		NbNodes:    nbNodes,
		IslandSize: make([]int, nbNodes),
		Parent:     make([]int, nbNodes),
	}
	for i := range dsu.IslandSize {
		dsu.IslandSize[i] = 1
	}
	// set itself as parent
	for i := range dsu.Parent {
		dsu.Parent[i] = i
	}

	return &dsu
}

// find the ultimate parent of the node, while compressing the path
func (dsu *IslandDSU) Find(node int) int {
	if dsu.Parent[node] == node {
		return node
	}
	dsu.Parent[node] = dsu.Find(dsu.Parent[node])
	return dsu.Parent[node]
}

// return true if u and v belong to different sets
func (dsu *IslandDSU) Union(u, v int) bool {
	u = dsu.Find(u)
	v = dsu.Find(v)

	if u == v {
		return false
	}

	if dsu.IslandSize[u] > dsu.IslandSize[v] {
		dsu.Parent[v] = u
		dsu.IslandSize[u] += dsu.IslandSize[v]
	} else {
		dsu.Parent[u] = v
		dsu.IslandSize[v] += dsu.IslandSize[u]
	}
	return true
}
