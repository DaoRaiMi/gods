package stack

import (
	"fmt"
	"testing"
)

func TestNewLinkedStack(t *testing.T) {
	s := NewLinkedStack()
	fmt.Println("Peek empty stack ", s.Peek())
	fmt.Println("Pop empty stack ", s.Pop())

	s.Push("a")
	s.Push("b")
	s.Push("c")

	fmt.Println("Peek ", s.Peek())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println("LinkedStack is empty? ",s.Empty())

	a := [40]int{}
	a[0] = 0
	a[1] = 1
	for i := 2; i<40;i++ {
		a[i] = a[i-1] + a[i-2]
	}
	fmt.Println(a)

	first, second, result := 1,1,0
	for i := 3; i<40;i++ {
		result = first+second
		first = second
		second = result
	}
	fmt.Println(result)
}