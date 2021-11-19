package DeleteNode

import "github.com/Spuxy/BST/Data"

func delete(root *data.BSTNode, data int) *data.BSTNode {
	if (*root).Data > data {
		// Recursion to append to left sub-tree
		root.Left = delete(root.Left, data)
	} else if (*root).Data < data {
		// Recursion to append to right sub-tree
		root.Right = delete(root.Right, data)
	} else {
		// case 1: leaf node
		if root.Left == nil && root.Right == nil {
			return nil
		}
		// case 2: one child
		if root.Left != nil && root.Right == nil {
			return root.Left
		} else if root.Left == nil && root.Right != nil {
			return root.Right
		}
		//case 3: two childs
		tmpNode := findMin(root.Right)
		rightSubTree := root.Right
		tmpNode.Right = rightSubTree
		tmpNode.DeleteMin(tmpNode.Data)
		return tmpNode
	}
	return root
}
func DeleteNodeInternal(bst *data.BST, data int) *data.BSTNode {
	bst.Size--
	return delete(bst.Root, data)
}

func findMin(root *data.BSTNode) *data.BSTNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}
