package main

import "log"

type Node struct {
	data int
	next *Node
}

// func Revert(src *Node) (new *Node) {

// 	nodeMap := make(map[int]*Node)

// 	nextNode := src
// 	count := 0
// 	nodeMap[count] = nextNode

// 	for {
// 		if nextNode.next != nil {
// 			nextNode = nextNode.next
// 			count += 1
// 			nodeMap[count] = nextNode
// 		} else {
// 			break
// 		}
// 	}

// 	for n := count; n > 0; n-- {
// 		nodeMap[n].next = nodeMap[n-1]
// 	}
// 	nodeMap[0].next = nil

// 	return nodeMap[count]
// }

func Revert(src *Node) (new *Node) {

	newNode := &Node{0, nil}

	cursor := src
	for cursor != nil {

		q := cursor
		cursor = cursor.next

		q.next = newNode.next
		newNode.next = q
	}

	return newNode.next
}

func PrintLinkNode(node *Node) {
	if node == nil {
		return
	}
	nextNode := node
	log.Println("PrintLinkNode:")
	for {
		println(nextNode.data)
		if nextNode.next == nil {
			break
		}
		nextNode = nextNode.next
	}
}

func main() {
	first := &Node{data: 1}
	second := &Node{data: 2}
	third := &Node{data: 3}
	four := &Node{data: 4, next: nil}

	first.next = second
	second.next = third
	third.next = four

	PrintLinkNode(first)

	netHead := Revert(first)

	PrintLinkNode(netHead)
}
