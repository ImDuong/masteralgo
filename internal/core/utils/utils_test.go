package utils_test

import (
	"masteralgo/internal/core/utils"
	"testing"
)

type (
	arrayTest struct {
		arg1     []int
		arg2     int
		arg3     int
		arg4     int
		expected int
	}
)

var testcaseBinarySearch = []arrayTest{
	{[]int{-1, 0, 3, 5, 9, 12}, 0, 5, -1, 0},
	{[]int{-1, 0, 3, 5, 9, 12}, 0, 5, 0, 1},
	{[]int{-1, 0, 3, 5, 9, 12}, 0, 5, 3, 2},
	{[]int{-1, 0, 3, 5, 9, 12}, 0, 5, 5, 3},
	{[]int{-1, 0, 3, 5, 9, 12}, 0, 5, 9, 4},
	{[]int{-1, 0, 3, 5, 9, 12}, 0, 5, 12, 5},
	{[]int{-1, 0, 3, 5, 9, 12}, 0, 5, 2, -1},
	{[]int{-1, 0, 3, 5, 9, 12}, 0, 5, 69, -1},
	{[]int{0}, 0, 0, 69, -1},
	{[]int{0}, 0, 0, 0, 0},
	{[]int{0, 1}, 0, 1, 2, -1},
	{[]int{0, 1}, 0, 1, 1, 1},
}

func TestBinarySearch(t *testing.T) {
	for idx, test := range testcaseBinarySearch {
		if output := utils.BinarySearch(test.arg1, test.arg2, test.arg3, test.arg4); output != test.expected {
			t.Errorf("TEST ID: %d. Expected %d but got %d", idx, test.expected, output)
		}
	}
}
