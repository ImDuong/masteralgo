package tree

import (
	"github.com/emirpasic/gods/trees/redblacktree"
)

func Challenge3013(nums []int, k int, dist int) int64 {
	return minimumCost(nums, k, dist)
}

// 1. nums[0] always be counted to the final result
//  2. call nums[i] as the first ele of last subarray
//     -> we need to choose k - 2 ele left so that nums[0] + sum(nums[j]) + nums[i] is min
//     -> with greedy, we just need to find smallest k - 2 ele within this window
//     -> and the size of the window is dist
//  3. For perf, we need to maintain smallest k - 2 ele in the window
//     -> use 2 ordered sets. st1 has smallest k - 2 ele, and st2 has remaining ele
//     -> why 2 sets? If 2 sets separatedly, we can easily track the sum of smallest k - 2 ele throughout modifications
func minimumCost(nums []int, k int, dist int) int64 {
	ctn := NewContainer3013(k - 2)

	// add enough k - 2 ele
	for i := 1; i < k-1; i++ {
		ctn.add(nums[i])
	}

	rs := ctn.sum() + int64(nums[k-1])
	for i := k; i < len(nums); i++ {
		j := i - dist - 1
		// remove the first element of previous window
		if j > 0 {
			ctn.remove(nums[j])
		}
		// expand the window
		ctn.add(nums[i-1])
		rs = min(rs, ctn.sum()+int64(nums[i]))
	}
	return rs + int64(nums[0])
}

type Container3013 struct {
	// maintain smallest k - 2 ele
	st1 *OrderedSet3013
	// store remaining ele
	st2 *OrderedSet3013

	// size of st1
	length int
}

func NewContainer3013(length int) *Container3013 {
	return &Container3013{
		length: length,
		st1:    NewOrderedSet3013(),
		st2:    NewOrderedSet3013(),
	}
}

func (ctn *Container3013) sum() int64 {
	return ctn.st1.sum()
}

func (ctn *Container3013) add(ele int) {
	if ctn.st2.size() > 0 {
		if first2, ok := ctn.st2.first(); ok && ele >= first2 {
			ctn.st2.add(ele)
		} else {
			// add new smaller element to st1
			ctn.st1.add(ele)
		}
	} else {
		ctn.st1.add(ele)
	}
	ctn.balance()
}

func (ctn *Container3013) remove(ele int) {
	isExisted := ctn.st1.remove(ele)
	if !isExisted {
		isExisted = ctn.st2.remove(ele)
	}
	if isExisted {
		ctn.balance()
	}
}

func (ctn *Container3013) balance() {
	// if st1 is not full, move from st2 to st1
	for ctn.st1.size() < ctn.length && ctn.st2.size() > 0 {
		if first, ok := ctn.st2.first(); ok {
			ctn.st2.remove(first)
			ctn.st1.add(first)
		}
	}

	// if st1 is overfilled, then move to st2
	for ctn.st1.size() > ctn.length {
		if last, ok := ctn.st1.last(); ok {
			ctn.st1.remove(last)
			ctn.st2.add(last)
		}
	}
}

type OrderedSet3013 struct {
	// any balanced BST can be used here because of fast insert, delete, shift element
	// in this case, we choose RB-tree for familiar
	tree *redblacktree.Tree

	// to fast check an element is in the list
	eleToFreq map[int]int
	curSize   int
	curSum    int64
}

func NewOrderedSet3013() *OrderedSet3013 {
	return &OrderedSet3013{
		tree:      redblacktree.NewWithIntComparator(),
		eleToFreq: make(map[int]int),
	}
}

func (st *OrderedSet3013) size() int {
	return st.curSize
}

func (st *OrderedSet3013) sum() int64 {
	return st.curSum
}

func (st *OrderedSet3013) add(ele int) {
	st.curSize++
	st.curSum += int64(ele)

	if _, isExisted := st.eleToFreq[ele]; isExisted {
		st.eleToFreq[ele]++
	} else {
		st.eleToFreq[ele] = 1
		st.tree.Put(ele, struct{}{})
	}
}

func (st *OrderedSet3013) remove(ele int) bool {
	if freq, isExisted := st.eleToFreq[ele]; isExisted {
		if freq == 1 {
			delete(st.eleToFreq, ele)
			st.tree.Remove(ele)
		} else {
			st.eleToFreq[ele]--
		}
		st.curSize--
		st.curSum -= int64(ele)
		return true
	}
	return false
}

func (st *OrderedSet3013) first() (int, bool) {
	if st.curSize == 0 {
		return 0, false
	}
	return st.tree.Left().Key.(int), true
}

func (st *OrderedSet3013) last() (int, bool) {
	if st.curSize == 0 {
		return 0, false
	}
	return st.tree.Right().Key.(int), true
}
