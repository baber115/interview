package main

import "fmt"

/*
下面代码输出什么？
答案：
[0 11 12]
[21 12 13]
*/
func main() {
	a := [3]int{0, 1, 2}
	s := a[1:2]

	s[0] = 11
	s = append(s, 12)
	s = append(s, 13)
	s[0] = 21

	fmt.Println(a)
	fmt.Println(s)
}
