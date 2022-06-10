package main

import "fmt"

// 以下代码能否通过编译
func main() {
	/**
	不能
	new初始化返回的是指针类型，append不能对指针进行操作，建议使用make
	*/
	list := new([]int)
	list = append(list, 1)
	fmt.Println(list)
}
