package main

import (
	"fmt"
	"reflect"
)

/*
反射：获取一个对象的类型，属性及方法
两种重要的类型：reflect.Type 接口 和 reflect.Value 结构体

反射三大定律
  1. 反射可以将 “接口类型变量” 转换为 “反射类型对象”；
  2. 反射可以将 “反射类型对象” 转换为 “接口类型变量”；
  3. 如果要修改 “反射类型对象” 其类型必须是 可写的；
*/

func main() {
	myFunc1()
	myFunc2()
	myFunc3()
	myFunc4()
	myFunc5()
	myFunc6()
	myFunc7()
	myFunc8()
	myFunc9()
	myFunc10()
}

/*
接口变量 <-> 反射对象
TypeOf() / ValueOf() / Interface()

Interface() 得到的是静态类型为 interface{} 的变量，要转换为原始类型需要使用 类型断言 再转换一次
*/
func myFunc1() {
	var age interface{} = 22
	fmt.Printf("原始接口变量的类型是 %T，值为 %v\n", age, age) // int 22

	t := reflect.TypeOf(age)
	v := reflect.ValueOf(age)
	fmt.Printf("接口变量 -> 反射对象: Type 的类型是 %T\n", t)  // *reflect.rtype
	fmt.Printf("接口变量 -> 反射对象: Value 的类型是 %T\n", v) // reflect.Value

	i := v.Interface()
	fmt.Printf("反射对象 -> 接口变量: 新对象的类型为 %T，值为 %v\n", i, i) // int 22

	j := v.Interface().(int)
	fmt.Printf("接口变量 -> 原始类型: 新变量的类型为 %T，值为 %v\n", j, j) // int 22
}

/*
Go 语言中的函数都是值传递，只要传递的不是变量指针，函数内部对变量的修改不会影响原始变量的值。
回到反射，当使用 ValueOf、TypeOf 的时候，如果不是传递的接口变量指针，反射得到的变量值始终只是真实值的一个拷贝。

因此在反射的规则里：
  - 不是接收变量指针创建的反射对象，是不具备『可写性』的
  - 是否具备『可写性』，可使用 CanSet() 来获取得知
  - 对不具备『可写性』的对象进行修改，是没有意义的，也认为是不合法的，因此会报错。

让反射对象具有可写性需要注意：
  1. 创建反射对象时传入变量指针
  2. 使用 Elem() 函数返回指针指向的数据

修改反射对象的值，使用以 Set 开头的几个方法： Set / SetBool / SetInt / SetString 等等
*/
func myFunc2() {
	var name string = "go"
	fmt.Println("name 的原始值为：", name) // go

	v1 := reflect.ValueOf(name)
	fmt.Println("v1 的可写性为：", v1.CanSet()) // false

	v2 := reflect.ValueOf(&name)
	fmt.Println("v2 的可写性为：", v2.CanSet()) // false

	v3 := v2.Elem()
	fmt.Println("v3 的可写性为：", v3.CanSet()) // true

	v3.SetString("go here")
	fmt.Println("通过反射对象修改后，name 的值变为：", name) // go here
}

type Person1 struct {
	name   string
	age    int
	gender string
}

func (p Person1) SayBye() {
	fmt.Println("Bye")
}

func (p Person1) SayHello() {
	fmt.Println("Hello")
}

/*
Kind() 获取类别
  注意与 TypeOf() 的区别
*/
func myFunc3() {
	m := Person1{}

	t1 := reflect.TypeOf(m)
	fmt.Println("t1 Type:", t1)        // main.Person1
	fmt.Println("t1 Kind:", t1.Kind()) // struct

	t2 := reflect.TypeOf(&m)
	fmt.Println("t2 Type:", t2)        // *main.Person1
	fmt.Println("t2 Kind:", t2.Kind()) // ptr

	v := reflect.ValueOf(&m)
	fmt.Println("&m Type: ", v.Type())       // *main.Person1
	fmt.Println("&m Kind: ", v.Kind())       // ptr
	fmt.Println("m Type: ", v.Elem().Type()) // main.Person1
	fmt.Println("m Kind: ", v.Elem().Kind()) // struct
}

