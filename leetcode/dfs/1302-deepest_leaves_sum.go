package dfs

// binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Challenge1302(root *TreeNode) int {
	return deepestLeavesSum(root)
}

func deepestLeavesSum(root *TreeNode) int {
	rs := 0
	deepestLevel := 0
	dls(root, 0, &rs, &deepestLevel)
	return rs
}

func dls(node *TreeNode, currentLevel int, globalSum *int, deepestLevel *int) {
	if node.Left == nil && node.Right == nil {
		if currentLevel > *deepestLevel {
			*deepestLevel = currentLevel
			*globalSum = node.Val
		} else if currentLevel == *deepestLevel {
			*globalSum += node.Val
		} else {
			return
		}
	}
	if node.Left != nil {
		dls(node.Left, currentLevel+1, globalSum, deepestLevel)
	}
	if node.Right != nil {
		dls(node.Right, currentLevel+1, globalSum, deepestLevel)
	}
}
