package graph

// similar to Challenge 310
func Challenge3203(edges1 [][]int, edges2 [][]int) int {
	return minimumDiameterAfterMerge(edges1, edges2)
}

// find the diameter of each tree -> find the min diameter after merge
func minimumDiameterAfterMerge(edges1 [][]int, edges2 [][]int) int {
	n, m := len(edges1)+1, len(edges2)+1
	d1 := getDiameter(n, edges1)
	d2 := getDiameter(m, edges2)
	return max(max(d1, d2), (d1+2-1)/2+1+(d2+2-1)/2)
}

// traverse from leaves to find the minimum height trees -> find the diameter
func getDiameter(n int, edges [][]int) int {
	if n == 1 {
		return 0
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
	layerCount := 0
	for n > 2 {
		layerCount++
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
	diameter := 2 * layerCount
	if len(leaves) == 2 {
		// count the path between 2 leaves
		diameter = 2*layerCount + 1
	}
	return diameter
}

func getFirstKey(m map[int]bool) int {
	for k := range m {
		return k
	}
	return 0
}
