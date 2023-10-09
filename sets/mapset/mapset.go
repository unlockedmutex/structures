package mapset

import (
	//"github.com/unlockedmutex/structures/list"
	"fmt"
)

// empty structs take no memory
type void struct{}

type hashfunc func(interface{}) int

// keys are ints, values are the members of the set
// users need to pass in a hashfunc
type Set struct {
	members map[int]interface{}
	size    int
	hasher  hashfunc
}

// hashfunc does not need to be an actual hash
// it can be anything that is an int, such as user IDs
func New(hasher hashfunc) *Set {
	var s Set
	s.hasher = hasher
	s.members = make(map[int]interface{})
	return &s
}

func FromSlice(values ...interface{}) *Set {
	var s Set
	for v := range values {
		s.Add(v)
		s.size++
	}
	return &s
}

func (s *Set) ToSlice() []interface{} {
	a := make([]interface{}, s.size)
	i := 0
	for _, value := range s.members {
		a[i] = value
		i++
	}
	return a
}

func (s *Set) Add(value interface{}) {
	fmt.Printf(`%v`, value)
	fmt.Printf(`%v`, s.hasher)
	hashed := s.hasher(value)
	_, exists := s.members[hashed]
	if !exists {
		s.AddUnchecked(value, hashed)
	} else {
		panic(fmt.Sprintf(`Inserting member with duplicate hash value: %d`, s.hasher(value)))
	}
}

func (s *Set) AddUnchecked(value interface{}, hashed int) {
	s.members[hashed] = value
	s.size++
}

func (s *Set) Remove(value interface{}) {
	hashed := s.hasher(value)
	_, exists := s.members[hashed]
	if exists {
		s.RemoveUnchecked(value, hashed)
	} else {
		panic(fmt.Sprintf(`Removing member with hash value not in set: %d`, s.hasher(value)))
	}
}

func (s *Set) RemoveIfExists(value interface{}) bool {
	hashed := s.hasher(value)
	_, exists := s.members[hashed]
	if exists {
		s.RemoveUnchecked(value, hashed)
	}
	return exists
}

func (s *Set) RemoveUnchecked(value interface{}, hashed int) {
	// Will cause panic if removing member not in set
	delete(s.members, hashed)
	s.size--
}

func (s *Set) Update(value interface{}) {
	hashed := s.hasher(value)
	_, exists := s.members[hashed]
	if exists {
		s.UpdateUnchecked(value, hashed)
	} else {
		panic(fmt.Sprintf(`Updating member with hash value not in set: %d`, hashed))
	}
}

func (s *Set) UpdateIfExists(value interface{}) bool {
	hashed := s.hasher(value)
	_, exists := s.members[hashed]
	if exists {
		return s.UpdateUnchecked(value, hashed)
	} else {
		return false
	}
}

func (s *Set) UpdateUnchecked(value interface{}, hashed int) bool {
	_, exists := s.members[hashed]
	s.members[hashed] = value
	return exists
}

func (s *Set) Exists(value interface{}) bool {
	hashed := s.hasher(value)
	_, exists := s.members[hashed]
	return exists
}

func (s *Set) Get(hashed int) interface{} {
	_, exists := s.members[hashed]
	if exists {
		return s.GetUnchecked(hashed)
	} else {
		panic(fmt.Sprintf(`Getting member with hash value not in set: %d`, hashed))
	}
}

func (s *Set) GetIfExists(hashed int) (interface{}, bool) {
	_, exists := s.members[hashed]
	if exists {
		return s.GetUnchecked(hashed), true
	} else {
		return nil, false
	}
}

func (s *Set) GetUnchecked(hashed int) interface{} {
	return s.members[hashed]
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
