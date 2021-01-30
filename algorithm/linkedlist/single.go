package linkedlist

import (
	"errors"
	"fmt"
	"math/rand"
)

/*
实现一个单链表，包括以下操作：
- Clear()
- IsEmpty()
- Length()
- Traverse()
- Add()
- Append()
- Insert()
- Get()
- Locate()
- Delete()
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

// 清空链表
func (l *List) Clear() {
	l.head = nil
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
func (l *List) Insert(index int, data object) {
	length := l.Length()
	if index <= 0 || index >= length {
		l.Append(data)
	} else {
		pos := 0
		cur := l.head
		for pos < index-1 {
			cur = cur.next
			pos ++
		}
		node := &Node{data: data}
		node.next = cur.next
		cur.next = node
	}
}

// 获取指定位置的元素
func (l *List) Get(index int) (data object, err error) {
	if index < 0 || index >= l.Length() {
		return 0, errors.New("不在链表内")
	}
	pos := 0
	cur := l.head
	for pos < index {
		cur = cur.next
		pos++
	}
	return cur.data, nil
}

// 查找指定元素，返回第一个相等的值
func (l *List) Locate(data object) (index int, err error) {
	// TODO 类型比较
	return
}

// 删除指定位置的元素
func (l *List) Delete(index int) {
	if l.IsEmpty() {
		// TODO 异常处理
		fmt.Println("空链表")
		return
	}

	cur := l.head
	if index <= 0 {
		l.head = cur.next
	} else if index >= l.Length() {
		// TODO 异常处理
		fmt.Println("超出链表长度")
		return
	} else {
		pos := 0
		for pos < index - 1 {
			cur = cur.next
			pos++
		}
		cur.next = cur.next.next
	}
}

// 创建一个指定长度的链表
func Create(length int) *List {
	if length <= 0 {
		return &List{}
	}

	list := List{}
	index := 0
	for index = 0; index < length; index++ {
		data := rand.Int()
		node := &Node{data: data}
		list.Append(node.data)
	}
	return &list
}
