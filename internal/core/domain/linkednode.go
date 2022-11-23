package domain

type (
	ListNode struct {
		Val  int
		Next *ListNode
	}

	TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
)
