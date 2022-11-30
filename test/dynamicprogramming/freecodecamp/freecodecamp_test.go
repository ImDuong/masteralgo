package dp_freecodecamp_test

import (
	dp_freecodecamp "masteralgo/dynamicprogramming/freecodecamp"
	"testing"
)

type (
	arrayTest struct {
		arg1       []bool
		arg2, arg3 int
		expected   bool
	}
)

var testcaseLandingRunway = []arrayTest{
	{[]bool{true, false, true, true, true, true, false}, 2, 0, true},
	{[]bool{true, false, true, true, false}, 2, 0, true},
	{[]bool{true, false, true, true, false}, 1, 0, true},
	{[]bool{false, false, true, true, false}, 1, 0, false},
}

func TestLandingRunwayRecusive(t *testing.T) {
	for idx, test := range testcaseLandingRunway {
		if output := dp_freecodecamp.CanStopRecursive(test.arg1, test.arg2, test.arg3); output != test.expected {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, output)
		}
	}
}

func TestLandingRunwayIterative(t *testing.T) {
	for idx, test := range testcaseLandingRunway {
		if output := dp_freecodecamp.CanStopIterative(test.arg1, test.arg2, test.arg3); output != test.expected {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, output)
		}
	}
}
