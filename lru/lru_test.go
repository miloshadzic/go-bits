package lru_test

import (
	"gobits/lru"
	"testing"
)

func TestLRU(t *testing.T) {
	cache := lru.Init(2)

	cache.Put("first", "A")
	cache.Put("second", "B")
	cache.Put("third", "C")

	if cache.Count() > 2 {
		t.Fatal("Cache should not have more than two elements")
	}

	if cache.Get("first") == "A" {
		t.Fatal("First not evicted")
	}
}
