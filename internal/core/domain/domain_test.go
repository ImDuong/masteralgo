package domain_test

import (
	"fmt"
	"masteralgo/internal/core/domain"
	"masteralgo/pkg/helpers"
	"reflect"
	"testing"
)

type (
	treeTest struct {
		arg1     domain.ITreeNode
		expected []interface{}
	}

	treeTestV2 struct {
		arg1     []interface{}
		expected domain.ITreeNode
	}

	probTest struct {
		arg1           []interface{}
		arg2, expected []float32
	}

	probTestV2 struct {
		arg1     []interface{}
		expected domain.ProbDensityFunc
	}
)

var testcaseGetListFromBinaryTree = []treeTest{
	{
		&domain.TreeNode{
			Val: 2,
			Left: &domain.TreeNode{
				Val: 2,
			},
			Right: &domain.TreeNode{
				Val: 1,
				Right: &domain.TreeNode{
					Val: 1,
				},
			},
		},
		[]interface{}{2, 2, 1, nil, nil, nil, 1},
	},
	{
		&domain.TreeNode{
			Val: 1,
			Right: &domain.TreeNode{
				Val: 2,
				Right: &domain.TreeNode{
					Val: 3,
				},
			},
		},
		[]interface{}{1, nil, 2, nil, 3},
	},
	{
		&domain.TreeNode{
			Val: 1,
			Left: &domain.TreeNode{
				Val: 2,
				Left: &domain.TreeNode{
					Val: 3,
				},
			},
		},
		[]interface{}{1, 2, nil, 3},
	},
	{
		&domain.TreeNode{
			Val: 2,
			Left: &domain.TreeNode{
				Val: 2,
				Left: &domain.TreeNode{
					Val: 3,
				},
			},
			Right: &domain.TreeNode{
				Val: 2,
				Right: &domain.TreeNode{
					Val: 3,
				},
			},
		},
		[]interface{}{2, 2, 2, 3, nil, nil, 3},
	},
	{
		&domain.TreeNode{
			Val: 1,
			Left: &domain.TreeNode{
				Val: 2,
				Left: &domain.TreeNode{
					Val: 4,
					Left: &domain.TreeNode{
						Val: 6,
					},
				},
			},
			Right: &domain.TreeNode{
				Val: 3,
				Right: &domain.TreeNode{
					Val: 5,
					Right: &domain.TreeNode{
						Val: 7,
					},
				},
			},
		},
		[]interface{}{1, 2, 3, 4, nil, nil, 5, 6, nil, nil, 7},
	},
	{
		&domain.TreeNode{
			Val: 2,
			Left: &domain.TreeNode{
				Val: 2,
			},
		},
		[]interface{}{2, 2},
	},
	{
		&domain.TreeNode{
			Val: 1,
			Left: &domain.TreeNode{
				Val: 2,
				Left: &domain.TreeNode{
					Val: 4,
				},
			},
			Right: &domain.TreeNode{
				Val: 3,
			},
		},
		[]interface{}{1, 2, 3, 4},
	},
	{
		&domain.TreeNode{
			Val: 1,
			Left: &domain.TreeNode{
				Val: 3,
				Left: &domain.TreeNode{
					Val: 5,
				},
			},
			Right: &domain.TreeNode{
				Val: 2,
			},
		},
		[]interface{}{1, 3, 2, 5},
	},
	{
		&domain.TreeNode{
			Val: 2,
			Left: &domain.TreeNode{
				Val: 1,
				Right: &domain.TreeNode{
					Val: 4,
				},
			},
			Right: &domain.TreeNode{
				Val: 3,
				Right: &domain.TreeNode{
					Val: 7,
				},
			},
		},
		[]interface{}{2, 1, 3, nil, 4, nil, 7},
	},
	{
		&domain.TreeNode{
			Val: 3,
			Left: &domain.TreeNode{
				Val: 4,
				Left: &domain.TreeNode{
					Val: 5,
				},
				Right: &domain.TreeNode{
					Val: 4,
				},
			},
			Right: &domain.TreeNode{
				Val: 5,
				Right: &domain.TreeNode{
					Val: 7,
				},
			},
		},
		[]interface{}{3, 4, 5, 5, 4, nil, 7},
	},
	{
		&domain.TreeNode{
			Val: 1,
			Left: &domain.TreeNode{
				Val: 2,
				Left: &domain.TreeNode{
					Val: 3,
				},
			},
		},
		[]interface{}{1, 2, nil, 3},
	},
	{
		&domain.TreeNode{
			Val: 1,
			Right: &domain.TreeNode{
				Val: 2,
				Right: &domain.TreeNode{
					Val: 3,
				},
			},
		},
		[]interface{}{1, nil, 2, nil, 3},
	},
}

func TestGetListFromBinaryTree(t *testing.T) {
	for idx, test := range testcaseGetListFromBinaryTree {
		output := test.arg1.GetListFromTree()
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, output)
		}
	}
}

