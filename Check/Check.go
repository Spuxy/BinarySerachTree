package check

import (
	"math"

	Data "github.com/Spuxy/BST/Data"
)

func IsBST(root *Data.BSTNode) bool {
	return isBST(root, math.MinInt64, math.MaxInt64)
}

func isBST(root *Data.BSTNode, min int, max int) bool {
	if root.Data < min || root.Data > max {
		return false
	}
	if root.Left != nil && !isBST(root.Left, min, root.Data) {
		return false
	}
	if root.Right != nil && !isBST(root.Right, root.Data, max) {
		return false
	}
	return true
}
