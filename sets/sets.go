package sets

type Set interface {
	ToSlice() []interface{}
	Add(value interface{})
	Remove(value interface{})
	Exists(value interface{}) bool
	Get(hashed int) interface{}
	Union(s2 *Set) *Set
	Intersection(s2 *Set) *Set
	Difference(s2 *Set) *Set
	Size() int
	String() string
}
