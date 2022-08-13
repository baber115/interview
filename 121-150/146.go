package main

import "fmt"

/*
下面代码输出什么？
答案：9 3 2
*/
func main() {
	var k = 9
	for k = range []int{} {
	}
	fmt.Println(k)

	for k = 0; k < 3; k++ {
	}
	fmt.Println(k)

	for k = range (*[3]int)(nil) {
	}
	fmt.Println(k)
}
