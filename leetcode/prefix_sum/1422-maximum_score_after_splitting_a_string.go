package prefixsum

func Challenge1422(s string) int {
	return maxScore(s)
}

// max(zeroesL + onesR) = max(zeroesL - onesL + onesL + onesR) = max(zeroesL - onesL) + onesTotal
func maxScore(s string) int {
	rs, zeroes, ones := 0, 0, 0
	for i := range s {
		if s[i] == '0' {
			zeroes++
		} else {
			ones++
		}
		if i < len(s)-1 {
			rs = max(rs, zeroes-ones)
		}
	}
	return rs + ones
}
