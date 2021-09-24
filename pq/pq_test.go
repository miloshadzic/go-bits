package pq_test

import (
	"gobits/pq"
	"testing"
)

func TestPeek(t *testing.T) {
	pq := pq.Init()

	_, err := pq.Peek()

	if err == nil {
		t.Fatal("Expected Peek to return an error when called on an empty Queue")
	}

	pq.Insert(5)

	val, err := pq.Peek()

	if val != 5 {
		t.Fatalf("Expected Peek to return 5, got %d", val)
	}

	pq.Insert(1)
	pq.Insert(3)

	val, err = pq.Peek()

	if val != 1 {
		t.Fatalf("Expected Peek to return 1, got %d", val)
	}
}

func TestTake(t *testing.T) {
	pq := pq.Init()

	_, err := pq.Take()

	if err == nil {
		t.Fatal("Expected Take to return an error when called on an empty Queue")
	}

	pq.Insert(5)

	val, err := pq.Take()

	if val != 5 {
		t.Fatalf("Expected Take to return 5, got %d", val)
	}

	pq.Insert(1)
	pq.Insert(3)

	val, err = pq.Take()

	if val != 1 {
		t.Fatalf("Expected Take to return 1, got %d", val)
	}

	val, err = pq.Take()

	if val != 3 {
		t.Fatalf("Expected Take to return 3, got %d", val)
	}

	_, err = pq.Take()

	if err == nil {
		t.Fatal("Expected Take to return an error when called on an empty Queue")
	}

}
