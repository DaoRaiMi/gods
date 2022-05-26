package stack

const (
	// 默认栈的大小
	defaultStackSize = 5
)

// 栈的抽象接口
type Stack interface {
	Push(v interface{})
	Pop() interface{}
	Peek() interface{}
	Full() bool
	Empty() bool
}

type ArrayStack struct {
	// 用数组来实际的栈，使用下标0来作为栈底
	data [defaultStackSize]interface{}
	// 栈顶的下标
	top int
}

func NewArrayStack() Stack {
	return &ArrayStack{
		data: [defaultStackSize]interface{}{},
		top: -1, // -1 表示栈为空
	}
}

// 向栈中添加一个元素
func (s *ArrayStack) Push(value interface{}) {
	if !s.Full() {
		s.top++
		s.data[s.top] = value
	}
}

// 获取栈顶元素，并将其从栈中删除
func (s *ArrayStack) Pop() interface{} {
	if s.Empty() {
		return nil
	}
	v := s.data[s.top]
	s.data[s.top] = nil
	s.top--
	return v
}

// 获取栈顶元素，但是不删除它
func (s *ArrayStack) Peek() interface{} {
	if s.Empty() {
		return nil
	}

	return s.data[s.top]
}

// 栈是否为空
func (s *ArrayStack) Empty() bool {
	return s.top < 0
}

// 栈是否满了
func (s *ArrayStack) Full() bool {
	return s.top >= defaultStackSize-1
}
