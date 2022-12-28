package collections

import "sync"

type SyncMap[K comparable, V any] struct {
	m sync.Map
}

func (s *SyncMap[K, V]) Store(key K, value V) {
	s.m.Store(key, value)
}

func (s *SyncMap[K, V]) Load(key K) (V, bool) {
	if v, ok := s.m.Load(key); ok {
		return v.(V), ok
	}
	return *new(V), false
}

func (s *SyncMap[K, V]) LoadAndDelete(key K) (V, bool) {
	if v, ok := s.m.LoadAndDelete(key); ok {
		return v.(V), ok
	}
	return *new(V), false
}

func (s *SyncMap[K, V]) Delete(key K) {
	s.m.Delete(key)
}

func (s *SyncMap[K, V]) LoadOrStore(key K, value V) (V, bool) {
	if actual, loaded := s.m.LoadOrStore(key, value); loaded {
		return actual.(V), loaded
	}
	return *new(V), false
}

func (s *SyncMap[K, V]) Range(f func(key K, value V) bool) {
	rangeFn := func(k, v any) bool {
		return f(k.(K), v.(V))
	}
	s.m.Range(rangeFn)
}
