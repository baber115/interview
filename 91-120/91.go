package main

import "fmt"

/*
下面代码输出什么？
答案：0 1 2
*/
func main() {
	x := []string{"a", "b", "c"}
	for v := range x {
		fmt.Print(v)
	}
}
