package containers

type Container[T comparable] interface {
	Empty() bool
	Size() int
	Clear()
	Values() []T
	String() string
}