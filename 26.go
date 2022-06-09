package main

import "fmt"

// 以下代码输出什么

/**
答案：nil
当接口动态值和动态类型都为nil时，接口才为nil
*/
func main() {
	var i interface{}
	if i == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}
