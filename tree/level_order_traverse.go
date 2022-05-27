package tree

import (
	"fmt"
	"gods/queue"
)

// 使用队列来实现层序遍历
func levelOrderTraverse(root *Node) {
	if root == nil {
		return
	}

	q := queue.NewLinkedQueue()
	q.EnQueue(root)
	for q.Size() > 0 {
		head := q.DeQueue()
		node := head.(*Node)
		fmt.Printf("%v", node.data)
		if node.left != nil {
			q.EnQueue(node.left)
		}
		if node.right != nil {
			q.EnQueue(node.right)
		}
	}
}