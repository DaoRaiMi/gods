package tree

import "fmt"

func inOrderTraverse(root *Node) {
	if root == nil {
		return
	}
	inOrderTraverse(root.left)
	fmt.Printf("%v", root.data)
	inOrderTraverse(root.right)
}
