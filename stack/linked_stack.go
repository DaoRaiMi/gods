package stack

type Node struct {
	data interface{}
	next *Node
}

// 使用链表来实现一个栈
type LinkedStack struct {
	list *Node // 链表用于存储数据，头结点就表示栈顶;
	count int // 栈中元素个数
}

func NewLinkedStack() Stack {
	return &LinkedStack{
		count: -1, // -1 表示栈为空
	}
}
func (s *LinkedStack) Push(value interface{}) {
	oldTop := s.list
	s.list = &Node{data: value,next: oldTop}
	s.count++
}

func (s *LinkedStack) Pop() interface{} {
	if !s.Empty() {
		v := s.list
		s.list = v.next
		s.count--
		return v.data
	}
	return nil
}

func (s *LinkedStack) Peek() interface{} {
	if !s.Empty() {
		return s.list.data
	}
	return nil
}

func (s *LinkedStack) Empty() bool {
	return s.count < 0
}

// 链栈理论上是不会被存满的，除非内存不够
func (s *LinkedStack) Full() bool {
	return false
}
