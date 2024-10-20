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
