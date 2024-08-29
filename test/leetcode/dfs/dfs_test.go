package dfs_test

import (
	"masteralgo/leetcode/dfs"
	"testing"
)

type (
	array2dTest struct {
		arg1     [][]int
		expected int
	}
)

var testcase947 = []array2dTest{
	{[][]int{{0, 0}, {2, 2}, {10000, 2}}, 1},
	{[][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}, 5},
	{[][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}}, 3},
	{[][]int{{0, 0}}, 0},
}

func Test947(t *testing.T) {
	for idx, test := range testcase947 {
		if output := dfs.RemoveStones(test.arg1); output != test.expected {
			t.Errorf("TEST ID: %d. Expected %d but got %d", idx, test.expected, output)
		}
	}
}
