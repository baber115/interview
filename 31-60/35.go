package main

import "fmt"

// 以下输出什么
/**
答案：
2
传的指针
*/
func incr(p *int) int {
	*p++
	return *p
}

func main() {
	p := 1
	incr(&p)
	fmt.Println(p)
}
