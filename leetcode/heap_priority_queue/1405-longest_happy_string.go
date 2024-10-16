package heap_priority_queue

import "container/heap"

func Challenge1405(a int, b int, c int) string {
	return longestDiverseString(a, b, c)
}

func longestDiverseString(a int, b int, c int) string {
	h := &heap1405{}
	heap.Init(h)
	if a > 0 {
		heap.Push(h, node1405{
			value: a,
			chr:   'a',
		})
	}

	if b > 0 {
		heap.Push(h, node1405{
			value: b,
			chr:   'b',
		})
	}

	if c > 0 {
		heap.Push(h, node1405{
			value: c,
			chr:   'c',
		})
	}

	rs := []byte{}
	i := 0
	for h.Len() > 0 {
		node := heap.Pop(h).(node1405)
		if i > 1 && rs[i-2] == rs[i-1] && rs[i-1] == node.chr {
			if h.Len() == 0 {
				break
			}
			secondNode := heap.Pop(h).(node1405)
			rs = append(rs, secondNode.chr)
			secondNode.value--
			if secondNode.value > 0 {
				heap.Push(h, secondNode)
			}
		} else {
			rs = append(rs, node.chr)
			node.value--
		}
		if node.value > 0 {
			heap.Push(h, node)
		}
		i++
	}
	return string(rs)
}

type node1405 struct {
	value int
	chr   byte
}

type heap1405 []node1405

func (h heap1405) Len() int           { return len(h) }
func (h heap1405) Less(i, j int) bool { return h[i].value > h[j].value }
func (h heap1405) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *heap1405) Push(x interface{}) {
	*h = append(*h, x.(node1405))
}

func (h *heap1405) Pop() interface{} {
	poppedItem := (*h)[h.Len()-1]
	*h = (*h)[0 : h.Len()-1]
	return poppedItem
}

func (h heap1405) Top() node1405 {
	return h[0]
}
