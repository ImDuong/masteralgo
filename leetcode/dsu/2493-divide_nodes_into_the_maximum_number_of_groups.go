package dsu

import "masteralgo/internal/boilerplates"

func Challenge2493(n int, edges [][]int) int {
	return magnificentSets(n, edges)
}

// Challenge can be broken down to 2 subproblems:
// 1. Find connected components: because graph can be disconnected
// 2. For each connected component, we check if it can form valid set. In the end to get the max result
// The 2nd subproblem is actually the extended version of Bipartite: instead of 2 parts, this time it forms n parts
// -> Denote node u and nodes in graph[u],
// we can put u to part A, where nodes in graph[u] will be put to part B
// -> To find the max result, we need to bruteforce each node in a connected component
func magnificentSets(n int, edges [][]int) int {
	dsu := boilerplates.NewDSU(n)
	adj := make([][]int, n)
	for _, pair := range edges {
		adj[pair[0]-1] = append(adj[pair[0]-1], pair[1]-1)
		adj[pair[1]-1] = append(adj[pair[1]-1], pair[0]-1)
		dsu.Union(pair[0]-1, pair[1]-1)
	}
	nbGroupsForComp := make(map[int]int)
	for i := 0; i < n; i++ {
		nbGroups := getNbGroups(adj, i, n)
		if nbGroups == -1 {
			return -1
		}
		rootNode := dsu.Find(i)
		if _, ok := nbGroupsForComp[rootNode]; ok {
			nbGroupsForComp[rootNode] = max(nbGroupsForComp[rootNode], nbGroups)
		} else {
			nbGroupsForComp[rootNode] = nbGroups
		}
	}
	rs := 0
	for i := range nbGroupsForComp {
		rs += nbGroupsForComp[i]
	}
	return rs
}

func getNbGroups(adj [][]int, srcNode int, n int) int {
	// queue for dfs to track layer
	q := []int{srcNode}
	layerSeen := make([]int, n)
	for i := range layerSeen {
		layerSeen[i] = -1
	}
	deepestLayer := 0
	layerSeen[srcNode] = deepestLayer
	for len(q) > 0 {
		for _ = range q {
			cur := q[0]
			q = q[1:]

			for j := range adj[cur] {
				neighbor := adj[cur][j]
				if layerSeen[neighbor] == -1 {
					// not seen -> push it to next layer
					layerSeen[neighbor] = deepestLayer + 1
					q = append(q, neighbor)
					continue
				}
				if layerSeen[neighbor] == deepestLayer {
					// neighbor been visited and in same layer with current
					// -> which means there's an edge in the same group
					return -1
				}
			}
		}
		deepestLayer++
	}
	return deepestLayer
}
