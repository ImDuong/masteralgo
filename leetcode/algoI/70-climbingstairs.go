package algoI

// recursive solution with memoization
func ClimbStairs(n int) int {
	mapPositionToNbWays := make(map[int]int)
	return canReachTop(n, 0, mapPositionToNbWays)
}

func canReachTop(n, curIdx int, mapPositionToNbWays map[int]int) int {
	// init the map when needed
	if _, ok := mapPositionToNbWays[curIdx]; ok {
		return mapPositionToNbWays[curIdx]
	} else {
		mapPositionToNbWays[curIdx] = 0
	}

	if curIdx == n {
		mapPositionToNbWays[curIdx]++
		return mapPositionToNbWays[curIdx]
	}

	climbOptions := []int{1, 2}
	for i := range climbOptions {
		if curIdx+climbOptions[i] <= n {
			mapPositionToNbWays[curIdx] += canReachTop(n, curIdx+climbOptions[i], mapPositionToNbWays)
		}
	}
	return mapPositionToNbWays[curIdx]
}
