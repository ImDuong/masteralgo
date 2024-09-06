package linkedlist

func Challenge3217(nums []int, head *ListNode) *ListNode {
	return modifiedList(nums, head)
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	wipedOut := make(map[int]bool)
	for i := range nums {
		wipedOut[nums[i]] = true
	}

	// really nice trick to deal when traversaling with linked list
	preNode := &ListNode{
		Next: head,
	}
	cur := preNode
	for cur.Next != nil {
		_, ok := wipedOut[cur.Next.Val]
		if !ok {
			cur = cur.Next
			continue
		}
		cur.Next = cur.Next.Next
	}
	return preNode.Next
}
