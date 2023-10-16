package hw04lrucache

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
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	_, wasInCache := c.items[key]

	elem := c.queue.PushFront(value)
	c.items[key] = elem

	return wasInCache
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if el, ok := c.items[key]; ok {
		return el.Value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
}
