package graph

// Beside coloring method, this challenge can be solved by DSU method
func Challenge785(graph [][]int) bool {
	return isBipartite(graph)
}

// Use coloring method to check if u and graph[u] are divided to 2 different sets
// In addition, nodes in graph[u] must belong to a same set
// Denote:
// 0: not visited
// 1 (blue): node u
// -1 (red): adj of node u
func isBipartite(graph [][]int) bool {
	colors := make([]int, len(graph))
	for i := range graph {
		// if not colored, we set blue for it
		if colors[i] == 0 && !isValidGraph(i, 1, graph, colors) {
			return false
		}
	}
	return true
}

func isValidGraph(node int, checkedColor int, graph [][]int, colors []int) bool {
	if colors[node] != 0 {
		return colors[node] == checkedColor
	}
	colors[node] = checkedColor
	for i := range graph[node] {
		// set the opposite color for the neighbors of node u
		if !isValidGraph(graph[node][i], -checkedColor, graph, colors) {
			return false
		}
	}
	return true
}
