/*

So a LFU cache evicts the one... least frequently used. Seems like the
most straightforward way to go about this is to track the LFU item in a
priority queue.

https://leetcode.com/problems/lfu-cache

*/
package lfu

import "fmt"

type LFUCache struct {
	capacity int
	data     map[int]*node
	pq       *list
}

func Constructor(capacity int) LFUCache {
	data := make(map[int]*node, capacity)
	list := &list{0, nil}

	return LFUCache{capacity, data, list}
}

func (this *LFUCache) Get(key int) int {
	e, exists := this.data[key]
	if !exists {
		return -1
	}

	e.count++
	this.pq.bubbleDown(e)

	return e.val
}

func (this *LFUCache) Put(key, val int) {
	if this.capacity == 0 {
		return
	}

	e, exists := this.data[key]

	if exists {
		e.val = val
		e.count++
	} else {
		e = &node{1, key, val, nil, nil}

		if this.pq.count == this.capacity {
			// Have to evict head
			old := this.pq.Pop()
			delete(this.data, old.key)
		}

		this.pq.Push(e)
		this.data[key] = this.pq.head

	}

	this.pq.bubbleDown(e)
}

// Given a node n, bubbleDown will move it forward until it encounters
// either the end or the next node which has a count greater than
// its count. This is so the priority queue can maintain it's strict
// ordering.
func (this *list) bubbleDown(n *node) {
	// If the n is nil, or there's no next node, we just need to return.
	if n == nil || n.next == nil {
		return
	}

	// If the passed in node is the head node, we need to store a reference
	// to it's next node so if we end up swapping, we can move the head pointer
	// to the new node.
	newHead := this.head.next
	swapHead := n == this.head

	// Here's the meat. We want to move forward until we encounter either nil or
	// a node that has a higher count than n.
	t := n

	for {
		if t.next == nil || t.count < t.next.count {
			break
		}

		t = t.next
	}

	// When we're at this boundary, we have two cases:
	//
	//   1. We're at the end, in which case we need to perform a head/tail
	//      swap
	//
	//   2. We've arrived up to a node which has a higher
	//      count. We need to insert n between the current node t, and that
	//      node.
	if n != t {
		// We need to "remove" n from its old position
		next := n.next
		prev := n.prev

		if n.prev != nil {
			n.prev.next = next
		}

		if n.next != nil {
			n.next.prev = prev
		}

		// Then insert into new position which is after node t
		n.next = t.next
		if n.next != nil {
			n.next.prev = n
		}

		t.next = n
		n.prev = t

		// Update the head pointer if there was a change
		if swapHead {
			this.head = newHead
			newHead.prev = nil
		}
	}
}

type list struct {
	count int
	head  *node
}

type node struct {
	count int
	key   int
	val   int
	prev  *node
	next  *node
}

func (this *list) Push(n *node) {
	if this.head != nil {
		old := this.head
		n.next = old
		old.prev = n
	}

	this.head = n
	this.count++
}

func (this *list) Pop() *node {
	if this.empty() {
		return nil
	}

	n := this.head

	this.head = n.next

	this.count--
	return n
}

func (this *list) empty() bool {
	return this.count == 0
}

func (this *LFUCache) DebugPrint() {
	fmt.Println("HEAD:", this.pq.head)
	fmt.Print("|")

	count := 0
	for e := this.pq.head; e != nil; e = e.next {
		count++
	}

	for e := this.pq.head; e != nil; e = e.next {
		fmt.Print(" [", e.key, ":", e.val, "-", e.count, "]")
	}

	fmt.Print(" | NODES:", count, " CAP:", this.capacity, " PQ.Count:", this.pq.count, "\n")
}
