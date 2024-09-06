package dfs

func Challenge2265(root *TreeNode) int {
	return averageOfSubtree(root)
}

func averageOfSubtree(root *TreeNode) int {
	rs := 0
	aos(&rs, root)
	return rs
}

func aos(nbAvgNodes *int, curNode *TreeNode) (branchSum int, nbNodesInBranch int) {
	// tree always has at least 1 node -> no need to check for case curNode is nil

	// self-count for this current node
	branchSum += curNode.Val
	nbNodesInBranch++
	if curNode.Left != nil {
		nextBranchSum, nextBranchNbNodes := aos(nbAvgNodes, curNode.Left)
		branchSum += nextBranchSum
		nbNodesInBranch += nextBranchNbNodes
	}
	if curNode.Right != nil {
		nextBranchSum, nextBranchNbNodes := aos(nbAvgNodes, curNode.Right)
		branchSum += nextBranchSum
		nbNodesInBranch += nextBranchNbNodes
	}
	if branchSum/nbNodesInBranch == curNode.Val {
		*nbAvgNodes++
	}
	return
}
