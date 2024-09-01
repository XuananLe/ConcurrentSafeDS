package arraylist

import (
	"cmp"
	"sync"

	"github.com/XuananLe/ConcurrentSafeDS/containers"
)

var _ containers.Iterator[int] = &ArrayListIterator[int]{}

type ArrayListIterator[T cmp.Ordered] struct {
	lock  sync.RWMutex
	l     *ArrayList[T]
	index int // Iteration State
}

func (it *ArrayListIterator[T]) Next() {
	it.lock.Lock()
	defer it.lock.Unlock()
	if it.index+1 <= len(it.l.slice)-1 {
		it.index += 1
	} else {
		panic("OUT OF BOUND INDEX ITERATOR")
	}
}

func (it *ArrayListIterator[T]) Begin() {
	it.lock.Lock()
	defer it.lock.Unlock()
	it.index = 0
}

func (it *ArrayListIterator[T]) Prev() {
	it.lock.Lock()
	defer it.lock.Unlock()
	if it.index >= 1 {
		it.index -= 1
	} else {
		panic("OUT OF BOUND INDEX ITERATOR")
	}
}

func (it *ArrayListIterator[T]) Value() T {
	it.lock.RLock()
	defer it.lock.RUnlock()
	return it.l.slice[it.index]
}

func (it *ArrayListIterator[T]) Sum() T {
	it.lock.RLock()
	defer it.lock.RUnlock()
	var sum T
	for _, value := range it.l.slice {
		sum += value
	}
	return sum
}

func (it *ArrayListIterator[T]) Product() T {
	panic("NOT IMPLEMENTED")
}

func (it *ArrayListIterator[T]) HasNext() bool {
	it.lock.RLock()
	defer it.lock.RUnlock()
	return it.index != len(it.l.slice)-1
}

func (it *ArrayListIterator[T]) AdvanceBy(n int) {
	it.lock.Lock()
	defer it.lock.Unlock()
	it.index = min(it.index+n, len(it.l.slice)-1)
}

func (it *ArrayListIterator[T]) Count() int {
	it.lock.RLock()
	defer it.lock.RLock()
	return len(it.l.slice)
}

func (it *ArrayListIterator[T]) Last() {
	it.lock.Lock()
	defer it.lock.Unlock()
	it.index = len(it.l.slice) - 1
}

func (it *ArrayListIterator[T]) Nth(n int) {
	it.lock.Lock()
	defer it.lock.Unlock()
	if 0 <= n && n <= len(it.l.slice)-1 {
		it.index = n
	} else {
		panic("OUT OF BOUND INDEX")
	}
}
