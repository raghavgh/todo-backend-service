package lru

import (
	"sync"
	"todoapp/config"
	"todoapp/ds/linkedlist"
)

type LRU struct {
	items    map[string]*linkedlist.Node
	eviction *linkedlist.LinkedList
	limit    int
	mu       *sync.RWMutex
}

type entry struct {
	key   string
	value any
}

func (l *LRU) Get(key string) (any, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if node, ok := l.items[key]; ok {
		l.eviction.MoveToFront(node)
		return node.Val.(*entry).value, true
	}
	return nil, false
}

func (l *LRU) Put(key string, val any) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.eviction.Len() >= l.limit {
		delete(l.items, l.eviction.Tail.Val.(*entry).key)
		l.eviction.Remove(l.eviction.Tail)
	}
	l.eviction.PushFront(&entry{key: key, value: val})
	l.items[key] = l.eviction.Head
}

func (l *LRU) Remove(key string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if node, ok := l.items[key]; ok {
		delete(l.items, key)
		l.eviction.Remove(node)
	}
}

func (l *LRU) Len() int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.eviction.Len()
}

func (l *LRU) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.items = nil
	l.eviction = nil
	l.items = make(map[string]*linkedlist.Node)
	l.eviction = linkedlist.New()
}

func (l *LRU) Contains(key string) bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	_, ok := l.items[key]
	return ok
}

func NewLRU() *LRU {
	limit := config.Config.CacheConfig.Limit
	return &LRU{
		items:    make(map[string]*linkedlist.Node, limit),
		eviction: linkedlist.New(),
		limit:    limit,
		mu:       &sync.RWMutex{},
	}
}
