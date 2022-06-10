package main

import "fmt"

const (
	x = iota // 初始值是0，下面依次+1
	_        // 1，但是不赋值
	y        // 2
	z = "zz" // zz
	k        // zz
	p = iota // 5，第5个元素
)

func main() {
	fmt.Println(x, y, z, k, p)
}
