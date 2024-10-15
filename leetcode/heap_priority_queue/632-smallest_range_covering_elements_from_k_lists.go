package heap_priority_queue

import "container/heap"

func Challenge632(nums [][]int) []int {
	return smallestRange(nums)
}

// Inituition: maintain a min heap contains only one element from each list (e.g., have 3 lists -> size of heap is 3)
// -> from the heap, we always get the range as the maxValue - minValue
// where minValue is the root of the min heap, and maxValue can be tracked individually
// Operation:
//  1. init first element from each list to the min heap; and get the maxValue
//  2. continuing pop the minValue to check if we get a smaller range
//     2.a. whenever we pop element that originally come from a list, push next element in that list to the heap
//     -> this make sure, we always have only one element from each list
//
// Time complexity: O(nlogk), where n is total nb elements in all lists and k is nb of lists
// Space complexity: O(k)
func smallestRange(nums [][]int) []int {
	h := &heap632{}
	heap.Init(h)
	maxValue := -100001
	for i := range nums {
		heap.Push(h, node632{
			value:     nums[i][0],
			noList:    i,
			idxInList: 0,
		})
		maxValue = max(maxValue, nums[i][0])
	}
	startRange, endRange := h.Top().value, maxValue

	// if heap cannot have at least one element from each list, we stop
	for h.Len() == len(nums) {
		minEle := heap.Pop(h).(node632)
		if maxValue-minEle.value < endRange-startRange {
			startRange = minEle.value
			endRange = maxValue
		}
		if minEle.idxInList+1 < len(nums[minEle.noList]) {
			newEle := node632{
				value:     nums[minEle.noList][minEle.idxInList+1],
				noList:    minEle.noList,
				idxInList: minEle.idxInList + 1,
			}
			maxValue = max(maxValue, newEle.value)
			heap.Push(h, newEle)
		}
	}
	return []int{startRange, endRange}
}

type node632 struct {
	value     int
	noList    int
	idxInList int
}

type heap632 []node632

func (h heap632) Len() int           { return len(h) }
func (h heap632) Less(i, j int) bool { return h[i].value < h[j].value }
func (h heap632) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *heap632) Push(x interface{}) {
	*h = append(*h, x.(node632))
}

func (h *heap632) Pop() interface{} {
	poppedItem := (*h)[h.Len()-1]
	*h = (*h)[0 : h.Len()-1]
	return poppedItem
}

func (h heap632) Top() node632 {
	return h[0]
}
