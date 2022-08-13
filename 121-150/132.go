package main

import "fmt"

/*
下面的代码输出什么？
答案：10 11 12
解析：
*/
type N int

func (n N) test() {
	fmt.Println(n)
}

func main() {
	var n N = 10
	fmt.Println(n)

	n++
	f1 := N.test
	f1(n)

	n++
	f2 := (*N).test
	f2(&n)
}
