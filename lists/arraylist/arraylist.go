package arraylist

const (
	defaultSize = 10
)
// 使用数组来实现的话，会比较麻烦一些还要考虑扩容的问题；
type ArrayList struct {
	elements []interface{}
	size int
}

func New() *ArrayList {
	return &ArrayList{}
}

func (l *ArrayList) Get(index int) (interface{},bool) {
	if index < 0|| l.size <= 0|| index > l.size {
		return nil, false
	}

	return l.elements[index], true
}

func (l *ArrayList) Add(value interface{}) {
	l.elements = append(l.elements, value)
	l.size++
	return
}

func (l *ArrayList) Insert(index int, value interface{}) {
	
}

func (l *ArrayList) Remove(index int) {}