var testcaseGetBinaryTreeFromList = []treeTestV2{
	{
		[]interface{}{1, nil, 2, nil, 3},
		&domain.TreeNode{
			Val: 1,
			Right: &domain.TreeNode{
				Val: 2,
				Right: &domain.TreeNode{
					Val: 3,
				},
			},
		},
	},
	{
		[]interface{}{1, 2, nil, 3},
		&domain.TreeNode{
			Val: 1,
			Left: &domain.TreeNode{
				Val: 2,
				Left: &domain.TreeNode{
					Val: 3,
				},
			},
		},
	},
	{
		[]interface{}{1, 2, 3, 4, nil, nil, 5, 6, nil, nil, 7},
		&domain.TreeNode{
			Val: 1,
			Left: &domain.TreeNode{
				Val: 2,
				Left: &domain.TreeNode{
					Val: 4,
					Left: &domain.TreeNode{
						Val: 6,
					},
				},
			},
			Right: &domain.TreeNode{
				Val: 3,
				Right: &domain.TreeNode{
					Val: 5,
					Right: &domain.TreeNode{
						Val: 7,
					},
				},
			},
		},
	},
	{
		[]interface{}{2, 2, 2, 3, nil, nil, 3},
		&domain.TreeNode{
			Val: 2,
			Left: &domain.TreeNode{
				Val: 2,
				Left: &domain.TreeNode{
					Val: 3,
				},
			},
			Right: &domain.TreeNode{
				Val: 2,
				Right: &domain.TreeNode{
					Val: 3,
				},
			},
		},
	},
	{
		[]interface{}{2, 2},
		&domain.TreeNode{
			Val: 2,
			Left: &domain.TreeNode{
				Val: 2,
			},
		},
	},
	{
		[]interface{}{1, nil, 2, nil, 3},
		&domain.TreeNode{
			Val: 1,
			Right: &domain.TreeNode{
				Val: 2,
				Right: &domain.TreeNode{
					Val: 3,
				},
			},
		},
	},
	{
		[]interface{}{1, 3, 2, 5},
		&domain.TreeNode{
			Val: 1,
			Left: &domain.TreeNode{
				Val: 3,
				Left: &domain.TreeNode{
					Val: 5,
				},
			},
			Right: &domain.TreeNode{
				Val: 2,
			},
		},
	},
	{
		[]interface{}{2, 1, 3, nil, 4, nil, 7},
		&domain.TreeNode{
			Val: 2,
			Left: &domain.TreeNode{
				Val: 1,
				Right: &domain.TreeNode{
					Val: 4,
				},
			},
			Right: &domain.TreeNode{
				Val: 3,
				Right: &domain.TreeNode{
					Val: 7,
				},
			},
		},
	},
	{
		[]interface{}{3, 4, 5, 5, 4, nil, 7},
		&domain.TreeNode{
			Val: 3,
			Left: &domain.TreeNode{
				Val: 4,
				Left: &domain.TreeNode{
					Val: 5,
				},
				Right: &domain.TreeNode{
					Val: 4,
				},
			},
			Right: &domain.TreeNode{
				Val: 5,
				Right: &domain.TreeNode{
					Val: 7,
				},
			},
		},
	},
	{
		[]interface{}{1, 2, nil, 3},
		&domain.TreeNode{
			Val: 1,
			Left: &domain.TreeNode{
				Val: 2,
				Left: &domain.TreeNode{
					Val: 3,
				},
			},
		},
	},
}

func TestGetBinaryTreeFromList(t *testing.T) {
	for idx, test := range testcaseGetBinaryTreeFromList {
		var input *domain.TreeNode
		output, err := input.GetTreeFromList(test.arg1)
		if err != nil {
			t.Errorf("TEST ID: %d. Error when parsing input: %v", idx, err)
			continue
		}
		outputParsed := output.GetListFromTree()
		expectedParsed := test.expected.GetListFromTree()
		if !reflect.DeepEqual(outputParsed, expectedParsed) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, expectedParsed, outputParsed)
		}
	}
}

var testcaseNewPDFFromValues = []probTestV2{
	{[]interface{}{1, 2, 2, 2}, domain.ProbDensityFunc{
		Vars: []domain.DiscreteRandVar{
			{
				Prob:   0.25,
				Choice: 1,
			},
			{
				Prob:   0.75,
				Choice: 2,
			},
		},
	}},
}

func TestNewPDFFromValues(t *testing.T) {
	for idx, test := range testcaseNewPDFFromValues {
		output, err := domain.NewPDFFromValues(test.arg1)
		if err != nil {
			t.Errorf("TEST ID: %d. Error when parsing input: %v", idx, err)
			continue
		}
		if !helpers.IsEqualWithoutOrder(output.Vars, test.expected.Vars) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected.Vars, output.Vars)
		}
	}
}

// note that arg2 and expected must be equal
var testcaseProbRand = []probTest{
	{
		arg1:     []interface{}{0, 1, 2},
		arg2:     []float32{0.2, 0.5, 0.3},
		expected: []float32{0.2, 0.5, 0.3},
	},
	{
		arg1:     []interface{}{true, false},
		arg2:     []float32{0.1, 0.9},
		expected: []float32{0.1, 0.9},
	},
}

func TestProbRand(t *testing.T) {
	// the bigger nbSampling is, the more certain the result is
	nbSampling := 10000000
	for idx, test := range testcaseProbRand {
		pdf, err := domain.NewPDF(test.arg1, test.arg2)
		if err != nil {
			t.Errorf("TEST ID: %d. Error when parsing input: %v", idx, err)
			continue
		}

		samples := make([]float32, len(test.arg1))
		for i := 0; i < nbSampling; i++ {
			var rIdx int
			rIdx, err = pdf.Rand()
			if err != nil {
				t.Errorf("TEST ID: %d. Error when getting random: %v", idx, err)
				break
			}
			samples[rIdx]++
		}
		if err != nil {
			continue
		}
		// normalize sample values
		for i := range samples {
			samples[i] /= float32(nbSampling)
		}

		for i := range samples {
			if !helpers.IsRoughlyEqual(float64(samples[i]), float64(test.expected[i]), 0.001) {
				t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, samples)
				break
			}
		}
		fmt.Println(samples)
	}
}
