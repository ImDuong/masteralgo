package algoI

func Combine(n int, k int) [][]int {
	combinations := [][]int{}
	visited := make([]bool, n)
	x := make([]int, k+1)
	tryCombine(1, x, visited, &combinations)
	return combinations
}

func tryCombine(i int, comb []int, visited []bool, combinations *[][]int) {
	for v := comb[i-1] + 1; v <= len(visited); v++ {
		if !visited[v-1] {
			visited[v-1] = true
			comb[i] = v
			if i == len(comb)-1 {
				added := make([]int, len(comb)-1)
				copy(added, comb[1:])
				*combinations = append(*combinations, added)
			} else {
				tryCombine(i+1, comb, visited, combinations)
			}

			visited[v-1] = false
		}
	}
}
