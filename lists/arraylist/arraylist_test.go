package arraylist

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayList(t *testing.T) {
	t.Run("Initialize and Get", func(t *testing.T) {
		list := &ArrayList[int]{}
		list.Initialize([]int{1, 2, 3, 4, 5})

		val, err := list.Get(2)
		assert.NoError(t, err)
		assert.Equal(t, 3, val)

		_, err = list.Get(10)
		assert.Error(t, err)
	})

	t.Run("PopBack", func(t *testing.T) {
		list := &ArrayList[int]{}
		list.Initialize([]int{1, 2, 3})

		err := list.PopBack()
		assert.NoError(t, err)
		assert.Equal(t, 2, list.Size())

		list.PopBack()
		list.PopBack()
		err = list.PopBack()
		assert.Error(t, err)
	})

	t.Run("Append", func(t *testing.T) {
		list := &ArrayList[int]{}
		list.Append([]int{1, 2, 3})
		list.Append([]int{4, 5})

		assert.Equal(t, 5, list.Size())
		val, _ := list.Get(4)
		assert.Equal(t, 5, val)
	})

	t.Run("Find", func(t *testing.T) {
		list := &ArrayList[int]{}
		list.Initialize([]int{1, 2, 3, 4, 5})

		index, found := list.Find(3)
		assert.True(t, found)
		assert.Equal(t, 2, index)

		_, found = list.Find(10)
		assert.False(t, found)
	})

	t.Run("Delete", func(t *testing.T) {
		list := &ArrayList[int]{}
		list.Initialize([]int{1, 2, 3, 4, 5})

		list.Delete(3)
		assert.Equal(t, 4, list.Size())

		val, _ := list.Get(2)
		assert.Equal(t, 4, val)

		list.Delete(5)
		assert.Equal(t, 3, list.Size())
	})

	t.Run("Set", func(t *testing.T) {
		list := &ArrayList[int]{}
		list.Initialize([]int{1, 2, 3, 4, 5})

		err := list.Set(10, 2)
		assert.NoError(t, err)

		val, _ := list.Get(2)
		assert.Equal(t, 10, val)

		err = list.Set(20, 10)
		assert.Error(t, err)
	})

	t.Run("Insert", func(t *testing.T) {
		list := &ArrayList[int]{}
		list.Initialize([]int{1, 2, 3, 4, 5})

		err := list.Insert(10, 2)
		assert.NoError(t, err)
		assert.Equal(t, 6, list.Size())

		val, _ := list.Get(2)
		assert.Equal(t, 10, val)

		err = list.Insert(20, 10)
		assert.Error(t, err)
	})

	t.Run("Sort", func(t *testing.T) {
		list := &ArrayList[int]{}
		list.Initialize([]int{5, 2, 8, 1, 9})

		list.Sort(func(a, b int) int { return a - b })

		expected := []int{1, 2, 5, 8, 9}
		for i, v := range expected {
			val, _ := list.Get(i)
			assert.Equal(t, v, val)
		}
	})

	t.Run("Clear and Empty", func(t *testing.T) {
		list := &ArrayList[int]{}
		list.Initialize([]int{1, 2, 3})

		assert.False(t, list.Empty())

		list.Clear()
		assert.True(t, list.Empty())
		assert.Equal(t, 0, list.Size())
	})

	t.Run("String", func(t *testing.T) {
		list := &ArrayList[int]{}
		list.Initialize([]int{1, 2, 3})

		str := list.String()
		assert.Contains(t, str, "l.slice: [1 2 3]")
	})

	t.Run("Values", func(t *testing.T) {
		list := &ArrayList[int]{}
		list.Initialize([]int{1, 2, 3})

		values := list.Values()
		assert.Equal(t, []int{1, 2, 3}, values)
	})

	t.Run("Concurrent operations", func(t *testing.T) {
		list := &ArrayList[int]{}
		list.Initialize([]int{1, 2, 3, 4, 5})

		var wg sync.WaitGroup
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(val int) {
				defer wg.Done()
				list.Append([]int{val})
			}(i)
		}

		for i := 0; i < 50; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				list.PopBack()
			}()
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			list.Delete(1)
		}()

		wg.Wait()

		assert.Equal(t, 54, list.Size())
	})

	t.Run("Test Serialize", func(t *testing.T) {
		list := &ArrayList[int]{}
		list.Initialize([]int{})
		list.Append([]int{12, 13, 14})
		_, err := list.MarshalJson()
		assert.Nil(t, err)
	})

}



func TestArrayListIterator(t *testing.T) {
	t.Run("Basic Iterator Operations", func(t *testing.T) {
		list := &ArrayList[int]{}
		list.Initialize([]int{1, 2, 3, 4, 5})
		iter := list.Iter()

		// Test Begin and Value
		iter.Begin()
		assert.Equal(t, 1, iter.Value())

		// Test Next and HasNext
		iter.Next()
		assert.Equal(t, 2, iter.Value())
		assert.True(t, iter.HasNext())

		// Test Last
		iter.Last()
		assert.Equal(t, 5, iter.Value())
		assert.False(t, iter.HasNext())

		// Test Prev
		iter.Prev()
		assert.Equal(t, 4, iter.Value())
	})

	t.Run("AdvanceBy and Nth", func(t *testing.T) {
		list := &ArrayList[int]{}
		list.Initialize([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
		iter := list.Iter()

		// Test AdvanceBy
		iter.AdvanceBy(3)
		assert.Equal(t, 3, iter.Value())

	})

	t.Run("Sum and Count", func(t *testing.T) {
		list := &ArrayList[int]{}
		list.Initialize([]int{1, 2, 3, 4, 5})
		iter := list.Iter()

		// Test Sum
		assert.Equal(t, 15, iter.Sum())

		// Test Count
		assert.Equal(t, 5, iter.Count())
	})

	t.Run("Edge Cases", func(t *testing.T) {
		list := &ArrayList[int]{}
		list.Initialize([]int{1})
		iter := list.Iter()

		iter.Next();
		assert.Panics(t, func() { iter.Next() })

		// Test Prev on first element
		iter.Begin()
		assert.Panics(t, func() { iter.Prev() })
	})

	t.Run("Concurrent Operations", func(t *testing.T) {
		list := &ArrayList[int]{}
		list.Initialize([]int{1, 2, 3, 4, 5})
		iter := list.Iter()

		// Simulate concurrent reads
		for i := 0; i < 5; i++ {
			go func() {
				iter.Value()
				iter.Sum()
				iter.Count()
			}()
		}
	})

	t.Run("Change the original slice", func(t *testing.T) {
		list := &ArrayList[int]{}
		list.Initialize([]int{1,2,3,4})
		iter := list.Iter();
		list.Set(12, 0);
		iter.Next();
		assert.Equal(t, 12, iter.Value())
	})

}