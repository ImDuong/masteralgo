package graph

import "container/heap"

func MaxProbability(nbNodes int, edges [][]int, succProb []float64, startNode int, endNode int) float64 {
	// Step 1: Construct the graph
	graph := make(map[int][]NodeProb, nbNodes)
	for i := range edges {
		u, v := edges[i][0], edges[i][1]
		graph[u] = append(graph[u], NodeProb{v, succProb[i]})
		graph[v] = append(graph[v], NodeProb{u, succProb[i]})
	}

	// Step 2: Dijkstra's Algorithm using a max-heap (priority queue)
	probGoingToNode := make([]float64, nbNodes)
	probGoingToNode[startNode] = 1

	maxHeap := &MaxHeap{NodeProb{startNode, 1}}
	heap.Init(maxHeap)

	for maxHeap.Len() > 0 {
		current := heap.Pop(maxHeap).(NodeProb)
		u, uProb := current.Node, current.Prob

		if u == endNode {
			return uProb
		}

		if uProb < probGoingToNode[u] {
			continue
		}

		for _, neighbor := range graph[u] {
			v, vProb := neighbor.Node, neighbor.Prob
			newProb := uProb * vProb

			if newProb > probGoingToNode[v] {
				probGoingToNode[v] = newProb
				heap.Push(maxHeap, NodeProb{v, newProb})
			}
		}
	}

	return 0
}

type NodeProb struct {
	Node int
	Prob float64
}

// MaxHeap is a priority queue that prioritizes larger probabilities
type MaxHeap []NodeProb

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	return h[i].Prob > h[j].Prob // Max-heap based on probabilities
}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(NodeProb))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}
