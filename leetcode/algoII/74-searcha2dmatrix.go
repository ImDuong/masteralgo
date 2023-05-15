package algoII

// https://leetcode.com/problems/search-a-2d-matrix

func SearchMatrix(matrix [][]int, target int) bool {
	// binary search for row, and then for column

	// step 1. binary search for row (aka all elements in the first column)
	var foundRowIdx int
	lowB := 0
	upB := len(matrix) - 1
	for upB >= lowB {
		if matrix[lowB][0] == target {
			return true
		}
		if matrix[upB][0] == target {
			return true
		}
		foundRowIdx = (upB + lowB) / 2
		if matrix[foundRowIdx][0] < target {
			if foundRowIdx < len(matrix)-1 && matrix[foundRowIdx+1][0] > target {
				break
			}
			lowB = foundRowIdx + 1
		} else if matrix[foundRowIdx][0] > target {
			upB = foundRowIdx - 1
		} else {
			return true
		}
	}
	if upB < 0 || lowB > len(matrix) {
		return false
	}

	// step 2. binary search for column (aka all elements in the found row)
	lowB = 0
	upB = len(matrix[0]) - 1
	for upB >= lowB {
		if matrix[foundRowIdx][lowB] == target {
			return true
		}
		if matrix[foundRowIdx][upB] == target {
			return true
		}
		jmp := (upB + lowB) / 2
		if matrix[foundRowIdx][jmp] < target {
			lowB = jmp + 1
		} else if matrix[foundRowIdx][jmp] > target {
			upB = jmp - 1
		} else {
			return true
		}
	}

	return false
}
