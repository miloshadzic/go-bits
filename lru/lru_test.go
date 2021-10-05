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
	cache.Put("second", "G")
	cache.Put("second", "b")
	cache.Put("third", "a")
	cache.Put("third", "E")

	if cache.Count() > 2 {
		t.Fatal("Cache should not have more than two elements")
	}

	if _, err := cache.Get("first"); err == nil {
		t.Fatal("First not evicted")
	}

	if val, _ := cache.Get("third"); val != "E" {
		t.Fatal("Third not E", val)
	}
}
