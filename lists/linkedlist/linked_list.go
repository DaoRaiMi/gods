package linkedlist

import (
	"gods/lists"
)

type Node struct {
	value interface{}
	next *Node
}

type LinkedList struct {
	first, last *Node // 首尾节点
	size int // 链表元素个数，
}

func New(values ...interface{}) lists.List {
	list := &LinkedList{}
	if len(values) > 0 {
		list.Append(values)
	}
	return list
}

// 向链表中追加元素
func (l *LinkedList) Append(values ...interface{}) {
	for _, v := range values {
		if l.size == 0 {
			l.first = &Node{value: v}
			l.last = l.first
		}else {
			last := &Node{value: v}
			l.last.next = last
			l.last = last
		}

		l.size++
	}
}

// 从链表的前面进行追加
// list = ["c","d"]; Prepend("a","b") -> list = ["a","b","c","d"]
func (l *LinkedList) Prepend(values ...interface{}) {
	vLen := len(values)
	for i := vLen-1; i >= 0;i-- {
		node := &Node{value: values[i], next: l.first}
		l.first = node
		if l.size == 0 {
			l.last = l.first
		}
		l.size++
	}
}

// Get
func (l *LinkedList) Get(index int) (interface{}, bool) {
	if !l.withinRange(index) {
		return nil, false
	}

	for n, node := 0, l.first; n < l.size; n, node = n+1, node.next {
		if n == index {
			return node.value, true
		}
	}
	return nil, false
}

// Remove
func (l *LinkedList) Remove(index int) {
	if !l.withinRange(index) {
		return
	}
	var preNode *Node // 待删除节点的上一个节点
	currentNode := l.first
	for i := 0 ; i != index; i, currentNode = i+1, currentNode.next {
		preNode = currentNode
	}

	if currentNode == l.first { // 要删除的是头节点
		l.first = currentNode.next
	}

	if currentNode == l.last { // 要删除的是尾节点
		l.last = preNode
	}

	if preNode != nil { // 删除中间节点
		preNode.next = currentNode.next
	}

	l.size--
	/*
	if index == 0 {
		l.first = l.first.next
		l.size--
		return
	}

	tmpNode := l.first
	for n, node := 0, l.first; n < l.size; n, node = n+1,node.next {
		if n == index {
			tmpNode.next = node.next
			l.size--
			return
		}
		tmpNode = node
	}
	 */
}

// Insert
func (l *LinkedList) Insert(index int, values ...interface{}) {
	if !l.withinRange(index) {
		// 在链表尾部插入，就用append
		if index == l.size {
			l.Append(values...)
			return
		}
	}

	l.size += len(values)

	var preNode *Node
	currentNode := l.first
	for i := 0; i<l.size; i++ {
		if i == index {
			for _, v := range values {
				tmpNode := &Node{value: v}
				if preNode == nil { // 在链表头部插入
					preNode = tmpNode
					l.first = preNode
					continue
				}
				preNode.next = tmpNode
				preNode = tmpNode
			}
			preNode.next = currentNode
			return
		}
		preNode = currentNode
		currentNode = currentNode.next
	}
}

//
func (l *LinkedList) Values() []interface{} {
	v := make([]interface{},l.size,l.size)
	for n, node := 0, l.first; node != nil; n,node = n+1, node.next {
		v[n] = node.value
	}

	return v
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) Clear() {
	l.first = nil
	l.last = nil
	l.size = 0
}

func (l *LinkedList) withinRange(index int) bool {
	return index >= 0 && index < l.size
}

func (l *LinkedList) empty() bool {
	return l.size == 0
}

