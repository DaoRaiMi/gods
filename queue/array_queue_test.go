package queue

import (
	"fmt"
	"testing"
)

func TestArrayQueue(t *testing.T) {
	aq := NewArrayQueue()
	aq.EnQueue("a")
	aq.EnQueue("b")
	aq.EnQueue("c")
	aq.EnQueue("d")
	aq.EnQueue("e")

	fmt.Println(aq.(*ArrayQueue))

	fmt.Println(aq.DeQueue())

	fmt.Println(aq.(*ArrayQueue))

	aq.EnQueue("f")

	fmt.Println(aq.(*ArrayQueue))

	aq.EnQueue("g")

	fmt.Println(aq.DeQueue())
	fmt.Println(aq.(*ArrayQueue))

	aq.EnQueue("g")
	fmt.Println(aq.(*ArrayQueue))

	fmt.Println(aq.DeQueue(),aq.DeQueue(),aq.DeQueue(),aq.DeQueue())
	fmt.Println(aq.(*ArrayQueue))
	aq.DeQueue()
}