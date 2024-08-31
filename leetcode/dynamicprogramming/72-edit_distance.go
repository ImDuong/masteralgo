package dynamicprogramming

// this is a problem of minimum edit distance (MED)
// nice document: https://web.stanford.edu/class/cs124/lec/med.pdf
func MinDistance(word1 string, word2 string) int {
	len1 := len(word1)
	len2 := len(word2)
	if len1 == 0 {
		return len2
	} else if len2 == 0 {
		return len1
	}

	dist := make([][]int, len1+1)
	dist[0] = make([]int, len2+1)
	for i := 1; i <= len1; i++ {
		dist[i] = make([]int, len2+1)
		dist[i][0] = i
		for j := 1; j <= len2; j++ {
			dist[0][j] = j
			deleteDist := 1 + dist[i-1][j]
			insertDist := 1 + dist[i][j-1]
			replaceDist := dist[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				replaceDist++
			}

			dist[i][j] = min(deleteDist, insertDist, replaceDist)
		}
	}
	return dist[len1][len2]
}

// only 2 rows are using each loop, hence, use only 2 rows
// Achievement: 0ms & 2.78MB
func MinDistanceSpaceOptimized(word1 string, word2 string) int {
	len1 := len(word1)
	len2 := len(word2)
	if len1 == 0 {
		return len2
	} else if len2 == 0 {
		return len1
	}

	prevDist := make([]int, len2+1)
	curDist := make([]int, len2+1)

	for j := 0; j <= len2; j++ {
		prevDist[j] = j
	}

	for i := 1; i <= len1; i++ {
		curDist[0] = i
		for j := 1; j <= len2; j++ {
			if word1[i-1] == word2[j-1] {
				curDist[j] = prevDist[j-1]
				continue
			}

			curDist[j] = 1 + min(prevDist[j], curDist[j-1], prevDist[j-1])
		}
		// instead of re-initiate new row, we can simply swap
		prevDist, curDist = curDist, prevDist
	}
	return prevDist[len2]
}
