package _46_lru_cache

/*
146. LRU Cache
Medium

3850

157

Favorite

Share
Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.

get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.

The cache is initialized with a positive capacity.

Follow up:
Could you do both operations in O(1) time complexity?

Example:

LRUCache cache = new LRUCache( 2  )

cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // returns 1
cache.put(3, 3);    // evicts key 2
cache.get(2);       // returns -1 (not found)
cache.put(4, 4);    // evicts key 1
cache.get(1);       // returns -1 (not found)
cache.get(3);       // returns 3
cache.get(4);       // returns 4
*/
type cacheItem struct {
	key, value, getCount int
	next                 *cacheItem
}

type LRUCache struct {
	head     *cacheItem
	capacity int
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{}
	cache.capacity = capacity
	return cache
}

func (this *LRUCache) Get(key int) int {
	if this.head == nil {
		return -1
	}
	current := this.head
	for current != nil {
		if current.key == key {
			this.bringKeyFront(current)
			return current.value
		}
		current = current.next
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	exit := this.getItem(key)
	if exit != nil {
		exit.value = value
		this.bringKeyFront(exit)
		return
	}
	item := cacheItem{key, value, 0, nil}
	item.next = this.head
	this.head = &item
	this.evicts()
}

func (this *LRUCache) getItem(key int) *cacheItem {
	if this.head == nil {
		return nil
	}
	current := this.head
	for current != nil {
		if current.key == key {
			this.bringKeyFront(current)
			return current
		}
		current = current.next
	}
	return nil
}

func (this *LRUCache) bringKeyFront(item *cacheItem) {
	if item == nil || this.head == nil || this.head == item {
		return
	}
	new := &cacheItem{}
	head := new
	head.key = item.key
	head.value = item.value

	current := this.head
	for current != nil {
		next := current.next
		current.next = nil
		if current.key != item.key {
			head.next = current
			head = head.next
		}
		current = next
	}
	this.head = new
}

func (this *LRUCache) evicts() {
	current := this.head
	for count := 0; count < this.capacity-1 && current != nil; count++ {
		current = current.next
	}
	if current != nil {
		current.next = nil
	}
}
