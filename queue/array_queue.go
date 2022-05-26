package queue

import "fmt"

const (
	maxQueueSize = 5
)

type Queue interface {
	EnQueue(value interface{})
	DeQueue() interface{}
	Size() int
}

// ArrayQueue 是一个循环队列，必须事先知道队列的大小
type ArrayQueue struct {
	data [maxQueueSize]interface{}
	front int
	rear int
}
func NewArrayQueue() Queue {
	return &ArrayQueue{
		data: [maxQueueSize]interface{}{},
	}
}

func (q *ArrayQueue) Size() int {
	return (q.rear-q.front+maxQueueSize)%maxQueueSize
}

func (q *ArrayQueue) EnQueue(value interface{}) {
	if (q.rear+1)%maxQueueSize == q.front {
		// 队列已经满了
		fmt.Println("queue full...")
		return
	}
	q.data[q.rear] = value
	q.rear = (q.rear+1)%maxQueueSize
}

func (q *ArrayQueue) DeQueue() interface{} {
	if q.front == q.rear {
		// 队列为空
		fmt.Println("queue empty...")
		return nil
	}

	// 从队列头部取出元素，并将front指针向前移动一步
	// 当达到数组的最后面时，再从数组的头部开始
	v := q.data[q.front]
	q.front = (q.front+1)%maxQueueSize
	return v
}