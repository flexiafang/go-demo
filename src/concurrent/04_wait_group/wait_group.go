package main

import (
	"fmt"
	"sync"
)

/*
1. 使用信道来标记子协程的任务完成
*/
func myFunc1() {
	done := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d ", i)
		}
		fmt.Printf("\n\n")
		done <- true
	}()
	<-done
}

/*
2. 使用 WaitGroup
sync 包提供的 WaitGroup 类型，实例化完成后，就可以使用它的几个方法：
	- Add：初始值为 0，你传入的值会往计数器上加，这里直接传入子协程的数量
	- Done：当某个子协程完成后，可调用此方法，会从计数器上减一，通常可以使用 defer 来调用。
	- Wait：阻塞当前协程，直到实例里的计数器归零。
*/
func myFunc2() {
	var wg sync.WaitGroup

	wg.Add(2)
	go worker(1, &wg)
	go worker(2, &wg)

	wg.Wait()
}

func worker(x int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("worker %d: %d\n", x, i)
	}
}

func main() {
	myFunc1()
	myFunc2()
}

/*
0 1 2 3 4

worker 2: 0
worker 2: 1
worker 2: 2
worker 2: 3
worker 2: 4
worker 1: 0
worker 1: 1
worker 1: 2
worker 1: 3
worker 1: 4
*/
