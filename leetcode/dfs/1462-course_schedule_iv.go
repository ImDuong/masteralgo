package dfs

// a basic dfs challenge
func Challenge1462(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	return checkIfPrerequisite(numCourses, prerequisites, queries)
}

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	// track prequesites nodes
	pre := make([][]int, numCourses)
	for _, pair := range prerequisites {
		pre[pair[0]] = append(pre[pair[0]], pair[1])
	}

	rs := make([]bool, len(queries))
	for i, pair := range queries {
		q := []int{pair[0]}
		visited := make([]bool, numCourses)
		visited[pair[0]] = true
		target := pair[1]
		isFound := false
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]

			for j := range pre[cur] {
				if pre[cur][j] == target {
					isFound = true
					break
				}
				if visited[pre[cur][j]] {
					continue
				}
				q = append(q, pre[cur][j])
				visited[pre[cur][j]] = true
			}
			if isFound {
				break
			}
		}
		if isFound {
			rs[i] = true
		}
	}
	return rs
}
