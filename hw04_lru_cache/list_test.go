package hw04lrucache

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("new list item", func(t *testing.T) {
		value := gofakeit.Word()
		prev := ListItem{}
		next := ListItem{}

		result := NewListItem(value, &prev, &next)

		require.Equal(t, value, result.Value)
		require.Equal(t, &prev, result.Prev)
		require.Equal(t, &next, result.Next)
	})

	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("simple", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushFront(20) // [20, 10]
		l.PushBack(30)  // [20, 10, 30]

		require.Equal(t, 3, l.Len())

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{20, 10, 30}, elems)
	})

	t.Run("remove only single elem", func(t *testing.T) {
		l := NewList()

		elemToRemove := l.PushBack(10)
		l.Remove(elemToRemove)

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("remove head elem", func(t *testing.T) {
		l := NewList()

		_ = l.PushFront(30)             // [20]
		_ = l.PushFront(20)             // [20 30]
		elemToRemove := l.PushFront(10) // [10 20 30]
		l.Remove(elemToRemove)          // [20 30]

		require.Equal(t, 2, l.Len())
		require.Equal(t, 20, l.Front().Value)
		require.Nil(t, l.Front().Prev)
		require.Equal(t, 30, l.Back().Value)
	})

	t.Run("remove tail elem", func(t *testing.T) {
		l := NewList()

		elemToRemove := l.PushFront(30) // [20]
		_ = l.PushFront(20)             // [20 30]
		_ = l.PushFront(10)             // [10 20 30]
		l.Remove(elemToRemove)          // [10 20]

		require.Equal(t, 2, l.Len())
		require.Equal(t, 10, l.Front().Value)

		require.Equal(t, 20, l.Back().Value)
		require.Nil(t, l.Back().Next)
	})

	t.Run("remove middle elem", func(t *testing.T) {
		l := NewList()

		_ = l.PushFront(30)             // [20]
		elemToRemove := l.PushFront(20) // [20 30]
		_ = l.PushFront(10)             // [10 20 30]
		l.Remove(elemToRemove)          // [10 30]

		require.Equal(t, 2, l.Len())

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{10, 30}, elems)
	})

	t.Run("move elem to front", func(t *testing.T) {
		l := NewList()

		elemToMove := l.PushFront(30) // [30]
		_ = l.PushFront(20)           // [20 30]
		_ = l.PushFront(10)           // [10 20 30]

		l.MoveToFront(elemToMove) // [30 10 20]

		require.Equal(t, 3, l.Len())

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{30, 10, 20}, elems)
		require.Nil(t, l.Back().Next)
		require.Nil(t, l.Front().Prev)

		// check that elem was moved and now its Next And Prev elements referenced to correct elems
		require.Nil(t, elemToMove.Prev)
		require.Equal(t, elemToMove.Next.Value, 10)
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

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

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}
