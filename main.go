package main

import "fmt"

type lnode struct {
	key, val   int
	prev, next *lnode
}

// LRUCache is a fixed-capacity key-value cache evicting the least recently
// used entry: a hash map for O(1) lookup, a doubly linked list for O(1)
// recency ordering, sentinels at both ends so no nil checks exist.
type LRUCache struct {
	capacity   int
	items      map[int]*lnode
	head, tail *lnode // sentinels: head.next is most recent, tail.prev least
}

// NewLRUCache returns a cache holding at most capacity entries (capacity ≥ 1).
func NewLRUCache(capacity int) *LRUCache {
	head, tail := &lnode{}, &lnode{}
	head.next, tail.prev = tail, head
	return &LRUCache{capacity: capacity, items: make(map[int]*lnode), head: head, tail: tail}
}

func (c *LRUCache) unlink(n *lnode) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (c *LRUCache) pushFront(n *lnode) {
	n.next = c.head.next
	n.prev = c.head
	c.head.next.prev = n
	c.head.next = n
}

// Get returns the value for key and marks it most recently used.
func (c *LRUCache) Get(key int) (int, bool) {
	n, ok := c.items[key]
	if !ok {
		return 0, false
	}
	c.unlink(n)
	c.pushFront(n)
	return n.val, true
}

// Put inserts or updates key, evicting the least recently used entry when full.
func (c *LRUCache) Put(key, val int) {
	if n, ok := c.items[key]; ok {
		n.val = val
		c.unlink(n)
		c.pushFront(n)
		return
	}
	if len(c.items) == c.capacity {
		lru := c.tail.prev
		c.unlink(lru)
		delete(c.items, lru.key)
	}
	n := &lnode{key: key, val: val}
	c.items[key] = n
	c.pushFront(n)
}

func main() {
	c := NewLRUCache(2)
	c.Put(1, 1)
	c.Put(2, 2)
	fmt.Println(c.Get(1)) // 1
	c.Put(3, 3)           // evicts 2
	fmt.Println(c.Get(2)) // miss
}
