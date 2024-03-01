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

	arrayTest struct {
		arg1     []int
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
