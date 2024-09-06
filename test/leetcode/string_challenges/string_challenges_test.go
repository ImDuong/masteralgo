package stringchallenges_test

import (
	"masteralgo/leetcode/string_challenges"
	"testing"
)

type stringTest struct {
	arg1     string
	expected int
}

var testcase420 = []stringTest{
	{"a", 5},
	{"aA1", 3},
	{"abB1", 2},
	{"1337C0d3", 0},
	{"1337C0d33333", 1},
	{"aaaaabbbbbbbbbbderereFereGGerGGGHHDFGF", 19},
	{"aaabbbbbbbbbbbbbbbbccc", 7},
	{"bbaaaaaaaaaaaaaaacccccc", 8},
}

func Test420(t *testing.T) {
	for idx, test := range testcase420 {
		output := string_challenges.Challenge420(test.arg1)
		if test.expected != output {
			t.Errorf("TEST ID: %d. Expected %v but got %v in %s", idx, test.expected, output, test.arg1)
		}
	}
}
