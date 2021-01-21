package example

import (
	"container/list"
	"fmt"
	"testing"
)

func TestIterate(t *testing.T)  {
	l := list.New()

	fmt.Println(l)


	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
