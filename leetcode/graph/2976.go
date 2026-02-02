package graph

import "math"

func Challenge2976(source string, target string, original []byte, changed []byte, cost []int) int64 {
	return minimumCost(source, target, original, changed, cost)
}

// build graph and find the cost when change a letter x -> y
// then use it to query the min cost
func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	// dist matrix: dist[i][j] is the cost to change i to j
	dist := make([][]int, 26)
	for i := 0; i < 26; i++ {
		dist[i] = make([]int, 26)
		for j := 0; j < 26; j++ {
			if i == j {
				// the diagonal: cost is zero
				dist[i][j] = 0
			} else {
				dist[i][j] = math.MaxInt32
			}
		}
	}

	// set the edges
	for i := range original {
		u := original[i] - 'a'
		v := changed[i] - 'a'
		if cost[i] < dist[u][v] {
			dist[u][v] = cost[i]
		}
	}

	// apply Floyd-Marshall to find the min distance from i to j
	for k := range 26 {
		for i := range 26 {
			for j := range 26 {
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}

	var rs int64
	for i := range source {
		u := source[i] - 'a'
		v := target[i] - 'a'
		if dist[u][v] == math.MaxInt32 {
			return -1
		}
		rs += int64(dist[u][v])
	}
	return rs
}
