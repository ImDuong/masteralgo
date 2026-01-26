package graph

import (
	"container/heap"
	"math"
	"sort"
)

func Challenge3651(grid [][]int, k int) int {
	// return minCost3651(grid, k)
	// return minCost3651WithPrune(grid, k)
	return minCost3651WithFlatten(grid, k)
}

// Optimized Dijkstra solution with flattened dist array
// Take a look at minCost3651WithPrune for explanation of the algorithm
func minCost3651WithFlatten(grid [][]int, k int) int {
	m := len(grid)
	n := len(grid[0])
	nbCells := m * n

	if k > nbCells {
		k = nbCells
	}

	// dist[t][i][j] = min cost to reach (i, j) from (0, 0) having t teleported used
	// -> flaten to dist[t * m * n + i * n + j]
	dist := make([]int, (k+1)*m*n)
	for i := range dist {
		dist[i] = math.MaxInt32
	}

	// sort cells -> we have a place to look for candidates to teleport because we can teleport from (i, j) to (x, y) such that grid[x][y] <= grid[i][j]
	cells := make([]Cell3651, 0, nbCells)
	for i := range m {
		for j := range n {
			cells = append(cells, Cell3651{i, j, grid[i][j]})
		}
	}
	sort.Slice(cells, func(i, j int) bool {
		return cells[i].value < cells[j].value
	})

	// for each teleport layer, track how many cells are unlocked
	unlockIdx := make([]int, k+1)

	pq := &MinHeap3651{}
	heap.Init(pq)

	getDistIdx := func(t, i, j int) int {
		return t*nbCells + i*n + j
	}

	// starting point
	dist[getDistIdx(0, 0, 0)] = 0
	heap.Push(pq, &Item3651{0, 0, 0, 0})
	dirs := [][2]int{
		{0, 1}, // right
		{1, 0}, // down
	}

	rs := math.MaxInt32
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(*Item3651)

		i, j := cur.i, cur.j
		if cur.dist > dist[getDistIdx(cur.t, i, j)] {
			continue
		}

		if cur.dist >= rs {
			continue
		}

		if i == m-1 && j == n-1 {
			rs = cur.dist
			continue
		}

		// normal moves (in the same teleport layer)
		for _, dir := range dirs {
			ni, nj := i+dir[0], j+dir[1]
			if ni < m && nj < n {
				newDist := cur.dist + grid[ni][nj]
				if newDist < dist[getDistIdx(cur.t, ni, nj)] {
					dist[getDistIdx(cur.t, ni, nj)] = newDist
					heap.Push(pq, &Item3651{ni, nj, cur.t, newDist})
				}
			}
		}

		// teleport moves (next layer)
		if cur.t < k {
			idx := unlockIdx[cur.t+1]
			for idx < len(cells) && cells[idx].value <= grid[i][j] {
				x, y := cells[idx].i, cells[idx].j
				if cur.dist < dist[getDistIdx(cur.t+1, x, y)] {
					dist[getDistIdx(cur.t+1, x, y)] = cur.dist
					heap.Push(pq, &Item3651{x, y, cur.t + 1, cur.dist})
				}
				idx++
			}

			unlockIdx[cur.t+1] = idx
		}
	}
	return rs
}

// Pruned Dijkstra solution
// Take a look at minCost3651 for explanation of the algorithm
func minCost3651WithPrune(grid [][]int, k int) int {
	m := len(grid)
	n := len(grid[0])

	// dist[t][i][j] = min cost to reach (i, j) from (0, 0) having t teleported used
	dist := make([][][]int, k+1)
	for t := range k + 1 {
		dist[t] = make([][]int, m)
		for i := range dist[t] {
			dist[t][i] = make([]int, n)
			for j := range dist[t][i] {
				dist[t][i][j] = math.MaxInt32
			}
		}
	}

	// sort cells
	cells := make([]Cell3651, 0, m*n)
	for i := range m {
		for j := range n {
			cells = append(cells, Cell3651{i, j, grid[i][j]})
		}
	}

	sort.Slice(cells, func(i, j int) bool {
		return cells[i].value < cells[j].value
	})

	// for each teleport layer, track how many cells are unlocked
	unlockIdx := make([]int, k+1)

	pq := &MinHeap3651{}
	heap.Init(pq)

	// starting point
	dist[0][0][0] = 0
	heap.Push(pq, &Item3651{0, 0, 0, 0})

	dirs := [][2]int{
		{0, 1}, // right
		{1, 0}, // down
	}

	rs := math.MaxInt32
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(*Item3651)

		i, j := cur.i, cur.j
		if cur.dist > dist[cur.t][i][j] {
			continue
		}

		// prune: when distance from (0, 0) to current node is gte than distance from (0, 0) to (m - 1, n - 1) -> we dont need to go further with this node
		if cur.dist >= rs {
			continue
		}

		if i == m-1 && j == n-1 {
			// with using priority queue, we're traversing the most optimized path
			rs = cur.dist
			continue
		}

		// normal moves (in the same teleport layer)
		for _, dir := range dirs {
			ni, nj := i+dir[0], j+dir[1]
			if ni < m && nj < n {
				newDist := cur.dist + grid[ni][nj]
				if newDist < dist[cur.t][ni][nj] {
					dist[cur.t][ni][nj] = newDist
					heap.Push(pq, &Item3651{ni, nj, cur.t, newDist})
				}
			}
		}

		// teleport moves (next layer)
		if cur.t < k {
			idx := unlockIdx[cur.t+1]
			for idx < len(cells) && cells[idx].value <= grid[i][j] {
				x, y := cells[idx].i, cells[idx].j
				if cur.dist < dist[cur.t+1][x][y] {
					dist[cur.t+1][x][y] = cur.dist
					heap.Push(pq, &Item3651{x, y, cur.t + 1, cur.dist})
				}
				idx++
			}

			unlockIdx[cur.t+1] = idx
		}
	}
	return rs
}

