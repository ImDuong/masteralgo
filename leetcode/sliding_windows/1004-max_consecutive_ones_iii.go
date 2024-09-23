package slidingwindows

func Challenge1004(nums []int, k int) int {
	return longestOnes(nums, k)
}

// The size of window is always the longest ones (brilliant solution from @lee215).
// Intuition:
// - k acts as the indicator whether window size needs to expand or hold
// Steps:
//  1. Set i runs in [0; n - 1], j starts from 0
//     -> The size of window in each loop would be i + 1 - j
//  2. If nums[i] is 1, we need to expand the window size by incrementing i
//     - Otherwise, decrease k by 1, since, we need to flip 0 to 1
//  3. If we ran out of energy to flip (aka k < 0), we need to maintain the window size by incrementing both i and j
//     - Moreover, when nums[j] == 0, we need to increase k by 1. Why?
//     + because after each loop, the window is sliding to the right
//     + -> the window is moving out of the 0 element -> need to recover the k by 1 unit
func longestOnes(nums []int, k int) int {
	i, j := 0, 0
	n := len(nums)
	for ; i < n; i++ {
		if nums[i] == 0 {
			k--
		}
		if k < 0 {
			if nums[j] == 0 {
				k++
			}

			// whenever we ran out energy to flip, we need to slide the window to the right
			// by incrementing both start and end of the window (aka j and i, respectively)
			j++
		}
	}

	// i is incremented already in the last statement of the loop
	// hence, no need to i - j + 1
	return i - j
}
