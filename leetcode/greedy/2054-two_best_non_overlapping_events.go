package greedy

import "sort"

func Challenge2054(events [][]int) int {
	return maxTwoEvents(events)
}

func maxTwoEvents(events [][]int) int {
	calendar := make([]*eventTime, len(events)*2)
	j := 0
	for i := range events {
		calendar[j] = &eventTime{
			time:    events[i][0],
			isStart: true,
			value:   events[i][2],
		}
		j++
		calendar[j] = &eventTime{
			time:  events[i][1],
			value: events[i][2],
		}
		j++
	}
	sort.Slice(calendar, func(i, j int) bool {
		if calendar[i].time == calendar[j].time {
			return calendar[i].isStart
		}
		return calendar[i].time < calendar[j].time
	})
	rs, curMax := 0, 0
	for i := range calendar {
		if calendar[i].isStart {
			// because the curMax is the maximum value of the event before current event
			// we never afraid of the overlapping
			rs = max(rs, curMax+calendar[i].value)
		} else {
			// we gradually get the maximum value of the event from the past up to this event
			curMax = max(curMax, calendar[i].value)
		}
	}
	return rs
}

type eventTime struct {
	time    int
	isStart bool
	value   int
}
