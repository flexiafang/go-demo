package main

import (
	"fmt"
	"strconv"
)

/*
类型断言 Type Assertion

可以通过类型断言：
  1. 检查接口对象 i 是否为 nil
  2. 检查接口对象 i 存储的值是否为某个类型

使用方式：
  1. t := i.(T)
     可以断言一个接口对象 i 里不是 nil，并且接口对象 i 存储的值的类型是 T
     - 断言成功会返回值给 t
     - 断言失败会触发 panic
  2. t, ok := i.(T)
     可以断言一个接口对象 i 里不是 nil，并且接口对象 i 存储的值的类型是 T
     - 断言成功会返回其类型给 t，并且此时 ok 的值 为 true
     - 断言失败不会触发 panic，而是将 ok 的值设为 false，此时 t 为 T 的零值
*/

func main() {
	myFunc1()
	myFunc2()
	myFunc3()
}

func myFunc1() {
	var i interface{} = 10

	t1 := i.(int)
	fmt.Println(t1) // 10

	// t2 := i.(string)
	// fmt.Println(t2)
	/*
		panic: interface conversion: interface {} is int, not string

		goroutine 1 [running]:
		main.myFunc1()
			go-demo/src/object/04_type_assertion/type_assertion.go:33 +0xb3
		main.main()
			go-demo/src/object/04_type_assertion/type_assertion.go:24 +0x27
	*/
}

func myFunc2() {

	var i interface{} = 10

	t1, ok := i.(int)
	fmt.Printf("%d-%t\n", t1, ok) // 10-true

	t2, ok := i.(string)
	fmt.Printf("%s-%t\n", t2, ok) // -false

	var k interface{} // nil

	t3, ok := k.(interface{})
	fmt.Println(t3, "-", ok) // <nil> - false

	k = 10

	t4, ok := k.(interface{})
	fmt.Printf("%d-%t\n", t4, ok) // 10-true

	t5, ok := k.(int)
	fmt.Printf("%d-%t\n", t5, ok) // 10-true
}

/*
如果需要区分多种类型，可以使用 type switch 断言，这个将会比一个一个进行类型断言更简单、直接、高效
*/

func myFunc3() {
	findType(10)      // 10 is int
	findType("hello") // hello is string

	var k interface{}
	findType(k) // <nil> is nil

	findType(0.5) // 0.5 not type matched
}

func findType(i interface{}) {
	var res string

	switch x := i.(type) {
	case int:
		res = strconv.Itoa(x) + " is int"
	case string:
		res = x + " is string"
	case nil:
		res = fmt.Sprintf("%v", x) + " is nil"
	default:
		res = fmt.Sprintf("%v", x) + " no type matched"
	}

	fmt.Println(res)
}
