package linesweep

func Challenge2779(nums []int, k int) int {
	return maximumBeauty(nums, k)
}

// challenge involving finding overlapping ranges -> use line sweep
func maximumBeauty(nums []int, k int) int {
	maxVal := 0
	for i := range nums {
		maxVal = max(nums[i], maxVal)
	}
	line := make([]int, maxVal+1)

	// mark 1 and -1 for starting and ending point of a range
	// hence, number of value in an index will be the number of elements having range overlapped, which be the beauth of the array
	for i := range nums {
		line[max(nums[i]-k, 0)]++

		// mark -1 for nums[i] + k + 1
		// but make sure we don't go out of bound
		if nums[i]+k+1 < len(line) {
			line[nums[i]+k+1]--
		}
	}

	rs := 0
	curSum := 0
	for i := range line {
		curSum += line[i]
		rs = max(rs, curSum)
	}
	return rs
}
