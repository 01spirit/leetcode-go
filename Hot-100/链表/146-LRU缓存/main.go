package main

type LRUCache struct {
	cap        int
	size       int
	cache      map[int]*DlinkedList
	head, tail *DlinkedList
}

type DlinkedList struct {
	key, value int
	prev, next *DlinkedList
}

func InitDLinkedList(k, v int) *DlinkedList {
	return &DlinkedList{
		key:   k,
		value: v,
		prev:  nil,
		next:  nil,
	}
}

func Constructor(capacity int) LRUCache {
	lc := LRUCache{
		cap:   capacity,
		size:  0,
		cache: make(map[int]*DlinkedList),
		head:  InitDLinkedList(0, 0),
		tail:  InitDLinkedList(0, 0),
	}

	lc.head.next = lc.tail
	lc.tail.prev = lc.head

	return lc
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}

	this.moveToHead(node)

	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if !ok {
		node = InitDLinkedList(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.cap {
			tail := this.removeTail()
			delete(this.cache, tail.key)
			this.size--
		}

		return
	}

	node.value = value
	this.moveToHead(node)
}

func (this *LRUCache) moveToHead(node *DlinkedList) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) addToHead(node *DlinkedList) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DlinkedList) {
	node.next.prev = node.prev
	node.prev.next = node.next
}

func (this *LRUCache) removeTail() *DlinkedList {
	node := this.tail.prev
	this.removeNode(node)

	return node
}
