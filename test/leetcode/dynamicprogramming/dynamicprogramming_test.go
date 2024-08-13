package dynamicprogramming_test

import (
	"masteralgo/leetcode/dynamicprogramming"
	"slices"
	"strings"
	"testing"
)

type (
	intTest struct {
		arg1     int
		expected int
	}

	intTestV1 struct {
		arg1, arg2 int
		expected   int
	}

	arrayTest struct {
		arg1     []int
		expected int
	}

	array2dTest struct {
		arg1     [][]int
		expected int
	}

	array2dByteTest struct {
		arg1     [][]byte
		expected int
	}

	stringTest struct {
		arg1     string
		expected []string
	}
)

var testcase1137 = []intTest{
	{3, 2},
	{4, 4},
	{25, 1389537},
}

func Test1137(t *testing.T) {
	for idx, test := range testcase1137 {
		if output := dynamicprogramming.Tribonacci(test.arg1); output != test.expected {
			t.Errorf("TEST ID: %d. Expected %d but got %d", idx, test.expected, output)
		}
	}
}

var testcase746 = []arrayTest{
	{[]int{10, 15, 20}, 15},
	{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
}

func Test746(t *testing.T) {
	for idx, test := range testcase746 {
		if output := dynamicprogramming.MinCostClimbingStairs(test.arg1); output != test.expected {
			t.Errorf("TEST ID: %d. Expected %d but got %d", idx, test.expected, output)
		}
	}
}

var testcase740 = []arrayTest{
	{[]int{3, 4, 2}, 6},
	{[]int{2, 2, 3, 3, 3, 4}, 9},
}

func Test740(t *testing.T) {
	for idx, test := range testcase740 {
		if output := dynamicprogramming.DeleteAndEarn(test.arg1); output != test.expected {
			t.Errorf("TEST ID: %d. Expected %d but got %d", idx, test.expected, output)
		}
	}
}

var testcase62 = []intTestV1{
	{3, 7, 28},
	{3, 2, 3},
}

func Test62(t *testing.T) {
	for idx, test := range testcase62 {
		if output := dynamicprogramming.UniquePathsV2(test.arg1, test.arg2); output != test.expected {
			t.Errorf("TEST ID: %d. Expected %d but got %d", idx, test.expected, output)
		}
	}
}

var testcase63 = []array2dTest{
	{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, 2},
	{[][]int{{0, 1}, {0, 0}}, 1},
	{[][]int{{0, 0}, {0, 1}}, 0},
	{[][]int{{0}}, 1},
}

func Test63(t *testing.T) {
	for idx, test := range testcase63 {
		if output := dynamicprogramming.UniquePathsWithObstacles(test.arg1); output != test.expected {
			t.Errorf("TEST ID: %d. Expected %d but got %d", idx, test.expected, output)
		}
	}
}

var testcase64 = []array2dTest{
	{[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 7},
	{[][]int{{1, 2, 3}, {4, 5, 6}}, 12},
}

func Test64(t *testing.T) {
	for idx, test := range testcase64 {
		if output := dynamicprogramming.MinPathSum(test.arg1); output != test.expected {
			t.Errorf("TEST ID: %d. Expected %d but got %d", idx, test.expected, output)
		}
	}
}

var testcase120 = []array2dTest{
	{[][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}, 11},
	{[][]int{{-10}}, -10},
	{[][]int{{-1}, {2, 3}, {1, -1, -3}}, -1},
}

func Test120(t *testing.T) {
	for idx, test := range testcase120 {
		if output := dynamicprogramming.MinimumTotal(test.arg1); output != test.expected {
			t.Errorf("TEST ID: %d. Expected %d but got %d", idx, test.expected, output)
		}
	}
}

var testcase931 = []array2dTest{
	{[][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}, 13},
	{[][]int{{-19, 57}, {-40, -5}}, -59},
	{[][]int{{-48}}, -48},
}

func Test931(t *testing.T) {
	for idx, test := range testcase931 {
		if output := dynamicprogramming.MinFallingPathSum(test.arg1); output != test.expected {
			t.Errorf("TEST ID: %d. Expected %d but got %d", idx, test.expected, output)
		}
	}
}

var testcase221 = []array2dByteTest{
	{[][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}, 4},
	{[][]byte{{'0', '1'}, {'1', '0'}}, 1},
	{[][]byte{{'0'}}, 0},
	{[][]byte{{'0', '1', '0'}}, 1},
	{[][]byte{{'0'}, {'1'}}, 1},
	{[][]byte{{'1', '0', '1', '1', '0', '1'}, {'1', '1', '1', '1', '1', '1'}, {'0', '1', '1', '0', '1', '1'}, {'1', '1', '1', '0', '1', '0'}, {'0', '1', '1', '1', '1', '1'}, {'1', '1', '0', '1', '1', '1'}}, 4},
}

func Test221(t *testing.T) {
	for idx, test := range testcase221 {
		if output := dynamicprogramming.MaximalSquare(test.arg1); output != test.expected {
			t.Errorf("TEST ID: %d. Expected %d but got %d", idx, test.expected, output)
		}
	}
}

var testcase5 = []stringTest{
	{"babad", []string{"bab", "aba"}},
	{"cbbd", []string{"bb"}},
	{"bb", []string{"bb"}},
}

func Test5(t *testing.T) {
	for idx, test := range testcase5 {
		if output := dynamicprogramming.LongestPalindrome(test.arg1); !slices.Contains(test.expected, output) {
			t.Errorf("TEST ID: %d. Expected %s but got %s", idx, strings.Join(test.expected, ","), output)
		}
	}
}
