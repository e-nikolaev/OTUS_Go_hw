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

// Cache element must have key for 1:1 relation to map.
type cacheElement struct {
	key   Key
	value interface{}
}

func (lc *lruCache) Set(key Key, value interface{}) bool {
	newCacheElementValue := cacheElement{
		key:   key,
		value: value,
	}
	cacheItem, ok := lc.items[key]
	if ok {
		cacheItem.Value = newCacheElementValue
		lc.queue.MoveToFront(cacheItem)
		return ok
	}
	if lc.queue.Len() == lc.capacity {
		lastCacheItem := lc.queue.Back()
		lastCacheElement, _ := lastCacheItem.Value.(cacheElement) // TODO: Parse _ if needed
		lc.queue.Remove(lastCacheItem)
		delete(lc.items, lastCacheElement.key)
	}
	newCacheItem := lc.queue.PushFront(newCacheElementValue)
	lc.items[key] = newCacheItem

	return ok
}

func (lc *lruCache) Get(key Key) (interface{}, bool) {
	cacheItem, ok := lc.items[key]
	if !ok {
		return nil, false
	}
	lc.queue.MoveToFront(cacheItem)
	CacheElement, _ := cacheItem.Value.(cacheElement) // TODO: Parse _ if needed
	return CacheElement.value, ok
}

func (lc *lruCache) Clear() {
	lc.items = make(map[Key]*ListItem, lc.capacity)
	lc.queue.Front().Next = lc.queue.Back()
	lc.queue.Back().Prev = lc.queue.Front()
	lc.queue.Remove(lc.queue.Front())
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
