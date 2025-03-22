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
// in short, find index of the smallest x such that x >= target
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

// equivalents to std::upper_bound in c++
// return the index of the first element in range [first, last)
// - where all elements in this range is strictly greater than the target
// find index of the smallest x such that x > target
func UpperBound(nums []int, target int) int {
	low, high, mid := 0, len(nums)-1, 0
	for low <= high {
		mid = low + (high-low)/2
		if nums[mid] <= target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

// find index of the biggest x such that x < target
func LowerBoundNegation(nums []int, target int) int {
	low, high, mid := 0, len(nums)-1, 0
	for low <= high {
		mid = low + (high-low)/2
		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return high
}

// find index of the biggest x such that x <= target
func UpperBoundNegation(nums []int, target int) int {
	low, high, mid := 0, len(nums)-1, 0
	for low <= high {
		mid = low + (high-low)/2
		if nums[mid] <= target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return high
}

// Usually BinarySearch calculates mid as (low + high) / 2
// but what about when we need to get the ceiling of mid?
// To get ceil, we use the formula: (dividend + divisor - 1) / divisor
// -> in this case, ceil = (low + high + 2 - 1) / 2 = (low + high + 1) / 2
// And, to avoid overflow, we use the formula: low + (high - low + 1) / 2
func GetCeilMid(low, high int) int {
	return low + (high-low+1)/2
}
