package main

import "fmt"

/*
下面代码输出什么，请说明？
答案：2 1
*/

func main() {
	defer func() {
		fmt.Print(recover())
	}()
	defer func() {
		defer fmt.Print(recover())
		panic(1)
	}()
	defer recover()
	panic(2)
}
