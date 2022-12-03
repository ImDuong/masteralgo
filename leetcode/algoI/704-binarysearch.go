package algoI

// runtime: 52ms, memory 6.6MB
func BinarySearch(nums []int, target int) int {
	lowB := 0
	upB := len(nums) - 1
	var i int
	for upB >= lowB {
		if nums[lowB] == target {
			return lowB
		}
		if nums[upB] == target {
			return upB
		}
		jmp := int((upB - lowB) / 2)
		if jmp == 0 {
			return -1
		}
		i = jmp + lowB
		if nums[i] < target {
			lowB = i + 1
		} else if nums[i] > target {
			upB = i - 1
		} else {
			return i
		}
	}
	return -1
}

// optimize the formula: however the same?
// runtime: 30ms, memory 6.7MB
func BinarySearchV2(nums []int, target int) int {
	lowB := 0
	upB := len(nums) - 1
	for upB >= lowB {
		if nums[lowB] == target {
			return lowB
		}
		if nums[upB] == target {
			return upB
		}
		jmp := (upB + lowB) / 2
		if nums[jmp] < target {
			lowB = jmp + 1
		} else if nums[jmp] > target {
			upB = jmp - 1
		} else {
			return jmp
		}
	}
	return -1
}
