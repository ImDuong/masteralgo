package dynamicprogramming

func Challenge1639(words []string, target string) int {
	return numWays(words, target)
}

// For an index, selecting a character in each word is independent
// -> we can use map to avoid looping through words to find the match character
// -> this map stores frequency of characters in each row of word
func numWays(words []string, target string) int {
	nbChr := len(words[0])
	freq := make([][]int, nbChr)
	for i := range freq {
		freq[i] = make([]int, 26)
		for j := range words {
			freq[i][int(words[j][i]-'a')]++
		}
	}

	dp := make([][]int, len(target))
	for i := range dp {
		dp[i] = make([]int, nbChr)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	return getNbWaysRecursively(0, 0, nbChr, target, freq, dp)
}

// When standing at a position in a word, we have 2 choices:
// 1. skip the current match
// 2. match and go to the next one (we use frequency to calculate the nb of ways)
// -> this can use recursion.
// To avoid re-calculating, we can use a 2D array to store the number of ways related to a word index and target index
func getNbWaysRecursively(curTargetIdx, curWordIdx int, wordLen int, target string, freq, dp [][]int) int {
	if curTargetIdx == len(target) {
		return 1
	}

	if curWordIdx == wordLen || wordLen-curWordIdx < len(target)-curTargetIdx {
		return 0
	}

	if dp[curTargetIdx][curWordIdx] != -1 {
		return dp[curTargetIdx][curWordIdx]
	}

	nbWays := 0
	// skip the current word index -> move to the next one
	nbWays += getNbWaysRecursively(curTargetIdx, curWordIdx+1, wordLen, target, freq, dp)

	// match current word index -> use freq to add the nbWays
	nbWays += freq[curWordIdx][int(target[curTargetIdx]-'a')] * getNbWaysRecursively(curTargetIdx+1, curWordIdx+1, wordLen, target, freq, dp)

	dp[curTargetIdx][curWordIdx] = nbWays % 1_000_000_007
	return dp[curTargetIdx][curWordIdx]
}
