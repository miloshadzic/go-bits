package lru

import "container/list"

type Entry struct {
	key string
	val string
}

type LRUCache struct {
	cap  int
	data map[string]*list.Element
	list *list.List
}

func Init(capacity int) *LRUCache {
	data := make(map[string]*list.Element, capacity)
	list := list.New()

	return &LRUCache{capacity, data, list}
}

func (self *LRUCache) Put(key string, val string) {
	// Is there anything already under that key in the map?
	e := self.data[key]

	if e == nil {
		if len(self.data) == self.cap {
			last := self.list.Back()

			delete(self.data, last.Value.(Entry).key)

			self.list.Remove(last)
		}

		self.data[key] = self.list.PushFront(Entry{key, val})
	} else {
		e.Value = val
		self.list.MoveToFront(e)
	}
}

func (self *LRUCache) Get(key string) string {
	e := self.data[key]
	if e != nil {
		self.list.MoveToFront(e)

		return e.Value.(string)
	} else {
		return ""
	}
}

func (self *LRUCache) Count() int {
	return len(self.data)
}
