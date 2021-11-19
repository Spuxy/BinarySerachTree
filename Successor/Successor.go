// package Successor

// import "github.com/Spuxy/BST/data"

// func (tree *data.BST) Find(data int) *Data.BSTNode {
// 	node := tree.root
// 	for node != nil {
// 		if node.Data > data {
// 			node = node.Left
// 		} else if node.Data < data {
// 			node = node.Right
// 		} else if node.Data == data {
// 			break
// 		}
// 	}
// 	return node
// }

// func (tree *Data.BST) FindSuccessor(data int) *BSTNode {
// 	foundNode := tree.Find(data)
// 	if foundNode.Right != nil {
// 		return foundNode.Right.IDK()
// 	}
// 	return foundNode
// }

// func (n *Data.BSTNode) IDK() *Data.BSTNode {
// 	for n != nil {
// 		if n.Left == nil {
// 			return n
// 		}
// 		n = n.Left
// 	}
// 	return n
// }
