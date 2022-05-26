package tree

import (
	"testing"
)

func TestPreOrderTraverse(t *testing.T) {
	root := createBinTree()
	preOrderTraverse(root)
}
