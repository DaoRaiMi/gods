package queue

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

// 链式队列可以不用考虑队列的大小，理论上空间是无限的。它的进出队列也都是O（1）的操作。唯一的不足就是，
// 在队列长度相同时，会比循环队列占用更多的内存
type LinkedQueue struct {
	front *Node
	rear *Node
	size int
}

func NewLinkedQueue() Queue {
	// 链式队列中front, rear一开始都指向链的头结点
	head := &Node{}
	return &LinkedQueue{
		front: head,
		rear: head,
		size: 0,
	}
}

func (q *LinkedQueue) Size() int {
	return q.size
}

// 直接在rear的后面进行拼接
func (q *LinkedQueue) EnQueue(value interface{}) {
	node := &Node{data: value}
	q.rear.next = node
	q.rear = node
	q.size++
}

// 直接把头结点后面的一个结点返回即可
func (q *LinkedQueue) DeQueue() interface{} {
	if q.front == q.rear {
		fmt.Println("queue empty...")
		return nil
	}
	firstNode := q.front.next
	if firstNode.next == nil { // 如果当前元素没有子结点时，表示队列在此之后为空。
		q.rear = q.front  // 将rear指向头结点
		q.front.next = nil
	}else {
		q.front.next = firstNode.next
	}
	q.size--

	return firstNode.data
}

