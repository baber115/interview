package main

import "fmt"

/*
下面代码最后一行输出什么？请说明原因。
答案：1
解析：作用域，func里的作用域，不影响外面的
*/

func main() {
	x := 1
	fmt.Println(x)
	{
		fmt.Println(x)
		i, x := 2, 2
		fmt.Println(i, x)
	}
	fmt.Println(x) // print ?
}
