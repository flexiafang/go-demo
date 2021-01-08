package main

import "fmt"

// 结构体 struct

// 定义一个结构体
type Profile struct {
	name   string
	age    int
	gender string
	mother *Profile // 指针
	father *Profile // 指针
}

/*
无法在结构体内部定义方法，可以使用组合方法的方式来定义结构体方法。
  - FmtProfile: 方法名。
  - (person Profile): 表示将 FmtProfile 方法与 Profile 的实例 person 绑定。
*/
func (person Profile) FmtProfile() {
	fmt.Printf("姓名：%s\n", person.name)
	fmt.Printf("年龄：%d\n", person.age)
	fmt.Printf("性别：%s\n", person.gender)
}

/*
两种定义方法的方式：
  1. 以值作为方法接收者。
  2. 以指针作为方法接收者。
     (1) 需要在方法内部改变结构体内容。
     (2) 结构体过大。

在 Go 语言中，方法名的首字母大小写非常重要，它被来实现控制对方法的访问权限。
  - 当方法的首字母为大写时，这个方法对于所有包都是Public，其他包可以随意调用。
  - 当方法的首字母为小写时，这个方法是Private，其他包是无法访问的。
*/
func (person *Profile) increaseAge() {
	person.age += 1
}

/*
使用组合实现类似继承的效果。
在 Go 语言中，把一个结构体嵌入到另一个结构体的方法，称之为组合。
可以将某个结构体嵌入到另一个结构体中作为一个匿名字段，另一个结构体就直接拥有了此结构体的所有属性了。
*/
type company struct {
	companyName string
	companyAddr string
}

type staff struct {
	name     string
	age      int
	gender   string
	position string
	company
}

func (staffInfo *staff) fmtCompanyName() {
	// staffInfo.companyName 和 staffInfo.company.companyName 的效果是一样的
	fmt.Printf("%s 在 %s 工作\n", staffInfo.name, staffInfo.companyName)
	fmt.Printf("%s 在 %s 工作\n", staffInfo.name, staffInfo.company.companyName)
}

func main() {
	// 实例化
	myself := Profile{name: "flexia", age: 22, gender: "male"}
	// 调用方法
	myself.FmtProfile()

	myself.increaseAge()
	myself.FmtProfile()

	myCompany := company{
		companyName: "Tencent",
		companyAddr: "深圳市南山区",
	}

	staffInfo := staff{
		name:     "flexia",
		age:      22,
		gender:   "男",
		position: "后台开发工程师",
		company:  myCompany,
	}

	staffInfo.fmtCompanyName()

	/*
		姓名：flexia
		年龄：22
		性别：male
		姓名：flexia
		年龄：23
		性别：male
		flexia 在 Tencent 工作
		flexia 在 Tencent 工作
	*/
}
