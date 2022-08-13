package main

import "fmt"

/*
下面的代码输出什么？
答案：13 13 13
解析：赋值的是指针
*/

type N int

func (n *N) test() {
	fmt.Println(*n)
}

func main() {
	var n N = 10
	p := &n

	n++
	f1 := n.test

	n++
	f2 := p.test

	n++
	fmt.Println(n)

	f1()
	f2()
}
