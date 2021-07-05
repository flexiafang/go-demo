package main

import (
	"fmt"
	"time"
)

/*
池化技术就是利用复用来提升性能的，那在 Golang 中需要协程池吗？
在 Golang 中，goroutine 是一个轻量级的线程，他的创建、调度都是在用户态进行，并不需要进入内核，这意味着创建销毁协程带来的开销是非常小的。
因此，大多数情况下，开发人员是不太需要使用协程池的。

抛开是否必要这个问题，单纯从技术的角度来看，怎样实现一个通用的协程池呢？
*/

// Pool 结构体
type Pool struct {
	work chan func()   // 用于接收 task 任务
	sem  chan struct{} // 用于设置协程池大小，即可同时执行的协程数量
}

// New 函数用于创建协程池对象
func New(size int) *Pool {
	return &Pool{
		make(chan func()),
		make(chan struct{}, size),
	}
}

// NewTask 往协程池中添加任务。
// 当第一次调用 NewTask 添加任务的时候，由于 work 是无缓冲通道，所以会一定会走第二个 case 的分支：使用 go worker 开启一个协程。
func (p *Pool) NewTask(task func()) {
	select {
	case p.work <- task:
	case p.sem <- struct{}{}:
		go p.worker(task)
	}
}

// worker 用于执行任务。
// 为了能够实现协程的复用，这个使用了 for 无限循环，使这个协程在执行完任务后，也不退出，而是一直在接收新的任务。
func (p *Pool) worker(task func()) {
	defer func() { <-p.sem }()
	for {
		task()
		task = <-p.work
	}
}

/*
两个函数的逻辑：
1. 如果设定的协程池数大于 2，此时第二次传入往 NewTask 传入 task，select case 的时候，
   如果第一个协程还在运行中，就一定会走第二个case，重新创建一个协程执行 task
2. 如果传入的任务数大于设定的协程池数，并且此时所有的任务都还在运行中，
   那此时再调用 NewTask 传入 task ，这两个 case 都不会命中，会一直阻塞直到有任务执行完成，
   worker 函数里的 work 通道才能接收到新的任务，继续执行。
*/

func main() {
	pool := New(2)
	for i := 0; i < 4; i++ {
		pool.NewTask(func() {
			time.Sleep(2 * time.Second)
			fmt.Println(time.Now())
		})
	}
	time.Sleep(5 * time.Second)
}

/*
2021-07-05 23:54:34.8046383 +0800 CST m=+2.024002001
2021-07-05 23:54:34.8046383 +0800 CST m=+2.024002001
2021-07-05 23:54:36.8662934 +0800 CST m=+4.085657101
2021-07-05 23:54:36.8662934 +0800 CST m=+4.085657101
可以看到四个任务分两批执行
*/
