package dfs

// Viewing problem as matching nodes as disjoin sets
// For each disjoin set, number of removable stones is equal to number of edge minus 1
// -> total # removable stones = # stones - (# disjoin sets * 1)
// Solution to find disjoin sets:
// We break 2d array to 1d array (with length as 2n. Where n is the upper bound of xi, yi)
// , where x coordinates in [0, n - 1], y coord in [n, 2n - 1]
// , where each element in 1d array is a coordination
// Because stones are connected when they share the same row / column
// -> try to unify each x coordinate in [0, n - 1] with equivalent y coord in [n, 2n - 1]
func RemoveStones(stones [][]int) int {
	// 0 <= xi, yi <= 10^4, hence, n is 10001
	nbNodes := 10001
	nodeParents := make([]int, nbNodes*2)
	// make each node as parent of itself
	for i := range nodeParents {
		nodeParents[i] = i
	}

	// unify graph of each x coordinate in the first half to the equivalent graph in the other half
	for _, stone := range stones {
		nodeParents[findRoot(stone[0], nodeParents)] = findRoot(stone[1]+nbNodes, nodeParents)
	}

	// get root of each disjoin sets to find out how many sets there are
	disjoinSets := make(map[int]bool)
	for _, stone := range stones {
		disjoinSets[findRoot(stone[0], nodeParents)] = true
	}
	return len(stones) - len(disjoinSets)
}

func findRoot(node int, parent []int) int {
	if parent[node] != node {
		parent[node] = findRoot(parent[node], parent)
	}
	return parent[node]
}
