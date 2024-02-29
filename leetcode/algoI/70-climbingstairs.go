package algoI

// recursive solution with memoization
// 1 <= n <= 45
func ClimbStairs(n int) int {
	mapPositionToNbWays := make([]int, 46)
	return canReachTop(n, 0, mapPositionToNbWays)
}

func canReachTop(n, curIdx int, mapPositionToNbWays []int) int {
	// init the map when needed
	if mapPositionToNbWays[curIdx] != 0 {
		return mapPositionToNbWays[curIdx]
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
