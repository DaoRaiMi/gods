package tree

import "fmt"

func postOrderTraverse(root *Node) {
	if root == nil {
		return
	}
	postOrderTraverse(root.left)
	postOrderTraverse(root.right)
	fmt.Printf("%v", root.data)
}
