package main

import "fmt"

/*
下面的代码输出什么？
答案：10
*/
type T struct {
	x int
	y *int
}

func main() {
	i := 20
	t := T{10, &i}

	p := &t.x

	*p++
	*p--

	t.y = p

	fmt.Println(*t.y)
}
