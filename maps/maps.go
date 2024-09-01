package maps

import (
	"cmp"
	"slices"
	"sync"
)

type ConcurrentMap[K cmp.Ordered, V any] struct {
	Mu sync.RWMutex
	M  map[K]V
}

func (Map *ConcurrentMap[K, V]) Get(key K) (V, bool) {
	Map.Mu.RLock()
	defer Map.Mu.RUnlock()
	result, exists := Map.M[key]
	return result, exists
}

func (Map *ConcurrentMap[K, V]) Set(key K, value V) {
	Map.Mu.Lock()
	defer Map.Mu.Unlock()
	Map.M[key] = value
}

func (Map *ConcurrentMap[K, V]) Delete(key K) {
	Map.Mu.Lock()
	defer Map.Mu.Unlock()
	delete(Map.M, key)
}

func (Map *ConcurrentMap[K, V]) ClearAll() {
	Map.Mu.Lock()
	defer Map.Mu.Unlock()
	for k := range Map.M {
		delete(Map.M, k)
	}
}

func (Map *ConcurrentMap[K, V]) ListSorted() []K {
	Map.Mu.RLock()
	defer Map.Mu.RUnlock()
	res := make([]K, len(Map.M))
	for key := range Map.M {
		res = append(res, key)
	}
	slices.Sort(res)
	return res
}

func (Map *ConcurrentMap[K, V]) Transform(transFunc func(K, V) V) {
	Map.Mu.Lock();
	defer Map.Mu.Unlock()
	for k := range Map.M {
		Map.M[k] = transFunc(k, Map.M[k]); 
	}
}