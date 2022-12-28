package collections

import (
	"container/list"
	"errors"

	"go-utils/lock"
)

var KeyError = errors.New("keyFn error")

// List 列表
type List[T comparable] struct {
	mu lock.RWLocker
	li *list.List
}

func NewList[T comparable](values ...T) *List[T] {
	li := &List[T]{
		li: list.New(),
	}
	li.Extend(values)
	return li
}

func (li *List[T]) Length() (length int) {
	li.mu.RunR(
		func() {
			length = li.li.Len()
		},
	)
	return
}

func (li *List[T]) Empty() bool {
	return li.Length() == 0
}

func (li *List[T]) findElement(val T) *list.Element {
	for ele := li.li.Front(); ele != nil; ele = ele.Next() {
		if ele.Value.(T) == val {
			return ele
		}
	}
	return nil
}

func (li *List[T]) getElement(index int) *list.Element {
	var ele *list.Element
	if index >= 0 {
		if index >= li.li.Len() {
			return nil
		}
		ele = li.li.Front()
		for c := 0; c < index; c++ {
			ele = ele.Next()
		}
		return ele
	} else {
		if -index > li.li.Len() {
			return nil
		}
		ele = li.li.Back()
		for c := 1; c < -index; c++ {
			ele = ele.Prev()
		}
		return ele
	}
}

// Append 在后面加一个
func (li *List[T]) Append(val T) {
	li.mu.RunW(
		func() {
			li.li.PushBack(val)
		},
	)
}

// Extend 在后面加一个列表
func (li *List[T]) Extend(other []T) {
	li.mu.RunW(
		func() {
			for i := 0; i < len(other); i++ {
				li.li.PushBack(other[i])
			}
		},
	)
}

// Contains 删除个指定元素
func (li *List[T]) Contains(val T) bool {
	var ele *list.Element
	li.mu.RunR(
		func() {
			ele = li.findElement(val)
		},
	)
	return ele != nil
}

// Remove 删除一个指定元素
func (li *List[T]) Remove(val T) {
	li.mu.RunW(
		func() {
			ele := li.findElement(val)
			if ele != nil {
				li.li.Remove(ele)
			}
		},
	)
}

// Pop 删除个指定元素
func (li *List[T]) Pop(index int) (val T, err error) {
	li.mu.RunW(
		func() {
			ele := li.getElement(index)
			if ele != nil {
				val = li.li.Remove(ele).(T)
			} else {
				err = KeyError
			}
		},
	)
	return
}

// Insert 插入指定元素
func (li *List[T]) Insert(index int, val T) (err error) {
	li.mu.RunW(
		func() {
			ele := li.getElement(index)
			if ele == nil {
				err = KeyError
			}
			li.li.InsertBefore(val, ele)
		},
	)
	return
}

// Values 内容
func (li *List[T]) Values() (values []T) {
	li.mu.RunR(
		func() {
			values = make([]T, li.li.Len())
			ele := li.li.Front()
			for i := 0; i < li.li.Len(); i++ {
				values[i] = ele.Value.(T)
				ele = ele.Next()
			}
		},
	)
	return
}
