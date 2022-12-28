package lock

import "sync"

func With(l sync.Locker, f func()) {
	defer l.Unlock()
	l.Lock()
	f()
}
