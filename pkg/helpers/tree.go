package helpers

import (
	"fmt"
	"masteralgo/internal/core/domain"
)

func GetListFromBinaryTree(root *domain.TreeNode) []interface{} {
	if root == nil {
		return nil
	}
	var fullList []interface{}
	fullList = append(fullList, root.Val)

	queue := []*domain.TreeNode{}
	queue = append(queue, root.Left)
	queue = append(queue, root.Right)
	for len(queue) > 0 {
		var poppedNode *domain.TreeNode
		poppedNode, queue = queue[0], queue[1:]
		if poppedNode == nil {
			fullList = append(fullList, nil)
		} else {
			fullList = append(fullList, poppedNode.Val)
			if !(poppedNode.Left == nil && poppedNode.Right == nil) {
				queue = append(queue, poppedNode.Left)
				queue = append(queue, poppedNode.Right)
			}
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

func GetBinaryTreeFromList(tree []interface{}) *domain.TreeNode {
	root := domain.TreeNode{}
	if len(tree) == 0 {
		return nil
	}
	queue := []*domain.TreeNodeWrapper{}
	queue = append(queue, &domain.TreeNodeWrapper{
		Node: &root,
	})
	i := 0
	for len(queue) > 0 && i < len(tree) {
		var poppedNode *domain.TreeNodeWrapper
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

				poppedNode.Node.Left = &domain.TreeNode{}
				poppedNode.Node.Right = &domain.TreeNode{}
				queue = append(queue, &domain.TreeNodeWrapper{
					ParentNode: poppedNode.Node,
					Node:       poppedNode.Node.Left,
				})
				queue = append(queue, &domain.TreeNodeWrapper{
					ParentNode: poppedNode.Node,
					Node:       poppedNode.Node.Right,
				})
			} else {
				fmt.Println("Node value: ", tree[i])
				panic("input is not a node value")
			}
		}
		i++
	}

	for len(queue) > 0 {
		var poppedNode *domain.TreeNodeWrapper
		poppedNode, queue = queue[0], queue[1:]
		if i%2 == 0 {
			poppedNode.ParentNode.Right = nil
		} else {
			poppedNode.ParentNode.Left = nil
		}
		i++
	}

	return &root
}
