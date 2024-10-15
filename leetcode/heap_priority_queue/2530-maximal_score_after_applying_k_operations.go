package heap_priority_queue

import (
	"container/heap"
)

func Challenge2530(nums []int, k int) int64 {
	return maxKelements(nums, k)
}

func maxKelements(nums []int, k int) int64 {
	h2 := IntMaxHeap(nums)
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

type IntMaxHeap []int

func (h IntMaxHeap) Len() int           { return len(h) }
func (h IntMaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMaxHeap) Pop() interface{} {
	poppedItem := (*h)[h.Len()-1]
	*h = (*h)[0 : h.Len()-1]
	return poppedItem
}

func (h *IntMaxHeap) update(idx int, value int) {
	(*h)[idx] = value
	heap.Fix(h, idx)
}
