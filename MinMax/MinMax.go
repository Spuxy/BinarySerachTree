package MinMax

import (
	Data "github.com/Spuxy/BST/Data"
)

func GetMin(root *Data.BSTNode) int {
	if root == nil {
		return -1
	}
	if root.Left == nil {
		return root.Data
	}
	return GetMin(root.Left)
}

func GetMax(root *Data.BSTNode) int {
	if root == nil {
		return -1
	}
	if root.Right == nil {
		return root.Data
	}
	return GetMax(root.Right)
}
