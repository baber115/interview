package main

import "fmt"

/**
下面这段代码正确的输出是什么？
A. F M D
B. D F M
C. F D M

答案：C
被调用函数里的defer，会在当被调用函数返回之前执行
*/
func f() {
	defer fmt.Println("D")
	fmt.Println("F")
}

func main() {
	f()
	fmt.Println("M")
}
