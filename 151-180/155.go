package main

import "fmt"

/*
下面的代码输出什么？
答案：0 [{0} {9}]
解析：for range循环使用的是数组ts的副本，所以t.n = 3的操作不会影响原数组
*/
type T struct {
	n int
}

func main() {
	ts := [2]T{}
	for i, t := range ts {
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
