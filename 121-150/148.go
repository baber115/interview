package main

import "fmt"

/*
下面代码输出什么？
答案：-128
解析：溢出
*/
func main() {
	var x int8 = -128
	var y = x / -1
	fmt.Println(y)
}
