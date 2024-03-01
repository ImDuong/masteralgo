package dynamicprogramming

import "masteralgo/pkg/helpers"

// 2 <= cost.length <= 1000
// 0 <= cost[i] <= 999
func MinCostClimbingStairs(cost []int) int {
	var first, second int = cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		curMinCost := helpers.MinInt(first, second) + cost[i]
		first = second
		second = curMinCost
	}
	return helpers.MinInt(first, second)
}

func MinCostClimbingStairsWithMemoiz(cost []int) int {
	cost = append(cost, 0)
	minCost := make([]int, len(cost))
	for i := range minCost {
		minCost[i] = -1
	}
	minCost[0] = cost[0]
	minCost[1] = helpers.MinInt(minCost[0]+cost[1], cost[1])
	return calMinCostClimbingStairs(len(minCost)-1, minCost, cost)
}

func calMinCostClimbingStairs(curIdx int, dp []int, cost []int) int {
	if dp[curIdx] != -1 {
		return dp[curIdx]
	}
	first, second := 1000, 1000
	if curIdx-1 >= 0 {
		first = calMinCostClimbingStairs(curIdx-1, dp, cost)
	}
	if curIdx-2 >= 0 {
		second = calMinCostClimbingStairs(curIdx-2, dp, cost)
	}
	dp[curIdx] = helpers.MinInt(first, second) + cost[curIdx]
	return dp[curIdx]
}
