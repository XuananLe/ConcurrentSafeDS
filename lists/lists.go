package lists

import (
	"cmp"

	"github.com/XuananLe/ConcurrentSafeDS/containers"
)

type List[T cmp.Ordered] interface {
	Get(index int) (T, error)
	Initialize(initial []T)
	Append(value []T)
	Find(value T) (int, bool)
	Delete(value T)
	Clone() List[T]
	Insert(value T, index int) error
	Set(value T, index int) error
	Sort(stb func(a, b T) int)
	Swap(index1, index2 int)
	PopBack() error
	Iter() containers.Iterator[T]
	containers.Container[T]
	
}
