package main

import "fmt"

// Go 语言的流程控制
func main() {

	/*
		3. for 循环
		for [condition |  ( init; condition; increment ) | Range] {
		   statement(s);
		}
	*/

	// 接一个表达式
	a := 1
	for a <= 5 {
		fmt.Print(a, " ")
		a++
	}
	fmt.Println()

	// 接三个表达式
	for b := 1; b <= 5; b++ {
		fmt.Print(b, " ")
	}
	fmt.Println()

	// 不接表达式，表示无限循环
	var c int = 1
	for {
		if c > 5 {
			break
		}
		fmt.Print(c, " ")
		c++
	}
	fmt.Println()

	// 接 for-range 语句，range 会返回两个数据，索引和数据
	d := [...]string{"hello", "world", "go"}

	for _, item := range d {
		fmt.Print(item, " ")
	}
	fmt.Println()

	// 用一个变量只会接收到索引
	for i := range d {
		fmt.Print(i, " ")
	}
	fmt.Println()

	/*
		4. goto 语句
		goto 后接一个标签，告诉Go程序下一步要执行哪里的代码
		goto 语句与标签之间不能有变量声明，否则编译错误
	*/

	i := 1
flag:
	for i <= 10 {
		if i%2 == 1 {
			i++
			goto flag
		}
		fmt.Print(i, " ")
		i++
	}
	fmt.Println()
}
