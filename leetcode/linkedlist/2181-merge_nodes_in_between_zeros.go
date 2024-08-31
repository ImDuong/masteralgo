package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func Challenge2181(head *ListNode) *ListNode {
	return mergeNodes(head)
}

func mergeNodes(head *ListNode) *ListNode {
	rsHeadNode := &ListNode{}
	curRsNode := rsHeadNode
	curInpNode := head.Next

	mergedSum := 0
	firstTime := true
	for curInpNode != nil {
		if curInpNode.Val == 0 {
			if firstTime {
				curRsNode.Val = mergedSum
				firstTime = false
			} else {
				if curRsNode.Next == nil {
					curRsNode.Next = &ListNode{}
				}
				curRsNode.Next.Val = mergedSum
				curRsNode = curRsNode.Next
			}
			curInpNode = curInpNode.Next
			// reset
			mergedSum = 0
			continue
		}

		mergedSum += curInpNode.Val
		curInpNode = curInpNode.Next
	}
	return rsHeadNode
}
