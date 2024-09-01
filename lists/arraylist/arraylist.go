package arraylist

import (
	"cmp"
	"errors"
	"fmt"
	"slices"
	"sync"

	"github.com/XuananLe/ConcurrentSafeDS/containers"
	"github.com/XuananLe/ConcurrentSafeDS/lists"
)

type ArrayList[T cmp.Ordered] struct {
	lock  sync.RWMutex
	slice []T
}

// Check if implement the List interface
var _ lists.List[int] = &ArrayList[int]{}

func (l *ArrayList[T]) Initialize(initial []T) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.slice = make([]T, len(initial))
	copy(l.slice, initial)
}

func (l *ArrayList[T]) Iter() containers.Iterator[T] {
	return &ArrayListIterator[T] {
			sync.RWMutex{},
			&ArrayList[T]{
				lock: sync.RWMutex{},
				slice: l.slice,
			},
			-1,
		}
}

func (l *ArrayList[T]) Get(index int) (T, error) {
	var zeroValue T
	l.lock.RLock()
	defer l.lock.RUnlock()
	if index <= -1 || index >= len(l.slice) {
		return zeroValue, errors.New("OUT OF BOUND INDEX")
	}
	return l.slice[index], nil
}

func (l *ArrayList[T]) PopBack() error {
	l.lock.Lock()
	defer l.lock.Unlock()
	if len(l.slice) == 0 {
		l.slice = nil
		return errors.New("EMPTY SLICE, CAN'T POP BACK ANY MORE")
	}
	l.slice = l.slice[:len(l.slice)-1]
	return nil
}

func (l *ArrayList[T]) Append(value []T) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.slice == nil || len(l.slice) == 0 {
		l.slice = make([]T, 0)
	}
	l.slice = append(l.slice, value...)

}

func (l *ArrayList[T]) Clone() lists.List[T] {
	l.lock.RLock() 
	defer l.lock.RUnlock()

	newSlice := make([]T, len(l.slice))
	copy(newSlice, l.slice) 

	return &ArrayList[T]{
		slice: newSlice,
	}
}



func (l *ArrayList[T]) Find(value T) (int, bool) {
	l.lock.RLock()
	defer l.lock.RUnlock()
	idx, exist := -1, false
	for index := range l.slice {
		if l.slice[index] == value {
			idx, exist = index, true
		}
	}
	return idx, exist
}

func (l *ArrayList[T]) Delete(value T) {
	l.lock.Lock()
	defer l.lock.Unlock()
	idx, exist := -1, false

	for index := range l.slice {
		if l.slice[index] == value {
			idx, exist = index, true
		}
	}

	if exist {
		if idx == len(l.slice)-1 {
			l.slice = l.slice[:len(l.slice)-1]
		} else {
			copy(l.slice[idx:], l.slice[idx+1:])
			l.slice = l.slice[:len(l.slice)-1]
		}
	}
}

func (l *ArrayList[T]) Set(value T, index int) error {
	l.lock.Lock()
	defer l.lock.Unlock()
	if index <= -1 || index >= len(l.slice) {
		return errors.New("SET VALUE WITH INDEX OUT OF BOUND")
	}
	l.slice[index] = value
	return nil
}

func (l *ArrayList[T]) Swap(index1, index2 int) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.slice[index1], l.slice[index2] = l.slice[index2], l.slice[index1]
}

func (l *ArrayList[T]) Insert(value T, index int) error {
	l.lock.Lock()
	defer l.lock.Unlock()
	if index <= -1 || index >= len(l.slice) {
		return errors.New("INSERT INDEX OUT OF BOUND")
	}
	if index == len(l.slice)-1 {
		l.slice = append(l.slice, value)
		return nil
	}
	var dumVal T
	l.slice = append(l.slice, dumVal)
	copy(l.slice[index+1:], l.slice[index:])
	l.slice[index] = value

	return nil
}

func (l *ArrayList[T]) Sort(stb func(a, b T) int) {
	l.lock.Lock()
	defer l.lock.Unlock()
	slices.SortStableFunc(l.slice, stb)
}
func (l *ArrayList[T]) Size() int {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return len(l.slice)
}

func (l *ArrayList[T]) Clear() {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.slice = make([]T, 0)
}

func (l *ArrayList[T]) String() string {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return fmt.Sprintf("l.slice: %v\n", l.slice)
}

func (l *ArrayList[T]) Values() []T {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return l.slice
}

func (l *ArrayList[T]) Empty() bool {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return len(l.slice) == 0
}
