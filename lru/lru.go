package lru

import "errors"

type LRUCache struct {
	cap  int
	data map[string]*entry
	head *entry
	last *entry
}

type entry struct {
	next *entry
	prev *entry
	key  string
	val  string
}

func NewLRUCache(cap int) *LRUCache {
	data := make(map[string]*entry, cap)

	return &LRUCache{cap, data, nil, nil}
}

func (this *LRUCache) Get(key string) (string, error) {
	e := this.data[key]
	if e != nil {
		this.moveToFront(e)
		return e.val, nil
	} else {
		return "", errors.New("Key not in cache.")
	}
}

func (this *LRUCache) Put(key string, val string) {
	_, err := this.Get(key)
	// Key was not in map already
	if err != nil {
		e := entry{nil, nil, key, val}
		this.prepend(&e)

		if len(this.data) == this.cap {
			this.truncate()
		}

		this.data[key] = &e
	} else {
		this.data[key].val = val
	}
}

func (this *LRUCache) Count() int {
	return len(this.data)
}

func (this *LRUCache) moveToFront(e *entry) {
	if this.head == e {
		return
	} else {
		e.prev.next = e.next
	}

	if this.last != e {
		e.next.prev = e.prev

		this.last = e.prev
	}

	e.prev = nil

	this.prepend(e)
}

func (this *LRUCache) prepend(e *entry) {
	e.next = this.head

	if this.head != nil {
		this.head.prev = e
	} else {
		this.last = e
	}

	for this.last.next != nil {
		this.last = this.last.next
	}

	this.head = e
}

func (this *LRUCache) truncate() *entry {
	e := this.last

	this.last = e.prev
	e.prev = nil
	this.last.next = nil

	delete(this.data, e.key)

	return e
}
