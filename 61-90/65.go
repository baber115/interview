package main

import "fmt"

/**
下面代码有什么问题

一、不同类型不能比较
二、切片和map不能直接做比较
*/
func main() {
	fmt.Println([...]int{1} == [2]int{1})
	fmt.Println([]int{1} == []int{1})
}
