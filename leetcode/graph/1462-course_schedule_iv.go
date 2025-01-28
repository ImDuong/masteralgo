package graph

// apply Floyd Warshall Algorithm
// The main idea is to check if there’s a path from src to dst by looking at all possible intermediate nodes.
// For each intermediate node, we check if there’s a path from src to that node and a path from that node to dst.
// If both conditions hold, then we can confirm that a path exists between src and dst
func Challenge1462(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	return checkIfPrerequisite(numCourses, prerequisites, queries)
}

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	isPre := make([][]bool, numCourses)
	for i := range isPre {
		isPre[i] = make([]bool, numCourses)
	}
	for i := range prerequisites {
		src, dst := prerequisites[i][0], prerequisites[i][1]
		isPre[src][dst] = true
	}
	for intermediate := 0; intermediate < numCourses; intermediate++ {
		for src := 0; src < numCourses; src++ {
			for dst := 0; dst < numCourses; dst++ {
				isPre[src][dst] = isPre[src][dst] || (isPre[src][intermediate] && isPre[intermediate][dst])
			}
		}
	}

	rs := make([]bool, len(queries))
	for i := range queries {
		rs[i] = isPre[queries[i][0]][queries[i][1]]
	}

	return rs
}
