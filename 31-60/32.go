package main

import "fmt"

// 以下代码输出什么
/**
答案：5
运行到defer的时候，已经把5放进内存了，后面的加法不会影响内存里的，所以答案是5
*/
func mai() {
	i := 5
	defer hello(i)
	i = i + 10
}

func hello(i int) {
	fmt.Println(i)
}
