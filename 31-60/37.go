package main

import "fmt"

/**
下面代码下划线处可以填入哪个选项以输出yes nil？
A. s1
B. s2
C. s1、s2 都可以

答案：A
s1是nil切片
s2是空切片
*/
func main() {
	var s1 []int
	var s2 = []int{}
	if __ == nil {
		fmt.Println("yes nil")
	} else {
		fmt.Println("no nil")
	}
}
