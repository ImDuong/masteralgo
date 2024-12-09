package binarysearch

func Challenge3152(nums []int, queries [][]int) []bool {
	return isArraySpecial(nums, queries)
}

// split the array into list of continuous special subarrays
// then use binary search to check if each query stay in which subarray
func isArraySpecial(nums []int, queries [][]int) []bool {
	specialSubs := [][]int{}

	start := 0
	for i := 1; i < len(nums); i++ {
		if isDifParity(nums[i-1], nums[i]) {
			if i == len(nums)-1 {
				specialSubs = append(specialSubs, []int{start, i})
			}
			continue
		}
		specialSubs = append(specialSubs, []int{start, i - 1})
		start = i
	}
	rs := make([]bool, len(queries))
	for i := range queries {
		start, end := queries[i][0], queries[i][1]
		if start == end {
			rs[i] = true
			continue
		}
		low, high, mid := 0, len(specialSubs)-1, 0
		for low <= high {
			mid = low + (high-low)/2
			midS, midE := specialSubs[mid][0], specialSubs[mid][1]
			if midS > end {
				high = mid - 1
			} else if midE < start {
				low = mid + 1
			} else {
				low = mid
				break
			}
		}
		foundS, foundE := specialSubs[low][0], specialSubs[low][1]
		if foundS <= start && foundE >= end {
			rs[i] = true
		}
	}
	return rs
}

func isDifParity(a, b int) bool {
	return (a%2 == 0 && b%2 != 0) || (a%2 != 0 && b%2 == 0)
}
