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
)

var testcaseGetListFromTreeNode = []treeTest{
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
}

func TestGetListFromTreeNode(t *testing.T) {
	for idx, test := range testcaseGetListFromTreeNode {
		output := helpers.GetListFromTreeNode(test.arg1)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, output)
		}
	}
}
