package string_challenges

func Challenge1545(n int, k int) byte {
	return findKthBit(n, k)
}

func findKthBit(n int, k int) byte {
	// length of S is 2^n - 1
	len := (1 << n) - 1
	invertCnt := 0
	for k > 1 {
		// index starts from 1 instead of 0
		middle := (len + 1) / 2

		if middle == k {
			// middle always 1, hence if invertCnt is even, means original kth bit is "1" too
			if invertCnt%2 == 0 {
				return '1'
			}
			return '0'
		}

		if k > middle {
			k = len + 1 - k
			invertCnt++
		}

		// shorten length
		len /= 2
	}
	// if invertCnt is even, this means it is the same with the s1,
	// where s1 always "0"
	if invertCnt%2 == 0 {
		return '0'
	}
	return '1'
}
