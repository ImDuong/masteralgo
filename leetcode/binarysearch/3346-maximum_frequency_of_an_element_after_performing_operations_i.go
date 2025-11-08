package binarysearch

import (
	"masteralgo/internal/core/utils"
	"sort"
)

func Challenge3346(nums []int, k int, numOperations int) int {
	return maxFrequency(nums, k, numOperations)
}

// bruteforce: for each element a, the answer is how many elements in the range [a - k; a + k]
// -> ans = min(r - l + 1, freq[a] + numOperations)
// where, r and l are right and left bound of range
// and `a` does not necessarily exist in nums!!!
// To speed up the process, we use binary search to look for upperbound and lowerbound
func maxFrequency(nums []int, k int, numOperations int) int {
	sort.Ints(nums)
	freq := make(map[int]int)
	for i := range nums {
		freq[nums[i]]++
	}
	rs := 0
	// because `a` does not necessarily exist in nums, we enumerate in the range
	for a := nums[0]; a <= nums[len(nums)-1]; a++ {
		l := utils.LowerBound(nums, a-k)
		r := utils.UpperBoundNegation(nums, a+k)
		rs = max(rs, min(r-l+1, freq[a]+numOperations))
	}
	return rs
}
