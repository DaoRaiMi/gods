package tree

import "testing"

func TestPostOrderTraverse(t *testing.T) {
	root := createBinTree()
	postOrderTraverse(root)
}
