package domain_test

import (
	"masteralgo/internal/core/domain"
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
