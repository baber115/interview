package main

import "fmt"

/*
下面代码是否能编译通过？如果通过，输出什么？

答案：non-empty interface
解析：接口包含静态类型、动态类型和动态值，当且仅当动态值和动态类型都为nil时，接口才为nil
这里的x的动态类型为*int，所以不是nil
*/

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}
func main() {
	var x *int = nil
	Foo(x)
}
