package linkedlist

import (
	"fmt"
	"testing"
)

/*
单链表反转
1 -> 2 -> 3 -> 4 -> 5 -> nil
5 -> 4 -> 3 -> 2 -> 1 -> nil
*/

// 迭代法 双指针
func reverseList(head *Node) *Node  {
	var cur *Node
	pre := head
	for pre != nil {
		cur, pre, pre.next = pre, pre.next, cur
	}
	return cur
}

// 递归法
func reverseList2(head *Node) *Node {
	// 1. 递归终止条件
	if head == nil || head.next == nil {
		return head
	}
	// 2. 处理当前层逻辑
	// head = head.next
	// 3. 下钻
	newHead := reverseList2(head.next)
	//
	head.next.next = head
	head.next = nil
	return newHead
}


func TestReverseList(t *testing.T)  {
	l := List{}
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Append(5)
	l.Traverse()

	cur := reverseList2(l.head)
	fmt.Println(cur.data)
}


/*
链表中环路检测
*/
// 哈希表法
func detectCycle(head *Node) *Node {
	seen := map[*Node]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return head
		}
		seen[head] = struct{}{}
		head = head.next
	}
	return nil
}

// 双指针法
func detectCycle2(head *Node) *Node {
	slow, fast := head, head
	for fast != nil {
		slow = slow.next
		if fast.next == nil {
			return nil
		}
		fast = fast.next.next
		if fast == slow {
			p := head
			for p != slow {
				p = p.next
				slow = slow.next
			}
			return p
		}
	}
	return nil
}

/*
两个有序的链表合并
1 -> 2 -> 4
1 -> 3 -> 4
1 -> 1 -> 2 -> 3 -> 4 -> 4
*/

func mergeTwoLists(l1 *Node, l2 *Node) *Node {

	return nil
}

// 删除链表倒数第n个结点

// 求链表的中间结点


