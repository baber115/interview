package main

import "fmt"

/**
发生panic时，会倒序执行defer（先进后出原则）
输出：
打印后
打印中
打印前
panic: 触发异常
*/
func main() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}
