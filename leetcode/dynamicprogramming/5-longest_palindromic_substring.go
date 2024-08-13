package dynamicprogramming

func LongestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	gMaxStr := s[0:1]
	gMaxLen := 1
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			isP := true
			for k := i; k <= (i+j)/2; k++ {
				// transpose to array where i starts from 0 -> get the formula
				if s[k] != s[j-k+i] {
					isP = false
					break
				}
			}
			if isP && j-i+1 > gMaxLen {
				gMaxLen = j - i + 1
				gMaxStr = s[i : j+1]
			}
		}
	}
	return gMaxStr
}
