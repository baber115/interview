package main

import "fmt"

// 以下代码能否通过编译

/**
不能通过
此题考察类型定义和类型别名
*/
type MyInt1 int // 创建了新的MyInt1类型

type MyInt2 = int // 定义类型别名MyInt2

func main() {
	var i int = 0
	var i1 MyInt1 = i // 这里含义是将int类型的变量i，赋值给MyInt类型，GO是强类型语言，不能这么赋值
	//var i1 MyInt1 = MyInt1(i) 上面的代码可以用这种方式，强制转换变量i的类型
	var i2 MyInt2 = i // 这里的MyInt2只是int类型的别名，所以同为int类型的i变量，可以赋值
	fmt.Println(i1, i2)
}
