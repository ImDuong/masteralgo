package dynamicprogramming

func LongestPalindromeSubseq(s string) int {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	revS := string(runes)
	return LongestCommonSubsequenceOptimizedSpace(s, revS)
}
