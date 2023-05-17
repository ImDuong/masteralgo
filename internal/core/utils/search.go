package utils

func BinarySearch(nums []int, startIdx, endIdx, target int) int {
	for endIdx >= startIdx {
		if nums[startIdx] == target {
			return startIdx
		}
		if nums[endIdx] == target {
			return endIdx
		}
		jmp := (endIdx + startIdx) / 2
		if nums[jmp] < target {
			startIdx = jmp + 1
		} else if nums[jmp] > target {
			endIdx = jmp - 1
		} else {
			return jmp
		}
	}
	return -1
}
