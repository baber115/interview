package main

import "fmt"

/**
return 之后的 defer 语句会执行吗，下面这段代码输出什么？

return 之后 的defer不会执行，所以答案是2，1
*/
var a54 bool = true

func main() {
	defer func() {
		fmt.Println("1")
	}()
	if a54 == true {
		fmt.Println("2")
		return
	}
	defer func() {
		fmt.Println("3")
	}()
}
