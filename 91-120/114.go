package main

import "fmt"

/*
下面的代码有什么问题？
答案：编译通过，常量可以不被使用
*/
func main() {
	const x = 123
	const y = 1.23
	fmt.Println(x)
}
