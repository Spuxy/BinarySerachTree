package Height

import (
	Data "github.com/Spuxy/BST/Data"
)

func Height(root *Data.BSTNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	leftHeight := Height(root.Left)
	rightHeight := Height(root.Right)

	if leftHeight >= rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}
