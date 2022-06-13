package main

import "fmt"

/**
下面代码输出什么？
答案：报错

*/
type a interface {
	showA() int
}

type b interface {
	showB() int
}

type work struct {
	i int
}

func (w work) showA() int {
	return w.i + 10
}

func (w work) showB() int {
	return w.i + 20
}

func main() {
	c := Work{3}
	var a A = c
	var b B = c
	// a接口实现的是showA，不是showB
	fmt.Println(a.showB())
	fmt.Println(b.showA())
}
