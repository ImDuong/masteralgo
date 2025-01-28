package topological_sort

// Kahn's algo to solve
// Why? Because we need to process nodes in the correct order, ensuring each node is encountered after its prerequisites are handled
// In the end, we maintain a list of dependencies for each node to fast queries
// Which mean we take time to preprocess, but fast retrieval
func Challenge1462(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	return checkIfPrerequisite(numCourses, prerequisites, queries)
}

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	// track incoming connections
	inDegree := make([]int, numCourses)

	// track next nodes for each node
	nexts := make([][]int, numCourses)
	for _, pair := range prerequisites {
		inDegree[pair[1]]++
		nexts[pair[0]] = append(nexts[pair[0]], pair[1])
	}

	// topological sort: preprocess to order nodes with inDegree = 0 to traverse first
	q := []int{}
	for i := range inDegree {
		if inDegree[i] == 0 {
			q = append(q, i)
		}
	}

	// track prerequisites of each node (including indirect)
	pre := make([]map[int]bool, numCourses)
	// traverse zero-inDegree nodes and maintain a list of dependencies for each node
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, next := range nexts[cur] {
			if len(pre[next]) == 0 {
				pre[next] = make(map[int]bool)
				pre[next][cur] = true
			} else if _, ok := pre[next][cur]; !ok {
				pre[next][cur] = true
			}

			for indirectPre := range pre[cur] {
				if _, ok := pre[next][indirectPre]; !ok {
					pre[next][indirectPre] = true
				}
			}

			inDegree[next]--
			if inDegree[next] == 0 {
				q = append(q, next)
			}
		}
	}

	rs := make([]bool, len(queries))
	for i := range queries {
		if len(pre[queries[i][1]]) == 0 {
			continue
		}
		if _, ok := pre[queries[i][1]][queries[i][0]]; ok {
			rs[i] = true
		}
	}
	return rs
}
