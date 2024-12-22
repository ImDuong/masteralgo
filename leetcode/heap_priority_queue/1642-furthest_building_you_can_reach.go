package heap_priority_queue

import (
	"container/heap"
	"masteralgo/internal/boilerplates"
)

func Challenge1642(heights []int, bricks int, ladders int) int {
	return furthestBuilding(heights, bricks, ladders)
}

// it's always optimal to use ladders for longest distance
// => maintain a min heap for storing differences that required ladders
// - whenever heap size is bigger than nb of ladders -> we use bricks for the smallest difference (aka top of min heap)
func furthestBuilding(heights []int, bricks int, ladders int) int {
	h := &boilerplates.IntMinHeap{}
	heap.Init(h)
	for i := 1; i < len(heights); i++ {
		if heights[i] <= heights[i-1] {
			continue
		}
		dist := heights[i] - heights[i-1]
		heap.Push(h, dist)
		if h.Len() > ladders {
			minDist := heap.Pop(h).(int)
			bricks -= minDist
		}
		if bricks < 0 {
			return i - 1
		}
	}
	return len(heights) - 1
}
