package main

import "fmt"

// 以下代码输出什么
/**
答案：
编译失败

GO语言中，字符串是只读的
*/
func main() {
	str := "hello"
	str[0] = 'x'
	fmt.Println(str)
}
