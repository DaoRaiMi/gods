package searching

import (
	"fmt"
	"testing"
)

func TestSearchBST(t *testing.T) {
	num1,num3,num5 := &Node{data: new(int)},&Node{data: new(int)},&Node{data: new(int)}
	num7,num9,num11 := &Node{data: new(int)},&Node{data: new(int)},&Node{data: new(int)}
	num13 := &Node{data: new(int)}
	*num1.data,*num3.data,*num5.data = 1,3,5
	num3.left, num3.right = num1,num5

	*num7.data,*num9.data,*num11.data,*num13.data = 7,9,11,13
	num7.left,num7.right = num3,num11

	num11.left,num11.right = num9,num13

	root := num7

	root2 := &Node{data: new(int)}
	*root2.data = 7
	v,ok := SearchBST(root, 13)
	fmt.Printf("%+v,%v\n",*v.data,ok)

	v2,ok2 := SearchBST(root2, 1)
	fmt.Println(*v2.data, ok2)
}

func TestInsertBST(t *testing.T) {
	root := &Node{}
	InsertBST(root, 7)
	InsertBST(root, 1)
	InsertBST(root, 3)
	InsertBST(root, 5)

	fmt.Println(*root.data)
	num1 := root.left
	num3 := num1.right
	num5 := num3.right
	fmt.Println(*num1.data,*num3.data)
	fmt.Println(*num5.data)

}

func showTree(root *Node) {

}