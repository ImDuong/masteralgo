package bfs

func Challenge773(board [][]int) int {
	return slidingPuzzle(board)
}

// similar to challenge 752
// -> just convert the board to string
// 1. use bfs to try each combination
// 2. use map to track visited combinations
func slidingPuzzle(board [][]int) int {
	initState := []byte{}
	for i := range board {
		for j := range board[i] {
			initState = append(initState, byte('0'+board[i][j]))
		}
	}
	q := []string{string(initState)}
	visited := map[string]bool{string(initState): true}
	rs := 0
	for len(q) != 0 {
		for range q {
			curState := q[0]
			q = q[1:]

			if curState == "123450" {
				return rs
			}

			// find index of empty square
			zIdx := -1
			for i := range curState {
				if curState[i] == '0' {
					zIdx = i
					break
				}
			}

			// try 4 directions
			var tmpStr string
			for i := 0; i < 4; i++ {
				tmp := []byte(curState)
				if i == 0 && zIdx > 2 {
					// swap up
					tmp[zIdx], tmp[zIdx-3] = tmp[zIdx-3], tmp[zIdx]
				} else if i == 1 && zIdx < 3 {
					// swap down
					tmp[zIdx], tmp[zIdx+3] = tmp[zIdx+3], tmp[zIdx]
				} else if i == 2 && (zIdx == 1 || zIdx == 2 || zIdx == 4 || zIdx == 5) {
					// swap left
					tmp[zIdx], tmp[zIdx-1] = tmp[zIdx-1], tmp[zIdx]
				} else if i == 3 && (zIdx == 0 || zIdx == 1 || zIdx == 3 || zIdx == 4) {
					// swap right
					tmp[zIdx], tmp[zIdx+1] = tmp[zIdx+1], tmp[zIdx]
				}

				tmpStr = string(tmp)
				if _, ok := visited[tmpStr]; !ok {
					q = append(q, tmpStr)
					visited[tmpStr] = true
				}
			}
		}
		rs++
	}
	return -1
}
