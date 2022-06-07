package main

import "fmt"

// 以下代码能否通过编译
func main() {
	/**
	不能
	append第2个参数，不能直接操作切片，需要加上...
		list = append(list1, list2...)
	*/
	list1 := []int{1, 2, 3}
	list2 := []int{4, 5}
	list1 = append(list1, list2)
	fmt.Println(list1)
}
