package dynamicprogramming

import "masteralgo/pkg/helpers"

// 1 <= nums.length <= 2 * 10^4
// 1 <= nums[i] <= 10^4
func DeleteAndEarn(nums []int) int {
	buckets := make([]int, 10_001)
	for i := range nums {
		buckets[nums[i]] += nums[i]
	}

	prevTake, prevSkip := 0, 0
	for i := range buckets {
		curTake := prevSkip + buckets[i]
		curSkip := helpers.MaxInt(prevSkip, prevTake)
		prevTake = curTake
		prevSkip = curSkip
	}
	return helpers.MaxInt(prevTake, prevSkip)
}
