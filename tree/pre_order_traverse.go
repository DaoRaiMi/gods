package tree

import "fmt"

/*
二叉树的前序遍历
 */

func preOrderTraverse(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%v", root.data)
	preOrderTraverse(root.left)
	preOrderTraverse(root.right)
}
