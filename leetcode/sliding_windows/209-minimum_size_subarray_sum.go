package slidingwindows

func Challenge209(target int, nums []int) int {
	return minSubArrayLen(target, nums)
}

func minSubArrayLen(target int, nums []int) int {
	curSum := 0
	const MaxInt = int(^uint(0) >> 1)
	rs := MaxInt
	left := 0
	for i := range nums {
		curSum += nums[i]
		for curSum >= target {
			rs = min(rs, i-left+1)
			curSum -= nums[left]
			left++
		}
	}

	if rs == MaxInt {
		return 0
	}
	return rs
}

// use the boilerplates.findSubString template
// as you can see, it's the same with the above solution
func minSubArrayLenUsingTemplate(target int, nums []int) int {
	sum := 0
	// left, right for the boundary of two pointers
	left, right := 0, 0
	const MaxInt = int(^uint(0) >> 1)
	rs := MaxInt

	for ; right < len(nums); right++ {
		sum += nums[right]

		for sum >= target {
			if rs > right-left+1 {
				rs = right - left + 1
			}

			sum -= nums[left]
			left++
		}
	}
	if rs == MaxInt {
		return 0
	}
	return rs
}
