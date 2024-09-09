package linkedlist

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Challenge1367(head *ListNode, root *TreeNode) bool {
	return isSubPath(head, root)
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	return dfs(root, head) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func dfs(curTreeNode *TreeNode, curHead *ListNode) bool {
	if curHead == nil {
		return true
	}

	if curTreeNode == nil {
		return false
	}

	return curTreeNode.Val == curHead.Val && (dfs(curTreeNode.Left, curHead.Next) || dfs(curTreeNode.Right, curHead.Next))
}
