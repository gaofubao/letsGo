package linkedlist

import (
	"fmt"
	"testing"
)

func TestList_Clear(t *testing.T) {
	l := List{}
	l.Clear()
	l.Traverse()

	l.Append(1)
	l.Append(2)
	l.Clear()
	l.Traverse()
}

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

func TestList_Insert(t *testing.T) {
	l := List{}
	l.Insert(-1, 1)
	l.Insert(0, 2)
	l.Insert(3, 4)
	l.Insert(2, 3)
	l.Traverse()
}

func TestList_Get(t *testing.T) {
	l := List{}
	data1, ok1 := l.Get(1)
	fmt.Println(data1, ok1)

	l.Add(1)
	l.Add(2)
	data2, _ := l.Get(1)
	fmt.Println(data2)

	l.Add(3)
	data3, _ := l.Get(1)
	fmt.Println(data3)
}

func TestList_Delete(t *testing.T) {
	l := List{}
	l.Delete(0)

	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Delete(0)
	l.Delete(2)
	l.Traverse()
}

func TestCreate(t *testing.T) {
	l := Create(0)
	l.Traverse()

	l1 := Create(5)
	l1.Traverse()
}



