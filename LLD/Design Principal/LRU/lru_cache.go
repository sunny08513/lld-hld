package lru

import (
	"container/list"
	"sync"
)

type LRUCache struct {
	capacity int
	lruList  *list.List
	cache    map[int]*list.Element
	mutex    *sync.RWMutex
}

type Item struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	lruList := list.New()
	cache := make(map[int]*list.Element)
	mutex := &sync.RWMutex{}
	return LRUCache{
		capacity: capacity,
		lruList:  lruList,
		cache:    cache,
		mutex:    mutex,
	}
}

func (this *LRUCache) Get(key int) int {
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	if element, ok := this.cache[key]; ok {
		this.lruList.MoveToFront(element)
		return element.Value.(*Item).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if element, ok := this.cache[key]; ok {
		this.lruList.MoveToFront(element)
		element.Value.(*Item).value = value
		return
	}

	if len(this.cache) >= this.capacity {
		backElement := this.lruList.Back()
		delete(this.cache, backElement.Value.(*Item).key)
		this.lruList.Remove(backElement)
	}

	element := this.lruList.PushFront(&Item{key, value})
	this.cache[key] = element
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

 
