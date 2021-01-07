package main

import "fmt"

// Go语言中的数据类型
func main() {

	/*
		7. map字典
		若干个键值对映射组合在一起的数据结构
		哈希表的一种实现，所以key必须是可哈希的，不能是切片、字典、函数
	*/

	// 声明初始化map
	var map1 map[string]int = map[string]int{"english": 80, "chinese": 90}
	map2 := map[string]int{"english": 80, "chinese": 90}
	map3 := make(map[string]int)
	map3["english"] = 80
	map3["chinese"] = 90
	fmt.Println(map1, map2, map3) // map[chinese:90 english:80]

	// 添加、更新或访问元素
	map3["math"] = 100
	fmt.Println(map3) // map[chinese:90 english:80 math:100]

	// 删除元素
	delete(map3, "math")
	fmt.Println(map3) // map[chinese:90 english:80]

	// 访问不存在的key时，返回value类型对应的零值
	fmt.Println(map3["physics"]) // 0

	// 判断key是否存在
	// 下标访问时会返回两个值，第二个返回值表示对应的key是否存在
	if math, ok := map3["math"]; ok { // false
		fmt.Printf("math 的值是 %d\n", math)
	} else {
		fmt.Println("math 不存在")
	}

	// 如何对map进行循环
	// 获取key和value
	for subject, score := range map3 {
		fmt.Println(subject, score)
	}

	// 只获取key，不使用占位符
	for subject := range map3 {
		fmt.Println(subject)
	}

	// 只获取value，用一个占位符代替
	for _, score := range map3 {
		fmt.Println(score)
	}

	/*
		8. bool型
		Go语言中true和false与1和0并不相等，并且更加严格，不同类型无法进行比较
		- 与 &&
		- 或 ||
		- 非 ！
		- 短路
	*/

	var male bool = true
	// fmt.Println(male == 1)
	fmt.Println(!male == false || male != false)

	/*
		9. 指针
		指针就是内存地址
		- 普通变量：存数据值本身
		- 指针变量：存值的内存地址
		指针的两种操作：
		- &：从一个变量中取得内存地址
		- *：取得指针指向的变量或变量值
		切片与指针一样，都是引用类型。如果我们想通过一个函数改变一个数组的值，有两种方法
		- 将这个数组的切片做为参数传给函数
		- 将这个数组的指针做为参数传给函数
		建议使用第一种方法，因为第一种方法，写出来的代码会更加简洁，易读
	*/

	// 指针的创建
	num := 1
	ptr1 := &num
	ptr2 := new(int)
	*ptr2 = num
	var ptr3 *int
	ptr3 = &num
	fmt.Println(ptr1, ptr2, ptr3)

	// 指针的类型，* + 所指向变量值的数据类型
	fmt.Printf("指针类型是 %T\n", ptr1) // *int

	// 指针与切片
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)
	modify1(&arr)
	fmt.Println(arr)
	modify2(arr[:])
	fmt.Println(arr)
}

func modify1(arr *[3]int) {
	(*arr)[0] = 0
}

func modify2(arr []int) {
	arr[2] = 0
}
