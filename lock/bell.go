package lock

// Bell 用来等待通知的一个模型
type Bell struct {
	locker    Locker
	blockFlag bool // 如果true,直接通过
	waiter    chan<- struct{}
}

// Wait 等待消息
// 调用后会阻塞
func (r *Bell) Wait() {
	ready := make(chan struct{})
	r.locker.Run(
		func() {
			r.blockFlag = true
			r.waiter = ready
		},
	)
	<-ready
}

// Ring 通知调用`Bell.Wait`的地方
func (r *Bell) Ring() {
	r.locker.Run(
		func() {
			if r.blockFlag {
				close(r.waiter)
			}
			r.blockFlag = false
		},
	)
}
