package linkedlist

import (
	"errors"
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

type object interface{}

// 定义节点
type Node struct {
	data object
	next *Node
}

// 定义单链表
type List struct {
	head *Node
}

// 判断链表是否为空
func (l *List) IsEmpty() bool {
	if l.head == nil {
		return true
	}
	return false
}

// 获取链表的长度
func (l *List) Length() int {
	length := 0
	cur := l.head

	for cur != nil {
		length++
		cur = cur.next
	}
	return length
}

// 遍历链表节点
func (l *List) Traverse() {
	if l.head == nil {
		fmt.Println("No element")
	} else {
		for cur := l.head; cur != nil; cur = cur.next {
			fmt.Println(cur.data)
		}

	}
}

// 在链表头添加节点
func (l *List) Add(data object) {
	node := &Node{data: data}
	node.next = l.head
	l.head = node
}

// 在链表尾添加节点
func (l *List) Append(data object) {
	node := &Node{data: data}
	if l.IsEmpty() {
		l.head = node
	} else {
		cur := l.head
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = node
	}
}



// 在链表中插入一个节点
func (list *List) Insert(i int, node *Node) {
	length := list.Length()
	if i <= 0 || i >= length {
		list.Append(node)
	} else {
		var j = 0
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
func (list *List) GetElement(i int) (object, error) {
	if i >= list.Length() {
		return 0, errors.New("超出链表长度")
	}

	var j = 0
	var data interface{}
	for cur := list.head; cur != nil; cur = cur.next {
		if i == j {
			data = cur.data
			break
		}
		j++
	}
	return data, nil
}

// 查找指定元素
func (list *List) Locate(e *Node) (i uint, err error) {
	return
}

// 删除指定位置的元素
func (list *List) Delete() {

}

// 创建一个指定长度的链表
func Create(l uint) *List {
	list := List{nil}
	var i uint = 0

	for i = 0; i < l; i++ {
		data := rand.Int()
		node := Node{data: data, next: nil}
		list.Append(&node)
	}
	return &list
}
