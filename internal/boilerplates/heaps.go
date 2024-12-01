package boilerplates

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
