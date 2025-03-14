package boilerplates

import "masteralgo/internal/core/utils"

func lengthOfLIS(nums []int) int {
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
	return lengths[len(lengths)-1]
}
