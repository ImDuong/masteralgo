package heap_priority_queue

import (
	"container/heap"
	"masteralgo/internal/boilerplates"
)

func Challenge2530(nums []int, k int) int64 {
	return maxKelements(nums, k)
}

func maxKelements(nums []int, k int) int64 {
	h2 := boilerplates.IntMaxHeap(nums)
	h := &h2
	heap.Init(h)
	var rs int64 = 0
	for k > 0 {
		curMax := (*h)[0]
		rs += int64(curMax)

		h.Update(0, (curMax+3-1)/3)
		k--
	}
	return rs
}
