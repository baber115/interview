package main

import "fmt"

/**
以下代码输出什么

答案：0112

解析：
iota是常量计数器，只能在const里使用
每个const都会重置iota，但是const中每一行，无论是否用iota都会计一次
第二个const,name占了一行，所以c是1
*/
const (
	a = iota
	b = iota
)

const (
	name = "name"
	c    = iota
	d    = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
