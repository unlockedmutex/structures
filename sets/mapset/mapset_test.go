package mapset

import (
	"testing"
)

func intHashFunc(value interface{}) int {
	return 100 - value.(int)
}

func CheckSetValid(t *testing.T, answer []int, s *Set) {
	for _, member := range answer {
		if !s.Exists(member) {
			t.Fatalf(`member %d missing from set`, member)
		}
	}
	if s.size > len(answer) {
		t.Fatalf(`set too big`)
	} else if s.size < len(answer) {
		t.Fatalf(`set too small`)
	}

	if s.size > len(s.members) {
		t.Fatalf(`set size does not match internal size`)
	} else if s.size < len(s.members) {
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
	s.RemoveIfExists(7)
	s.RemoveIfExists(28)
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
	s.Update(7)
	CheckSetValid(t, should_exist, s)
	s.UpdateIfExists(9)
	CheckSetValid(t, should_exist, s)
}
