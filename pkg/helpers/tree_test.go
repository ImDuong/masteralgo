package helpers_test

import (
	"masteralgo/internal/core/domain"
	"masteralgo/pkg/helpers"
	"reflect"
	"testing"
)

type (
	treeTest struct {
		arg1     *domain.TreeNode
		expected []interface{}
	}

	treeTestV2 struct {
		arg1     []interface{}
		expected *domain.TreeNode
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
		output := helpers.GetListFromBinaryTree(test.arg1)
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
		output := helpers.GetBinaryTreeFromList(test.arg1)
		outputParsed := helpers.GetListFromBinaryTree(output)
		expectedParsed := helpers.GetListFromBinaryTree(test.expected)
		if !reflect.DeepEqual(outputParsed, expectedParsed) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, expectedParsed, outputParsed)
		}
	}
}
