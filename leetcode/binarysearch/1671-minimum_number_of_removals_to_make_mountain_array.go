package binarysearch

import (
	"masteralgo/internal/core/utils"
	"math"
	"slices"
)

func Challenge1671(nums []int) int {
	return minimumMountainRemovals(nums)
}

// challenge equivalents to:
// - find max length of strictly increasing subsequence in the half left
// - find max length of strictly decreasing subsequence in the half right
func minimumMountainRemovals(nums []int) int {
	lis := getLISForEachIdx(nums)
	slices.Reverse(nums)

	// to find the max length of strictly decreasing subsequence -> reverse the string, then re-use LIS algo
	lds := getLISForEachIdx(nums)
	// then reverse the obtained lis to get the lds
	slices.Reverse(lds)

	rs := math.MaxInt64
	for i := range nums {
		if lis[i] > 1 && lds[i] > 1 {
			// assume x is the number of ele needed to remove
			// and lis, lds both contains ele at index i.
			// x + lis + lds - 1 = len
			// => x = len - lis - lds + 1
			rs = min(rs, len(nums)-lis[i]-lds[i]+1)
		}
	}
	return rs
}

// use template from boilerplates.lengthOfLIS
func getLISForEachIdx(nums []int) []int {
	lengths := make([]int, len(nums))
	lis := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		lb := utils.LowerBound(lis, nums[i])
		if lb >= len(lis) {
			lis = append(lis, nums[i])
		} else {
			// replace with new value to keep the subsequence strictly increasing
			lis[lb] = nums[i]
		}
		lengths[i] = len(lis)
	}
	return lengths
}
