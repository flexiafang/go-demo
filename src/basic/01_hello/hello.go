// 文件所属的包，Go语言中主函数所在的包一定是main
package main

// 导入标准输入输出包
import "fmt"

// func表示函数，Hello是函数名
func hello(str string) {
	fmt.Println(str)
	println(str)
}

// 主函数，程序有且只有一个主函数入口
func main() {
	var str string = "hello, world!"
	hello(str)
}
