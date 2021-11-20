package LevelOrder

import (
	"fmt"

	Data "github.com/Spuxy/BST/Data"
)

var Queue []*Data.BSTNode

func GetLevelOrder(root *Data.BSTNode) {
	getLevelOrder(root)
	// GetLevelOrderWithWhile(root)
}
func getLevelOrder(root *Data.BSTNode) {
	if root.Left != nil {
		Enqueue(root.Left)
	}
	if root.Right != nil {
		Enqueue(root.Right)
	}
	fmt.Println("Level Order Traversal ====> ", root.Data)
	if node := Dequeue(); node != nil {
		getLevelOrder(node)
	}
}

func Enqueue(node *Data.BSTNode) {
	Queue = append(Queue, node)
}

func Dequeue() *Data.BSTNode {
	if len(Queue) <= 0 {
		return nil
	}
	node := Queue[0]
	Queue = Queue[1:]
	return node
}

// Another implementation with while

func GetLevelOrderWithWhile(root *Data.BSTNode) {

	Enqueue(root)

	for {

		node := Dequeue()

		if node == nil {
			break
		}

		fmt.Println(node.Data)

		if node.Left != nil {
			Enqueue(node.Left)
		}

		if node.Right != nil {
			Enqueue(node.Right)
		}
	}

}
