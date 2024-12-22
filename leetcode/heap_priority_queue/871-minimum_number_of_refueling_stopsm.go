package heap_priority_queue

import (
	"container/heap"
	"masteralgo/internal/boilerplates"
)

// similar idea with challenge 1642
func Challenge871(target int, startFuel int, stations [][]int) int {
	return minRefuelStops(target, startFuel, stations)
}

// to get the min stops, we encounter the stations with max gas as possible (greedy)
// -> use a max heap to store fuels
func minRefuelStops(target int, startFuel int, stations [][]int) int {
	rs := 0
	stationIdx := 0

	h := &boilerplates.IntMaxHeap{}
	heap.Init(h)

	// cur: position in the road
	// init with startFuel miles
	cur := startFuel
	for cur < target {
		// add stations that possible to visited in previous iteration
		for stationIdx < len(stations) && stations[stationIdx][0] <= cur {
			heap.Push(h, stations[stationIdx][1])
			stationIdx++
		}

		// if car does not have fuel left -> cannot reach the dst
		if h.Len() == 0 {
			return -1
		}

		cur += heap.Pop(h).(int)
		rs++
	}
	return rs
}
