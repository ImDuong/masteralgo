package helpers

import "masteralgo/internal/core/domain"

func GetListFromTreeNode(root *domain.TreeNode) []interface{} {
	if root == nil {
		return nil
	}
	var fullList []interface{}
	fullList = append(fullList, root.Val)

	queue := []*domain.TreeNode{}
	if root.Left != nil {
		queue = append(queue, root.Left)
	}
	if root.Right != nil {
		queue = append(queue, root.Right)
	}
	for len(queue) > 0 {
		var poppedNode *domain.TreeNode
		poppedNode, queue = queue[0], queue[1:]
		if poppedNode == nil {
			fullList = append(fullList, nil)
			continue
		}
		fullList = append(fullList, poppedNode.Val)
		if poppedNode.Left != nil {
			queue = append(queue, poppedNode.Left)
		}
		if poppedNode.Right != nil {
			// show nil for left node when there's right node
			if poppedNode.Left == nil {
				queue = append(queue, nil)
			}
			queue = append(queue, poppedNode.Right)
		}
	}

	return fullList
}
