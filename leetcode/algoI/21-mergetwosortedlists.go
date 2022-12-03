package algoI

import "masteralgo/internal/core/domain"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func MergeTwoLists(list1 *domain.ListNode, list2 *domain.ListNode) *domain.ListNode {
	var mergedList *domain.ListNode
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		mergedList = list1
		list1 = list1.Next
	} else {
		mergedList = list2
		list2 = list2.Next
	}
	curP := mergedList
	for list1 != nil || list2 != nil {
		if list1 == nil {
			curP.Next = list2
			break
		}
		if list2 == nil {
			curP.Next = list1
			break
		}
		if list1.Val < list2.Val {
			curP.Next = list1
			list1 = list1.Next
		} else {
			curP.Next = list2
			list2 = list2.Next
		}
		curP = curP.Next
	}
	return mergedList
}
