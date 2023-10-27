package bst

func createNewTree(mas []int) *BST {
	bst := new(BST)
	bst.size = len(mas)
	bst.values = mas
	if bst.Root == nil {
		bst.Root = &Node{nil, 0, nil}
	}
	rebalanced(bst.Root, mas)
	return bst
}

func rebalanced(root *Node, mas []int) {

	if len(mas) >= 3 {
		index := (len(mas) - 1) / 2
		root.Value = mas[index]

		left := mas[:index]
		right := mas[index+1:]
		if root.Left == nil {
			root.Left = &Node{nil, 0, nil}
		}

		rebalanced(root.Left, left)
		if root.Right == nil {
			root.Right = &Node{nil, 0, nil}
		}
		rebalanced(root.Right, right)
	} else if len(mas) == 2 {
		root.Value = mas[0]
		if root.Right == nil {
			root.Right = &Node{nil, 0, nil}
		}
		rebalanced(root.Right, mas[1:])
	} else if len(mas) == 1 {
		root.Value = mas[0]
		return
	}
}
