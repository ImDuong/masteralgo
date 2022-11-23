package algoI

import (
	"masteralgo/internal/core/domain"
	"masteralgo/pkg/helpers"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func MiddleNode(head *domain.ListNode) *domain.ListNode {
	nbNode := helpers.GetNumberOfNodesFromLinkedList(head)
	middleIdx := nbNode / 2
	travNode := head
	for i := 0; i < middleIdx; i++ {
		travNode = travNode.Next
	}
	return travNode
}
