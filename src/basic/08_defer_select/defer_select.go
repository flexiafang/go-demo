package main

import (
	"fmt"
	"sync"
	"time"
)

func myFunc1() {
	fmt.Println("B")
}

var name string = "go"

func myFunc2() string {
	defer func() {
		name = "python"
	}()

	fmt.Printf("myFunc2 函数里的 name： %s\n", name)
	return name
}

func main() {

	/*
		5. defer 延迟语句
		只要在defer后面跟一个函数的调用，就能将这个函数调用延迟到当前函数执行完后再执行
		defer只是延时函数调用，但是传递给函数的变量值不受后续程序的影响
		defer是反序调用的
		defer是在return之后再调用的
	*/

	defer myFunc1()
	fmt.Println("A")

	str := "go"
	defer fmt.Println(str)
	str = "hust"
	fmt.Println(str)

	myName := myFunc2()
	fmt.Printf("main 函数里的 name： %s\n", name)
	fmt.Printf("main 函数里的 myName： %s\n", myName)

	/*
		6. select 信道/通道
		select-case 用法比较单一，它仅能用于 信道/通道 的相关操作
		select 在执行过程中，必须命中其中的某一分支。若没有命中任何一个 case，就会进入 default 里的代码分支
		但如果没有写 default 分支，select 就会阻塞，直到有某个 case 可以命中，而如果一直没有命中，select 就会抛出 deadlock 的错误
		select 具有随机性
		当 case 里的信道始终没有接收到数据，也没有 default 语句时，select 整体就会阻塞，这时候就可以手动设置一个超时时间
	*/

	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	// c2 <- "hello"

	// 开启一个协程给信道发送数据
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		c2 <- "hello"
		wg.Done()
	}()

	wg.Wait()

	select {
	case msg1 := <-c1:
		fmt.Println("c1 received: ", msg1)
	case msg2 := <-c2:
		fmt.Println("c2 received: ", msg2)
	default:
	}

	// select的超时
	c3 := make(chan string, 1)
	c4 := make(chan string, 1)
	timeout := make(chan bool, 1)

	go makeTimeout(timeout, 0)

	select {
	case msg1 := <-c3:
		fmt.Println("c3 received: ", msg1)
	case msg2 := <-c4:
		fmt.Println("c4 received: ", msg2)
	case <-timeout:
		fmt.Println("Timeout, exit.")
	}

	// select用于channel的读取/写入
	c5 := make(chan int, 2)
	c5 <- 2
	select {
	case c5 <- 4:
		fmt.Println("c5 received: ", <-c5)
		fmt.Println("c5 received: ", <-c5)
	default:
		fmt.Println("Channel blocking.")
	}
}

func makeTimeout(ch chan bool, t int) {
	time.Sleep(time.Second * time.Duration(t))
	ch <- true
}
