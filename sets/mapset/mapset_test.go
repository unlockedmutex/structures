package mapset

import (
	"testing"
)

func intHashFunc(value interface{}) interface{} {
	return 100 - value.(int)
}

func CheckSetValid(t *testing.T, answer []int, s *Set) {
	for _, member := range answer {
		if !s.Exists(member) {
			t.Fatalf(`member %d missing from set`, member)
		}
	}
	if s.Size() > len(answer) {
		t.Fatalf(`set too big`)
	} else if s.Size() < len(answer) {
		t.Fatalf(`set too small`)
	}

	if s.Size() > len(s.members) {
		t.Fatalf(`set size does not match internal size`)
	} else if s.Size() < len(s.members) {
		t.Fatalf(`set size does not match internal size`)
	}
}

func TestAdd(t *testing.T) {
	s := New(intHashFunc)
	s.Add(5)
	s.Add(7)
	s.Add(9)
	should_exist := []int{5, 7, 9}
	CheckSetValid(t, should_exist, s)
}

func TestRemove(t *testing.T) {
	s := New(intHashFunc)
	s.Add(5)
	s.Add(7)
	s.Add(9)
	should_exist := []int{5, 7, 9}
	CheckSetValid(t, should_exist, s)
	s.Remove(7)
	should_exist = []int{5, 9}
	CheckSetValid(t, should_exist, s)
}

func TestRemoveIfExists(t *testing.T) {
	s := New(intHashFunc)
	s.Add(5)
	s.Add(7)
	s.Add(9)
	should_exist := []int{5, 7, 9}
	CheckSetValid(t, should_exist, s)
	found := s.Remove(7)
	if !found {
		t.Fatalf(`remove returned false for member in set`)
	}
	found = s.Remove(28)
	if found {
		t.Fatalf(`remove returned true for member not in set`)
	}
	should_exist = []int{5, 9}
	CheckSetValid(t, should_exist, s)
}

func TestUpdate(t *testing.T) {
	s := New(intHashFunc)
	s.Add(5)
	s.Add(7)
	s.Add(9)
	should_exist := []int{5, 7, 9}
	CheckSetValid(t, should_exist, s)
	found := s.Add(7)
	if !found {
		t.Fatalf(`remove returned false for member in set`)
	}
	CheckSetValid(t, should_exist, s)
	found = s.Add(28)
	if found {
		t.Fatalf(`remove returned true for member not in set`)
	}
	should_exist = []int{5, 7, 9, 28}
	CheckSetValid(t, should_exist, s)
}

func TestExists(t *testing.T) {
	s := New(intHashFunc)
	s.Add(5)
	s.Add(7)
	s.Add(9)
	should_exist := []int{5, 7, 9}
	CheckSetValid(t, should_exist, s)
	if s.Exists(10) {
		t.Fatalf(`Exists returned true for a member not in the set`)
	}
	if s.Exists(0) {
		t.Fatalf(`Exists returned true for a member not in the set`)
	}
	s.Remove(7)
	if s.Exists(7) {
		t.Fatalf(`Exists returned true for a member not in the set`)
	}
	s.Remove(5)
	if s.Exists(5) {
		t.Fatalf(`Exists returned true for a member not in the set`)
	}
}

func TestUnion(t *testing.T) {
	s1 := New(intHashFunc)
	s1.Add(5)
	s1.Add(7)
	s1.Add(9)
	s2 := New(intHashFunc)
	s2.Add(100)
	s2.Add(200)
	s2.Add(300)
	s3 := s1.Union(s2)
	CheckSetValid(t, []int{5, 7, 9, 100, 200, 300}, s3)
}
