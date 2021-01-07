package main

import (
	"fmt"
	"time"
)

// 异常机制 panic 和 recover
func main() {

	// myFunc1()

	myFunc2(20)

	defer fmt.Println("in main")
	go myFunc3()
	time.Sleep(time.Second)

	fmt.Println("everything is ok")
}

// 手动触发 panic，直接报错宕机
func myFunc1() {

	panic("crash")

	/*
		panic: crash

		goroutine 1 [running]:
		main.myFunc1(...)
			go-demo/src/basic/09_panic_recover/panic_recover.go:12
		main.main()
			go-demo/src/basic/09_panic_recover/panic_recover.go:6 +0x45
	*/
}

// 在 defer 函数中使用 recover 捕获 panic
func myFunc2(x int) {
	defer func() {
		// recover() 可以捕获到 panic
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	// 数组越界触发 panic
	var arr [10]int
	arr[x] = 88

	/*
		runtime error: index out of range [20] with length 10
	*/
}

// panic 无法跨协程，子协程触发 panic 只能触发自己协程内的 defer，而不能调用 main 协程里的 defer
func myFunc3() {
	defer println("in goroutine")
	panic("panic here")

	/*
		in goroutine
		panic: panic here

		goroutine 6 [running]:
		main.myFunc3()
			go-demo/src/basic/09_panic_recover/panic_recover.go:55 +0x78
		created by main.main
			go-demo/src/basic/09_panic_recover/panic_recover.go:16 +0xbd
	*/
}
