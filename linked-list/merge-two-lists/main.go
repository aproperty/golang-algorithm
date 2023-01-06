package main

// 将两个升序链表合并为一个新的 升序 链表并返回。
// 新链表 是通过拼接给定的两个链表 的所有节点 组成的。

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 取两个有序链表中较小的那一个数，链接到新链表上，
// 直到其中一个链表为空。再把另一个链表的剩余部分，链接到新链表上。
func mergeTwoListsMethodOne(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var l3 = &ListNode{}
	var current = l3

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	if l1 == nil {
		current.Next = l2
	} else {
		current.Next = l1
	}

	return l3.Next
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	// 递归出口

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// 递归逻辑
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}
