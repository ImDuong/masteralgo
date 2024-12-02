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

// equivalents to std::lower_bound in c++
// return the index of the first element in range [first, last)
// - where all elements in this range does not less than target
// - if range is empty, return last
func LowerBound(nums []int, target int) int {
	low, high, mid := 0, len(nums)-1, 0
	for low <= high {
		mid = low + (high-low)/2
		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}
