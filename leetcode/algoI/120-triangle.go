package algoI

func MinimumTotal(triangle [][]int) int {
	minPath := make([][]int, len(triangle))
	for i := range minPath {
		minPath[i] = make([]int, len(triangle[i]))

		// init for setting threshold
		for j := range minPath[i] {
			minPath[i][j] = (len(triangle) - 1 - i) * -10000
		}
	}
	return getMinimumTotalRecursive(triangle, 0, 0, minPath)
}

func getMinimumTotalRecursive(triangle [][]int, Pr, Pc int, minPath [][]int) int {
	if minPath[Pr][Pc] != (len(triangle)-1-Pr)*-10000 {
		return minPath[Pr][Pc]
	}

	if Pr == len(triangle)-1 {
		minPath[Pr][Pc] = triangle[Pr][Pc]
		return minPath[Pr][Pc]
	}

	routeOptions := []int{0, 1}
	// upper bound: 10^4 * 200 + 1
	minPathTotal := 2000001
	for i := range routeOptions {
		// prune
		if minPathTotal < minPath[Pr][Pc] {
			continue
		}

		minPathI := getMinimumTotalRecursive(triangle, Pr+1, Pc+routeOptions[i], minPath)
		if minPathTotal > minPathI {
			minPathTotal = minPathI
		}
	}
	minPath[Pr][Pc] = minPathTotal + triangle[Pr][Pc]
	return minPath[Pr][Pc]
}
