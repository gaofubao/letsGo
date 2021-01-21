package linkedlist

import (
	"fmt"
	"testing"
)

func TestList_IsEmpty(t *testing.T)  {
	l := List{}
	fmt.Println(l.IsEmpty())
}

func TestList_Length(t *testing.T) {
	l := List{}
	fmt.Println(l.Length())
}

func TestList_Traverse(t *testing.T) {
	l := List{}
	l.Traverse()
}

func TestList_Add(t *testing.T) {
	l := List{}
	l.Add(1)
	l.Traverse()
}

func TestList_Append(t *testing.T) {
	l := List{}
	l.Append(1)
	l.Traverse()

	l.Append(2)
	l.Traverse()
}


//func TestLinkedList_Insert(t *testing.T) {
//	l0 := Create(5)
//	fmt.Println(l0.GetLength())
//
//	node0 := Node{data: 10, next: nil}
//	l0.Insert(12, &node0)
//	l0.Traverse()
//}
//
//func TestLinkedList_GetElement(t *testing.T) {
//
//}
//
//func TestA(t *testing.T) {
//	for i := 0; i != 3; i++ {
//		fmt.Println(i)
//	}
//}




