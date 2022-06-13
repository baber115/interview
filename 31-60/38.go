package main

import "fmt"

/**
下面这段代码输出什么？
A. A
B. 65
C. compilation error

答案A，UTF-8编码中，十进制65表示A，但是不推荐这么写
*/
func main() {
	//i := 65
	var i rune = 65
	fmt.Println(string(i))
	//fmt.Println(string(rune(i))) 正确的写法一
	// var i rune = 65
	// var i unit8 = 65
	// var i byte = 65
}
