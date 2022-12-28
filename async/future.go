package async

type (
	// ResultFn 异步处理使用ResultFn作为返回
	ResultFn[T any] func() T

	// Future 异步的FUTURE
	Future interface {
		// Result 阻塞直到返回结果
		Result() error
	}
)

// EnsureFuture 传入一个函数，这个函数将在goroutine中运行
// 返回一个Result函数，调用这个函数会阻塞直到返回结果
func EnsureFuture[T any](syncFn ResultFn[T]) ResultFn[T] {
	// 包装
	var async *Async[T]
	go func() {
		res := syncFn()
		async.Done(res)
	}()
	return async.Result
}
