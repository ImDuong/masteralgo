package algoI

func LengthOfLongestSubstring(s string) int {
	maxLen := 0
	curStr := make([]byte, len(s))
	curLen := 0
	for i := range s {
		foundIdx := -1
		for j := 0; j < curLen; j++ {
			if curStr[j] == s[i] {
				foundIdx = j
				break
			}
		}
		if foundIdx != -1 {
			if maxLen < curLen {
				maxLen = curLen
			}
			curStr = curStr[foundIdx+1:]
			curLen -= foundIdx + 1
		}
		curStr[curLen] = s[i]
		curLen++
	}
	if maxLen < curLen {
		return curLen
	}
	return maxLen
}

func LengthOfLongestSubstringV2(s string) int {
	var maxLen, headIdx int = 0, 0
	chars := make([]int, 128)
	for i := range s {
		lastIdx := chars[s[i]]
		if lastIdx > 0 {
			if maxLen < i-headIdx {
				maxLen = i - headIdx
			}
			if headIdx < lastIdx {
				headIdx = lastIdx
			}
		}

		chars[s[i]] = i + 1
	}
	if maxLen < len(s)-headIdx {
		maxLen = len(s) - headIdx
	}
	return maxLen
}
