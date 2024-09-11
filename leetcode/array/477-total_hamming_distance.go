package array

func Challenge477(nums []int) int {
	return totalHammingDistance(nums)
}

// instead of calculate hamming dist of each pair (which is O(n^2))
// we check how many 1-bits by each layer of all numbers
// where layer means each bit by position
func totalHammingDistance(nums []int) int {
	rs := 0
	nbEles := len(nums)
	// nums[i] <= 10^9 -> max number of bits are log2(10^9) + 1 = 31
	for j := 0; j < 31; j++ {
		nbElesHavingBit1 := 0
		for i := 0; i < nbEles; i++ {
			// AND with 1 to get the last bit
			nbElesHavingBit1 += nums[i] & 1

			// shift right
			nums[i] >>= 1
		}
		rs += nbElesHavingBit1 * (nbEles - nbElesHavingBit1)
	}
	return rs
}
