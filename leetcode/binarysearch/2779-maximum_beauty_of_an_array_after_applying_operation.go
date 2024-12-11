package binarysearch

import (
	"masteralgo/internal/core/utils"
	"sort"
)

func Challenge2779(nums []int, k int) int {
	return maximumBeauty(nums, k)
}

// This challenge relates to find the longest subsequence containing elements lies in a suitable spectrum
// Hence, we need to fix the start and find the end.
// Given x, y, where x + k >= y - k
// -> y <= x + 2 * k
// -> we sort the array, and binary search to look for y
// All elements that lie in the range between x and y would also satisfy the condition

// Why in this challenge, sorting does not affect the result?
// Because in the original array, we look for the LENGTH OF the subsequence.
// Hence, we only need to know how many elements lie in that spectrum
// Therefore, when sorting, we can know exactly the length of the subsequence we looking for
func maximumBeauty(nums []int, k int) int {
	sort.Ints(nums)

	rs := 0
	for i := range nums {
		j := utils.UpperBoundNegation(nums, nums[i]+2*k)
		rs = max(rs, j-i+1)
	}
	return rs
}
