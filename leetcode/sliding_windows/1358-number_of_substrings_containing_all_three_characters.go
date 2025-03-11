package slidingwindows

// initution: use dynamic sliding window to find the number of substrings that contain all characters of a, b, c
// 1. we expand the window until we find a valid substring.
// Instead of continuing to expand the window, we can count number of substrings having the current window as prefix
// In details, we has s[l:r] is a valid substring, then we have n - r - 1 substrings that have s[l:r] as prefix (which is also valid)
// 2. we shrink the window until it is no longer valid
// Note: this challenge is a fundamental problem for the challenge 3306
func Challenge1358(s string) int {
	return numberOfSubstrings(s)
}

func numberOfSubstrings(s string) int {
	rs := 0

	freq := make([]int, 3)
	for l, r := 0, 0; r < len(s); r++ {
		freq[chrToIdx(s[r])]++

		for l <= r && isValid(freq) {
			// 1: itself [l, r]
			// n - r - 1: number of chars left over
			rs += 1 + len(s) - r - 1

			// shrink the window+
			l++
			freq[chrToIdx(s[l-1])]--
		}
	}
	return rs
}

func isValid(freq []int) bool {
	return freq[0] != 0 && freq[1] != 0 && freq[2] != 0
}

func chrToIdx(chr byte) int {
	switch chr {
	case 'a':
		return 0
	case 'b':
		return 1
	case 'c':
		return 2
	default:
		return -1
	}
}
