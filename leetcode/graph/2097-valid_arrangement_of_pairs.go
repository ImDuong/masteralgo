package graph

func Challenge2097(pairs [][]int) [][]int {
	return validArrangement(pairs)
}

// challenge similars to find an Eulerian path, where all edges are visited once only
// use postorder DFS (runs with O(N + E) time)
// below is the recursive postorder DFS
func validArrangement(pairs [][]int) [][]int {
	adjacents := make(map[int][]int)
	outDegree := make(map[int]int)
	inDegree := make(map[int]int)
	for i := range pairs {
		adjacents[pairs[i][0]] = append(adjacents[pairs[i][0]], pairs[i][1])
		outDegree[pairs[i][0]]++
		inDegree[pairs[i][1]]++
	}

	// find starting node
	startNode := pairs[0][0]
	for i := range outDegree {
		if outDegree[i] == inDegree[i]+1 {
			startNode = i
			break
		}
	}

	paths := []int{}

	var visit func(node int)
	visit = func(node int) {
		for len(adjacents[node]) != 0 {
			nextNode := adjacents[node][0]
			adjacents[node] = adjacents[node][1:]
			visit(nextNode)
		}
		paths = append(paths, node)
	}

	visit(startNode)

	rs := make([][]int, len(pairs))
	cnt := 0
	for i := len(paths) - 1; i > 0; i-- {
		rs[cnt] = []int{paths[i], paths[i-1]}
		cnt++
	}
	return rs
}

// postorder DFS using stack for iterative: Hierholdzer's algo
func validArrangementWithIterativeDFS(pairs [][]int) [][]int {
	adjacents := make(map[int][]int)
	outDegree := make(map[int]int)
	inDegree := make(map[int]int)
	for i := range pairs {
		adjacents[pairs[i][0]] = append(adjacents[pairs[i][0]], pairs[i][1])
		outDegree[pairs[i][0]]++
		inDegree[pairs[i][1]]++
	}

	// find starting node
	startNode := pairs[0][0]
	for i := range outDegree {
		if outDegree[i] == inDegree[i]+1 {
			startNode = i
			break
		}
	}

	paths := []int{}

	// stacks for postorder DFS
	stk := []int{startNode}
	for len(stk) > 0 {
		curNode := stk[len(stk)-1]
		if len(adjacents[curNode]) > 0 {
			// pop the next node from adjacents & add to the stack
			nextNode := adjacents[curNode][len(adjacents[curNode])-1]
			tmp := adjacents[curNode]
			tmp = tmp[:len(tmp)-1]
			adjacents[curNode] = tmp
			stk = append(stk, nextNode)
		} else {
			paths = append(paths, curNode)
			stk = stk[:len(stk)-1]
		}
	}

	rs := make([][]int, len(pairs))
	cnt := 0
	for i := len(paths) - 1; i > 0; i-- {
		rs[cnt] = []int{paths[i], paths[i-1]}
		cnt++
	}
	return rs
}
