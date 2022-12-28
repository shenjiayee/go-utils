package collections

import "sync/atomic"

// AtomicValue 对`atomic.Value`的一个泛型封装
type AtomicValue[T any] struct {
	value atomic.Value
}

func (a *AtomicValue[T]) Store(val T) {
	a.value.Store(val)
}

func (a *AtomicValue[T]) Load() T {
	return a.value.Load().(T)
}

func (a *AtomicValue[T]) Swap(new T) T {
	return a.value.Swap(new).(T)
}

func (a *AtomicValue[T]) CompareAndSwap(old, new T) bool {
	return a.value.CompareAndSwap(old, new)
}
