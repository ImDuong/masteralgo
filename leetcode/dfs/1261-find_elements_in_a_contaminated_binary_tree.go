package dfs

import "math"

type FindElements struct {
	root *TreeNode
}

// straight forward solution:
// 1. build up the tree with dfs
// 2. when finding target, use math to trace back the path to the root
// 3. check if the path is valid
func Constructor(root *TreeNode) FindElements {
	rs := FindElements{
		root: root,
	}
	rs.root.Val = 0
	stk := []*TreeNode{root}
	for len(stk) > 0 {
		for _ = range stk {
			cur := stk[0]
			stk = stk[1:]
			if cur.Left != nil {
				cur.Left.Val = cur.Val*2 + 1
				stk = append(stk, cur.Left)
			}
			if cur.Right != nil {
				cur.Right.Val = cur.Val*2 + 2
				stk = append(stk, cur.Right)
			}
		}
	}
	return rs
}

func (this *FindElements) Find(target int) bool {
	// Each layer d contains 2^d nodes, and their indices range from: [2^d - 1, 2^(d+1) - 2]
	// -> The depth of the tree is log2(target + 1)
	depth := int(math.Log2(float64(target + 1)))

	// true: right
	// false: left
	path := make([]bool, depth+1)
	k := target
	for i := depth; i > 0; i-- {
		// check if node is right child
		if k%2 == 0 {
			path[i] = true
		}

		// update k to its parent
		k = (k - 1) / 2
	}
	cur := this.root
	for i := 1; i <= depth; i++ {
		isGoingRight := path[i]
		if isGoingRight {
			if cur.Right == nil {
				return false
			}
			cur = cur.Right
		} else {
			if cur.Left == nil {
				return false
			}
			cur = cur.Left
		}
	}
	return true
}

// Solution 2: Amazing solution by adding 1 to each node to make the path easier to find the target
// Ref: https://leetcode.com/problems/find-elements-in-a-contaminated-binary-tree/solutions/431229/python-special-way-for-find-without-hashset-o-1-space-o-logn-time
