package collections

import "github.com/shenjiayee/go-utils/slices"

// Set 集合
type Set[T comparable] map[T]struct{}

func NewSet[T comparable](values ...T) Set[T] {
	set := Set[T](make(map[T]struct{}))
	for _, val := range values {
		set.Add(val)
	}
	return set
}

func (s Set[T]) Add(val T) {
	s[val] = struct{}{}
}

func (s Set[T]) Pop(val T) {
	delete(s, val)
}

func (s Set[T]) Has(val T) bool {
	_, ok := s[val]
	return ok
}

func (s Set[T]) Length() int {
	return len(s)
}

func (s Set[T]) Values() []T {
	values := make([]T, s.Length())
	i := 0
	for val := range s {
		values[i] = val
		i++
	}
	return values
}

func (s Set[T]) Copy() Set[T] {
	return NewSet(s.Values()...)
}

// ComplexSet 复杂集合
// 解决非comparable不可用的问题
type ComplexSet[T any, K comparable] struct {
	// 取Key的函数
	// keyFn(x)相等的两个元素视为同一个元素
	// 后置优先
	keyFn func(item T) K
	store map[K]T
}

func NewComplexSet[T any, K comparable](key func(item T) K, values ...T) *ComplexSet[T, K] {
	set := &ComplexSet[T, K]{
		keyFn: key,
		store: make(map[K]T),
	}
	slices.ForEach(
		values,
		slices.W[T](set.Add),
	)
	return set
}

func (s *ComplexSet[T, K]) Add(val T) {
	key := s.keyFn(val)
	s.store[key] = val
}

// Pop 即使`val`不同，但是`keyFn(val)`相等就行
func (s *ComplexSet[T, K]) Pop(val T) {
	delete(s.store, s.keyFn(val))
}

func (s *ComplexSet[T, K]) Has(val T) bool {
	_, ok := s.store[s.keyFn(val)]
	return ok
}

func (s *ComplexSet[T, K]) Length() int {
	return len(s.store)
}

func (s *ComplexSet[T, K]) Values() []T {
	values := make([]T, s.Length())
	i := 0
	for _, val := range s.store {
		values[i] = val
		i++
	}
	return values
}

func (s *ComplexSet[T, K]) Copy() *ComplexSet[T, K] {
	set := &ComplexSet[T, K]{
		keyFn: s.keyFn,
		store: make(map[K]T),
	}
	for key, val := range s.store {
		set.store[key] = val
	}
	return set
}
