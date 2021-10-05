package lru_test

import (
	"gobits/lru"
	"testing"
)

func TestLRU(t *testing.T) {
	cache := lru.NewLRUCache(2)

	cache.Put("first", "A")
	cache.Put("second", "B")
	cache.Put("third", "C")
	cache.Put("second", "B")
	cache.Put("second", "B")
	cache.Put("third", "C")
	cache.Put("third", "C")

	if cache.Count() > 2 {
		t.Fatal("Cache should not have more than two elements")
	}

	if _, err := cache.Get("first"); err == nil {
		t.Fatal("First not evicted")
	}
}
