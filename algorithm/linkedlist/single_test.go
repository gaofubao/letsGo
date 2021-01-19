package linkedlist

import (
	"fmt"
	"testing"
)


func TestLinkedList_IsEmpty(t *testing.T)  {
	l0 := LinkedList{nil}
	fmt.Println(l0.IsEmpty())

	node1 := Node{data: 1, next: nil}
	l1 := LinkedList{&node1}
	fmt.Println(l1.IsEmpty())
}

func TestLinkedList_Traverse(t *testing.T) {
	l0 := LinkedList{nil}
	l0.Traverse()

	node1 := Node{data: 1, next: nil}
	l1 := LinkedList{&node1}
	l1.Traverse()
}

func TestLinkedList_Append(t *testing.T) {
	node1 := Node{data: 1, next: nil}
	l0 := LinkedList{nil}
	l0.Append(&node1)
	l0.Traverse()

	node2 := Node{data: 2, next: nil}
	l0.Append(&node2)
	l0.Traverse()
}

func TestLinkedList_Add(t *testing.T) {
	node1 := Node{data: 1, next: nil}
	l0 := LinkedList{nil}
	l0.Add(&node1)
	l0.Traverse()
}

func TestLinkedList_Insert(t *testing.T) {
	l0 := Create(5)
	fmt.Println(l0.GetLength())

	node0 := Node{data: 10, next: nil}
	l0.Insert(12, &node0)
	l0.Traverse()
}

func TestA(t *testing.T) {
	for i := 0; i != 3; i++ {
		fmt.Println(i)
	}
}




