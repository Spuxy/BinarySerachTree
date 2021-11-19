package main

import (
	PreInPost "github.com/Spuxy/BST/Preorder_Inorder_Postorder"
	"github.com/Spuxy/BST/data"
)

func main() {
	// Lets create node as a root !
	root := data.NewNode(30)

	// Lets create BST itself with root
	bst := data.NewBST(root)

	// Binart tree insert
	//                10
	//                /\
	//  lesser -->   4  12 <-- greater
	// lets complete BST!
	bst.Insert(10)
	bst.Insert(19)
	bst.Insert(15)
	bst.Insert(14)
	bst.Insert(5)
	bst.Insert(7)
	bst.Insert(6)

	// Binary tree traversal
	// Preoder --> |root| |left subtree| |right subtree|
	PreInPost.Preorder(bst)

	// Inorder --> |left subtree| |root| |right subtree|
	PreInPost.Inorder(bst)

	// Postorder --> |left subtree| |right subtree| |root|
	PreInPost.Postorder(bst)
}
