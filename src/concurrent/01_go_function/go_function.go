package main

import "fmt"

/*
1. 关于函数
函数是基于功能或逻辑进行封装的可复用的代码结构。
- 带有名字的普通函数
- 没有名字的匿名函数
Go语言是编译型语言，所以函数编写的顺序无关紧要，不像Python函数在位置上需要定义在调用之前。

2. 函数声明
func 函数名(参数列表) (返回值列表) {
	函数体
}

3. 可变参数
多个类型一致的参数：使用 ...Type 表示一个元素为 Type 类型的切片，用来接收调用者传入的参数。
多个类型不一致的参数：使用 ...interface{} 接收。
... 是Go语言的语法糖，只能在定义函数时使用，如果在该函数下有多个类型的参数，这个语法糖必须是最后一个参数。
... 还有一个用法，就是用来解序列，也只能在给函数传递参数时使用，将函数的可变参数（一个切片）一个一个取出来，传递给另一个可变参数的函数，而不是传递可变参数变量本身。

4. 函数的返回值
没有指明返回值类型时，函数体可以用 return 来结束函数的运行，但 return 后面不能跟任何对象。
Go支持一个函数返回多个值，也支持返回带有变量名的值。

5. 匿名函数
func(参数列表) (返回值列表) {
	函数体
}
匿名函数一般定义后立即使用，亦或是作为回调函数使用。
*/

// 使用 ... 接收可变参数
func sum(args ...int) int {
	var sum int
	for _, v := range args {
		sum += v
	}
	return sum
}

func Sum(args ...int) int {
	// 利用 ... 解序列
	result := sum(args...)
	return result
}

func double(a int) (b int) {
	// 不能使用 :=，因为在返回值那里已经声明为 int
	b = a * 2
	// 不需要指明写回哪个变量，在返回值类型那里已经指定了
	return
}

func visit(list []int, f func(int)) {
	for _, v := range list {
		// 执行回调函数
		f(v)
	}
}

func main() {
	fmt.Println(Sum(1, 2, 3)) // 6

	fmt.Println(double(2)) // 4

	visit([]int{1, 2, 3, 4}, func(v int) { fmt.Println(v) }) // 1 2 3 4
}
