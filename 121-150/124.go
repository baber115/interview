package main

import "fmt"

/*
下面代码输出什么？
答案：false true
解析：switch默认没有break，或者case完成程序默认会break，可以在case中加上fallthrough，这样会继续走后面的case
*/
func main() {
	isMatch := func(i int) bool {
		switch i {
		case 1:
		case 2:
			return true
		}
		return false
	}

	fmt.Println(isMatch(1))
	fmt.Println(isMatch(2))
}
