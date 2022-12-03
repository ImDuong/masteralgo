package algoI

import (
	"masteralgo/internal/core/domain"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func ReverseList(head *domain.ListNode) *domain.ListNode {
	var tmpP *domain.ListNode
	for head != nil {
		nextNode := head.Next
		head.Next = tmpP
		tmpP = head
		head = nextNode
	}
	return tmpP
}
