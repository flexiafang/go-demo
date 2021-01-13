package main

/*
go.mod 文件
module：模块的引用路径
go：项目使用的 go 版本
require：项目所需的直接依赖包及其版本
exclude：忽略指定版本的依赖包
replace：由于在国内访问golang.org/x的各个包都需要翻墙，可以在go.mod中使用replace替换成github上对应的库。

go.sum 文件
<module> <version> <hash>
<module> <version>/go.mod <hash>
每一行都是由 模块路径，模块版本，哈希检验值 组成，其中哈希检验值是用来保证当前缓存的模块不会被篡改。
hash 是以h1:开头的字符串，表示生成checksum的算法是第一版的hash算法
*/
func main() {

}
