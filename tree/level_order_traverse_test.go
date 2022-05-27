package tree

import "testing"

func TestLevelOrderTraverse(t *testing.T) {
	tree := createBinTree()
	levelOrderTraverse(tree)
}
