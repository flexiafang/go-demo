package main

import (
	"fmt"
)

// Go语言中的数据类型
func main() {

	/*
		5. 数组
		数组是一个由固定长度的特定类型元素组成的序列，一个数组可以由零个或多个元素组成
		可以使用 ... 让 Go 语言自己根据实际情况来分配空间
		使用 type 关键字可以定义一个类型字面量，也就是别名类型
	*/

	// 声明数组
	var arr01 [3]int
	// 给数组元素赋值
	arr01[0] = 1
	arr01[1] = 2
	arr01[2] = 3
	fmt.Println("arr01: ", arr01)

	// 声明并直接初始化数组
	var arr02 = [3]int{1, 2, 3}
	fmt.Println("arr02: ", arr02)

	// 使用 ... 避免硬编码
	arr03 := [...]int{1, 2, 3, 4}
	fmt.Println("arr03: ", arr03)

	fmt.Printf("arr02 的类型是 %T\n", arr02) // [3]int
	fmt.Printf("arr03 的类型是 %T\n", arr03) // [4]int

	// 定义类型字面量
	type arr3 [3]int
	arr04 := arr3{1, 2, 3}
	fmt.Printf("数组 %d 的类型是 %T\n", arr04, arr04) // main.arr3

	// 偷懒的定义数组方式，指明第几个元素是多少
	arr05 := [4]int{1: 3, 2: 2}
	fmt.Println(arr05) // [0 3 2 0]

	/*
		6. 切片
		切片是对数组的一个连续片段的引用
		无法通过切片类型来确定其值的长度
	*/

	// 对数组进行片段截取获得切片
	myArr := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("切片 %d 的类型是 %T\n", myArr[0:2], myArr[0:2]) // [1 2] []int

	// 在切片时，若不指定第三个数，那么切片终止索引会一直到原数组的最后一个数
	sli1 := myArr[1:3]
	// 而如果指定了第三个数，那么切片终止索引只会到原数组的该索引值
	sli2 := myArr[1:3:4]
	// 切片的第三个数，影响的只是切片的容量，而不会影响长度
	fmt.Printf("sli1 的长度为：%d，容量为：%d\n", len(sli1), cap(sli1)) // 2 4
	fmt.Printf("sli2 的长度为：%d，容量为：%d\n", len(sli2), cap(sli2)) // 2 3

	// 直接声明切片
	var sli3 []int
	fmt.Println(sli3 == nil) // true

	// make函数构造切片，make([]type, size, cap)
	sli4 := make([]int, 2)
	sli5 := make([]int, 2, 10)
	fmt.Println(sli4, len(sli4), cap(sli4)) // [0 0] 2 2
	fmt.Println(sli5, len(sli5), cap(sli5)) // [0 0] 2 10

	// 使用类似数组的偷懒方法构造切片
	sli6 := []int{4: 2}
	fmt.Println(sli6, len(sli6), cap(sli6)) // [0 0 0 0 2] 5 5

	// 可以向切片中 append 元素
	sli7 := []int{1}
	// 追加一个元素
	sli7 = append(sli7, 2)
	// 追加多个元素
	sli7 = append(sli7, 3, 4)
	// 追加一个切片，...表示解包，不能省略
	sli7 = append(sli7, []int{7, 8}...)
	// 在第一个位置插入元素
	sli7 = append([]int{0}, sli7...)
	// 在中间插入一个切片
	sli7 = append(sli7[:5], append([]int{5, 6}, sli7[5:]...)...)
	fmt.Println(sli7) // [0 1 2 3 4 5 6 7 8]

	// 对切片进行切片
	nums := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sli8 := nums[4:6:8]
	fmt.Printf("sli8 为 %d，其长度为 %d\n", sli8, len(sli8)) // [5 6] 2
	sli8 = sli8[:cap(sli8)]
	fmt.Printf("sli8 变为 %d，其第 4 个元素为 %d\n", sli8, sli8[3]) // [5 6 7 8] 8
}
