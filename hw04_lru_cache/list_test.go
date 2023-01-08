package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		// Add nodes to front/back
		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]

		require.Equal(t, 3, l.Len())
		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, 30, l.Back().Value)

		// Delete a node in middle
		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]

		require.Equal(t, 2, l.Len())
		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, 30, l.Back().Value)

		// Add nodes to front/back
		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		// Move front/back nodes to front
		// Initial: [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)

		// Move non front/back nodes to front
		// Initial: [70, 80, 60, 40, 10, 30, 50]
		l.MoveToFront(l.Front().Next) // [80, 70, 60, 40, 10, 30, 50]
		l.MoveToFront(l.Back().Prev)  // [30, 80, 70, 60, 40, 10, 50]

		elems = make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{30, 80, 70, 60, 40, 10, 50}, elems)

		// Edge cases

		// One node list, remove the node
		l = NewList()

		l.PushFront(10) // [10]
		l.Remove(l.Front())

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())

		// Two nodes list, remove front node
		l = NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.Remove(l.Front())

		require.Equal(t, 1, l.Len())
		require.Equal(t, 20, l.Front().Value)
		require.Nil(t, l.Front().Next)
		require.Nil(t, l.Front().Prev)

		// Two nodes list, remove back node
		l = NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.Remove(l.Back())

		require.Equal(t, 1, l.Len())
		require.Equal(t, 10, l.Front().Value)
		require.Nil(t, l.Front().Next)
		require.Nil(t, l.Front().Prev)

		// One node list, move to front
		l = NewList()

		l.PushFront(10) // [10]
		l.MoveToFront(l.Front())

		require.Equal(t, 1, l.Len())
		require.Equal(t, l.Back(), l.Front())
		require.Equal(t, 10, l.Front().Value)

		// Two nodes list, move front to front
		l = NewList()

		l.PushFront(10)          // [10]
		l.PushBack(20)           // [10, 20]
		l.MoveToFront(l.Front()) // [10, 20]

		require.Equal(t, 2, l.Len())
		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, 20, l.Back().Value)

		// Two nodes list, move back to front
		l = NewList()

		l.PushFront(10)         // [10]
		l.PushBack(20)          // [10, 20]
		l.MoveToFront(l.Back()) // [20, 10]

		require.Equal(t, 2, l.Len())
		require.Equal(t, 20, l.Front().Value)
		require.Equal(t, 10, l.Back().Value)
	})
}
