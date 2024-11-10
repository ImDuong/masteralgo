package slidingwindows

func Challenge3097(nums []int, k int) int {
	return minimumSubarrayLength(nums, k)
}

// Intuition: the array contains a special subarray will also be a special array
// -> this also means, when we reach a special array, the last element of this array must exist in the special subarray (if there is)
// -> when we find an array is special array, we can
// remove each element one by one from left to right to find out a smaller special subarray

// Approach: use dynamic sliding windows to find the special array
// 1. Find a window that ORvalue >= k
// 2. Shrink the window until ORvalue < k
// 3. Freeze the left pivot and continue expand the window as in step 1

// Tactics:
// 1. 0 <= nums[i] <= 10^9, hence we need log2(10^9) + 1 = 30 bits
// 2. Use bitcount to count how each element contribute to all 30 bits in the Orvalue
// 3. Nullify step: to remove an element's contribution from the Orvalue, we need to decrease its bitcount for Orvalue's 30 bits.
// 3.1. If the bitcount in a bit-th is zero, we flip that bit in Orvalue

// Time Complexity:
// - all elements are visited twice only -> O(n)
// Space complexity:
// - requires 30-long array -> O(1)
func minimumSubarrayLength(nums []int, k int) int {
	if k == 0 {
		return 1
	}
	const MaxInt = int(^uint(0) >> 1)
	rs := MaxInt
	// 30 bits track for Orvalue
	bitCounts := make([]int, 30)
	Orvalue := 0

	// i, j are right and left pivot of the sliding window
	for i, j := 0, 0; i < len(nums); i++ {
		Orvalue |= nums[i]
		for idx := range bitCounts {
			if nums[i]&(1<<idx) != 0 {
				bitCounts[idx]++
			}
		}

		// if the window is not valid, just continue
		if Orvalue < k {
			continue
		}

		// now a valid array is found, it's time to shrink the window
		for ; j <= i && Orvalue >= k; j++ {
			for idx := range bitCounts {
				if nums[j]&(1<<idx) == 0 {
					// skip if bit in idx-th does not contribute
					continue
				}

				bitCounts[idx]--
				if bitCounts[idx] == 0 {
					// flip the idx-th bit in Orvalue
					// in go: we does not have NOT, hence we use 1 XOR X instead
					Orvalue &= (^(1 << idx))
				}
			}

			rs = min(rs, i-j+1)
		}
	}
	if rs != MaxInt {
		return rs
	}
	return -1
}
