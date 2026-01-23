package heap_priority_queue

import "container/heap"

func Challenge3510(nums []int) int {
	return minimumPairRemoval(nums)
}

// 1. For each operation, we must merge the pair with minimum sum -> use min heap
// 2. Deleting a pair from the heap is too slow -> we can mark the pair as dirty to avoid merging later.
//     2.1. In a pair, we merge the second to the first -> mark the second as dirty. e.g., (i, j) -> k. Then isMerged[j] = true
// 3. To know when to finish, we track the total number of decreasing pairs thorough the merging process.
func minimumPairRemoval(nums []int) int {
	h := &PairMinHeap{}
	heap.Init(h)
	head := &Node{
		value: int64(nums[0]),
		idx:   0,
	}
	prevNode := head

	// count nb of decreasing pair
	nbDecPair := 0
	for i := 1; i < len(nums); i++ {
		curNode := &Node{
			value: int64(nums[i]),
			idx:   i,
		}
		curNode.prev = prevNode
		prevNode.next = curNode

		heap.Push(h, &Pair{
			x:   prevNode,
			y:   curNode,
			sum: prevNode.value + curNode.value,
		})

		if nums[i] < nums[i-1] {
			nbDecPair++
		}

		prevNode = curNode
	}

	rs := 0
	// track the second node in a pair that has been merged
	// note: in a pair, we merge the second node to the first node
	isMerged := make([]bool, len(nums))
	for nbDecPair > 0 {
		curPair := heap.Pop(h).(*Pair)
		x, y := curPair.x, curPair.y

		// instead of removing pair from the heap, we can mark it as merged and skip the operation
		// there're 3 cases that a pair is dirty
		// 1. x or y is already merged
		// 2. x and y form a pair, however, after that, pair containing y is merged, which change the value of y -> make this old pair is dirty
		// 3. same situation of second case, but if x.value + y.value still equals to curPair.sum, then we can still merge the pair. Now, the pair (x, y) will be duplicated, but we already have the isMerged to make sure the pair is only processed once.
		if isMerged[x.idx] || isMerged[y.idx] || x.value+y.value != curPair.sum {
			continue
		}

		// always has to merge in an operation
		rs++
		if x.value > y.value {
			nbDecPair--
		}

		prevNode, nextNode := x.prev, y.next
		x.next = nextNode
		if nextNode != nil {
			nextNode.prev = x
		}

		// when (i, j) is merged to k, then form a new pair (i - 1, k)
		if prevNode != nil {
			heap.Push(h, &Pair{
				x:   prevNode,
				y:   x,
				sum: prevNode.value + curPair.sum,
			})

			// consider the before and after difference between i - 1 and k
			if prevNode.value > x.value && prevNode.value <= curPair.sum {
				nbDecPair--
			} else if prevNode.value <= x.value && prevNode.value > curPair.sum {
				nbDecPair++
			}
		}

		// when (i, j) is merged to k, then form a new pair (k, j + 1)
		if nextNode != nil {
			heap.Push(h, &Pair{
				x:   x,
				y:   nextNode,
				sum: curPair.sum + nextNode.value,
			})

			// consider the before and after difference between j + 1 and k
			if nextNode.value < y.value && nextNode.value >= curPair.sum {
				nbDecPair--
			} else if nextNode.value >= y.value && nextNode.value < curPair.sum {
				nbDecPair++
			}
		}

		x.value = curPair.sum
		isMerged[y.idx] = true
	}
	return rs
}

type Node struct {
	value int64
	idx   int
	prev  *Node
	next  *Node
}

type Pair struct {
	x   *Node
	y   *Node
	sum int64
}

type PairMinHeap []*Pair

func (h PairMinHeap) Len() int { return len(h) }
func (h PairMinHeap) Less(i, j int) bool {
	if h[i].sum == h[j].sum {
		return h[i].x.idx < h[j].x.idx
	}
	return h[i].sum < h[j].sum
}
func (h PairMinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *PairMinHeap) Push(x interface{}) {
	*h = append(*h, x.(*Pair))
}

func (h *PairMinHeap) Pop() interface{} {
	poppedItem := (*h)[h.Len()-1]
	*h = (*h)[0 : h.Len()-1]
	return poppedItem
}
