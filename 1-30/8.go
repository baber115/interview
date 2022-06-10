package main

import "fmt"

// 以下代码能否编译通过

/**
不能
:=是简短的定义模式
1、只能在函数内部使用
2、不能指定数据类型，编译的时候会自动推导

*/
var (
	size     = 1024
	max_size = size * 2
)

func main() {
	fmt.Println(size)
	fmt.Println(max_size)
}
