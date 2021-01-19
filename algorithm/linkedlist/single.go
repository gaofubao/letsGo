package linkedlist

import (
	"fmt"
	"math/rand"
)

/*
实现一个单链表，包括以下操作：
- Init(*L)
- isEmpty(L)
- clearList(*L)
- getElem(L, i, *e)
- locateElem(L, e)
- listInsert()
- listDelete()
- listLength()
*/

// 1 -> 2 -> 3 -> 4 -> nil

type object interface {}

// 定义节点
type Node struct {
	data object
	next *Node
}

// 定义链表
type LinkedList struct {
	head *Node
}

// 判断链表是否为空
func (list *LinkedList) IsEmpty() bool {
	if list.head == nil {
		return true
	}
	return false
}

// 获取链表的长度
func (list *LinkedList) GetLength() uint {
	if list.IsEmpty() {
		return 0
	} else {
		var length uint = 0
		for cur := list.head; cur.next != nil; cur = cur.next {
			length++
		}
		return length
	}
}

// 遍历链表节点
func (list *LinkedList) Traverse()  {
	if list.head == nil {
		fmt.Println("No element")
	} else {
		for cur := list.head; cur != nil; cur = cur.next {
			fmt.Println(cur.data)
		}

	}
}

// 在链表尾添加节点
func (list *LinkedList) Append(node *Node)  {
	if list.IsEmpty() {
		list.head = node
	} else {
		for cur := list.head; cur != nil; cur = cur.next {
			if cur.next == nil {
				cur.next = node
				break
			}
		}
	}
}

// 在链表头添加节点
func (list *LinkedList) Add(node *Node)  {
	if list.IsEmpty() {
		list.head = node
	} else {
		node.next = list.head
		list.head = node
	}
}

// 在链表中插入一个节点
func (list *LinkedList) Insert(i uint, node *Node)  {
	length := list.GetLength()
	if i == 0 || i >= length {
		list.Append(node)
	} else {
		var j uint = 0
		for cur := list.head; cur != nil; cur = cur.next {
			j++
			if i == j {
				node.next = cur.next
				cur.next = node
				break
			}
		}
	}
}

// 获取指定位置的元素

// 查找指定元素

// 删除指定位置的元素

// 删除链表

// 创建一个指定长度的链表
func Create(l uint) *LinkedList {
	list := LinkedList{nil}
	var i uint = 0

	for i = 0; i < l; i++ {
		data := rand.Int()
		node := Node{data: data, next: nil}
		list.Append(&node)
	}
	return &list
}
