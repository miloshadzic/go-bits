package lfu_test

import (
	"fmt"
	"gobits/lfu"
	"testing"
)

func TestLFU(t *testing.T) {
	cache := lfu.Constructor(2)

	cache.Put(2, 1)
	cache.Put(2, 2)
	cache.Put(1, 1)
	cache.Put(4, 1)

	cache.DebugPrint()
	test(t, &cache, 2, 2)
}

func TestIteration(t *testing.T) {
	cache := lfu.Constructor(2)
	cache.Put(1, 1)
	cache.DebugPrint()
	cache.Put(2, 1)
	cache.DebugPrint()
	cache.Put(2, 1)
	cache.DebugPrint()
	cache.Put(3, 1)
	cache.DebugPrint()
	test(t, &cache, 2, 1)
}

func TestLFULong(t *testing.T) {
	commandInputs := []string{
		"LFUCache",
		"put",
		"put",
		"put",
		"put",
		"put",
		"get",
		"put",
		"get",
		"get",
		"put",
		"get",
		"put",
		"put",
		"put",
		"get",
		"put",
		"get",
		"get",
		"get",
		"get",
		"put",
		"put",
		"get",
		"get",
		"get",
		"put",
		"put",
		"get",
		"put",
		"get",
		"put",
		"get",
		"get",
		"get",
		"put",
		"put",
		"put",
		"get",
		"put",
		"get",
		"get",
		"put",
		"put",
		"get",
		"put",
		"put",
		"put",
		"put",
		"get",
		"put",
		"put",
		"get",
		"put",
		"put",
		"get",
		"put",
		"put",
		"put",
		"put",
		"put",
		"get",
		"put",
		"put",
		"get",
		"put",
		"get",
		"get",
		"get",
		"put",
		"get",
		"get",
		"put",
		"put",
		"put",
		"put",
		"get",
		"put",
		"put",
		"put",
		"put",
		"get",
		"get",
		"get",
		"put",
		"put",
		"put",
		"get",
		"put",
		"put",
		"put",
		"get",
		"put",
		"put",
		"put",
		"get",
		"get",
		"get",
		"put",
		"put",
		"put",
		"put",
		"get",
		"put",
		"put",
		"put",
		"put",
		"put",
		"put",
		"put",
	}
	valueInputs := [][]int{
		{10},
		{10, 13},
		{3, 17},
		{6, 11},
		{10, 5},
		{9, 10},
		{13},
		{2, 19},
		{2},
		{3},
		{5, 25},
		{8},
		{9, 22},
		{5, 5},
		{1, 30},
		{11},
		{9, 12},
		{7},
		{5},
		{8},
		{9},
		{4, 30},
		{9, 3},
		{9},
		{10},
		{10},
		{6, 14},
		{3, 1},
		{3},
		{10, 11},
		{8},
		{2, 14},
		{1},
		{5},
		{4},
		{11, 4},
		{12, 24},
		{5, 18},
		{13},
		{7, 23},
		{8},
		{12},
		{3, 27},
		{2, 12},
		{5},
		{2, 9},
		{13, 4},
		{8, 18},
		{1, 7},
		{6},
		{9, 29},
		{8, 21},
		{5},
		{6, 30},
		{1, 12},
		{10},
		{4, 15},
		{7, 22},
		{11, 26},
		{8, 17},
		{9, 29},
		{5},
		{3, 4},
		{11, 30},
		{12},
		{4, 29},
		{3},
		{9},
		{6},
		{3, 4},
		{1},
		{10},
		{3, 29},
		{10, 28},
		{1, 20},
		{11, 13},
		{3},
		{3, 12},
		{3, 8},
		{10, 9},
		{3, 26},
		{8},
		{7},
		{5},
		{13, 17},
		{2, 27},
		{11, 15},
		{12},
		{9, 19},
		{2, 15},
		{3, 16},
		{1},
		{12, 17},
		{9, 1},
		{6, 19},
		{4},
		{5},
		{5},
		{8, 1},
		{11, 7},
		{5, 2},
		{9, 28},
		{1},
		{2, 2},
		{7, 4},
		{4, 22},
		{7, 24},
		{9, 26},
		{13, 28},
		{11, 26},
	}

	expectedOutputs := []int{
		-99,
		-99,
		-99,
		-99,
		-99,
		-99,
		-1,
		-99,
		19,
		17,
		-99,
		-1,
		-99,
		-99,
		-99,
		-1,
		-99,
		-1,
		5,
		-1,
		12,
		-99,
		-99,
		3,
		5,
		5,
		-99,
		-99,
		1,
		-99,
		-1,
		-99,
		30,
		5,
		30,
		-99,
		-99,
		-99,
		-1,
		-99,
		-1,
		24,
		-99,
		-99,
		18,
		-99,
		-99,
		-99,
		-99,
		14,
		-99,
		-99,
		18,
		-99,
		-99,
		11,
		-99,
		-99,
		-99,
		-99,
		-99,
		18,
		-99,
		-99,
		-1,
		-99,
		4,
		29,
		30,
		-99,
		12,
		11,
		-99,
		-99,
		-99,
		-99,
		29,
		-99,
		-99,
		-99,
		-99,
		17,
		-1,
		18,
		-99,
		-99,
		-99,
		-1,
		-99,
		-99,
		-99,
		20,
		-99,
		-99,
		-99,
		29,
		18,
		18,
		-99,
		-99,
		-99,
		-99,
		20,
		-99,
		-99,
		-99,
		-99,
		-99,
		-99,
		-99,
	}

	var cache *lfu.LFUCache

	exampleRunner(t, cache, commandInputs, valueInputs, expectedOutputs)
}

func TestLC18(t *testing.T) {
	var cache *lfu.LFUCache

	commands := []string{
		"LFUCache",
		"put",
		"put",
		"put",
		"put",
		"get",
		"get",
		"get",
		"get",
		"put",
		"get",
		"get",
		"get",
		"get",
		"get",
	}

	values := [][]int{{3}, {1, 1}, {2, 2}, {3, 3}, {4, 4}, {4}, {3}, {2}, {1}, {5, 5}, {1}, {2}, {3}, {4}, {5}}
	outputs := []int{-1, -1, -1, -1, -1, 4, 3, 2, -1, -1, -1, 2, 3, -1, 5}

	exampleRunner(t, cache, commands, values, outputs)
}

func exampleRunner(t *testing.T, cache *lfu.LFUCache, commands []string, values [][]int, outputs []int) {
	for i, command := range commands {
		switch command {
		case "LFUCache":
			c := lfu.Constructor(values[i][0])
			cache = &c
		case "put":
			cache.Put(values[i][0], values[i][1])
		case "get":
			test(t, cache, values[i][0], outputs[i])
		}
		fmt.Println(command, values[i])
		cache.DebugPrint()
	}
}

func test(t *testing.T, cache *lfu.LFUCache, pos int, expected int) {
	actual := cache.Get(pos)

	if actual != expected {
		t.Fatal("Expected val at", pos, "to be", expected, "got", actual)
	}
}
