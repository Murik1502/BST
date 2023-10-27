package bst

func find(root *Node, value int) bool {
	if root != nil {
		if value == root.Value {
			return true
		} else if value > root.Value {
			return find(root.Right, value)
		} else {
			return find(root.Left, value)
		}
	}
	return false
}

func (bst *BST) Exists(value int) bool {
	bst.size--
	return find(bst.Root, value)
}
