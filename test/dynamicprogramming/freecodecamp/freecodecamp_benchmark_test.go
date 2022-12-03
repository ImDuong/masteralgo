package dp_freecodecamp_test

import (
	dp_freecodecamp "masteralgo/dynamicprogramming/freecodecamp"
	"masteralgo/internal/core/domain"
	"testing"
)

func prepareSpikeyRunwayTestCases(nbRunways, runwayLength, startIdx, initSpeed int, spikeProbPerSpot float32) []arrayTest {
	testcases := make([]arrayTest, nbRunways)
	choices := []bool{true, false}
	pdf, err := domain.NewPDF(choices, []float32{1 - spikeProbPerSpot, spikeProbPerSpot})
	if err != nil {
		panic(err)
	}
	for i := range testcases {
		runway := make([]bool, runwayLength)
		for j := range testcases {
			choiceIdx, err := pdf.Rand()
			if err != nil {
				panic(err)
			}
			runway[j] = choices[choiceIdx]
		}
		testcases[i] = arrayTest{
			arg1: runway,
			arg2: initSpeed,
			arg3: startIdx,
		}
	}
	return testcases
}

var spikeyRunways = prepareSpikeyRunwayTestCases(1, 1000, 0, 30, 0.2)

// use a global variable to avoid compiler optimization when running benchmark: https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
// more about optimization can be referred from https://livebook.manning.com/book/100-go-mistakes-and-how-to-avoid-them/chapter-12/
var isPossibleRunway bool

func BenchmarkCanStopIterative(b *testing.B) {
	var output bool
	for i := 0; i < b.N; i++ {
		for _, test := range spikeyRunways {
			output = dp_freecodecamp.CanStopIterative(test.arg1, test.arg2, test.arg3)
		}
	}
	isPossibleRunway = output
}

func BenchmarkCanStopRecursiveWithMemo(b *testing.B) {
	var output bool
	for i := 0; i < b.N; i++ {
		for _, test := range spikeyRunways {
			output = dp_freecodecamp.CanStopRecursiveWithMemo(test.arg1, test.arg2, test.arg3)
		}
	}
	isPossibleRunway = output
}

func BenchmarkCanStopRecursive(b *testing.B) {
	var output bool
	for i := 0; i < b.N; i++ {
		for _, test := range spikeyRunways {
			output = dp_freecodecamp.CanStopRecursive(test.arg1, test.arg2, test.arg3)
		}
	}
	isPossibleRunway = output
}
