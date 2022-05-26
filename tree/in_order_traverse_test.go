package tree

import "testing"

func TestInOrderTraverse(t *testing.T) {
	root := createBinTree()
	inOrderTraverse(root)
}
