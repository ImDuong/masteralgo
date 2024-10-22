package heap_priority_queue

import (
	"container/heap"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Challenge2583(root *TreeNode, k int) int64 {
	return kthLargestLevelSum(root, k)
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	// min heap for holding k largest values -> top of the heap will be the kth largest value
	h := &Int64MinHeap{}
	heap.Init(h)

	// queue for level order traversal
	q := []*TreeNode{}
	q = append(q, root)
	depth := 1
	for len(q) != 0 {
		sumLevel := int64(0)
		nbNodeInCurLevel := len(q)
		for i := 0; i < nbNodeInCurLevel; i++ {
			// pop immediately
			// why not pop by batch: so that size of queue is optimized before appending new elements
			poppedNode := q[0]
			q = q[1:]

			sumLevel += int64(poppedNode.Val)
			// push nodes from next level
			if poppedNode.Left != nil {
				q = append(q, poppedNode.Left)
			}
			if poppedNode.Right != nil {
				q = append(q, poppedNode.Right)
			}
		}

		// push sum of current level to the min heap
		if h.Len() == k && h.Top() < sumLevel {
			// if heap is full and the smallest value is smaller than current level sum
			// -> pop the top of heap
			heap.Pop(h)
		}
		if h.Len() < k {
			heap.Push(h, sumLevel)
		}

		depth++
	}

	if depth <= k {
		return -1
	}

	return int64(h.Top())
}

type Int64MinHeap []int64

func (h Int64MinHeap) Len() int           { return len(h) }
func (h Int64MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h Int64MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Int64MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int64))
}

func (h *Int64MinHeap) Pop() interface{} {
	poppedItem := (*h)[h.Len()-1]
	*h = (*h)[0 : h.Len()-1]
	return poppedItem
}

func (h *Int64MinHeap) Top() int64 {
	return (*h)[0]
}

func (h *Int64MinHeap) update(idx int, value int64) {
	(*h)[idx] = value
	heap.Fix(h, idx)
}
