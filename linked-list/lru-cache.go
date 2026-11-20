package main

/*
Implement the Least Recently Used (LRU) cache class LRUCache. The class should support the following operations
LRUCache(int capacity) Initialize the LRU cache of size capacity.
int get(int key) Return the value corresponding to the key if the key exists, otherwise return -1.
void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the introduction of the new pair causes the cache to exceed its capacity, remove the least recently used key.
A key is considered used if a get or a put operation is called on it.
Ensure that get and put each run in O(n) average time complexity.
*/

/*
Hash map + doubly linked list
The hash map gives O(1) lookup from key to node.
The doubly linked list keeps usage order. The most recently used node sits next to head, and the least recently used node sits next to tail.
Dummy head and tail nodes remove the edge cases for empty lists and for the first and last nodes.
*/
type node struct {
	key, value int
	prev, next *node
}

type LRUCache struct {
	capacity int
	cache    map[int]*node
	head     *node // dummy: head.next is the most recently used
	tail     *node // dummy: tail.prev is the least recently used
}

func Constructor(capacity int) LRUCache {
	head := &node{}
	tail := &node{}
	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*node, capacity),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) remove(n *node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (this *LRUCache) addToFront(n *node) {
	n.next = this.head.next
	n.prev = this.head

	this.head.next.prev = n
	this.head.next = n
}

func (this *LRUCache) Get(key int) int {
	n, exists := this.cache[key]
	if !exists {
		return -1
	}

	// mark as most recently used
	this.remove(n)
	this.addToFront(n)

	return n.value
}

func (this *LRUCache) Put(key, value int) {
	if n, exists := this.cache[key]; exists {
		n.value = value

		this.remove(n)
		this.addToFront(n)
		return
	}

	// Evict the least recently used node when full
	if len(this.cache) == this.capacity {
		lru := this.tail.prev

		this.remove(lru)
		delete(this.cache, lru.key)
	}

	n := &node{key: key, value: value}
	this.cache[key] = n
	this.addToFront(n)
}
