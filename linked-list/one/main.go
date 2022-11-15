package main

type Node struct {
	data int
	next *Node
}

// 给定一个单向链表，求出该链表中倒数第 m 个元素。
func getItem(head *Node, m int) *Node {

	if head == nil {
		return nil
	}

	if m == 0 { // 如果 m=0，则链表的最后一个元素就是所求的元素
		m = 1
	}

	pointOne := head
	for i := 0; i < m-1; i++ { // 第一个指针 从 链表 头指针 开始遍历，向前走 m-1 步，第二个指针不动。
		if pointOne.next != nil {
			pointOne = pointOne.next
		} else {
			return nil
		}
	}

	// 从第 m 步 开始，第二个指针 也开始 从 链表 头指针 开始遍历
	// 由于 两个指针 的距离 保持 m-1，
	// 当第一个指针 到达链表的 尾节点时，第二个指针 指向的 正好是 倒数 第 m 个节点
	pointTwo := head
	for pointOne.next != nil {
		pointOne = pointOne.next
		pointTwo = pointTwo.next
	}

	return pointTwo
}
