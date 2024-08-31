package array

func Challenge1828(points [][]int, queries [][]int) []int {
	return countPoints(points, queries)
}

func countPoints(points [][]int, queries [][]int) []int {
	rs := make([]int, len(queries))
	for j := range queries {
		nbPts := 0
		for i := range points {
			xCoor := points[i][0]
			yCoor := points[i][1]
			if (queries[j][0]-xCoor)*(queries[j][0]-xCoor)+(queries[j][1]-yCoor)*(queries[j][1]-yCoor) <= queries[j][2]*queries[j][2] {
				nbPts += 1
			}
		}

		rs[j] = nbPts
	}
	return rs
}