// This solution gives TLE on leetcode
// Algorithm: Dijkstra with layers (each layer represents number of teleports used)
// Each time we use a teleport, we can only teleport to cells with value <= current cell value
func minCost3651(grid [][]int, k int) int {
	m := len(grid)
	n := len(grid[0])

	// dist[t][i][j] = min cost to reach (i, j) from (0, 0) having t teleported used
	dist := make([][][]int, k+1)
	for t := range k + 1 {
		dist[t] = make([][]int, m)
		for i := range dist[t] {
			dist[t][i] = make([]int, n)
			for j := range dist[t][i] {
				dist[t][i][j] = math.MaxInt32
			}
		}
	}

	// sort cells -> we have a place to look for candidates to teleport because we can teleport from (i, j) to (x, y) such that grid[x][y] <= grid[i][j]
	cells := make([]Cell3651, 0, m*n)
	for i := range m {
		for j := range n {
			cells = append(cells, Cell3651{i, j, grid[i][j]})
		}
	}
	sort.Slice(cells, func(i, j int) bool {
		return cells[i].value < cells[j].value
	})

	// for each teleport layer, track how many cells are unlocked
	unlockIdx := make([]int, k+1)

	pq := &MinHeap3651{}
	heap.Init(pq)

	// starting point
	dist[0][0][0] = 0
	heap.Push(pq, &Item3651{0, 0, 0, 0})

	dirs := [][2]int{
		{0, 1}, // right
		{1, 0}, // down
	}

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(*Item3651)

		i, j := cur.i, cur.j
		if cur.dist > dist[cur.t][i][j] {
			continue
		}

		if i == m-1 && j == n-1 {
			return cur.dist
		}

		// normal moves (in the same teleport layer)
		for _, dir := range dirs {
			ni, nj := i+dir[0], j+dir[1]
			if ni < m && nj < n {
				newDist := cur.dist + grid[ni][nj]
				if newDist < dist[cur.t][ni][nj] {
					dist[cur.t][ni][nj] = newDist
					heap.Push(pq, &Item3651{ni, nj, cur.t, newDist})
				}
			}
		}

		// teleport moves (next layer)
		if cur.t < k {
			idx := unlockIdx[cur.t+1]
			for idx < len(cells) && cells[idx].value <= grid[i][j] {
				x, y := cells[idx].i, cells[idx].j
				if cur.dist < dist[cur.t+1][x][y] {
					dist[cur.t+1][x][y] = cur.dist
					heap.Push(pq, &Item3651{x, y, cur.t + 1, cur.dist})
				}
				idx++
			}

			unlockIdx[cur.t+1] = idx
		}
	}

	rs := math.MaxInt32
	for t := range k + 1 {
		rs = min(rs, dist[t][m-1][n-1])
	}

	return rs
}

type Cell3651 struct {
	i, j  int
	value int
}

type Item3651 struct {
	i, j int
	t    int // nb of teleports used
	dist int
}

type MinHeap3651 []*Item3651

func (h MinHeap3651) Len() int           { return len(h) }
func (h MinHeap3651) Less(a, b int) bool { return h[a].dist < h[b].dist }
func (h MinHeap3651) Swap(a, b int)      { h[a], h[b] = h[b], h[a] }
func (h *MinHeap3651) Push(x interface{}) {
	*h = append(*h, x.(*Item3651))
}

func (h *MinHeap3651) Pop() interface{} {
	poppedItem := (*h)[h.Len()-1]
	*h = (*h)[0 : h.Len()-1]
	return poppedItem
}
