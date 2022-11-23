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
func RemoveNthFromEnd(head *domain.ListNode, n int) *domain.ListNode {
	if head.Next == nil && n == 1 {
		return nil
	}
	nbNode := helpers.GetNumberOfNodesFromLinkedList(head)
	travNode := head
	if n == nbNode {
		return head.Next
	}

	for i := 0; i < nbNode-n-1; i++ {
		travNode = travNode.Next
	}
	if travNode.Next != nil && travNode.Next.Next != nil {
		travNode.Next = travNode.Next.Next
	} else {
		travNode.Next = nil
	}
	return head
}
