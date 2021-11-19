package data

type BSTNode struct {
	Data  int
	Left  *BSTNode
	Right *BSTNode
}

type BST struct {
	Root *BSTNode
	size int
}

func (b *BST) Size() int {
	return b.size
}

func NewNode(data int) *BSTNode {
	node := new(BSTNode)
	node.Data = data
	return node
}

func NewBST(root *BSTNode) *BST {
	return &BST{Root: root, size: 0}
}

func (node *BSTNode) insert(newNode *BSTNode) {
	if node.Data > newNode.Data {
		if node.Left != nil {
			node.Left.insert(newNode)
		} else {
			node.Left = newNode
		}
	} else if node.Data < newNode.Data {
		if node.Right != nil {
			node.Right.insert(newNode)
		} else {
			node.Right = newNode
		}
	}
}

func (tree *BST) Insert(data int) {
	if tree == nil {
		tree.Root = &BSTNode{}
	}
	tree.size++
	tree.Root.insert(&BSTNode{Data: data, Left: nil, Right: nil})

}

func (tree *BST) Find(data int) *BSTNode {
	node := tree.Root
	for node != nil {
		if node.Data > data {
			node = node.Left
		} else if node.Data < data {
			node = node.Right
		} else if node.Data == data {
			break
		}
	}
	return node
}

func (tree *BST) FindSuccessor(data int) *BSTNode {
	foundNode := tree.Find(data)
	if foundNode.Right != nil {
		return foundNode.Right.IDK()
	}
	return foundNode
}

func (n *BSTNode) IDK() *BSTNode {
	for n != nil {
		if n.Left == nil {
			return n
		}
		n = n.Left
	}
	return n
}

func (root *BSTNode) DeleteMin(data int) {
	root = root.Right
	for root.Left != nil {
		if root.Left.Data == data {
			root.Left = nil
		} else {
			root = root.Left
		}
	}
}
