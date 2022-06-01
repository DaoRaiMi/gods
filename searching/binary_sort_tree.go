package searching

type Node struct {
	left *Node
	right *Node
	data *int
}
// 向二叉排序树中插入结点
func InsertBST(t *Node, n int) {
	// 当二叉查找树为空时，直接把n当作树的根结点。
	currentNode := Node{data: new(int)}
	*currentNode.data = n
	if t == nil || t.data == nil {
		*t = currentNode
		return
	}

	node, ok := SearchBST(t, n)
	if !ok {
		// 未找到，则把n当作node的子结点
		if *node.data > n {
			node.left = &currentNode
		}else if *node.data < n {
			node.right = &currentNode
		}
	}
}

/*
	从二叉查找树中删除值为n的结点
	在删除结点时，分为以下几种情况
	1. 被删除结点是叶子结点，可以直接删除;
	2. 被删除结点只有左子树或只有右子树，此时只需要让其子树成为当前结点父结点的孩子即可
	3. 被删除结点左右子树都不为空，那么需要在左右子树中找一个最接近当前结点的结点来替换被删除的结点。最接近其实就是在
		左子树中找到最靠右的那结点，或者在右子树中找到那个最靠左的结点。
 */

func DeleteBST(t *Node, n int) {
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
		if t.left == nil {
			// left为空时，没必要再向下查找了，直接返回当前所在的结点。
			return t, false
		}
		return SearchBST(t.left, n)
	}else {
		if t.right == nil {
			return t, false
		}
		return SearchBST(t.right, n)
	}
}
