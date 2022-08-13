package main

import "fmt"

/*
下面的代码输出什么？
答案：9 [{0} {9}]
解析：for range循环中，循环变量t是原数组的副本，如果数组元素是结构体，则副本字段和原数组字段是两个不同的值
*/
type T struct {
	n int
}

func main() {
	ts := [2]T{}
	for i, t := range &ts {
		switch i {
		case 0:
			t.n = 3
			ts[1].n = 9
		case 1:
			fmt.Print(t.n, " ")
		}
	}
	fmt.Print(ts)
}
