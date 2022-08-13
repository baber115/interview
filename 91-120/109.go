package main

/*
下面代码有什么问题？
答案：nil代表interface、函数、map、slice和channel的”零值“，如果不指定类型，go语言不知道nil代表哪个类型
*/
func main() {
	var x = nil
	_ = x
}
