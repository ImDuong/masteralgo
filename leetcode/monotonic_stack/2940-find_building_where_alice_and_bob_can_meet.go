package monotonicstack

func Challenge2940(heights []int, queries [][]int) []int {
	return leftmostBuildingQueries(heights, queries)
}

// preprocess the queries, so that we will answer them when traversing heights with monotonic stack
func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	// processed queries
	pq := make(map[int][]Query)
	rs := make([]int, len(queries))
	for i := range queries {
		a, b := queries[i][0], queries[i][1]
		if a == b {
			rs[i] = a
			continue
		}
		if b < a {
			a, b = b, a
		}
		if heights[b] > heights[a] {
			rs[i] = b
			continue
		}

		// store bigger height between a & b
		pq[b] = append(pq[b], Query{
			index:  i,
			height: heights[a],
		})

		// init rs
		rs[i] = -1
	}

	// decreasing monotonic stack storing index of building
	stk := []int{}
	// traverse from the end to construct decreasing monotonic stack for finding next higher building (next greater element - NGE)
	for i := len(heights) - 1; i >= 0; i-- {
		for _, query := range pq[i] {
			found := search(query.height, stk, heights)
			if found >= 0 && found < len(stk) {
				rs[query.index] = stk[found]
			}
		}

		for len(stk) > 0 && heights[i] >= heights[stk[len(stk)-1]] {
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, i)
	}
	return rs
}

type Query struct {
	index  int
	height int
}

// return index in stk for element x that x > target
func search(targetHeight int, stk, heights []int) int {
	low, high, mid := 0, len(stk)-1, 0
	for low <= high {
		mid = low + (high-low)/2
		if heights[stk[mid]] <= targetHeight {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return high
}
