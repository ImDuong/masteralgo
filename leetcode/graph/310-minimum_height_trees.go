package graph

func Challenge310(n int, edges [][]int) []int {
	return findMinHeightTrees(n, edges)
}

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	adj := make([]map[int]bool, n)
	for i := range edges {
		a, b := edges[i][0], edges[i][1]
		if adj[a] == nil {
			adj[a] = make(map[int]bool)
		}
		adj[a][b] = true

		if adj[b] == nil {
			adj[b] = make(map[int]bool)
		}
		adj[b][a] = true
	}
	leaves := []int{}
	for i := range adj {
		if len(adj[i]) == 1 {
			leaves = append(leaves, i)
		}
	}
	for n > 2 {
		n -= len(leaves)
		nextLeaves := []int{}
		for _, curNode := range leaves {
			nextNode := getFirstKey(adj[curNode])
			delete(adj[curNode], nextNode)

			if _, ok := adj[nextNode][curNode]; ok {
				delete(adj[nextNode], curNode)
			}

			if len(adj[nextNode]) == 1 {
				nextLeaves = append(nextLeaves, nextNode)
			}
		}
		leaves = nextLeaves
	}
	return leaves
}
