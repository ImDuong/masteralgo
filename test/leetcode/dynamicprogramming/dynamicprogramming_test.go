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
