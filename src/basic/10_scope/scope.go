package main

import "fmt"

// 理解语句块与作用域
func main() {

	/*
		显式语句块：由花括弧 {} 所包含的一系列语句
		隐式语句块：
		- 主语句块：包括所有源码，对应内置作用域
		- 包语句块：包括该包中所有的源码（一个包可能会包括一个目录下的多个文件），对应包级作用域
		- 文件语句块：包括该文件中的所有源码，对应文件级作用域
		- for、if、switch 等语句本身也在它自身的隐式语句块中，对应局部作用域

		作用域：语句块内部声明的变量是无法被外部块访问，这个块决定了内部声明的名字的作用域范围，在哪里可用，在哪里不可用
		- 内置作用域：不需要自己声明，所有的关键字和内置类型、函数都拥有全局作用域
		- 包级作用域：必須函数外声明，在该包内的所有文件都可以访问
		- 文件级作用域：不需要声明，导入即可。一个文件中通过 import 导入的包名，只在该文件内可用
		- 局部作用域：在自己的语句块内声明，包括函数、for、if 等语句块，或自定义的 {} 语句块形成的作用域，只在自己的局部作用域内可用

		不要将作用域和生命周期混为一谈
		- 声明语句的作用域对应的是一个源代码的文本区域，是一个编译时的属性
		- 而一个变量的生命周期是指程序运行时变量存在的有效时间段，在此时间区域内它可以被程序的其他部分引用；是一个运行时的概念
	*/

	switch i := 2; i * 4 {
	case 8:
		j := 0
		fmt.Println(i, j)
	default:
		// "j" is undefined here
		fmt.Println("default")
	}
	// "i"/"j" is undefined here

	var name string = "local scope"
	fmt.Println("在 main 函数中的 name：", name)
	myFunc()

	/*
		2 0
		在 main 函数中的 name： local scope
		在 myFunc 函数中的 name： file scope
	*/
}

var name string = "file scope"

func myFunc() {
	fmt.Println("在 myFunc 函数中的 name：", name)
}
