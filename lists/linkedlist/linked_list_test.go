package linkedlist

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	l1 := New()
	l1.Append(100,400,500)
	l1.Insert(0,200,300)
	fmt.Println("size = ", l1.Size(),"values = ", l1.Values())

	l2 := New()
	// Append
	l2.Append("300",400,"500")
	fmt.Println(l2.Values())

	// Prepend
	l2.Prepend("100",200)
	fmt.Println(l2.Values())

	// Get
	fmt.Println(l2.Get(2))

	// Remove
	//l2.Remove(4)
	l2.Remove(0)
	fmt.Println(l2.Values())

	// Insert

	// Clear

	// Size()

}