package collections

import "sync"

// RWMap 读写加锁的Map，适用于大量写的场景
type RWMap[K comparable, V any] struct {
	mu    sync.RWMutex
	store map[K]V
}

func NewRWMap[K comparable, V any]() *RWMap[K, V] {
	return &RWMap[K, V]{
		store: make(map[K]V),
	}
}

func (m *RWMap[K, V]) Get(key K, default_ V) (V, bool) {
	defer m.mu.RUnlock()
	m.mu.RLock()
	if val, ok := m.store[key]; ok {
		return val, true
	}
	return default_, false
}

func (m *RWMap[K, V]) Set(key K, value V) {
	defer m.mu.Unlock()
	m.mu.Lock()
	m.store[key] = value
}

func (m *RWMap[K, V]) Pop(key K, default_ V) (V, bool) {
	defer m.mu.Unlock()
	m.mu.Lock()
	if val, ok := m.store[key]; ok {
		delete(m.store, key)
		return val, true
	}
	return default_, false
}
