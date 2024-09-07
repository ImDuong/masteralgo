package string_challenges

func Challenge28(haystack string, needle string) int {
	return strStr(haystack, needle)
}

func strStr(inp string, pat string) int {
	patLen, inpLen := len(pat), len(inp)
	if patLen > inpLen || patLen == 0 || inpLen == 0 {
		return -1
	}

	lps := getLPS(pat)
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

func getLPS(pat string) []int {
	if len(pat) == 0 {
		return []int{}
	}
	patLen := len(pat)
	lps := make([]int, patLen)

	// always start at 0, because there is no proper prefix
	// for string having length = 1
	lps[0] = 0

	maxMatch := 0
	for i := 1; i < patLen; i++ {
		if pat[i] == pat[maxMatch] {
			maxMatch++
			lps[i] = maxMatch
			continue
		} else if maxMatch > 0 {
			maxMatch = lps[maxMatch-1]
			i--
		}
	}

	return lps
}
