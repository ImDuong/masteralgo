package dp_freecodecamp_test

import (
	dp_freecodecamp "masteralgo/dynamicprogramming/freecodecamp"
	"masteralgo/internal/core/domain"
	"testing"
)

type (
	arrayTest struct {
		arg1       []bool
		arg2, arg3 int
		expected   bool
	}
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
		for j := 0; j < runwayLength; j++ {
			choiceIdx, err := pdf.Rand()
			if err != nil {
				panic(err)
			}
			runway[j] = choices[choiceIdx]
		}
		testcases[i] = arrayTest{
			arg1:     runway,
			arg2:     initSpeed,
			arg3:     startIdx,
			expected: dp_freecodecamp.CanStopRecursiveWithMemo(runway, initSpeed, startIdx),
		}
	}
	return testcases
}

var spikeyRunways = prepareSpikeyRunwayTestCases(1, 1000, 0, 30, 0.4)

// use a global variable to avoid compiler optimization when running benchmark: https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go or https://teivah.medium.com/how-to-write-accurate-benchmarks-in-go-4266d7dd1a95
// more about optimization can be referred from https://livebook.manning.com/book/100-go-mistakes-and-how-to-avoid-them/chapter-12/
var isPossibleRunway bool

func BenchmarkCanStopRecursive(b *testing.B) {
	var output bool
	sR := spikeyRunways
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		output = dp_freecodecamp.CanStopRecursive(sR[0].arg1, sR[0].arg2, sR[0].arg3)
	}
	isPossibleRunway = output
}

func BenchmarkCanStopIterative(b *testing.B) {
	var output bool
	sR := spikeyRunways
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		output = dp_freecodecamp.CanStopIterative(sR[0].arg1, sR[0].arg2, sR[0].arg3)
	}
	isPossibleRunway = output
}

func BenchmarkCanStopRecursiveWithMemo(b *testing.B) {
	var output bool
	sR := spikeyRunways
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		output = dp_freecodecamp.CanStopRecursiveWithMemo(sR[0].arg1, sR[0].arg2, sR[0].arg3)
	}
	isPossibleRunway = output
}
