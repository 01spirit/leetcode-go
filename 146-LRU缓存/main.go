package main

import "fmt"

type LRUCache struct {
	cap        int
	size       int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func InitDLinkedList(k, v int) *DLinkedNode {
	return &DLinkedNode{
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
		cache: make(map[int]*DLinkedNode),
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

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.next.prev = node.prev
	node.prev.next = node.next
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)

	return node
}

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
}
