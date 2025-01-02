package twopointers

func Challenge689(nums []int, k int) []int {
	return maxSumOfThreeSubarrays(nums, k)
}

// use 3 pointers and prefix sum
func maxSumOfThreeSubarrays(nums []int, k int) []int {
	p := make([]int, len(nums)+1)
	for i := range nums {
		p[i+1] = p[i] + nums[i]
	}
	left, right := make([]int, len(nums)), make([]int, len(nums))

	// store starting index of subarray (that max sum) ending at i
	for i, curMax := k, p[k]-p[0]; i < len(nums); i++ {
		if curMax < p[i+1]-p[i-k+1] {
			left[i] = i - k + 1
			curMax = p[i+1] - p[i-k+1]
		} else {
			left[i] = left[i-1]
		}
	}

	// store starting index of subarray (that max sum) starts at / after i
	right[len(nums)-k] = len(nums) - k
	for i, curMax := len(nums)-k-1, p[len(nums)]-p[len(nums)-k]; i >= 0; i-- {
		if curMax <= p[i+k]-p[i] {
			right[i] = i
			curMax = p[i+k] - p[i]
		} else {
			right[i] = right[i+1]
		}
	}

	// iterate to find mid region
	maxSum := 0
	rs := make([]int, 3)
	for i := k; i <= len(nums)-2*k; i++ {
		leftIdx := left[i-1]
		rightIdx := right[i+k]
		curSum := p[leftIdx+k] - p[leftIdx] + p[i+k] - p[i] + p[rightIdx+k] - p[rightIdx]
		if curSum > maxSum {
			maxSum = curSum
			rs[0] = leftIdx
			rs[1] = i
			rs[2] = rightIdx
		}
	}
	return rs
}
