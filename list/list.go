package list

import "fmt"

type Node struct {
	next  *Node
	value int
}

// Returns a pointer to the last element of the list.
func (self *Node) Last() *Node {
	var temp *Node = self

	for temp.next != nil {
		temp = temp.next
	}

	return temp
}

// Append an element to the end of the list. Returns a pointer to the new
// element.
func (self *Node) Append(value int) *Node {
	node := &Node{nil, value}

	self.Last().next = node

	return node
}

// Prepend an element to the list and return a pointer to the new element.
func (self *Node) Push(value int) *Node {
	return &Node{self, value}
}

// Removes the first element. Returns the value of the removed element and the new head.
func (self *Node) Pop() (head *Node, value int) {
	head = self.next
	self.next = nil

	return head, self.value
}

// Reverses the list in place. Returns a pointer new head of the list.
func (self *Node) Reverse() *Node {
	var temp, prev, next *Node = self, nil, nil

	for temp != nil {
		next = temp.next
		temp.next = prev
		prev = temp
		temp = next
	}

	return prev
}

// Two lists are equal if all of their values are the same.
func (self *Node) Equal(other *Node) bool {
	if other == nil {
		return false
	}

	left, right := self, other

	for left != nil {
		if left.value == right.value {
			left = left.next
			right = right.next
		} else {
			return false
		}
	}

	return true
}

// Prints out the linked list on a new line.
// Format: [ a, b, c ]
func (self *Node) Debug() {
	var temp *Node = self

	fmt.Print("[ ")

	for temp != nil {
		fmt.Printf("%d", temp.value)

		if temp.next != nil {
			fmt.Print(", ")
		} else {
			fmt.Print(" ")
		}

		temp = temp.next
	}

	fmt.Print("]\n")
}
