package prefixsum

// relates to Challenge560
func Challenge862(nums []int, k int) int {
	return shortestSubarray(nums, k)
}

func shortestSubarray(nums []int, k int) int {
	rs := len(nums) + 1

	// assume B[i] is the prefix sum of subarray endings at i
	b := make([]int, len(nums))
	b[0] = nums[0]
	// sum[j, i] = B[i] - B[j - 1]
	// -> we need to find biggest j that B[i] - B[j - 1] >= k

	// dequeue tracking for increasing j
	dq := []int{}
	for i := range nums {
		if i > 0 {
			b[i] = b[i-1] + nums[i]
		}
		if b[i] >= k {
			rs = min(rs, i+1)
		}

		// find the biggest j that B[i] - B[j - 1] >= k
		for len(dq) > 0 && b[i]-b[dq[0]] >= k {
			rs = min(rs, i-dq[0])

			// pop first, cause we no longer need it (greedy approach)
			dq = dq[1:]
		}

		// keep the dequeue sorted -> pop in the back before pushing i in the end
		// why pop but not insert i?
		// because if b[i] <= b[dq.last] and i is biggest, the shorter subarray will always start from i, rather than dq.last
		// hence, just remove dq.last
		for len(dq) > 0 && b[i] <= b[dq[len(dq)-1]] {
			// pop last
			dq = dq[:len(dq)-1]
		}

		dq = append(dq, i)
	}
	if rs == len(nums)+1 {
		return -1
	}
	return rs
}
