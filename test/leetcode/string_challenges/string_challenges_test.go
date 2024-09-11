package stringchallenges_test

import (
	"masteralgo/leetcode/string_challenges"
	"masteralgo/pkg/helpers"
	"testing"
)

type (
	stringTest struct {
		arg1     string
		expected int
	}

	stringTestV2 struct {
		arg1     []string
		expected [][]int
	}
)

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

var testcase336 = []stringTestV2{
	{[]string{"abcd", "dcba", "lls", "s", "sssll"}, [][]int{{0, 1}, {1, 0}, {3, 2}, {2, 4}}},
	{[]string{"bat", "tab", "cat"}, [][]int{{0, 1}, {1, 0}}},
	{[]string{"a", ""}, [][]int{{0, 1}, {1, 0}}},
	{[]string{"sl", "s"}, [][]int{{0, 1}}},
	{[]string{"a", "abc", "aba", ""}, [][]int{{0, 3}, {3, 0}, {2, 3}, {3, 2}}},
	{[]string{"a", "b", "c", "ab", "ac", "aa"}, [][]int{{3, 0}, {1, 3}, {4, 0}, {2, 4}, {5, 0}, {0, 5}}},
	{[]string{"abaaaaa", "aba"}, [][]int{{0, 1}}},
	{[]string{"a", "aa", "aaa"}, [][]int{{1, 0}, {0, 1}, {2, 0}, {1, 2}, {2, 1}, {0, 2}}},
	{[]string{"aba", "ba", "a", "caba"}, [][]int{{0, 1}, {2, 1}, {0, 3}}},
}

func Test336(t *testing.T) {
	for idx, test := range testcase336 {
		output := string_challenges.Challenge336(test.arg1)
		if !helpers.IsEqualWithoutOrder(test.expected, output) {
			t.Errorf("TEST ID: %d. Expected %v but got %v in %s", idx, test.expected, output, test.arg1)
		}
	}
}
