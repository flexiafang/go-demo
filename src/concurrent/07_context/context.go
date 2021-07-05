package main

/*
1. 什么是 Context 上下文

Context 接口定义如下：
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}

可以看到 Context 接口共有 4 个方法：

*/

func myFunc1() {

}

func main() {
	myFunc1()
}
