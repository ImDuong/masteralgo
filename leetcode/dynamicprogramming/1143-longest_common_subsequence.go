package dynamicprogramming

// using dp
func LongestCommonSubsequence(text1 string, text2 string) int {
	len1 := len(text1)
	len2 := len(text2)
	f := make([][]int, len1)
	for i := 0; i < len1; i++ {
		f[i] = make([]int, len2)
		for j := 0; j < len2; j++ {
			if text1[i] == text2[j] {
				if i-1 >= 0 && j-1 >= 0 {
					f[i][j] = 1 + f[i-1][j-1]
				} else {
					f[i][j] = 1
				}
			} else {
				var leftF, rightF int
				if i-1 >= 0 {
					leftF = f[i-1][j]
				}
				if j-1 >= 0 {
					rightF = f[i][j-1]
				}
				f[i][j] = max(leftF, rightF)
			}
		}
	}

	return f[len1-1][len2-1]
}

// in dp solution, only the previous and current row is required to calculate the current f[i]
// e.g.: f[i][j] = 1 + f[i-1][j-1] -> need row i - 1
// e.g.: f[i][j] = max(f[i-1][j],  f[i][j-1]) -> need row i - 1
// Hence, we reduce from f[n][m] to f[2][m]
// We will use two rows alternately.
// e.g.: i is even, we use row 0
// e.g.: i is odd, we use row 1
func LongestCommonSubsequenceOptimizedSpace(text1 string, text2 string) int {
	len1 := len(text1)
	len2 := len(text2)
	f := make([][]int, 2)
	f[0] = make([]int, max(len1, len2))
	f[1] = make([]int, max(len1, len2))
	for i := 0; i < len1; i++ {
		// i is even, curRowIdx is 0, else 1
		curRowIdx := i & 1
		for j := 0; j < len2; j++ {
			if text1[i] == text2[j] {
				var prevF int
				if j-1 >= 0 {
					prevF = f[1-curRowIdx][j-1]
				}
				f[curRowIdx][j] = 1 + prevF
			} else {
				var rightF int
				if j-1 >= 0 {
					rightF = f[curRowIdx][j-1]
				}
				f[curRowIdx][j] = max(f[1-curRowIdx][j], rightF)
			}
		}
	}
	return f[1-(len1&1)][len2-1]
}

func LongestCommonSubsequenceMemoi(text1 string, text2 string) int {
	return lcsMemoi(0, 0, text1, text2)
}

// memoization
func lcsMemoi(i, j int, text1, text2 string) int {
	if i >= len(text1) {
		return 0
	}
	if j >= len(text2) {
		return 0
	}
	if text1[i] == text2[j] {
		return 1 + lcsMemoi(i+1, j+1, text1, text2)
	}
	return max(lcsMemoi(i, j+1, text1, text2), lcsMemoi(i+1, j, text1, text2))
}
