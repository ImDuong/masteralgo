package dynamicprogramming_test

import (
	"masteralgo/leetcode/dynamicprogramming"
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
