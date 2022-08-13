package main

import "fmt"

/*
下面这段代码输出什么？
*/
func main() {
	count := 0
	for i := range [256]struct{}{} {
		m, n := byte(i), int8(i)
		if n == -n {
			count++
		}
		if m == -m {
			count++
		}
	}
	fmt.Println(count)
}
