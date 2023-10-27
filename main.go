package main

import (
	"awesomeProject/bst"
	"fmt"
)

func main() {
	test := bst.NewBST()

	test.Insert(4)
	test.Insert(1)
	test.Insert(22)
	test.Insert(2321)
	test.Insert(3)
	test.Insert(23)
	test.Insert(234)
	test.Insert(445)

	fmt.Println(bst.DFS(test))
	fmt.Println(bst.DelElement(test, 4))
	fmt.Println("new tree:", bst.DFS(test))
	fmt.Println(bst.DelElement(test, 2321))
	fmt.Println("new tree:", bst.DFS(test))
}
