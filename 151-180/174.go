package main

import "fmt"

/*
下面代码输出什么，为什么？
答案：recover:1
解析：程序遇到panic就不会再往下执行，可以通过recover捕获panic的信息
*/
func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recover:%#v", r)
		}
	}()
	panic(1)
	panic(2)
}

func main() {
	f()
}
