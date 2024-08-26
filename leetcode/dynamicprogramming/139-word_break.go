package dynamicprogramming

import "slices"

func WordBreak(s string, wordDict []string) bool {
	lengthToWords := make(map[int][]string, len(wordDict))
	minLen := 20
	maxLen := 1
	for i := range wordDict {
		wordLen := len(wordDict[i])
		if wordLen < minLen {
			minLen = wordLen
		}
		if wordLen > maxLen {
			maxLen = wordLen
		}
		lengthToWords[wordLen] = append(lengthToWords[wordLen], wordDict[i])
	}
	// true as non-breakable
	// false as breakable
	memoiNonbreakable := make([]bool, len(s)+1)
	return isWordBreakable(0, &s, minLen, maxLen, lengthToWords, memoiNonbreakable)
}

func isWordBreakable(startIdx int, s *string, minLen, maxLen int, lengthToWords map[int][]string, memoiNonbreakable []bool) bool {
	sLen := len(*s)
	if sLen-startIdx < minLen {
		return false
	}
	for subLen := minLen; subLen <= maxLen; subLen++ {
		wordDictByLen, ok := lengthToWords[subLen]
		if !ok {
			continue
		}
		if startIdx+subLen > sLen {
			return false
		}
		startSubStr := (*s)[startIdx : startIdx+subLen]
		if !slices.Contains(wordDictByLen, startSubStr) {
			continue
		}
		nextIdx := startIdx + subLen
		if memoiNonbreakable[nextIdx] {
			continue
		}
		if nextIdx == sLen {
			return true
		}
		if isWordBreakable(nextIdx, s, minLen, maxLen, lengthToWords, memoiNonbreakable) {
			return true
		}
		memoiNonbreakable[nextIdx] = true
	}
	return false
}
