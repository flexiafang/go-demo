package main

import (
	"fmt"
	"unsafe"
)

// Go 语言中的数据类型
func main() {

	/*
		1. 整型
		10个类型的整数：
		- 有符号（int, int8, int16, int32, int64）
		- 无符号（uint, uint8, uint16, uint32, uint64）
		不同进制的表示：
		- 10进制
		- 2进制 0B/0b
		- 8进制 0O/0o
		- 16进制 0X/0x
	*/
	var int01 = 0b1100
	var int02 = 0o14
	var int03 = 0xC
	fmt.Printf("2  进制数 %b 表示的是：%d\n", int01, int01)
	fmt.Printf("8  进制数 %o 表示的是：%d\n", int02, int02)
	fmt.Printf("16 进制数 %x 表示的是：%d\n", int03, int03)

	/*
		2. 浮点型
		两种精度的浮点数类型
		- float32: 1位表示符号，8位表示指数，23位表示尾数，精确到小数点后6位
		- float64: 1位表示符号，11位表示指数，52位表示尾数，精确到小数点后15位
	*/
	var float01 float32 = 10000018
	fmt.Println("float01：  ", float01)   // 1.0000018e+07
	fmt.Println("float01+1：", float01+1) // 1.0000019e+07

	var float02 float32 = 100000182
	var float03 float32 = 100000187
	fmt.Println("float02+5：", float02+5) // 1.0000019e+08
	fmt.Println("float03：  ", float03)   // 1.00000184e+08
	fmt.Println(float02+5 == float03)    // false

	/*
		3. byte 与 rune
		- byte: 占用一个字节，与uint8本质上没有区别，表示的是ASCII表中的一个字符
		- rune:	占用4个字节，共32位，与uint32本质上没有区别，表示的是一个Unicode字符
	*/
	var a byte = 'A'
	var b uint8 = 'B'
	fmt.Printf("a 的值：%c\nb 的值：%c\n", a, b)

	var c byte = 'A'
	var d rune = 'B'
	fmt.Printf("c 占用 %d 个字节数\nd 占用 %d 个字节数\n", unsafe.Sizeof(c), unsafe.Sizeof(d)) // 1 4

	/*
		4. 字符串
		多个字符组成字符串，string 本质是一个 byte 数组
		Go 语言的 string 使用 utf-8 进行编码，英文字母占用一个字节，中文字符占用3个字节
		除了双引号 "" 外，还可以使用反引号 ``，会忽略转义，所见即所得
	*/
	var str01 string = "hello"
	var str02 [5]byte = [5]byte{104, 101, 108, 108, 111}
	fmt.Printf("str01: %s\nstr02: %s\n", str01, str02) // hello hello

	var str03 = "hello,中国"
	fmt.Printf("str03 占用 %d 个字节\n", len(str03)) // 12

	var str04 = "\\r\\n"
	var str05 = `\r\n`
	fmt.Println("str04: ", str04)
	fmt.Println("str05: ", str05)
	fmt.Printf("%s 的解释型字符串是 %q\n", str05, str05) // "\r\n" "\\r\\n"
}
