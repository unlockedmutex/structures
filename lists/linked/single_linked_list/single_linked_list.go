package single_linked

import (
	//"github.com/unlockedmutex/structures/list"
	"slices"
)

type Node struct {
	data interface{}
	next *Node
}

type List struct {
	head *Node
	size int
}

// Check for all index operations
func (l *List) notInRangePanic(index int) {
	if index >= l.size {
		panic("Index Greater Than List Size")
	}
	if index < 0 {
		panic("Index Below 0")
	}
}

func New() *List {
	var l List
	l.size = 0
	return &l
}

func FromSlice(values ...interface{}) *List {
	var l List
	for v := range values {
		l.Insert(v)
		l.size++
	}
	return &l
}

func (l *List) ToSlice() []interface{} {
	a := make([]interface{}, l.size)
	for curr, i := l.head, 0; curr.next != nil; curr, i = curr.next, i+1 {
		a[i] = curr.data
	}
	return a
}

// Adding, Inserting, Removing
func (l *List) Insert(value interface{}) {
	newNode := &Node{data: value}
	l.size++

	if l.head == nil {
		l.head = newNode
		return
	}

	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode
}

func (l *List) InsertAtIndex(value interface{}, index int) {
	// if index == l.size, we append
	// otherwise normal checks apply
	if index != l.size {
		l.notInRangePanic(index)
	}
	l.InsertAtIndexUnchecked(value, index)
}

// Only use if you GUARANTEE that you will pass in a valid index
// Passing in an invalid index results in undefined behavior
// No guarantees are made for this function if indices are out of bounds
func (l *List) InsertAtIndexUnchecked(value interface{}, index int) {
	newNode := &Node{data: value}
	l.size++

	if l.head == nil && index == 0 {
		l.head = newNode
		return
	}

	curr := l.head
	for i := 0; i < index-1 && curr.next != nil; i++ {
		curr = curr.next
	}

	prevNext := curr.next
	curr.next = newNode
	newNode.next = prevNext
}

func (l *List) Remove(value interface{}) {
	if l.head.data == value {
		l.head = l.head.next
		l.size--
		return
	}

	curr := l.head
	for curr.next != nil && curr.next.data != value {
		curr = curr.next
	}

	if curr.next != nil {
		curr.next = curr.next.next
	} else {
		panic("Value does not exist in list")
	}
	l.size--
}

func (l *List) RemoveAtIndex(index int) {
	l.notInRangePanic(index)
	l.RemoveAtIndexUnchecked(index)
}

// Only use if you GUARANTEE that you will pass in a valid index
// Passing in an invalid index results in undefined behavior
// No guarantees are made for this function if indices are out of bounds
func (l *List) RemoveAtIndexUnchecked(index int) {
	l.size--
	if index == 0 {
		l.head = l.head.next
		return
	}

	curr := l.head
	for i := 0; i < index-1 && curr.next != nil; i++ {
		curr = curr.next
	}

	if curr.next != nil {
		curr.next = curr.next.next
	}
}

// Getting & Finding values
func (l *List) Get(index int) interface{} {
	l.notInRangePanic(index)
	return l.GetUnchecked(index)
}

func (l *List) GetUnchecked(index int) interface{} {
	curr := l.head
	for i := 0; i < index; i++ {
		curr = curr.next
	}
	return curr.data
}

func (l *List) find(value interface{}) int {
	curr := l.head
	for i := 0; i < l.size; i++ {
		if curr.data == value {
			return i
		}
		curr = curr.next
	}
	return -1
}

func (l *List) sort(less func(i, j interface{}) int) {
	//TODO: Optimize to be actual sort.
	//TODO: Measure speed between this + different implementations
	a := l.ToSlice()
	slices.SortFunc(a, less)
	l = FromSlice(a)
}
