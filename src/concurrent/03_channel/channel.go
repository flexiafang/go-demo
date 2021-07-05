package main

import (
	"fmt"
	"time"
)

/*
1. 信道
Go语言之所以开始流行起来，很大一部分原因是因为它自带的并发机制。
如果说 goroutine 是 Go语言程序的并发体的话，那么 channel（信道） 就是 它们之间的通信机制。
channel，是一个可以让一个 goroutine 与另一个 goroutine 传输信息的通道，连接多个 goroutine 程序 ，它是一种队列式的数据结构，遵循先入先出的规则。

2. 信道的定义和使用
每个信道都只能传递一种数据类型的数据，声明的时候得指定数据类型。
声明后的信道，其零值是 nil，无法直接使用，必须配合 make 函进行初始化。
信道的数据操作，无非就两种：发送数据与读取数据。
信道用完了，可以对其进行关闭，避免有人一直在等待。但是关闭信道后，接收方仍然可以从信道中取到数据，只是接收到的会永远是 0。
对一个已关闭的信道再关闭，是会报错的。当从信道中读取数据时，可以有多个返回值，其中第二个可以表示 信道是否被关闭。
*/
func myFunc1() {
	// 声明并初始化信道
	var pipeline chan int
	pipeline = make(chan int, 1)
	// 往信道中发送数据
	pipeline <- 1
	// 从信道中取出数据
	data := <-pipeline
	fmt.Println(data) // 1
	// 关闭信道
	close(pipeline)
	// 判断信道是否关闭
	if _, ok := <-pipeline; !ok {
		fmt.Println("信道被关闭")
	}
	fmt.Println()
}

/*
3. 信道的长度和容量
一般创建信道都是使用 make 函数，make 函数接收两个参数：
  - 第一个参数：必填，指定信道类型
  - 第二个参数：选填，不填默认为 0，指定信道的容量（可缓存多少数据）
信道的容量：
  - 当容量为 0 时，说明信道中不能存放数据，在发送数据时，必须要求立马有人接收，否则会报错。此时的信道称之为无缓冲信道。
  - 当容量为 1 时，说明信道只能缓存一个数据，若信道中已有一个数据，此时再往里发送数据，会造成程序阻塞。 利用这点可以利用信道来做锁。
  - 当容量大于 1 时，信道中可以存放多个数据，可以用于多个协程之间的通信管道，共享资源。
至此我们知道，信道就是一个容器。
*/
func myFunc2() {
	pipeline := make(chan int, 10)
	fmt.Printf("信道可缓冲 %d 个数据\n", cap(pipeline)) // 10
	pipeline <- 1
	fmt.Printf("信道中当前有 %d 个数据\n", len(pipeline)) // 1
	fmt.Println()
}

/*
4. 缓冲信道与无缓冲信道
缓冲信道：允许信道里存储一个或多个数据，发送端和接收端可以处于异步的状态。
无缓冲信道：在信道里无法存储数据，接收端必须先于发送端准备好，以确保发送完数据，有人立马接收数据，否则发送端就会造成阻塞，也就是说发送端和接收端是同步运行的。
*/
func myFunc3() {
	pipeline := make(chan int)

	go func() {
		num := 100
		fmt.Printf("准备发送数据: %d\n", num)
		pipeline <- num
	}()

	go func() {
		num := <-pipeline
		fmt.Printf("接收到的数据是: %d\n", num)
	}()

	// 函数sleep，使得上面两个goroutine有机会执行
	time.Sleep(1)
	fmt.Println()
}

/*
5. 双向信道与单向信道
默认情况下定义的信道都是双向的，可发送数据，也可接收数据。
单向信道，可以细分为 只读信道 和 只写信道。
关键代码：定义别名类型。
<-chan 表示这个信道，只能从里发出数据，对于程序来说就是只读。
chan<- 表示这个信道，只能从外面接收数据，对于程序来说就是只写。
*/

//定义只写信道类型
type Sender chan<- int

//定义只读信道类型
type Receiver <-chan int

func myFunc4() {
	pipeline := make(chan int)

	go func() {
		var sender Sender = pipeline
		num := 1
		fmt.Printf("准备发送数据：%d\n", num)
		sender <- num
	}()

	go func() {
		var receiver Receiver = pipeline
		num := <-receiver
		fmt.Printf("接收到的数据是：%d\n", num)
	}()

	time.Sleep(1)
	fmt.Println()
}

/*
6. 遍历信道
遍历信道，可以使用 for 搭配 range 关键字，在 range 时，要确保信道是处于关闭状态，否则循环会阻塞。
*/
func myFunc5() {
	pipeline := make(chan int, 10)
	go fibonacci(pipeline)
	for k := range pipeline {
		fmt.Printf("%d ", k)
	}
	fmt.Print("\n\n")
}

func fibonacci(c chan int) {
	n := cap(c)
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// 记得 close 信道
	// 不然主函数中遍历完并不会结束，而是会阻塞。
	close(c)
}

/*
7. 用信道做锁
当信道里的数据量已经达到设定的容量时，此时再往里发送数据会阻塞整个程序。
利用这个特性，可以用来当程序的锁。
*/
func myFunc7() {
	// 设置容量为 1 的缓冲信道
	pipeline := make(chan bool, 1)

	var x int
	for i := 0; i < 1000; i++ {
		go increment(pipeline, &x)
	}

	time.Sleep(time.Second)
	// 如果不加锁，结果可能小于1000
	fmt.Println("x 的值为", x)
}

func increment(ch chan bool, x *int) {
	ch <- true
	// 由于 x=x+1 不是原子操作
	// 所以应避免多个协程对x进行操作
	// 使用容量为1的信道可以达到锁的效果
	*x = *x + 1
	<-ch
}

/*
8. 几个注意事项
关闭一个未初始化的 channel 会产生 panic
重复关闭同一个 channel 会产生 panic
向一个已关闭的 channel 发送消息会产生 panic
从已关闭的 channel 读取消息不会产生 panic，且能读出 channel 中还未被读取的消息，若消息均已被读取，则会读取到该类型的零值。
从已关闭的 channel 读取消息永远不会阻塞，并且会返回一个为 false 的值，用以判断该 channel 是否已关闭（x,ok := <-ch）
关闭 channel 会产生一个广播机制，所有向 channel 读取消息的 goroutine 都会收到消息
channel 在 Golang 中是一等公民，它是线程安全的，面对并发问题，应首先想到 channel。
*/

func main() {
	myFunc1()
	myFunc2()
	myFunc3()
	myFunc4()
	myFunc5()
	myFunc7()
}

/*
1
信道被关闭

信道可缓冲 10 个数据
信道中当前有 1 个数据

准备发送数据: 100
接收到的数据是: 100

准备发送数据：1
接收到的数据是：1

1 1 2 3 5 8 13 21 34 55

x 的值为 1000
*/
