package linkedlist

func Challenge2626(m int, n int, head *ListNode) [][]int {
	return spiralMatrix(m, n, head)
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	rs := make([][]int, m)
	for i := range rs {
		rs[i] = make([]int, n)
	}
	i, j := 0, 0
	// init: go right first
	x, y := 1, 0
	// additional start buffer
	k := 0
	nbSlots := m * n
	for head != nil || nbSlots > 0 {
		if head == nil {
			rs[i][j] = -1
		} else {
			rs[i][j] = head.Val
		}
		if i == k+1 && j == k && x == 0 && y == -1 {
			// reach top left corner -> go right
			x = 1
			y = 0

			// go inside
			k++
		} else if i == k && j == n-1-k && x == 1 && y == 0 {
			// reach top right corner -> go down
			x = 0
			y = 1
		} else if i == m-1-k && j == n-1-k && x == 0 && y == 1 {
			// reach bottom right corner -> go left
			x = -1
			y = 0
		} else if i == m-1-k && j == k && x == -1 && y == 0 {
			// reach bottom left corner -> go up
			x = 0
			y = -1
		}
		i += y
		j += x

		if head != nil {
			head = head.Next
		}
		nbSlots--
	}
	return rs
}
