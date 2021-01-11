package main

import (
	"fmt"
	"os"
)

func main() {
	myFunc1()
	myFunc2()
	_, _ = myFunc3()
}

/*
静态类型：变量声明时的类型
动态类型：程序运行时才能看见的类型

空接口可以承接任意类型的值，当赋值后，变量的静态类型永远是 interface{}，但是动态类型会改变
*/
func myFunc1() {
	var age int = 20

	var i interface{}
	i = age
	i = "go"

	fmt.Println(i)
}

/*
每个接口变量，实际上都是由一个 pair 对（type 和 value）组合而成，pair 对中记录着实际变量的值和类型。
*/
func myFunc2() {
	age := (interface{})(22)
	fmt.Printf("type: %T, value: %v\n", age, age) // type: int, value: 22
}

/*
两种接口：
  - 带有一组方法的接口 iface
  - 不带有方法的接口 eface
*/
func myFunc3() (res interface{}, err error) {
	var eface interface{}

	file, err := os.OpenFile("demo.txt", os.O_RDONLY, 0)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	eface = file
	fmt.Println(file)
	return eface, nil
}
