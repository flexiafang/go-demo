package main

import (
	"fmt"
	"time"
)

/*
一个 goroutine 本身就是一个函数，如果在调用前加一个关键字 go ，就开启了一个 goroutine。
一个 Go 程序的入口通常是 main 函数,程序启动后， main 函数最先运行，称之为 main goroutine。
在 main 中或者其下调用的代码中才可以使用 go func() 的方法来启动协程。
main 的地位相当于主线程，当 main 函数执行完成后，这个线程也就终结了，其下的运行着的所有协程也不管代码是不是还在跑，也得乖乖退出。
*/

var times = 0

func mygo(name string) {
	for i := 0; i < 10; i++ {
		times++
		fmt.Println(times, ": In goroutine", name)
		time.Sleep(time.Millisecond)
	}
}

func main() {
	// 启动两个协程
	go mygo("协程1")
	go mygo("协程2")
	// 可以使用 time.Sleep 来使 main 阻塞，使其他协程能够有机会运行完全，但要注意的是，这并不是推荐的方式。
	time.Sleep(time.Second)

	/*
		1 : In goroutine 协程1
		2 : In goroutine 协程2
		4 : In goroutine 协程2
		3 : In goroutine 协程1
		5 : In goroutine 协程2
		6 : In goroutine 协程1
		7 : In goroutine 协程1
		8 : In goroutine 协程2
		9 : In goroutine 协程2
		10 : In goroutine 协程1
		11 : In goroutine 协程1
		12 : In goroutine 协程2
		13 : In goroutine 协程2
		14 : In goroutine 协程1
		15 : In goroutine 协程1
		16 : In goroutine 协程2
		17 : In goroutine 协程2
		18 : In goroutine 协程1
		19 : In goroutine 协程2
		20 : In goroutine 协程1
	*/
}
