package main

type LRUCache struct {
	capacity int
	len      int

	hashMap map[int]*Node

	head *Node
	tail *Node
}

type Node struct {
	Key int
	Val int

	Prev *Node
	Next *Node
}

func Constructor(capacity int) LRUCache {

	m := make(map[int]*Node)

	lru := LRUCache{
		capacity: capacity,
		hashMap:  m,
		head:     &Node{},
		tail:     &Node{},
	}

	lru.head.Next = lru.tail
	lru.tail.Prev = lru.head

	return lru
}

func (this *LRUCache) Get(key int) int {

	if v, ok := this.hashMap[key]; ok {

		v.Prev.Next = v.Next
		v.Next.Prev = v.Prev

		n := this.head.Next

		this.head.Next = v
		v.Prev = this.head

		n.Prev = v
		v.Next = n

		return v.Val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {

	if v, ok := this.hashMap[key]; ok {

		v.Prev.Next = v.Next
		v.Next.Prev = v.Prev

		n := this.head.Next

		this.head.Next = v
		v.Prev = this.head

		n.Prev = v
		v.Next = n

		v.Val = value

		return
	}

	if this.len < this.capacity {

		this.len++

		node := &Node{
			Key: key,
			Val: value,
		}
		this.hashMap[key] = node

		n := this.head.Next

		this.head.Next = node
		node.Prev = this.head

		node.Next = n
		n.Prev = node

	} else {

		t := this.tail.Prev

		this.tail.Prev.Prev.Next = this.tail
		this.tail.Prev = this.tail.Prev.Prev

		t.Val = value
		delete(this.hashMap, t.Key)
		t.Key = key

		this.hashMap[key] = t

		hn := this.head.Next

		this.head.Next = t
		t.Prev = this.head

		t.Next = hn
		hn.Prev = t
	}

	return
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);

 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
