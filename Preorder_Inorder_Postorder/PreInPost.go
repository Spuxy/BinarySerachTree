package PreInPost

import (
	"fmt"

	"github.com/Spuxy/BST/Data"
)

func Preorder(tree *data.BST) {
	preorder(tree.Root)
}

func preorder(node *data.BSTNode) {
	fmt.Println("VISITED: ", node.Data)
	if node.Left != nil {
		preorder(node.Left)
	}
	if node.Right != nil {
		preorder(node.Right)
	}
}

func Inorder(tree *data.BST) {
	inorder(tree.Root)
}

func inorder(node *data.BSTNode) {
	if node.Left != nil {
		inorder(node.Left)
	}
	fmt.Println("VISITED: ", node.Data)
	if node.Right != nil {
		inorder(node.Right)
	}
}

func Postorder(tree *data.BST) {
	postorder(tree.Root)
}

func postorder(node *data.BSTNode) {
	if node.Left != nil {
		postorder(node.Left)
	}
	if node.Right != nil {
		postorder(node.Right)
	}
	fmt.Println("VISITED: ", node.Data)
}
