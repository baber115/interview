package main

import "fmt"

// 下面这段代码输出什么？

/**
答案：
18
hello21函数改变了第一个值
*/
func hello21(num ...int) {
	num[0] = 18
}

func main() {
	i := []int{5, 6, 7}
	hello21(i...)
	fmt.Println(i[0])
}
