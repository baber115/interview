package main

import "fmt"

/*
下面这段代码输出什么？
答案：100
解析：虽然foo的函数是传值，但是t.ls[0]，是给底层数组赋值
*/
type T struct {
	ls []int
}

func foo(t T) {
	t.ls[0] = 100
}

func main() {
	var t = T{
		ls: []int{1, 2, 3},
	}

	foo(t)
	fmt.Println(t.ls[0])
}
