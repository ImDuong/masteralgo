package backtracking

import "sort"

func Challenge90(nums []int) [][]int {
	return subsetsWithDup(nums)
}

func subsetsWithDup(nums []int) [][]int {
	// sort to bring duplicates together
	// so we can skip duplicates later
	sort.Ints(nums)

	rs := [][]int{}
	try(nums, &rs, []int{}, 0)
	return rs
}

func try(nums []int, rs *[][]int, cur []int, numIdx int) {
	copiedCur := make([]int, len(cur))
	copy(copiedCur, cur)
	*rs = append(*rs, copiedCur)
	for i := numIdx; i < len(nums); i++ {
		// for i == numIdx, we always add the number and this is the only one case we don't skip when we see duplicates
		// for i > numIdx, we skip duplicates because we already added the number at numIdx
		if i > numIdx && nums[i] == nums[i-1] {
			continue
		}

		cur = append(cur, nums[i])

		try(nums, rs, cur, i+1)

		// backtrack
		cur = cur[:len(cur)-1]
	}
}
