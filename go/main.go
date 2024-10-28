package main

import (
	"fmt"
)

func main() {
	root := TreeNode{
		Val: 1,
	}
	node1 := TreeNode{
		Val: 2,
	}
	node2 := TreeNode{
		Val: 3,
	}
	node3 := TreeNode{
		Val: 4,
	}
	root.Left = &node1
	root.Right = &node2
	node1.Right = &node3
	output := printTree(&root)
	fmt.Println("results: ", output)
}
