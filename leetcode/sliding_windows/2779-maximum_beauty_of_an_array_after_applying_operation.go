package slidingwindows

import "sort"

func Challenge2779(nums []int, k int) int {
	return maximumBeauty(nums, k)
}

// This challenge relates to find the longest subsequence containing elements lies in a suitable spectrum
// Hence, we need to fix the start and find the end.
// Given x, y, where x + k >= y - k
// -> y <= x + 2 * k
// -> we sort the array, and use dynamic sliding windows to look for the valid window
// All elements that lie in the range between x and y would also satisfy the condition

// Why in this challenge, sorting does not affect the result?
// Because in the original array, we look for the LENGTH OF the subsequence.
// Hence, we only need to know how many elements lie in that spectrum
// Therefore, when sorting, we can know exactly the length of the subsequence we looking for
func maximumBeauty(nums []int, k int) int {
	sort.Ints(nums)

	rs := 0
	start, end := 0, 0
	for end < len(nums) {
		// shrink the window until we find a valid subsequence
		for start < len(nums) && nums[end]-nums[start] > 2*k {
			start++
		}
		rs = max(rs, end-start+1)
		end++
	}
	return rs
}

// Initial idea is same as the above solution
// However, when the window reachs the max size, we can make the window doesn't shrink
// Instead, we slide the window to the end. Hence, at the end, the current window size is the length we want
func maximumBeautySimplified(nums []int, k int) int {
	sort.Ints(nums)
	start, end := 0, 0
	for end < len(nums) {
		if nums[end]-nums[start] > 2*k {
			start++
		}
		end++
	}
	return end - start
}
