package funcset

import (
	//"github.com/unlockedmutex/structures/list"
	"fmt"
)

// empty structs take no memory
type void struct{}

type hashfunc func(interface{}) interface{}

// keys are ints, values are the members of the set
// users need to pass in a hashfunc
type Set struct {
	members map[interface{}]interface{}
	hasher  hashfunc
}

// hashfunc does not need to return an actual hash
// it can be anything that is an int, such as user IDs
func New(hasher hashfunc) *Set {
	var s Set
	s.hasher = hasher
	s.members = make(map[interface{}]interface{})
	return &s
}

func FromSlice(hasher hashfunc, values ...interface{}) *Set {
	s := New(hasher)
	for value := range values {
		hashed := s.hasher(value)
		s.members[hashed] = value
	}
	return s
}

func (s *Set) ToSlice() []interface{} {
	a := make([]interface{}, s.Size())
	i := 0
	for _, value := range s.members {
		a[i] = value
		i++
	}
	return a
}

// Returns True if member already existed
func (s *Set) Add(value interface{}) bool {
	hashed := s.hasher(value)
	_, exists := s.members[hashed]
	s.members[hashed] = value
	return exists
}

// Returns True if member already existed
func (s *Set) Remove(value interface{}) bool {
	hashed := s.hasher(value)
	_, exists := s.members[hashed]
	if exists {
		delete(s.members, hashed)
	}
	return exists
}

func (s *Set) Exists(value interface{}) bool {
	hashed := s.hasher(value)
	_, exists := s.members[hashed]
	return exists
}

func (s *Set) Get(hashed interface{}) interface{} {
	member, _ := s.members[hashed]
	return member
}

// If any hashes collide, s2 values overwride s1 values
func (s1 *Set) Union(s2 *Set) *Set {
	// Weird but checks if function addresses are the same
	if fmt.Sprintf(`%v`, s1.hasher) != fmt.Sprintf(`%v`, s2.hasher) {
		panic(fmt.Sprintf(`Union of two sets with different hash functions`))
	}
	ret := New(s1.hasher)
	for hashed, member := range s1.members {
		ret.members[hashed] = member
	}
	for hashed, member := range s2.members {
		ret.members[hashed] = member
	}
	return ret
}

// If any hashes collide, s2 values overwride s1 values
func (s1 *Set) Intersection(s2 *Set) *Set {
	// Weird but checks if function addresses are the same
	if fmt.Sprintf(`%v`, s1.hasher) != fmt.Sprintf(`%v`, s2.hasher) {
		panic(fmt.Sprintf(`Intersection of two sets with different hash functions`))
	}
	ret := New(s1.hasher)
	for hashed, member := range s2.members {
		if s1.Exists(hashed) {
			ret.members[hashed] = member
		}
	}
	return ret
}

func (s1 *Set) Difference(s2 *Set) *Set {
	// Weird but checks if function addresses are the same
	if fmt.Sprintf(`%v`, s1.hasher) != fmt.Sprintf(`%v`, s2.hasher) {
		panic(fmt.Sprintf(`Intersection of two sets with different hash functions`))
	}
	ret := New(s1.hasher)
	for hashed, member := range s1.members {
		if !s2.Exists(hashed) {
			ret.members[hashed] = member
		}
	}
	for hashed, member := range s2.members {
		if !s1.Exists(hashed) {
			ret.members[hashed] = member
		}
	}
	return ret
}

func (s *Set) Size() int {
	return len(s.members)
}

func (s *Set) String() string {
	printstr := ""
	first := false
	for _, value := range s.members {
		if !first {
			printstr = printstr + ", " + fmt.Sprintf(`%v`, value)
		} else {
			printstr = fmt.Sprintf(`%v`, value)
		}
		first = true
	}
	return printstr
}
