package main

import "fmt"

// 流程控制
func main() {

	/*
		1. if-else 条件语句
		由于 Go 是强类型，所以要求条件表达式必须严格返回布尔型的数据
		Go 编译器，对于 { 和 } 的位置有严格的要求，必须这样写 } else (if) {
	*/

	var age int
	age = 22
	if age > 18 {
		fmt.Println("已经成年了")
	} else if age > 12 {
		fmt.Println("已经是青少年了")
	} else {
		fmt.Println("还是个小孩子")
	}

	// 高级写法
	if score := 100; score >= 60 {
		fmt.Println("及格了")
	}

	/*
		2. switch-case 选择语句
		只要有一个 case 满足条件，就会执行对应的代码块，然后直接退出 switch-case
	*/

	// 一个 case 可接多个条件
	month := 1
	var monthStr string
	switch month {
	case 3, 4, 5:
		monthStr = "春天"
	case 6, 7, 8:
		monthStr = "夏天"
	case 9, 10, 11:
		monthStr = "秋天"
	case 12, 1, 2:
		monthStr = "冬天"
	default:
		monthStr = "输入有误..."
	}
	fmt.Println(monthStr)

	// switch 后可接函数
	chinese := 80
	english := 50
	math := 100
	switch getResult(chinese, english, math) {
	case true:
		fmt.Println("全科通过")
	case false:
		fmt.Println("有挂科记录")
	}

	// switch 可不接表达式，相当于 if-elseif-else
	var score int
	score = 100
	var scoreRank string
	switch {
	case score >= 95 && score <= 100:
		scoreRank = "优秀"
	case score >= 80:
		scoreRank = "良好"
	case score >= 60:
		scoreRank = "合格"
	case score >= 0:
		scoreRank = "不合格"
	default:
		scoreRank = "输入有误..."
	}
	fmt.Println(scoreRank)

	// switch 的穿透能力，使用 fallthrough 穿透一层
	var s string
	s = "hello"
	switch {
	case s == "hello":
		fmt.Println("hello")
		fallthrough
	case s == "hust":
		fmt.Println("hust")
	case s != "world":
		fmt.Println("world")
	}
}

func getResult(args ...int) bool {
	var result = true
	for _, i := range args {
		if i < 60 {
			result = false
			break
		}
	}
	return result
}
