package heap_priority_queue

import (
	"container/heap"
)

func Challenge2530(nums []int, k int) int64 {
	return maxKelements(nums, k)
}

func maxKelements(nums []int, k int) int64 {
	h2 := IntHeap(nums)
	h := &h2
	heap.Init(h)
	var rs int64 = 0
	for k > 0 {
		curMax := (*h)[0]
		rs += int64(curMax)

		h.update(0, (curMax+3-1)/3)
		k--
	}
	return rs
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	poppedItem := (*h)[h.Len()-1]
	*h = (*h)[0 : h.Len()-1]
	return poppedItem
}

func (h *IntHeap) update(idx int, value int) {
	(*h)[idx] = value
	heap.Fix(h, idx)
}
