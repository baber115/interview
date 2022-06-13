package main

import "fmt"

/**
下面代码输出什么？
A. 13 23
B. compilation error
答：A

解析：

知识点：类型断言。
*/
type A45 interface {
	ShowA45() int
}

type B45 interface {
	ShowB45() int
}

type Work45 struct {
	i int
}

func (w Work45) ShowA45() int {
	return w.i + 10
}

func (w Work45) ShowB45() int {
	return w.i + 20
}

func main() {
	var a A45 = Work45{3}
	s := a.(Work45)
	fmt.Println(s.ShowA45())
	fmt.Println(s.ShowB45())
}
