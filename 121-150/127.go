package main

/*
下面的代码有什么问题？
type X struct {}

func (x *X) test()  {
    println(x)
}

func main() {

    var a *X
    a.test()

    X{}.test()
}
答案：X{}是不可寻址的，不能直接调用方法
解析：知识点：在方法中，指针类型的接收者必须是合法指针（包括 nil）,或能获取实例地址。
*/

// 修改后
type X struct{}

func (x *X) test() {
	println(x)
}

func main() {
	var a *X
	a.test()
	x := X{}
	x.test()
}
