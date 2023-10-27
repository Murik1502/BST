package bst

import "fmt"

func DFS(root *BST) []int {
	treeClear(root)
	TreePrint(root, root.Root)
	fmt.Println()
	return root.values
}

func treeClear(root *BST) {
	root.values = []int{}
}

func TreePrint(tree *BST, root *Node) {
	if root == nil {
		return
	}
	TreePrint(tree, root.Left)
	tree.values = append(tree.values, root.Value)
	TreePrint(tree, root.Right)
}
