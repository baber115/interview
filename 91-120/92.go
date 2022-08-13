package main

import "fmt"

/*
下面这段代码能否编译通过？如果通过，输出什么？

答案：不能通过

解析：
User1是基于User创建的新类型
User2是User的别名
所以User2具有User的所有方法，但是User1做为新类型，没有m1(),m2()方法，i1.m1()会调用失败
*/

type User struct{}
type User1 User
type User2 = User

func (i User) m1() {
	fmt.Println("m1")
}
func (i User) m2() {
	fmt.Println("m2")
}

func main() {
	var i1 User1
	var i2 User2
	i1.m1()
	i2.m2()
}
