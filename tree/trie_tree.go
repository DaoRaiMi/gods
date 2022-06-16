package tree

import "fmt"

type TrieTree interface {
	Insert(v string)
	Search(v string) bool
	Size() int
}

type TrieNode struct {
	data int32 //
	isEnding bool
	children map[int32]*TrieNode
}

type TrieRoot struct {
	size int
	root *TrieNode
}

func (t *TrieRoot) Insert(value string) {
	if len(value) == 0 {
		return
	}

	node := t.root
	for _, s := range value {
		fmt.Printf("%+v\n", node)
		v, ok := node.children[s]
		if !ok {
			tmpNode := &TrieNode{data: s,children: make(map[int32]*TrieNode)}
			node.children[s] = tmpNode
			node = tmpNode
			continue
		}
		node = v
	}
	node.isEnding = true
	t.size++
}

func (t *TrieRoot) Search(value string) bool {
	if len(value) == 0 {
		return false
	}

	node := t.root
	for _, s := range value {
		v, ok := node.children[s]
		if !ok {
			return false
		}
		node = v
	}
	if node.isEnding {
		return true
	}
	return false
}

func (t *TrieRoot) Size() int {
	return t.size
}

func NewTrieTree() TrieTree {
	return &TrieRoot{
		size: 0,
		root: &TrieNode{
			data: -1,
			isEnding: false,
			children: make(map[int32]*TrieNode),
		},
	}
}




