package single_linked

import (
	"testing"
	//"github.com/unlockedmutex/structures/list/linked/single_linked"
)

func TestCreate(t *testing.T) {
	l := New()
	if l.head != nil {
		t.Fatalf(`Head is not null`)
	} else if l.size != 0 {
		t.Fatalf(`Size is not 0`)
	}
}

func TestInsert(t *testing.T) {
	l := New()
	l.Insert(4)
	l.Insert(3)
	l.Insert(2)
	l.Insert(1)
	l.Insert(0)
	for i := 0; i < 5; i++ {
		if l.Get(i) != 4-i {
			t.Fatalf(`Error, %d value at %d index`, i, l.Get(i))
		}
	}
	if l.size != 5 {
		t.Fatalf(`incorrect size %d`, l.size)
	}
}

func TestInsertAtIndex(t *testing.T) {
	l := New()
	l.Insert(4)
	l.Insert(3)
	l.Insert(2)
	l.Insert(1)
	l.Insert(0)

	// [4, 3, 2, 1, 0]

	l.InsertAtIndex(100, 2)
	l.InsertAtIndex(200, 4)
	l.InsertAtIndex(300, 7)

	values := [8]int{4, 3, 100, 2, 200, 1, 0, 300}

	for i := 0; i < 8; i++ {
		if l.Get(i) != values[i] {
			t.Fatalf(`Error, %d value at %d index`, l.Get(i), i)
		}
	}
	if l.size != 8 {
		t.Fatalf(`incorrect size %d`, l.size)
	}
}

func TestInsertAtIndexPanic(t *testing.T) {
	l := New()
	l.Insert(4)
	l.Insert(3)
	l.Insert(2)
	l.Insert(1)
	l.Insert(0)

	// [4, 3, 2, 1, 0]

	// [4, 3, 2, 1, 0]
	panicked := false
	panicWrapper := func(l *List) {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		l.InsertAtIndex(100, 9)
	}

	panicWrapper(l)

	if !panicked {
		t.Fatalf(`List.InsertAtIndex should panic when index is > list size`)
	}

	panicked = false
	panicWrapper = func(l *List) {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		l.InsertAtIndex(100, -1)
	}

	panicWrapper(l)

	if !panicked {
		t.Fatalf(`List.InsertAtIndex should panic when index is < 0`)
	}

	values := [5]int{4, 3, 2, 1, 0}

	for i := 0; i < 5; i++ {
		if l.Get(i) != values[i] {
			t.Fatalf(`Error, %d value at %d index`, l.Get(i), i)
		}
	}
	if l.size != 5 {
		t.Fatalf(`incorrect size %d`, l.size)
	}
}

func TestInsertAtIndexUnchecked(t *testing.T) {
	l := New()
	l.Insert(4)
	l.Insert(3)
	l.Insert(2)
	l.Insert(1)
	l.Insert(0)

	// [4, 3, 2, 1, 0]

	l.InsertAtIndexUnchecked(100, 2)
	l.InsertAtIndexUnchecked(200, 4)
	l.InsertAtIndexUnchecked(300, 7)

	values := [8]int{4, 3, 100, 2, 200, 1, 0, 300}

	for i := 0; i < 8; i++ {
		if l.Get(i) != values[i] {
			t.Fatalf(`Error, %d value at %d index`, l.Get(i), i)
		}
	}
	if l.size != 8 {
		t.Fatalf(`incorrect size %d`, l.size)
	}

}

func TestRemove(t *testing.T) {
	l := New()
	l.Insert(4)
	l.Insert(3)
	l.Insert(2)
	l.Insert(1)
	l.Insert(0)

	// [4, 3, 2, 1,  0]
	l.Remove(2)
	l.Remove(4)

	values := [3]int{3, 1, 0}

	for i := 0; i < 3; i++ {
		if l.Get(i) != values[i] {
			t.Fatalf(`Error, %d value at %d index`, l.Get(i), i)
		}
	}
	if l.size != 3 {
		t.Fatalf(`incorrect size %d`, l.size)
	}
}

func TestRemovePanic(t *testing.T) {
	l := New()
	l.Insert(4)
	l.Insert(3)
	l.Insert(2)
	l.Insert(1)
	l.Insert(0)

	// [4, 3, 2, 1, 0]
	panicked := false
	panicWrapper := func(l *List) {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		l.Remove(100)
	}

	panicWrapper(l)

	if !panicked {
		t.Fatalf(`List.Remove should panic when removing a value that doesn't exist`)
	}
}

func TestRemoveAtIndex(t *testing.T) {
	l := New()
	l.Insert(4)
	l.Insert(3)
	l.Insert(2)
	l.Insert(1)
	l.Insert(0)

	// [4, 3, 2, 1,  0]
	l.RemoveAtIndex(2)
	l.RemoveAtIndex(3)

	values := [3]int{4, 3, 1}

	for i := 0; i < 3; i++ {
		if l.Get(i) != values[i] {
			t.Fatalf(`Error, %d value at %d index`, l.Get(i), i)
		}
	}
	if l.size != 3 {
		t.Fatalf(`incorrect size %d`, l.size)
	}
}
