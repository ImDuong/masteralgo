package boilerplates

import "container/heap"

type StringMinHeap []string

func (h StringMinHeap) Len() int           { return len(h) }
func (h StringMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h StringMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *StringMinHeap) Push(x interface{}) {
	*h = append(*h, x.(string))
}

func (h *StringMinHeap) Pop() interface{} {
	poppedItem := (*h)[h.Len()-1]
	*h = (*h)[0 : h.Len()-1]
	return poppedItem
}

func (h *StringMinHeap) Top() string {
	return (*h)[0]
}

// h := &IntMinHeap{}
// heap.Init(h)
type IntMinHeap []int

func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() interface{} {
	poppedItem := (*h)[h.Len()-1]
	*h = (*h)[0 : h.Len()-1]
	return poppedItem
}

func (h *IntMinHeap) Top() int {
	return (*h)[0]
}

type IntMaxHeap []int

func (h IntMaxHeap) Len() int           { return len(h) }
func (h IntMaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMaxHeap) Pop() interface{} {
	poppedItem := (*h)[h.Len()-1]
	*h = (*h)[0 : h.Len()-1]
	return poppedItem
}

func (h *IntMaxHeap) Update(idx int, value int) {
	(*h)[idx] = value
	heap.Fix(h, idx)
}
