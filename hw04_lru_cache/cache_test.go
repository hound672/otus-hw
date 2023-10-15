package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("new cache item", func(t *testing.T) {
		key := Key(gofakeit.Word())
		value := gofakeit.Word()

		result := newCacheItem(key, value)

		require.Equal(t, key, result.key)
		require.Equal(t, value, result.value)
	})

	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("capacity limit", func(t *testing.T) {
		c := NewCache(3)

		c.Set("key1", 1)
		c.Set("key2", 2)
		c.Set("key3", 3)
		c.Set("key4", 4)

		res, wasInCache := c.Get("key1")
		require.Nil(t, res)
		require.False(t, wasInCache)

		res, wasInCache = c.Get("key4")
		require.Equal(t, 4, res)
		require.True(t, wasInCache)
	})

	t.Run("update queue by getting from cache", func(t *testing.T) {
		c := NewCache(3)

		c.Set("key1", 1)
		c.Set("key2", 2)
		c.Set("key3", 3)

		c.Get("key1")
		c.Get("key2")
		c.Get("key2")
		c.Get("key2")
		c.Get("key1")
		c.Get("key1")

		c.Set("key4", 4)

		result, wasInCache := c.Get("key3")

		require.Nil(t, result)
		require.False(t, wasInCache)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(3)

		c.Set("key1", 1)
		c.Set("key2", 2)
		c.Set("key3", 3)

		c.Clear()

		_, wasInCache := c.Get("key1")
		require.False(t, wasInCache)

		_, wasInCache = c.Get("key2")
		require.False(t, wasInCache)

		_, wasInCache = c.Get("key3")
		require.False(t, wasInCache)
	})
}

func TestCacheMultithreading(t *testing.T) {
	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
