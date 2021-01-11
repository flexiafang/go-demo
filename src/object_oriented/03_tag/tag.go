package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	myFunc1()
	myFunc2()
}

type Person1 struct {
	Name string
	Age  int
	Addr string
}

/*
结构体每个字段都由名字和字段类型组成。
字段上还可以额外再加一个属性，用反引号 `` 包含的字符串，称之为 Tag，也就是标签。
Tag 由反引号包含，由一对或几对的键值对组成，通过空格来分割键值。
`key01:"value01" key02:"value02" key03:"value03"`
*/

type Person2 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Addr string `json:"addr,omitempty"`
}

/*
使用 encoding/json 抛砖引玉
字段有 omitempty 属性，因此 encoding/json 在将对象转化 json 字符串时，只要发现对象里的
字段值 为 false、0、空指针、空接口、空数组、空切片、空映射、空字符串中的一种，就会被忽略。
*/
func myFunc1() {
	p1 := Person1{
		Name: "Jack",
		Age:  20,
	}

	fmt.Println(p1)

	data1, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", data1)

	p2 := Person2{
		Name: "Jack",
		Age:  20,
	}

	fmt.Println(p2)

	data2, err := json.Marshal(p2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", data2)

	/*
		{Jack 20 }
		{"Name":"Jack","Age":20,"Addr":""}
		{Jack 20 }
		{"name":"Jack","age":20}
	*/
}

/*
从结构体中，获取 Tag 可以分为三个步骤：
  1. 获取字段 field
     field := reflect.TypeOf(obj).FieldByName("Name")
     field := reflect.ValueOf(obj).Type().Field(i)  // i 表示第几个字段
     field := reflect.ValueOf(&obj).Elem().Type().Field(i)  // i 表示第几个字段
  2. 获取标签 tag
     tag := field.Tag
  3. 获取键值对 key:value
     labelValue := tag.Get("label")
     labelValue,ok := tag.Lookup("label")
*/

type Person struct {
	Name   string `label:"Name is: "`
	Age    int    `label:"Age is: "`
	Gender string `label:"Gender is: " default:"unknown"`
}

// 打印对象
func Print(obj interface{}) {
	// 取 value
	v := reflect.ValueOf(obj)

	// 解析字段
	for i := 0; i < v.NumField(); i++ {

		// 取 Tag
		field := v.Type().Field(i)
		tag := field.Tag

		// 解析 label 和 default
		label := tag.Get("label")
		defaultValue := tag.Get("default")

		// Sprintf() 是把格式化字符串输出到指定的字符串中，可以用一个变量来接受，然后打印
		// %v 表示按值的本来值输出
		value := fmt.Sprintf("%v", v.Field(i))
		if value == "" {
			// 如果没有指定值，则使用默认值代替
			value = defaultValue
		}

		fmt.Println(label + value)
	}
}

func myFunc2() {
	person := Person{
		Name: "flexia",
		Age:  22,
	}

	Print(person)

	/*
		Name is: flexia
		Age is: 22
		Gender is: unknown
	*/
}
