package list

import "fmt"

type Node struct {
	next  *Node
	value int
}

func (self *Node) Last() *Node {
	var temp *Node = self

	for temp.next != nil {
		temp = temp.next
	}

	return temp
}

func (self *Node) Append(value int) error {
	self.Last().next = &Node{nil, value}

	return nil
}

func (self *Node) Reverse() *Node {
	var temp *Node = self
	var prev *Node
	var next *Node

	for temp != nil {
		next = temp.next
		temp.next = prev
		prev = temp
		temp = next
	}

	return prev
}

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
