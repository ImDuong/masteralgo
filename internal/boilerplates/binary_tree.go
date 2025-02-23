package boilerplates

import "masteralgo/internal/core/domain"

// Construct tree from preorder & postorder traversal
// The properties are: first node in preorder and the last node of the postorder are all the root node
// There are several ways to achieve this

// Method 1: For each index in the preorder traversal, we can find the index of the same element in the postorder traversal
// to find the left & right subtree of the current node and then recursively build the tree
// e.g. preorder = [1, 2, 4, 5, 3, 6, 7], postorder = [4, 5, 2, 6, 7, 3, 1])
// 1. With root node as 1, the cur node is 2
// 2. Look at the postorder, the left subtree is [4, 5, 2] and the right subtree is [6, 7, 3]
// The result of this method is always a full binary tree
// Time complexity: O(n^2) in the worst case, O(n) in the best case
func constructFromPrePost(preorder []int, postorder []int) *domain.TreeNode {
	// pointer to track cur pos in preorder
	preIdx := 0

	return makeTree(preorder, postorder, &preIdx, 0, len(postorder)-1)
}

func makeTree(preorder, postorder []int, preIdx *int, postLeft, postRight int) *domain.TreeNode {
	if postLeft > postRight || *preIdx >= len(preorder) {
		return nil
	}

	root := &domain.TreeNode{Val: preorder[*preIdx]}
	*preIdx++

	if postLeft == postRight {
		return root
	}

	nextChildVal := preorder[*preIdx]
	// find the pos of next child in postorder
	var nextChildPos int
	for i := postLeft; i <= postRight; i++ {
		if postorder[i] == nextChildVal {
			nextChildPos = i
			break
		}
	}

	if nextChildPos > postRight {
		return nil
	}
	root.Left = makeTree(preorder, postorder, preIdx, postLeft, nextChildPos)
	root.Right = makeTree(preorder, postorder, preIdx, nextChildPos+1, postRight-1)

	return root
}

// Above solution works for general case where values of node can be duplicated
// In the case where values of nodes are unique, we can bring up other optimzed solutions

// Method 2: Use recursive function to build the tree
// iterate through preorder to dfs
// while using postorder to find the bottom node to stop dfs -> then move up
// Note: in recursive solution, to move up, just return the node
// Time Complexity: because we only iterate through the preorder once, the time complexity is O(n)
func constructFromPrePostMethod2(preorder []int, postorder []int) *domain.TreeNode {
	preIdx, postIdx := 0, 0
	return makeTreeRecursively(preorder, postorder, &preIdx, &postIdx)
}

func makeTreeRecursively(preorder, postorder []int, preIdx *int, postIdx *int) *domain.TreeNode {
	root := &domain.TreeNode{Val: preorder[*preIdx]}

	// move to next node for dfs
	*preIdx++

	// the node in postIdx will always be the bottom node of the current subtree we're building
	// hence, if we don't meet that node yet, we continue to dfs
	if root.Val != postorder[*postIdx] {
		root.Left = makeTreeRecursively(preorder, postorder, preIdx, postIdx)
	}

	if root.Val != postorder[*postIdx] {
		root.Right = makeTreeRecursively(preorder, postorder, preIdx, postIdx)
	}

	// finish adding a bottom node -> move to next node
	*postIdx++
	return root
}

// Method 3: The same with method 2 but use stack to build the tree
func constructFromPrePostMethod3(preorder []int, postorder []int) *domain.TreeNode {
	stk := []*domain.TreeNode{{Val: preorder[0]}}

	// i, j tracks cur pos for preorder & postorder
	for i, j := 1, 0; i < len(preorder); i++ {
		newNode := &domain.TreeNode{Val: preorder[i]}

		// if we meet a same value in pre & post, it means we've reached the bottom node of the current subtree
		// hence we move up or right (by popping the stack)
		for stk[len(stk)-1].Val == postorder[j] {
			stk = stk[:len(stk)-1]
			j++
		}
		if stk[len(stk)-1].Left == nil {
			stk[len(stk)-1].Left = newNode
		} else {
			stk[len(stk)-1].Right = newNode
		}

		stk = append(stk, newNode)
	}
	return stk[0]
}

// Method 4: To optimize solution 1, we can use a hashmap to store the index of each element in the postorder traversal
func constructFromPrePostMethod4(preorder []int, postorder []int) *domain.TreeNode {
	preIdx := 0
	postDict := make(map[int]int)
	for i := range postorder {
		postDict[postorder[i]] = i
	}
	return makeTreeWithMap(preorder, postorder, &preIdx, 0, len(postorder)-1, postDict)
}

func makeTreeWithMap(preorder, postorder []int, preIdx *int, postLeft, postRight int, postDict map[int]int) *domain.TreeNode {
	if postLeft > postRight || *preIdx >= len(preorder) {
		return nil
	}

	root := &domain.TreeNode{Val: preorder[*preIdx]}
	*preIdx++

	if postLeft == postRight {
		return root
	}

	childNodePos := postDict[preorder[*preIdx]]
	if childNodePos > postRight {
		return nil
	}

	root.Left = makeTreeWithMap(preorder, postorder, preIdx, postLeft, childNodePos, postDict)
	root.Right = makeTreeWithMap(preorder, postorder, preIdx, childNodePos+1, postRight-1, postDict)

	return root
}
