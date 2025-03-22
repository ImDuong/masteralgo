package linesweep

func Challenge3355(nums []int, queries [][]int) bool {
	return isZeroArray(nums, queries)
}

func isZeroArray(nums []int, queries [][]int) bool {
	diff := make([]int, len(nums)+1)
	for i := range queries {
		diff[queries[i][0]]++
		diff[queries[i][1]+1]--
	}
	prefixSum := 0
	for i := range nums {
		prefixSum += diff[i]
		if nums[i] > prefixSum {
			return false
		}
	}

	return true
}
