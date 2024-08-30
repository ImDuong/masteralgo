package graph

// Find the STP without the negative-weight edges. There are 3 cases
// Case 1. STP's weight < target: no answers.
// 			Because, STP is always STP without negative-weights even though we modify negative-weight edges
// 			-> STP's weight always smaller than target no matter what
// Case 2. STP's weight = target: modify all negative-weight equal to target
// Case 3. STP's weight > target: we bruteforce change each negative-weight to 1
// For each modification, try to find the STP again. If the STP's weight ≤ target -> the new weight would be 1 + (target - STP's weight)
func ModifiedGraphEdges(n int, edges [][]int, source int, destination int, target int) [][]int {
	// find STP weight without negative-weight edges
	initSTP := GetSTPForPositiveWeights(n, edges, source, destination)
	if initSTP < target {
		return [][]int{}
	}

	isSolvable := false
	if initSTP == target {
		isSolvable = true
	}

	for i := range edges {
		if edges[i][2] > 0 {
			continue
		}

		// if already solved, just make all negative weights equals to target
		if isSolvable {
			edges[i][2] = target
			continue
		}

		// modify each neg weight to 1 and check for STP
		edges[i][2] = 1
		newSTP := GetSTPForPositiveWeights(n, edges, source, destination)

		if newSTP <= target {
			isSolvable = true
			edges[i][2] = 1 + target - newSTP
		}
	}
	if isSolvable {
		return edges
	}
	return [][]int{}
}

// use Dijikstra to find STP
// output: the STP's weight
func GetSTPForPositiveWeights(nbNodes int, edges [][]int, source, destination int) int {
	const inf int = 2e9
	g := make([][]int, nbNodes)
	dist := make([]int, nbNodes)
	visited := make([]bool, nbNodes)
	for i := range g {
		g[i] = make([]int, nbNodes)
		for j := range g[i] {
			g[i][j] = inf
		}
		dist[i] = inf
	}
	dist[source] = 0

	for i := range edges {
		if edges[i][2] == -1 {
			continue
		}
		g[edges[i][0]][edges[i][1]], g[edges[i][1]][edges[i][0]] = edges[i][2], edges[i][2]
	}

	for i := 0; i < nbNodes; i++ {
		nearestUnvisitedNode := -1

		// find the nearest (to source) univisted node
		for j := 0; j < nbNodes; j++ {
			if !visited[j] && (nearestUnvisitedNode == -1 || dist[nearestUnvisitedNode] > dist[j]) {
				nearestUnvisitedNode = j
			}
		}

		visited[nearestUnvisitedNode] = true

		// update distances to all other nodes
		for j := 0; j < nbNodes; j++ {
			dist[j] = min(dist[j], dist[nearestUnvisitedNode]+g[nearestUnvisitedNode][j])
		}
	}
	return dist[destination]
}
