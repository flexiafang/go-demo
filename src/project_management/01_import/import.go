package main

/*
关于包导入的8个知识点：

1. 单行导入与多行导入
	import "fmt"
	import "sync"
	或
	import (
		"fmt"
		"sync"
	)

2. 使用别名

(1) 导入具有同一包名的包时避免冲突
	import (
		"crypto/rand"
		mrand "math/rand"
	)

(2) 导入的包名过长
	import hw "helloworldtestmodule"

(3) 防止导入的包名与本地变量产生冲突
	import pathpkg "path"

3. 频繁使用的包使用点操作
	import . "fmt"
	func main() {
		Println("hello, world!")
	}

4. 包的初始化
	每个包都允许有一个 init 函数，当这个包被导入时，会执行该包的 init 函数，做一些初始化工作。
	需要注意，init 函数优先于 main 函数执行，在一个包的引用连中，包的初始化是深度优先的。

5. 包的匿名导入
	当导入的包没有被使用时，编译时会报错。
	使用下划线 _ 进行匿名导入，会执行包里的 init 函数，编译时会将这个包编译到可执行文件中，但是并不能被访问。
	import _ "image/png"

6. 导入的是路径还是包？
	导入时，是按照目录进行导入的，导入目录后，可以使用这个目录下的所有包。

7. 相对导入和绝对导入
	绝对导入：从 $GOPATH/src 或 $GOROOT 或者 $GOPATH/pkg/mod 目录下搜索包并导入。
	相对导入：从当前目录中搜索包并开始导入。
	使用相对导入的方式，项目可读性会大打折扣，不利用开发者理清整个引用关系。所以一般更推荐使用绝对引用的方式。

8. 包导入路径优先级

	(1) govendor 模式
		vendor -> $GOROOT/src -> $GOPATH/src

	(2) go modules 模式
		导入的包有域名：$GOPATH/pkg/mod -> 联网
		导入的包没有域名：$GOROOT
		项目下有 vendor 目录时，只会在 vendor 目录下查找。
*/

import (
	. "fmt"
)

func main() {
	Println("hello, world!")
}
