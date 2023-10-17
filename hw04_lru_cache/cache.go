package hw04lrucache

import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	m        sync.RWMutex
}

type cacheItemType struct {
	key   Key
	value interface{}
}

func newCacheItem(key Key, value interface{}) cacheItemType {
	return cacheItemType{
		key:   key,
		value: value,
	}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.m.Lock()
	defer c.m.Unlock()

	cacheItem := newCacheItem(key, value)

	if item, ok := c.items[key]; ok {
		item.Value = cacheItem
		c.queue.MoveToFront(item)
		return true
	}

	elem := c.queue.PushFront(cacheItem)
	c.items[key] = elem

	if c.queue.Len() > c.capacity {
		// stack overflow :)
		elemToRemove := c.queue.Back()

		cacheItem, ok := elemToRemove.Value.(cacheItemType)
		if !ok {
			// it's dev's issue, so we can use panic here
			panic("Get: item is not cacheItemType")
		}

		delete(c.items, cacheItem.key)
		c.queue.Remove(elemToRemove)
	}

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.m.RLock()
	defer c.m.RUnlock()

	item, ok := c.items[key]
	if !ok {
		return nil, false
	}

	c.queue.MoveToFront(item)

	cacheItem, ok := item.Value.(cacheItemType)
	if !ok {
		// it's dev's issue, so we can use panic here
		panic("Get: item is not cacheItemType")
	}

	return cacheItem.value, true
}

func (c *lruCache) Clear() {
	c.m.Lock()
	defer c.m.Unlock()

	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}
