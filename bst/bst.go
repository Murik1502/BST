package bst

type BST struct {
	Root   *Node
	size   int
	values []int
}

type Node struct {
	Left  *Node
	Value int
	Right *Node
}

func NewBST() *BST {
	bst := new(BST)
	bst.size = 0
	bst.values = []int{}
	return bst
}

func (bst *BST) Size() int {
	return bst.size
}

func (bst *BST) Values() []int {
	return bst.values
}
