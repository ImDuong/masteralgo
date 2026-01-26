package graph

import (
	"container/heap"
	"math"
)

func Challenge3650(n int, edges [][]int) int {
	return minCost(n, edges)
}

// Dijkstra: minimum cost from node 0 to node n-1
func minCost(n int, edges [][]int) int {
	// Build adjacency list
	in, out := make([][]Edge, n), make([][]Edge, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		out[u] = append(out[u], Edge{to: v, cost: w})
		in[v] = append(in[v], Edge{to: u, cost: w})
	}

	// Distance array:
	// dist[i] = distance between node 0 to node i
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt64
	}
	dist[0] = 0

	// Priority queue
	pq := &MinHeap{}
	heap.Init(pq)
	heap.Push(pq, &Item{node: 0, dist: 0})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(*Item)
		u := cur.node

		// Skip outdated entries
		if cur.dist > dist[u] {
			continue
		}

		// Early exit
		if u == n-1 {
			return cur.dist
		}

		// normal edges
		for _, e := range out[u] {
			v := e.to
			newDist := cur.dist + e.cost
			if newDist < dist[v] {
				dist[v] = newDist
				heap.Push(pq, &Item{node: v, dist: newDist})
			}
		}

		// reverse edges
		for _, e := range in[u] {
			v := e.to
			newDist := cur.dist + 2*e.cost
			if newDist < dist[v] {
				dist[v] = newDist
				heap.Push(pq, &Item{node: v, dist: newDist})
			}
		}
	}

	// If unreachable
	if dist[n-1] == math.MaxInt64 {
		return -1
	}
	return dist[n-1]
}

// Edge represents a graph edge
type Edge struct {
	to   int
	cost int
}

// Item: represent distance between node 0 and current node in item
type Item struct {
	node int
	dist int
}

// MinHeap implements heap.Interface
type MinHeap []*Item

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].dist < h[j].dist }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(*Item)) }
func (h *MinHeap) Pop() interface{} {
	poppedItem := (*h)[h.Len()-1]
	*h = (*h)[0 : h.Len()-1]
	return poppedItem
}
