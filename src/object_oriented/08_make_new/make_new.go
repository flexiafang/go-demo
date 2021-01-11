package main

import "fmt"

func main() {
	myFunc1()
	myFunc2()
}

type Person struct {
	name string
	age  int
}

/*
new(Type) 函数只能传递一个参数，该参数为一个任意类型
new 函数会分配内存、设置零值、返回指针
*/
func myFunc1() {
	num := new(int)
	fmt.Println(*num) // 0

	p := new(Person)
	p.name = "flexia"
	fmt.Println(*p) // {flexia 0}
}

/*
make(Type, size) 函数用来为 slice、map、chan 类型分配内存和初始化一个对象
make 函数返回类型的本身而不是指针，返回值依赖于具体传入的引用类型
*/
func myFunc2() {
	a := make([]int, 2, 10)
	fmt.Println(a, len(a), cap(a)) // [0 0] 2 10

	b := make(map[string]int)
	fmt.Println(b, len(b)) // map[] 0

	c := make(chan int, 10)
	fmt.Println(c, len(c), cap(c)) // 0xc0000d6000 0 10
}
