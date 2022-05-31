package searching

type Node struct {
	left *Node
	right *Node
	data *int
}
// 向二叉排序树中插入结点
func InsertBST(t *Node, n int) {
	// todo
}

/*
	查找二叉排序树t中是否有n
	如果查找到了则返回n所在的结点 和 True
	如果查找失败了，则返回最后一次查找所在的结点 和 False
 */

func SearchBST(t *Node, n int) (*Node,bool) {
	if t == nil {
		// 查找失败
		return nil, false
	}
	if *t.data == n {
		return t, true
	}else if *t.data > n {
		node, ok := SearchBST(t.left, n)
		if node == nil && !ok {
			// 查找失败
			return t.left, false
		}
		return node, ok
	}else {
		node, ok := SearchBST(t.right, n)
		if node == nil && !ok {
			// 查找失败
			return t.right, false
		}
		return node, ok
	}
}
