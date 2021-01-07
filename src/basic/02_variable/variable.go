package main

import "fmt"

// Go 语言声明变量的 5 种方式
func main() {

	// 1. 一行声明一个变量
	var sum = 100
	var rate float32 = 0.50
	fmt.Println(sum, rate)
	fmt.Println()

	// 2. 多个变量一起声明
	var (
		name   string = "flexia"
		age    int    = 22
		gender string = "男"
	)
	fmt.Println(name, age, gender)
	fmt.Println()

	// 3. 推导声明写法(声明和初始化一个变量）
	// 编译器会自动判断变量类型，只能用于函数内部
	student := "学go的学生"
	fmt.Println(student)
	fmt.Println()

	// 4. 一行声明和初始化多个变量
	a, b := 173, 125
	// 常用于变量交换
	b, a = a, b
	fmt.Println(a, b)

	school, score := "关山口男子职业技术学院", 100
	fmt.Println(school, score)
	fmt.Println()

	// 5. new 函数声明一个指针变量
	var num = 1
	var ptr1 = &num
	fmt.Println("ptr1 address: ", ptr1)
	fmt.Println("ptr1 value: ", *ptr1)

	ptr2 := new(int)
	fmt.Println("ptr2 address: ", ptr2)
	fmt.Println("ptr2 value: ", *ptr2)
	fmt.Println()

	// 上述方法声明变量，只能声明一次，多次声明会报错
	// 匿名变量，也称为占位符，不分配内存，也不占用内存，多次声明不会有问题
	c, _ := getData()
	_, d := getData()
	fmt.Println(c, d)
}

func getData() (int, int) {
	return 100, 200
}
