package graph

import (
	"container/heap"
	"masteralgo/internal/boilerplates"
	"slices"
)

func Challenge332(tickets [][]string) []string {
	return findItinerary(tickets)
}

// same with challenge 2097
// use postorder DFS iterative
func findItinerary(tickets [][]string) []string {
	// adjacents for storing edges
	adj := make(map[string]*boilerplates.StringMinHeap)
	outDegree := make(map[string]int)
	inDegree := make(map[string]int)

	for i := range tickets {
		src, dst := tickets[i][0], tickets[i][1]
		if _, ok := adj[src]; !ok {
			adj[src] = &boilerplates.StringMinHeap{}
			heap.Init(adj[src])
		}
		heap.Push(adj[src], dst)

		outDegree[src]++
		inDegree[dst]++
	}

	paths := []string{}
	// stack for DFS iterative
	stk := []string{"JFK"}
	for len(stk) != 0 {
		cur := stk[len(stk)-1]
		if adj[cur] != nil && (*adj[cur]).Len() > 0 {
			// add the smallest first, because it will be the last to traverse. Then when we reverse the paths, the smallest will appear first
			next := adj[cur].Top()
			heap.Pop(adj[cur])
			stk = append(stk, next)
		} else {
			stk = stk[:len(stk)-1]
			paths = append(paths, cur)
		}
	}

	slices.Reverse(paths)
	return paths
}
