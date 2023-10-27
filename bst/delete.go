package bst

func findIndex(tree *BST, value int) int {

	for i := 0; i < tree.size; i++ {
		if tree.Values()[i] == value {
			return i
		}
	}
	return -1
}

func del(root *Node, value int) bool {
	if root == nil {
		return false
	}
	if root.Value == value {

		if root.Left == nil && root.Right == nil {

			root = nil
			return true
		}

	} else if root.Value > value {
		del(root.Left, value)
	} else {
		del(root.Right, value)
	}
	return false
}

func DelElement(tree *BST, value int) bool {
	println("Deleting element", value)
	if del(tree.Root, value) {
		return true
	} else {
		index := findIndex(tree, value)
		if index >= 0 {
			mas := tree.values
			left := mas[:index]
			right := mas[index+1:]
			newTree := append(left, right...)
			newTr := createNewTree(newTree)
			tree.Root = newTr.Root
			tree.size = newTr.Size()
			tree.values = newTr.values
			return true
		} else {
			return false
		}
	}
}
