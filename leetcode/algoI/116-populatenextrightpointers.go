package algoI

import "masteralgo/internal/core/domain"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func Connect(root *domain.Node) *domain.Node {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		root.Left.Next = root.Right
	}
	if root.Right != nil {
		if root.Next == nil {
			root.Right.Next = nil
		} else {
			root.Right.Next = root.Next.Left
		}
	}
	Connect(root.Left)
	Connect(root.Right)

	return root
}
