package heap_priority_queue

import (
	"container/heap"
	"sort"
)

func Challenge2054(events [][]int) int {
	return maxTwoEvents(events)
}

// 1. sort events increasing based on start time
// 2. use min heap to track end time
// 3. for each event, we pop out valid end time events & find the maximum value
// -> by using min heap, we always make sure we meet all events that non overlapping with current event
func maxTwoEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	h := &endTimeMinHeap{}
	heap.Init(h)
	rs := 0
	curMaxVal := 0
	for i := range events {
		for h.Len() > 0 && h.Top().time < events[i][0] {
			curMaxVal = max(curMaxVal, heap.Pop(h).(endTime).value)
		}

		rs = max(rs, events[i][2]+curMaxVal)
		heap.Push(h, endTime{
			time:  events[i][1],
			value: events[i][2],
		})
	}
	return rs
}

type endTime struct {
	time, value int
}

type endTimeMinHeap []endTime

func (h endTimeMinHeap) Len() int           { return len(h) }
func (h endTimeMinHeap) Less(i, j int) bool { return h[i].time < h[j].time }
func (h endTimeMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *endTimeMinHeap) Push(x interface{}) {
	*h = append(*h, x.(endTime))
}

func (h *endTimeMinHeap) Pop() interface{} {
	poppedItem := (*h)[h.Len()-1]
	*h = (*h)[0 : h.Len()-1]
	return poppedItem
}

func (h *endTimeMinHeap) Top() endTime {
	return (*h)[0]
}
