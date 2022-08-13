package main

import "fmt"

/*
下面代码输出什么
答：
1 2 0 0 0
1 2 3 0 0
*/
func main() {
	slice := make([]int, 5, 5)
	slice[0] = 1
	slice[1] = 2
	change(slice...)
	fmt.Println(slice)
	change(slice[0:2]...)
	fmt.Println(slice)
}

func change(s ...int) {
	s = append(s, 3)
}
