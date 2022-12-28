package slices

// StreamWrapper 流式处理
type StreamWrapper[T any] struct {
	slice []T
}

func NewStream[T any](slice []T) *StreamWrapper[T] {
	return &StreamWrapper[T]{slice: slice}
}

// Value 值
func (s *StreamWrapper[T]) Value() []T {
	return s.slice
}

// Length 长度
func (s *StreamWrapper[T]) Length() int {
	return len(s.slice)
}

// Copy 复制
func (s *StreamWrapper[T]) Copy() *StreamWrapper[T] {
	newSlice := Copy(s.slice)
	return NewStream(newSlice)
}

// Any 任一
func (s *StreamWrapper[T]) Any(f func(T) bool) bool {
	return Any(s.slice, f)
}

// All 所有
func (s *StreamWrapper[T]) All(f func(T) bool) bool {
	return All(s.slice, f)
}

// ForEach 依次执行
func (s *StreamWrapper[T]) ForEach(f func(T, int)) *StreamWrapper[T] {
	ForEach(s.slice, f)
	return s
}

// Find 找到第一个
func (s *StreamWrapper[T]) Find(f func(T, int) bool) (T, bool) {
	return Find(s.slice, f)
}

// Filter 过滤
func (s *StreamWrapper[T]) Filter(f func(T, int) bool) *StreamWrapper[T] {
	newSlice := Filter(s.slice, f)
	copy(s.slice, newSlice)
	s.slice = s.slice[:len(newSlice)]
	return s
}

// Map
// 基于`Go-1.19.3`，因暂不支持泛型方法，`Map`方法只支持同类型
func (s *StreamWrapper[T]) Map(f func(T, int) T) *StreamWrapper[T] {
	newSlice := Map(s.slice, f)
	copy(s.slice, newSlice)
	return s
}
