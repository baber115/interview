package main

import "fmt"

/*
下面的代码有什么问题？
答案：编译错误
解析：不能使用多级指针调用方法
*/
type N int

func (n N) value() {
	n++
	fmt.Printf("v:%p,%v\n", &n, n)
}

func (n *N) pointer() {
	*n++
	fmt.Printf("v:%p,%v\n", n, *n)
}

func main() {

	var a N = 25

	p := &a
	p1 := &p

	p1.value()
	p1.pointer()
}
