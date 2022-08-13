package main

import "fmt"

/*
下面的代码输出什么，请说明？
答案：1 2
*/
func main() {
	defer func() {
		fmt.Print(recover())
	}()
	defer func() {
		defer func() {
			fmt.Print(recover())
		}()
		panic(1)
	}()
	defer recover()
	panic(2)
}
