package list

import (
	"testing"
)

func TestAppend(t *testing.T) {
	list := Node{nil, 1}
	list.Append(2)
	list.Append(3)

	var actual = list.Last().value
	var expected = 3

	if actual != expected {
		t.Fatalf("Expected last value of list to be %d, got %d", expected, actual)
	}
}

func TestLastNode(t *testing.T) {
	third := Node{nil, 3}
	second := Node{&third, 2}
	first := Node{&second, 1}

	var actual *Node = first.Last()
	var expected *Node = &third

	if actual != expected {
		t.Fatalf("Expected last element of list to be %+v, got %+v", *expected, *actual)
	}
}

func TestReverse(t *testing.T) {
	list := &Node{nil, 1}
	list.Append(2)
	list.Append(3)

	list.Debug()

	list = list.Reverse()

	list.Debug()

	// TODO: PROPER TEST
}
