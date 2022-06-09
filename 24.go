package main

import "fmt"

// 以下代码输出什么

/**
答案:编译报错
GO中数组是一种值类型,可以比较。但是数组长度是数组的一部分，不同长度的数组不能相互比较
*/
func main() {
	a := [2]int{5, 6}
	b := [3]int{5, 6}
	if a == b {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}
