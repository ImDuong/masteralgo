package dfs

func Challenge684(edges [][]int) []int {
	return findRedundantConnection(edges)
}

// removable edge must be an edge in a cycle
// -> convert challenge to detect cycle in graph
func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	visited := make([]bool, n+1)
	adj := make([][]int, n+1)
	for i := range edges {
		adj[edges[i][0]] = append(adj[edges[i][0]], edges[i][1])
		adj[edges[i][1]] = append(adj[edges[i][1]], edges[i][0])
	}

	// dfs to track parent of each node through a single traversal while finding the start of the cycle
	parent := make([]int, n+1)
	cycleStart := -1
	detectCycle(1, adj, visited, parent, &cycleStart)

	// backtrack to mark all node in the cycle
	cycleNodes := make([]bool, n+1)
	nextNode := cycleStart
	for {
		cycleNodes[nextNode] = true
		nextNode = parent[nextNode]
		if nextNode == cycleStart {
			break
		}
	}

	// check which edge is needed to removed
	for i := len(edges) - 1; i >= 0; i-- {
		a, b := edges[i][0], edges[i][1]
		if cycleNodes[a] && cycleNodes[b] {
			return edges[i]
		}
	}
	return []int{}
}

func detectCycle(cur int, adj [][]int, visited []bool, parent []int, cycleStart *int) {
	visited[cur] = true
	for i := range adj[cur] {
		nextNode := adj[cur][i]
		if visited[nextNode] {
			// avoid re-pointing to its parent (not a cycle actually)
			if nextNode != parent[cur] && *cycleStart == -1 {
				*cycleStart = nextNode
				parent[nextNode] = cur
			}
			continue
		}
		parent[nextNode] = cur
		detectCycle(nextNode, adj, visited, parent, cycleStart)
	}
}
