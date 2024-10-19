package math_challenges

func Challenge779(n int, k int) int {
	return kthGrammar(n, k)
}

// Similar to challenge 1545
// This time, the string can always be split into 2 halves, the second half is the NOT version of the first one.
// Whereas, in challenge 1545, there is extra 1 in the middle of these 2 halves, and the second half is inverse(not())
func kthGrammar(n int, k int) int {
	// length of S is 2^(n -1)
	len := 1 << (n - 1)
	invertCnt := 0
	for k > 1 {
		// index starts from 1 instead of 0
		middle := len / 2
		if k > middle {
			k = middle - (len - k)
			invertCnt++
		}

		// shorten length
		len /= 2
	}
	// if invertCnt is even, this means it is the same with the s1,
	// where s1 always 0
	if invertCnt%2 == 0 {
		return 0
	}
	return 1
}
