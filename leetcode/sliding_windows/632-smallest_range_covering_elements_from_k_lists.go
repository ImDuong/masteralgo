package slidingwindows

import "sort"

func Challenge632(nums [][]int) []int {
	return smallestRanges(nums)
}

// merge all list into single one, and try to find the sublist having minimum range,
// where it contains at least one number from each of the k lists
func smallestRanges(nums [][]int) []int {
	// merge into a single non-decreasing list
	nbEles := 0
	for i := range nums {
		nbEles += len(nums[i])
	}
	merged := make([]Ele, nbEles)
	mIdx := 0
	for i := range nums {
		for j := range nums[i] {
			merged[mIdx] = Ele{
				value:  nums[i][j],
				listNo: i,
			}
			mIdx++
		}
	}
	sort.Slice(merged, func(i, j int) bool {
		return merged[i].value < merged[j].value
	})

	// use a dynamic sliding window to find the minimum range sublist
	left := 0

	// use cnt to track how many lists the window containing number from
	cnt := 0
	freq := make(map[int]int, len(nums))
	startRange, endRange := 0, int(^uint(0)>>1)

	// window expanding to the end of the list
	for right := 0; right < len(merged); right++ {
		freq[merged[right].listNo]++
		if freq[merged[right].listNo] == 1 {
			cnt++
		}

		// shrink the window until the window does not contain at least one number from each of the k lists
		for cnt == len(nums) {
			if merged[right].value-merged[left].value < endRange-startRange {
				startRange = merged[left].value
				endRange = merged[right].value
			}
			freq[merged[left].listNo]--
			if freq[merged[left].listNo] == 0 {
				cnt--
			}
			left++
		}
	}
	return []int{startRange, endRange}
}

type Ele struct {
	value  int
	listNo int
}
