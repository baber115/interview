package main

import "fmt"

/*
下面这段代码输出什么？
答案：4
解析：多重赋值
1、先计算左边的表达式，再计算右边的表达式
2、赋值
*/
func main() {
	var k = 1
	var s = []int{1, 2}
	k, s[k] = 0, 3
	fmt.Println(s[0] + s[1])
}
