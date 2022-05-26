package stack

import (
	"fmt"
	"testing"
)

func TestArrayStack(t *testing.T) {
	s := NewArrayStack()
	s.Push("a")
	s.Push("b")
	s.Push("c")
	s.Push("d")
	s.Push("e")
	 s.Push("f")

	fmt.Println("empty?", s.Empty())
	fmt.Println("full?", s.Full())
	fmt.Println(s.Peek())
	fmt.Printf("%+v\n",s)
	fmt.Println(s.Pop())
	fmt.Printf("%+v\n",s)

}
