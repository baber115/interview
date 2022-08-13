package main

import "fmt"

/*
下面代码输出什么？
答案：0 1 1 2
*/
const (
	azero = iota
	aone  = iota
)

const (
	info  = "msg"
	bzero = iota
	bone  = iota
)

func main() {
	fmt.Println(azero, aone)
	fmt.Println(bzero, bone)
}
