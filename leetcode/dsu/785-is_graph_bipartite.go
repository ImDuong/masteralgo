package dsu

import "masteralgo/internal/boilerplates"

// Beside using DSU, this challenge can be solved with coloring method
func Challenge785(graph [][]int) bool {
	return isBipartite(graph)
}

// Because graph[u] is an array of nodes that node u is adjacent to
// -> u and nodes in graph[u] must belong to 2 different sets
// where nodes in graph[u] must belong to a same set (since we're looking for bipartite)
// -> hence, we use dsu to union all nodes in graph[u]
// Then, the graph will not be bipartite if there's a node in graph[u] belonging to the same component with u
func isBipartite(graph [][]int) bool {
	dsu := boilerplates.NewDSU(len(graph))
	for i := range graph {
		for j := range graph[i] {
			// if u is in the same set with a node in graph[u] -> not bipartite
			if dsu.Find(i) == dsu.Find(graph[i][j]) {
				return false
			}
			dsu.Union(graph[i][0], graph[i][j])
		}
	}
	return true
}
