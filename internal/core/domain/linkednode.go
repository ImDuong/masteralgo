package domain

type (
	ListNode struct {
		Val  int
		Next *ListNode
	}

	ITreeNode interface {
		GetListFromTree() []interface{}
		GetTreeFromList([]interface{}) (ITreeNode, error)
	}
)
