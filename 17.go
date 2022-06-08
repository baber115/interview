package main

import "fmt"

// 以下代码输出什么
/**
输出not nil
是将hello()这个函数给h变量，不是返回值给h，所以是not nil
*/
func hello() []string {
	return nil
}

func main() {
	h := hello
	if h == nil {
		fmt.Println(" nil")
	} else {
		fmt.Println("not nil")
	}
}
