package lock

import "sync"

// RWLocker 读写锁
type RWLocker struct {
	l sync.RWMutex
}

// RunW 写锁运行
func (rw *RWLocker) RunW(f func()) {
	With(&rw.l, f)
}

// RunR 读锁运行
func (rw *RWLocker) RunR(f func()) {
	With(rw.l.RLocker(), f)
}

// Locker 普通锁
type Locker struct {
	l sync.Mutex
}

func (lo *Locker) Run(f func()) {
	With(&lo.l, f)
}
