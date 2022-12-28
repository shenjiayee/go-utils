package async

import (
	"sync"
)

// Async
// 可以作为其它struct的组成
// 实现`Future`接口
// 通常情况下没有用处
// 使用
// ```
//
//	go func() {
//		func1()
//		func2()
//		...
//	}()
//
// ```
// 即可
// 仅在多个函数并发调用，且需要join时，可以做为waitGroup的替代方法
type Async[T any] struct {
	// 用于done的初始化
	init sync.Once

	// 写入结果
	flag   sync.Once
	result T
	done   chan struct{} // 最终结果的通道，实现阻塞
}

// 如果`Async.done`不存在，初始化一个
// 从而`Async`可以零值使用
func (a *Async[T]) doneChan() chan struct{} {
	a.init.Do(func() { a.done = make(chan struct{}) })
	return a.done
}

// Done 发出结束的信号
func (a *Async[T]) Done(res T) {
	done := a.doneChan()
	a.flag.Do(
		func() {
			// 先给`Async.result`写上结果
			a.result = res
			// 给信号
			close(done)
		},
	)
}

// Result 调用后阻塞，直到`Async.Done`被调用后，返回结果
func (a *Async[T]) Result() T {
	// 等信号
	<-a.doneChan()
	return a.result
}
