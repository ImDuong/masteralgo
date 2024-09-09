package string_challenges

import "masteralgo/pkg/helpers"

func Challenge28(haystack string, needle string) int {
	return strStr(haystack, needle)
}

func strStr(inp string, pat string) int {
	patLen, inpLen := len(pat), len(inp)
	if patLen > inpLen || patLen == 0 || inpLen == 0 {
		return -1
	}

	lps := helpers.GetLPS(pat)
	j := 0
	for i := 0; i < inpLen; i++ {
		if inp[i] == pat[j] {
			j++
			if j == patLen {
				return i - j + 1
			}
		} else if j > 0 {
			j = lps[j-1]

			// maintain i
			i--
		}
	}

	return -1
}
