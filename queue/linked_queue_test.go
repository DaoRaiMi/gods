package queue

import (
	"fmt"
	"testing"
)

func TestLinkedQueue(t *testing.T) {
	lq := NewLinkedQueue()
	lq.EnQueue("a")
	lq.EnQueue("b")
	lq.EnQueue("c")
	lq.EnQueue("d")

	fmt.Println(lq.Size())
	fmt.Println(lq.DeQueue())
	fmt.Println(lq.DeQueue())
	fmt.Println(lq.DeQueue())
	fmt.Println(lq.DeQueue())
	fmt.Println(lq.DeQueue())
	fmt.Println(lq.DeQueue())

	v, _ := lq.(*LinkedQueue)
	fmt.Printf("%+v\n",v)
}