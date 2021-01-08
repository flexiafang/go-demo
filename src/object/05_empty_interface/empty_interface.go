package main

import "fmt"

func main() {
	myFunc1()
	myFunc2()
	myFunc3()
	myFunc4()
}

/*
空接口：
是特殊形式的接口类型，空接口中没有定义任何方法
所有类型都至少实现了空接口

使用空接口：
  1. 直接使用 interface{} 声明一个实例，这个实例可以承载任意类型的值
  2. 使用空接口可以让方法接收任意个任意类型的值
  3. 定义可以接收任意类型的 array、slice、map、struct
*/

type emptyInterface interface {
}

func myFunc1() {
	var i emptyInterface
	fmt.Printf("type: %T, value: %v\n", i, i) // type: <nil>, value: <nil>

	i = 1
	fmt.Println(i) // 1

	i = "hello"
	fmt.Println(i) // hello

	i = false
	fmt.Println(i) // false
}

// 接收任意个任意类型的值
func func1(ifaces ...interface{}) {
	for _, iface := range ifaces {
		fmt.Println(iface)
	}
}

func myFunc2() {
	func1(0, "hello", true)

	/*
		0
		hello
		true
	*/
}

func myFunc3() {
	any := make([]interface{}, 5)
	any[0] = 1
	any[1] = "hello"
	any[2] = []int{1, 2, 3, 4}

	for _, value := range any {
		fmt.Println(value)
	}

	/*
		1
		hello
		[1 2 3 4]
		<nil>
		<nil>
	*/
}

/*
空接口需要注意的点：
  1. 不能将空接口类型的对象赋给某个固定类型的对象
  2. 空接口承载数组和切片之后，不能再进行切片
  3. 使用空接口接收任意类型的参数时，静态类型是 interface{}，但动态类型并不确定，需要使用类型断言
*/

func myFunc4() {
	var i interface{}

	i = 1
	// var a int = i
	/*
		cannot use i (type interface {}) as type int in assignment: need type assertion
	*/

	sli := []int{1, 2, 3, 4, 5}
	i = sli
	// g := i[1:3]
	/*
		cannot slice i (type interface {})
	*/

	fmt.Println(i)

	func2(10)
	func2("go")
	/*
		参数类型是 int
		参数类型是 string
	*/
}

func func2(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("参数类型是 int")
	case string:
		fmt.Println("参数类型是 string")
	}
}
