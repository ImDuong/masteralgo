package algoI

// recursive solution with memoization
// 1 <= n <= 45
func ClimbStairs(n int) int {
	mapPositionToNbWays := make([]int, n+1)
	return canReachTop(0, mapPositionToNbWays)
}

func canReachTop(curIdx int, mapPositionToNbWays []int) int {
	// init the map when needed
	if mapPositionToNbWays[curIdx] != 0 {
		return mapPositionToNbWays[curIdx]
	}

	n := len(mapPositionToNbWays) - 1
	if curIdx == n {
		mapPositionToNbWays[curIdx]++
		return mapPositionToNbWays[curIdx]
	}

	climbOptions := []int{1, 2}
	for i := range climbOptions {
		if curIdx+climbOptions[i] <= n {
			mapPositionToNbWays[curIdx] += canReachTop(curIdx+climbOptions[i], mapPositionToNbWays)
		}
	}
	return mapPositionToNbWays[curIdx]
}
