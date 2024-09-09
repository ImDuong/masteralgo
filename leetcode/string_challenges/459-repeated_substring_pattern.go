package string_challenges

import (
	"masteralgo/pkg/helpers"
	"strings"
)

func Challenge459(s string) bool {
	return repeatedSubstringPattern(s)
}

func repeatedSubstringPattern(s string) bool {
	strLen := len(s)
	if strLen == 1 {
		return false
	}
	if strLen == 2 {
		return s[0] == s[1]
	}

	lps := helpers.GetLPS(s)
	if lps[strLen-1] <= 1 {
		return false
	}
	repLen := strLen - lps[strLen-1]
	return strLen%repLen == 0
}

func repeatedSubstringPatternV2(s string) bool {
	ss := (s + s)[1 : len(s)*2-1]
	return strings.Contains(ss, s)
}
