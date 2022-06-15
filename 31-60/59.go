package main

import "fmt"

/**
下面这段代码输出什么？为什么？
A. 1
B. compilation error

答案: B

基于类型创建的方法定义，必须在同一个包内，上面的代码基于int创建了PrintInt方法，但是int和PrintInt方法不在同一个包内，所以编译出错

正确答案

type Myint int
func (i Myint) PrintInt ()  {
    fmt.Println(i)
}
func main() {
    var i Myint = 1
    i.PrintInt()
}
*/
func (i int) PrintInt() {
	fmt.Println(i)
}

func main() {
	var i int = 1
	i.PrintInt()
}
