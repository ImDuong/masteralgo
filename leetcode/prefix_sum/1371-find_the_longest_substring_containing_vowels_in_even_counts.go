package prefixsum

// similar to challenge 1915
func Challenge1371(s string) int {
	return findTheLongestSubstring(s)
}

func findTheLongestSubstring(s string) int {
	maskMap := make(map[int]int)
	preMask := 0
	rs := 0
	maskMap[preMask] = -1
	for i := range s {
		getBitMask(s[i], &preMask)
		if _, ok := maskMap[preMask]; !ok && preMask != 0 {
			maskMap[preMask] = i
		}
		rs = max(rs, i-maskMap[preMask])
	}
	return rs
}

func getBitMask(c byte, mask *int) {
	if c == 'a' {
		*mask ^= 1
	} else if c == 'e' {
		*mask ^= 1 << 1
	} else if c == 'i' {
		*mask ^= 1 << 2
	} else if c == 'o' {
		*mask ^= 1 << 3
	} else if c == 'u' {
		*mask ^= 1 << 4
	}
}