/*
类型转换
  Int()
  Float()
  String()
  Bool()
  Pointer()
  Interface()
*/
func myFunc4() {
	var age int = 25

	v1 := reflect.ValueOf(age)
	fmt.Printf("转换前，type：%T，value：%v\n", v1, v1) // reflect.Value 25

	v2 := v1.Int()
	fmt.Printf("转换后，type：%T，value：%v\n", v2, v2) // int64 25
}

/*
对切片的操作
  Slice() 对切片再切片（2下标）
  Slice3() 对切片再切片（3下标）
  Set()、Append() 更新切片
*/
func myFunc5() {
	var nums []int = []int{1, 2, 3, 4, 5}

	v1 := reflect.ValueOf(nums)
	fmt.Printf("切片前，type：%T，value：%v\n", v1, v1) // reflect.Value [1 2 3 4 5]

	v2 := v1.Slice(0, 2)
	fmt.Printf("切片后，type：%T，value：%v\n", v2, v2) // reflect.Value [1 2]

	v3 := v2.Slice3(2, 5, 5)
	fmt.Printf("切片后，type：%T，value：%v\n", v3, v3) // reflect.Value [3 4 5]

	appendToSlice(&nums)
	fmt.Println(nums) // [1 2 3 4 5 3]
}

func appendToSlice(arrPtr interface{}) {
	valuePtr := reflect.ValueOf(arrPtr)
	value := valuePtr.Elem()

	value.Set(reflect.Append(value, reflect.ValueOf(3)))

	fmt.Println(value)       // [1 2 3 4 5 3]
	fmt.Println(value.Len()) // 6
}

/*
对属性的操作
  NumField()
  Field(i)
*/
func myFunc6() {
	p := Person1{"flexia", 22, "male"}
	v := reflect.ValueOf(p)

	fmt.Println("字段数：", v.NumField())
	for i := 0; i < v.NumField(); i++ {
		fmt.Println("第", i, "个字段：", v.Field(i))
	}

	/*
		字段数： 3
		第 0 个字段： flexia
		第 1 个字段： 22
		第 2 个字段： male
	*/
}

/*
对方法的操作
  NumMethod()
  Method(i)
*/
func myFunc7() {
	p := &Person1{"flexia", 22, "male"}

	t := reflect.TypeOf(p)

	fmt.Println("方法数（可导出的）：", t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println("第", i, "个方法：", t.Method(i).Name)
	}

	/*
		方法数（可导出的）： 2
		第 0 个方法： SayBye
		第 1 个方法： SayHello
	*/
}

type Person2 struct {
	name string
	age  int
}

func (p Person2) SayHello() string {
	return "hello"
}

func (p Person2) SayBye() string {
	return "bye"
}

/*
动态调用函数（使用索引且无参数）

要调用 Call，注意要使用 ValueOf
*/
func myFunc8() {
	p := &Person2{"flexia", 22}

	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)

	for i := 0; i < v.NumMethod(); i++ {
		fmt.Println("调用第", i+1, "个方法：", t.Method(i).Name,
			"，调用结果：", v.Elem().Method(i).Call(nil))
	}

	/*
		调用第 1 个方法： SayBye ，调用结果： [bye]
		调用第 2 个方法： SayHello ，调用结果： [hello]
	*/
}

/*
动态调用函数（使用函数名且无参数）
*/
func myFunc9() {
	p := &Person1{"flexia", 22, "male"}

	v := reflect.ValueOf(p)

	v.MethodByName("SayBye").Call(nil)
	v.MethodByName("SayHello").Call(nil)

	/*
		Bye
		Hello
	*/
}

type Person3 struct {
}

func (p Person3) SelfIntroduction(name string, age int) {
	fmt.Printf("Hello, my name is %s and I'm %d years old\n", name, age)
}

/*
动态调用函数（使用函数且有参数）
*/
func myFunc10() {
	p := &Person3{}

	v := reflect.ValueOf(p)
	name := reflect.ValueOf("flexia")
	age := reflect.ValueOf(22)
	params := []reflect.Value{name, age}

	v.MethodByName("SelfIntroduction").Call(params) // Hello, my name is flexia and I'm 22 years old
}
