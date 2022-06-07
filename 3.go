package main

import "fmt"

func main() {
	/*
		1、make的时候设置了len=5，会有5个0做为占位符
		2、append在末尾增加数据，不影响前面
		输出：000001234
	*/
	fmt.Println("s1")
	s1 := make([]int, 5)
	s1 = append(s1, 1, 2, 3, 4)
	fmt.Println(s1)

	// 输出：1234
	fmt.Println("s2")
	s2 := make([]int, 0)
	s2 = append(s2, 1, 2, 3, 4)
	fmt.Println(s2)
}
