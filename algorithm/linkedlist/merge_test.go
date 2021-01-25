package linkedlist

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val int
	Next *ListNode
}

//1->2->4, 1->3->4
//1->1->2->3->4->4
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	prehead := new(ListNode)
	prehead.Val = -1

	pre := prehead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			pre.Next = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
		}
		pre = pre.Next
	}
	if l1 != nil {
		pre.Next = l1
	} else {
		pre.Next = l2
	}
	return prehead.Next
}

func TestMergeTwoLists(t *testing.T)  {
	a1 := new(ListNode)
	a1.Val = 1
	a2 := new(ListNode)
	a2.Val = 2
	a3 := new(ListNode)
	a3.Val = 4

	a1.Next = a2
	a2.Next = a3
	a3.Next = nil

	b1 := new(ListNode)
	b1.Val = 1
	b2 := new(ListNode)
	b2.Val = 3
	b3 := new(ListNode)
	b3.Val = 4

	b1.Next = b2
	b2.Next = b3
	a3.Next = nil

	head := mergeTwoLists(a1, b1)
	fmt.Println(head, head.Next, head.Next.Next, head.Next.Next.Next)
}
