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
