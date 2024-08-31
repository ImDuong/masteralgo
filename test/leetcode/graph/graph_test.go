package graph_test

import (
	"masteralgo/leetcode/graph"
	"strconv"
	"testing"
)

type test2699 struct {
	nbNodes     int
	edges       [][]int
	source      int
	destination int
	target      int
	isSolvable  bool
}

var testcase2699 = []test2699{
	{5, [][]int{{1, 4, 1}, {2, 4, -1}, {3, 0, 2}, {0, 4, -1}, {1, 3, 10}, {1, 0, 10}}, 0, 2, 15, true},
	{5, [][]int{{4, 1, -1}, {2, 0, -1}, {0, 3, -1}, {4, 3, -1}}, 0, 1, 5, true},
	{3, [][]int{{0, 1, -1}, {0, 2, 5}}, 0, 2, 6, false},
	{4, [][]int{{1, 0, 4}, {1, 2, 3}, {2, 3, 5}, {0, 3, -1}}, 0, 2, 6, true},
}

func Test2699(t *testing.T) {
	for idx, test := range testcase2699 {
		output := graph.ModifiedGraphEdges(test.nbNodes, test.edges, test.source, test.destination, test.target)
		if !test.isSolvable && len(output) == 0 {
			continue
		}

		actualWeight := graph.GetSTPForPositiveWeights(test.nbNodes, output, test.source, test.destination)
		if actualWeight != test.target {
			t.Errorf("TEST ID: %d. The actual STP weight is: %d. Wrong output %s", idx, actualWeight, func(array2D [][]int) string {
				str := "["
				for i := range array2D {
					str += "["
					for j := range array2D[i] {
						str += strconv.Itoa(array2D[i][j]) + " "
					}
					str += "],"
				}
				str += "]"
				return str
			}(output))
		}
	}
}

type test1514 struct {
	nbNodes   int
	edges     [][]int
	succProb  []float64
	startNode int
	endNode   int
	expected  float64
}

var testcase1514 = []test1514{
	{3, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{0.5, 0.5, 0.2}, 0, 2, 0.25},
	{3, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{0.5, 0.5, 0.3}, 0, 2, 0.3},
	{3, [][]int{{0, 1}}, []float64{0.5}, 0, 2, 0},
}

func Test1514(t *testing.T) {
	for idx, test := range testcase1514 {
		output := graph.MaxProbability(test.nbNodes, test.edges, test.succProb, test.startNode, test.endNode)
		if output != test.expected {
			t.Errorf("TEST ID: %d. Expected %f but got %f", idx, test.expected, output)
		}
	}
}
