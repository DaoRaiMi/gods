package tree

import (
	"fmt"
	"testing"
)

func TestNewTrieTree(t *testing.T) {
	prefixTree := NewTrieTree()
	prefixTree.Insert("abc")
	prefixTree.Insert("abcdef")
	prefixTree.Insert("ab世")
	prefixTree.Insert("ab世界")

	fmt.Println(prefixTree.Search("abc"))
}
