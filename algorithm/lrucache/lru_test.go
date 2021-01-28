package lrucache

import (
	"fmt"
	"testing"
)

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

type LRUCache struct {
	size int
	capacity int
	cache map[int]*DLinkedNode
	head, tail *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key: key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache: map[int]*DLinkedNode{},
		head: initDLinkedNode(0, 0),
		tail: initDLinkedNode(0, 0),
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int)  {
	if _, ok := this.cache[key]; !ok {
		node := initDLinkedNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

// 在头部新增节点
func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

// 移除节点
func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 添加节点到头部
func (this *LRUCache) moveToHead(node *DLinkedNode)  {
	this.removeNode(node)
	this.addToHead(node)
}

// 删除尾部的节点
func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

func TestLRUCache(t *testing.T)  {
	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))   // 返回  1
	cache.Put(3, 3)   		// 该操作会使得密钥 2 作废
	fmt.Println(cache.Get(2))   // 返回 -1 (未找到)
	cache.Put(4, 4)   		// 该操作会使得密钥 1 作废
	fmt.Println(cache.Get(1))	// 返回 -1 (未找到)
	fmt.Println(cache.Get(3))   // 返回  3
	fmt.Println(cache.Get(4))   // 返回  4
}

