package dynamicprogramming

var MAXVALUE = 10000

// MinimumTotal
//
// 2
// 3 4
// 6 5 7
// 4 1 8 3
func MinimumTotal(triangle [][]int) int {
	dp := make([]int, len(triangle))
	for i := range dp {
		dp[i] = MAXVALUE
	}
	layerth := len(dp) - 1
	minCost := MAXVALUE
	for j := 0; j < len(dp); j++ {
		minCost = min(minCost, recMinimumTotal(triangle, dp, layerth, j))
	}
	return minCost
}

// pivot = 0, dp[3] = x
// pivot = 1, dp[3] = y
// in row 3, pivot: 0 -> 3
// dp[3] = min(dp[3], T[3][pivot] + dp[2])
func recMinimumTotal(triangle [][]int, dp []int, layerth, pivot int) int {
	if layerth == 0 {
		dp[0] = triangle[0][0]
		return dp[0]
	}

	start := max(0, pivot-1)
	end := pivot
	if layerth == pivot-1 {
		end = layerth
	}
	for j := start; j <= end; j++ {
		dp[layerth] = triangle[layerth][j] + recMinimumTotal(triangle, dp, layerth-1, j)
	}
	return dp[layerth]
}
