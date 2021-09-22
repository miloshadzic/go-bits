package list

import (
	"testing"
)

func TestAppend(t *testing.T) {
	list := Node{nil, 1}
	list.Append(2)
	list.Append(3)

	actual := list.Last().value
	expected := 3

	if actual != expected {
		t.Fatalf("Expected last value of list to be %d, got %d", expected, actual)
	}
}

func TestLastNode(t *testing.T) {
	list := Node{nil, 1}
	list.Append(2)

	expected := list.Append(3)

	actual := list.Last()

	if actual != expected {
		t.Fatalf("Expected last element of list to be %+v, got %+v", *expected, *actual)
	}
}

func TestEqual(t *testing.T) {
	node := &Node{nil, 99}

	list := &Node{nil, 12}
	list.Append(2)
	list.Append(3)

	if node.Equal(list) {
		t.Fatalf("Single element list is equal to a longer list")
	}

	if list.Equal(node) {
		t.Fatalf("Longer list is equal to a single element list.")
	}

	if !list.Equal(list) {
		t.Fatalf("List should be equal to itself.")
	}

	if !node.Equal(node) {
		t.Fatalf("List should be equal to itself.")
	}

	if node.Equal(nil) {
		t.Fatalf("List should not be equal to nil")
	}
}

func TestReverse(t *testing.T) {
	list := &Node{nil, 1}
	list.Append(2)
	list.Append(3)

	actual := list.Reverse()

	expected := &Node{nil, 3}
	expected.Append(2)
	expected.Append(1)

	if !actual.Equal(expected) {
		t.Fatalf("List not reversed correctly")
	}
}
