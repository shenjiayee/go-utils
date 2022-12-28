package slices

import "sort"

func W[T any](f func(T)) func(T, int) {
	return func(t T, _ int) {
		f(t)
	}
}

func Wa[T any, U any](f func(T) U) func(T, int) U {
	return func(t T, _ int) U {
		return f(t)
	}
}

func Expand[T any](slice *[]T, count int) *[]T {
	for count > 0 {
		*slice = append(*slice, *new(T))
		count--
	}
	return slice
}

// Index 判断一个元素在slice中的位置
func Index[T comparable](slice []T, elem T) int {
	for index, val := range slice {
		if val == elem {
			return index
		}
	}
	return -1
}

// First 找到第一个符合条件的元素
func First[T any](slice []T, f func(T) bool) int {
	for index, val := range slice {
		if f(val) {
			return index
		}
	}
	return -1
}

// Contains 判断一个元素是否在slice中
func Contains[T comparable](slice []T, elem T) bool {
	return Any(slice, func(val T) bool { return val == elem })
}

// Copy 复制
func Copy[T any](origin []T) []T {
	newSlice := make([]T, len(origin))
	copy(newSlice, origin)
	return newSlice
}

// Any 任一
func Any[T any](slice []T, f func(T) bool) bool {
	for _, val := range slice {
		if f(val) {
			return true
		}
	}
	return false
}

// All 所有
func All[T any](slice []T, f func(T) bool) bool {
	for _, val := range slice {
		if !f(val) {
			return false
		}
	}
	return true
}

// Map 映射
func Map[T any, U any](slice []T, f func(T, int) U) []U {
	newSlice := make([]U, len(slice))
	for index, val := range slice {
		newSlice[index] = f(val, index)
	}
	return newSlice
}

// ForEach 依次执行
func ForEach[T any](slice []T, f func(T, int)) {
	for index, val := range slice {
		f(val, index)
	}
}

// ForEachP 依次执行，传指针
// 我觉得这个函数非常没有必要
func ForEachP[T any](slice []T, f func(*T, int)) {
	for i := 0; i < len(slice); i++ {
		f(&slice[i], i)
	}
}

// Find 找到第一个
func Find[T any](slice []T, f func(T, int) bool) (T, bool) {
	for index, val := range slice {
		if f(val, index) {
			return val, true
		}
	}
	return *new(T), false
}

// Filter 过滤
func Filter[T any](slice []T, f func(T, int) bool) []T {
	newSlice := make([]T, 0)
	ForEach(
		slice,
		func(elem T, index int) {
			if f(elem, index) {
				newSlice = append(newSlice, elem)
			}
		},
	)
	return newSlice
}

// Sort 排序 if compareFn(a, b) == true, a 在 b前
func Sort[T any](slice []T, compareFn func(T, T) bool) {
	lessFunc := func(i, j int) bool {
		return compareFn(slice[i], slice[j])
	}
	sort.SliceStable(slice, lessFunc)
}

// NewSlice 构造一个有默认值的slice
func NewSlice[T any](length int, fillWith T) []T {
	newSlice := make([]T, length)
	ForEach(
		newSlice,
		func(_ T, index int) {
			newSlice[index] = fillWith
		},
	)
	return newSlice
}
