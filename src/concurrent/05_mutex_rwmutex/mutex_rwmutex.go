package main

import (
	"fmt"
	"sync"
	"time"
)

/*
面对 golang 中的并发问题，首先应该考虑使用信道。
如果必须使用共享内存实现并发编程，则需要了解 golang 中的锁机制。

1. 互斥锁 Mutex
使用互斥锁是为了来保护一个资源不会因为并发操作而引起冲突导致数据不准确。
*/

func add1(count *int, wg *sync.WaitGroup) {
	for i := 0; i < 3333; i++ {
		*count = *count + 1
	}
	wg.Done()
}

/*
开启了三个协程，每个协程分别往 count 这个变量加 1000 次 1，理论上看，最终的 count 值应为 3000。
可运行多次的结果，都不相同，原因就在于这三个协程在执行时，先读取 count 再更新 count 的值，而这个过程并不具备原子性，所以导致了数据的不准确。
*/
func myFunc1() {
	var wg sync.WaitGroup
	count := 0

	wg.Add(3)
	go add1(&count, &wg)
	go add1(&count, &wg)
	go add1(&count, &wg)

	wg.Wait()
	fmt.Println("count:", count)
}

/*
给 add 函数加上 Mutex 互斥锁，要求同一时刻，仅能有一个协程能对 count 操作。

Mutex 锁的两种定义方法：
 - 第一种
	var lock *sync.Mutex
	lock = new(sync.Mutex)
 - 第二种
	lock := &sync.Mutex{}
*/
func add2(count *int, wg *sync.WaitGroup, lock *sync.Mutex) {
	for i := 0; i < 3333; i++ {
		lock.Lock()
		*count = *count + 1
		lock.Unlock()
	}
	wg.Done()
}

/*
使用 Mutext 锁虽然很简单，但仍然有几点需要注意：
 - 同一协程里，不要在尚未解锁时再次使加锁
 - 同一协程里，不要对已解锁的锁再次解锁
 - 加了锁后，别忘了解锁，必要时使用 defer 语句
*/
func myFunc2() {
	var wg sync.WaitGroup
	lock := &sync.Mutex{}
	count := 0

	wg.Add(3)
	go add2(&count, &wg, lock)
	go add2(&count, &wg, lock)
	go add2(&count, &wg, lock)

	wg.Wait()
	fmt.Println("count:", count)
}

/*
2. 读写锁 RWMutex
RWMutex 将程序对资源的访问分为读操作和写操作：
 - 为了保证数据的安全，它规定了当有人还在读取数据（即读锁占用）时，不允许有人更新这个数据（即写锁会阻塞）。
 - 为了保证程序的效率，多个人（线程）读取数据（拥有读锁）时，互不影响不会造成阻塞，它不会像 Mutex 那样只允许有一个人（线程）读取同一个数据。

定义一个 RWMuteux 锁，有两种方法
 - 第一种
	var lock *sync.RWMutex
	lock = new(sync.RWMutex)
 - 第二种
	lock := &sync.RWMutex{}

RWMutex 里提供了两种锁，每种锁分别对应两个方法，为了避免死锁，两个方法应成对出现，必要时请使用 defer。
 - 读锁：调用 RLock 方法开启锁，调用 RUnlock 释放锁
 - 写锁：调用 Lock 方法开启锁，调用 Unlock 释放锁（和 Mutex类似）
*/
func myFunc3() {
	lock := &sync.RWMutex{}
	fmt.Println("开启写锁")
	lock.Lock()

	for i := 0; i < 4; i++ {
		go func(i int) {
			fmt.Printf("第 %d 个协程准备开始...\n", i)
			lock.RLock()
			fmt.Printf("第 %d 个协程获得读锁, sleep 1s 后，释放锁\n", i)
			time.Sleep(time.Second)
			lock.RUnlock()
		}(i)
	}

	time.Sleep(time.Second * 2)
	fmt.Println("准备释放写锁，读锁不再阻塞")
	// 写锁一释放，读锁就自由了
	lock.Unlock()

	// 由于会等到读锁全部释放，才能获得写锁
	// 因为这里一定会在上面 4 个协程全部完成才能往下走
	lock.Lock()
	fmt.Println("程序退出...")
	lock.Unlock()
}

func main() {
	myFunc1()
	myFunc2()
	myFunc3()
}

/*
count: 7444
count: 9999
开启写锁
第 3 个协程准备开始...
第 1 个协程准备开始...
第 2 个协程准备开始...
第 0 个协程准备开始...
准备释放写锁，读锁不再阻塞
第 0 个协程获得读锁, sleep 1s 后，释放锁
第 3 个协程获得读锁, sleep 1s 后，释放锁
第 1 个协程获得读锁, sleep 1s 后，释放锁
第 2 个协程获得读锁, sleep 1s 后，释放锁
程序退出...
*/
