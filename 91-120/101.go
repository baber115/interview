package main

/*
下面这段代码输出什么？
答案：2
*/
import "fmt"

func incr(p *int) int {
	*p++
	return *p
}
func main() {
	v := 1
	incr(&v)
	fmt.Println(v)
}
