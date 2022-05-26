package lists

type List interface {
	Get(index int) (interface{}, bool)
	Append(values ...interface{})
	Prepend(values ...interface{})
	Insert(index int, value ...interface{})
	Remove(index int)
	Values() []interface{}
	Size() int
	Clear()
}