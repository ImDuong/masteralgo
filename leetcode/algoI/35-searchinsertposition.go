package algoI

func SearchInsert(nums []int, target int) int {
	if nums[0] > target {
		return 0
	}
	lowB := 0
	upB := len(nums) - 1
	var jmp int
	for upB >= lowB {
		jmp = (upB + lowB) / 2
		if nums[jmp] < target {
			lowB = jmp + 1
		} else if nums[jmp] > target {
			upB = jmp - 1
		} else {
			return jmp
		}
	}
	return (upB+lowB)/2 + 1
}
