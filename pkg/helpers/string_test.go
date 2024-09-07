package helpers

import (
	"reflect"
	"testing"
)

func TestComputeLPS(t *testing.T) {
	type testData struct {
		pat string
		lps []int
	}

	var tcs = []testData{
		{"A", []int{0}},
		{"AB", []int{0, 0}},
		{"AA", []int{0, 1}},
		{"AAAA", []int{0, 1, 2, 3}},
		{"ABCDE", []int{0, 0, 0, 0, 0}},
		{"AABAACAABAA", []int{0, 1, 0, 1, 2, 0, 1, 2, 3, 4, 5}},
		{"AAACAAAAAC", []int{0, 1, 2, 0, 1, 2, 3, 3, 3, 4}},
		{"AAABAAA", []int{0, 1, 2, 0, 1, 2, 3}},
	}

	for idx, test := range tcs {
		output := getLPS(test.pat)
		if !reflect.DeepEqual(output, test.lps) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.lps, output)
		}
	}
}

func TestKMPSearch(t *testing.T) {
	type testData struct {
		inp      string
		pat      string
		expected []int
	}

	var tcs = []testData{
		{"geeksforgeeks", "geeks", []int{0, 8}},
		{"abcabcabc", "abc", []int{0, 3, 6}},
		{"abcdef", "xyz", []int{}},            // No match
		{"ababa", "aba", []int{0, 2}},         // Partial overlap
		{"abc", "abc", []int{0}},              // Pattern equals the entire text
		{"aaaaa", "aa", []int{0, 1, 2, 3}},    // Pattern with repeating characters
		{"banana", "a", []int{1, 3, 5}},       // Single character pattern
		{"abxabxab", "abc", []int{}},          // Pattern not in text but characters appear
		{"", "abc", []int{}},                  // Empty text
		{"abcdef", "", []int{}},               // Empty pattern
		{"this is the end", "end", []int{12}}, // Pattern appears at the end
	}

	for idx, test := range tcs {
		output := KMPSearch(test.pat, test.inp)
		if len(output) == 0 && len(test.expected) == 0 {
			continue
		}
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, output)
		}
	}
}
