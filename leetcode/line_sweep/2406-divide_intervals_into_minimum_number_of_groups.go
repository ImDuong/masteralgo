package linesweep

import "masteralgo/internal/core/domain"

func Challenge2406(intervals [][]int) int {
	return minGroups(intervals)
}

func minGroups(intervals [][]int) int {
	m := domain.NewSortedKeysMap[int, int]()
	for i := range intervals {
		cnt, ok := m.Get(intervals[i][0])
		if !ok {
			cnt = 0
		}
		m.Set(intervals[i][0], cnt+1)
		cnt, ok = m.Get(intervals[i][1] + 1)
		if !ok {
			cnt = 0
		}
		m.Set(intervals[i][1]+1, cnt-1)
	}
	rs := 1
	currentNbRooms := 0
	for it := m.Iterator(); it.Valid(); it.Next() {
		currentNbRooms += it.Value()
		rs = max(rs, currentNbRooms)
	}
	return rs
}

// still minesweep but without sorted keys map
// suitable when the whole intervals is small. In this challenge, the constraints is 1 <= lefti <= righti <= 10^6
func minGroupsWithCountingSort(intervals [][]int) int {
	rangeStart, rangeEnd := 1000007, 1
	for i := range intervals {
		rangeStart = min(rangeStart, intervals[i][0])
		rangeEnd = max(rangeEnd, intervals[i][1]+1)
	}
	// instead of sorted keys map, just an array
	m := make([]int, rangeEnd+1)
	for i := range intervals {
		m[intervals[i][0]]++
		m[intervals[i][1]+1]--
	}
	rs := 1
	currentNbRooms := 0
	for i := rangeStart; i <= rangeEnd; i++ {
		currentNbRooms += m[i]
		rs = max(rs, currentNbRooms)
	}
	return rs
}
