package dsu

import (
	"masteralgo/internal/boilerplates"
)

func Challenge2685(n int, edges [][]int) int {
	return countCompleteComponents(n, edges)
}

func countCompleteComponents(n int, edges [][]int) int {
	dsu := boilerplates.NewDSUWithEdges(n)
	for i := range edges {
		dsu.Union(edges[i][0], edges[i][1])
	}
	visited := make(map[int]bool)
	rs := 0
	for i := 0; i < n; i++ {
		parent := dsu.Find(i)
		if visited[parent] {
			continue
		}

		// assume n is the number of nodes in the component
		// if the component is complete, it must have n*(n-1)/2 edges
		// proof: each node must connect to n-1 other nodes -> (n - 1) + (n - 2) + ... + 1 = n*(n-1)/2
		if dsu.Rank[parent]*(dsu.Rank[parent]-1)/2 == dsu.Edges[parent] {
			rs++
		}
		visited[parent] = true
	}
	return rs
}
