package containers

type Iterator[T comparable] interface {
	Next()
	Prev() 
	Value() T
	Sum() T
	Product() T
	HasNext() bool
	AdvanceBy(n int)
	Count() int 
	Last()
	Begin()
}