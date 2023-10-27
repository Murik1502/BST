package bst

func (root *Node) insert(newNode *Node) {
	if newNode.Value > root.Value {
		if root.Right == nil {
			root.Right = newNode
		} else {
			root.Right.insert(newNode)
		}
	} else if newNode.Value < root.Value {
		if root.Left == nil {
			root.Left = newNode
		} else {
			root.Left.insert(newNode)
		}
	}
}

func (tree *BST) Insert(value int) {
	if tree.Root == nil {
		tree.Root = &Node{nil, value, nil}
	}
	tree.size++
	tree.Root.insert(&Node{nil, value, nil})
}
