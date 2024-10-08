package helpers

// KMPSearch uses KMP to find indicies of pattern found in input string
func KMPSearch(pat string, inp string) []int {
	patLen, inpLen := len(pat), len(inp)
	if patLen > inpLen || patLen == 0 || inpLen == 0 {
		return []int{}
	}

	lps := GetLPS(pat)
	var rs []int
	j := 0
	for i := 0; i < inpLen; i++ {
		if inp[i] == pat[j] {
			j++
			if j == patLen {
				rs = append(rs, i-j+1)
				j = lps[j-1]
			}
		} else if j > 0 {
			j = lps[j-1]

			// maintain i
			i--
		}
	}

	return rs
}

// lps[i]: the longest proper prefix that also be a suffix, obtained from pat[0:i + 1]
// example: "aabcdaabc" -> "aabc" is both proper prefix and a suffix -> lps for "aabcdaabc" is 4
func GetLPS(pat string) []int {
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
