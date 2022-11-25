package domain

import (
	"errors"
	"fmt"
)

type (
	TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}

	TreeNodeWrapper struct {
		ParentNode *TreeNode
		Node       *TreeNode
	}
)

func (root *TreeNode) GetListFromTree() []interface{} {
	if root == nil {
		return nil
	}
	var fullList []interface{}
	fullList = append(fullList, root.Val)

	queue := []*TreeNode{}
	queue = append(queue, root.Left)
	queue = append(queue, root.Right)
	for len(queue) > 0 {
		var poppedNode *TreeNode
		poppedNode, queue = queue[0], queue[1:]
		if poppedNode == nil {
			fullList = append(fullList, nil)
		} else {
			fullList = append(fullList, poppedNode.Val)
			queue = append(queue, poppedNode.Left)
			queue = append(queue, poppedNode.Right)
		}

	}
	var i int
	for i = len(fullList) - 1; i >= 0; i-- {
		if fullList[i] != nil {
			break
		}
	}
	fullList = fullList[:i+1]
	return fullList
}

func (root *TreeNode) GetTreeFromList(tree []interface{}) (ITreeNode, error) {
	if len(tree) == 0 {
		return nil, nil
	}
	if root == nil {
		root = &TreeNode{}
	}
	queue := []*TreeNodeWrapper{}
	queue = append(queue, &TreeNodeWrapper{
		Node: root,
	})
	i := 0
	for len(queue) > 0 && i < len(tree) {
		var poppedNode *TreeNodeWrapper
		poppedNode, queue = queue[0], queue[1:]

		if tree[i] == nil {
			if i%2 == 0 {
				poppedNode.ParentNode.Right = nil
			} else {
				poppedNode.ParentNode.Left = nil
			}
		} else {
			if nodeVal, ok := tree[i].(int); ok {
				poppedNode.Node.Val = nodeVal

				poppedNode.Node.Left = &TreeNode{}
				poppedNode.Node.Right = &TreeNode{}
				queue = append(queue, &TreeNodeWrapper{
					ParentNode: poppedNode.Node,
					Node:       poppedNode.Node.Left,
				})
				queue = append(queue, &TreeNodeWrapper{
					ParentNode: poppedNode.Node,
					Node:       poppedNode.Node.Right,
				})
			} else {
				return nil, errors.New(fmt.Sprintf("%v", tree[i]) + " is not a node value")
			}
		}
		i++
	}

	for len(queue) > 0 {
		var poppedNode *TreeNodeWrapper
		poppedNode, queue = queue[0], queue[1:]
		if i%2 == 0 {
			poppedNode.ParentNode.Right = nil
		} else {
			poppedNode.ParentNode.Left = nil
		}
		i++
	}

	return root, nil
}
