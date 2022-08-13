package main

import "fmt"

/*
下面的代码输出什么？
答案：运行错误
解析：a没有0下标，先计算左边，再计算右边，再赋值
*/
func main() {
	var a []int = nil
	a, a[0] = []int{1, 2}, 9
	fmt.Println(a)
}
