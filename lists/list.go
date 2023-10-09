package list

type List interface {
	Get(index int) (interface{}, bool)
	Remove(index int)
	RemoveIndex(index int)
	RemoveValue(values ...interface{})
	Add(values ...interface{})
	Contains(values ...interface{}) bool
	// Sort(comparator utils.Comparator)
	Swap(index1, index2 int)
	Insert(index int, values ...interface{})
	Set(index int, value interface{})

	// containers.Container
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
	// String() string
}
