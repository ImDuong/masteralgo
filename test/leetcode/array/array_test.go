package array_test

import (
	"masteralgo/leetcode/array"
	"testing"
)

type arrayTest struct {
	arg1     [][]int
	expected []int
}

var testcase1380 = []arrayTest{
	{[][]int{
		{3, 7, 8}, {9, 11, 13}, {15, 16, 17},
	},
		[]int{15},
	},
	{[][]int{
		{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12},
	},
		[]int{12},
	},
	{[][]int{
		{7, 8}, {1, 2},
	},
		[]int{7},
	},
}

func Test1380(t *testing.T) {
	for idx, test := range testcase1380 {
		output := array.Challenge1380(test.arg1)
		if len(output) != len(test.expected) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, output)
		}
		for i := range output {
			isFound := false
			for j := range test.expected {
				if output[i] == test.expected[j] {
					isFound = true
					break
				}
			}
			if !isFound {
				t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, output)
			}
		}
	}
}
