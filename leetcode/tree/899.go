package tree

import (
	"masteralgo/internal/boilerplates"
	"masteralgo/internal/core/domain"
)

func Challenge899(preorder []int, postorder []int) *domain.TreeNode {
	return boilerplates.ConstructFromPrePostMethod4(preorder, postorder)
}
