package main

import "fmt"

/*
下面的代码输出什么？
答案：9 [{0} {9}]

*/
type T struct {
	n int
}

func main() {
	ts := [2]T{}
	for i := range ts[:] {
		switch i {
		case 0:
			ts[1].n = 9
		case 1:
			fmt.Print(ts[i].n, " ")
		}
	}
	fmt.Print(ts)
}
