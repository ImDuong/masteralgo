package linkedlist

func Challenge2487(head *ListNode) *ListNode {
	return removeNodes(head)
}

// if we go from highest position (from left to right), we will not know what is the next second highest value
// This means when traversing, we only know the local second highest value, not the global second highest value
// e.g., 5 -> 2 -> 3 -> 4 -> 1
//   - When we at 5, we only know the 2 is the local second highest value, while 4 should be the next second highest value
//   - However, if we go from the lowest to the highest: 1 -> 4 -> 3 -> 2 -> 5, we can easily obtain 1 -> 4 -> 5
//
// Hence, going from right to left of the linked list is the optimal way.
// But the linked list is one direction -> reverse it
func removeNodes(head *ListNode) *ListNode {
	head = reverseList(head)
	cur := &ListNode{Next: head}
	for cur.Next != nil {
		for cur.Val > cur.Next.Val {
			cur.Next = cur.Next.Next
			if cur.Next == nil {
				break
			}
		}
		cur = cur.Next
		if cur == nil {
			break
		}
	}
	return reverseList(head)
}

func reverseList(head *ListNode) *ListNode {
	var tmpP *ListNode
	for head != nil {
		nextNode := head.Next
		head.Next = tmpP
		tmpP = head
		head = nextNode
	}
	return tmpP
}
