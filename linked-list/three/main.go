package main

// Definition for singly-linked list.
type Node struct {
	Val  int
	Next *Node
}

// 思路是快慢指针，快指针比慢指针的步长快一步，那么快指针跑到链表末尾的时候，慢指针此时就位于链表的中间位置。
func middleNode(head *Node) *Node {

	slow := head
	fast := head

	for fast != nil {

		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			fast = fast.Next
			break
		}
		slow = slow.Next

	}

	return slow
}
